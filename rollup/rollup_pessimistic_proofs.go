package rollup

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/sync/errgroup"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/l2-sovereign-chain/polygonpessimisticconsensus"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/l2-sovereign-chain/polygonzkevmbridgev2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/l2-sovereign-chain/polygonzkevmglobalexitrootv2"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
)

// https://github.com/0xPolygonHermez/zkevm-commonjs/blob/bb0e77e9158a0fc3d06eb5de53b458bb87f77bc7/src/constants.js#L58
var L2GERManager = common.HexToAddress("0xa40d5f56745a118d0906a34e69aec8c0db1cb8fa")

type RollupPessimisticProofs struct {
	*RollupMetadata

	Contract *polygonpessimisticconsensus.Polygonpessimisticconsensus `json:"-"`
}

func (r *RollupPessimisticProofs) InitContract(ctx context.Context, client bind.ContractBackend) error {
	if r.Contract != nil {
		return nil
	}

	contract, err := polygonpessimisticconsensus.NewPolygonpessimisticconsensus(r.Address, client)
	if err != nil {
		return err
	}

	r.Contract = contract

	return nil
}

// GetBatchL2Data constructs bridge initialization transaction.
// The logic is implemented according to the following snippet:
// https://github.com/0xPolygonHermez/zkevm-contracts/blob/v9.0.0-rc.3-pp/deployment/v2/4_createRollup.ts#L404-L456
func (r *RollupPessimisticProofs) GetBatchL2Data(client bind.ContractBackend) (string, error) {
	const (
		maxTxDataLength = 65535
		gasLimit        = uint64(30000000)
	)

	bridgeAddr, err := r.Contract.BridgeAddress(nil)
	if err != nil {
		return "", err
	}

	bridge, err := polygonzkevmbridgev2.NewPolygonzkevmbridgev2(bridgeAddr, client)
	if err != nil {
		return "", err
	}

	gasTokenNetwork, err := bridge.GasTokenNetwork(nil)
	if err != nil {
		return "", err
	}

	gasTokenMetadata := []byte{}
	if r.GasToken != common.HexToAddress("0x0") {
		gasTokenMetadata, err = bridge.GetTokenMetadata(nil, r.GasToken)
		if err != nil {
			return "", err
		}
	}

	bridgeABI, err := polygonzkevmbridgev2.Polygonzkevmbridgev2MetaData.GetAbi()
	if err != nil {
		return "", err
	}

	bridgeInitTxData, err := bridgeABI.Pack("initialize",
		r.RollupID,
		r.GasToken,
		gasTokenNetwork,
		L2GERManager,
		common.Address{}, // rollup manager does not exist on L2
		gasTokenMetadata)
	if err != nil {
		return "", err
	}

	if len(bridgeInitTxData) > maxTxDataLength {
		return "", fmt.Errorf("bridge init tx data length exceeds maximum allowed size (%d bytes)", maxTxDataLength)
	}

	V := big.NewInt(27)
	R := common.BigToHash(big.NewInt(0x5ca1ab1e0))
	S := common.BigToHash(big.NewInt(0x5ca1ab1e))

	bridgeInitTx := &preEIP155Transaction{
		To:       &bridgeAddr,
		Nonce:    0,
		GasPrice: common.Big0,
		Value:    common.Big0,
		GasLimit: gasLimit,
		Data:     bridgeInitTxData,
	}

	var rawBridgeInitTx bytes.Buffer
	err = rlp.Encode(&rawBridgeInitTx, bridgeInitTx)
	if err != nil {
		return "", err
	}

	rawTxWithSignature := append(rawBridgeInitTx.Bytes(), encodeTxSignature(V, R, S)...)
	effectivePercentage := big.NewInt(0xff).Bytes() // effective percentage is custom transaction field, related to the zkevm
	rawTxWithSignature = append(rawTxWithSignature, effectivePercentage...)

	return hexutil.Encode(rawTxWithSignature), nil
}

// GetRollupGlobalExitRoot retrieves the actual global exit root at the rollup creation time
func (r *RollupPessimisticProofs) GetRollupGlobalExitRoot(rm *rollupmanager.RollupManager, client bind.ContractBackend) (common.Hash, error) {
	gerContract, err := polygonzkevmglobalexitrootv2.NewPolygonzkevmglobalexitrootv2(rm.GERAddr, client)
	if err != nil {
		return common.Hash{}, err
	}

	if r.CreationBlock == 0 {
		return common.Hash{}, errors.New("rollup creation block number is zero")
	}

	endBlock := r.CreationBlock - 1

	if endBlock < rm.UpdateToULxLyBlock {
		return common.Hash{}, fmt.Errorf("end block (%d) is less than starting block (%d) in UpdateL1InfoTree filter",
			rm.UpdateToULxLyBlock, endBlock)
	}

	const blocksChunkSize = 100000
	var (
		l1InfoTreeEvents []*polygonzkevmglobalexitrootv2.Polygonzkevmglobalexitrootv2UpdateL1InfoTree
		mu               sync.Mutex
	)

	g := errgroup.Group{}

	for start := rm.UpdateToULxLyBlock; start <= endBlock; start += blocksChunkSize {
		start := start
		g.Go(func() error {
			chunkEnd := min(start+blocksChunkSize-1, endBlock)

			filter := &bind.FilterOpts{
				Start: start,
				End:   &chunkEnd,
			}
			iter, err := gerContract.FilterUpdateL1InfoTree(filter, nil, nil)
			if err != nil {
				var dataErr rpc.DataError
				if errors.As(err, &dataErr) {
					return fmt.Errorf("%w (additional details: %s)", dataErr, dataErr.ErrorData())
				}
				return err
			}

			for iter.Next() {
				mu.Lock()
				l1InfoTreeEvents = append(l1InfoTreeEvents, iter.Event)
				mu.Unlock()
			}

			return iter.Close()
		})
	}

	if err := g.Wait(); err != nil {
		return common.Hash{}, err
	}

	if len(l1InfoTreeEvents) == 0 {
		return common.Hash{}, errors.New("no UpdateL1InfoTree events found")
	}

	sort.Slice(l1InfoTreeEvents, func(i, j int) bool {
		return l1InfoTreeEvents[i].Raw.BlockNumber > l1InfoTreeEvents[j].Raw.BlockNumber
	})

	latestEvent := l1InfoTreeEvents[0]
	globalExitRoot := common.BytesToHash(
		crypto.Keccak256(
			latestEvent.MainnetExitRoot[:],
			latestEvent.RollupExitRoot[:],
		))

	return globalExitRoot, nil
}

type preEIP155Transaction struct {
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
	To       *common.Address
	Value    *big.Int
	Data     []byte
}

// encodeTxSignature combines the v, r and s into r, s and v byte array
func encodeTxSignature(v *big.Int, r, s common.Hash) []byte {
	return append(r.Bytes(), append(s.Bytes(), byte(v.Int64()))...)
}

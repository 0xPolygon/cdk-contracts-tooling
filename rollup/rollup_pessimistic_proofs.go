package rollup

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/banana-paris/polygonzkevmbridgev2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmglobalexitrootv2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/pessimistic-proofs/polygonpessimisticconsensus"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
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
	const maxDataLength = 65535

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

	gasTokenMetadata, err := bridge.GetTokenMetadata(nil, r.GasToken)
	if err != nil {
		return "", err
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

	if len(bridgeInitTxData) > maxDataLength {
		return "", fmt.Errorf("bridge init tx data length exceeds maximum allowed size (%d bytes)", maxDataLength)
	}

	V := big.NewInt(27)
	R := common.BigToHash(big.NewInt(0x5ca1ab1e0))
	S := common.BigToHash(big.NewInt(0x5ca1ab1e))

	const gasLimit = uint64(30000000)

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

	filter := &bind.FilterOpts{
		Start: rm.UpdateToULxLyBlock,
		End:   &endBlock,
	}
	iter, err := gerContract.FilterUpdateL1InfoTree(filter, nil, nil)
	if err != nil {
		return common.Hash{}, err
	}

	// We need to grab the latest UpdateL1InfoTree event
	var globalExitRoot common.Hash
	for iter.Next() {
		globalExitRoot = common.BytesToHash(
			crypto.Keccak256(
				iter.Event.MainnetExitRoot[:],
				iter.Event.RollupExitRoot[:],
			))
	}

	err = iter.Close()
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(globalExitRoot[:]), nil
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

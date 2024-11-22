package rollup

import (
	"bytes"
	"context"
	"math/big"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/banana-paris/polygonzkevmbridgev2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/pessimistic-proofs/polygonpessimisticconsensus"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// https://github.com/0xPolygonHermez/zkevm-commonjs/blob/bb0e77e9158a0fc3d06eb5de53b458bb87f77bc7/src/constants.js#L58
var L2GERManager = common.HexToAddress("0xa40D5f56745a118D0906a34E69aeC8C0Db1cB8fA")

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
func (r *RollupPessimisticProofs) GetBatchL2Data(rm *rollupmanager.RollupManager, client bind.ContractBackend) (string, error) {
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

	gasTokenMetadata, err := bridge.GasTokenMetadata(nil)
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

	bridgeInitTx := types.NewTx(
		&types.LegacyTx{
			To:       &bridgeAddr,
			Nonce:    0,
			GasPrice: common.Big0,
			Gas:      30000000,
			Value:    common.Big0,
			Data:     bridgeInitTxData,
			V:        big.NewInt(0x1b),
			R:        big.NewInt(0x5ca1ab1e0),
			S:        big.NewInt(0x5ca1ab1e),
		})

	var buf bytes.Buffer
	err = bridgeInitTx.EncodeRLP(&buf)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(buf.Bytes()), nil
}

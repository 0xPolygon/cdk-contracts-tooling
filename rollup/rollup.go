package rollup

import (
	"context"
	"encoding/json"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevm"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Rollup struct {
	Contract      *polygonzkevm.Polygonzkevm `json:"-"`
	Address       common.Address
	GenesisRoot   common.Hash
	CreationBlock uint64
	ChainID       uint64
}

func LoadFromL1ByChainID(client *ethclient.Client, rm *rollupmanager.RollupManager, chainID uint64) (*Rollup, error) {
	rID, err := rm.Contract.ChainIDToRollupID(nil, chainID)
	if err != nil {
		return nil, err
	}
	rData, err := rm.Contract.RollupIDToRollupData(nil, rID)
	if err != nil {
		return nil, err
	}
	genesisRoot, creationBlock, err := rm.GetRollupCreationInfo(context.TODO(), rID)
	if err != nil {
		return nil, err
	}

	rollup, err := polygonzkevm.NewPolygonzkevm(rData.RollupContract, client)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		Contract:      rollup,
		Address:       rData.RollupContract,
		CreationBlock: creationBlock,
		GenesisRoot:   genesisRoot,
		ChainID:       chainID,
	}, nil
}

func LoadFromFile(client *ethclient.Client, filePath string) (*Rollup, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var r Rollup
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	if client != nil {
		contract, err := polygonzkevm.NewPolygonzkevm(r.Address, client)
		if err != nil {
			return nil, err
		}
		r.Contract = contract
	}
	return &r, nil
}

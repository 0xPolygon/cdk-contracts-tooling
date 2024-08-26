package rollup

import (
	"context"
	"encoding/json"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygonvalidiumetrog"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Rollup struct {
	Contract          *polygonvalidiumetrog.Polygonvalidiumetrog `json:"-"`
	Address           common.Address
	GenesisRoot       common.Hash
	CreationBlock     uint64
	CreationTimestamp uint64
	ChainID           uint64
	Name              string
	RollupID          uint32
	GasToken          common.Address
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
	info, err := rm.GetRollupCreationInfo(context.TODO(), rID)
	if err != nil {
		return nil, err
	}

	rollup, err := polygonvalidiumetrog.NewPolygonvalidiumetrog(rData.RollupContract, client)
	if err != nil {
		return nil, err
	}
	name, err := rollup.NetworkName(nil)
	if err != nil {
		return nil, err
	}
	return &Rollup{
		Contract:          rollup,
		Address:           rData.RollupContract,
		CreationBlock:     info.Block,
		CreationTimestamp: info.Timestamp,
		GenesisRoot:       info.Root,
		ChainID:           chainID,
		Name:              name,
		RollupID:          info.RollupID,
		GasToken:          info.GasToken,
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
		contract, err := polygonvalidiumetrog.NewPolygonvalidiumetrog(r.Address, client)
		if err != nil {
			return nil, err
		}
		r.Contract = contract
	}
	return &r, nil
}

// InitContract initializes the rollup contract if not already initialized
func (r *Rollup) InitContract(ctx context.Context, client bind.ContractBackend) error {
	if r.Contract != nil {
		return nil
	}

	contract, err := polygonvalidiumetrog.NewPolygonvalidiumetrog(r.Address, client)
	if err != nil {
		return err
	}

	r.Contract = contract

	return nil
}

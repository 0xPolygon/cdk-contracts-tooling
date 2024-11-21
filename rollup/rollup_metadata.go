package rollup

import (
	"context"
	"encoding/json"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/pessimistic-proofs/polygonconsensusbase"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RollupMetadata struct {
	Address                 common.Address
	GenesisRoot             common.Hash
	CreationBlock           uint64
	CreationBlockHash       common.Hash
	CreationParentBlockHash common.Hash
	CreationTimestamp       uint64
	ChainID                 uint64
	Name                    string
	RollupID                uint32
	GasToken                common.Address
	VerifierType            rollupmanager.VerifierType
}

func NewRollupMetadata(rollupName string, chainID uint64, rollupAddr common.Address, rollupInfo rollupmanager.CreateRollupInfo) *RollupMetadata {
	return &RollupMetadata{
		Address:                 rollupAddr,
		CreationBlock:           rollupInfo.Block,
		CreationBlockHash:       rollupInfo.BlockHash,
		CreationParentBlockHash: rollupInfo.ParentBlockHash,
		CreationTimestamp:       rollupInfo.Timestamp,
		GenesisRoot:             rollupInfo.Root,
		ChainID:                 chainID,
		Name:                    rollupName,
		RollupID:                rollupInfo.RollupID,
		GasToken:                rollupInfo.GasToken,
		VerifierType:            rollupInfo.VerifierType,
	}
}

// LoadMetadataFromFile reads the rollup metadata from the json file
func LoadMetadataFromFile(filePath string) (*RollupMetadata, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var r RollupMetadata
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// LoadMetadataFromL1ByChainID reads the on-chain rollup metadata by the given chain id
func LoadMetadataFromL1ByChainID(client *ethclient.Client, rm *rollupmanager.RollupManager, chainID uint64) (*RollupMetadata, error) {
	rollupAddr, rollupID, err := rm.GetRollupIdentityData(chainID)
	if err != nil {
		return nil, err
	}

	info, err := rm.GetRollupCreationInfo(context.TODO(), rollupID)
	if err != nil {
		return nil, err
	}

	rollupBaseContract, err := polygonconsensusbase.NewPolygonconsensusbase(rollupAddr, client)
	if err != nil {
		return nil, err
	}

	rollupName, err := rollupBaseContract.NetworkName(nil)
	if err != nil {
		return nil, err
	}

	return NewRollupMetadata(rollupName, chainID, rollupAddr, info), nil
}

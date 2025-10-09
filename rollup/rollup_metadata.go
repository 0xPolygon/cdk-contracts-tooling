package rollup

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/aggchain-multisig/aggchainbase"
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
	AggchainType            rollupmanager.AggchainType
}

func NewRollupMetadata(rollupName string, chainID uint64, rollupAddr common.Address,
	aggchainType rollupmanager.AggchainType, rollupInfo rollupmanager.CreateRollupInfo) *RollupMetadata {
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
		AggchainType:            aggchainType,
	}
}

// IsPessimistic returns true if the rollup uses pessimistic proofs
func (r *RollupMetadata) IsPessimistic() bool {
	return r.VerifierType == rollupmanager.Pessimistic ||
		(r.VerifierType == rollupmanager.ALGateway && r.AggchainType == rollupmanager.PP)
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
func LoadMetadataFromL1ByChainID(ctx context.Context, client *ethclient.Client, rm *rollupmanager.RollupManager, chainID uint64) (*RollupMetadata, error) {
	rollupAddr, rollupID, err := rm.GetRollupIdentityData(chainID)
	if err != nil {
		return nil, err
	}

	info, err := rm.GetRollupCreationInfo(ctx, rollupID)
	if err != nil {
		return nil, err
	}

	aggchainBaseSC, err := aggchainbase.NewAggchainbase(rollupAddr, client)
	if err != nil {
		return nil, err
	}

	rollupName, err := aggchainBaseSC.NetworkName(nil)
	if err != nil {
		return nil, err
	}

	if info.VerifierType != rollupmanager.ALGateway {
		aggchainType := rollupmanager.PP
		if info.VerifierType == rollupmanager.StateTransition {
			aggchainType = rollupmanager.FEP
		}
		return NewRollupMetadata(rollupName, chainID, rollupAddr, aggchainType, info), nil
	}

	aggchainTypeRaw, err := aggchainBaseSC.AGGCHAINTYPE(nil)
	if err != nil {
		return nil, err
	}

	aggchainType := binary.BigEndian.Uint16(aggchainTypeRaw[:])
	return NewRollupMetadata(rollupName, chainID, rollupAddr,
		rollupmanager.AggchainType(aggchainType), info), nil
}

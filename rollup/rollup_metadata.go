package rollup

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/aggchain-multisig/aggchainbase"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	multiECDSAFeatSP1 = 1
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
func LoadMetadataFromL1ByChainID(
	ctx context.Context,
	client *ethclient.Client,
	rm *rollupmanager.RollupManager,
	chainID uint64,
) (*RollupMetadata, error) {
	// --- Step 1: Retrieve rollup identity and creation info ---
	rollupAddr, rollupID, err := rm.GetRollupIdentityData(chainID)
	if err != nil {
		return nil, fmt.Errorf("get rollup identity: %w", err)
	}

	info, err := rm.GetRollupCreationInfo(ctx, rollupID)
	if err != nil {
		return nil, fmt.Errorf("get rollup creation info: %w", err)
	}

	// --- Step 2: Initialize the AggchainBase smart contract ---
	aggchainBase, err := aggchainbase.NewAggchainbase(rollupAddr, client)
	if err != nil {
		return nil, fmt.Errorf("init aggchain base: %w", err)
	}

	rollupName, err := aggchainBase.NetworkName(nil)
	if err != nil {
		return nil, fmt.Errorf("get network name: %w", err)
	}

	// --- Step 3: Determine aggchain type ---
	switch info.VerifierType {
	case rollupmanager.ALGateway:
		consensusType, err := aggchainBase.CONSENSUSTYPE(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch consensus type (rollup %d): %w", rollupID, err)
		}

		if consensusType != multiECDSAFeatSP1 {
			return nil, fmt.Errorf("unsupported consensus type %d (rollup %d)", consensusType, rollupID)
		}

		// Fetch additional info specific to ALGateway
		rawType, err := aggchainBase.AGGCHAINTYPE(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch aggchain type (rollup %d): %w", rollupID, err)
		}

		gasToken, err := aggchainBase.GasTokenAddress(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch gas token (rollup %d): %w", rollupID, err)
		}

		info.GasToken = gasToken
		aggchainType := rollupmanager.AggchainType(binary.BigEndian.Uint16(rawType[:]))

		return NewRollupMetadata(rollupName, chainID, rollupAddr, aggchainType, info), nil

	case rollupmanager.StateTransition:
		return NewRollupMetadata(rollupName, chainID, rollupAddr, rollupmanager.FEP, info), nil

	default:
		return NewRollupMetadata(rollupName, chainID, rollupAddr, rollupmanager.PP, info), nil
	}
}

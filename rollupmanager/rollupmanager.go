package rollupmanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/aggchain-multisig/aggchainbase"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/aggchain-multisig/agglayermanager"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/aggchain-multisig/polygonzkevm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VerifierType int

const (
	StateTransition VerifierType = iota
	Pessimistic
	ALGateway
)

type AggchainType uint32

const (
	PP AggchainType = iota
	FEP
)

const blocksChunkSize = uint64(50000)

type RollupManager struct {
	Address            common.Address
	BridgeAddress      common.Address
	GERAddr            common.Address
	POLAddr            common.Address
	CreationBlock      uint64
	UpdateToULxLyBlock uint64

	Client   *ethclient.Client                `json:"-"`
	Contract *agglayermanager.Agglayermanager `json:"-"`
}

func LoadFromL1(ctx context.Context, client *ethclient.Client, address common.Address) (*RollupManager, error) {
	rm := &RollupManager{
		Client:  client,
		Address: address,
	}

	if err := rm.InitContract(ctx, client); err != nil {
		return nil, fmt.Errorf("failed to bind AgglayerManager: %w", err)
	}

	bridgeAddr, err := rm.Contract.BridgeAddress(nil)
	if err != nil {
		return nil, err
	}
	rm.BridgeAddress = bridgeAddr

	gerAddr, err := rm.Contract.GlobalExitRootManager(nil)
	if err != nil {
		return nil, err
	}
	rm.GERAddr = gerAddr

	polAddr, err := rm.Contract.Pol(nil)
	if err != nil {
		return nil, err
	}
	rm.POLAddr = polAddr

	initializedBlocks, err := rm.GetInitializedBlocks(ctx)
	if err != nil {
		return nil, err
	}

	creationBlock, creationBlockFound := initializedBlocks[1]
	upgradeBlock, upgradeBlockFound := initializedBlocks[2]
	if creationBlockFound {
		rm.CreationBlock = creationBlock
		if upgradeBlockFound {
			// Rollup Manager used to be an isolated LxLy and upgraded to unified LxLy
			rm.UpdateToULxLyBlock = upgradeBlock
		} else {
			rm.UpdateToULxLyBlock = creationBlock
		}

		return rm, nil
	}

	return nil, errors.New("couldn't find the creation block of the rollup manager")
}

func LoadFromFile(client *ethclient.Client, filePath string) (*RollupManager, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rm RollupManager
	err = json.Unmarshal(data, &rm)
	if err != nil {
		return nil, err
	}

	if client != nil {
		contract, err := agglayermanager.NewAgglayermanager(rm.Address, client)
		if err != nil {
			return nil, err
		}

		rm.Contract = contract
		rm.Client = client
	}
	return &rm, nil
}

// GetInitializedBlocks returns a mapping of version => block number of the rollup manager
func (rm *RollupManager) GetInitializedBlocks(ctx context.Context) (map[uint8]uint64, error) {
	// Get the latest block to know where to stop
	latestBlock, err := rm.Client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block header: %w", err)
	}

	blocks := make(map[uint8]uint64)

	for start := uint64(1); start <= latestBlock; start += blocksChunkSize {
		end := min(start+blocksChunkSize-1, latestBlock)

		it, err := rm.Contract.FilterInitialized(
			&bind.FilterOpts{
				Start:   start,
				End:     &end,
				Context: ctx,
			})
		if err != nil {
			return nil, fmt.Errorf("failed to filter Initialized events (range %d-%d): %w", start, end, err)
		}

		for it.Next() {
			blocks[it.Event.Version] = it.Event.Raw.BlockNumber
		}

		if err := it.Close(); err != nil {
			return nil, fmt.Errorf("failed to close iterator: %w", err)
		}

		// Allow context cancellation between chunks
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	if len(blocks) == 0 {
		return nil, errors.New("no Initialized events found")
	}

	return blocks, nil
}

type CreateRollupInfo struct {
	Root            common.Hash
	Block           uint64
	BlockHash       common.Hash
	ParentBlockHash common.Hash
	Timestamp       uint64
	ChainID         uint64
	RollupID        uint32
	GasToken        common.Address
	VerifierType    VerifierType
}

// GetRollupIdentityData returns the rollup id and its address based on the provided chain id
func (rm *RollupManager) GetRollupIdentityData(chainID uint64) (common.Address, uint32, error) {
	rollupID, err := rm.Contract.ChainIDToRollupID(nil, chainID)
	if err != nil {
		return common.Address{}, 0, err
	}
	rollupData, err := rm.Contract.RollupIDToRollupData(nil, rollupID)
	if err != nil {
		return common.Address{}, 0, err
	}

	return rollupData.RollupContract, rollupID, nil
}

// GetRollupCreation returns genesis root and the block number in which the rollup was created
func (rm *RollupManager) GetRollupCreationInfo(ctx context.Context, rollupID uint32) (CreateRollupInfo, error) {
	if rollupID == 1 && rm.UpdateToULxLyBlock != rm.CreationBlock {
		// Original rollup, created before the upgrade
		rollup, err := polygonzkevm.NewPolygonzkevm(rm.Address, rm.Client)
		if err != nil {
			return CreateRollupInfo{}, err
		}
		callOpts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(rm.UpdateToULxLyBlock - 1)}
		root, err := rollup.BatchNumToStateRoot(callOpts, 0)
		if err != nil {
			fmt.Println("couldn't find genesis for batch 0 of the rollup 1")
			return CreateRollupInfo{}, err
		}
		chainID, err := rollup.ChainID(callOpts)
		if err != nil {
			return CreateRollupInfo{}, err
		}
		b, err := rm.Client.BlockByNumber(ctx, new(big.Int).SetUint64(rm.CreationBlock))
		if err != nil {
			return CreateRollupInfo{}, err
		}

		return CreateRollupInfo{
			Root:            common.Hash(root),
			Block:           rm.CreationBlock,
			BlockHash:       b.Hash(),
			ParentBlockHash: b.ParentHash(),
			Timestamp:       b.Time(),
			ChainID:         chainID,
			RollupID:        rollupID,
			GasToken:        common.Address{},
			VerifierType:    StateTransition,
		}, nil
	}

	latestBlock, err := rm.Client.BlockNumber(ctx)
	if err != nil {
		return CreateRollupInfo{}, fmt.Errorf("failed to get latest block: %w", err)
	}

	for from := rm.UpdateToULxLyBlock; from <= latestBlock; from += blocksChunkSize {
		to := min(from+blocksChunkSize-1, latestBlock)

		it, err := rm.Contract.FilterCreateNewRollup(
			&bind.FilterOpts{
				Start:   from,
				End:     &to,
				Context: ctx,
			},
			[]uint32{rollupID},
		)
		if err != nil {
			return CreateRollupInfo{}, fmt.Errorf("failed to filter CreateNewRollup events (%d–%d): %w", from, to, err)
		}

		for it.Next() {
			event := it.Event

			b, err := rm.Client.BlockByNumber(ctx, new(big.Int).SetUint64(event.Raw.BlockNumber))
			if err != nil {
				return CreateRollupInfo{}, fmt.Errorf("failed to fetch block %d: %w", event.Raw.BlockNumber, err)
			}
			if b == nil {
				continue
			}

			latestRollupData, err := rm.Contract.RollupIDToRollupDataV2(nil, rollupID)
			if err != nil {
				return CreateRollupInfo{}, fmt.Errorf("failed to fetch rollup data for rollup id %d: %w", rollupID, err)
			}

			latestRollupType, err := rm.Contract.RollupTypeMap(nil, uint32(latestRollupData.RollupTypeID))
			if err != nil {
				return CreateRollupInfo{}, fmt.Errorf("failed to fetch rollup type for rollup type id %d: %w", latestRollupData.RollupTypeID, err)
			}

			_ = it.Close()

			return CreateRollupInfo{
				Root:            common.Hash(latestRollupType.Genesis),
				Block:           event.Raw.BlockNumber,
				BlockHash:       b.Hash(),
				ParentBlockHash: b.ParentHash(),
				Timestamp:       b.Time(),
				ChainID:         event.ChainID,
				RollupID:        rollupID,
				GasToken:        event.GasTokenAddress,
				VerifierType:    VerifierType(latestRollupData.RollupVerifierType),
			}, nil
		}

		_ = it.Close()
	}
	return CreateRollupInfo{}, fmt.Errorf("no create new rollup event for ID %d", rollupID)
}

// GetAttachedRollups returns a list of attached rollups
func (rm *RollupManager) GetAttachedRollups(ctx context.Context) (map[uint64]string, error) {
	res := make(map[uint64]string)
	// Original rollup
	data, err := rm.Contract.RollupIDToRollupData(nil, 1)
	if err != nil {
		return nil, err
	}
	zeroAddr := common.Address{}
	if data.RollupContract == zeroAddr {
		fmt.Println("rollup manager has no attached rollups")
		return res, nil
	}
	rollup, err := aggchainbase.NewAggchainbase(data.RollupContract, rm.Client)
	if err != nil {
		return nil, err
	}
	name, err := rollup.NetworkName(nil)
	if err != nil {
		return nil, err
	}
	res[data.ChainID] = name

	// Attached rollups
	latestBlock, err := rm.Client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block: %w", err)
	}

	for from := rm.CreationBlock; from <= latestBlock; from += blocksChunkSize {
		to := min(from+blocksChunkSize-1, latestBlock)

		it, err := rm.Contract.FilterCreateNewRollup(
			&bind.FilterOpts{
				Start:   from,
				End:     &to,
				Context: ctx,
			}, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to filter CreateNewRollup events (%d–%d): %w", from, to, err)
		}

		for it.Next() {
			event := it.Event

			rollup, err := aggchainbase.NewAggchainbase(event.RollupAddress, rm.Client)
			if err != nil {
				_ = it.Close()
				return nil, fmt.Errorf("failed to bind rollup contract at %s: %w", event.RollupAddress.Hex(), err)
			}

			name, err := rollup.NetworkName(nil)
			if err != nil {
				_ = it.Close()
				return nil, fmt.Errorf("failed to read NetworkName for rollup %s: %w", event.RollupAddress.Hex(), err)
			}

			res[event.ChainID] = name
		}

		if err := it.Close(); err != nil {
			return nil, fmt.Errorf("failed to close iterator: %w", err)
		}
	}

	if len(res) == 0 {
		return nil, errors.New("no CreateNewRollup events found")
	}
	return res, nil
}

// InitContract initializes the rollup manager contract if not already initialized
func (rm *RollupManager) InitContract(ctx context.Context, client *ethclient.Client) error {
	if rm.Contract != nil {
		return nil
	}

	contract, err := agglayermanager.NewAgglayermanager(rm.Address, client)
	if err != nil {
		return err
	}

	rm.Contract = contract
	rm.Client = client

	return nil
}

// GetConsensusDescription returns the description of the consensus for a given rollup ID
func (rm *RollupManager) GetConsensusDescription(ctx context.Context, rollupID uint32) (string, error) {
	startBlock := rm.CreationBlock
	latestBlock, err := rm.Client.BlockNumber(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get latest block: %w", err)
	}

	// Find rollupTypeID by scanning CreateNewRollup events in chunks
	var rollupTypeID uint32
	found := false

	for from := startBlock; from <= latestBlock && !found; from += blocksChunkSize {
		to := min(from+blocksChunkSize-1, latestBlock)

		it, err := rm.Contract.FilterCreateNewRollup(
			&bind.FilterOpts{
				Start:   from,
				End:     &to,
				Context: ctx,
			},
			[]uint32{rollupID},
		)
		if err != nil {
			return "", fmt.Errorf("FilterCreateNewRollup failed (%d-%d): %w", from, to, err)
		}

		for it.Next() {
			rollupTypeID = it.Event.RollupTypeID
			found = true
			break
		}
		_ = it.Close()
	}

	if !found {
		return "", nil
	}

	// Find AddNewRollupType event by scanning in chunks
	for from := startBlock; from <= latestBlock; from += blocksChunkSize {
		to := min(from+blocksChunkSize-1, latestBlock)

		it, err := rm.Contract.FilterAddNewRollupType(
			&bind.FilterOpts{
				Start:   from,
				End:     &to,
				Context: ctx,
			},
			[]uint32{rollupTypeID},
		)
		if err != nil {
			return "", fmt.Errorf("FilterAddNewRollupType failed (%d-%d): %w", from, to, err)
		}

		for it.Next() {
			// sanity check: only match if IDs align
			if it.Event.RollupTypeID == rollupTypeID {
				desc := it.Event.Description
				_ = it.Close()
				return desc, nil
			}
		}
		_ = it.Close()
	}

	return "", nil
}

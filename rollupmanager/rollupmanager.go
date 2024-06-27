package rollupmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonrollupmanager"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RollupManager struct {
	Client             *ethclient.Client
	Contract           *polygonrollupmanager.Polygonrollupmanager `json:"-"`
	Address            common.Address
	BridgeAddress      common.Address
	GERAddr            common.Address
	POLAddr            common.Address
	CreationBlock      uint64
	UpdateToULxLyBlock uint64
}

func LoadFromL1(client *ethclient.Client, address common.Address) (*RollupManager, error) {
	contract, err := polygonrollupmanager.NewPolygonrollupmanager(address, client)
	if err != nil {
		return nil, err
	}
	rm := RollupManager{
		Client:   client,
		Address:  address,
		Contract: contract,
	}

	bridgeAddr, err := contract.BridgeAddress(nil)
	if err != nil {
		return nil, err
	}
	rm.BridgeAddress = bridgeAddr

	gerAddr, err := contract.GlobalExitRootManager(nil)
	if err != nil {
		return nil, err
	}
	rm.GERAddr = gerAddr

	polAddr, err := contract.Pol(nil)
	if err != nil {
		return nil, err
	}
	rm.POLAddr = polAddr

	ub, err := rm.GetUpgradeBlocks(context.TODO())
	if err != nil {
		return nil, err
	}
	if block, ok := ub[1]; ok {
		rm.CreationBlock = block
	} else {
		return nil, fmt.Errorf("creation block not found")
	}
	if block, ok := ub[2]; ok {
		rm.UpdateToULxLyBlock = block
	} else {
		return nil, fmt.Errorf("upgrade to uLxLy block not found")
	}
	return &rm, nil
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
		contract, err := polygonrollupmanager.NewPolygonrollupmanager(rm.Address, client)
		if err != nil {
			return nil, err
		}
		rm.Contract = contract
	}
	return &rm, nil
}

// GetUpgradeBlocks returns a mapping of version => block number of the rollup manager
func (rm *RollupManager) GetUpgradeBlocks(ctx context.Context) (map[uint8]uint64, error) {
	it, err := rm.Contract.FilterInitialized(&bind.FilterOpts{
		Start:   1,
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}
	res := make(map[uint8]uint64)
	for it.Next() {
		res[it.Event.Version] = it.Event.Raw.BlockNumber
	}
	return res, nil
}

type CreateRollupInfo struct {
	Root     common.Hash
	Block    uint64
	ChainID  uint64
	RollupID uint32
	GasToken common.Address
}

// GetRollupCreation returns genesis root and the block number in which the rollup was created
func (rm *RollupManager) GetRollupCreationInfo(ctx context.Context, rollupID uint32) (CreateRollupInfo, error) {
	if rollupID == 1 {
		rollup, err := polygonzkevm.NewPolygonzkevm(rm.Address, rm.Client)
		if err != nil {
			return CreateRollupInfo{}, err
		}
		callOpts := &bind.CallOpts{BlockNumber: big.NewInt(int64(rm.UpdateToULxLyBlock - 1))}
		root, err := rollup.BatchNumToStateRoot(callOpts, 0)
		if err != nil {
			fmt.Println("couldn't find genesis for batch 0 of the rollup 1")
			return CreateRollupInfo{}, err
		}
		chainID, err := rollup.ChainID(callOpts)
		if err != nil {
			return CreateRollupInfo{}, err
		}

		return CreateRollupInfo{
			Root:     common.Hash(root),
			Block:    rm.CreationBlock,
			ChainID:  chainID,
			RollupID: rollupID,
			GasToken: common.Address{},
		}, nil
	}
	it, err := rm.Contract.FilterCreateNewRollup(&bind.FilterOpts{
		Start:   rm.UpdateToULxLyBlock,
		Context: ctx,
	}, []uint32{rollupID})
	if err != nil {
		return CreateRollupInfo{}, err
	}
	for it.Next() {
		rollupType, err := rm.Contract.RollupTypeMap(nil, it.Event.RollupTypeID)
		if err != nil {
			return CreateRollupInfo{}, err
		}
		return CreateRollupInfo{
			Root:     common.Hash(rollupType.Genesis),
			Block:    it.Event.Raw.BlockNumber,
			ChainID:  it.Event.ChainID,
			RollupID: rollupID,
			GasToken: it.Event.GasTokenAddress,
		}, nil
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
	rollup, err := polygonzkevm.NewPolygonzkevm(data.RollupContract, rm.Client)
	if err != nil {
		return nil, err
	}
	name, err := rollup.NetworkName(nil)
	if err != nil {
		return nil, err
	}
	res[data.ChainID] = name

	// Attached rollups
	it, err := rm.Contract.FilterCreateNewRollup(&bind.FilterOpts{
		Start:   rm.CreationBlock,
		Context: ctx,
	}, nil)
	if err != nil {
		return nil, err
	}
	for it.Next() {
		rollup, err := polygonzkevm.NewPolygonzkevm(it.Event.RollupAddress, rm.Client)
		if err != nil {
			return nil, err
		}
		name, err := rollup.NetworkName(nil)
		if err != nil {
			return nil, err
		}
		res[it.Event.ChainID] = name
	}
	return res, nil
}

// InitContract initializes the rollup manager contract if not already initialized
func (rm *RollupManager) InitContract(ctx context.Context, client bind.ContractBackend) error {
	if rm.Contract != nil {
		return nil
	}

	contract, err := polygonrollupmanager.NewPolygonrollupmanager(rm.Address, client)
	if err != nil {
		return err
	}

	rm.Contract = contract

	return nil
}

// GetConsensusDescription returns the description of the consensus for a given rollup ID
func (rm *RollupManager) GetConsensusDescription(ctx context.Context, rollupID uint32) (string, error) {
	it, err := rm.Contract.FilterAddNewRollupType(&bind.FilterOpts{
		Start:   rm.CreationBlock,
		Context: ctx,
	}, []uint32{rollupID})
	if err != nil {
		return "", err
	}

	for it.Next() {
		if it.Event.RollupTypeID == rollupID {
			return it.Event.Description, nil
		}
	}

	return "", nil
}

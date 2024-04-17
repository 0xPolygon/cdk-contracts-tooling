package rollupmanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygonrollupmanagernotupgraded"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonrollupmanager"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ADD_ROLLUP_TYPE_ROLE = common.HexToHash("0xac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f590")
	CREATE_ROLLUP_ROLE   = common.HexToHash("0xa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08")
)

type RollupManager struct {
	Client             *ethclient.Client                          `json:"-"`
	Contract           *polygonrollupmanager.Polygonrollupmanager `json:"-"`
	Address            common.Address
	BridgeAddress      common.Address
	GERAddr            common.Address
	POLAddr            common.Address
	CreationBlock      uint64
	UpdateToULxLyBlock uint64
}

func LoadFromL1(ctx context.Context, client *ethclient.Client, address common.Address) (*RollupManager, error) {
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

	ub, err := rm.GetUpgradeBlocks(ctx)
	if err != nil {
		return nil, err
	}
	var creationBlock, upgradeBlock uint64
	var creationBlockFound, upgradeBlockFound bool
	creationBlock, creationBlockFound = ub[1]
	upgradeBlock, upgradeBlockFound = ub[1]
	if creationBlockFound && upgradeBlockFound {
		// Rollup Manager used to be a isolated LxLy and upgraded to uLxLy
		rm.CreationBlock = creationBlock
		rm.UpdateToULxLyBlock = upgradeBlock
	} else {
		// Rollup Manager deployed directly on uLxLy mode
		initBlock, err := rm.GetInitializedBlock(ctx)
		if err != nil {
			return nil, err
		}
		rm.CreationBlock = initBlock
		rm.UpdateToULxLyBlock = initBlock
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

// GetInitializedBlock returns the block in which the contract was initialized
func (rm *RollupManager) GetInitializedBlock(ctx context.Context) (uint64, error) {
	notUpgraded, err := polygonrollupmanagernotupgraded.NewPolygonrollupmanagernotupgraded(rm.Address, rm.Client)
	if err != nil {
		return 0, err
	}
	it, err := notUpgraded.FilterInitialized(&bind.FilterOpts{
		Start:   1,
		Context: ctx,
	})
	if err != nil {
		return 0, err
	}
	for it.Next() {
		return it.Event.Raw.BlockNumber, nil
	}
	return 0, errors.New("initialized event not found")
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
	// // WARNING: leaving this commented as it breaks my test.
	// // Comenting this code will make it work when rollup 1 was attached in already uLxLy
	// // But will break wen rollup 1 was created before uLxLt
	// if rollupID == 1 {
	// 	// TODO: this assumes that the first rollup is always created in unified bridge mode
	// 	// but this is not the case!
	// 	rollup, err := polygonzkevm.NewPolygonzkevm(rm.Address, rm.Client)
	// 	if err != nil {
	// 		return CreateRollupInfo{}, err
	// 	}
	// 	root, err := rollup.BatchNumToStateRoot(
	// 		&bind.CallOpts{BlockNumber: big.NewInt(int64(rm.UpdateToULxLyBlock - 1))},
	// 		0,
	// 	)
	// 	if err != nil {
	// 		if strings.Contains(err.Error(), "missing trie node") {
	// 			fmt.Println("your L1 node does not support this query. Probably need an archival node")
	// 		}
	// 		fmt.Println("couldn't find genesis for batch 0 of the rollup 1")
	// 		return CreateRollupInfo{}, err
	// 	}
	// 	chainID, err := rollup.ChainID(
	// 		&bind.CallOpts{BlockNumber: big.NewInt(int64(rm.UpdateToULxLyBlock - 1))},
	// 	)
	// 	if err != nil {
	// 		return CreateRollupInfo{}, err
	// 	}

	// 	return CreateRollupInfo{
	// 		Root:     common.Hash(root),
	// 		Block:    rm.CreationBlock,
	// 		ChainID:  chainID,
	// 		RollupID: rollupID,
	// 		GasToken: common.Address{},
	// 	}, nil
	// }
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
	zeroAddr := common.Address{}
	if data.RollupContract == zeroAddr {
		fmt.Println("rollup manager has no attached rollups")
		return res, nil
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

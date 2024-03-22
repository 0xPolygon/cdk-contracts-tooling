package rollupmanager

import (
	"context"
	"fmt"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonrollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RollupManager struct {
	Address            common.Address
	BridgeAddress      common.Address
	GERAddr            common.Address
	POLAddr            common.Address
	CreationBlock      uint64
	UpdateToULxLyBlock uint64
}

func GetRollupManager(client *ethclient.Client, address common.Address) (RollupManager, error) {
	rollupManager, err := polygonrollupmanager.NewPolygonrollupmanager(address, client)
	if err != nil {
		return RollupManager{}, err
	}
	bridgeAddr, err := rollupManager.BridgeAddress(nil)
	if err != nil {
		return RollupManager{}, err
	}
	gerAddr, err := rollupManager.GlobalExitRootManager(nil)
	if err != nil {
		return RollupManager{}, err
	}
	polAddr, err := rollupManager.Pol(nil)
	if err != nil {
		return RollupManager{}, err
	}
	ub, err := GetUpgradeBlocks(context.TODO(), rollupManager)
	if err != nil {
		return RollupManager{}, err
	}
	var creationBlock, upgradeToULxLyBlock uint64
	if block, ok := ub[1]; ok {
		creationBlock = block
	} else {
		return RollupManager{}, fmt.Errorf("creation block not found")
	}
	if block, ok := ub[2]; ok {
		upgradeToULxLyBlock = block
	} else {
		return RollupManager{}, fmt.Errorf("upgrade to uLxLy block not found")
	}
	return RollupManager{
		Address:            address,
		BridgeAddress:      bridgeAddr,
		GERAddr:            gerAddr,
		POLAddr:            polAddr,
		CreationBlock:      creationBlock,
		UpdateToULxLyBlock: upgradeToULxLyBlock,
	}, nil
}

// GetUpgradeBlocks returns a mapping of version => block number of the rollup manager
func GetUpgradeBlocks(ctx context.Context, rm *polygonrollupmanager.Polygonrollupmanager) (map[uint8]uint64, error) {
	it, err := rm.FilterInitialized(&bind.FilterOpts{
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

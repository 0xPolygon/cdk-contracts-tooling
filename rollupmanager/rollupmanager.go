package rollupmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonrollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RollupManager struct {
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

// GetRollupCreation returns genesis root and the block number in which the rollup was created
func (rm *RollupManager) GetRollupCreationInfo(ctx context.Context, rollupID uint32) (common.Hash, uint64, error) {
	it, err := rm.Contract.FilterCreateNewRollup(&bind.FilterOpts{
		Start:   rm.UpdateToULxLyBlock,
		Context: ctx,
	}, []uint32{rollupID})
	if err != nil {
		return common.Hash{}, 0, err
	}
	for it.Next() {
		rollupType, err := rm.Contract.RollupTypeMap(nil, it.Event.RollupTypeID)
		if err != nil {
			return common.Hash{}, 0, err
		}
		return common.Hash(rollupType.Genesis), it.Event.Raw.BlockNumber, nil
	}
	return common.Hash{}, 0, fmt.Errorf("no create new rollup event for ID %d", rollupID)
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

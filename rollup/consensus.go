package rollup

import (
	"fmt"

	validiumelderberry "github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygonvalidiumetrog"
	rollupelderberry "github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygonzkevmetrog"
	validiumetrog "github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonvalidiumetrog"
	rollupetrog "github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmetrog"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	Etrog             = "etrog"
	Elderberry        = "elderberry"
	RollupConsensus   = "rollup"
	ValidiumConsensus = "validium"
)

type Consensus struct {
	Deploy                        func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error)
	GenerateInitializeTransaction func(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error)
}

func GetConsensus(
	version, consensusType string,
	client *ethclient.Client, auth *bind.TransactOpts,
	rm *rollupmanager.RollupManager,
	consensusAddress *common.Address,
) (*Consensus, error) {
	var consensus Consensus
	switch version {
	case Etrog:
		switch consensusType {
		case RollupConsensus:
			consensus.Deploy = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
				_, tx, _, err := rollupetrog.DeployPolygonzkevmetrog(
					auth, client, rm.GERAddr, rm.POLAddr, rm.BridgeAddress, rm.Address,
				)
				return tx, err
			}
			if consensusAddress != nil {
				c, err := rollupetrog.NewPolygonzkevmetrog(*consensusAddress, client)
				if err != nil {
					return nil, err
				}
				consensus.GenerateInitializeTransaction = c.GenerateInitializeTransaction
			}
		case ValidiumConsensus:
			consensus.Deploy = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
				_, tx, _, err := validiumetrog.DeployPolygonvalidiumetrog(
					auth, client, rm.GERAddr, rm.POLAddr, rm.BridgeAddress, rm.Address,
				)
				return tx, err
			}
			if consensusAddress != nil {
				c, err := validiumetrog.NewPolygonvalidiumetrog(*consensusAddress, client)
				if err != nil {
					return nil, err
				}
				consensus.GenerateInitializeTransaction = c.GenerateInitializeTransaction
			}
		default:
			return nil, fmt.Errorf(
				"unsupported consensus type %s. Supported values are [%s, %s]",
				consensusType, RollupConsensus, ValidiumConsensus,
			)
		}
	case Elderberry:
		switch consensusType {
		case RollupConsensus:
			consensus.Deploy = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
				_, tx, _, err := rollupelderberry.DeployPolygonzkevmetrog(
					auth, client, rm.GERAddr, rm.POLAddr, rm.BridgeAddress, rm.Address,
				)
				return tx, err
			}
			if consensusAddress != nil {
				c, err := rollupelderberry.NewPolygonzkevmetrog(*consensusAddress, client)
				if err != nil {
					return nil, err
				}
				consensus.GenerateInitializeTransaction = c.GenerateInitializeTransaction
			}
		case ValidiumConsensus:
			consensus.Deploy = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
				_, tx, _, err := validiumelderberry.DeployPolygonvalidiumetrog(
					auth, client, rm.GERAddr, rm.POLAddr, rm.BridgeAddress, rm.Address,
				)
				return tx, err
			}
			if consensusAddress != nil {
				c, err := validiumelderberry.NewPolygonvalidiumetrog(*consensusAddress, client)
				if err != nil {
					return nil, err
				}
				consensus.GenerateInitializeTransaction = c.GenerateInitializeTransaction
			}
		default:
			return nil, fmt.Errorf(
				"unsupported consensus type %s. Supported values are [%s, %s]",
				consensusType, RollupConsensus, ValidiumConsensus,
			)
		}
	default:
		return nil, fmt.Errorf(
			"unsupported version %s. SUpported versions [%s, %s]",
			version, Etrog, Elderberry,
		)
	}
	return &consensus, nil
}

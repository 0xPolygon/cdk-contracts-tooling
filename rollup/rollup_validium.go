package rollup

import (
	"context"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/legacy/elderberry/polygonvalidiumetrog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RollupValidium struct {
	*RollupMetadata

	Contract *polygonvalidiumetrog.Polygonvalidiumetrog `json:"-"`
}

// InitContract initializes the rollup contract if not already initialized
func (r *RollupValidium) InitContract(ctx context.Context, client bind.ContractBackend) error {
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

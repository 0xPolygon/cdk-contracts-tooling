package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

const (
	rollupManagerAliasFlagName     = "rollup-manager-alias"
	rollupManagerAddressFlagName   = "rollup-manager-address"
	importRollupManagerCommandName = "import-rollup-manager"
)

var (
	importRollupManagerCommand = &cli.Command{
		Name:    importRollupManagerCommandName,
		Aliases: []string{"import-rm"},
		Usage:   "Import a rollup manager smart contract, adding a new uLxLy into the network files",
		Action:  importRollupManager,
		Flags: []cli.Flag{
			l1Flag,
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"alias"},
				Usage:    "Name that will be used to store the imported rollup manager, alias of the uLxLy being imported",
				Required: true,
			},
			&cli.StringFlag{
				Name:     rollupManagerAddressFlagName,
				Aliases:  []string{"address", "addr"},
				Usage:    "Address of the rollup manager smart contract",
				Required: true,
			},
		},
	}
)

func importRollupManager(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}

	fmt.Println("loading RPC")
	rpcs, err := config.LoadRPCs()
	if err != nil {
		return err
	}
	l1Network := cliCtx.String(l1FlagName)
	client, err := rpcs.GetClient(l1Network)
	if err != nil {
		return err
	}

	fmt.Println("fetching on-chain info for the rollup manager")
	addrStr := cliCtx.String(rollupManagerAddressFlagName)
	addr := common.HexToAddress(addrStr)
	rm, err := rollupmanager.LoadFromL1(cliCtx.Context, client, addr)
	if err != nil {
		return err
	}

	fmt.Println("storing info for the rollup manager")
	alias := cliCtx.String(rollupManagerAliasFlagName)
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, alias)
	err = os.MkdirAll(rollupManagerPath, 0744)
	if err != nil {
		return err
	}
	rmData, err := json.MarshalIndent(rm, "", "   ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(rollupManagerPath, "rollupManager.json"), rmData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("importing attached rollups")
	rollups, err := rm.GetAttachedRollups(context.TODO())
	if err != nil {
		return err
	}
	rollupPath := path.Join(rollupManagerPath, "rollups")
	err = os.MkdirAll(rollupPath, 0744)
	if err != nil {
		return err
	}
	for chainID, name := range rollups {
		fmt.Println("importing rollup", name, "with chainID", chainID)
		r, err := rollup.LoadMetadataFromL1ByChainID(cliCtx.Context, client, rm, chainID)
		if err != nil {
			return err
		}
		rData, err := json.MarshalIndent(r, "", "   ")
		if err != nil {
			return err
		}

		rollupFile := fmt.Sprintf("%d-%s.json", r.RollupID, name)
		if name == "networkName" || name == "" {
			rollupFile = fmt.Sprintf("%d.json", r.RollupID)
		}

		err = os.WriteFile(path.Join(rollupPath, rollupFile), rData, 0644)
		if err != nil {
			return err
		}

		if !r.IsPessimistic() && r.GenesisRoot != (common.Hash{}) {
			// for pessimistic proof chains, the genesis root should be an empty hash
			// (https://github.com/agglayer/agglayer-contracts/blob/c8659e6282340de7bdb8fdbf7924a9bd2996bc98/contracts/v2/PolygonRollupManager.sol#L443-L446)
			if _, err := os.Stat(path.Join(baseDir, "genesis", fepDirName, r.GenesisRoot.Hex()+".json")); errors.Is(err, os.ErrNotExist) {
				fmt.Printf(
					"WARNING: the rollup '%s' with chain id %d uses a genesis with root %s. But there is no such genesis file. Please manually import it into the ./genesis directory.\n",
					name, chainID, r.GenesisRoot.Hex(),
				)
			}
		}
	}
	return nil
}

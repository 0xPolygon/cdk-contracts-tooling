package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

const (
	rollupManagerAliasFlagName   = "rollup-manager-alias"
	rollupManagerAddressFlagName = "rollup-manager-address"
)

var (
	importRollupManagerCommand = &cli.Command{
		Name:    "import-rollup-manager",
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
	rm, err := rollupmanager.GetRollupManager(client, addr)
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
	rmData, err := json.MarshalIndent(rm, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(rollupManagerPath, "rollupManager.json"), rmData, 0644)
}

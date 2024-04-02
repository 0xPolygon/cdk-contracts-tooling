package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/urfave/cli/v2"
)

const (
	chainIDFlagName         = "rollup-chain-id"
	rollupAliasFlagName     = "rollup-alias"
	importRollupCommandName = "import-rollup"
)

var (
	importRollupCommand = &cli.Command{
		Name:    importRollupCommandName,
		Aliases: []string{"import-r"},
		Usage:   "Import a rollup smart contract, adding a new CDK into the network files",
		Action:  importRollup,
		Flags: []cli.Flag{
			l1Flag,
			&cli.StringFlag{
				Name:    rollupManagerAliasFlagName,
				Aliases: []string{"rm"},
				Usage: fmt.Sprintf(
					"Name of the rollup manager to which the rollup belongs. Needs to match an already imported rollup manager (can be done by running the %s command)",
					importRollupManagerCommandName,
				),
				Required: true,
			},
			&cli.Uint64Flag{
				Name:     chainIDFlagName,
				Aliases:  []string{"id", "chainid", "chain-id"},
				Usage:    "ChainID of the rollup",
				Required: true,
			},
			&cli.StringFlag{
				Name:     rollupAliasFlagName,
				Aliases:  []string{"r"},
				Usage:    "Name of the rollup. Will be used to store the file",
				Required: true,
			},
		},
	}
)

func importRollup(cliCtx *cli.Context) error {
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

	fmt.Println("loading the rollup manager info from file")
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(client, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return err
	}

	fmt.Println("fetching on-chain info for the rollup")
	chainID := cliCtx.Uint64(chainIDFlagName)
	r, err := rollup.LoadFromL1ByChainID(client, rm, chainID)
	if err != nil {
		return err
	}

	fmt.Println("storing info for the rollup")
	alias := cliCtx.String(rollupAliasFlagName)
	rollupPath := path.Join(rollupManagerPath, "rollups")
	err = os.MkdirAll(rollupPath, 0744)
	if err != nil {
		return err
	}
	rData, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(path.Join(rollupPath, alias+".json"), rData, 0644)
}

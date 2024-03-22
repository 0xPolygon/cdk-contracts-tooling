package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var (
	nodeGenesisCommand = &cli.Command{
		Name:    "generate-node-genesis",
		Aliases: []string{"node-genesis", "genesis"},
		Usage:   "Create the genesis configuration file as expected by the node",
		Action:  nodeGenesis,
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
			&cli.StringFlag{
				Name:    rollupAliasFlagName,
				Aliases: []string{"r"},
				Usage: fmt.Sprintf(
					"Name of the rollup. Needs to match an already imported rollup (can be done by running the %s command)",
					importRollupCommandName,
				),
				Required: true,
			},
		},
	}
)

func nodeGenesis(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}

	fmt.Println("loading the rollup manager info from file")
	l1Network := cliCtx.String(l1FlagName)
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(nil, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return err
	}

	fmt.Println("loading the rollup info from file")
	rAlias := cliCtx.String(rollupAliasFlagName)
	rollupPath := path.Join(rollupManagerPath, "rollups")
	r, err := rollup.LoadFromFile(nil, path.Join(rollupPath, rAlias+".json"))
	if err != nil {
		return err
	}

	fmt.Println("loading genesis file")
	genesisData, err := os.ReadFile(path.Join(baseDir, "genesis", r.GenesisRoot.Hex()+".json"))
	if err != nil {
		return err
	}
	var genesis interface{}
	err = json.Unmarshal(genesisData, &genesis)
	if err != nil {
		return err
	}

	fmt.Println("creating genesis file")
	type l1Config struct {
		chainId                           uint64
		polygonZkEVMAddress               common.Address
		polygonRollupManagerAddress       common.Address
		polTokenAddress                   common.Address
		polygonZkEVMGlobalExitRootAddress common.Address
	}
	type nodeGenesis struct {
		l1Config           l1Config
		genesisBlockNumber uint64
		genesis            interface{}
		root               common.Hash
	}
	// alias := cliCtx.String(rollupAliasFlagName)
	// rollupPath := path.Join(rollupManagerPath, "rollups")
	// err = os.MkdirAll(rollupPath, 0744)
	// if err != nil {
	// 	return err
	// }
	// rData, err := json.MarshalIndent(r, "", " ")
	// if err != nil {
	// 	return err
	// }
	// return os.WriteFile(path.Join(rollupPath, alias+".json"), rData, 0644)
}

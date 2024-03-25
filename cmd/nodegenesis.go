package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
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
			outputFlag,
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

	fmt.Println("loading RPC info")
	rpcs, err := config.LoadRPCs()
	if err != nil {
		return err
	}
	l1ChainID, err := rpcs.GetChainID(l1Network)
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
		ChainId       uint64         `json:"chainId"`
		Rollup        common.Address `json:"polygonZkEVMAddress"`
		RollupManager common.Address `json:"polygonRollupManagerAddress"`
		POL           common.Address `json:"polTokenAddress"`
		GER           common.Address `json:"polygonZkEVMGlobalExitRootAddress"`
	}
	type nodeGenesis struct {
		L1Config                                l1Config    `json:"l1Config"`
		RollupCreationBlockNumberUsedByRollup   uint64      `json:"genesisBlockNumber"`
		RollupCreationBlockNumberUsedByValidium uint64      `json:"rollupCreationBlockNumber"`
		UpdateToULxLyBlockNumber                uint64      `json:"rollupManagerCreationBlockNumber"`
		Genesis                                 interface{} `json:"genesis"`
		Root                                    common.Hash `json:"root"`
	}
	ng := nodeGenesis{
		L1Config: l1Config{
			ChainId:       l1ChainID,
			Rollup:        r.Address,
			RollupManager: rm.Address,
			POL:           rm.POLAddr,
			GER:           rm.GERAddr,
		},
		RollupCreationBlockNumberUsedByRollup:   r.CreationBlock,
		RollupCreationBlockNumberUsedByValidium: r.CreationBlock,
		UpdateToULxLyBlockNumber:                rm.UpdateToULxLyBlock,
		Genesis:                                 genesis,
		Root:                                    r.GenesisRoot,
	}
	outputFilePath := cliCtx.String(outputFileFlagName)
	data, err := json.MarshalIndent(&ng, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputFilePath, data, 0644)
}

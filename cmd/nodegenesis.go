package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

const (
	genesisAllocsFileName = "allocs.json"
	fepDirName            = "fep"
	ppDirName             = "pp"
)

var (
	nodeGenesisCommand = &cli.Command{
		Name:    "generate-node-genesis",
		Aliases: []string{"node-genesis", "genesis"},
		Usage:   "Create the genesis configuration file as expected by the node",
		Action:  createNodeGenesis,
		Flags: []cli.Flag{
			l1Flag,
			outputFlag,
			rollupManagerAliasFlag,
			rollupAliasFlag,
		},
	}
)

type l1Config struct {
	ChainId       uint64         `json:"chainId"`
	Rollup        common.Address `json:"polygonZkEVMAddress"`
	RollupManager common.Address `json:"polygonRollupManagerAddress"`
	POL           common.Address `json:"polTokenAddress"`
	GER           common.Address `json:"polygonZkEVMGlobalExitRootAddress"`
}

type nodeGenesis struct {
	L1Config                                l1Config                `json:"l1Config"`
	RollupCreationBlockNumberUsedByRollup   uint64                  `json:"genesisBlockNumber"`
	RollupCreationBlockNumberUsedByValidium uint64                  `json:"rollupCreationBlockNumber"`
	UpdateToULxLyBlockNumber                uint64                  `json:"rollupManagerCreationBlockNumber"`
	Genesis                                 []config.GenesisAccount `json:"genesis"`
	Root                                    common.Hash             `json:"root"`
}

func createNodeGenesis(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	l1Network := cliCtx.String(l1FlagName)
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rAlias := cliCtx.String(rollupAliasFlagName)
	outputFilePath := cliCtx.String(outputFileFlagName)

	rpcs, rm, r, err := config.Load(l1Network, rmAlias, rAlias, baseDir)
	if err != nil {
		return err
	}

	l1ChainID, err := rpcs.GetChainID(l1Network)
	if err != nil {
		return err
	}

	var genesisPath string
	if r.IsPessimistic() && r.GenesisRoot == (common.Hash{}) {
		// The genesis allocations file is uniquelly identified by: l1 network alias, rollup manager alias, and rollup alias.
		// This is the case, because for pessimistic consensus, rollup's genesis state root is an empty hash.
		genesisPath = filepath.Join("genesis", ppDirName, rmAlias, genesisAllocsFileName)
	} else {
		// The genesis allocations file is uniquelly identified by the genesis state root, written in the rollup
		genesisPath = filepath.Join("genesis", fepDirName, fmt.Sprintf("%s.json", r.GenesisRoot.Hex()))
	}

	genesis, err := config.LoadGenesisAllocs(baseDir, genesisPath)
	if err != nil {
		return err
	}

	fmt.Println("creating genesis file")

	stateRoot := r.GenesisRoot
	if stateRoot == (common.Hash{}) {
		fmt.Println("WARNING: the rollup genesis state root is an empty hash, fallbacking to the root from the genesis allocations file")
		stateRoot = genesis.Root
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
		Genesis:                                 genesis.Genesis,
		Root:                                    stateRoot,
	}

	data, err := json.MarshalIndent(&ng, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputFilePath, data, 0644)
}

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
	importCombinedJsonCmd = &cli.Command{
		Name:    "import-combined-json",
		Aliases: []string{"import-cj"},
		Usage:   "Imports the combined.json file to a specified rollup in network files",
		Action:  importCombinedJson,
		Flags: []cli.Flag{
			l1Flag,
			rollupManagerAliasFlag,
			rollupAliasFlag,
		},
	}
)

type CombinedJSON struct {
	RollupManagerAddress       common.Address `json:"polygonRollupManagerAddress"`
	BridgeAddress              common.Address `json:"polygonZkEVMBridgeAddress"`
	GlobalExitRootAddress      common.Address `json:"polygonZkEVMGlobalExitRootAddress"`
	PolTokenAddress            common.Address `json:"polTokenAddress"`
	RollupManagerCreationBlock uint64         `json:"deploymentRollupManagerBlockNumber"`
	UpdateToULxLyBlock         uint64         `json:"upgradeToULxLyBlockNumber"`
	Genesis                    common.Hash    `json:"genesis"`
	RollupCreationTimestamp    uint64         `json:"createRollupTimestamp"`
	RollupCreationBlock        uint64         `json:"createRollupBlockNumber"`
	RollupAddress              common.Address `json:"rollupAddress"`
	ConsensusContract          string         `json:"consensusContract"`
	RollupID                   uint32         `json:"rollupID"`
	L2ChainID                  uint64         `json:"l2ChainID"`
	RollupGasTokenAddress      common.Address `json:"gasTokenAddress"`
	DACAddress                 common.Address `json:"polygonDataCommitteeAddress"`
	BatchL2Data                string         `json:"batchL2Data,omitempty"`
	RollupGlobalExitRoot       common.Hash    `json:"globalExitRoot,omitempty"`
}

func importCombinedJson(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}

	l1Network := cliCtx.String(l1FlagName)
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rAlias := cliCtx.String(rollupAliasFlagName)

	rpcs, rollupManager, rollupMetadata, err := config.Load(l1Network, rmAlias, rAlias, baseDir)
	if err != nil {
		return err
	}

	client, err := rpcs.GetClient(l1Network)
	if err != nil {
		return err
	}

	if err := rollupManager.InitContract(cliCtx.Context, client); err != nil {
		return err
	}

	consensusDesc, err := rollupManager.GetConsensusDescription(cliCtx.Context, rollupMetadata.RollupID)
	if err != nil {
		return err
	}

	var (
		dacAddr              common.Address
		batchL2Data          string
		rollupGlobalExitRoot common.Hash
	)

	switch rollupMetadata.VerifierType {
	case rollupmanager.Pessimistic:
		r := &rollup.RollupPessimisticProofs{RollupMetadata: rollupMetadata}
		if err := r.InitContract(cliCtx.Context, client); err != nil {
			return err
		}

		batchL2Data, err = r.GetBatchL2Data(client)
		if err != nil {
			return fmt.Errorf("failed to retrieve batch l2 data %w", err)
		}

		rollupGlobalExitRoot, err = r.GetRollupGlobalExitRoot(rollupManager, client)
		if err != nil {
			return fmt.Errorf("failed to retrieve the rollup global exit root %w", err)
		}

	default:
		r := &rollup.RollupValidium{RollupMetadata: rollupMetadata}
		if err := r.InitContract(cliCtx.Context, client); err != nil {
			return err
		}

		dacAddr, err = r.Contract.DataAvailabilityProtocol(nil)
		if err != nil {
			fmt.Println("failed to retrieve data availability protocol address", "reason:", err)
		}
	}

	combinedJson := &CombinedJSON{
		RollupManagerAddress:       rollupManager.Address,
		BridgeAddress:              rollupManager.BridgeAddress,
		GlobalExitRootAddress:      rollupManager.GERAddr,
		PolTokenAddress:            rollupManager.POLAddr,
		RollupManagerCreationBlock: rollupManager.CreationBlock,
		UpdateToULxLyBlock:         rollupManager.UpdateToULxLyBlock,
		Genesis:                    rollupMetadata.GenesisRoot,
		RollupCreationBlock:        rollupMetadata.CreationBlock,
		RollupCreationTimestamp:    rollupMetadata.CreationTimestamp,
		RollupAddress:              rollupMetadata.Address,
		ConsensusContract:          consensusDesc,
		RollupID:                   rollupMetadata.RollupID,
		L2ChainID:                  rollupMetadata.ChainID,
		RollupGasTokenAddress:      rollupMetadata.GasToken,
		DACAddress:                 dacAddr,
		BatchL2Data:                batchL2Data,
		RollupGlobalExitRoot:       rollupGlobalExitRoot,
	}

	raw, err := json.MarshalIndent(combinedJson, "", "   ")
	if err != nil {
		return err
	}

	combinedJsonFullPath := path.Join(baseDir, "networks", l1Network, rmAlias, "rollups", rAlias+"-combined.json")

	return os.WriteFile(combinedJsonFullPath, raw, 0644)
}

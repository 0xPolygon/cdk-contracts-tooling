package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	createRollupFileFlagName = "create-rollup-parameters-file"
)

var (
	createRollupCommand = &cli.Command{
		Name:    "create-rollup",
		Aliases: []string{"create-r"},
		Usage:   "Create a new rollup and attach it to a rollup manager",
		Action:  createRollupCmd,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			smartContractVersionFlag,
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"alias"},
				Usage:    "Name of the rollup manager that the rollup to be created will be attached to",
				Required: true,
			},
			&cli.PathFlag{
				Name:     createRollupFileFlagName,
				Aliases:  []string{"params", "file-params", "i"},
				Usage:    `TODO: description`,
				Required: true,
			},
		},
	}
)

func createRollupCmd(cliCtx *cli.Context) error {
	// Load configs
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}
	params, err := loadCreateRollupRollupTypeParams(cliCtx.String(createRollupFileFlagName))
	if err != nil {
		return err
	}
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	l1Network := cliCtx.String(l1FlagName)
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(client, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return err
	}

	// Check permissions
	hasRole, err := rm.Contract.HasRole(nil, rollupmanager.CREATE_ROLLUP_ROLE, auth.From)
	if err != nil {
		return err
	}
	if !hasRole {
		_, err = sendTxWithConfirmation(
			cliCtx, client, fmt.Sprintf(
				"Send tx to grant CREATE_ROLLUP_ROLE role to %s",
				auth.From,
			),
			func() (*types.Transaction, error) {
				return rm.Contract.GrantRole(auth, rollupmanager.CREATE_ROLLUP_ROLE, auth.From)
			},
		)
		if err != nil {
			return err
		}
	}

	// Create rollup
	var (
		createRollupTx *types.Transaction
	)
	_, err = sendTxWithConfirmation(
		cliCtx, client, fmt.Sprintf(
			"Send create rollup tx? Following params are going to be used:\n"+
				"- RollupTypeID: %d\n"+
				"- ChainID: %d\n"+
				"- AdminZkEVM: %s\n"+
				"- TrustedSequencer: %s\n"+
				"- GasTokenAddress: %s\n"+
				"- TrustedSequencerURL: %s\n"+
				"- NetworkName: %s\n",
			params.RollupTypeID, params.ChainID,
			params.AdminZkEVM, params.TrustedSequencer, params.GasTokenAddress,
			params.TrustedSequencerURL, params.NetworkName,
		),
		func() (*types.Transaction, error) {
			createRollupTx, err = rm.Contract.CreateNewRollup(
				auth, params.RollupTypeID, params.ChainID,
				params.AdminZkEVM, params.TrustedSequencer, params.GasTokenAddress,
				params.TrustedSequencerURL, params.NetworkName,
			)
			return createRollupTx, err
		},
	)
	if err != nil {
		return err
	}
	rID, err := rm.Contract.ChainIDToRollupID(nil, params.ChainID)
	if err != nil {
		return err
	}
	rData, err := rm.Contract.RollupIDToRollupData(nil, rID)
	if err != nil {
		return err
	}
	fmt.Printf("Rollup deployed: %+v\n", rData)
	fmt.Println("Adding rollup info to the networks directory")
	rt, err := rm.Contract.RollupTypeMap(nil, 1)
	if err != nil {
		return err
	}
	fmt.Println(common.Hash(rt.Genesis).Hex())
	// TODO: assert that rollup has the expected params
	return importRollupManager(cliCtx.Context, l1Network, rm.Address.Hex(), rmAlias)
}

type CreateRollupParams struct {
	RollupTypeID        uint32         `json:"rollupTypeID"`
	ChainID             uint64         `json:"chainID"`
	AdminZkEVM          common.Address `json:"adminZkEVM"`
	TrustedSequencer    common.Address `json:"trustedSequencer"`
	GasTokenAddress     common.Address `json:"gasTokenAddress"`
	TrustedSequencerURL string         `json:"trustedSequencerURL"`
	NetworkName         string         `json:"networkName"`
	// Consensus type. Supported values: [rollup, validium]
	ConsensusType string `json:"consensusType"`
}

func loadCreateRollupRollupTypeParams(file string) (*CreateRollupParams, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var params CreateRollupParams
	return &params, json.Unmarshal(data, &params)
}

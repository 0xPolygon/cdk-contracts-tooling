package main

import (
	"fmt"
	"math/big"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygondatacommittee"
	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	dapAddrFlagName   = "data-availability-protocol-addr"
	initEmptyFlagName = "initialize-empty"
)

var (
	setDAPCommand = &cli.Command{
		Name:    "set-data-availablity-protocol",
		Aliases: []string{"set-dap", "set-da"},
		Usage:   "Setup the data availability protocol to a deployed validium",
		Action:  setDAP,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			&cli.StringFlag{
				Name:     dapAddrFlagName,
				Aliases:  []string{"dap-addr", "addr"},
				Usage:    `Smart contract address of a Data Availability Protocol deployment`,
				Required: true,
			},
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"rm-alias"},
				Usage:    "Name of the rollup manager where the rollup belongs",
				Required: true,
			},
			&cli.StringFlag{
				Name:     rollupAliasFlagName,
				Aliases:  []string{"r-alias"},
				Usage:    "Name of the rollup where the DAP will be setup to",
				Required: true,
			},
		},
	}
	deployAndsetDACCommand = &cli.Command{
		Name:    "deploy-and-set-data-availablity-committee",
		Aliases: []string{"deployset-dac"},
		Usage:   "Deploy a DAC and set it as the data availability protocol to a deployed validium",
		Action:  deployAndSetDAC,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"rm-alias"},
				Usage:    "Name of the rollup manager where the rollup belongs",
				Required: true,
			},
			&cli.StringFlag{
				Name:     rollupAliasFlagName,
				Aliases:  []string{"r-alias"},
				Usage:    "Name of the rollup where the DAP will be setup to",
				Required: true,
			},
			&cli.BoolFlag{
				Name:     initEmptyFlagName,
				Aliases:  []string{"init"},
				Usage:    "If set to true, the DAC will be initialized with an empty committee (0/0)",
				Required: false,
			},
		},
	}
)

func setDAP(cliCtx *cli.Context) error {
	// Load configs
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rAlias := cliCtx.String(rollupAliasFlagName)
	l1Network := cliCtx.String(l1FlagName)
	rollupPath := path.Join(baseDir, "networks", l1Network, rmAlias, "rollups", rAlias+".json")
	r, err := rollup.LoadFromFile(client, rollupPath)
	if err != nil {
		return err
	}
	dap := common.HexToAddress(cliCtx.String(dapAddrFlagName))

	// Set DAP
	_, err = sendTxWithConfirmation(
		cliCtx, client, fmt.Sprintf(
			"Send tx to set the data availability protocol of validium %s to the protocol deployed at %s?",
			r.Address, dap,
		),
		func() (*types.Transaction, error) {
			return r.SetDataAvailabilityProtocol(auth, dap)
		},
	)
	return err
}

func deployAndSetDAC(cliCtx *cli.Context) error {
	// Load configs
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rAlias := cliCtx.String(rollupAliasFlagName)
	l1Network := cliCtx.String(l1FlagName)
	rollupPath := path.Join(baseDir, "networks", l1Network, rmAlias, "rollups", rAlias+".json")
	r, err := rollup.LoadFromFile(client, rollupPath)
	if err != nil {
		return err
	}

	// Deploy DAC
	dacImpl, err := deployDACImpl(cliCtx, auth, client)
	if err != nil {
		return err
	}
	dacProxy, err := deployDACProxy(cliCtx, auth, client, dacImpl)
	if err != nil {
		return err
	}

	// Init DAC
	if cliCtx.Bool(initEmptyFlagName) {
		_, err = sendTxWithConfirmation(
			cliCtx, client,
			"Initialise the deployed DAC with an empty committee (0/0)?",
			func() (*types.Transaction, error) {
				dac, err := polygondatacommittee.NewPolygondatacommittee(dacProxy, client)
				if err != nil {
					return nil, err
				}
				return dac.SetupCommittee(auth, big.NewInt(0), nil, nil)
			},
		)
		if err != nil {
			return err
		}
	}

	// Set DAP
	_, err = sendTxWithConfirmation(
		cliCtx, client, fmt.Sprintf(
			"Send tx to set the data availability protocol of validium %s to the protocol deployed at %s?",
			r.Address, dacProxy,
		),
		func() (*types.Transaction, error) {
			return r.SetDataAvailabilityProtocol(auth, dacProxy)
		},
	)
	return err
}

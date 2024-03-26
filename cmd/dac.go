package main

import (
	"fmt"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygondatacommittee"
	"github.com/urfave/cli/v2"
)

var (
	deployDACCommand = &cli.Command{
		Name:    "deploy-dac",
		Aliases: []string{},
		Usage:   "Deploy the Data Availability smart contract",
		Action:  deployDAC,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
		},
	}
	setupDACCommand = &cli.Command{
		Name:    "setup-dac",
		Aliases: []string{},
		Usage:   "Setup a Data Availability smart contract",
		Action:  setupDAC,
		Flags: []cli.Flag{
			l1Flag,
			addressFlag,
			walletFlag,
		},
	}
)

func deployDAC(cliCtx *cli.Context) error {
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
	chainID, err := rpcs.GetChainID(l1Network)
	if err != nil {
		return err
	}

	fmt.Println("loading wallet")
	wallets, err := config.LoadWallets()
	if err != nil {
		return err
	}
	walletAddr := cliCtx.String(walletFlagName)
	pk, err := wallets.GetAuth(walletAddr, chainID)
	if err != nil {
		return err
	}

	fmt.Println("sending tx")
	addr, tx, _, err := polygondatacommittee.DeployPolygondatacommittee(pk, client)
	if err != nil {
		return err
	}
	fmt.Printf("DAC will be deployed at %s with the tx %s", addr.Hex(), tx.Hash())

	return nil
}

func setupDAC(cliCtx *cli.Context) error {
	return nil
}

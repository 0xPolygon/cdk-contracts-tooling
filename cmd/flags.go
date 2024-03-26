package main

import (
	"github.com/urfave/cli/v2"
)

const (
	l1FlagName         = "l1-network"
	outputFileFlagName = "output"
	walletFlagName     = "wallet"
	addressFlagName    = "address"
)

var (
	l1Flag = &cli.StringFlag{
		Name:     l1FlagName,
		Aliases:  []string{"l1"},
		Usage:    "L1 network, such as sepolia, local, goerli, mainnet, ... Should match an entry on the rpcs.toml file",
		Required: true,
	}
	outputFlag = &cli.StringFlag{
		Name:     outputFileFlagName,
		Aliases:  []string{"o"},
		Usage:    "Output file",
		Required: true,
	}
	walletFlag = &cli.StringFlag{
		Name:     walletFlagName,
		Aliases:  []string{"w"},
		Usage:    "Address of a wallet registered at wallets.toml",
		Required: true,
	}
	addressFlag = &cli.StringFlag{
		Name:     addressFlagName,
		Aliases:  []string{"a", "addr"},
		Usage:    "Address of the smart contract",
		Required: true,
	}
)

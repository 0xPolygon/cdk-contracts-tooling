package main

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	l1FlagName                    = "l1-network"
	outputFileFlagName            = "output"
	walletFlagName                = "wallet"
	walletPasswordFlagName        = "wallet-password"
	addressFlagName               = "address"
	skipConfirmationFlagName      = "skip-confirmation"
	timeoutFlagName               = "timeout"
	setupFilePathFlagName         = "setup-file"
	implementationAddressFlagName = "implementation-address"
	smartContractVersionFlagName  = "smart-contract-version"

	defaultTimeout = time.Minute * 10
	etrog          = "etrog"
	elderbeerry    = "elderberry"
)

var (
	l1Flag = &cli.StringFlag{
		Name:     l1FlagName,
		Aliases:  []string{"l1"},
		Usage:    "L1 network, such as sepolia, local, goerli, mainnet, ... Should match an entry on the rpcs.toml file",
		Required: true,
	}
	outputFlag = &cli.PathFlag{
		Name:     outputFileFlagName,
		Aliases:  []string{"o"},
		Usage:    "Output file",
		Required: true,
	}
	addressFlag = &cli.StringFlag{
		Name:     addressFlagName,
		Aliases:  []string{"a", "addr"},
		Usage:    "Address of the smart contract",
		Required: true,
	}
	walletFlag = &cli.StringFlag{
		Name:     walletFlagName,
		Aliases:  []string{"w"},
		Usage:    "Address of a wallet registered at wallets.toml",
		Required: true,
	}
	walletPasswordFlag = &cli.StringFlag{
		Name:     walletPasswordFlagName,
		Aliases:  []string{"wp", "pass", "password"},
		Usage:    "Password of the specified wallet. This is optional, it can be inputed through CLI later on. Only applies for password protected wallets",
		Required: false,
	}
	skipConfirmationFlag = &cli.BoolFlag{
		Name:     skipConfirmationFlagName,
		Aliases:  []string{"skip"},
		Usage:    "Skip confirmation promt",
		Required: false,
	}
	timeoutFlag = &cli.DurationFlag{
		Name:     timeoutFlagName,
		Aliases:  []string{"to"},
		Usage:    fmt.Sprintf("Time after which a given operation within a command will fail. Defaults to %s", defaultTimeout),
		Required: false,
	}
	smartContractVersionFlag = &cli.StringFlag{
		Name:    smartContractVersionFlagName,
		Aliases: []string{"scv", "contract-version"},
		Usage: fmt.Sprintf(
			`Version of the smart contracts to be used. Supported options: ["%s", "%s"]`,
			etrog, elderbeerry,
		),
		Required: false,
		Action: func(ctx *cli.Context, input string) error {
			switch input {
			case "etrog":
				return nil
			case "elderberry":
				return nil
			default:
				return fmt.Errorf(
					`unsupported %s value. Supported options are: ["%s", "%s"]`,
					walletPasswordFlagName, etrog, elderbeerry,
				)
			}
		},
	}
)

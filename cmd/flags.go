package main

import (
	"github.com/urfave/cli/v2"
)

const (
	l1FlagName         = "l1-network"
	outputFileFlagName = "output"
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
)

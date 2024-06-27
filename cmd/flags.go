package main

import (
	"fmt"

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
	rollupManagerAliasFlag = &cli.StringFlag{
		Name:    rollupManagerAliasFlagName,
		Aliases: []string{"rm"},
		Usage: fmt.Sprintf(
			"Name of the rollup manager to which the rollup belongs. Needs to match an already imported rollup manager (can be done by running the %s command)",
			importRollupManagerCommandName,
		),
		Required: true,
	}
	rollupAliasFlag = &cli.StringFlag{
		Name:    rollupAliasFlagName,
		Aliases: []string{"r"},
		Usage: fmt.Sprintf(
			"Name of the rollup. Needs to match an already imported rollup (can be done by running the %s command)",
			importRollupCommandName,
		),
		Required: true,
	}
)

package main

import (
	"fmt"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/urfave/cli/v2"
)

var (
	bridgeConfigCommand = &cli.Command{
		Name:    "generate-bridge-config",
		Aliases: []string{"bridge-config", "bridge"},
		Usage:   "Create the network section of the bridge configuration file",
		Action:  bridgeConfig,
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

func bridgeConfig(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	l1Network := cliCtx.String(l1FlagName)
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	rAlias := cliCtx.String(rollupAliasFlagName)
	outputFilePath := cliCtx.String(outputFileFlagName)
	_, rm, r, _, err := config.Load(l1Network, rmAlias, rAlias, baseDir)
	if err != nil {
		return err
	}

	fmt.Println("creating nridge service config file")
	bridgeConfigTemplate := `
[NetworkConfig]
GenBlockNumber = %d
PolygonBridgeAddress = "%s"
PolygonZkEVMGlobalExitRootAddress = "%s"
PolygonRollupManagerAddress = "%s"
PolygonZkEvmAddress = "%s"
L2PolygonBridgeAddresses = ["%s"]
`
	bridgeConfig := fmt.Sprintf(
		bridgeConfigTemplate,
		rm.CreationBlock,
		rm.BridgeAddress.Hex(),
		rm.GERAddr.Hex(),
		rm.Address.Hex(),
		r.Address.Hex(),
		rm.BridgeAddress.Hex(),
	)
	return os.WriteFile(outputFilePath, []byte(bridgeConfig), 0644)
}

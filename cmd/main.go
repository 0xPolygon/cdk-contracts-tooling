package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	appName = "cdk-contracts"
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Commands = []*cli.Command{
		importContractsCommand,
		importRollupManagerCommand,
		importRollupCommand,
		nodeGenesisCommand,
		bridgeConfigCommand,
		importCombinedJsonCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("error running the application:", err)
		os.Exit(1)
	}
}

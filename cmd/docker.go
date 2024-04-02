package main

import (
	"fmt"

	"github.com/0xPolygon/cdk-contracts-tooling/docker"
	"github.com/urfave/cli/v2"
)

const (
	startMockL1FlagName  = "start-mock-l1"
	stopMockL1FlagName   = "stop-mock-l1"
	exportMockL1FlagName = "export-mock-l1"
	imageFlagName        = "docker-image-name"
	defaultImageName     = "cdk-mock-l1"
)

var (
	startMockL1Command = &cli.Command{
		Name:    startMockL1FlagName,
		Aliases: []string{"start-l1"},
		Usage:   "Starts an L1 network (dockerized geth in dev mode). The first 20 accounts of the menmonic `test test test test test test test test test test test junk` will be funded",
		Action:  startMockL1,
		Flags:   []cli.Flag{},
	}
	stopMockL1Command = &cli.Command{
		Name:    stopMockL1FlagName,
		Aliases: []string{"stop-l1"},
		Usage:   "Stops the L1 mock",
		Action:  stopMockL1,
		Flags:   []cli.Flag{},
	}
	exportMockL1Command = &cli.Command{
		Name:    exportMockL1FlagName,
		Aliases: []string{"export-l1"},
		Usage:   "Create a docekr image with the same state as the running or the last run mock L1. If the mock is still running it will be shut down",
		Action:  exportMockL1,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     imageFlagName,
				Aliases:  []string{"image", "image-name", "name"},
				Usage:    fmt.Sprintf("Name of the resulting docker image. Defaults to `%s`", defaultImageName),
				Required: false,
			},
		},
	}
)

func startMockL1(cliCtx *cli.Context) error {
	return docker.StartMockL1Docker(cliCtx.Context)
}

func stopMockL1(cliCtx *cli.Context) error {
	return docker.StopMockL1Docker()
}

func exportMockL1(cliCtx *cli.Context) error {
	if err := docker.StopMockL1Docker(); err != nil {
		return err
	}
	imageName := cliCtx.String(imageFlagName)
	if imageName == "" {
		imageName = defaultImageName
	}
	return docker.BuildL1FromMock(imageName)
}

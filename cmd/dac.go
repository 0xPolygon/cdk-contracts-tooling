package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/erc1967proxy"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygondatacommittee"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	setupDACFilePathFlagName      = "setup-file"
	implementationAddressFlagName = "implementation-address"
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
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			&cli.StringFlag{
				Name:     implementationAddressFlagName,
				Aliases:  []string{"implementation", "impl"},
				Usage:    `Smart contract address of a DAC implementation. If provided, it will be used for the proxy implementation instead of deploying a new one`,
				Required: false,
			},
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
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			&cli.PathFlag{
				Name:     setupDACFilePathFlagName,
				Aliases:  []string{"f"},
				Usage:    `File path of a JSON that looks like {"requiredSingatures": X, "members": [{"address": "0x...", "url": "http://..."}]}`,
				Required: true,
			},
		},
	}
)

func deployDAC(cliCtx *cli.Context) error {
	_, err := checkWorkingDir()
	if err != nil {
		return err
	}
	walletAddr, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}

	timeout := getTimeout(cliCtx)
	dacAddrStr := cliCtx.String(implementationAddressFlagName)
	var dacAddr common.Address
	if dacAddrStr != "" {
		dacAddr = common.HexToAddress(dacAddrStr)
		fmt.Printf("Using %s as DAC implementation (previously deployed)\n", dacAddr)
	} else {
		err = askForConfirmation(cliCtx, fmt.Sprintf(
			"Do you want to send the tx that will deploy the DAC from the address %s?",
			walletAddr,
		))
		if err != nil {
			return err
		}
		fmt.Println("deploying DAC implementaiton")
		var tx *types.Transaction
		dacAddr, tx, _, err = polygondatacommittee.DeployPolygondatacommittee(auth, client)
		if err != nil {
			return err
		}
		fmt.Printf("DAC implementation will be deployed at %s with the tx %s\n", dacAddr.Hex(), tx.Hash())
		fmt.Printf("Waiting for the tx to be mined, this will timeout after %s\n", timeout)
		ok, err := waitTxToBeMined(cliCtx.Context, client, tx, timeout)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("the transaction was mined, but it was not executed successfuly")
		}
	}

	fmt.Println("deploying proxy")
	proxyAddr, tx, _, err := erc1967proxy.DeployErc1967proxy(
		auth,
		client,
		dacAddr,
		common.Hex2Bytes("8129fc1c00000000000000000000000000000000000000000000000000000000"), // initialize() signature
	)
	if err != nil {
		return err
	}
	fmt.Printf("DAC proxy will be deployed at %s with the tx %s\n", proxyAddr.Hex(), tx.Hash())
	fmt.Printf("Waiting for the tx to be mined, this will timeout after %s\n", timeout)
	ok, err := waitTxToBeMined(cliCtx.Context, client, tx, timeout)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("the transaction was mined, but it was not executed successfuly")
	}

	return nil
}

func setupDAC(cliCtx *cli.Context) error {
	_, err := checkWorkingDir()
	if err != nil {
		return err
	}
	walletAddr, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}

	fmt.Println("loading setup file")
	setupDACFilePath := cliCtx.String(setupDACFilePathFlagName)
	data, err := os.ReadFile(setupDACFilePath)
	if err != nil {
		return err
	}
	var setup DACSetup
	err = json.Unmarshal(data, &setup)
	if err != nil {
		return err
	}
	fmt.Printf(`
DAC Configuration:
required signatures %d / %d.
Members: %+v		

`, setup.RequiredSingatures, len(setup.Members), setup.Members,
	)
	dacAddr := cliCtx.String(addressFlagName)
	err = askForConfirmation(cliCtx, fmt.Sprintf(
		"Do you want to send the tx that will setup the DAC deployed at %s from the address %s with the configuration shown above?",
		dacAddr, walletAddr,
	))
	if err != nil {
		return err
	}

	fmt.Println("sending tx")
	dac, err := polygondatacommittee.NewPolygondatacommittee(common.HexToAddress(dacAddr), client)
	if err != nil {
		return err
	}
	addrs, urls := setup.MembersAndAddrsForSetup()
	tx, err := dac.SetupCommittee(auth, big.NewInt(setup.RequiredSingatures), urls, addrs)
	if err != nil {
		return err
	}
	fmt.Printf("DAC will be setup with the tx %s\n", tx.Hash())
	timeout := getTimeout(cliCtx)
	fmt.Printf("Waiting for the tx to be mined, this will timeout after %s\n", timeout)
	ok, err := waitTxToBeMined(cliCtx.Context, client, tx, timeout)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("the transaction was mined, but it was not executed successfuly")
	}

	return nil
}

type DACMember struct {
	Address common.Address `json:"address"`
	URL     string         `json:"url"`
}

type DACSetup struct {
	RequiredSingatures int64       `json:"requiredSingatures"`
	Members            []DACMember `json:"members"`
}

func (s DACSetup) Len() int { return len(s.Members) }
func (s DACSetup) Less(i, j int) bool {
	return strings.ToUpper(s.Members[i].Address.Hex()) < strings.ToUpper(s.Members[j].Address.Hex())
}
func (s DACSetup) Swap(i, j int) { s.Members[i], s.Members[j] = s.Members[j], s.Members[i] }

func (d *DACSetup) MembersAndAddrsForSetup() (addrsBytes []byte, urls []string) {
	sort.Sort(d)
	for _, m := range d.Members {
		addrsBytes = append(addrsBytes, m.Address.Bytes()...)
		urls = append(urls, m.URL)
	}
	return
}

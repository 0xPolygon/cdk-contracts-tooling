package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygondatacommittee"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
				Name:     setupFilePathFlagName,
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
		dacAddr, err = sendTxWithConfirmation(
			cliCtx, client,
			fmt.Sprintf(
				"Do you want to send the tx that will deploy the DAC from the address %s?",
				walletAddr,
			),
			func() (*types.Transaction, common.Address, error) {
				dacAddr, tx, _, err := polygondatacommittee.DeployPolygondatacommittee(auth, client)
				return tx, dacAddr, err
			},
		)
		if err != nil {
			return err
		}
		fmt.Println("DAC implementation deployed at", dacAddr)
	}

	proxyAddr, err := deployProxy(
		cliCtx,
		auth,
		client,
		dacAddr,
		common.Hex2Bytes("8129fc1c00000000000000000000000000000000000000000000000000000000"), // initialize() signature, timeout
		timeout,
	)
	if err != nil {
		return err
	}
	fmt.Println("DAC proxy deployed at", proxyAddr)
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
	setupDACFilePath := cliCtx.String(setupFilePathFlagName)
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
	_, err = sendTxWithConfirmation(
		cliCtx, client,
		fmt.Sprintf(
			"Do you want to send the tx that will setup the DAC deployed at %s from the address %s with the configuration shown above?",
			dacAddr, walletAddr,
		),
		func() (*types.Transaction, common.Address, error) {
			dac, err := polygondatacommittee.NewPolygondatacommittee(common.HexToAddress(dacAddr), client)
			if err != nil {
				return nil, common.Address{}, err
			}
			addrs, urls := setup.MembersAndAddrsForSetup()
			tx, err := dac.SetupCommittee(auth, big.NewInt(setup.RequiredSingatures), urls, addrs)
			return tx, common.Address{}, err
		},
	)

	return err
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

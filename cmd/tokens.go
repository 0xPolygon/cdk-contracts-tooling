package main

import (
	"fmt"
	"math/big"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/erc20permitmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	spenderAddrFlagName = "spender-address"
)

var (
	approveToken = &cli.Command{
		Name:    "approve-token",
		Aliases: []string{"approve"},
		Usage:   "Approve tokens",
		Action:  approveERC20,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			amountFlag,
			&cli.StringFlag{
				Name:     implementationAddressFlagName,
				Aliases:  []string{"implementation", "impl", "token", "t"},
				Usage:    `Smart contract address of the ERC20`,
				Required: true,
			},
			&cli.StringFlag{
				Name:     spenderAddrFlagName,
				Aliases:  []string{"spender", "s"},
				Usage:    `Address that will be able too spend the tokens on behalf of the holder`,
				Required: true,
			},
		},
	}
	mintToken = &cli.Command{
		Name:    "mint-token",
		Aliases: []string{"mint"},
		Usage:   "Mint tokens",
		Action:  mintERC20,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			amountFlag,
			&cli.StringFlag{
				Name:     implementationAddressFlagName,
				Aliases:  []string{"implementation", "impl", "token", "t"},
				Usage:    `Smart contract address of the ERC20`,
				Required: true,
			},
			&cli.StringFlag{
				Name:     spenderAddrFlagName,
				Aliases:  []string{"spender", "s"},
				Usage:    `Address that will be able too spend the tokens on behalf of the holder. If not provided, the tokens will be minted to the wallet address`,
				Required: false,
			},
		},
	}
)

func approveERC20(cliCtx *cli.Context) error {
	_, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}

	erc20AddrStr := cliCtx.String(implementationAddressFlagName)
	erc20Addr := common.HexToAddress(erc20AddrStr)
	spenderAddrStr := cliCtx.String(spenderAddrFlagName)
	spenderAddr := common.HexToAddress(spenderAddrStr)
	amountStr := cliCtx.String(amountFlagName)
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return fmt.Errorf("error decoding flag %s, it should be a base 10 integer", amountFlagName)
	}
	_, err = sendTxWithConfirmation(
		cliCtx, client,
		fmt.Sprintf(
			"Do you want to send the tx that will approve %s tokens of the contract %s to be spent by %s?",
			amountStr, erc20Addr, spenderAddr,
		),
		func() (*types.Transaction, error) {
			erc20, err := erc20permitmock.NewErc20permitmock(erc20Addr, client)
			if err != nil {
				return nil, err
			}
			return erc20.Approve(auth, spenderAddr, amount)
		},
	)
	return err
}

func mintERC20(cliCtx *cli.Context) error {
	_, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}

	erc20AddrStr := cliCtx.String(implementationAddressFlagName)
	erc20Addr := common.HexToAddress(erc20AddrStr)
	spenderAddrStr := cliCtx.String(spenderAddrFlagName)
	var spenderAddr common.Address
	if spenderAddrStr != "" {
		spenderAddr = common.HexToAddress(spenderAddrStr)
	} else {
		spenderAddr = auth.From
	}
	amountStr := cliCtx.String(amountFlagName)
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok || amount == nil {
		return fmt.Errorf("error decoding flag %s, it should be a base 10 integer", amountFlagName)
	}
	_, err = sendTxWithConfirmation(
		cliCtx, client,
		fmt.Sprintf(
			"Do you want to send the tx that will mint %s tokens of the contract %s to be spent by %s?",
			amountStr, erc20Addr, spenderAddr,
		),
		func() (*types.Transaction, error) {
			erc20, err := erc20permitmock.NewErc20permitmock(erc20Addr, client)
			if err != nil {
				return nil, err
			}
			return erc20.Mint(auth, spenderAddr, amount)
		},
	)
	return err
}

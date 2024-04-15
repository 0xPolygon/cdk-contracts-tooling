package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/0xPolygon/cdk-contracts-tooling/config"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/erc1967proxy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

const (
	repoName = "cdk-contracts-tooling"
)

func checkWorkingDir() (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	_, f := path.Split(baseDir)
	if f != repoName {
		return "", fmt.Errorf("run the command from the root of the (%s)", repoName)
	}
	return baseDir, nil
}

func askForConfirmation(cliCtx *cli.Context, msg string) error {
	fmt.Println(msg, " (yes / no): ")
	skip := cliCtx.Bool(skipConfirmationFlagName)
	if skip {
		fmt.Println("(CONFIRMATION SKIPPED)")
		return nil
	}
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	if answer != "yes" {
		return errors.New("confirmation rejected")
	}
	return nil
}

func loadAuthAndClient(cliCtx *cli.Context) (common.Address, *bind.TransactOpts, *ethclient.Client, error) {
	fmt.Println("loading RPC")
	rpcs, err := config.LoadRPCs()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	l1Network := cliCtx.String(l1FlagName)
	client, err := rpcs.GetClient(l1Network)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	chainID, err := rpcs.GetChainID(l1Network)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	fmt.Println("loading wallet")
	wallets, err := config.LoadWallets()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	walletAddr := cliCtx.String(walletFlagName)
	walletPass := cliCtx.String(walletPasswordFlagName)
	auth, err := wallets.GetAuth(walletAddr, walletPass, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return common.HexToAddress(walletAddr), auth, client, nil
}

// waitTxToBeMined boolean indicates if the tx was successful
func waitTxToBeMined(parentCtx context.Context, client *ethclient.Client, tx *types.Transaction, timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if errors.Is(err, context.DeadlineExceeded) {
		return false, err
	} else if err != nil {
		return false, err
	}
	return receipt.Status == types.ReceiptStatusSuccessful, nil
}

func getTimeout(cliCtx *cli.Context) time.Duration {
	timeout := cliCtx.Duration(timeoutFlagName)
	if timeout == 0 {
		return defaultTimeout
	}
	return timeout
}

func deployProxy(
	cliCtx *cli.Context,
	auth *bind.TransactOpts,
	client *ethclient.Client,
	implementationAddr common.Address,
	initializeParams []byte,
	timeout time.Duration,
) error {
	_, err := sendTxWithConfirmation(
		cliCtx, client,
		"Do you want to send the tx that will deploy the proxy?",
		func() (*types.Transaction, common.Address, error) {
			proxyAddr, tx, _, err := erc1967proxy.DeployErc1967proxy(
				auth,
				client,
				implementationAddr,
				initializeParams,
			)
			return tx, proxyAddr, err
		},
	)
	return err
}

func sendTxWithConfirmation(
	cliCtx *cli.Context,
	client *ethclient.Client,
	confirmationQuestion string,
	sendTx func() (*types.Transaction, common.Address, error),
) (common.Address, error) {
	err := askForConfirmation(cliCtx, confirmationQuestion)
	if err != nil {
		return common.Address{}, err
	}
	fmt.Println("sending tx")
	tx, deployAddr, err := sendTx()
	if err != nil {
		return common.Address{}, err
	}

	zeroAddr := common.Address{}
	if zeroAddr != deployAddr {
		fmt.Printf(
			"Tx sent %s, smart contract will be deployed at %s\n",
			tx.Hash(), deployAddr,
		)
	} else {
		fmt.Printf("Tx %s sent", tx.Hash())
	}
	timeout := getTimeout(cliCtx)
	fmt.Printf("Waiting for the tx to be mined, this will timeout after %s\n", timeout)
	ok, err := waitTxToBeMined(cliCtx.Context, client, tx, timeout)
	if err != nil {
		return common.Address{}, err
	}
	if !ok {
		return common.Address{}, errors.New("the transaction was mined, but it was not executed successfuly")
	}
	return deployAddr, nil
}

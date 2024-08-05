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

var (
	zeroAddr = common.Address{}
)

func checkWorkingDir() (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir, f := path.Split(baseDir)
	if f == "cmd" {
		os.Chdir(dir)
		return checkWorkingDir()
	}
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
) (common.Address, error) {
	return sendTxWithConfirmation(
		cliCtx, client,
		"Do you want to send the tx that will deploy the proxy?",
		func() (*types.Transaction, error) {
			_, tx, _, err := erc1967proxy.DeployErc1967proxy(
				auth,
				client,
				implementationAddr,
				initializeParams,
			)
			return tx, err
		},
	)
}

func sendTxWithConfirmation(
	cliCtx *cli.Context,
	client *ethclient.Client,
	confirmationQuestion string,
	sendTx func() (*types.Transaction, error),
) (common.Address, error) {
	err := askForConfirmation(cliCtx, confirmationQuestion)
	if err != nil {
		return common.Address{}, err
	}
	fmt.Println("sending tx")
	tx, err := sendTx()
	if err != nil {
		return common.Address{}, err
	}

	fmt.Printf("Tx %s sent\n", tx.Hash())
	timeout := getTimeout(cliCtx)
	fmt.Printf("Waiting for the tx to be mined, this will timeout after %s\n", timeout)
	ok, err := waitTxToBeMined(cliCtx.Context, client, tx, timeout)
	if err != nil {
		return common.Address{}, err
	}
	if !ok {
		return common.Address{}, errors.New("the transaction was mined, but it was not executed successfuly")
	}
	receipt, err := client.TransactionReceipt(cliCtx.Context, tx.Hash())
	if err != nil {
		return common.Address{}, err
	}
	return receipt.ContractAddress, nil
}

func deploy(
	cliCtx *cli.Context,
	auth *bind.TransactOpts,
	client *ethclient.Client,
	deploy func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error),
	confirmationQuestion string,
) (common.Address, error) {
	// Deploy verifier implementation
	verifierImplAddr, err := sendTxWithConfirmation(
		cliCtx, client, confirmationQuestion,
		func() (*types.Transaction, error) {
			return deploy(auth, client)
		},
	)
	if err != nil {
		return common.Address{}, err
	}

	return verifierImplAddr, nil
}

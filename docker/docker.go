package docker

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartMockL1Docker(ctx context.Context) error {
	var err error
	err = os.Chdir("./docker")
	if err != nil {
		return err
	}
	defer func() { err = os.Chdir("..") }()

	err = os.RemoveAll("gethData")
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join("gethData", "geth_data", "geth"), 0744)
	if err != nil {
		return err
	}
	err = exec.Command("docker", "compose", "up", "-d", "cdk-mock-l1").Run()
	if err != nil {
		return fmt.Errorf("error running dockerized geth: %s", err)
	}
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return fmt.Errorf("error dialing dockerized geth: %s", err)
	}
	var l1Started bool
	for i := 0; i < 10; i++ {
		bn, _ := client.BlockNumber(ctx)
		if bn > 1 {
			l1Started = true
			break
		}
		time.Sleep(time.Second)
	}
	if !l1Started {
		return errors.New("mock L1 network failed to start")
	}

	fundCmd := exec.Command("docker", "compose", "up", "cdk-l1-funder")
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	fundCmd.Dir = currentDir
	out, err := fundCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running funding script: %s", err)
	}
	fmt.Println(string(out))
	var accountHasBalance bool
	for i := 0; i < 10; i++ {
		balance, err := client.BalanceAt(ctx, common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"), nil)
		if err != nil {
			return fmt.Errorf("error getting balance: %s", err)
		}
		if balance.Cmp(big.NewInt(0)) == 1 {
			accountHasBalance = true
			break
		}
		time.Sleep(time.Second)
	}
	if !accountHasBalance {
		return errors.New("failed to fund account")
	}

	return err
}

func StopMockL1Docker() error {
	var err error
	err = os.Chdir("./docker")
	if err != nil {
		return err
	}
	defer func() { err = os.Chdir("..") }()

	return exec.Command("docker", "compose", "down").Run()
}

func BuildL1FromMock() error {
	var err error
	err = os.Chdir("./docker")
	if err != nil {
		return err
	}
	defer func() { err = os.Chdir("..") }()

	err = exec.Command("docker", "compose", "up", "-d", "cdk-mock-l1").Run()
	if err != nil {
		return fmt.Errorf("error running dockerized geth: %s", err)
	}

	return err
}

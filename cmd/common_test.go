package main

import (
	"os"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	thisDir    = "cmd"
	walletAddr = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
	walletPass = "testonly"
)

func TestMain(t *testing.M) {
	baseDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_, f := path.Split(baseDir)
	if f == thisDir {
		err = os.Chdir("..")
		if err != nil {
			panic(err)
		}
	}
	t.Run()
}

func getClient() (*ethclient.Client, error) {
	return ethclient.Dial("http://localhost:8545")
}

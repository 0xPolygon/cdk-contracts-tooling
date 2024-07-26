package main

import (
	"fmt"
	"os"
	"path"
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
	if f == "cmd" {
		return "", fmt.Errorf("run the command from the root of the (%s), not from (%s)", repoName, baseDir)
	}
	return baseDir, nil
}

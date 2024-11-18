package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
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

// runCommand executes given command and extracts output from standard error in case of failure.
// Also in case command is executed successfully, the standard output is printed to the terminal.
func runCommand(cmdName string, args ...string) error {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(cmdName, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s %s", err, stderr.String())
	}

	cmdAndArgs := cmdName + " " + strings.Join(args, " ")
	fmt.Printf("command '%s' executed successfully\n%s", cmdAndArgs, stdout.String())

	return nil
}

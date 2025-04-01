package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	repoName = "cdk-contracts-tooling"
)

// checkWorkingDir checks if the current working directory is the root of the project
func checkWorkingDir() (string, error) {
	baseDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Check if the current directory is the root of the project by looking for specific files
	rootPathIndicators := []string{"go.mod", ".gitignore"}

	for _, indicator := range rootPathIndicators {
		if _, err := os.Stat(filepath.Join(baseDir, indicator)); err == nil {
			return baseDir, nil
		}
	}

	return "", fmt.Errorf("run the command from the root of the (%s). Current directory (%s) is not the project root", repoName, baseDir)
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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

// extractContractsRepoName returns the name of the contracts directory based on the proving schema
func extractContractsRepoName() (string, error) {
	parts := strings.Split(strings.TrimSuffix(contractsRepoURL, ".git"), "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid repo URL format: %s", contractsRepoURL)
	}

	return parts[len(parts)-1], nil
}

const (
	// Flags
	contractsVersionFlagName = "contracts-version"
	contractsAliasFlagName   = "contracts-alias"
	nodeVersionFlagName      = "node-version"
	buildParisFlagName       = "build-paris"
	replaceFlagName          = "replace"
	// Constants
	contractsRepoURL = "https://github.com/agglayer/agglayer-contracts.git"
	artifactsPath    = "artifacts/contracts"
	readmeTemplate   = `# %s contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [%s repo](%s), using the version %s (commit %s)

The CLI command used to generate the contracts: ` + "`$ go run ./cmd %s`" + `
`
)

var (
	importContractsCommand = &cli.Command{
		Name:    "import-contracts",
		Aliases: []string{"import-c"},
		Usage:   "Import the smart contracts from agglayer-contracts repo and generate Go bindings",
		Action:  importContracts,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     contractsVersionFlagName,
				Aliases:  []string{"cv"},
				Usage:    "Version of the smart contracts (git tag or branch)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     contractsAliasFlagName,
				Aliases:  []string{"alias"},
				Usage:    "Name that will be used to store the contracts (name of directory under contracts)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     nodeVersionFlagName,
				Aliases:  []string{"node"},
				Usage:    "Version of Node to use",
				Required: false,
				Value:    "16",
			},
			&cli.BoolFlag{
				Name:     buildParisFlagName,
				Aliases:  []string{"paris"},
				Usage:    "Build target PARIS to avoid PUSH0",
				Required: false,
				Value:    false,
			},
			&cli.BoolFlag{
				Name: replaceFlagName,
				Usage: "Replace the alias directory if it already exists. The previous one is kept aside " +
					"while the new one is generated, and it is restored if the import fails",
				Required: false,
				Value:    false,
			},
		},
	}
)

// aliasSwap puts the directory of an alias aside while it is regenerated, so it can be restored if the import fails.
//
// The alias can't be removed before running the command: the packages of the rest of this repository (rollup,
// rollupmanager...) import the generated bindings, so without the alias the command doesn't compile. It is moved
// once the command is running instead.
type aliasSwap struct {
	// path is the directory of the alias: contracts/<alias>
	path string
	// backup is where the previous directory is kept. It starts with a dot, so Go tooling ignores it, and it is in
	// the same file system as the alias, so renaming it is atomic
	backup string
	moved  bool
}

func newAliasSwap(aliasPath string) *aliasSwap {
	return &aliasSwap{
		path:   aliasPath,
		backup: filepath.Join(filepath.Dir(aliasPath), "."+filepath.Base(aliasPath)+".replaced"),
	}
}

// moveAside renames the directory of the alias to the backup. It does nothing if the alias doesn't exist
func (s *aliasSwap) moveAside() error {
	// A backup that is already there is the previous alias of an import that was interrupted (Ctrl-C, a crash): it is
	// never overwritten, and it is checked first so it is not ignored when the alias is gone
	if _, err := os.Stat(s.backup); err == nil {
		return fmt.Errorf("%s already exists: it is the previous alias of an import that was interrupted. "+
			"Remove %s (what that import left), rename %s to it, and run the command again; "+
			"or remove %s if the previous alias is not needed", s.backup, s.path, s.backup, s.backup)
	} else if !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if _, err := os.Stat(s.path); errors.Is(err, fs.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}
	if err := os.Rename(s.path, s.backup); err != nil {
		return fmt.Errorf("failed to move %s aside: %w", s.path, err)
	}
	s.moved = true
	return nil
}

// restore removes whatever was generated and puts the previous directory back. It does nothing if there was none
func (s *aliasSwap) restore() error {
	if !s.moved {
		return nil
	}
	if err := os.RemoveAll(s.path); err != nil {
		return fmt.Errorf("failed to remove %s: %w", s.path, err)
	}
	if err := os.Rename(s.backup, s.path); err != nil {
		return fmt.Errorf("failed to restore %s from %s: %w", s.path, s.backup, err)
	}
	s.moved = false
	return nil
}

// discard removes the previous directory, once the new one is complete
func (s *aliasSwap) discard() error {
	if !s.moved {
		return nil
	}
	if err := os.RemoveAll(s.backup); err != nil {
		return fmt.Errorf("failed to remove %s: %w", s.backup, err)
	}
	s.moved = false
	return nil
}

func importContracts(cliCtx *cli.Context) (err error) {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}

	contractsRepoName, err := extractContractsRepoName()
	if err != nil {
		return err
	}

	alias := cliCtx.String(contractsAliasFlagName)
	contractsPath := path.Join(baseDir, "contracts", alias)
	replace := cliCtx.Bool(replaceFlagName)

	// Fail now, not after cloning and compiling the contracts (it takes minutes), if the alias can't be written
	if _, statErr := os.Stat(contractsPath); statErr == nil && !replace {
		return fmt.Errorf("%s already exists: remove it, use another alias, or use the flag -%s to replace it",
			contractsPath, replaceFlagName)
	} else if statErr != nil && !errors.Is(statErr, fs.ErrNotExist) {
		return statErr
	}

	// clone repo in tmp
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return err
	}
	err = os.Chdir(tmpDir)
	if err != nil {
		return err
	}

	// Defer deletion of the temporary directory and all its contents
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Printf("warning: failed to remove tmp dir %s: %v\n", tmpDir, err)
		}
	}()

	err = runCommand("ls", contractsRepoName)
	absolutePath := path.Join(tmpDir, contractsRepoName)
	if err != nil {
		if strings.Contains(err.Error(), "exit status 2") || strings.Contains(err.Error(), "No such file or directory") {
			fmt.Println("cloning contracts repo into temporary directory", absolutePath)
			err = runCommand("git", "clone", contractsRepoURL)
			if err != nil {
				fmt.Printf("error cloning contracts repo from %s\n", contractsRepoURL)
				return err
			}
		} else {
			fmt.Printf("unexpected error checking if %s is already cloned\n", contractsRepoName)
			return err
		}
	}
	err = os.Chdir(absolutePath)
	if err != nil {
		return err
	}

	checkoutVersion := cliCtx.String(contractsVersionFlagName)
	fmt.Println("checking out version ", checkoutVersion)
	err = runCommand("git", "checkout", checkoutVersion)
	if err != nil {
		fmt.Println("error checking out to: ", checkoutVersion)
		return err
	}
	gitCommit, err := exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		fmt.Printf("error retrieving current Git commit hash: %v\nOutput: %s\n", err, string(gitCommit))
		return err
	}
	nodeVersion := cliCtx.String(nodeVersionFlagName)
	flagParis := cliCtx.Bool(buildParisFlagName)
	if flagParis {
		err = prepareParisMode()
		if err != nil {
			return err
		}
	}
	fmt.Println("compiling contracts node version", nodeVersion)

	err = runCommand("bash", "-l", "-c", "NODE_VERSION="+nodeVersion+" $NVM_DIR/nvm-exec npm i && npm run compile")
	if err != nil {
		fmt.Println("error compiling contracts")
		return err
	}

	if replace {
		// The command is already compiled and running, so it can move the alias that the rest of the repository
		// imports. If anything fails from here on, the previous alias is put back
		swap := newAliasSwap(contractsPath)
		if err := swap.moveAside(); err != nil {
			return err
		}
		defer func() {
			// The generation changes the working directory to the alias, which is about to be removed or kept
			if chdirErr := os.Chdir(baseDir); chdirErr != nil {
				fmt.Printf("warning: failed to go back to %s: %v\n", baseDir, chdirErr)
			}
			if err != nil {
				if swap.moved {
					fmt.Printf("the import failed: restoring the previous %s\n", contractsPath)
				}
				if restoreErr := swap.restore(); restoreErr != nil {
					err = fmt.Errorf("%w. The previous alias could not be restored, it is in %s: %v",
						err, swap.backup, restoreErr)
				}
				return
			}
			if discardErr := swap.discard(); discardErr != nil {
				fmt.Printf("warning: %v\n", discardErr)
			}
		}()
	}

	fmt.Println("creating target directory and abi and bin subdirectories")
	err = os.Mkdir(contractsPath, 0744)
	if err != nil {
		fmt.Println("error creating directory ", contractsPath)
		fmt.Printf("remove the directory, use another alias or use the flag -%s\n", replaceFlagName)
		return err
	}
	err = os.Mkdir(path.Join(contractsPath, "/bin"), 0744)
	if err != nil {
		fmt.Println("error creating directory ", path.Join(contractsPath, "/bin"))
		return err
	}
	err = os.Mkdir(path.Join(contractsPath, "/abi"), 0744)
	if err != nil {
		fmt.Println("error creating directory ", path.Join(contractsPath, "/abi"))
		return err
	}

	fmt.Println("scrapping files to find contracts to compile")
	contractPaths, err := getContractJSONFilePaths(path.Join(tmpDir, fmt.Sprintf("%s/%s", contractsRepoName, artifactsPath)))
	if err != nil {
		return err
	}
	contractsFound := len(contractPaths)
	err = os.Chdir(contractsPath)
	if err != nil {
		return err
	}
	compilationFailures := 0
	fmt.Printf("found %d contracts, compiling...\n", contractsFound)
	baseAbiPath := filepath.Join(contractsPath, "abi")
	baseBinPath := filepath.Join(contractsPath, "bin")
	for _, contractPath := range contractPaths {
		if err := importContract(contractPath, contractsPath, baseAbiPath, baseBinPath); err != nil {
			fmt.Println("WARNING: error importing contract ", contractPath, ": ", err)
			compilationFailures++
		}
	}
	fmt.Printf("%d / %d contracts compiled successfuly\n", contractsFound-compilationFailures, contractsFound)
	if replace && compilationFailures > 0 {
		// Without -replace the failures are only warnings. Here an incomplete alias would replace a complete one
		return fmt.Errorf("%d of %d contracts failed to import: the previous alias is kept, run the command without -%s "+
			"if an incomplete alias is acceptable", compilationFailures, contractsFound, replaceFlagName)
	}

	fmt.Println("cleaning up abi and bin transient folders...")
	if err := os.RemoveAll(baseAbiPath); err != nil {
		fmt.Println("error removing abi directory ", baseAbiPath)
		return err
	}

	if err := os.RemoveAll(baseBinPath); err != nil {
		fmt.Println("error removing bin directory ", baseAbiPath)
		return err
	}

	fmt.Println("running go mod tidy to fix potential dependency issues related to generated Go code")
	err = runCommand("go", "mod", "tidy")
	if err != nil {
		fmt.Println("error running go mod tidy")
		return err
	}

	fmt.Println("generating README.md")
	commandlineParams := strings.Join(os.Args[1:], " ")
	return os.WriteFile(
		contractsPath+"/README.md",
		fmt.Appendf(nil,
			readmeTemplate, alias,
			contractsRepoName, contractsRepoURL,
			checkoutVersion, strings.TrimSuffix(string(gitCommit), "\n"), commandlineParams,
		),
		0644,
	)
}

func prepareParisMode() error {
	fmt.Println("preparing PARIS mode...")
	err := runCommand("bash", "-l", "-c", "cp docker/scripts/v2/hardhat.example.paris hardhat.config.ts")
	if err != nil {
		fmt.Println("error copying hardhat.example.paris", "reason", err)
	}
	return err
}

// read file
type contractJSON struct {
	ABI      []any  `json:"abi"`
	Bytecode string `json:"bytecode"`
}

func importContract(contractPath, storingPath, baseAbiPath, baseBinPath string) error {
	data, err := os.ReadFile(contractPath)
	if err != nil {
		return err
	}
	var contract contractJSON
	if err := json.Unmarshal(data, &contract); err != nil {
		return err
	}

	// add abi and bin files
	splitted := strings.Split(contractPath, "/")
	name := strings.TrimSuffix(splitted[len(splitted)-1], ".json")
	abiData, err := json.MarshalIndent(contract.ABI, "", "   ")
	if err != nil {
		return err
	}

	abiPath := filepath.Join(baseAbiPath, name+".abi")
	binPath := filepath.Join(baseBinPath, name+".bin")

	if err := os.WriteFile(abiPath, abiData, 0644); err != nil {
		return err
	}

	if err := os.WriteFile(binPath, []byte(strings.TrimPrefix(contract.Bytecode, "0x")), 0644); err != nil {
		return err
	}

	// create directory for Go binding
	goName := strings.ToLower(name)
	if err := os.Mkdir(filepath.Join(storingPath, goName), 0744); err != nil {
		fmt.Println("error creating directory ", filepath.Join(storingPath, goName))
		return err
	}

	// generate Go binding
	if err := runCommand(
		"bash", "-l", "-c",
		fmt.Sprintf(
			"abigen --bin bin/%s.bin --abi abi/%s.abi --pkg=%s --out=%s/%s.go",
			name, name, goName, goName, goName,
		),
	); err != nil {
		fmt.Println("error generating go bindings for smart contracts")
		return err
	}

	// generate errors mapping
	return nil
}

func getContractJSONFilePaths(pathToFetchContracts string) ([]string, error) {
	filePaths := []string{}
	err := filepath.Walk(pathToFetchContracts, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, ".sol") && strings.HasSuffix(path, ".json") && !strings.Contains(path, ".dbg") {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths, err
}

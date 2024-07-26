package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	contractsVersionFlagName = "contracts-version"
	contractsAliasFlagName   = "contracts-alias"
	nodeVersionFlagName      = "node-version"
	buildParisFlagName       = "build-paris"
	pathToFetchContracts     = "zkevm-contracts/artifacts/contracts"
	readmeTemplate           = `# %s contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [zkevm-contracts repo](https://github.com/0xPolygonHermez/zkevm-contracts), using the version %s (commit %s)

Commandline used: ` + "` $ go run ./cmd %s `" + `

`
)

var (
	importContractsCommand = &cli.Command{
		Name:    "import-contracts",
		Aliases: []string{"import-c"},
		Usage:   "Import the smart contracts from zkevm-contracts repo and generate Go bindings",
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
				Usage:    "Build targe PARIS to avoid PUSH0",
				Required: false,
				Value:    false,
			},
		},
	}
)

func importContracts(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
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
	err = exec.Command("ls", "zkevm-contracts").Run()
	if err != nil {
		if err.Error() == "exit status 2" {
			fmt.Println("cloning contracts repo into temporary directory")
			err = exec.Command("git", "clone", "https://github.com/0xPolygonHermez/zkevm-contracts.git").Run()
			if err != nil {
				fmt.Println("error cloning zkevm-contracts repo")
				return err
			}
		} else {
			fmt.Println("unexpected error checking if zkevm-contracts is already cloned")
			return err
		}
	}
	err = os.Chdir(path.Join(tmpDir, "zkevm-contracts"))
	if err != nil {
		return err
	}

	checkoutVersion := cliCtx.String(contractsVersionFlagName)
	fmt.Println("checking out version ", checkoutVersion)
	err = exec.Command("git", "checkout", checkoutVersion).Run()
	if err != nil {
		fmt.Println("error checking out to: ", checkoutVersion)
		return err
	}
	gitCommit, err := exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		fmt.Println("error checking out to: ", checkoutVersion)
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
	fmt.Println("compiling contracts node version ", nodeVersion)

	err = exec.Command("bash", "-l", "-c", "NODE_VERSION="+nodeVersion+" $NVM_DIR/nvm-exec npm i && npm run compile").Run()
	if err != nil {
		fmt.Println("error compiling contracts")
		return err
	}

	fmt.Println("creating target directory and abi and bin subdirectories")
	alias := cliCtx.String(contractsAliasFlagName)
	contractsPath := path.Join(baseDir, "contracts", alias)
	err = os.Mkdir(contractsPath, 0744)
	if err != nil {
		fmt.Println("error creating directory ", contractsPath)
		fmt.Println("remove the directory or use another alias for this contracts")
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
	contractPaths, err := getContractJSONFilePaths(path.Join(tmpDir, pathToFetchContracts))
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
	for _, contractPath := range contractPaths {
		if err := importContract(contractPath, contractsPath); err != nil {
			fmt.Println("WARNING: error importing contract ", contractPath, ": ", err)
			compilationFailures++
		}
	}
	fmt.Printf("%d / %d contracts compiled successfuly\n", contractsFound-compilationFailures, contractsFound)

	fmt.Println("running go mod tidy to fix potential dependency issues related to generated Go code")
	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		fmt.Println("error running go mod tidy")
		return err
	}

	fmt.Println("generating README.md")
	commandlineParams := strings.Join(os.Args[1:], " ")
	err = os.WriteFile(
		contractsPath+"/README.md",
		[]byte(fmt.Sprintf(
			readmeTemplate, alias, checkoutVersion, strings.TrimSuffix(string(gitCommit), "\n"), commandlineParams,
		)),
		0644,
	)
	if err != nil {
		return err
	}
	return nil
}

func prepareParisMode() error {
	fmt.Println("preparing PARIS mode...")
	err := exec.Command("bash", "-l", "-c", "cp docker/scripts/v2/hardhat.example.paris hardhat.config.ts").Run()
	if err != nil {
		fmt.Println("error copying hardhat.example.paris")
	}
	return err
}

func importContract(contractPath, storingPath string) error {
	// read file
	type contractJSON struct {
		ABI      []interface{} `json:"abi"`
		Bytecode string        `json:"bytecode"`
	}
	data, err := os.ReadFile(contractPath)
	if err != nil {
		return err
	}
	var contract contractJSON
	err = json.Unmarshal(data, &contract)
	if err != nil {
		return err
	}

	// add abi and bin files
	splitted := strings.Split(contractPath, "/")
	name := strings.TrimSuffix(splitted[len(splitted)-1], ".json")
	abiData, err := json.MarshalIndent(contract.ABI, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(storingPath+"/abi/"+name+".abi", abiData, 0644)
	if err != nil {
		return err
	}
	err = os.WriteFile(storingPath+"/bin/"+name+".bin", []byte(strings.TrimPrefix(contract.Bytecode, "0x")), 0644)
	if err != nil {
		return err
	}

	// create directory for Go binding
	goName := strings.ToLower(name)
	err = os.Mkdir(path.Join(storingPath, goName), 0744)
	if err != nil {
		fmt.Println("error creating directory ", storingPath+"/"+goName)
		return err
	}
	// generate Go binding
	err = exec.Command(
		"bash", "-l", "-c",
		fmt.Sprintf(
			"abigen --bin bin/%s.bin --abi abi/%s.abi --pkg=%s --out=%s/%s.go",
			name, name, goName, goName, goName,
		),
	).Run()
	if err != nil {
		fmt.Println("error compiling contracts")
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

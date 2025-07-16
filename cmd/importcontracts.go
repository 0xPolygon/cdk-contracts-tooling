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

type ProvingSchema string

const (
	FullExecutionProofs ProvingSchema = "fep"
	PessimisticProofs   ProvingSchema = "pp"
)

// GetContractsRepoURL returns the URL of the contracts repository based on the proving schema
func (ps ProvingSchema) GetContractsRepoURL() (string, error) {
	return contractsRepoURL, nil
}

// GetContractsRepoName returns the name of the contracts directory based on the proving schema
func (ps ProvingSchema) GetContractsRepoName() (string, error) {
	url, err := ps.GetContractsRepoURL()
	if err != nil {
		return "", err
	}

	parts := strings.Split(strings.TrimSuffix(url, ".git"), "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid repo URL format: %s", url)
	}

	return parts[len(parts)-1], nil
}

func (ps ProvingSchema) String() string {
	return string(ps)
}

var validProvingSchemas = map[ProvingSchema]struct{}{
	FullExecutionProofs: {},
	PessimisticProofs:   {},
}

const (
	// Flags
	contractsVersionFlagName = "contracts-version"
	contractsAliasFlagName   = "contracts-alias"
	nodeVersionFlagName      = "node-version"
	buildParisFlagName       = "build-paris"
	provingSystemFlagName    = "proving-schema"

	// Contracts repo URL
	contractsRepoURL = "https://github.com/agglayer/agglayer-contracts.git"

	artifactsPath  = "artifacts/contracts"
	readmeTemplate = `# %s contracts

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
			&cli.StringFlag{
				Name:     provingSystemFlagName,
				Aliases:  []string{"ps"},
				Usage:    "Proving system: 'fep' (full execution proofs) or 'pp' (pessimistic proofs) - used for organizing output directories",
				Required: false,
				Value:    string(FullExecutionProofs),
			},
		},
	}
)

func importContracts(cliCtx *cli.Context) error {
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}

	provingSchemaRaw := cliCtx.String(provingSystemFlagName)
	if err := validateProvingSchema(provingSchemaRaw); err != nil {
		return err
	}

	provingSchema := ProvingSchema(provingSchemaRaw)
	contractsRepoName, err := provingSchema.GetContractsRepoName()
	if err != nil {
		return err
	}
	contractsRepoURL, err := provingSchema.GetContractsRepoURL()
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
	err = runCommand("ls", contractsRepoName)
	if err != nil {
		if strings.Contains(err.Error(), "exit status 2") || strings.Contains(err.Error(), "No such file or directory") {
			fmt.Println("cloning contracts repo into temporary directory")
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
	err = os.Chdir(path.Join(tmpDir, contractsRepoName))
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

	fmt.Println("creating target directory and abi and bin subdirectories")
	alias := cliCtx.String(contractsAliasFlagName)
	contractsPath := path.Join(baseDir, "contracts", provingSchemaRaw, alias)
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
	for _, contractPath := range contractPaths {
		if err := importContract(contractPath, contractsPath); err != nil {
			fmt.Println("WARNING: error importing contract ", contractPath, ": ", err)
			compilationFailures++
		}
	}
	fmt.Printf("%d / %d contracts compiled successfuly\n", contractsFound-compilationFailures, contractsFound)

	fmt.Println("running go mod tidy to fix potential dependency issues related to generated Go code")
	err = runCommand("go", "mod", "tidy")
	if err != nil {
		fmt.Println("error running go mod tidy")
		return err
	}

	fmt.Println("generating README.md")
	commandlineParams := strings.Join(os.Args[1:], " ")
	err = os.WriteFile(
		contractsPath+"/README.md",
		fmt.Appendf(nil,
			readmeTemplate, alias,
			contractsRepoName, contractsRepoURL,
			checkoutVersion, strings.TrimSuffix(string(gitCommit), "\n"), commandlineParams,
		),
		0644,
	)
	if err != nil {
		return err
	}
	return nil
}

// validateProvingSchema ensures only valid values are used
func validateProvingSchema(value string) error {
	if _, exists := validProvingSchemas[ProvingSchema(value)]; !exists {
		return fmt.Errorf("invalid proving-schema value provided: %s, must be either %s or %s", value, FullExecutionProofs, PessimisticProofs)
	}
	return nil
}

func prepareParisMode() error {
	fmt.Println("preparing PARIS mode...")
	err := runCommand("bash", "-l", "-c", "cp docker/scripts/v2/hardhat.example.paris hardhat.config.ts")
	if err != nil {
		fmt.Println("error copying hardhat.example.paris", "reason", err)
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
	abiData, err := json.MarshalIndent(contract.ABI, "", "   ")
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
	err = runCommand(
		"bash", "-l", "-c",
		fmt.Sprintf(
			"abigen --bin bin/%s.bin --abi abi/%s.abi --pkg=%s --out=%s/%s.go",
			name, name, goName, goName, goName,
		),
	)
	if err != nil {
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

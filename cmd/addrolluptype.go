package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	verifierelderberry "github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/fflonkverifier"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/verifierrolluphelpermock"
	verifieretrog "github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/fflonkverifier"
	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	mockVerifierFlagName      = "mock-verifier"
	addRollupTypeFileFlagName = "add-rollup-type-parameters-file"
	rollupConsensus           = "rollup"
	validiumConsensus         = "validium"
)

var (
	addRollupTypeCommand = &cli.Command{
		Name:    "add-rollup-type",
		Aliases: []string{"add-rt"},
		Usage:   "Deploy necessary implementations and add type to a rollup manager",
		Action:  addRollupTypeCmd,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			smartContractVersionFlag,
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"alias"},
				Usage:    "Name of the rollup manager where the rollup type will be added to",
				Required: true,
			},
			&cli.PathFlag{
				Name:     addRollupTypeFileFlagName,
				Aliases:  []string{"params", "file-params", "i"},
				Usage:    `TODO: description`,
				Required: true,
			},
		},
	}
)

func addRollupTypeCmd(cliCtx *cli.Context) error {
	// Load configs
	baseDir, err := checkWorkingDir()
	if err != nil {
		return err
	}
	_, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}
	params, err := loadAddRollupTypeParams(cliCtx.String(addRollupTypeFileFlagName))
	if err != nil {
		return err
	}
	rmAlias := cliCtx.String(rollupManagerAliasFlagName)
	l1Network := cliCtx.String(l1FlagName)
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(client, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return err
	}
	if _, err := os.Stat(path.Join(baseDir, "genesis", params.GenesisRoot.Hex()+".json")); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf(
			"no genesis file found for root %s. Please add genesis file to the genesis directory",
			params.GenesisRoot.Hex(),
		)
	}
	contractsVersion := cliCtx.String(smartContractVersionFlagName)

	// Check permissions
	hasRole, err := rm.Contract.HasRole(nil, rollupmanager.ADD_ROLLUP_TYPE_ROLE, auth.From)
	if err != nil {
		return err
	}
	if !hasRole {
		_, err = sendTxWithConfirmation(
			cliCtx, client, fmt.Sprintf(
				"Send tx to grant ADD_ROLLUP_TYPE_ROLE role to %s",
				auth.From,
			),
			func() (*types.Transaction, error) {
				return rm.Contract.GrantRole(auth, rollupmanager.ADD_ROLLUP_TYPE_ROLE, auth.From)
			},
		)
		if err != nil {
			return err
		}
	}

	// if ((await rollupManagerContract.hasRole(ADD_ROLLUP_TYPE_ROLE, deployer.address)) == false)
	//     await rollupManagerContract.grantRole(ADD_ROLLUP_TYPE_ROLE, deployer.address);

	// Verifier
	if params.VerifierImplementation != zeroAddr {
		fmt.Println("using verifier:", params.VerifierImplementation)
	} else {
		fmt.Println("deploying verifier")
		var deployFn func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error)
		if params.UseMockVerifier {
			deployFn = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
				_, tx, _, err := verifierrolluphelpermock.DeployVerifierrolluphelpermock(auth, client)
				return tx, err
			}
		} else {
			switch contractsVersion {
			case etrog:
				deployFn = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
					_, tx, _, err := verifieretrog.DeployFflonkverifier(auth, client)
					return tx, err
				}
			case elderbeerry:
				deployFn = func(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.Transaction, error) {
					_, tx, _, err := verifierelderberry.DeployFflonkverifier(auth, client)
					return tx, err
				}
			}
		}
		addr, err := deploy(
			cliCtx, auth, client, deployFn,
			"Do you want to send the tx that will deploy the verifier?",
		)
		if err != nil {
			return err
		}
		fmt.Println("Verifier implementation deployed at:", addr)
		params.VerifierImplementation = addr
	}

	// Consensus
	if params.ConsensusImplementation != zeroAddr {
		fmt.Println("using consensus:", params.ConsensusImplementation)
	} else {
		consensus, err := rollup.GetConsensus(
			contractsVersion, params.ConsensusType, client, auth, rm, nil,
		)
		if err != nil {
			return err
		}

		addr, err := deploy(
			cliCtx, auth, client, consensus.Deploy,
			"Do you want to send the tx that will deploy the consensus?",
		)
		if err != nil {
			return err
		}
		fmt.Println("Consensus implementation deployed at:", addr)
		params.ConsensusImplementation = addr
	}

	// Last rollup type ID
	rollupTypeCounterBeforeTx, err := rm.Contract.RollupTypeCount(nil)
	if err != nil {
		return err
	}

	// Add rollup type
	_, err = sendTxWithConfirmation(
		cliCtx, client, fmt.Sprintf(
			"Send add rollup type tx? Following params are going to be used:\n"+
				"- consensus implementation: %s\n"+
				"- verifier implementation: %s\n"+
				"- fork ID: %d\n"+
				"- rollup compatibility ID: 0\n"+
				"- genesis root: %s\n"+
				"- description: %s\n",
			params.ConsensusImplementation, params.VerifierImplementation,
			params.ForkID, params.GenesisRoot, params.Description,
		),
		func() (*types.Transaction, error) {
			return rm.Contract.AddNewRollupType(
				auth,
				params.ConsensusImplementation, params.VerifierImplementation,
				params.ForkID, 0, params.GenesisRoot, params.Description,
			)
		},
	)
	if err != nil {
		return err
	}
	rollupTypeCounterAfterTx, err := rm.Contract.RollupTypeCount(nil)
	if err != nil {
		return err
	}
	if rollupTypeCounterBeforeTx == rollupTypeCounterAfterTx {
		return errors.New("something went wrong, rollup type counter is still the same after sending the tx")
	}
	fmt.Println("Rollup type added with ID:", rollupTypeCounterAfterTx)
	rt, err := rm.Contract.RollupTypeMap(nil, rollupTypeCounterAfterTx)
	if err != nil {
		return err
	}
	fmt.Println(common.Hash(rt.Genesis).Hex())
	// TODO: assert that the rt output matches the expected inputs
	return nil
}

type AddRollupTypeParams struct {
	// If provided, will use this instead of deploying a new verifier implementation
	VerifierImplementation common.Address `json:"verifierImplementation"`
	// If true will deploy a mock verifier instead of a real one
	UseMockVerifier bool `json:"useMockVerifier"`
	// If provided, will use this instead of deploying a new consensus implementation
	ConsensusImplementation common.Address `json:"consensusImplementation"`
	// Consensus type. Supported values: [rollup, validium]
	ConsensusType string      `json:"consensusType"`
	ForkID        uint64      `json:"forkID"`
	GenesisRoot   common.Hash `json:"genesisRoot"`
	Description   string      `json:"description"`
}

func loadAddRollupTypeParams(file string) (*AddRollupTypeParams, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var params AddRollupTypeParams
	return &params, json.Unmarshal(data, &params)
}

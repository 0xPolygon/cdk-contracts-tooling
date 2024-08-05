package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/proxyadmin"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/transparentupgradableproxy"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/erc20permitmock"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonrollupmanagernotupgraded"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmbridgev2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmdeployer"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmglobalexitrootv2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmtimelock"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

const (
	deployParametersFileFlagName = "rollup-manager-parameters-file"
)

var (
	deployRollupManagerCommand = &cli.Command{
		Name:    "deploy-rollup-manager",
		Aliases: []string{"deploy-rm", "drm"},
		Usage:   "Deploy the Rollup Manager smart contract",
		Action:  deployRollupManager,
		Flags: []cli.Flag{
			l1Flag,
			walletFlag,
			walletPasswordFlag,
			skipConfirmationFlag,
			timeoutFlag,
			smartContractVersionFlag,
			&cli.PathFlag{
				Name:     deployParametersFileFlagName,
				Aliases:  []string{"params", "file-params", "i"},
				Usage:    `TODO: description`,
				Required: true,
			},
			outputFlag,
			&cli.StringFlag{
				Name:     rollupManagerAliasFlagName,
				Aliases:  []string{"alias"},
				Usage:    "Name that will be used to store the rollup manager info once deployed, alias of the uLxLy being imported",
				Required: true,
			},
		},
	}
)

func deployRollupManager(cliCtx *cli.Context) error {
	_, err := checkWorkingDir()
	if err != nil {
		return err
	}
	walletAddr, auth, client, err := loadAuthAndClient(cliCtx)
	if err != nil {
		return err
	}

	// Load params
	params, output, err := loadDeployRollupManagerFiles(
		cliCtx.Path(deployParametersFileFlagName),
		cliCtx.Path(outputFileFlagName),
	)
	if err != nil {
		return err
	}
	if output.DeploymentCompleted {
		return errors.New(
			"the provided output file indicates that the deployment is completed. " +
				"Use a different file if you want to do a new deployment",
		)
	}
	defer storeOutputFile(cliCtx.Path(outputFileFlagName), output)
	output.Salt = params.Salt
	output.TrustedAggregator = params.TrustedAggregator
	output.Admin = params.Admin
	output.DeployerAddress = walletAddr
	if output.PolTokenAddress == zeroAddr && params.PolTokenAddress != zeroAddr {
		output.PolTokenAddress = params.PolTokenAddress
	}

	// Deploy POL
	if output.PolTokenAddress == zeroAddr {
		fmt.Println("Deploying POL token")
		polTokenAddr, err := sendTxWithConfirmation(
			cliCtx, client,
			"Do you want to send the tx that will deploy the POL token?",
			func() (*types.Transaction, error) {
				_, tx, _, err := erc20permitmock.DeployErc20permitmock(
					auth, client, "Pol Token", "POL",
					walletAddr, big.NewInt(20000000),
				)
				return tx, err
			},
		)
		if err != nil {
			return err
		}
		fmt.Println("POL deployed at ", polTokenAddr)
		output.PolTokenAddress = polTokenAddr
	} else {
		fmt.Println("Using already deployed POL token at", output.PolTokenAddress)
	}

	// Deploy PolygonZkEVMDeployer
	fmt.Println("Deploying PolygonZkEVMDeployer...")
	_, tx, _, err := polygonzkevmdeployer.DeployPolygonzkevmdeployer(
		&bind.TransactOpts{
			NoSend:   true,
			Nonce:    big.NewInt(0),
			Signer:   scalableSigner{}.SignerFn,
			GasLimit: 1_000_000,                   // gas limit 1M
			GasPrice: big.NewInt(100_000_000_000), // gas price 100 gWEI

		},
		client, params.Admin,
	)
	if err != nil {
		return err
	}

	// check if already deployed
	deterministicSender, err := scalableSigner{}.Sender(tx)
	if err != nil {
		return err
	}
	zkevmDeployerAddr := crypto.CreateAddress(deterministicSender, 0)

	code, err := client.CodeAt(cliCtx.Context, zkevmDeployerAddr, nil)
	if err != nil {
		return err
	}
	if len(code) > 0 {
		fmt.Println("PolygonZkEVMDeployer was already deployed")
	} else {
		// Fund deployer
		neededFunds := big.NewInt(0).Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas())))
		dsBalance, err := client.BalanceAt(cliCtx.Context, deterministicSender, nil)
		if err != nil {
			return err
		}
		if dsBalance.Cmp(neededFunds) == -1 {
			_, err = sendTxWithConfirmation(
				cliCtx, client,
				fmt.Sprintf(
					"The address used for the deterministic deployment %s doesn't have enough funds for the tx. Send %s ETH from %s?",
					deterministicSender, neededFunds.String(), walletAddr,
				),
				func() (*types.Transaction, error) {
					nonce, err := client.PendingNonceAt(cliCtx.Context, walletAddr)
					if err != nil {
						return nil, err
					}
					gasPrice, err := client.SuggestGasPrice(cliCtx.Context)
					if err != nil {
						return nil, err
					}
					fundTx := types.NewTransaction(nonce, deterministicSender, neededFunds, 21000, gasPrice, nil)
					fundTx, err = auth.Signer(walletAddr, fundTx)
					if err != nil {
						return nil, err
					}
					err = client.SendTransaction(cliCtx.Context, fundTx)
					return fundTx, err
				},
			)
			if err != nil {
				return err
			}
		}
		// Deploy deployer
		actualDeployerAddr, err := sendTxWithConfirmation(
			cliCtx, client,
			fmt.Sprintf(
				"Do you want to send the tx that will deploy the PolygonZkEVMDeployer with owner %s? Note that this tx will be sent from %s",
				params.Admin.Hex(), deterministicSender,
			),
			func() (*types.Transaction, error) {
				err := client.SendTransaction(cliCtx.Context, tx)
				return tx, err
			},
		)
		if err != nil {
			return err
		}
		if zkevmDeployerAddr != actualDeployerAddr {
			return fmt.Errorf("unexpcted deployer addr. Expected %s, actual: %s", zkevmDeployerAddr, actualDeployerAddr)
		}
	}
	deployer, err := polygonzkevmdeployer.NewPolygonzkevmdeployer(zkevmDeployerAddr, client)
	if err != nil {
		return err
	}
	fmt.Println("PolygonZkEVMDeployer addr:", zkevmDeployerAddr)
	output.ZkEVMDeployerContract = zkevmDeployerAddr

	// Deploy proxy admin
	fmt.Println("Deploying proxy admin...")
	proxyABI, err := proxyadmin.ProxyadminMetaData.GetAbi()
	if err != nil {
		return err
	}
	if proxyABI == nil {
		return errors.New("GetABI returned nil")
	}
	dataCallAdmin, err := proxyABI.Pack("transferOwnership", walletAddr)
	if err != nil {
		return err
	}
	_, proxyTx, _, err := proxyadmin.DeployProxyadmin(
		&bind.TransactOpts{
			NoSend: true,
			From:   auth.From,   // Irrelevant, we just want tx data
			Signer: auth.Signer, // Irrelevant, we just want tx data
		},
		client,
	)
	if err != nil {
		return err
	}
	proxyAdminAddr, err := create2Deployment(
		cliCtx, client, auth, deployer, zkevmDeployerAddr, params.Salt, proxyTx.Data(), dataCallAdmin,
		fmt.Sprintf(
			"Do you want to send the tx that will deploy the proxy (using the deployer)? Admin of the proxy will be set to %s",
			walletAddr,
		),
	)
	if err != nil {
		return err
	}
	proxy, err := proxyadmin.NewProxyadmin(proxyAdminAddr, client)
	if err != nil {
		return err
	}
	fmt.Println("proxy admin addr:", proxyAdminAddr)
	output.ProxyAdminAddress = proxyAdminAddr

	// Deploy implementation PolygonZkEVMBridge
	fmt.Println("Deploying PolygonZkEVMBridge implementation...")
	auth.GasLimit = 5500000 // gas limit with create are mess up D:
	_, bridgeImplTx, _, err := polygonzkevmbridgev2.DeployPolygonzkevmbridgev2(
		&bind.TransactOpts{
			NoSend: true,
			From:   auth.From,   // Irrelevant, we just want tx data
			Signer: auth.Signer, // Irrelevant, we just want tx data
		},
		client,
	)
	if err != nil {
		return err
	}
	bridgeImplementationAddr, err := create2Deployment(
		cliCtx, client, auth, deployer, zkevmDeployerAddr, params.Salt, bridgeImplTx.Data(), nil,
		fmt.Sprintf(
			"Do you want to send the tx that will deploy the proxy (using the deployer)? Admin of the proxy will be set to %s",
			walletAddr,
		),
	)
	if err != nil {
		return err
	}
	auth.GasLimit = 0
	bridgeImplementation, err := polygonzkevmbridgev2.NewPolygonzkevmbridgev2(bridgeImplementationAddr, client)
	if err != nil {
		return err
	}
	// Sanity check, just checking the right code was deployed on the expected addr
	checkCounter, err := bridgeImplementation.DepositCount(nil)
	if err != nil {
		return err
	}
	if checkCounter.Cmp(big.NewInt(0)) != 0 {
		return fmt.Errorf("counter should be 0 but it's %s", checkCounter.String())
	}
	fmt.Println("bridge implementation addr:", bridgeImplementationAddr)

	// Useless tx to match zkevm-contracts scripts
	_, err = sendTxWithConfirmation(
		cliCtx, client,
		"Do you want to send a useless tx just to increase the nonce (required step to move forward)?",
		func() (*types.Transaction, error) {
			nonce, err := client.PendingNonceAt(cliCtx.Context, walletAddr)
			if err != nil {
				return nil, err
			}
			gasPrice, err := client.SuggestGasPrice(cliCtx.Context)
			if err != nil {
				return nil, err
			}
			fundTx := types.NewTransaction(nonce, auth.From, big.NewInt(0), 21000, gasPrice, nil)
			fundTx, err = auth.Signer(walletAddr, fundTx)
			if err != nil {
				return nil, err
			}
			err = client.SendTransaction(cliCtx.Context, fundTx)
			return fundTx, err
		},
	)
	if err != nil {
		return err
	}

	// Deploy PolygonZkEVMTimelock
	fmt.Println("Deploying PolygonZkEVMTimelock...")
	currentNonce, err := client.NonceAt(cliCtx.Context, walletAddr, nil)
	if err != nil {
		return err
	}
	// Nonce globalExitRoot: currentNonce + 1 (deploy bridge proxy) + 1(impl globalExitRoot)
	// + 1 (deployTimelock) + 1 (transfer Ownership Admin) = +4
	nonceProxyGlobalExitRoot := currentNonce + 4
	fmt.Println("nonceProxyGlobalExitRoot: ", nonceProxyGlobalExitRoot)
	// nonceProxyRollupManager :Nonce globalExitRoot + 1 (proxy globalExitRoot) + 1 (impl rollupManager) = +2
	nonceProxyRollupManager := nonceProxyGlobalExitRoot + 2
	precalculateGlobalExitRootAddress := crypto.CreateAddress(walletAddr, nonceProxyGlobalExitRoot)
	precalculateRollupManager := crypto.CreateAddress(walletAddr, nonceProxyRollupManager)

	// Deploy timelock
	var timelockAddr common.Address
	if output.TimelockContractAddress != zeroAddr &&
		output.PolygonRollupManager != zeroAddr &&
		output.PolygonZkEVMGlobalExitRootAddress != zeroAddr {
		timelockAddr = output.TimelockContractAddress
		fmt.Println("using already deployed timelock", timelockAddr)
	} else {
		if output.TimelockContractAddress != zeroAddr {
			fmt.Println("I'm afraid we can't recover this situation. Time to change the salt")
		}
		// If both are not deployed, it's better to deploy them both again
		output.TimelockContractAddress = zeroAddr
		output.PolygonRollupManager = zeroAddr
		output.PolygonZkEVMGlobalExitRootAddress = zeroAddr

		timelockAdmins := []common.Address{params.TimelockAdminAddress}
		timelockAddr, err = sendTxWithConfirmation(
			cliCtx, client,
			fmt.Sprintf(
				"Do you want to send the tx that will deploy the PolygonZkEVMTimelock with admin %s? This timelock will be used for the Rollup Manager with precalculated address %s",
				params.TimelockAdminAddress, precalculateRollupManager,
			),
			func() (*types.Transaction, error) {
				_, tx, _, err := polygonzkevmtimelock.DeployPolygonzkevmtimelock(
					auth, client, big.NewInt(params.MinDelayTimelock), timelockAdmins, timelockAdmins, params.TimelockAdminAddress, precalculateRollupManager,
				)
				return tx, err
			},
		)
		if err != nil {
			return err
		}
		// Sanity check
		timelock, err := polygonzkevmtimelock.NewPolygonzkevmtimelock(timelockAddr, client)
		if err != nil {
			return err
		}
		actualRMAddr, err := timelock.PolygonZkEVM(nil)
		if err != nil {
			return err
		}
		if precalculateRollupManager != actualRMAddr {
			return fmt.Errorf("timelock rollup manager was expected to be %s instead of %s", precalculateRollupManager, actualRMAddr)
		}
		fmt.Println("PolygonZkEVMTimelock addr", timelockAddr)
		output.TimelockContractAddress = timelockAddr
	}

	// Transfer ownership of the proxyAdmin to timelock
	proxyOwner, err := proxy.Owner(nil)
	if err != nil {
		return err
	}
	if proxyOwner != timelockAddr {
		_, err = sendTxWithConfirmation(
			cliCtx, client,
			"Do you want to send the tx that will transfer the ownership of the proxy to the timelock?",
			func() (*types.Transaction, error) {
				return proxy.TransferOwnership(auth, timelockAddr)
			},
		)
		if err != nil {
			return err
		}
		newProxyOwner, err := proxy.Owner(nil)
		if err != nil {
			return err
		}
		if timelockAddr != newProxyOwner {
			return fmt.Errorf(
				"owner of the proxy should be the timelock (%s), but it's %s",
				timelockAddr, newProxyOwner,
			)
		}
	}

	// Deploy PolygonZkEVMBridgeV2 proxy
	fmt.Println("Deploying PolygonZkEVMBridgeV2 proxy...")
	_, bridgeProxyTx, _, err := transparentupgradableproxy.DeployTransparentupgradableproxy(
		&bind.TransactOpts{
			NoSend: true,
			From:   auth.From,   // Irrelevant, we just want tx data
			Signer: auth.Signer, // Irrelevant, we just want tx data
		},
		client,
		bridgeImplementationAddr,
		proxyAdminAddr,
		[]byte{},
	)
	if err != nil {
		return err
	}

	bridgeABI, err := polygonzkevmbridgev2.Polygonzkevmbridgev2MetaData.GetAbi()
	if err != nil {
		return err
	}
	if bridgeABI == nil {
		return errors.New("GetABI returned nil")
	}
	dataCallProxy, err := bridgeABI.Pack("initialize",
		uint32(0),        // networkIDMainnet
		common.Address{}, // gasTokenAddressMainnet"
		uint32(0),        // gasTokenNetworkMainnet
		precalculateGlobalExitRootAddress,
		precalculateRollupManager,
		[]byte{}, // gasTokenMetadata
	)
	if err != nil {
		return err
	}
	bridgeProxyAddr, err := create2Deployment(
		cliCtx, client, auth, deployer, zkevmDeployerAddr, params.Salt, bridgeProxyTx.Data(), dataCallProxy,
		fmt.Sprintf(
			"Do you want to send the tx that will deploy the proxy (using the deployer)? Admin of the proxy will be set to %s",
			walletAddr,
		),
	)
	if err != nil {
		return err
	}
	bridge, err := polygonzkevmbridgev2.NewPolygonzkevmbridgev2(bridgeProxyAddr, client)
	if err != nil {
		return err
	}
	actualGERAddr, err := bridge.GlobalExitRootManager(nil)
	if err != nil {
		return err
	}
	if precalculateGlobalExitRootAddress != actualGERAddr {
		return fmt.Errorf("GER addr should be %s instead of %s", precalculateGlobalExitRootAddress, actualGERAddr)
	}
	actualRMAddr, err := bridge.PolygonRollupManager(nil)
	if err != nil {
		return err
	}
	if precalculateRollupManager != actualRMAddr {
		return fmt.Errorf("rollup Manager addr should be %s instead of %s", precalculateRollupManager, actualRMAddr)
	}
	fmt.Println("PolygonZkEVMBridgeV2 proxy addr:", bridgeProxyAddr)
	output.PolygonZkEVMBridgeAddress = bridgeProxyAddr

	// Deploy PolygonZkEVMGlobalExitRootV2 proxy
	fmt.Println("Deploying PolygonZkEVMGlobalExitRootV2 proxy...")
	if output.PolygonZkEVMGlobalExitRootAddress != zeroAddr {
		fmt.Println("PolygonZkEVMGlobalExitRootV2 already deployed")
	} else {
		// Deploy GER implementation
		gerImplAddr, err := sendTxWithConfirmation(
			cliCtx, client,
			"Do you want to send the tx that will deploy the PolygonZkEVMGlobalExitRootV2 implementation?",
			func() (*types.Transaction, error) {
				_, tx, _, err := polygonzkevmglobalexitrootv2.DeployPolygonzkevmglobalexitrootv2(
					auth,
					client,
					precalculateRollupManager,
					bridgeProxyAddr,
				)
				return tx, err
			},
		)
		if err != nil {
			return err
		}

		// Deploy GER proxy
		GERProxyAddr, err := deployProxy(
			cliCtx,
			auth,
			client,
			gerImplAddr,
			[]byte{},
		)
		if err != nil {
			return err
		}
		if GERProxyAddr != precalculateGlobalExitRootAddress {
			return fmt.Errorf(
				"GER was expected to be deployed at %s, but was deployed at %s instead",
				precalculateGlobalExitRootAddress, GERProxyAddr,
			)
		}
		ger, err := polygonzkevmglobalexitrootv2.NewPolygonzkevmglobalexitrootv2(GERProxyAddr, client)
		if err != nil {
			return err
		}
		actualRMAddr, err := ger.RollupManager(nil)
		if err != nil {
			return err
		}
		if actualRMAddr != precalculateRollupManager {
			return fmt.Errorf("rollup manager addr is %s but should be %s", actualRMAddr, precalculateRollupManager)
		}
		actualBridgeAddr, err := ger.BridgeAddress(nil)
		if err != nil {
			return err
		}
		if actualBridgeAddr != bridgeProxyAddr {
			return fmt.Errorf("bridge addr is %s but should be %s", actualBridgeAddr, bridgeProxyAddr)
		}
		fmt.Println("PolygonZkEVMGlobalExitRootV2 proxy deployed at", GERProxyAddr)
		output.PolygonZkEVMGlobalExitRootAddress = GERProxyAddr
	}

	if params.UseDeployerAsAdminForRollupManager {
		fmt.Println("Using deployer as admin of the rollup manager insetad of the timelock!")
		timelockAddr = walletAddr
	}

	// Deploy PolygonRollupManagerNotUpgraded proxy
	fmt.Println("Deploying PolygonRollupManagerNotUpgraded proxy...")
	if output.PolygonRollupManager != zeroAddr {
		fmt.Println("PolygonRollupManagerNotUpgraded already deployed")
	} else {
		// Deploy PolygonRollupManagerNotUpgraded implementation
		rmImplAddr, err := sendTxWithConfirmation(
			cliCtx, client,
			fmt.Sprintf(
				"Do you want to send the tx that will deploy the PolygonRollupManagerNotUpgraded implementation? "+
					"Following values are going to be used:\n"+
					"trustedAggregator: %s\n"+
					"pendingStateTimeout: %d\n"+
					"trustedAggregatorTimeout: %d\n"+
					"admin: %s\n"+
					"timelockAddressRollupManager: %s\n"+
					"emergencyCouncilAddress: %s\n"+
					"PolygonZkEVMGlobalExitRootAddress: %s\n"+
					"PolTokenAddress: %s\n"+
					"bridgeProxyAddr: %s\n",
				params.TrustedAggregator, params.PendingStateTimeout, params.TrustedAggregatorTimeout,
				params.Admin, timelockAddr, params.EmergencyCouncilAddress,
				output.PolygonZkEVMGlobalExitRootAddress, output.PolTokenAddress, bridgeProxyAddr,
			),

			func() (*types.Transaction, error) {
				_, tx, _, err := polygonrollupmanagernotupgraded.DeployPolygonrollupmanagernotupgraded(
					auth,
					client,
					output.PolygonZkEVMGlobalExitRootAddress,
					output.PolTokenAddress,
					bridgeProxyAddr,
				)
				return tx, err
			},
		)
		if err != nil {
			return err
		}
		// Deploy PolygonRollupManagerNotUpgraded proxy
		rmABI, err := polygonrollupmanagernotupgraded.PolygonrollupmanagernotupgradedMetaData.GetAbi()
		if err != nil {
			return err
		}
		if rmABI == nil {
			return errors.New("GetABI returned nil")
		}
		initializeData, err := rmABI.Pack(
			"initialize",
			params.TrustedAggregator,
			params.PendingStateTimeout,
			params.TrustedAggregatorTimeout,
			params.Admin,
			timelockAddr,
			params.EmergencyCouncilAddress,
			zeroAddr, zeroAddr, uint64(0), uint64(0), // unused parameters
		)
		if err != nil {
			return err
		}
		RMProxyAddr, err := deployProxy(
			cliCtx,
			auth,
			client,
			rmImplAddr,
			initializeData,
		)
		if err != nil {
			return err
		}
		if RMProxyAddr != precalculateRollupManager {
			return fmt.Errorf(
				"rollup Manager was expected to be deployed at %s, but was deployed at %s instead",
				precalculateRollupManager, RMProxyAddr,
			)
		}
		rm, err := polygonrollupmanagernotupgraded.NewPolygonrollupmanagernotupgraded(RMProxyAddr, client)
		if err != nil {
			return err
		}
		actualBridgeAddr, err := rm.BridgeAddress(nil)
		if err != nil {
			return err
		}
		if actualBridgeAddr != bridgeProxyAddr {
			return fmt.Errorf("bridge addr is %s but should be %s", actualBridgeAddr, bridgeProxyAddr)
		}
		fmt.Println("PolygonRollupManagerNotUpgraded proxy deployed at", RMProxyAddr)
		output.PolygonRollupManager = RMProxyAddr

		rmInfo, err := rollupmanager.LoadFromL1(cliCtx.Context, client, RMProxyAddr)
		if err != nil {
			return err
		}
		output.DeploymentBlockNumber = rmInfo.CreationBlock
	}

	// TODO: checks, a lot of checks and asserts https://github.com/0xPolygonHermez/zkevm-contracts/blob/main/deployment/v2/3_deployContracts.ts#L488

	output.DeploymentCompleted = true
	l1Network := cliCtx.String(l1FlagName)
	alias := cliCtx.String(rollupManagerAliasFlagName)
	return importRollupManager(cliCtx.Context, l1Network, output.PolygonRollupManager.Hex(), alias)
}

type DeployRollupManagerParams struct {
	UseDeployerAsAdminForRollupManager bool           `json:"useDeployerAsAdminForRollupManager"`
	TimelockAdminAddress               common.Address `json:"timelockAdminAddress"`
	MinDelayTimelock                   int64          `json:"minDelayTimelock"`
	Salt                               common.Hash    `json:"salt"`
	InitialZkEVMDeployerOwner          common.Address `json:"initialZkEVMDeployerOwner"`
	Admin                              common.Address `json:"admin"`
	TrustedAggregator                  common.Address `json:"trustedAggregator"`
	TrustedAggregatorTimeout           uint64         `json:"trustedAggregatorTimeout"`
	PendingStateTimeout                uint64         `json:"pendingStateTimeout"`
	EmergencyCouncilAddress            common.Address `json:"emergencyCouncilAddress"`
	PolTokenAddress                    common.Address `json:"polTokenAddress"`
}

type DeployRollupManagerOutput struct {
	DeploymentCompleted               bool           `json:"deploymentCompleted"`
	PolygonRollupManager              common.Address `json:"polygonRollupManager"`
	PolygonZkEVMBridgeAddress         common.Address `json:"polygonZkEVMBridgeAddress"`
	PolygonZkEVMGlobalExitRootAddress common.Address `json:"polygonZkEVMGlobalExitRootAddress"`
	PolTokenAddress                   common.Address `json:"polTokenAddress"`
	ZkEVMDeployerContract             common.Address `json:"zkEVMDeployerContract"`
	DeployerAddress                   common.Address `json:"deployerAddress"`
	TimelockContractAddress           common.Address `json:"timelockContractAddress"`
	DeploymentBlockNumber             uint64         `json:"deploymentBlockNumber"`
	Admin                             common.Address `json:"admin"`
	TrustedAggregator                 common.Address `json:"trustedAggregator"`
	ProxyAdminAddress                 common.Address `json:"proxyAdminAddress"`
	Salt                              common.Hash    `json:"salt"`
}

func loadDeployRollupManagerFiles(paramsFile, outputFile string) (*DeployRollupManagerParams, *DeployRollupManagerOutput, error) {
	data, err := os.ReadFile(paramsFile)
	if err != nil {
		return nil, nil, err
	}
	var params DeployRollupManagerParams
	err = json.Unmarshal(data, &params)
	if err != nil {
		return nil, nil, err
	}
	output, err := loadRollupManagerOutput(outputFile)
	if err != nil {
		return nil, nil, err
	}
	return &params, output, nil
}

func loadRollupManagerOutput(outputFile string) (*DeployRollupManagerOutput, error) {
	data, err := os.ReadFile(outputFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &DeployRollupManagerOutput{}, nil
		}
		return nil, err
	}
	var output DeployRollupManagerOutput
	return &output, json.Unmarshal(data, &output)
}

func storeOutputFile(outputFile string, output *DeployRollupManagerOutput) error {
	outputData, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputFile, outputData, 0644)
}

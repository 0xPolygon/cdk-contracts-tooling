package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/proxyadmin"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/elderberry/polygonzkevmbridgev2"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmdeployer"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmtimelock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

const (
	ownerFlagName = "owner"
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
			// &cli.StringFlag{
			// 	Name:     implementationAddressFlagName,
			// 	Aliases:  []string{"implementation", "impl"},
			// 	Usage:    `Smart contract address of a DAC implementation. If provided, it will be used for the proxy implementation instead of deploying a new one`,
			// 	Required: false,
			// },
			&cli.StringFlag{
				Name:     ownerFlagName,
				Aliases:  []string{},
				Usage:    `Address with owner role of the deployment`,
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

	// Deploy PolygonZkEVMDeployer
	fmt.Println("Deploying PolygonZkEVMDeployer...")
	// Deterministic address deployment
	owner := common.HexToAddress(cliCtx.String(ownerFlagName))
	// txData, err := getPolygonZkEVMDeployerDeployTransactionData(owner)
	// if err != nil {
	// 	return err
	// }
	_, tx, _, err := polygonzkevmdeployer.DeployPolygonzkevmdeployer(
		&bind.TransactOpts{
			NoSend:   true,
			Nonce:    big.NewInt(0),
			Signer:   scalableSigner{}.SignerFn,
			GasLimit: 1_000_000,                   // gas limit 1M
			GasPrice: big.NewInt(100_000_000_000), // gas price 100 gWEI

		},
		client, owner,
	)
	if err != nil {
		return err
	}

	// check if already deployed
	deterministicSender, err := scalableSigner{}.Sender(tx)
	if err != nil {
		return err
	}
	deployerAddr := crypto.CreateAddress(deterministicSender, 0)

	code, err := client.CodeAt(cliCtx.Context, deployerAddr, nil)
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
			nonce, err := client.PendingNonceAt(cliCtx.Context, deterministicSender)
			if err != nil {
				return err
			}
			fmt.Printf("deterministic addr = %s\nbalance = %s\n nonce = %d\n\n\n", deterministicSender, dsBalance.String(), nonce)
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
				owner.Hex(), deterministicSender,
			),
			func() (*types.Transaction, error) {
				err := client.SendTransaction(cliCtx.Context, tx)
				return tx, err
			},
		)
		if err != nil {
			return err
		}
		if deployerAddr != actualDeployerAddr {
			return fmt.Errorf("unexpcted deployer addr. Expected %s, actual: %s", deployerAddr, actualDeployerAddr)
		}
	}
	deployer, err := polygonzkevmdeployer.NewPolygonzkevmdeployer(deployerAddr, client)
	if err != nil {
		return err
	}
	// Sanity check
	actualOwner, err := deployer.Owner(nil)
	if err != nil {
		return err
	}
	if actualOwner != owner {
		return fmt.Errorf("owner was supposed to be %s but is %s", owner, actualOwner)
	}
	fmt.Println("PolygonZkEVMDeployer addr:", deployerAddr)

	// Deploy proxy admin
	fmt.Println("Deploying proxy admin...")
	salt := common.HexToHash("TODOooo") // TODO: input via config / CLI
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
		cliCtx, client, auth, deployer, deployerAddr, salt, proxyTx.Data(), dataCallAdmin,
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
	actualOwner, err = proxy.Owner(nil)
	if err != nil {
		return err
	}
	if actualOwner != walletAddr {
		return fmt.Errorf("owner was supposed to be %s but is %s", owner, actualOwner)
	}
	fmt.Println("proxy admin addr:", proxyAdminAddr)

	// Deploy implementation PolygonZkEVMBridge
	fmt.Println("Deploying PolygonZkEVMBridge implementation...")
	auth.GasLimit = 5500000 // gas limit with create are mess up D:
	_, bridgeTx, _, err := polygonzkevmbridgev2.DeployPolygonzkevmbridgev2(
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
	bridgeImplementationAddress, err := create2Deployment(
		cliCtx, client, auth, deployer, deployerAddr, salt, bridgeTx.Data(), nil,
		fmt.Sprintf(
			"Do you want to send the tx that will deploy the proxy (using the deployer)? Admin of the proxy will be set to %s",
			walletAddr,
		),
	)
	if err != nil {
		return err
	}
	auth.GasLimit = 0
	bridgeImplementation, err := polygonzkevmbridgev2.NewPolygonzkevmbridgev2(bridgeImplementationAddress, client)
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
	fmt.Println("bridge implementation addr:", bridgeImplementationAddress)

	// Deploy PolygonZkEVMTimelock
	fmt.Println("Deploying PolygonZkEVMTimelock...")
	currentNonce, err := client.NonceAt(cliCtx.Context, walletAddr, nil)
	if err != nil {
		return err
	}
	// Nonce globalExitRoot: currentNonce + 1 (deploy bridge proxy) + 1(impl globalExitRoot)
	// + 1 (deployTimelock) + 1 (transfer Ownership Admin) = +4
	nonceProxyGlobalExitRoot := currentNonce + 4
	// nonceProxyRollupManager :Nonce globalExitRoot + 1 (proxy globalExitRoot) + 1 (impl rollupManager) = +2
	nonceProxyRollupManager := nonceProxyGlobalExitRoot + 2
	precalculateGlobalExitRootAddress := crypto.CreateAddress(walletAddr, nonceProxyGlobalExitRoot)
	precalculateRollupManager := crypto.CreateAddress(walletAddr, nonceProxyRollupManager)
	fmt.Println("TODO: remove this log", precalculateGlobalExitRootAddress)

	// TODO: input via config / CLI
	minDelayTimelock := big.NewInt(10000)
	timelockAdmin := common.Address{}
	timelockAdmins := []common.Address{timelockAdmin}

	// Deploy deployer
	timelockAddr, err := sendTxWithConfirmation(
		cliCtx, client,
		fmt.Sprintf(
			"Do you want to send the tx that will deploy the PolygonZkEVMTimelock with admin %s? This timelock will be used for the Rollup Manager with precalculated address %s",
			timelockAdmin, precalculateRollupManager,
		),
		func() (*types.Transaction, error) {
			_, tx, _, err := polygonzkevmtimelock.DeployPolygonzkevmtimelock(
				auth, client, minDelayTimelock, timelockAdmins, timelockAdmins, timelockAdmin, precalculateRollupManager,
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
	// Transfer ownership of the proxyAdmin to timelock
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

	return nil
}

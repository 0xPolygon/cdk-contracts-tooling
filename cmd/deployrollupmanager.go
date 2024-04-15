package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/common/proxyadmin"
	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmdeployer"
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
				Required: false,
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
	// Deterministic address deployment
	owner := common.HexToAddress(cliCtx.String(ownerFlagName))
	txData, err := getPolygonZkEVMDeployerDeployTransactionData(owner)
	if err != nil {
		return err
	}
	tx := types.NewTransaction(
		0,                           // nonce
		common.Address{},            // to: 0x00...0 means deploy
		big.NewInt(0),               // amount
		1_000_000,                   // gas limit 1M
		big.NewInt(100_000_000_000), // gas price 100 gWEI
		txData,
	)
	tx, err = tx.WithSignature(scalableSigner{}, nil)
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
		fmt.Println("Deployer was already deployed")
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
					deterministicSender, walletAddr, neededFunds.String(),
				),
				func() (*types.Transaction, common.Address, error) {
					nonce, err := client.PendingNonceAt(cliCtx.Context, walletAddr)
					if err != nil {
						return nil, common.Address{}, err
					}
					fundTx := types.NewTransaction(nonce, deterministicSender, neededFunds, 21000, nil, nil)
					fundTx, err = auth.Signer(walletAddr, fundTx)
					if err != nil {
						return nil, common.Address{}, err
					}
					err = client.SendTransaction(cliCtx.Context, fundTx)
					return fundTx, common.Address{}, err
				},
			)
		}
		// Deploy deployer
		deployerAddr, err = sendTxWithConfirmation(
			cliCtx, client,
			fmt.Sprintf(
				"Do you want to send the tx that will deploy the PolygonZkEVMDeployer with owner %s? Note that this tx will be sent from %s",
				owner.Hex(), deterministicSender,
			),
			func() (*types.Transaction, common.Address, error) {
				return nil, common.Address{}, err
			},
		)
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

	// Deploy proxy admin
	salt := common.HexToHash("TODO")
	dataCallAdmin := common.Hex2Bytes("TODO")
	deployTx, err := getProxyAdminDeployTransactionData()
	proxyAdminAddr, err := create2Deployment(
		cliCtx, client, auth, deployer, deployerAddr, salt, deployTx, dataCallAdmin,
		"Do you want to send the tx that will deploy the proxy (using the deployer)?",
	)
	fmt.Println("TODO: remove this log", proxyAdminAddr)

	// const proxyAdminFactory = await ethers.getContractFactory(
	//     "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol:ProxyAdmin",
	//     deployer
	// );
	// const deployTransactionAdmin = (await proxyAdminFactory.getDeployTransaction()).data;
	// const dataCallAdmin = proxyAdminFactory.interface.encodeFunctionData("transferOwnership", [deployer.address]);
	// const [proxyAdminAddress, isProxyAdminDeployed] = await create2Deployment(
	//     zkEVMDeployerContract,
	//     salt,
	//     deployTransactionAdmin,
	//     dataCallAdmin,
	//     deployer
	// );

	return nil
}

func getPolygonZkEVMDeployerDeployTransactionData(owner common.Address) ([]byte, error) {
	abi, err := polygonzkevmdeployer.PolygonzkevmdeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if abi == nil {
		return nil, errors.New("GetABI returned nil")
	}
	return abi.Pack("", owner)
}

func getProxyAdminDeployTransactionData() ([]byte, error) {
	abi, err := proxyadmin.ProxyadminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if abi == nil {
		return nil, errors.New("GetABI returned nil")
	}
	return abi.Pack("")
}

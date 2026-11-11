// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayerbridgel2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Iagglayerbridgel2MetaData contains all meta data concerning the Iagglayerbridgel2 contract.
var Iagglayerbridgel2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpectedLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLBTLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeafType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeavesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubtreeFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Iagglayerbridgel2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iagglayerbridgel2MetaData.ABI instead.
var Iagglayerbridgel2ABI = Iagglayerbridgel2MetaData.ABI

// Iagglayerbridgel2 is an auto generated Go binding around an Ethereum contract.
type Iagglayerbridgel2 struct {
	Iagglayerbridgel2Caller     // Read-only binding to the contract
	Iagglayerbridgel2Transactor // Write-only binding to the contract
	Iagglayerbridgel2Filterer   // Log filterer for contract events
}

// Iagglayerbridgel2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iagglayerbridgel2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayerbridgel2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iagglayerbridgel2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayerbridgel2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iagglayerbridgel2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayerbridgel2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iagglayerbridgel2Session struct {
	Contract     *Iagglayerbridgel2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Iagglayerbridgel2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iagglayerbridgel2CallerSession struct {
	Contract *Iagglayerbridgel2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Iagglayerbridgel2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iagglayerbridgel2TransactorSession struct {
	Contract     *Iagglayerbridgel2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Iagglayerbridgel2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iagglayerbridgel2Raw struct {
	Contract *Iagglayerbridgel2 // Generic contract binding to access the raw methods on
}

// Iagglayerbridgel2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iagglayerbridgel2CallerRaw struct {
	Contract *Iagglayerbridgel2Caller // Generic read-only contract binding to access the raw methods on
}

// Iagglayerbridgel2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iagglayerbridgel2TransactorRaw struct {
	Contract *Iagglayerbridgel2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayerbridgel2 creates a new instance of Iagglayerbridgel2, bound to a specific deployed contract.
func NewIagglayerbridgel2(address common.Address, backend bind.ContractBackend) (*Iagglayerbridgel2, error) {
	contract, err := bindIagglayerbridgel2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayerbridgel2{Iagglayerbridgel2Caller: Iagglayerbridgel2Caller{contract: contract}, Iagglayerbridgel2Transactor: Iagglayerbridgel2Transactor{contract: contract}, Iagglayerbridgel2Filterer: Iagglayerbridgel2Filterer{contract: contract}}, nil
}

// NewIagglayerbridgel2Caller creates a new read-only instance of Iagglayerbridgel2, bound to a specific deployed contract.
func NewIagglayerbridgel2Caller(address common.Address, caller bind.ContractCaller) (*Iagglayerbridgel2Caller, error) {
	contract, err := bindIagglayerbridgel2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iagglayerbridgel2Caller{contract: contract}, nil
}

// NewIagglayerbridgel2Transactor creates a new write-only instance of Iagglayerbridgel2, bound to a specific deployed contract.
func NewIagglayerbridgel2Transactor(address common.Address, transactor bind.ContractTransactor) (*Iagglayerbridgel2Transactor, error) {
	contract, err := bindIagglayerbridgel2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iagglayerbridgel2Transactor{contract: contract}, nil
}

// NewIagglayerbridgel2Filterer creates a new log filterer instance of Iagglayerbridgel2, bound to a specific deployed contract.
func NewIagglayerbridgel2Filterer(address common.Address, filterer bind.ContractFilterer) (*Iagglayerbridgel2Filterer, error) {
	contract, err := bindIagglayerbridgel2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iagglayerbridgel2Filterer{contract: contract}, nil
}

// bindIagglayerbridgel2 binds a generic wrapper to an already deployed contract.
func bindIagglayerbridgel2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iagglayerbridgel2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayerbridgel2 *Iagglayerbridgel2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayerbridgel2.Contract.Iagglayerbridgel2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayerbridgel2 *Iagglayerbridgel2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Iagglayerbridgel2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayerbridgel2 *Iagglayerbridgel2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Iagglayerbridgel2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayerbridgel2 *Iagglayerbridgel2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayerbridgel2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.contract.Transact(opts, method, params...)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Caller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iagglayerbridgel2.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) GetProxiedTokensManager() (common.Address, error) {
	return _Iagglayerbridgel2.Contract.GetProxiedTokensManager(&_Iagglayerbridgel2.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2CallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Iagglayerbridgel2.Contract.GetProxiedTokensManager(&_Iagglayerbridgel2.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Caller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Iagglayerbridgel2.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Iagglayerbridgel2.Contract.GetTokenMetadata(&_Iagglayerbridgel2.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Iagglayerbridgel2 *Iagglayerbridgel2CallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Iagglayerbridgel2.Contract.GetTokenMetadata(&_Iagglayerbridgel2.CallOpts, token)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Caller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iagglayerbridgel2.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Iagglayerbridgel2.Contract.GetWrappedTokenBridgeImplementation(&_Iagglayerbridgel2.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2CallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Iagglayerbridgel2.Contract.GetWrappedTokenBridgeImplementation(&_Iagglayerbridgel2.CallOpts)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Caller) WrappedTokenToTokenInfo(opts *bind.CallOpts, destinationAddress common.Address) (uint32, common.Address, error) {
	var out []interface{}
	err := _Iagglayerbridgel2.contract.Call(opts, &out, "wrappedTokenToTokenInfo", destinationAddress)

	if err != nil {
		return *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Iagglayerbridgel2.Contract.WrappedTokenToTokenInfo(&_Iagglayerbridgel2.CallOpts, destinationAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address destinationAddress) view returns(uint32, address)
func (_Iagglayerbridgel2 *Iagglayerbridgel2CallerSession) WrappedTokenToTokenInfo(destinationAddress common.Address) (uint32, common.Address, error) {
	return _Iagglayerbridgel2.Contract.WrappedTokenToTokenInfo(&_Iagglayerbridgel2.CallOpts, destinationAddress)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) ActivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ActivateEmergencyState(&_Iagglayerbridgel2.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ActivateEmergencyState(&_Iagglayerbridgel2.TransactOpts)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeAsset(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeAsset(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeMessage(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeMessage(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeMessageWETH(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.BridgeMessageWETH(&_Iagglayerbridgel2.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ClaimAsset(&_Iagglayerbridgel2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ClaimAsset(&_Iagglayerbridgel2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ClaimMessage(&_Iagglayerbridgel2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.ClaimMessage(&_Iagglayerbridgel2.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.DeactivateEmergencyState(&_Iagglayerbridgel2.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.DeactivateEmergencyState(&_Iagglayerbridgel2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Initialize(&_Iagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Initialize(&_Iagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) Initialize0(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "initialize0", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Initialize0(&_Iagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) Initialize0(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.Initialize0(&_Iagglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Transactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerbridgel2.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2Session) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.UpdateGlobalExitRoot(&_Iagglayerbridgel2.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Iagglayerbridgel2 *Iagglayerbridgel2TransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Iagglayerbridgel2.Contract.UpdateGlobalExitRoot(&_Iagglayerbridgel2.TransactOpts)
}

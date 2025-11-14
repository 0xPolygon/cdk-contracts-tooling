// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iinitializeragglayerbridgel2

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

// Iinitializeragglayerbridgel2MetaData contains all meta data concerning the Iinitializeragglayerbridgel2 contract.
var Iinitializeragglayerbridgel2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iinitializeragglayerbridgel2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iinitializeragglayerbridgel2MetaData.ABI instead.
var Iinitializeragglayerbridgel2ABI = Iinitializeragglayerbridgel2MetaData.ABI

// Iinitializeragglayerbridgel2 is an auto generated Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2 struct {
	Iinitializeragglayerbridgel2Caller     // Read-only binding to the contract
	Iinitializeragglayerbridgel2Transactor // Write-only binding to the contract
	Iinitializeragglayerbridgel2Filterer   // Log filterer for contract events
}

// Iinitializeragglayerbridgel2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iinitializeragglayerbridgel2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iinitializeragglayerbridgel2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iinitializeragglayerbridgel2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iinitializeragglayerbridgel2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iinitializeragglayerbridgel2Session struct {
	Contract     *Iinitializeragglayerbridgel2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Iinitializeragglayerbridgel2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iinitializeragglayerbridgel2CallerSession struct {
	Contract *Iinitializeragglayerbridgel2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// Iinitializeragglayerbridgel2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iinitializeragglayerbridgel2TransactorSession struct {
	Contract     *Iinitializeragglayerbridgel2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// Iinitializeragglayerbridgel2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2Raw struct {
	Contract *Iinitializeragglayerbridgel2 // Generic contract binding to access the raw methods on
}

// Iinitializeragglayerbridgel2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2CallerRaw struct {
	Contract *Iinitializeragglayerbridgel2Caller // Generic read-only contract binding to access the raw methods on
}

// Iinitializeragglayerbridgel2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iinitializeragglayerbridgel2TransactorRaw struct {
	Contract *Iinitializeragglayerbridgel2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIinitializeragglayerbridgel2 creates a new instance of Iinitializeragglayerbridgel2, bound to a specific deployed contract.
func NewIinitializeragglayerbridgel2(address common.Address, backend bind.ContractBackend) (*Iinitializeragglayerbridgel2, error) {
	contract, err := bindIinitializeragglayerbridgel2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iinitializeragglayerbridgel2{Iinitializeragglayerbridgel2Caller: Iinitializeragglayerbridgel2Caller{contract: contract}, Iinitializeragglayerbridgel2Transactor: Iinitializeragglayerbridgel2Transactor{contract: contract}, Iinitializeragglayerbridgel2Filterer: Iinitializeragglayerbridgel2Filterer{contract: contract}}, nil
}

// NewIinitializeragglayerbridgel2Caller creates a new read-only instance of Iinitializeragglayerbridgel2, bound to a specific deployed contract.
func NewIinitializeragglayerbridgel2Caller(address common.Address, caller bind.ContractCaller) (*Iinitializeragglayerbridgel2Caller, error) {
	contract, err := bindIinitializeragglayerbridgel2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iinitializeragglayerbridgel2Caller{contract: contract}, nil
}

// NewIinitializeragglayerbridgel2Transactor creates a new write-only instance of Iinitializeragglayerbridgel2, bound to a specific deployed contract.
func NewIinitializeragglayerbridgel2Transactor(address common.Address, transactor bind.ContractTransactor) (*Iinitializeragglayerbridgel2Transactor, error) {
	contract, err := bindIinitializeragglayerbridgel2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iinitializeragglayerbridgel2Transactor{contract: contract}, nil
}

// NewIinitializeragglayerbridgel2Filterer creates a new log filterer instance of Iinitializeragglayerbridgel2, bound to a specific deployed contract.
func NewIinitializeragglayerbridgel2Filterer(address common.Address, filterer bind.ContractFilterer) (*Iinitializeragglayerbridgel2Filterer, error) {
	contract, err := bindIinitializeragglayerbridgel2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iinitializeragglayerbridgel2Filterer{contract: contract}, nil
}

// bindIinitializeragglayerbridgel2 binds a generic wrapper to an already deployed contract.
func bindIinitializeragglayerbridgel2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iinitializeragglayerbridgel2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iinitializeragglayerbridgel2.Contract.Iinitializeragglayerbridgel2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.Iinitializeragglayerbridgel2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.Iinitializeragglayerbridgel2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iinitializeragglayerbridgel2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.contract.Transact(opts, method, params...)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2Transactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2Session) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.Initialize(&_Iinitializeragglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Iinitializeragglayerbridgel2 *Iinitializeragglayerbridgel2TransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Iinitializeragglayerbridgel2.Contract.Initialize(&_Iinitializeragglayerbridgel2.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

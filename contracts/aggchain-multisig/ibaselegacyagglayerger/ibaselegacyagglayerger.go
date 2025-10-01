// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibaselegacyagglayerger

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

// IbaselegacyagglayergerMetaData contains all meta data concerning the Ibaselegacyagglayerger contract.
var IbaselegacyagglayergerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IbaselegacyagglayergerABI is the input ABI used to generate the binding from.
// Deprecated: Use IbaselegacyagglayergerMetaData.ABI instead.
var IbaselegacyagglayergerABI = IbaselegacyagglayergerMetaData.ABI

// Ibaselegacyagglayerger is an auto generated Go binding around an Ethereum contract.
type Ibaselegacyagglayerger struct {
	IbaselegacyagglayergerCaller     // Read-only binding to the contract
	IbaselegacyagglayergerTransactor // Write-only binding to the contract
	IbaselegacyagglayergerFilterer   // Log filterer for contract events
}

// IbaselegacyagglayergerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbaselegacyagglayergerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbaselegacyagglayergerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbaselegacyagglayergerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbaselegacyagglayergerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbaselegacyagglayergerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbaselegacyagglayergerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbaselegacyagglayergerSession struct {
	Contract     *Ibaselegacyagglayerger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IbaselegacyagglayergerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbaselegacyagglayergerCallerSession struct {
	Contract *IbaselegacyagglayergerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IbaselegacyagglayergerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbaselegacyagglayergerTransactorSession struct {
	Contract     *IbaselegacyagglayergerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IbaselegacyagglayergerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbaselegacyagglayergerRaw struct {
	Contract *Ibaselegacyagglayerger // Generic contract binding to access the raw methods on
}

// IbaselegacyagglayergerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbaselegacyagglayergerCallerRaw struct {
	Contract *IbaselegacyagglayergerCaller // Generic read-only contract binding to access the raw methods on
}

// IbaselegacyagglayergerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbaselegacyagglayergerTransactorRaw struct {
	Contract *IbaselegacyagglayergerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbaselegacyagglayerger creates a new instance of Ibaselegacyagglayerger, bound to a specific deployed contract.
func NewIbaselegacyagglayerger(address common.Address, backend bind.ContractBackend) (*Ibaselegacyagglayerger, error) {
	contract, err := bindIbaselegacyagglayerger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibaselegacyagglayerger{IbaselegacyagglayergerCaller: IbaselegacyagglayergerCaller{contract: contract}, IbaselegacyagglayergerTransactor: IbaselegacyagglayergerTransactor{contract: contract}, IbaselegacyagglayergerFilterer: IbaselegacyagglayergerFilterer{contract: contract}}, nil
}

// NewIbaselegacyagglayergerCaller creates a new read-only instance of Ibaselegacyagglayerger, bound to a specific deployed contract.
func NewIbaselegacyagglayergerCaller(address common.Address, caller bind.ContractCaller) (*IbaselegacyagglayergerCaller, error) {
	contract, err := bindIbaselegacyagglayerger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbaselegacyagglayergerCaller{contract: contract}, nil
}

// NewIbaselegacyagglayergerTransactor creates a new write-only instance of Ibaselegacyagglayerger, bound to a specific deployed contract.
func NewIbaselegacyagglayergerTransactor(address common.Address, transactor bind.ContractTransactor) (*IbaselegacyagglayergerTransactor, error) {
	contract, err := bindIbaselegacyagglayerger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbaselegacyagglayergerTransactor{contract: contract}, nil
}

// NewIbaselegacyagglayergerFilterer creates a new log filterer instance of Ibaselegacyagglayerger, bound to a specific deployed contract.
func NewIbaselegacyagglayergerFilterer(address common.Address, filterer bind.ContractFilterer) (*IbaselegacyagglayergerFilterer, error) {
	contract, err := bindIbaselegacyagglayerger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbaselegacyagglayergerFilterer{contract: contract}, nil
}

// bindIbaselegacyagglayerger binds a generic wrapper to an already deployed contract.
func bindIbaselegacyagglayerger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbaselegacyagglayergerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibaselegacyagglayerger.Contract.IbaselegacyagglayergerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.IbaselegacyagglayergerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.IbaselegacyagglayergerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibaselegacyagglayerger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.contract.Transact(opts, method, params...)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.GlobalExitRootMap(&_Ibaselegacyagglayerger.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.GlobalExitRootMap(&_Ibaselegacyagglayerger.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.UpdateExitRoot(&_Ibaselegacyagglayerger.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibaselegacyagglayerger *IbaselegacyagglayergerTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibaselegacyagglayerger.Contract.UpdateExitRoot(&_Ibaselegacyagglayerger.TransactOpts, newRollupExitRoot)
}

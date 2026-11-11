// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayerger

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

// IagglayergerMetaData contains all meta data concerning the Iagglayerger contract.
var IagglayergerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IagglayergerABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergerMetaData.ABI instead.
var IagglayergerABI = IagglayergerMetaData.ABI

// Iagglayerger is an auto generated Go binding around an Ethereum contract.
type Iagglayerger struct {
	IagglayergerCaller     // Read-only binding to the contract
	IagglayergerTransactor // Write-only binding to the contract
	IagglayergerFilterer   // Log filterer for contract events
}

// IagglayergerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergerSession struct {
	Contract     *Iagglayerger     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IagglayergerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergerCallerSession struct {
	Contract *IagglayergerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IagglayergerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergerTransactorSession struct {
	Contract     *IagglayergerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IagglayergerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergerRaw struct {
	Contract *Iagglayerger // Generic contract binding to access the raw methods on
}

// IagglayergerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergerCallerRaw struct {
	Contract *IagglayergerCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergerTransactorRaw struct {
	Contract *IagglayergerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayerger creates a new instance of Iagglayerger, bound to a specific deployed contract.
func NewIagglayerger(address common.Address, backend bind.ContractBackend) (*Iagglayerger, error) {
	contract, err := bindIagglayerger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayerger{IagglayergerCaller: IagglayergerCaller{contract: contract}, IagglayergerTransactor: IagglayergerTransactor{contract: contract}, IagglayergerFilterer: IagglayergerFilterer{contract: contract}}, nil
}

// NewIagglayergerCaller creates a new read-only instance of Iagglayerger, bound to a specific deployed contract.
func NewIagglayergerCaller(address common.Address, caller bind.ContractCaller) (*IagglayergerCaller, error) {
	contract, err := bindIagglayerger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergerCaller{contract: contract}, nil
}

// NewIagglayergerTransactor creates a new write-only instance of Iagglayerger, bound to a specific deployed contract.
func NewIagglayergerTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergerTransactor, error) {
	contract, err := bindIagglayerger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergerTransactor{contract: contract}, nil
}

// NewIagglayergerFilterer creates a new log filterer instance of Iagglayerger, bound to a specific deployed contract.
func NewIagglayergerFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergerFilterer, error) {
	contract, err := bindIagglayerger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergerFilterer{contract: contract}, nil
}

// bindIagglayerger binds a generic wrapper to an already deployed contract.
func bindIagglayerger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayerger *IagglayergerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayerger.Contract.IagglayergerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayerger *IagglayergerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerger.Contract.IagglayergerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayerger *IagglayergerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayerger.Contract.IagglayergerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayerger *IagglayergerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayerger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayerger *IagglayergerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayerger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayerger *IagglayergerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayerger.Contract.contract.Transact(opts, method, params...)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayerger.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Iagglayerger.Contract.GetLastGlobalExitRoot(&_Iagglayerger.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Iagglayerger.Contract.GetLastGlobalExitRoot(&_Iagglayerger.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayerger.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerSession) GetRoot() ([32]byte, error) {
	return _Iagglayerger.Contract.GetRoot(&_Iagglayerger.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Iagglayerger *IagglayergerCallerSession) GetRoot() ([32]byte, error) {
	return _Iagglayerger.Contract.GetRoot(&_Iagglayerger.CallOpts)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Iagglayerger *IagglayergerCaller) L1InfoRootMap(opts *bind.CallOpts, depositCount uint32) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayerger.contract.Call(opts, &out, "l1InfoRootMap", depositCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Iagglayerger *IagglayergerSession) L1InfoRootMap(depositCount uint32) ([32]byte, error) {
	return _Iagglayerger.Contract.L1InfoRootMap(&_Iagglayerger.CallOpts, depositCount)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Iagglayerger *IagglayergerCallerSession) L1InfoRootMap(depositCount uint32) ([32]byte, error) {
	return _Iagglayerger.Contract.L1InfoRootMap(&_Iagglayerger.CallOpts, depositCount)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayerger *IagglayergerTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayerger *IagglayergerSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.Contract.GlobalExitRootMap(&_Iagglayerger.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayerger *IagglayergerTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.Contract.GlobalExitRootMap(&_Iagglayerger.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayerger *IagglayergerTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayerger *IagglayergerSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.Contract.UpdateExitRoot(&_Iagglayerger.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayerger *IagglayergerTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayerger.Contract.UpdateExitRoot(&_Iagglayerger.TransactOpts, newRollupExitRoot)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibasepolygonzkevmglobalexitroot

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

// IbasepolygonzkevmglobalexitrootMetaData contains all meta data concerning the Ibasepolygonzkevmglobalexitroot contract.
var IbasepolygonzkevmglobalexitrootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IbasepolygonzkevmglobalexitrootABI is the input ABI used to generate the binding from.
// Deprecated: Use IbasepolygonzkevmglobalexitrootMetaData.ABI instead.
var IbasepolygonzkevmglobalexitrootABI = IbasepolygonzkevmglobalexitrootMetaData.ABI

// Ibasepolygonzkevmglobalexitroot is an auto generated Go binding around an Ethereum contract.
type Ibasepolygonzkevmglobalexitroot struct {
	IbasepolygonzkevmglobalexitrootCaller     // Read-only binding to the contract
	IbasepolygonzkevmglobalexitrootTransactor // Write-only binding to the contract
	IbasepolygonzkevmglobalexitrootFilterer   // Log filterer for contract events
}

// IbasepolygonzkevmglobalexitrootCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbasepolygonzkevmglobalexitrootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbasepolygonzkevmglobalexitrootSession struct {
	Contract     *Ibasepolygonzkevmglobalexitroot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbasepolygonzkevmglobalexitrootCallerSession struct {
	Contract *IbasepolygonzkevmglobalexitrootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbasepolygonzkevmglobalexitrootTransactorSession struct {
	Contract     *IbasepolygonzkevmglobalexitrootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootRaw struct {
	Contract *Ibasepolygonzkevmglobalexitroot // Generic contract binding to access the raw methods on
}

// IbasepolygonzkevmglobalexitrootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootCallerRaw struct {
	Contract *IbasepolygonzkevmglobalexitrootCaller // Generic read-only contract binding to access the raw methods on
}

// IbasepolygonzkevmglobalexitrootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootTransactorRaw struct {
	Contract *IbasepolygonzkevmglobalexitrootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbasepolygonzkevmglobalexitroot creates a new instance of Ibasepolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitroot(address common.Address, backend bind.ContractBackend) (*Ibasepolygonzkevmglobalexitroot, error) {
	contract, err := bindIbasepolygonzkevmglobalexitroot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibasepolygonzkevmglobalexitroot{IbasepolygonzkevmglobalexitrootCaller: IbasepolygonzkevmglobalexitrootCaller{contract: contract}, IbasepolygonzkevmglobalexitrootTransactor: IbasepolygonzkevmglobalexitrootTransactor{contract: contract}, IbasepolygonzkevmglobalexitrootFilterer: IbasepolygonzkevmglobalexitrootFilterer{contract: contract}}, nil
}

// NewIbasepolygonzkevmglobalexitrootCaller creates a new read-only instance of Ibasepolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootCaller(address common.Address, caller bind.ContractCaller) (*IbasepolygonzkevmglobalexitrootCaller, error) {
	contract, err := bindIbasepolygonzkevmglobalexitroot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootCaller{contract: contract}, nil
}

// NewIbasepolygonzkevmglobalexitrootTransactor creates a new write-only instance of Ibasepolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootTransactor(address common.Address, transactor bind.ContractTransactor) (*IbasepolygonzkevmglobalexitrootTransactor, error) {
	contract, err := bindIbasepolygonzkevmglobalexitroot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootTransactor{contract: contract}, nil
}

// NewIbasepolygonzkevmglobalexitrootFilterer creates a new log filterer instance of Ibasepolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootFilterer(address common.Address, filterer bind.ContractFilterer) (*IbasepolygonzkevmglobalexitrootFilterer, error) {
	contract, err := bindIbasepolygonzkevmglobalexitroot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootFilterer{contract: contract}, nil
}

// bindIbasepolygonzkevmglobalexitroot binds a generic wrapper to an already deployed contract.
func bindIbasepolygonzkevmglobalexitroot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbasepolygonzkevmglobalexitrootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibasepolygonzkevmglobalexitroot.Contract.IbasepolygonzkevmglobalexitrootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.IbasepolygonzkevmglobalexitrootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.IbasepolygonzkevmglobalexitrootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibasepolygonzkevmglobalexitroot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.contract.Transact(opts, method, params...)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.GlobalExitRootMap(&_Ibasepolygonzkevmglobalexitroot.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.GlobalExitRootMap(&_Ibasepolygonzkevmglobalexitroot.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.UpdateExitRoot(&_Ibasepolygonzkevmglobalexitroot.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitroot *IbasepolygonzkevmglobalexitrootTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitroot.Contract.UpdateExitRoot(&_Ibasepolygonzkevmglobalexitroot.TransactOpts, newRollupExitRoot)
}

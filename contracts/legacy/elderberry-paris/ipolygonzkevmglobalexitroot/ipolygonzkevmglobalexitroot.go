// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmglobalexitroot

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

// IpolygonzkevmglobalexitrootMetaData contains all meta data concerning the Ipolygonzkevmglobalexitroot contract.
var IpolygonzkevmglobalexitrootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonzkevmglobalexitrootABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonzkevmglobalexitrootMetaData.ABI instead.
var IpolygonzkevmglobalexitrootABI = IpolygonzkevmglobalexitrootMetaData.ABI

// Ipolygonzkevmglobalexitroot is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitroot struct {
	IpolygonzkevmglobalexitrootCaller     // Read-only binding to the contract
	IpolygonzkevmglobalexitrootTransactor // Write-only binding to the contract
	IpolygonzkevmglobalexitrootFilterer   // Log filterer for contract events
}

// IpolygonzkevmglobalexitrootCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonzkevmglobalexitrootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmglobalexitrootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonzkevmglobalexitrootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmglobalexitrootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonzkevmglobalexitrootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmglobalexitrootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonzkevmglobalexitrootSession struct {
	Contract     *Ipolygonzkevmglobalexitroot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IpolygonzkevmglobalexitrootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonzkevmglobalexitrootCallerSession struct {
	Contract *IpolygonzkevmglobalexitrootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IpolygonzkevmglobalexitrootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonzkevmglobalexitrootTransactorSession struct {
	Contract     *IpolygonzkevmglobalexitrootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IpolygonzkevmglobalexitrootRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonzkevmglobalexitrootRaw struct {
	Contract *Ipolygonzkevmglobalexitroot // Generic contract binding to access the raw methods on
}

// IpolygonzkevmglobalexitrootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonzkevmglobalexitrootCallerRaw struct {
	Contract *IpolygonzkevmglobalexitrootCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonzkevmglobalexitrootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonzkevmglobalexitrootTransactorRaw struct {
	Contract *IpolygonzkevmglobalexitrootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmglobalexitroot creates a new instance of Ipolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitroot(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmglobalexitroot, error) {
	contract, err := bindIpolygonzkevmglobalexitroot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmglobalexitroot{IpolygonzkevmglobalexitrootCaller: IpolygonzkevmglobalexitrootCaller{contract: contract}, IpolygonzkevmglobalexitrootTransactor: IpolygonzkevmglobalexitrootTransactor{contract: contract}, IpolygonzkevmglobalexitrootFilterer: IpolygonzkevmglobalexitrootFilterer{contract: contract}}, nil
}

// NewIpolygonzkevmglobalexitrootCaller creates a new read-only instance of Ipolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootCaller(address common.Address, caller bind.ContractCaller) (*IpolygonzkevmglobalexitrootCaller, error) {
	contract, err := bindIpolygonzkevmglobalexitroot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmglobalexitrootCaller{contract: contract}, nil
}

// NewIpolygonzkevmglobalexitrootTransactor creates a new write-only instance of Ipolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonzkevmglobalexitrootTransactor, error) {
	contract, err := bindIpolygonzkevmglobalexitroot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmglobalexitrootTransactor{contract: contract}, nil
}

// NewIpolygonzkevmglobalexitrootFilterer creates a new log filterer instance of Ipolygonzkevmglobalexitroot, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonzkevmglobalexitrootFilterer, error) {
	contract, err := bindIpolygonzkevmglobalexitroot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmglobalexitrootFilterer{contract: contract}, nil
}

// bindIpolygonzkevmglobalexitroot binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmglobalexitroot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonzkevmglobalexitrootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmglobalexitroot.Contract.IpolygonzkevmglobalexitrootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.IpolygonzkevmglobalexitrootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.IpolygonzkevmglobalexitrootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmglobalexitroot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.contract.Transact(opts, method, params...)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonzkevmglobalexitroot.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.GetLastGlobalExitRoot(&_Ipolygonzkevmglobalexitroot.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.GetLastGlobalExitRoot(&_Ipolygonzkevmglobalexitroot.CallOpts)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.GlobalExitRootMap(&_Ipolygonzkevmglobalexitroot.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.GlobalExitRootMap(&_Ipolygonzkevmglobalexitroot.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.UpdateExitRoot(&_Ipolygonzkevmglobalexitroot.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitroot *IpolygonzkevmglobalexitrootTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitroot.Contract.UpdateExitRoot(&_Ipolygonzkevmglobalexitroot.TransactOpts, newRollupExitRoot)
}

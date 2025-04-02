// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmglobalexitrootv2

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

// Ipolygonzkevmglobalexitrootv2MetaData contains all meta data concerning the Ipolygonzkevmglobalexitrootv2 contract.
var Ipolygonzkevmglobalexitrootv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ipolygonzkevmglobalexitrootv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ipolygonzkevmglobalexitrootv2MetaData.ABI instead.
var Ipolygonzkevmglobalexitrootv2ABI = Ipolygonzkevmglobalexitrootv2MetaData.ABI

// Ipolygonzkevmglobalexitrootv2 is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2 struct {
	Ipolygonzkevmglobalexitrootv2Caller     // Read-only binding to the contract
	Ipolygonzkevmglobalexitrootv2Transactor // Write-only binding to the contract
	Ipolygonzkevmglobalexitrootv2Filterer   // Log filterer for contract events
}

// Ipolygonzkevmglobalexitrootv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmglobalexitrootv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmglobalexitrootv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ipolygonzkevmglobalexitrootv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonzkevmglobalexitrootv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ipolygonzkevmglobalexitrootv2Session struct {
	Contract     *Ipolygonzkevmglobalexitrootv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Ipolygonzkevmglobalexitrootv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ipolygonzkevmglobalexitrootv2CallerSession struct {
	Contract *Ipolygonzkevmglobalexitrootv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// Ipolygonzkevmglobalexitrootv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ipolygonzkevmglobalexitrootv2TransactorSession struct {
	Contract     *Ipolygonzkevmglobalexitrootv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// Ipolygonzkevmglobalexitrootv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2Raw struct {
	Contract *Ipolygonzkevmglobalexitrootv2 // Generic contract binding to access the raw methods on
}

// Ipolygonzkevmglobalexitrootv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2CallerRaw struct {
	Contract *Ipolygonzkevmglobalexitrootv2Caller // Generic read-only contract binding to access the raw methods on
}

// Ipolygonzkevmglobalexitrootv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ipolygonzkevmglobalexitrootv2TransactorRaw struct {
	Contract *Ipolygonzkevmglobalexitrootv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmglobalexitrootv2 creates a new instance of Ipolygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootv2(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmglobalexitrootv2, error) {
	contract, err := bindIpolygonzkevmglobalexitrootv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmglobalexitrootv2{Ipolygonzkevmglobalexitrootv2Caller: Ipolygonzkevmglobalexitrootv2Caller{contract: contract}, Ipolygonzkevmglobalexitrootv2Transactor: Ipolygonzkevmglobalexitrootv2Transactor{contract: contract}, Ipolygonzkevmglobalexitrootv2Filterer: Ipolygonzkevmglobalexitrootv2Filterer{contract: contract}}, nil
}

// NewIpolygonzkevmglobalexitrootv2Caller creates a new read-only instance of Ipolygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootv2Caller(address common.Address, caller bind.ContractCaller) (*Ipolygonzkevmglobalexitrootv2Caller, error) {
	contract, err := bindIpolygonzkevmglobalexitrootv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmglobalexitrootv2Caller{contract: contract}, nil
}

// NewIpolygonzkevmglobalexitrootv2Transactor creates a new write-only instance of Ipolygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Ipolygonzkevmglobalexitrootv2Transactor, error) {
	contract, err := bindIpolygonzkevmglobalexitrootv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmglobalexitrootv2Transactor{contract: contract}, nil
}

// NewIpolygonzkevmglobalexitrootv2Filterer creates a new log filterer instance of Ipolygonzkevmglobalexitrootv2, bound to a specific deployed contract.
func NewIpolygonzkevmglobalexitrootv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Ipolygonzkevmglobalexitrootv2Filterer, error) {
	contract, err := bindIpolygonzkevmglobalexitrootv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmglobalexitrootv2Filterer{contract: contract}, nil
}

// bindIpolygonzkevmglobalexitrootv2 binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmglobalexitrootv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ipolygonzkevmglobalexitrootv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmglobalexitrootv2.Contract.Ipolygonzkevmglobalexitrootv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.Ipolygonzkevmglobalexitrootv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.Ipolygonzkevmglobalexitrootv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmglobalexitrootv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.contract.Transact(opts, method, params...)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Caller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonzkevmglobalexitrootv2.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Session) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GetLastGlobalExitRoot(&_Ipolygonzkevmglobalexitrootv2.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2CallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GetLastGlobalExitRoot(&_Ipolygonzkevmglobalexitrootv2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Caller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonzkevmglobalexitrootv2.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Session) GetRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GetRoot(&_Ipolygonzkevmglobalexitrootv2.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2CallerSession) GetRoot() ([32]byte, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GetRoot(&_Ipolygonzkevmglobalexitrootv2.CallOpts)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Transactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Session) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GlobalExitRootMap(&_Ipolygonzkevmglobalexitrootv2.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2TransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.GlobalExitRootMap(&_Ipolygonzkevmglobalexitrootv2.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2Session) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.UpdateExitRoot(&_Ipolygonzkevmglobalexitrootv2.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ipolygonzkevmglobalexitrootv2 *Ipolygonzkevmglobalexitrootv2TransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonzkevmglobalexitrootv2.Contract.UpdateExitRoot(&_Ipolygonzkevmglobalexitrootv2.TransactOpts, newRollupExitRoot)
}

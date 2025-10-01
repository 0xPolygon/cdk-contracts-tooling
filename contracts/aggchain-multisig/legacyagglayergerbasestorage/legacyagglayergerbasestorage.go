// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legacyagglayergerbasestorage

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

// LegacyagglayergerbasestorageMetaData contains all meta data concerning the Legacyagglayergerbasestorage contract.
var LegacyagglayergerbasestorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"l1InfoRootMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LegacyagglayergerbasestorageABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyagglayergerbasestorageMetaData.ABI instead.
var LegacyagglayergerbasestorageABI = LegacyagglayergerbasestorageMetaData.ABI

// Legacyagglayergerbasestorage is an auto generated Go binding around an Ethereum contract.
type Legacyagglayergerbasestorage struct {
	LegacyagglayergerbasestorageCaller     // Read-only binding to the contract
	LegacyagglayergerbasestorageTransactor // Write-only binding to the contract
	LegacyagglayergerbasestorageFilterer   // Log filterer for contract events
}

// LegacyagglayergerbasestorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyagglayergerbasestorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyagglayergerbasestorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyagglayergerbasestorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyagglayergerbasestorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyagglayergerbasestorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyagglayergerbasestorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyagglayergerbasestorageSession struct {
	Contract     *Legacyagglayergerbasestorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LegacyagglayergerbasestorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyagglayergerbasestorageCallerSession struct {
	Contract *LegacyagglayergerbasestorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// LegacyagglayergerbasestorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyagglayergerbasestorageTransactorSession struct {
	Contract     *LegacyagglayergerbasestorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// LegacyagglayergerbasestorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyagglayergerbasestorageRaw struct {
	Contract *Legacyagglayergerbasestorage // Generic contract binding to access the raw methods on
}

// LegacyagglayergerbasestorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyagglayergerbasestorageCallerRaw struct {
	Contract *LegacyagglayergerbasestorageCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyagglayergerbasestorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyagglayergerbasestorageTransactorRaw struct {
	Contract *LegacyagglayergerbasestorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyagglayergerbasestorage creates a new instance of Legacyagglayergerbasestorage, bound to a specific deployed contract.
func NewLegacyagglayergerbasestorage(address common.Address, backend bind.ContractBackend) (*Legacyagglayergerbasestorage, error) {
	contract, err := bindLegacyagglayergerbasestorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Legacyagglayergerbasestorage{LegacyagglayergerbasestorageCaller: LegacyagglayergerbasestorageCaller{contract: contract}, LegacyagglayergerbasestorageTransactor: LegacyagglayergerbasestorageTransactor{contract: contract}, LegacyagglayergerbasestorageFilterer: LegacyagglayergerbasestorageFilterer{contract: contract}}, nil
}

// NewLegacyagglayergerbasestorageCaller creates a new read-only instance of Legacyagglayergerbasestorage, bound to a specific deployed contract.
func NewLegacyagglayergerbasestorageCaller(address common.Address, caller bind.ContractCaller) (*LegacyagglayergerbasestorageCaller, error) {
	contract, err := bindLegacyagglayergerbasestorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyagglayergerbasestorageCaller{contract: contract}, nil
}

// NewLegacyagglayergerbasestorageTransactor creates a new write-only instance of Legacyagglayergerbasestorage, bound to a specific deployed contract.
func NewLegacyagglayergerbasestorageTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyagglayergerbasestorageTransactor, error) {
	contract, err := bindLegacyagglayergerbasestorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyagglayergerbasestorageTransactor{contract: contract}, nil
}

// NewLegacyagglayergerbasestorageFilterer creates a new log filterer instance of Legacyagglayergerbasestorage, bound to a specific deployed contract.
func NewLegacyagglayergerbasestorageFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyagglayergerbasestorageFilterer, error) {
	contract, err := bindLegacyagglayergerbasestorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyagglayergerbasestorageFilterer{contract: contract}, nil
}

// bindLegacyagglayergerbasestorage binds a generic wrapper to an already deployed contract.
func bindLegacyagglayergerbasestorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyagglayergerbasestorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyagglayergerbasestorage.Contract.LegacyagglayergerbasestorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.LegacyagglayergerbasestorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.LegacyagglayergerbasestorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyagglayergerbasestorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.contract.Transact(opts, method, params...)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.GetLastGlobalExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.GetLastGlobalExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) GetRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.GetRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) GetRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.GetRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Legacyagglayergerbasestorage.Contract.GlobalExitRootMap(&_Legacyagglayergerbasestorage.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Legacyagglayergerbasestorage.Contract.GlobalExitRootMap(&_Legacyagglayergerbasestorage.CallOpts, arg0)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) L1InfoRootMap(opts *bind.CallOpts, depositCount uint32) ([32]byte, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "l1InfoRootMap", depositCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) L1InfoRootMap(depositCount uint32) ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.L1InfoRootMap(&_Legacyagglayergerbasestorage.CallOpts, depositCount)
}

// L1InfoRootMap is a free data retrieval call binding the contract method 0xef4eeb35.
//
// Solidity: function l1InfoRootMap(uint32 depositCount) view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) L1InfoRootMap(depositCount uint32) ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.L1InfoRootMap(&_Legacyagglayergerbasestorage.CallOpts, depositCount)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.LastMainnetExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.LastMainnetExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Legacyagglayergerbasestorage.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) LastRollupExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.LastRollupExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Legacyagglayergerbasestorage.Contract.LastRollupExitRoot(&_Legacyagglayergerbasestorage.CallOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.UpdateExitRoot(&_Legacyagglayergerbasestorage.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Legacyagglayergerbasestorage *LegacyagglayergerbasestorageTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Legacyagglayergerbasestorage.Contract.UpdateExitRoot(&_Legacyagglayergerbasestorage.TransactOpts, newRollupExitRoot)
}

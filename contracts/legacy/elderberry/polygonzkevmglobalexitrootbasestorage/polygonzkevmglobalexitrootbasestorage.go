// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootbasestorage

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

// PolygonzkevmglobalexitrootbasestorageMetaData contains all meta data concerning the Polygonzkevmglobalexitrootbasestorage contract.
var PolygonzkevmglobalexitrootbasestorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PolygonzkevmglobalexitrootbasestorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmglobalexitrootbasestorageMetaData.ABI instead.
var PolygonzkevmglobalexitrootbasestorageABI = PolygonzkevmglobalexitrootbasestorageMetaData.ABI

// Polygonzkevmglobalexitrootbasestorage is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootbasestorage struct {
	PolygonzkevmglobalexitrootbasestorageCaller     // Read-only binding to the contract
	PolygonzkevmglobalexitrootbasestorageTransactor // Write-only binding to the contract
	PolygonzkevmglobalexitrootbasestorageFilterer   // Log filterer for contract events
}

// PolygonzkevmglobalexitrootbasestorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootbasestorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootbasestorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootbasestorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootbasestorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmglobalexitrootbasestorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmglobalexitrootbasestorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmglobalexitrootbasestorageSession struct {
	Contract     *Polygonzkevmglobalexitrootbasestorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// PolygonzkevmglobalexitrootbasestorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmglobalexitrootbasestorageCallerSession struct {
	Contract *PolygonzkevmglobalexitrootbasestorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// PolygonzkevmglobalexitrootbasestorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmglobalexitrootbasestorageTransactorSession struct {
	Contract     *PolygonzkevmglobalexitrootbasestorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// PolygonzkevmglobalexitrootbasestorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootbasestorageRaw struct {
	Contract *Polygonzkevmglobalexitrootbasestorage // Generic contract binding to access the raw methods on
}

// PolygonzkevmglobalexitrootbasestorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootbasestorageCallerRaw struct {
	Contract *PolygonzkevmglobalexitrootbasestorageCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmglobalexitrootbasestorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmglobalexitrootbasestorageTransactorRaw struct {
	Contract *PolygonzkevmglobalexitrootbasestorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootbasestorage creates a new instance of Polygonzkevmglobalexitrootbasestorage, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootbasestorage(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootbasestorage, error) {
	contract, err := bindPolygonzkevmglobalexitrootbasestorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootbasestorage{PolygonzkevmglobalexitrootbasestorageCaller: PolygonzkevmglobalexitrootbasestorageCaller{contract: contract}, PolygonzkevmglobalexitrootbasestorageTransactor: PolygonzkevmglobalexitrootbasestorageTransactor{contract: contract}, PolygonzkevmglobalexitrootbasestorageFilterer: PolygonzkevmglobalexitrootbasestorageFilterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootbasestorageCaller creates a new read-only instance of Polygonzkevmglobalexitrootbasestorage, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootbasestorageCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmglobalexitrootbasestorageCaller, error) {
	contract, err := bindPolygonzkevmglobalexitrootbasestorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootbasestorageCaller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootbasestorageTransactor creates a new write-only instance of Polygonzkevmglobalexitrootbasestorage, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootbasestorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmglobalexitrootbasestorageTransactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootbasestorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootbasestorageTransactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootbasestorageFilterer creates a new log filterer instance of Polygonzkevmglobalexitrootbasestorage, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootbasestorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmglobalexitrootbasestorageFilterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootbasestorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmglobalexitrootbasestorageFilterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootbasestorage binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootbasestorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmglobalexitrootbasestorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.PolygonzkevmglobalexitrootbasestorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.PolygonzkevmglobalexitrootbasestorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.PolygonzkevmglobalexitrootbasestorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.contract.Transact(opts, method, params...)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCaller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootbasestorage.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootbasestorage.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GetRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GetRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootbasestorage.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootbasestorage.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootbasestorage.CallOpts, arg0)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCaller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootbasestorage.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.LastMainnetExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootbasestorage.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootbasestorage.CallOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootbasestorage.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Polygonzkevmglobalexitrootbasestorage *PolygonzkevmglobalexitrootbasestorageTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootbasestorage.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootbasestorage.TransactOpts, newRollupExitRoot)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergerl2

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

// Iagglayergerl2MetaData contains all meta data concerning the Iagglayergerl2 contract.
var Iagglayergerl2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGlobalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iagglayergerl2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Iagglayergerl2MetaData.ABI instead.
var Iagglayergerl2ABI = Iagglayergerl2MetaData.ABI

// Iagglayergerl2 is an auto generated Go binding around an Ethereum contract.
type Iagglayergerl2 struct {
	Iagglayergerl2Caller     // Read-only binding to the contract
	Iagglayergerl2Transactor // Write-only binding to the contract
	Iagglayergerl2Filterer   // Log filterer for contract events
}

// Iagglayergerl2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iagglayergerl2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayergerl2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iagglayergerl2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayergerl2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iagglayergerl2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iagglayergerl2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iagglayergerl2Session struct {
	Contract     *Iagglayergerl2   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iagglayergerl2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iagglayergerl2CallerSession struct {
	Contract *Iagglayergerl2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Iagglayergerl2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iagglayergerl2TransactorSession struct {
	Contract     *Iagglayergerl2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Iagglayergerl2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iagglayergerl2Raw struct {
	Contract *Iagglayergerl2 // Generic contract binding to access the raw methods on
}

// Iagglayergerl2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iagglayergerl2CallerRaw struct {
	Contract *Iagglayergerl2Caller // Generic read-only contract binding to access the raw methods on
}

// Iagglayergerl2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iagglayergerl2TransactorRaw struct {
	Contract *Iagglayergerl2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergerl2 creates a new instance of Iagglayergerl2, bound to a specific deployed contract.
func NewIagglayergerl2(address common.Address, backend bind.ContractBackend) (*Iagglayergerl2, error) {
	contract, err := bindIagglayergerl2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergerl2{Iagglayergerl2Caller: Iagglayergerl2Caller{contract: contract}, Iagglayergerl2Transactor: Iagglayergerl2Transactor{contract: contract}, Iagglayergerl2Filterer: Iagglayergerl2Filterer{contract: contract}}, nil
}

// NewIagglayergerl2Caller creates a new read-only instance of Iagglayergerl2, bound to a specific deployed contract.
func NewIagglayergerl2Caller(address common.Address, caller bind.ContractCaller) (*Iagglayergerl2Caller, error) {
	contract, err := bindIagglayergerl2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iagglayergerl2Caller{contract: contract}, nil
}

// NewIagglayergerl2Transactor creates a new write-only instance of Iagglayergerl2, bound to a specific deployed contract.
func NewIagglayergerl2Transactor(address common.Address, transactor bind.ContractTransactor) (*Iagglayergerl2Transactor, error) {
	contract, err := bindIagglayergerl2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iagglayergerl2Transactor{contract: contract}, nil
}

// NewIagglayergerl2Filterer creates a new log filterer instance of Iagglayergerl2, bound to a specific deployed contract.
func NewIagglayergerl2Filterer(address common.Address, filterer bind.ContractFilterer) (*Iagglayergerl2Filterer, error) {
	contract, err := bindIagglayergerl2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iagglayergerl2Filterer{contract: contract}, nil
}

// bindIagglayergerl2 binds a generic wrapper to an already deployed contract.
func bindIagglayergerl2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iagglayergerl2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergerl2 *Iagglayergerl2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergerl2.Contract.Iagglayergerl2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergerl2 *Iagglayergerl2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.Iagglayergerl2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergerl2 *Iagglayergerl2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.Iagglayergerl2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergerl2 *Iagglayergerl2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergerl2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergerl2 *Iagglayergerl2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergerl2 *Iagglayergerl2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.contract.Transact(opts, method, params...)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Caller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iagglayergerl2.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Session) GlobalExitRootRemover() (common.Address, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootRemover(&_Iagglayergerl2.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2CallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootRemover(&_Iagglayergerl2.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Caller) GlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iagglayergerl2.contract.Call(opts, &out, "globalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Session) GlobalExitRootUpdater() (common.Address, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootUpdater(&_Iagglayergerl2.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2CallerSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootUpdater(&_Iagglayergerl2.CallOpts)
}

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Iagglayergerl2 *Iagglayergerl2Caller) InsertedGERHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayergerl2.contract.Call(opts, &out, "insertedGERHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Iagglayergerl2 *Iagglayergerl2Session) InsertedGERHashChain() ([32]byte, error) {
	return _Iagglayergerl2.Contract.InsertedGERHashChain(&_Iagglayergerl2.CallOpts)
}

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Iagglayergerl2 *Iagglayergerl2CallerSession) InsertedGERHashChain() ([32]byte, error) {
	return _Iagglayergerl2.Contract.InsertedGERHashChain(&_Iagglayergerl2.CallOpts)
}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Caller) PendingGlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iagglayergerl2.contract.Call(opts, &out, "pendingGlobalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2Session) PendingGlobalExitRootUpdater() (common.Address, error) {
	return _Iagglayergerl2.Contract.PendingGlobalExitRootUpdater(&_Iagglayergerl2.CallOpts)
}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Iagglayergerl2 *Iagglayergerl2CallerSession) PendingGlobalExitRootUpdater() (common.Address, error) {
	return _Iagglayergerl2.Contract.PendingGlobalExitRootUpdater(&_Iagglayergerl2.CallOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iagglayergerl2 *Iagglayergerl2Transactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergerl2.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iagglayergerl2 *Iagglayergerl2Session) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.AcceptGlobalExitRootUpdater(&_Iagglayergerl2.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iagglayergerl2 *Iagglayergerl2TransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.AcceptGlobalExitRootUpdater(&_Iagglayergerl2.TransactOpts)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayergerl2 *Iagglayergerl2Transactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayergerl2 *Iagglayergerl2Session) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootMap(&_Iagglayergerl2.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iagglayergerl2 *Iagglayergerl2TransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.GlobalExitRootMap(&_Iagglayergerl2.TransactOpts, globalExitRootNum)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2Transactor) InsertGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.contract.Transact(opts, "insertGlobalExitRoot", _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2Session) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.InsertGlobalExitRoot(&_Iagglayergerl2.TransactOpts, _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2TransactorSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.InsertGlobalExitRoot(&_Iagglayergerl2.TransactOpts, _newRoot)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iagglayergerl2 *Iagglayergerl2Transactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iagglayergerl2.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iagglayergerl2 *Iagglayergerl2Session) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.TransferGlobalExitRootUpdater(&_Iagglayergerl2.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iagglayergerl2 *Iagglayergerl2TransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.TransferGlobalExitRootUpdater(&_Iagglayergerl2.TransactOpts, _newGlobalExitRootUpdater)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2Session) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.UpdateExitRoot(&_Iagglayergerl2.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iagglayergerl2 *Iagglayergerl2TransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayergerl2.Contract.UpdateExitRoot(&_Iagglayergerl2.TransactOpts, newRollupExitRoot)
}

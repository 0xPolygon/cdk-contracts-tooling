// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibasepolygonzkevmglobalexitrootpessimistic

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

// IbasepolygonzkevmglobalexitrootpessimisticMetaData contains all meta data concerning the Ibasepolygonzkevmglobalexitrootpessimistic contract.
var IbasepolygonzkevmglobalexitrootpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IbasepolygonzkevmglobalexitrootpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use IbasepolygonzkevmglobalexitrootpessimisticMetaData.ABI instead.
var IbasepolygonzkevmglobalexitrootpessimisticABI = IbasepolygonzkevmglobalexitrootpessimisticMetaData.ABI

// Ibasepolygonzkevmglobalexitrootpessimistic is an auto generated Go binding around an Ethereum contract.
type Ibasepolygonzkevmglobalexitrootpessimistic struct {
	IbasepolygonzkevmglobalexitrootpessimisticCaller     // Read-only binding to the contract
	IbasepolygonzkevmglobalexitrootpessimisticTransactor // Write-only binding to the contract
	IbasepolygonzkevmglobalexitrootpessimisticFilterer   // Log filterer for contract events
}

// IbasepolygonzkevmglobalexitrootpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbasepolygonzkevmglobalexitrootpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbasepolygonzkevmglobalexitrootpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbasepolygonzkevmglobalexitrootpessimisticSession struct {
	Contract     *Ibasepolygonzkevmglobalexitrootpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                               // Call options to use throughout this session
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbasepolygonzkevmglobalexitrootpessimisticCallerSession struct {
	Contract *IbasepolygonzkevmglobalexitrootpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                     // Call options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbasepolygonzkevmglobalexitrootpessimisticTransactorSession struct {
	Contract     *IbasepolygonzkevmglobalexitrootpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                     // Transaction auth options to use throughout this session
}

// IbasepolygonzkevmglobalexitrootpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootpessimisticRaw struct {
	Contract *Ibasepolygonzkevmglobalexitrootpessimistic // Generic contract binding to access the raw methods on
}

// IbasepolygonzkevmglobalexitrootpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootpessimisticCallerRaw struct {
	Contract *IbasepolygonzkevmglobalexitrootpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// IbasepolygonzkevmglobalexitrootpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbasepolygonzkevmglobalexitrootpessimisticTransactorRaw struct {
	Contract *IbasepolygonzkevmglobalexitrootpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbasepolygonzkevmglobalexitrootpessimistic creates a new instance of Ibasepolygonzkevmglobalexitrootpessimistic, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootpessimistic(address common.Address, backend bind.ContractBackend) (*Ibasepolygonzkevmglobalexitrootpessimistic, error) {
	contract, err := bindIbasepolygonzkevmglobalexitrootpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibasepolygonzkevmglobalexitrootpessimistic{IbasepolygonzkevmglobalexitrootpessimisticCaller: IbasepolygonzkevmglobalexitrootpessimisticCaller{contract: contract}, IbasepolygonzkevmglobalexitrootpessimisticTransactor: IbasepolygonzkevmglobalexitrootpessimisticTransactor{contract: contract}, IbasepolygonzkevmglobalexitrootpessimisticFilterer: IbasepolygonzkevmglobalexitrootpessimisticFilterer{contract: contract}}, nil
}

// NewIbasepolygonzkevmglobalexitrootpessimisticCaller creates a new read-only instance of Ibasepolygonzkevmglobalexitrootpessimistic, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootpessimisticCaller(address common.Address, caller bind.ContractCaller) (*IbasepolygonzkevmglobalexitrootpessimisticCaller, error) {
	contract, err := bindIbasepolygonzkevmglobalexitrootpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootpessimisticCaller{contract: contract}, nil
}

// NewIbasepolygonzkevmglobalexitrootpessimisticTransactor creates a new write-only instance of Ibasepolygonzkevmglobalexitrootpessimistic, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*IbasepolygonzkevmglobalexitrootpessimisticTransactor, error) {
	contract, err := bindIbasepolygonzkevmglobalexitrootpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootpessimisticTransactor{contract: contract}, nil
}

// NewIbasepolygonzkevmglobalexitrootpessimisticFilterer creates a new log filterer instance of Ibasepolygonzkevmglobalexitrootpessimistic, bound to a specific deployed contract.
func NewIbasepolygonzkevmglobalexitrootpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*IbasepolygonzkevmglobalexitrootpessimisticFilterer, error) {
	contract, err := bindIbasepolygonzkevmglobalexitrootpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbasepolygonzkevmglobalexitrootpessimisticFilterer{contract: contract}, nil
}

// bindIbasepolygonzkevmglobalexitrootpessimistic binds a generic wrapper to an already deployed contract.
func bindIbasepolygonzkevmglobalexitrootpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbasepolygonzkevmglobalexitrootpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.IbasepolygonzkevmglobalexitrootpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.IbasepolygonzkevmglobalexitrootpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.IbasepolygonzkevmglobalexitrootpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.contract.Transact(opts, method, params...)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.GlobalExitRootMap(&_Ibasepolygonzkevmglobalexitrootpessimistic.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.GlobalExitRootMap(&_Ibasepolygonzkevmglobalexitrootpessimistic.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.UpdateExitRoot(&_Ibasepolygonzkevmglobalexitrootpessimistic.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Ibasepolygonzkevmglobalexitrootpessimistic *IbasepolygonzkevmglobalexitrootpessimisticTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Ibasepolygonzkevmglobalexitrootpessimistic.Contract.UpdateExitRoot(&_Ibasepolygonzkevmglobalexitrootpessimistic.TransactOpts, newRollupExitRoot)
}

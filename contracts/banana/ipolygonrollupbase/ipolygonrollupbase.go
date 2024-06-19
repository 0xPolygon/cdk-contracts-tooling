// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonrollupbase

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

// IpolygonrollupbaseMetaData contains all meta data concerning the Ipolygonrollupbase contract.
var IpolygonrollupbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonrollupbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonrollupbaseMetaData.ABI instead.
var IpolygonrollupbaseABI = IpolygonrollupbaseMetaData.ABI

// Ipolygonrollupbase is an auto generated Go binding around an Ethereum contract.
type Ipolygonrollupbase struct {
	IpolygonrollupbaseCaller     // Read-only binding to the contract
	IpolygonrollupbaseTransactor // Write-only binding to the contract
	IpolygonrollupbaseFilterer   // Log filterer for contract events
}

// IpolygonrollupbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonrollupbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonrollupbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonrollupbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonrollupbaseSession struct {
	Contract     *Ipolygonrollupbase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IpolygonrollupbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonrollupbaseCallerSession struct {
	Contract *IpolygonrollupbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IpolygonrollupbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonrollupbaseTransactorSession struct {
	Contract     *IpolygonrollupbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IpolygonrollupbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonrollupbaseRaw struct {
	Contract *Ipolygonrollupbase // Generic contract binding to access the raw methods on
}

// IpolygonrollupbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonrollupbaseCallerRaw struct {
	Contract *IpolygonrollupbaseCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonrollupbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonrollupbaseTransactorRaw struct {
	Contract *IpolygonrollupbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonrollupbase creates a new instance of Ipolygonrollupbase, bound to a specific deployed contract.
func NewIpolygonrollupbase(address common.Address, backend bind.ContractBackend) (*Ipolygonrollupbase, error) {
	contract, err := bindIpolygonrollupbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonrollupbase{IpolygonrollupbaseCaller: IpolygonrollupbaseCaller{contract: contract}, IpolygonrollupbaseTransactor: IpolygonrollupbaseTransactor{contract: contract}, IpolygonrollupbaseFilterer: IpolygonrollupbaseFilterer{contract: contract}}, nil
}

// NewIpolygonrollupbaseCaller creates a new read-only instance of Ipolygonrollupbase, bound to a specific deployed contract.
func NewIpolygonrollupbaseCaller(address common.Address, caller bind.ContractCaller) (*IpolygonrollupbaseCaller, error) {
	contract, err := bindIpolygonrollupbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbaseCaller{contract: contract}, nil
}

// NewIpolygonrollupbaseTransactor creates a new write-only instance of Ipolygonrollupbase, bound to a specific deployed contract.
func NewIpolygonrollupbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonrollupbaseTransactor, error) {
	contract, err := bindIpolygonrollupbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbaseTransactor{contract: contract}, nil
}

// NewIpolygonrollupbaseFilterer creates a new log filterer instance of Ipolygonrollupbase, bound to a specific deployed contract.
func NewIpolygonrollupbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonrollupbaseFilterer, error) {
	contract, err := bindIpolygonrollupbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbaseFilterer{contract: contract}, nil
}

// bindIpolygonrollupbase binds a generic wrapper to an already deployed contract.
func bindIpolygonrollupbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonrollupbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupbase *IpolygonrollupbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupbase.Contract.IpolygonrollupbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupbase *IpolygonrollupbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.IpolygonrollupbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupbase *IpolygonrollupbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.IpolygonrollupbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupbase *IpolygonrollupbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbase.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbase *IpolygonrollupbaseSession) Admin() (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.Admin(&_Ipolygonrollupbase.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorSession) Admin() (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.Admin(&_Ipolygonrollupbase.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbase.contract.Transact(opts, "initialize", _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.Initialize(&_Ipolygonrollupbase.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.Initialize(&_Ipolygonrollupbase.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbase.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.OnVerifyBatches(&_Ipolygonrollupbase.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.OnVerifyBatches(&_Ipolygonrollupbase.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactor) RollbackBatches(opts *bind.TransactOpts, targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupbase.contract.Transact(opts, "rollbackBatches", targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.RollbackBatches(&_Ipolygonrollupbase.TransactOpts, targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_Ipolygonrollupbase *IpolygonrollupbaseTransactorSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupbase.Contract.RollbackBatches(&_Ipolygonrollupbase.TransactOpts, targetBatch, accInputHashToRollback)
}

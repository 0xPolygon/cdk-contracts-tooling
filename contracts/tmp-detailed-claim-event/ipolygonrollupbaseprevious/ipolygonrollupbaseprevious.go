// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonrollupbaseprevious

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

// IpolygonrollupbasepreviousMetaData contains all meta data concerning the Ipolygonrollupbaseprevious contract.
var IpolygonrollupbasepreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonrollupbasepreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonrollupbasepreviousMetaData.ABI instead.
var IpolygonrollupbasepreviousABI = IpolygonrollupbasepreviousMetaData.ABI

// Ipolygonrollupbaseprevious is an auto generated Go binding around an Ethereum contract.
type Ipolygonrollupbaseprevious struct {
	IpolygonrollupbasepreviousCaller     // Read-only binding to the contract
	IpolygonrollupbasepreviousTransactor // Write-only binding to the contract
	IpolygonrollupbasepreviousFilterer   // Log filterer for contract events
}

// IpolygonrollupbasepreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonrollupbasepreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbasepreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonrollupbasepreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbasepreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonrollupbasepreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupbasepreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonrollupbasepreviousSession struct {
	Contract     *Ipolygonrollupbaseprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IpolygonrollupbasepreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonrollupbasepreviousCallerSession struct {
	Contract *IpolygonrollupbasepreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IpolygonrollupbasepreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonrollupbasepreviousTransactorSession struct {
	Contract     *IpolygonrollupbasepreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IpolygonrollupbasepreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonrollupbasepreviousRaw struct {
	Contract *Ipolygonrollupbaseprevious // Generic contract binding to access the raw methods on
}

// IpolygonrollupbasepreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonrollupbasepreviousCallerRaw struct {
	Contract *IpolygonrollupbasepreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonrollupbasepreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonrollupbasepreviousTransactorRaw struct {
	Contract *IpolygonrollupbasepreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonrollupbaseprevious creates a new instance of Ipolygonrollupbaseprevious, bound to a specific deployed contract.
func NewIpolygonrollupbaseprevious(address common.Address, backend bind.ContractBackend) (*Ipolygonrollupbaseprevious, error) {
	contract, err := bindIpolygonrollupbaseprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonrollupbaseprevious{IpolygonrollupbasepreviousCaller: IpolygonrollupbasepreviousCaller{contract: contract}, IpolygonrollupbasepreviousTransactor: IpolygonrollupbasepreviousTransactor{contract: contract}, IpolygonrollupbasepreviousFilterer: IpolygonrollupbasepreviousFilterer{contract: contract}}, nil
}

// NewIpolygonrollupbasepreviousCaller creates a new read-only instance of Ipolygonrollupbaseprevious, bound to a specific deployed contract.
func NewIpolygonrollupbasepreviousCaller(address common.Address, caller bind.ContractCaller) (*IpolygonrollupbasepreviousCaller, error) {
	contract, err := bindIpolygonrollupbaseprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbasepreviousCaller{contract: contract}, nil
}

// NewIpolygonrollupbasepreviousTransactor creates a new write-only instance of Ipolygonrollupbaseprevious, bound to a specific deployed contract.
func NewIpolygonrollupbasepreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonrollupbasepreviousTransactor, error) {
	contract, err := bindIpolygonrollupbaseprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbasepreviousTransactor{contract: contract}, nil
}

// NewIpolygonrollupbasepreviousFilterer creates a new log filterer instance of Ipolygonrollupbaseprevious, bound to a specific deployed contract.
func NewIpolygonrollupbasepreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonrollupbasepreviousFilterer, error) {
	contract, err := bindIpolygonrollupbaseprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupbasepreviousFilterer{contract: contract}, nil
}

// bindIpolygonrollupbaseprevious binds a generic wrapper to an already deployed contract.
func bindIpolygonrollupbaseprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonrollupbasepreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupbaseprevious.Contract.IpolygonrollupbasepreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.IpolygonrollupbasepreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.IpolygonrollupbasepreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupbaseprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousSession) Admin() (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.Admin(&_Ipolygonrollupbaseprevious.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactorSession) Admin() (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.Admin(&_Ipolygonrollupbaseprevious.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.contract.Transact(opts, "initialize", _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.Initialize(&_Ipolygonrollupbaseprevious.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.Initialize(&_Ipolygonrollupbaseprevious.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.OnVerifyBatches(&_Ipolygonrollupbaseprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Ipolygonrollupbaseprevious *IpolygonrollupbasepreviousTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Ipolygonrollupbaseprevious.Contract.OnVerifyBatches(&_Ipolygonrollupbaseprevious.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

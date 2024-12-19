// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonconsensusbase

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

// IpolygonconsensusbaseMetaData contains all meta data concerning the Ipolygonconsensusbase contract.
var IpolygonconsensusbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonconsensusbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonconsensusbaseMetaData.ABI instead.
var IpolygonconsensusbaseABI = IpolygonconsensusbaseMetaData.ABI

// Ipolygonconsensusbase is an auto generated Go binding around an Ethereum contract.
type Ipolygonconsensusbase struct {
	IpolygonconsensusbaseCaller     // Read-only binding to the contract
	IpolygonconsensusbaseTransactor // Write-only binding to the contract
	IpolygonconsensusbaseFilterer   // Log filterer for contract events
}

// IpolygonconsensusbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonconsensusbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonconsensusbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonconsensusbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonconsensusbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonconsensusbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonconsensusbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonconsensusbaseSession struct {
	Contract     *Ipolygonconsensusbase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IpolygonconsensusbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonconsensusbaseCallerSession struct {
	Contract *IpolygonconsensusbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IpolygonconsensusbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonconsensusbaseTransactorSession struct {
	Contract     *IpolygonconsensusbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IpolygonconsensusbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonconsensusbaseRaw struct {
	Contract *Ipolygonconsensusbase // Generic contract binding to access the raw methods on
}

// IpolygonconsensusbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonconsensusbaseCallerRaw struct {
	Contract *IpolygonconsensusbaseCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonconsensusbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonconsensusbaseTransactorRaw struct {
	Contract *IpolygonconsensusbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonconsensusbase creates a new instance of Ipolygonconsensusbase, bound to a specific deployed contract.
func NewIpolygonconsensusbase(address common.Address, backend bind.ContractBackend) (*Ipolygonconsensusbase, error) {
	contract, err := bindIpolygonconsensusbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonconsensusbase{IpolygonconsensusbaseCaller: IpolygonconsensusbaseCaller{contract: contract}, IpolygonconsensusbaseTransactor: IpolygonconsensusbaseTransactor{contract: contract}, IpolygonconsensusbaseFilterer: IpolygonconsensusbaseFilterer{contract: contract}}, nil
}

// NewIpolygonconsensusbaseCaller creates a new read-only instance of Ipolygonconsensusbase, bound to a specific deployed contract.
func NewIpolygonconsensusbaseCaller(address common.Address, caller bind.ContractCaller) (*IpolygonconsensusbaseCaller, error) {
	contract, err := bindIpolygonconsensusbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonconsensusbaseCaller{contract: contract}, nil
}

// NewIpolygonconsensusbaseTransactor creates a new write-only instance of Ipolygonconsensusbase, bound to a specific deployed contract.
func NewIpolygonconsensusbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonconsensusbaseTransactor, error) {
	contract, err := bindIpolygonconsensusbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonconsensusbaseTransactor{contract: contract}, nil
}

// NewIpolygonconsensusbaseFilterer creates a new log filterer instance of Ipolygonconsensusbase, bound to a specific deployed contract.
func NewIpolygonconsensusbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonconsensusbaseFilterer, error) {
	contract, err := bindIpolygonconsensusbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonconsensusbaseFilterer{contract: contract}, nil
}

// bindIpolygonconsensusbase binds a generic wrapper to an already deployed contract.
func bindIpolygonconsensusbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonconsensusbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonconsensusbase.Contract.IpolygonconsensusbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.IpolygonconsensusbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.IpolygonconsensusbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonconsensusbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonconsensusbase *IpolygonconsensusbaseSession) Admin() (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.Admin(&_Ipolygonconsensusbase.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactorSession) Admin() (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.Admin(&_Ipolygonconsensusbase.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.contract.Transact(opts, "initialize", _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonconsensusbase *IpolygonconsensusbaseSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.Initialize(&_Ipolygonconsensusbase.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Ipolygonconsensusbase *IpolygonconsensusbaseTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Ipolygonconsensusbase.Contract.Initialize(&_Ipolygonconsensusbase.TransactOpts, _admin, sequencer, networkID, gasTokenAddress, sequencerURL, _networkName)
}

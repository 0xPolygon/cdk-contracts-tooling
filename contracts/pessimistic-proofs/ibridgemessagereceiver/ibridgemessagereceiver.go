// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibridgemessagereceiver

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

// IbridgemessagereceiverMetaData contains all meta data concerning the Ibridgemessagereceiver contract.
var IbridgemessagereceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onMessageReceived\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IbridgemessagereceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbridgemessagereceiverMetaData.ABI instead.
var IbridgemessagereceiverABI = IbridgemessagereceiverMetaData.ABI

// Ibridgemessagereceiver is an auto generated Go binding around an Ethereum contract.
type Ibridgemessagereceiver struct {
	IbridgemessagereceiverCaller     // Read-only binding to the contract
	IbridgemessagereceiverTransactor // Write-only binding to the contract
	IbridgemessagereceiverFilterer   // Log filterer for contract events
}

// IbridgemessagereceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbridgemessagereceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbridgemessagereceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbridgemessagereceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbridgemessagereceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbridgemessagereceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbridgemessagereceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbridgemessagereceiverSession struct {
	Contract     *Ibridgemessagereceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IbridgemessagereceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbridgemessagereceiverCallerSession struct {
	Contract *IbridgemessagereceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IbridgemessagereceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbridgemessagereceiverTransactorSession struct {
	Contract     *IbridgemessagereceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IbridgemessagereceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbridgemessagereceiverRaw struct {
	Contract *Ibridgemessagereceiver // Generic contract binding to access the raw methods on
}

// IbridgemessagereceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbridgemessagereceiverCallerRaw struct {
	Contract *IbridgemessagereceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbridgemessagereceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbridgemessagereceiverTransactorRaw struct {
	Contract *IbridgemessagereceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbridgemessagereceiver creates a new instance of Ibridgemessagereceiver, bound to a specific deployed contract.
func NewIbridgemessagereceiver(address common.Address, backend bind.ContractBackend) (*Ibridgemessagereceiver, error) {
	contract, err := bindIbridgemessagereceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibridgemessagereceiver{IbridgemessagereceiverCaller: IbridgemessagereceiverCaller{contract: contract}, IbridgemessagereceiverTransactor: IbridgemessagereceiverTransactor{contract: contract}, IbridgemessagereceiverFilterer: IbridgemessagereceiverFilterer{contract: contract}}, nil
}

// NewIbridgemessagereceiverCaller creates a new read-only instance of Ibridgemessagereceiver, bound to a specific deployed contract.
func NewIbridgemessagereceiverCaller(address common.Address, caller bind.ContractCaller) (*IbridgemessagereceiverCaller, error) {
	contract, err := bindIbridgemessagereceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbridgemessagereceiverCaller{contract: contract}, nil
}

// NewIbridgemessagereceiverTransactor creates a new write-only instance of Ibridgemessagereceiver, bound to a specific deployed contract.
func NewIbridgemessagereceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbridgemessagereceiverTransactor, error) {
	contract, err := bindIbridgemessagereceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbridgemessagereceiverTransactor{contract: contract}, nil
}

// NewIbridgemessagereceiverFilterer creates a new log filterer instance of Ibridgemessagereceiver, bound to a specific deployed contract.
func NewIbridgemessagereceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbridgemessagereceiverFilterer, error) {
	contract, err := bindIbridgemessagereceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbridgemessagereceiverFilterer{contract: contract}, nil
}

// bindIbridgemessagereceiver binds a generic wrapper to an already deployed contract.
func bindIbridgemessagereceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbridgemessagereceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibridgemessagereceiver *IbridgemessagereceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibridgemessagereceiver.Contract.IbridgemessagereceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibridgemessagereceiver *IbridgemessagereceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.IbridgemessagereceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibridgemessagereceiver *IbridgemessagereceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.IbridgemessagereceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibridgemessagereceiver *IbridgemessagereceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibridgemessagereceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibridgemessagereceiver *IbridgemessagereceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibridgemessagereceiver *IbridgemessagereceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.contract.Transact(opts, method, params...)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Ibridgemessagereceiver *IbridgemessagereceiverTransactor) OnMessageReceived(opts *bind.TransactOpts, originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.contract.Transact(opts, "onMessageReceived", originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Ibridgemessagereceiver *IbridgemessagereceiverSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.OnMessageReceived(&_Ibridgemessagereceiver.TransactOpts, originAddress, originNetwork, data)
}

// OnMessageReceived is a paid mutator transaction binding the contract method 0x1806b5f2.
//
// Solidity: function onMessageReceived(address originAddress, uint32 originNetwork, bytes data) payable returns()
func (_Ibridgemessagereceiver *IbridgemessagereceiverTransactorSession) OnMessageReceived(originAddress common.Address, originNetwork uint32, data []byte) (*types.Transaction, error) {
	return _Ibridgemessagereceiver.Contract.OnMessageReceived(&_Ibridgemessagereceiver.TransactOpts, originAddress, originNetwork, data)
}

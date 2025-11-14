// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iversion

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

// IversionMetaData contains all meta data concerning the Iversion contract.
var IversionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// IversionABI is the input ABI used to generate the binding from.
// Deprecated: Use IversionMetaData.ABI instead.
var IversionABI = IversionMetaData.ABI

// Iversion is an auto generated Go binding around an Ethereum contract.
type Iversion struct {
	IversionCaller     // Read-only binding to the contract
	IversionTransactor // Write-only binding to the contract
	IversionFilterer   // Log filterer for contract events
}

// IversionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IversionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IversionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IversionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IversionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IversionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IversionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IversionSession struct {
	Contract     *Iversion         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IversionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IversionCallerSession struct {
	Contract *IversionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IversionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IversionTransactorSession struct {
	Contract     *IversionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IversionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IversionRaw struct {
	Contract *Iversion // Generic contract binding to access the raw methods on
}

// IversionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IversionCallerRaw struct {
	Contract *IversionCaller // Generic read-only contract binding to access the raw methods on
}

// IversionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IversionTransactorRaw struct {
	Contract *IversionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIversion creates a new instance of Iversion, bound to a specific deployed contract.
func NewIversion(address common.Address, backend bind.ContractBackend) (*Iversion, error) {
	contract, err := bindIversion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iversion{IversionCaller: IversionCaller{contract: contract}, IversionTransactor: IversionTransactor{contract: contract}, IversionFilterer: IversionFilterer{contract: contract}}, nil
}

// NewIversionCaller creates a new read-only instance of Iversion, bound to a specific deployed contract.
func NewIversionCaller(address common.Address, caller bind.ContractCaller) (*IversionCaller, error) {
	contract, err := bindIversion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IversionCaller{contract: contract}, nil
}

// NewIversionTransactor creates a new write-only instance of Iversion, bound to a specific deployed contract.
func NewIversionTransactor(address common.Address, transactor bind.ContractTransactor) (*IversionTransactor, error) {
	contract, err := bindIversion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IversionTransactor{contract: contract}, nil
}

// NewIversionFilterer creates a new log filterer instance of Iversion, bound to a specific deployed contract.
func NewIversionFilterer(address common.Address, filterer bind.ContractFilterer) (*IversionFilterer, error) {
	contract, err := bindIversion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IversionFilterer{contract: contract}, nil
}

// bindIversion binds a generic wrapper to an already deployed contract.
func bindIversion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IversionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iversion *IversionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iversion.Contract.IversionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iversion *IversionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iversion.Contract.IversionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iversion *IversionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iversion.Contract.IversionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iversion *IversionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iversion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iversion *IversionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iversion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iversion *IversionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iversion.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iversion *IversionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Iversion.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iversion *IversionSession) Version() (string, error) {
	return _Iversion.Contract.Version(&_Iversion.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iversion *IversionCallerSession) Version() (string, error) {
	return _Iversion.Contract.Version(&_Iversion.CallOpts)
}

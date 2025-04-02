// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonvalidium

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

// IpolygonvalidiumMetaData contains all meta data concerning the Ipolygonvalidium contract.
var IpolygonvalidiumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"}]",
}

// IpolygonvalidiumABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonvalidiumMetaData.ABI instead.
var IpolygonvalidiumABI = IpolygonvalidiumMetaData.ABI

// Ipolygonvalidium is an auto generated Go binding around an Ethereum contract.
type Ipolygonvalidium struct {
	IpolygonvalidiumCaller     // Read-only binding to the contract
	IpolygonvalidiumTransactor // Write-only binding to the contract
	IpolygonvalidiumFilterer   // Log filterer for contract events
}

// IpolygonvalidiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonvalidiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonvalidiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonvalidiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonvalidiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonvalidiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonvalidiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonvalidiumSession struct {
	Contract     *Ipolygonvalidium // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IpolygonvalidiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonvalidiumCallerSession struct {
	Contract *IpolygonvalidiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IpolygonvalidiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonvalidiumTransactorSession struct {
	Contract     *IpolygonvalidiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IpolygonvalidiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonvalidiumRaw struct {
	Contract *Ipolygonvalidium // Generic contract binding to access the raw methods on
}

// IpolygonvalidiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonvalidiumCallerRaw struct {
	Contract *IpolygonvalidiumCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonvalidiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonvalidiumTransactorRaw struct {
	Contract *IpolygonvalidiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonvalidium creates a new instance of Ipolygonvalidium, bound to a specific deployed contract.
func NewIpolygonvalidium(address common.Address, backend bind.ContractBackend) (*Ipolygonvalidium, error) {
	contract, err := bindIpolygonvalidium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonvalidium{IpolygonvalidiumCaller: IpolygonvalidiumCaller{contract: contract}, IpolygonvalidiumTransactor: IpolygonvalidiumTransactor{contract: contract}, IpolygonvalidiumFilterer: IpolygonvalidiumFilterer{contract: contract}}, nil
}

// NewIpolygonvalidiumCaller creates a new read-only instance of Ipolygonvalidium, bound to a specific deployed contract.
func NewIpolygonvalidiumCaller(address common.Address, caller bind.ContractCaller) (*IpolygonvalidiumCaller, error) {
	contract, err := bindIpolygonvalidium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonvalidiumCaller{contract: contract}, nil
}

// NewIpolygonvalidiumTransactor creates a new write-only instance of Ipolygonvalidium, bound to a specific deployed contract.
func NewIpolygonvalidiumTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonvalidiumTransactor, error) {
	contract, err := bindIpolygonvalidium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonvalidiumTransactor{contract: contract}, nil
}

// NewIpolygonvalidiumFilterer creates a new log filterer instance of Ipolygonvalidium, bound to a specific deployed contract.
func NewIpolygonvalidiumFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonvalidiumFilterer, error) {
	contract, err := bindIpolygonvalidium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonvalidiumFilterer{contract: contract}, nil
}

// bindIpolygonvalidium binds a generic wrapper to an already deployed contract.
func bindIpolygonvalidium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonvalidiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonvalidium *IpolygonvalidiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonvalidium.Contract.IpolygonvalidiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonvalidium *IpolygonvalidiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonvalidium.Contract.IpolygonvalidiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonvalidium *IpolygonvalidiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonvalidium.Contract.IpolygonvalidiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonvalidium *IpolygonvalidiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonvalidium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonvalidium *IpolygonvalidiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonvalidium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonvalidium *IpolygonvalidiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonvalidium.Contract.contract.Transact(opts, method, params...)
}

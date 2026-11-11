// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iproxyadmin

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

// IproxyadminMetaData contains all meta data concerning the Iproxyadmin contract.
var IproxyadminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IproxyadminABI is the input ABI used to generate the binding from.
// Deprecated: Use IproxyadminMetaData.ABI instead.
var IproxyadminABI = IproxyadminMetaData.ABI

// Iproxyadmin is an auto generated Go binding around an Ethereum contract.
type Iproxyadmin struct {
	IproxyadminCaller     // Read-only binding to the contract
	IproxyadminTransactor // Write-only binding to the contract
	IproxyadminFilterer   // Log filterer for contract events
}

// IproxyadminCaller is an auto generated read-only Go binding around an Ethereum contract.
type IproxyadminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IproxyadminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IproxyadminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IproxyadminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IproxyadminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IproxyadminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IproxyadminSession struct {
	Contract     *Iproxyadmin      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IproxyadminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IproxyadminCallerSession struct {
	Contract *IproxyadminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IproxyadminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IproxyadminTransactorSession struct {
	Contract     *IproxyadminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IproxyadminRaw is an auto generated low-level Go binding around an Ethereum contract.
type IproxyadminRaw struct {
	Contract *Iproxyadmin // Generic contract binding to access the raw methods on
}

// IproxyadminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IproxyadminCallerRaw struct {
	Contract *IproxyadminCaller // Generic read-only contract binding to access the raw methods on
}

// IproxyadminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IproxyadminTransactorRaw struct {
	Contract *IproxyadminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIproxyadmin creates a new instance of Iproxyadmin, bound to a specific deployed contract.
func NewIproxyadmin(address common.Address, backend bind.ContractBackend) (*Iproxyadmin, error) {
	contract, err := bindIproxyadmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iproxyadmin{IproxyadminCaller: IproxyadminCaller{contract: contract}, IproxyadminTransactor: IproxyadminTransactor{contract: contract}, IproxyadminFilterer: IproxyadminFilterer{contract: contract}}, nil
}

// NewIproxyadminCaller creates a new read-only instance of Iproxyadmin, bound to a specific deployed contract.
func NewIproxyadminCaller(address common.Address, caller bind.ContractCaller) (*IproxyadminCaller, error) {
	contract, err := bindIproxyadmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IproxyadminCaller{contract: contract}, nil
}

// NewIproxyadminTransactor creates a new write-only instance of Iproxyadmin, bound to a specific deployed contract.
func NewIproxyadminTransactor(address common.Address, transactor bind.ContractTransactor) (*IproxyadminTransactor, error) {
	contract, err := bindIproxyadmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IproxyadminTransactor{contract: contract}, nil
}

// NewIproxyadminFilterer creates a new log filterer instance of Iproxyadmin, bound to a specific deployed contract.
func NewIproxyadminFilterer(address common.Address, filterer bind.ContractFilterer) (*IproxyadminFilterer, error) {
	contract, err := bindIproxyadmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IproxyadminFilterer{contract: contract}, nil
}

// bindIproxyadmin binds a generic wrapper to an already deployed contract.
func bindIproxyadmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IproxyadminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iproxyadmin *IproxyadminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iproxyadmin.Contract.IproxyadminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iproxyadmin *IproxyadminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iproxyadmin.Contract.IproxyadminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iproxyadmin *IproxyadminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iproxyadmin.Contract.IproxyadminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iproxyadmin *IproxyadminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iproxyadmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iproxyadmin *IproxyadminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iproxyadmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iproxyadmin *IproxyadminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iproxyadmin.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Iproxyadmin *IproxyadminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iproxyadmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Iproxyadmin *IproxyadminSession) Owner() (common.Address, error) {
	return _Iproxyadmin.Contract.Owner(&_Iproxyadmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Iproxyadmin *IproxyadminCallerSession) Owner() (common.Address, error) {
	return _Iproxyadmin.Contract.Owner(&_Iproxyadmin.CallOpts)
}

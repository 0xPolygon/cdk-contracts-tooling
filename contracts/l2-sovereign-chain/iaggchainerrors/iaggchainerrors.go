// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainerrors

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

// IaggchainerrorsMetaData contains all meta data concerning the Iaggchainerrors contract.
var IaggchainerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"}]",
}

// IaggchainerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainerrorsMetaData.ABI instead.
var IaggchainerrorsABI = IaggchainerrorsMetaData.ABI

// Iaggchainerrors is an auto generated Go binding around an Ethereum contract.
type Iaggchainerrors struct {
	IaggchainerrorsCaller     // Read-only binding to the contract
	IaggchainerrorsTransactor // Write-only binding to the contract
	IaggchainerrorsFilterer   // Log filterer for contract events
}

// IaggchainerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainerrorsSession struct {
	Contract     *Iaggchainerrors  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IaggchainerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainerrorsCallerSession struct {
	Contract *IaggchainerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IaggchainerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainerrorsTransactorSession struct {
	Contract     *IaggchainerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IaggchainerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainerrorsRaw struct {
	Contract *Iaggchainerrors // Generic contract binding to access the raw methods on
}

// IaggchainerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainerrorsCallerRaw struct {
	Contract *IaggchainerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainerrorsTransactorRaw struct {
	Contract *IaggchainerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainerrors creates a new instance of Iaggchainerrors, bound to a specific deployed contract.
func NewIaggchainerrors(address common.Address, backend bind.ContractBackend) (*Iaggchainerrors, error) {
	contract, err := bindIaggchainerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainerrors{IaggchainerrorsCaller: IaggchainerrorsCaller{contract: contract}, IaggchainerrorsTransactor: IaggchainerrorsTransactor{contract: contract}, IaggchainerrorsFilterer: IaggchainerrorsFilterer{contract: contract}}, nil
}

// NewIaggchainerrorsCaller creates a new read-only instance of Iaggchainerrors, bound to a specific deployed contract.
func NewIaggchainerrorsCaller(address common.Address, caller bind.ContractCaller) (*IaggchainerrorsCaller, error) {
	contract, err := bindIaggchainerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainerrorsCaller{contract: contract}, nil
}

// NewIaggchainerrorsTransactor creates a new write-only instance of Iaggchainerrors, bound to a specific deployed contract.
func NewIaggchainerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainerrorsTransactor, error) {
	contract, err := bindIaggchainerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainerrorsTransactor{contract: contract}, nil
}

// NewIaggchainerrorsFilterer creates a new log filterer instance of Iaggchainerrors, bound to a specific deployed contract.
func NewIaggchainerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainerrorsFilterer, error) {
	contract, err := bindIaggchainerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainerrorsFilterer{contract: contract}, nil
}

// bindIaggchainerrors binds a generic wrapper to an already deployed contract.
func bindIaggchainerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainerrors *IaggchainerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainerrors.Contract.IaggchainerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainerrors *IaggchainerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainerrors.Contract.IaggchainerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainerrors *IaggchainerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainerrors.Contract.IaggchainerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainerrors *IaggchainerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainerrors *IaggchainerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainerrors *IaggchainerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainerrors.Contract.contract.Transact(opts, method, params...)
}

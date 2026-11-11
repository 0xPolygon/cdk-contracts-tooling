// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggoraclecommitteeerrors

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

// IaggoraclecommitteeerrorsMetaData contains all meta data concerning the Iaggoraclecommitteeerrors contract.
var IaggoraclecommitteeerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AgglayerGERL2CannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERAlreadyInjected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignaturesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggOracleProposerCanCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RenounceOwnershipDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeGreaterThanValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeThisContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorsNotOrdered\",\"type\":\"error\"}]",
}

// IaggoraclecommitteeerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggoraclecommitteeerrorsMetaData.ABI instead.
var IaggoraclecommitteeerrorsABI = IaggoraclecommitteeerrorsMetaData.ABI

// Iaggoraclecommitteeerrors is an auto generated Go binding around an Ethereum contract.
type Iaggoraclecommitteeerrors struct {
	IaggoraclecommitteeerrorsCaller     // Read-only binding to the contract
	IaggoraclecommitteeerrorsTransactor // Write-only binding to the contract
	IaggoraclecommitteeerrorsFilterer   // Log filterer for contract events
}

// IaggoraclecommitteeerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggoraclecommitteeerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggoraclecommitteeerrorsSession struct {
	Contract     *Iaggoraclecommitteeerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggoraclecommitteeerrorsCallerSession struct {
	Contract *IaggoraclecommitteeerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IaggoraclecommitteeerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggoraclecommitteeerrorsTransactorSession struct {
	Contract     *IaggoraclecommitteeerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggoraclecommitteeerrorsRaw struct {
	Contract *Iaggoraclecommitteeerrors // Generic contract binding to access the raw methods on
}

// IaggoraclecommitteeerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeerrorsCallerRaw struct {
	Contract *IaggoraclecommitteeerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IaggoraclecommitteeerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeerrorsTransactorRaw struct {
	Contract *IaggoraclecommitteeerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggoraclecommitteeerrors creates a new instance of Iaggoraclecommitteeerrors, bound to a specific deployed contract.
func NewIaggoraclecommitteeerrors(address common.Address, backend bind.ContractBackend) (*Iaggoraclecommitteeerrors, error) {
	contract, err := bindIaggoraclecommitteeerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggoraclecommitteeerrors{IaggoraclecommitteeerrorsCaller: IaggoraclecommitteeerrorsCaller{contract: contract}, IaggoraclecommitteeerrorsTransactor: IaggoraclecommitteeerrorsTransactor{contract: contract}, IaggoraclecommitteeerrorsFilterer: IaggoraclecommitteeerrorsFilterer{contract: contract}}, nil
}

// NewIaggoraclecommitteeerrorsCaller creates a new read-only instance of Iaggoraclecommitteeerrors, bound to a specific deployed contract.
func NewIaggoraclecommitteeerrorsCaller(address common.Address, caller bind.ContractCaller) (*IaggoraclecommitteeerrorsCaller, error) {
	contract, err := bindIaggoraclecommitteeerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeerrorsCaller{contract: contract}, nil
}

// NewIaggoraclecommitteeerrorsTransactor creates a new write-only instance of Iaggoraclecommitteeerrors, bound to a specific deployed contract.
func NewIaggoraclecommitteeerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggoraclecommitteeerrorsTransactor, error) {
	contract, err := bindIaggoraclecommitteeerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeerrorsTransactor{contract: contract}, nil
}

// NewIaggoraclecommitteeerrorsFilterer creates a new log filterer instance of Iaggoraclecommitteeerrors, bound to a specific deployed contract.
func NewIaggoraclecommitteeerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggoraclecommitteeerrorsFilterer, error) {
	contract, err := bindIaggoraclecommitteeerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeerrorsFilterer{contract: contract}, nil
}

// bindIaggoraclecommitteeerrors binds a generic wrapper to an already deployed contract.
func bindIaggoraclecommitteeerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggoraclecommitteeerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommitteeerrors.Contract.IaggoraclecommitteeerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommitteeerrors.Contract.IaggoraclecommitteeerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommitteeerrors.Contract.IaggoraclecommitteeerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommitteeerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommitteeerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommitteeerrors *IaggoraclecommitteeerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommitteeerrors.Contract.contract.Transact(opts, method, params...)
}

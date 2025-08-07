// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygondatacommitteeerrors

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

// IpolygondatacommitteeerrorsMetaData contains all meta data concerning the Ipolygondatacommitteeerrors contract.
var IpolygondatacommitteeerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CommitteeAddressDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyURLNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyRequiredSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsAndSignaturesSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedAddrsBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedCommitteeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAddrOrder\",\"type\":\"error\"}]",
}

// IpolygondatacommitteeerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygondatacommitteeerrorsMetaData.ABI instead.
var IpolygondatacommitteeerrorsABI = IpolygondatacommitteeerrorsMetaData.ABI

// Ipolygondatacommitteeerrors is an auto generated Go binding around an Ethereum contract.
type Ipolygondatacommitteeerrors struct {
	IpolygondatacommitteeerrorsCaller     // Read-only binding to the contract
	IpolygondatacommitteeerrorsTransactor // Write-only binding to the contract
	IpolygondatacommitteeerrorsFilterer   // Log filterer for contract events
}

// IpolygondatacommitteeerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygondatacommitteeerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygondatacommitteeerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygondatacommitteeerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygondatacommitteeerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygondatacommitteeerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygondatacommitteeerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygondatacommitteeerrorsSession struct {
	Contract     *Ipolygondatacommitteeerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IpolygondatacommitteeerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygondatacommitteeerrorsCallerSession struct {
	Contract *IpolygondatacommitteeerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IpolygondatacommitteeerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygondatacommitteeerrorsTransactorSession struct {
	Contract     *IpolygondatacommitteeerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IpolygondatacommitteeerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygondatacommitteeerrorsRaw struct {
	Contract *Ipolygondatacommitteeerrors // Generic contract binding to access the raw methods on
}

// IpolygondatacommitteeerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygondatacommitteeerrorsCallerRaw struct {
	Contract *IpolygondatacommitteeerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygondatacommitteeerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygondatacommitteeerrorsTransactorRaw struct {
	Contract *IpolygondatacommitteeerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygondatacommitteeerrors creates a new instance of Ipolygondatacommitteeerrors, bound to a specific deployed contract.
func NewIpolygondatacommitteeerrors(address common.Address, backend bind.ContractBackend) (*Ipolygondatacommitteeerrors, error) {
	contract, err := bindIpolygondatacommitteeerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygondatacommitteeerrors{IpolygondatacommitteeerrorsCaller: IpolygondatacommitteeerrorsCaller{contract: contract}, IpolygondatacommitteeerrorsTransactor: IpolygondatacommitteeerrorsTransactor{contract: contract}, IpolygondatacommitteeerrorsFilterer: IpolygondatacommitteeerrorsFilterer{contract: contract}}, nil
}

// NewIpolygondatacommitteeerrorsCaller creates a new read-only instance of Ipolygondatacommitteeerrors, bound to a specific deployed contract.
func NewIpolygondatacommitteeerrorsCaller(address common.Address, caller bind.ContractCaller) (*IpolygondatacommitteeerrorsCaller, error) {
	contract, err := bindIpolygondatacommitteeerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygondatacommitteeerrorsCaller{contract: contract}, nil
}

// NewIpolygondatacommitteeerrorsTransactor creates a new write-only instance of Ipolygondatacommitteeerrors, bound to a specific deployed contract.
func NewIpolygondatacommitteeerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygondatacommitteeerrorsTransactor, error) {
	contract, err := bindIpolygondatacommitteeerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygondatacommitteeerrorsTransactor{contract: contract}, nil
}

// NewIpolygondatacommitteeerrorsFilterer creates a new log filterer instance of Ipolygondatacommitteeerrors, bound to a specific deployed contract.
func NewIpolygondatacommitteeerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygondatacommitteeerrorsFilterer, error) {
	contract, err := bindIpolygondatacommitteeerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygondatacommitteeerrorsFilterer{contract: contract}, nil
}

// bindIpolygondatacommitteeerrors binds a generic wrapper to an already deployed contract.
func bindIpolygondatacommitteeerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygondatacommitteeerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygondatacommitteeerrors.Contract.IpolygondatacommitteeerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygondatacommitteeerrors.Contract.IpolygondatacommitteeerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygondatacommitteeerrors.Contract.IpolygondatacommitteeerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygondatacommitteeerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygondatacommitteeerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygondatacommitteeerrors *IpolygondatacommitteeerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygondatacommitteeerrors.Contract.contract.Transact(opts, method, params...)
}

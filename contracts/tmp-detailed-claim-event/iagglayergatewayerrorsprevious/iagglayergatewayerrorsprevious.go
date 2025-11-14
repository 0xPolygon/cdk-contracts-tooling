// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergatewayerrorsprevious

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

// IagglayergatewayerrorspreviousMetaData contains all meta data concerning the Iagglayergatewayerrorsprevious contract.
var IagglayergatewayerrorspreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"}]",
}

// IagglayergatewayerrorspreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewayerrorspreviousMetaData.ABI instead.
var IagglayergatewayerrorspreviousABI = IagglayergatewayerrorspreviousMetaData.ABI

// Iagglayergatewayerrorsprevious is an auto generated Go binding around an Ethereum contract.
type Iagglayergatewayerrorsprevious struct {
	IagglayergatewayerrorspreviousCaller     // Read-only binding to the contract
	IagglayergatewayerrorspreviousTransactor // Write-only binding to the contract
	IagglayergatewayerrorspreviousFilterer   // Log filterer for contract events
}

// IagglayergatewayerrorspreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewayerrorspreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorspreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewayerrorspreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorspreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewayerrorspreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorspreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewayerrorspreviousSession struct {
	Contract     *Iagglayergatewayerrorsprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IagglayergatewayerrorspreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewayerrorspreviousCallerSession struct {
	Contract *IagglayergatewayerrorspreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// IagglayergatewayerrorspreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewayerrorspreviousTransactorSession struct {
	Contract     *IagglayergatewayerrorspreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// IagglayergatewayerrorspreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewayerrorspreviousRaw struct {
	Contract *Iagglayergatewayerrorsprevious // Generic contract binding to access the raw methods on
}

// IagglayergatewayerrorspreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewayerrorspreviousCallerRaw struct {
	Contract *IagglayergatewayerrorspreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewayerrorspreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewayerrorspreviousTransactorRaw struct {
	Contract *IagglayergatewayerrorspreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergatewayerrorsprevious creates a new instance of Iagglayergatewayerrorsprevious, bound to a specific deployed contract.
func NewIagglayergatewayerrorsprevious(address common.Address, backend bind.ContractBackend) (*Iagglayergatewayerrorsprevious, error) {
	contract, err := bindIagglayergatewayerrorsprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergatewayerrorsprevious{IagglayergatewayerrorspreviousCaller: IagglayergatewayerrorspreviousCaller{contract: contract}, IagglayergatewayerrorspreviousTransactor: IagglayergatewayerrorspreviousTransactor{contract: contract}, IagglayergatewayerrorspreviousFilterer: IagglayergatewayerrorspreviousFilterer{contract: contract}}, nil
}

// NewIagglayergatewayerrorspreviousCaller creates a new read-only instance of Iagglayergatewayerrorsprevious, bound to a specific deployed contract.
func NewIagglayergatewayerrorspreviousCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewayerrorspreviousCaller, error) {
	contract, err := bindIagglayergatewayerrorsprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorspreviousCaller{contract: contract}, nil
}

// NewIagglayergatewayerrorspreviousTransactor creates a new write-only instance of Iagglayergatewayerrorsprevious, bound to a specific deployed contract.
func NewIagglayergatewayerrorspreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewayerrorspreviousTransactor, error) {
	contract, err := bindIagglayergatewayerrorsprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorspreviousTransactor{contract: contract}, nil
}

// NewIagglayergatewayerrorspreviousFilterer creates a new log filterer instance of Iagglayergatewayerrorsprevious, bound to a specific deployed contract.
func NewIagglayergatewayerrorspreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewayerrorspreviousFilterer, error) {
	contract, err := bindIagglayergatewayerrorsprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorspreviousFilterer{contract: contract}, nil
}

// bindIagglayergatewayerrorsprevious binds a generic wrapper to an already deployed contract.
func bindIagglayergatewayerrorsprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewayerrorspreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayerrorsprevious.Contract.IagglayergatewayerrorspreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayerrorsprevious.Contract.IagglayergatewayerrorspreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayerrorsprevious.Contract.IagglayergatewayerrorspreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayerrorsprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayerrorsprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayerrorsprevious *IagglayergatewayerrorspreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayerrorsprevious.Contract.contract.Transact(opts, method, params...)
}

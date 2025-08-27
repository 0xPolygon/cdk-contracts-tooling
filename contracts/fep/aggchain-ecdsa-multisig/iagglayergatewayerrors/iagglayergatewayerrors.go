// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayergatewayerrors

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

// IagglayergatewayerrorsMetaData contains all meta data concerning the Iagglayergatewayerrors contract.
var IagglayergatewayerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofBytesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggLayerAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PPSelectorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"RouteAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsAlreadyFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteIsFrozen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"RouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VKeyCannotBeZero\",\"type\":\"error\"}]",
}

// IagglayergatewayerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayergatewayerrorsMetaData.ABI instead.
var IagglayergatewayerrorsABI = IagglayergatewayerrorsMetaData.ABI

// Iagglayergatewayerrors is an auto generated Go binding around an Ethereum contract.
type Iagglayergatewayerrors struct {
	IagglayergatewayerrorsCaller     // Read-only binding to the contract
	IagglayergatewayerrorsTransactor // Write-only binding to the contract
	IagglayergatewayerrorsFilterer   // Log filterer for contract events
}

// IagglayergatewayerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayergatewayerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayergatewayerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayergatewayerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayergatewayerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayergatewayerrorsSession struct {
	Contract     *Iagglayergatewayerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IagglayergatewayerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayergatewayerrorsCallerSession struct {
	Contract *IagglayergatewayerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IagglayergatewayerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayergatewayerrorsTransactorSession struct {
	Contract     *IagglayergatewayerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IagglayergatewayerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayergatewayerrorsRaw struct {
	Contract *Iagglayergatewayerrors // Generic contract binding to access the raw methods on
}

// IagglayergatewayerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayergatewayerrorsCallerRaw struct {
	Contract *IagglayergatewayerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayergatewayerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayergatewayerrorsTransactorRaw struct {
	Contract *IagglayergatewayerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayergatewayerrors creates a new instance of Iagglayergatewayerrors, bound to a specific deployed contract.
func NewIagglayergatewayerrors(address common.Address, backend bind.ContractBackend) (*Iagglayergatewayerrors, error) {
	contract, err := bindIagglayergatewayerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayergatewayerrors{IagglayergatewayerrorsCaller: IagglayergatewayerrorsCaller{contract: contract}, IagglayergatewayerrorsTransactor: IagglayergatewayerrorsTransactor{contract: contract}, IagglayergatewayerrorsFilterer: IagglayergatewayerrorsFilterer{contract: contract}}, nil
}

// NewIagglayergatewayerrorsCaller creates a new read-only instance of Iagglayergatewayerrors, bound to a specific deployed contract.
func NewIagglayergatewayerrorsCaller(address common.Address, caller bind.ContractCaller) (*IagglayergatewayerrorsCaller, error) {
	contract, err := bindIagglayergatewayerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorsCaller{contract: contract}, nil
}

// NewIagglayergatewayerrorsTransactor creates a new write-only instance of Iagglayergatewayerrors, bound to a specific deployed contract.
func NewIagglayergatewayerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayergatewayerrorsTransactor, error) {
	contract, err := bindIagglayergatewayerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorsTransactor{contract: contract}, nil
}

// NewIagglayergatewayerrorsFilterer creates a new log filterer instance of Iagglayergatewayerrors, bound to a specific deployed contract.
func NewIagglayergatewayerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayergatewayerrorsFilterer, error) {
	contract, err := bindIagglayergatewayerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayergatewayerrorsFilterer{contract: contract}, nil
}

// bindIagglayergatewayerrors binds a generic wrapper to an already deployed contract.
func bindIagglayergatewayerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayergatewayerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayerrors.Contract.IagglayergatewayerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayerrors.Contract.IagglayergatewayerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayerrors.Contract.IagglayergatewayerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayergatewayerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayergatewayerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayergatewayerrors *IagglayergatewayerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayergatewayerrors.Contract.contract.Transact(opts, method, params...)
}

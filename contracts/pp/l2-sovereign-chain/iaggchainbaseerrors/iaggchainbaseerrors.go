// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainbaseerrors

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

// IaggchainbaseerrorsMetaData contains all meta data concerning the Iaggchainbaseerrors contract.
var IaggchainbaseerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"}]",
}

// IaggchainbaseerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainbaseerrorsMetaData.ABI instead.
var IaggchainbaseerrorsABI = IaggchainbaseerrorsMetaData.ABI

// Iaggchainbaseerrors is an auto generated Go binding around an Ethereum contract.
type Iaggchainbaseerrors struct {
	IaggchainbaseerrorsCaller     // Read-only binding to the contract
	IaggchainbaseerrorsTransactor // Write-only binding to the contract
	IaggchainbaseerrorsFilterer   // Log filterer for contract events
}

// IaggchainbaseerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainbaseerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainbaseerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainbaseerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainbaseerrorsSession struct {
	Contract     *Iaggchainbaseerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IaggchainbaseerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainbaseerrorsCallerSession struct {
	Contract *IaggchainbaseerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IaggchainbaseerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainbaseerrorsTransactorSession struct {
	Contract     *IaggchainbaseerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IaggchainbaseerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainbaseerrorsRaw struct {
	Contract *Iaggchainbaseerrors // Generic contract binding to access the raw methods on
}

// IaggchainbaseerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainbaseerrorsCallerRaw struct {
	Contract *IaggchainbaseerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainbaseerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainbaseerrorsTransactorRaw struct {
	Contract *IaggchainbaseerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainbaseerrors creates a new instance of Iaggchainbaseerrors, bound to a specific deployed contract.
func NewIaggchainbaseerrors(address common.Address, backend bind.ContractBackend) (*Iaggchainbaseerrors, error) {
	contract, err := bindIaggchainbaseerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainbaseerrors{IaggchainbaseerrorsCaller: IaggchainbaseerrorsCaller{contract: contract}, IaggchainbaseerrorsTransactor: IaggchainbaseerrorsTransactor{contract: contract}, IaggchainbaseerrorsFilterer: IaggchainbaseerrorsFilterer{contract: contract}}, nil
}

// NewIaggchainbaseerrorsCaller creates a new read-only instance of Iaggchainbaseerrors, bound to a specific deployed contract.
func NewIaggchainbaseerrorsCaller(address common.Address, caller bind.ContractCaller) (*IaggchainbaseerrorsCaller, error) {
	contract, err := bindIaggchainbaseerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorsCaller{contract: contract}, nil
}

// NewIaggchainbaseerrorsTransactor creates a new write-only instance of Iaggchainbaseerrors, bound to a specific deployed contract.
func NewIaggchainbaseerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainbaseerrorsTransactor, error) {
	contract, err := bindIaggchainbaseerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorsTransactor{contract: contract}, nil
}

// NewIaggchainbaseerrorsFilterer creates a new log filterer instance of Iaggchainbaseerrors, bound to a specific deployed contract.
func NewIaggchainbaseerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainbaseerrorsFilterer, error) {
	contract, err := bindIaggchainbaseerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorsFilterer{contract: contract}, nil
}

// bindIaggchainbaseerrors binds a generic wrapper to an already deployed contract.
func bindIaggchainbaseerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainbaseerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseerrors.Contract.IaggchainbaseerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseerrors.Contract.IaggchainbaseerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseerrors.Contract.IaggchainbaseerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseerrors *IaggchainbaseerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseerrors.Contract.contract.Transact(opts, method, params...)
}

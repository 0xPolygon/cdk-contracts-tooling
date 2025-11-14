// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchainbaseerrorsprevious

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

// IaggchainbaseerrorspreviousMetaData contains all meta data concerning the Iaggchainbaseerrorsprevious contract.
var IaggchainbaseerrorspreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAgglayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"}]",
}

// IaggchainbaseerrorspreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainbaseerrorspreviousMetaData.ABI instead.
var IaggchainbaseerrorspreviousABI = IaggchainbaseerrorspreviousMetaData.ABI

// Iaggchainbaseerrorsprevious is an auto generated Go binding around an Ethereum contract.
type Iaggchainbaseerrorsprevious struct {
	IaggchainbaseerrorspreviousCaller     // Read-only binding to the contract
	IaggchainbaseerrorspreviousTransactor // Write-only binding to the contract
	IaggchainbaseerrorspreviousFilterer   // Log filterer for contract events
}

// IaggchainbaseerrorspreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainbaseerrorspreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorspreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainbaseerrorspreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorspreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainbaseerrorspreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainbaseerrorspreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainbaseerrorspreviousSession struct {
	Contract     *Iaggchainbaseerrorsprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IaggchainbaseerrorspreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainbaseerrorspreviousCallerSession struct {
	Contract *IaggchainbaseerrorspreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IaggchainbaseerrorspreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainbaseerrorspreviousTransactorSession struct {
	Contract     *IaggchainbaseerrorspreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IaggchainbaseerrorspreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainbaseerrorspreviousRaw struct {
	Contract *Iaggchainbaseerrorsprevious // Generic contract binding to access the raw methods on
}

// IaggchainbaseerrorspreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainbaseerrorspreviousCallerRaw struct {
	Contract *IaggchainbaseerrorspreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainbaseerrorspreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainbaseerrorspreviousTransactorRaw struct {
	Contract *IaggchainbaseerrorspreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchainbaseerrorsprevious creates a new instance of Iaggchainbaseerrorsprevious, bound to a specific deployed contract.
func NewIaggchainbaseerrorsprevious(address common.Address, backend bind.ContractBackend) (*Iaggchainbaseerrorsprevious, error) {
	contract, err := bindIaggchainbaseerrorsprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchainbaseerrorsprevious{IaggchainbaseerrorspreviousCaller: IaggchainbaseerrorspreviousCaller{contract: contract}, IaggchainbaseerrorspreviousTransactor: IaggchainbaseerrorspreviousTransactor{contract: contract}, IaggchainbaseerrorspreviousFilterer: IaggchainbaseerrorspreviousFilterer{contract: contract}}, nil
}

// NewIaggchainbaseerrorspreviousCaller creates a new read-only instance of Iaggchainbaseerrorsprevious, bound to a specific deployed contract.
func NewIaggchainbaseerrorspreviousCaller(address common.Address, caller bind.ContractCaller) (*IaggchainbaseerrorspreviousCaller, error) {
	contract, err := bindIaggchainbaseerrorsprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorspreviousCaller{contract: contract}, nil
}

// NewIaggchainbaseerrorspreviousTransactor creates a new write-only instance of Iaggchainbaseerrorsprevious, bound to a specific deployed contract.
func NewIaggchainbaseerrorspreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainbaseerrorspreviousTransactor, error) {
	contract, err := bindIaggchainbaseerrorsprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorspreviousTransactor{contract: contract}, nil
}

// NewIaggchainbaseerrorspreviousFilterer creates a new log filterer instance of Iaggchainbaseerrorsprevious, bound to a specific deployed contract.
func NewIaggchainbaseerrorspreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainbaseerrorspreviousFilterer, error) {
	contract, err := bindIaggchainbaseerrorsprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainbaseerrorspreviousFilterer{contract: contract}, nil
}

// bindIaggchainbaseerrorsprevious binds a generic wrapper to an already deployed contract.
func bindIaggchainbaseerrorsprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainbaseerrorspreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseerrorsprevious.Contract.IaggchainbaseerrorspreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseerrorsprevious.Contract.IaggchainbaseerrorspreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseerrorsprevious.Contract.IaggchainbaseerrorspreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchainbaseerrorsprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchainbaseerrorsprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchainbaseerrorsprevious *IaggchainbaseerrorspreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchainbaseerrorsprevious.Contract.contract.Transact(opts, method, params...)
}

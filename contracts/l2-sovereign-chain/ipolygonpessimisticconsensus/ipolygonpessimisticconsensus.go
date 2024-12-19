// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonpessimisticconsensus

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

// IpolygonpessimisticconsensusMetaData contains all meta data concerning the Ipolygonpessimisticconsensus contract.
var IpolygonpessimisticconsensusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getConsensusHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IpolygonpessimisticconsensusABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonpessimisticconsensusMetaData.ABI instead.
var IpolygonpessimisticconsensusABI = IpolygonpessimisticconsensusMetaData.ABI

// Ipolygonpessimisticconsensus is an auto generated Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensus struct {
	IpolygonpessimisticconsensusCaller     // Read-only binding to the contract
	IpolygonpessimisticconsensusTransactor // Write-only binding to the contract
	IpolygonpessimisticconsensusFilterer   // Log filterer for contract events
}

// IpolygonpessimisticconsensusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonpessimisticconsensusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonpessimisticconsensusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonpessimisticconsensusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonpessimisticconsensusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonpessimisticconsensusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonpessimisticconsensusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonpessimisticconsensusSession struct {
	Contract     *Ipolygonpessimisticconsensus // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IpolygonpessimisticconsensusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonpessimisticconsensusCallerSession struct {
	Contract *IpolygonpessimisticconsensusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// IpolygonpessimisticconsensusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonpessimisticconsensusTransactorSession struct {
	Contract     *IpolygonpessimisticconsensusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// IpolygonpessimisticconsensusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonpessimisticconsensusRaw struct {
	Contract *Ipolygonpessimisticconsensus // Generic contract binding to access the raw methods on
}

// IpolygonpessimisticconsensusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonpessimisticconsensusCallerRaw struct {
	Contract *IpolygonpessimisticconsensusCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonpessimisticconsensusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonpessimisticconsensusTransactorRaw struct {
	Contract *IpolygonpessimisticconsensusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonpessimisticconsensus creates a new instance of Ipolygonpessimisticconsensus, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensus(address common.Address, backend bind.ContractBackend) (*Ipolygonpessimisticconsensus, error) {
	contract, err := bindIpolygonpessimisticconsensus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonpessimisticconsensus{IpolygonpessimisticconsensusCaller: IpolygonpessimisticconsensusCaller{contract: contract}, IpolygonpessimisticconsensusTransactor: IpolygonpessimisticconsensusTransactor{contract: contract}, IpolygonpessimisticconsensusFilterer: IpolygonpessimisticconsensusFilterer{contract: contract}}, nil
}

// NewIpolygonpessimisticconsensusCaller creates a new read-only instance of Ipolygonpessimisticconsensus, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusCaller(address common.Address, caller bind.ContractCaller) (*IpolygonpessimisticconsensusCaller, error) {
	contract, err := bindIpolygonpessimisticconsensus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonpessimisticconsensusCaller{contract: contract}, nil
}

// NewIpolygonpessimisticconsensusTransactor creates a new write-only instance of Ipolygonpessimisticconsensus, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonpessimisticconsensusTransactor, error) {
	contract, err := bindIpolygonpessimisticconsensus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonpessimisticconsensusTransactor{contract: contract}, nil
}

// NewIpolygonpessimisticconsensusFilterer creates a new log filterer instance of Ipolygonpessimisticconsensus, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonpessimisticconsensusFilterer, error) {
	contract, err := bindIpolygonpessimisticconsensus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonpessimisticconsensusFilterer{contract: contract}, nil
}

// bindIpolygonpessimisticconsensus binds a generic wrapper to an already deployed contract.
func bindIpolygonpessimisticconsensus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonpessimisticconsensusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonpessimisticconsensus.Contract.IpolygonpessimisticconsensusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensus.Contract.IpolygonpessimisticconsensusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensus.Contract.IpolygonpessimisticconsensusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonpessimisticconsensus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensus.Contract.contract.Transact(opts, method, params...)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusCaller) GetConsensusHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonpessimisticconsensus.contract.Call(opts, &out, "getConsensusHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusSession) GetConsensusHash() ([32]byte, error) {
	return _Ipolygonpessimisticconsensus.Contract.GetConsensusHash(&_Ipolygonpessimisticconsensus.CallOpts)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xad1edf34.
//
// Solidity: function getConsensusHash() view returns(bytes32)
func (_Ipolygonpessimisticconsensus *IpolygonpessimisticconsensusCallerSession) GetConsensusHash() ([32]byte, error) {
	return _Ipolygonpessimisticconsensus.Contract.GetConsensusHash(&_Ipolygonpessimisticconsensus.CallOpts)
}

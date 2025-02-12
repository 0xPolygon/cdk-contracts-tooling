// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonpessimisticconsensusv2

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

// Ipolygonpessimisticconsensusv2MetaData contains all meta data concerning the Ipolygonpessimisticconsensusv2 contract.
var Ipolygonpessimisticconsensusv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getConsensusHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onCustomChainData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ipolygonpessimisticconsensusv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ipolygonpessimisticconsensusv2MetaData.ABI instead.
var Ipolygonpessimisticconsensusv2ABI = Ipolygonpessimisticconsensusv2MetaData.ABI

// Ipolygonpessimisticconsensusv2 is an auto generated Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2 struct {
	Ipolygonpessimisticconsensusv2Caller     // Read-only binding to the contract
	Ipolygonpessimisticconsensusv2Transactor // Write-only binding to the contract
	Ipolygonpessimisticconsensusv2Filterer   // Log filterer for contract events
}

// Ipolygonpessimisticconsensusv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonpessimisticconsensusv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonpessimisticconsensusv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ipolygonpessimisticconsensusv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ipolygonpessimisticconsensusv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ipolygonpessimisticconsensusv2Session struct {
	Contract     *Ipolygonpessimisticconsensusv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Ipolygonpessimisticconsensusv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ipolygonpessimisticconsensusv2CallerSession struct {
	Contract *Ipolygonpessimisticconsensusv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// Ipolygonpessimisticconsensusv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ipolygonpessimisticconsensusv2TransactorSession struct {
	Contract     *Ipolygonpessimisticconsensusv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// Ipolygonpessimisticconsensusv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2Raw struct {
	Contract *Ipolygonpessimisticconsensusv2 // Generic contract binding to access the raw methods on
}

// Ipolygonpessimisticconsensusv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2CallerRaw struct {
	Contract *Ipolygonpessimisticconsensusv2Caller // Generic read-only contract binding to access the raw methods on
}

// Ipolygonpessimisticconsensusv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ipolygonpessimisticconsensusv2TransactorRaw struct {
	Contract *Ipolygonpessimisticconsensusv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonpessimisticconsensusv2 creates a new instance of Ipolygonpessimisticconsensusv2, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusv2(address common.Address, backend bind.ContractBackend) (*Ipolygonpessimisticconsensusv2, error) {
	contract, err := bindIpolygonpessimisticconsensusv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonpessimisticconsensusv2{Ipolygonpessimisticconsensusv2Caller: Ipolygonpessimisticconsensusv2Caller{contract: contract}, Ipolygonpessimisticconsensusv2Transactor: Ipolygonpessimisticconsensusv2Transactor{contract: contract}, Ipolygonpessimisticconsensusv2Filterer: Ipolygonpessimisticconsensusv2Filterer{contract: contract}}, nil
}

// NewIpolygonpessimisticconsensusv2Caller creates a new read-only instance of Ipolygonpessimisticconsensusv2, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusv2Caller(address common.Address, caller bind.ContractCaller) (*Ipolygonpessimisticconsensusv2Caller, error) {
	contract, err := bindIpolygonpessimisticconsensusv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonpessimisticconsensusv2Caller{contract: contract}, nil
}

// NewIpolygonpessimisticconsensusv2Transactor creates a new write-only instance of Ipolygonpessimisticconsensusv2, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Ipolygonpessimisticconsensusv2Transactor, error) {
	contract, err := bindIpolygonpessimisticconsensusv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ipolygonpessimisticconsensusv2Transactor{contract: contract}, nil
}

// NewIpolygonpessimisticconsensusv2Filterer creates a new log filterer instance of Ipolygonpessimisticconsensusv2, bound to a specific deployed contract.
func NewIpolygonpessimisticconsensusv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Ipolygonpessimisticconsensusv2Filterer, error) {
	contract, err := bindIpolygonpessimisticconsensusv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ipolygonpessimisticconsensusv2Filterer{contract: contract}, nil
}

// bindIpolygonpessimisticconsensusv2 binds a generic wrapper to an already deployed contract.
func bindIpolygonpessimisticconsensusv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ipolygonpessimisticconsensusv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonpessimisticconsensusv2.Contract.Ipolygonpessimisticconsensusv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.Ipolygonpessimisticconsensusv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.Ipolygonpessimisticconsensusv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonpessimisticconsensusv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.contract.Transact(opts, method, params...)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xa27aed88.
//
// Solidity: function getConsensusHash(bytes data) view returns(bytes32)
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Caller) GetConsensusHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonpessimisticconsensusv2.contract.Call(opts, &out, "getConsensusHash", data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConsensusHash is a free data retrieval call binding the contract method 0xa27aed88.
//
// Solidity: function getConsensusHash(bytes data) view returns(bytes32)
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Session) GetConsensusHash(data []byte) ([32]byte, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.GetConsensusHash(&_Ipolygonpessimisticconsensusv2.CallOpts, data)
}

// GetConsensusHash is a free data retrieval call binding the contract method 0xa27aed88.
//
// Solidity: function getConsensusHash(bytes data) view returns(bytes32)
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2CallerSession) GetConsensusHash(data []byte) ([32]byte, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.GetConsensusHash(&_Ipolygonpessimisticconsensusv2.CallOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Transactor) Initialize(opts *bind.TransactOpts, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.contract.Transact(opts, "initialize", initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Session) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.Initialize(&_Ipolygonpessimisticconsensusv2.TransactOpts, initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2TransactorSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.Initialize(&_Ipolygonpessimisticconsensusv2.TransactOpts, initializeBytesCustomChain)
}

// OnCustomChainData is a paid mutator transaction binding the contract method 0x1772a08b.
//
// Solidity: function onCustomChainData(bytes data) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Transactor) OnCustomChainData(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.contract.Transact(opts, "onCustomChainData", data)
}

// OnCustomChainData is a paid mutator transaction binding the contract method 0x1772a08b.
//
// Solidity: function onCustomChainData(bytes data) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2Session) OnCustomChainData(data []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.OnCustomChainData(&_Ipolygonpessimisticconsensusv2.TransactOpts, data)
}

// OnCustomChainData is a paid mutator transaction binding the contract method 0x1772a08b.
//
// Solidity: function onCustomChainData(bytes data) returns()
func (_Ipolygonpessimisticconsensusv2 *Ipolygonpessimisticconsensusv2TransactorSession) OnCustomChainData(data []byte) (*types.Transaction, error) {
	return _Ipolygonpessimisticconsensusv2.Contract.OnCustomChainData(&_Ipolygonpessimisticconsensusv2.TransactOpts, data)
}

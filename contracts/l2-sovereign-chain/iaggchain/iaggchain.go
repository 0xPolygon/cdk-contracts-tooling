// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggchain

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

// IaggchainMetaData contains all meta data concerning the Iaggchain contract.
var IaggchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IaggchainABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggchainMetaData.ABI instead.
var IaggchainABI = IaggchainMetaData.ABI

// Iaggchain is an auto generated Go binding around an Ethereum contract.
type Iaggchain struct {
	IaggchainCaller     // Read-only binding to the contract
	IaggchainTransactor // Write-only binding to the contract
	IaggchainFilterer   // Log filterer for contract events
}

// IaggchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggchainSession struct {
	Contract     *Iaggchain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IaggchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggchainCallerSession struct {
	Contract *IaggchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IaggchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggchainTransactorSession struct {
	Contract     *IaggchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IaggchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggchainRaw struct {
	Contract *Iaggchain // Generic contract binding to access the raw methods on
}

// IaggchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggchainCallerRaw struct {
	Contract *IaggchainCaller // Generic read-only contract binding to access the raw methods on
}

// IaggchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggchainTransactorRaw struct {
	Contract *IaggchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggchain creates a new instance of Iaggchain, bound to a specific deployed contract.
func NewIaggchain(address common.Address, backend bind.ContractBackend) (*Iaggchain, error) {
	contract, err := bindIaggchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggchain{IaggchainCaller: IaggchainCaller{contract: contract}, IaggchainTransactor: IaggchainTransactor{contract: contract}, IaggchainFilterer: IaggchainFilterer{contract: contract}}, nil
}

// NewIaggchainCaller creates a new read-only instance of Iaggchain, bound to a specific deployed contract.
func NewIaggchainCaller(address common.Address, caller bind.ContractCaller) (*IaggchainCaller, error) {
	contract, err := bindIaggchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainCaller{contract: contract}, nil
}

// NewIaggchainTransactor creates a new write-only instance of Iaggchain, bound to a specific deployed contract.
func NewIaggchainTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggchainTransactor, error) {
	contract, err := bindIaggchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggchainTransactor{contract: contract}, nil
}

// NewIaggchainFilterer creates a new log filterer instance of Iaggchain, bound to a specific deployed contract.
func NewIaggchainFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggchainFilterer, error) {
	contract, err := bindIaggchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggchainFilterer{contract: contract}, nil
}

// bindIaggchain binds a generic wrapper to an already deployed contract.
func bindIaggchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchain *IaggchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchain.Contract.IaggchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchain *IaggchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchain.Contract.IaggchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchain *IaggchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchain.Contract.IaggchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggchain *IaggchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggchain *IaggchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggchain *IaggchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggchain.Contract.contract.Transact(opts, method, params...)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Iaggchain *IaggchainCaller) GetAggchainHash(opts *bind.CallOpts, customChainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Iaggchain.contract.Call(opts, &out, "getAggchainHash", customChainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Iaggchain *IaggchainSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Iaggchain.Contract.GetAggchainHash(&_Iaggchain.CallOpts, customChainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes customChainData) view returns(bytes32)
func (_Iaggchain *IaggchainCallerSession) GetAggchainHash(customChainData []byte) ([32]byte, error) {
	return _Iaggchain.Contract.GetAggchainHash(&_Iaggchain.CallOpts, customChainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes data) returns()
func (_Iaggchain *IaggchainTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Iaggchain.contract.Transact(opts, "onVerifyPessimistic", data)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes data) returns()
func (_Iaggchain *IaggchainSession) OnVerifyPessimistic(data []byte) (*types.Transaction, error) {
	return _Iaggchain.Contract.OnVerifyPessimistic(&_Iaggchain.TransactOpts, data)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes data) returns()
func (_Iaggchain *IaggchainTransactorSession) OnVerifyPessimistic(data []byte) (*types.Transaction, error) {
	return _Iaggchain.Contract.OnVerifyPessimistic(&_Iaggchain.TransactOpts, data)
}

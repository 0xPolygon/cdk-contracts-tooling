// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iglobalexitrootmanagerl2sovereignchain

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

// Iglobalexitrootmanagerl2sovereignchainMetaData contains all meta data concerning the Iglobalexitrootmanagerl2sovereignchain contract.
var Iglobalexitrootmanagerl2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRootNum\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Iglobalexitrootmanagerl2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Iglobalexitrootmanagerl2sovereignchainMetaData.ABI instead.
var Iglobalexitrootmanagerl2sovereignchainABI = Iglobalexitrootmanagerl2sovereignchainMetaData.ABI

// Iglobalexitrootmanagerl2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchain struct {
	Iglobalexitrootmanagerl2sovereignchainCaller     // Read-only binding to the contract
	Iglobalexitrootmanagerl2sovereignchainTransactor // Write-only binding to the contract
	Iglobalexitrootmanagerl2sovereignchainFilterer   // Log filterer for contract events
}

// Iglobalexitrootmanagerl2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iglobalexitrootmanagerl2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iglobalexitrootmanagerl2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iglobalexitrootmanagerl2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iglobalexitrootmanagerl2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iglobalexitrootmanagerl2sovereignchainSession struct {
	Contract     *Iglobalexitrootmanagerl2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                           // Call options to use throughout this session
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// Iglobalexitrootmanagerl2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iglobalexitrootmanagerl2sovereignchainCallerSession struct {
	Contract *Iglobalexitrootmanagerl2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                 // Call options to use throughout this session
}

// Iglobalexitrootmanagerl2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iglobalexitrootmanagerl2sovereignchainTransactorSession struct {
	Contract     *Iglobalexitrootmanagerl2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                 // Transaction auth options to use throughout this session
}

// Iglobalexitrootmanagerl2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchainRaw struct {
	Contract *Iglobalexitrootmanagerl2sovereignchain // Generic contract binding to access the raw methods on
}

// Iglobalexitrootmanagerl2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchainCallerRaw struct {
	Contract *Iglobalexitrootmanagerl2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Iglobalexitrootmanagerl2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iglobalexitrootmanagerl2sovereignchainTransactorRaw struct {
	Contract *Iglobalexitrootmanagerl2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIglobalexitrootmanagerl2sovereignchain creates a new instance of Iglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewIglobalexitrootmanagerl2sovereignchain(address common.Address, backend bind.ContractBackend) (*Iglobalexitrootmanagerl2sovereignchain, error) {
	contract, err := bindIglobalexitrootmanagerl2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iglobalexitrootmanagerl2sovereignchain{Iglobalexitrootmanagerl2sovereignchainCaller: Iglobalexitrootmanagerl2sovereignchainCaller{contract: contract}, Iglobalexitrootmanagerl2sovereignchainTransactor: Iglobalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Iglobalexitrootmanagerl2sovereignchainFilterer: Iglobalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// NewIglobalexitrootmanagerl2sovereignchainCaller creates a new read-only instance of Iglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewIglobalexitrootmanagerl2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Iglobalexitrootmanagerl2sovereignchainCaller, error) {
	contract, err := bindIglobalexitrootmanagerl2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iglobalexitrootmanagerl2sovereignchainCaller{contract: contract}, nil
}

// NewIglobalexitrootmanagerl2sovereignchainTransactor creates a new write-only instance of Iglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewIglobalexitrootmanagerl2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Iglobalexitrootmanagerl2sovereignchainTransactor, error) {
	contract, err := bindIglobalexitrootmanagerl2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iglobalexitrootmanagerl2sovereignchainTransactor{contract: contract}, nil
}

// NewIglobalexitrootmanagerl2sovereignchainFilterer creates a new log filterer instance of Iglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewIglobalexitrootmanagerl2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Iglobalexitrootmanagerl2sovereignchainFilterer, error) {
	contract, err := bindIglobalexitrootmanagerl2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iglobalexitrootmanagerl2sovereignchainFilterer{contract: contract}, nil
}

// bindIglobalexitrootmanagerl2sovereignchain binds a generic wrapper to an already deployed contract.
func bindIglobalexitrootmanagerl2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Iglobalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.Iglobalexitrootmanagerl2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.Iglobalexitrootmanagerl2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.Iglobalexitrootmanagerl2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainCaller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainSession) GlobalExitRootRemover() (common.Address, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Iglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Iglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactor) GlobalExitRootMap(opts *bind.TransactOpts, globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "globalExitRootMap", globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Iglobalexitrootmanagerl2sovereignchain.TransactOpts, globalExitRootNum)
}

// GlobalExitRootMap is a paid mutator transaction binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 globalExitRootNum) returns(uint256)
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactorSession) GlobalExitRootMap(globalExitRootNum [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Iglobalexitrootmanagerl2sovereignchain.TransactOpts, globalExitRootNum)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "updateExitRoot", newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Iglobalexitrootmanagerl2sovereignchain.TransactOpts, newRollupExitRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRollupExitRoot) returns()
func (_Iglobalexitrootmanagerl2sovereignchain *Iglobalexitrootmanagerl2sovereignchainTransactorSession) UpdateExitRoot(newRollupExitRoot [32]byte) (*types.Transaction, error) {
	return _Iglobalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Iglobalexitrootmanagerl2sovereignchain.TransactOpts, newRollupExitRoot)
}

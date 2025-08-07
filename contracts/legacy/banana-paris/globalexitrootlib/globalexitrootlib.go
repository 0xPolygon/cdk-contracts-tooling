// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootlib

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

// GlobalexitrootlibMetaData contains all meta data concerning the Globalexitrootlib contract.
var GlobalexitrootlibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a567d4ea19d8ea8c8817e40abbeca2bddd76293fdacac4614b28a8a6d0c64ccd64736f6c63430008140033",
}

// GlobalexitrootlibABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalexitrootlibMetaData.ABI instead.
var GlobalexitrootlibABI = GlobalexitrootlibMetaData.ABI

// GlobalexitrootlibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlobalexitrootlibMetaData.Bin instead.
var GlobalexitrootlibBin = GlobalexitrootlibMetaData.Bin

// DeployGlobalexitrootlib deploys a new Ethereum contract, binding an instance of Globalexitrootlib to it.
func DeployGlobalexitrootlib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Globalexitrootlib, error) {
	parsed, err := GlobalexitrootlibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlobalexitrootlibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootlib{GlobalexitrootlibCaller: GlobalexitrootlibCaller{contract: contract}, GlobalexitrootlibTransactor: GlobalexitrootlibTransactor{contract: contract}, GlobalexitrootlibFilterer: GlobalexitrootlibFilterer{contract: contract}}, nil
}

// Globalexitrootlib is an auto generated Go binding around an Ethereum contract.
type Globalexitrootlib struct {
	GlobalexitrootlibCaller     // Read-only binding to the contract
	GlobalexitrootlibTransactor // Write-only binding to the contract
	GlobalexitrootlibFilterer   // Log filterer for contract events
}

// GlobalexitrootlibCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalexitrootlibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalexitrootlibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalexitrootlibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalexitrootlibSession struct {
	Contract     *Globalexitrootlib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GlobalexitrootlibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalexitrootlibCallerSession struct {
	Contract *GlobalexitrootlibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GlobalexitrootlibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalexitrootlibTransactorSession struct {
	Contract     *GlobalexitrootlibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GlobalexitrootlibRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalexitrootlibRaw struct {
	Contract *Globalexitrootlib // Generic contract binding to access the raw methods on
}

// GlobalexitrootlibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalexitrootlibCallerRaw struct {
	Contract *GlobalexitrootlibCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalexitrootlibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalexitrootlibTransactorRaw struct {
	Contract *GlobalexitrootlibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootlib creates a new instance of Globalexitrootlib, bound to a specific deployed contract.
func NewGlobalexitrootlib(address common.Address, backend bind.ContractBackend) (*Globalexitrootlib, error) {
	contract, err := bindGlobalexitrootlib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootlib{GlobalexitrootlibCaller: GlobalexitrootlibCaller{contract: contract}, GlobalexitrootlibTransactor: GlobalexitrootlibTransactor{contract: contract}, GlobalexitrootlibFilterer: GlobalexitrootlibFilterer{contract: contract}}, nil
}

// NewGlobalexitrootlibCaller creates a new read-only instance of Globalexitrootlib, bound to a specific deployed contract.
func NewGlobalexitrootlibCaller(address common.Address, caller bind.ContractCaller) (*GlobalexitrootlibCaller, error) {
	contract, err := bindGlobalexitrootlib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibCaller{contract: contract}, nil
}

// NewGlobalexitrootlibTransactor creates a new write-only instance of Globalexitrootlib, bound to a specific deployed contract.
func NewGlobalexitrootlibTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalexitrootlibTransactor, error) {
	contract, err := bindGlobalexitrootlib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibTransactor{contract: contract}, nil
}

// NewGlobalexitrootlibFilterer creates a new log filterer instance of Globalexitrootlib, bound to a specific deployed contract.
func NewGlobalexitrootlibFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalexitrootlibFilterer, error) {
	contract, err := bindGlobalexitrootlib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibFilterer{contract: contract}, nil
}

// bindGlobalexitrootlib binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootlib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlobalexitrootlibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootlib *GlobalexitrootlibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootlib.Contract.GlobalexitrootlibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootlib *GlobalexitrootlibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootlib.Contract.GlobalexitrootlibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootlib *GlobalexitrootlibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootlib.Contract.GlobalexitrootlibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootlib *GlobalexitrootlibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootlib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootlib *GlobalexitrootlibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootlib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootlib *GlobalexitrootlibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootlib.Contract.contract.Transact(opts, method, params...)
}

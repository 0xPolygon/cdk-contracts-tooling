// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itokenwrappedbridgeinitcode

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

// ItokenwrappedbridgeinitcodeMetaData contains all meta data concerning the Itokenwrappedbridgeinitcode contract.
var ItokenwrappedbridgeinitcodeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BASE_INIT_BYTECODE_WRAPPED_TOKEN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ItokenwrappedbridgeinitcodeABI is the input ABI used to generate the binding from.
// Deprecated: Use ItokenwrappedbridgeinitcodeMetaData.ABI instead.
var ItokenwrappedbridgeinitcodeABI = ItokenwrappedbridgeinitcodeMetaData.ABI

// Itokenwrappedbridgeinitcode is an auto generated Go binding around an Ethereum contract.
type Itokenwrappedbridgeinitcode struct {
	ItokenwrappedbridgeinitcodeCaller     // Read-only binding to the contract
	ItokenwrappedbridgeinitcodeTransactor // Write-only binding to the contract
	ItokenwrappedbridgeinitcodeFilterer   // Log filterer for contract events
}

// ItokenwrappedbridgeinitcodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeinitcodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeinitcodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeinitcodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeinitcodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItokenwrappedbridgeinitcodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeinitcodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItokenwrappedbridgeinitcodeSession struct {
	Contract     *Itokenwrappedbridgeinitcode // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ItokenwrappedbridgeinitcodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItokenwrappedbridgeinitcodeCallerSession struct {
	Contract *ItokenwrappedbridgeinitcodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ItokenwrappedbridgeinitcodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItokenwrappedbridgeinitcodeTransactorSession struct {
	Contract     *ItokenwrappedbridgeinitcodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ItokenwrappedbridgeinitcodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItokenwrappedbridgeinitcodeRaw struct {
	Contract *Itokenwrappedbridgeinitcode // Generic contract binding to access the raw methods on
}

// ItokenwrappedbridgeinitcodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeinitcodeCallerRaw struct {
	Contract *ItokenwrappedbridgeinitcodeCaller // Generic read-only contract binding to access the raw methods on
}

// ItokenwrappedbridgeinitcodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeinitcodeTransactorRaw struct {
	Contract *ItokenwrappedbridgeinitcodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItokenwrappedbridgeinitcode creates a new instance of Itokenwrappedbridgeinitcode, bound to a specific deployed contract.
func NewItokenwrappedbridgeinitcode(address common.Address, backend bind.ContractBackend) (*Itokenwrappedbridgeinitcode, error) {
	contract, err := bindItokenwrappedbridgeinitcode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itokenwrappedbridgeinitcode{ItokenwrappedbridgeinitcodeCaller: ItokenwrappedbridgeinitcodeCaller{contract: contract}, ItokenwrappedbridgeinitcodeTransactor: ItokenwrappedbridgeinitcodeTransactor{contract: contract}, ItokenwrappedbridgeinitcodeFilterer: ItokenwrappedbridgeinitcodeFilterer{contract: contract}}, nil
}

// NewItokenwrappedbridgeinitcodeCaller creates a new read-only instance of Itokenwrappedbridgeinitcode, bound to a specific deployed contract.
func NewItokenwrappedbridgeinitcodeCaller(address common.Address, caller bind.ContractCaller) (*ItokenwrappedbridgeinitcodeCaller, error) {
	contract, err := bindItokenwrappedbridgeinitcode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeinitcodeCaller{contract: contract}, nil
}

// NewItokenwrappedbridgeinitcodeTransactor creates a new write-only instance of Itokenwrappedbridgeinitcode, bound to a specific deployed contract.
func NewItokenwrappedbridgeinitcodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ItokenwrappedbridgeinitcodeTransactor, error) {
	contract, err := bindItokenwrappedbridgeinitcode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeinitcodeTransactor{contract: contract}, nil
}

// NewItokenwrappedbridgeinitcodeFilterer creates a new log filterer instance of Itokenwrappedbridgeinitcode, bound to a specific deployed contract.
func NewItokenwrappedbridgeinitcodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ItokenwrappedbridgeinitcodeFilterer, error) {
	contract, err := bindItokenwrappedbridgeinitcode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeinitcodeFilterer{contract: contract}, nil
}

// bindItokenwrappedbridgeinitcode binds a generic wrapper to an already deployed contract.
func bindItokenwrappedbridgeinitcode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ItokenwrappedbridgeinitcodeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenwrappedbridgeinitcode.Contract.ItokenwrappedbridgeinitcodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenwrappedbridgeinitcode.Contract.ItokenwrappedbridgeinitcodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenwrappedbridgeinitcode.Contract.ItokenwrappedbridgeinitcodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenwrappedbridgeinitcode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenwrappedbridgeinitcode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenwrappedbridgeinitcode.Contract.contract.Transact(opts, method, params...)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() pure returns(bytes)
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeCaller) BASEINITBYTECODEWRAPPEDTOKEN(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeinitcode.contract.Call(opts, &out, "BASE_INIT_BYTECODE_WRAPPED_TOKEN")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() pure returns(bytes)
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Itokenwrappedbridgeinitcode.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Itokenwrappedbridgeinitcode.CallOpts)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() pure returns(bytes)
func (_Itokenwrappedbridgeinitcode *ItokenwrappedbridgeinitcodeCallerSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Itokenwrappedbridgeinitcode.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Itokenwrappedbridgeinitcode.CallOpts)
}

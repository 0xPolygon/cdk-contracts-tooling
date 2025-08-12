// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibytecodestorer

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

// IbytecodestorerMetaData contains all meta data concerning the Ibytecodestorer contract.
var IbytecodestorerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// IbytecodestorerABI is the input ABI used to generate the binding from.
// Deprecated: Use IbytecodestorerMetaData.ABI instead.
var IbytecodestorerABI = IbytecodestorerMetaData.ABI

// Ibytecodestorer is an auto generated Go binding around an Ethereum contract.
type Ibytecodestorer struct {
	IbytecodestorerCaller     // Read-only binding to the contract
	IbytecodestorerTransactor // Write-only binding to the contract
	IbytecodestorerFilterer   // Log filterer for contract events
}

// IbytecodestorerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbytecodestorerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbytecodestorerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbytecodestorerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbytecodestorerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbytecodestorerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbytecodestorerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbytecodestorerSession struct {
	Contract     *Ibytecodestorer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbytecodestorerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbytecodestorerCallerSession struct {
	Contract *IbytecodestorerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IbytecodestorerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbytecodestorerTransactorSession struct {
	Contract     *IbytecodestorerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IbytecodestorerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbytecodestorerRaw struct {
	Contract *Ibytecodestorer // Generic contract binding to access the raw methods on
}

// IbytecodestorerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbytecodestorerCallerRaw struct {
	Contract *IbytecodestorerCaller // Generic read-only contract binding to access the raw methods on
}

// IbytecodestorerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbytecodestorerTransactorRaw struct {
	Contract *IbytecodestorerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbytecodestorer creates a new instance of Ibytecodestorer, bound to a specific deployed contract.
func NewIbytecodestorer(address common.Address, backend bind.ContractBackend) (*Ibytecodestorer, error) {
	contract, err := bindIbytecodestorer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ibytecodestorer{IbytecodestorerCaller: IbytecodestorerCaller{contract: contract}, IbytecodestorerTransactor: IbytecodestorerTransactor{contract: contract}, IbytecodestorerFilterer: IbytecodestorerFilterer{contract: contract}}, nil
}

// NewIbytecodestorerCaller creates a new read-only instance of Ibytecodestorer, bound to a specific deployed contract.
func NewIbytecodestorerCaller(address common.Address, caller bind.ContractCaller) (*IbytecodestorerCaller, error) {
	contract, err := bindIbytecodestorer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbytecodestorerCaller{contract: contract}, nil
}

// NewIbytecodestorerTransactor creates a new write-only instance of Ibytecodestorer, bound to a specific deployed contract.
func NewIbytecodestorerTransactor(address common.Address, transactor bind.ContractTransactor) (*IbytecodestorerTransactor, error) {
	contract, err := bindIbytecodestorer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbytecodestorerTransactor{contract: contract}, nil
}

// NewIbytecodestorerFilterer creates a new log filterer instance of Ibytecodestorer, bound to a specific deployed contract.
func NewIbytecodestorerFilterer(address common.Address, filterer bind.ContractFilterer) (*IbytecodestorerFilterer, error) {
	contract, err := bindIbytecodestorer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbytecodestorerFilterer{contract: contract}, nil
}

// bindIbytecodestorer binds a generic wrapper to an already deployed contract.
func bindIbytecodestorer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbytecodestorerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibytecodestorer *IbytecodestorerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibytecodestorer.Contract.IbytecodestorerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibytecodestorer *IbytecodestorerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibytecodestorer.Contract.IbytecodestorerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibytecodestorer *IbytecodestorerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibytecodestorer.Contract.IbytecodestorerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ibytecodestorer *IbytecodestorerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ibytecodestorer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ibytecodestorer *IbytecodestorerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ibytecodestorer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ibytecodestorer *IbytecodestorerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ibytecodestorer.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Ibytecodestorer *IbytecodestorerCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Ibytecodestorer.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Ibytecodestorer *IbytecodestorerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Ibytecodestorer.Contract.INITBYTECODETRANSPARENTPROXY(&_Ibytecodestorer.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() pure returns(bytes)
func (_Ibytecodestorer *IbytecodestorerCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Ibytecodestorer.Contract.INITBYTECODETRANSPARENTPROXY(&_Ibytecodestorer.CallOpts)
}

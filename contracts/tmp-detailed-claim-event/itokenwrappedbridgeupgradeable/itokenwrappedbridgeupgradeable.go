// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package itokenwrappedbridgeupgradeable

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

// ItokenwrappedbridgeupgradeableMetaData contains all meta data concerning the Itokenwrappedbridgeupgradeable contract.
var ItokenwrappedbridgeupgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ItokenwrappedbridgeupgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ItokenwrappedbridgeupgradeableMetaData.ABI instead.
var ItokenwrappedbridgeupgradeableABI = ItokenwrappedbridgeupgradeableMetaData.ABI

// Itokenwrappedbridgeupgradeable is an auto generated Go binding around an Ethereum contract.
type Itokenwrappedbridgeupgradeable struct {
	ItokenwrappedbridgeupgradeableCaller     // Read-only binding to the contract
	ItokenwrappedbridgeupgradeableTransactor // Write-only binding to the contract
	ItokenwrappedbridgeupgradeableFilterer   // Log filterer for contract events
}

// ItokenwrappedbridgeupgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeupgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeupgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeupgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeupgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ItokenwrappedbridgeupgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ItokenwrappedbridgeupgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ItokenwrappedbridgeupgradeableSession struct {
	Contract     *Itokenwrappedbridgeupgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ItokenwrappedbridgeupgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ItokenwrappedbridgeupgradeableCallerSession struct {
	Contract *ItokenwrappedbridgeupgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ItokenwrappedbridgeupgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ItokenwrappedbridgeupgradeableTransactorSession struct {
	Contract     *ItokenwrappedbridgeupgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ItokenwrappedbridgeupgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ItokenwrappedbridgeupgradeableRaw struct {
	Contract *Itokenwrappedbridgeupgradeable // Generic contract binding to access the raw methods on
}

// ItokenwrappedbridgeupgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeupgradeableCallerRaw struct {
	Contract *ItokenwrappedbridgeupgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ItokenwrappedbridgeupgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ItokenwrappedbridgeupgradeableTransactorRaw struct {
	Contract *ItokenwrappedbridgeupgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewItokenwrappedbridgeupgradeable creates a new instance of Itokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewItokenwrappedbridgeupgradeable(address common.Address, backend bind.ContractBackend) (*Itokenwrappedbridgeupgradeable, error) {
	contract, err := bindItokenwrappedbridgeupgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Itokenwrappedbridgeupgradeable{ItokenwrappedbridgeupgradeableCaller: ItokenwrappedbridgeupgradeableCaller{contract: contract}, ItokenwrappedbridgeupgradeableTransactor: ItokenwrappedbridgeupgradeableTransactor{contract: contract}, ItokenwrappedbridgeupgradeableFilterer: ItokenwrappedbridgeupgradeableFilterer{contract: contract}}, nil
}

// NewItokenwrappedbridgeupgradeableCaller creates a new read-only instance of Itokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewItokenwrappedbridgeupgradeableCaller(address common.Address, caller bind.ContractCaller) (*ItokenwrappedbridgeupgradeableCaller, error) {
	contract, err := bindItokenwrappedbridgeupgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeupgradeableCaller{contract: contract}, nil
}

// NewItokenwrappedbridgeupgradeableTransactor creates a new write-only instance of Itokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewItokenwrappedbridgeupgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ItokenwrappedbridgeupgradeableTransactor, error) {
	contract, err := bindItokenwrappedbridgeupgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeupgradeableTransactor{contract: contract}, nil
}

// NewItokenwrappedbridgeupgradeableFilterer creates a new log filterer instance of Itokenwrappedbridgeupgradeable, bound to a specific deployed contract.
func NewItokenwrappedbridgeupgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ItokenwrappedbridgeupgradeableFilterer, error) {
	contract, err := bindItokenwrappedbridgeupgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeupgradeableFilterer{contract: contract}, nil
}

// bindItokenwrappedbridgeupgradeable binds a generic wrapper to an already deployed contract.
func bindItokenwrappedbridgeupgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ItokenwrappedbridgeupgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenwrappedbridgeupgradeable.Contract.ItokenwrappedbridgeupgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.ItokenwrappedbridgeupgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.ItokenwrappedbridgeupgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Itokenwrappedbridgeupgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.DOMAINSEPARATOR(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.DOMAINSEPARATOR(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Allowance(&_Itokenwrappedbridgeupgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Allowance(&_Itokenwrappedbridgeupgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.BalanceOf(&_Itokenwrappedbridgeupgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.BalanceOf(&_Itokenwrappedbridgeupgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Decimals() (uint8, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Decimals(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) Decimals() (uint8, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Decimals(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Name() (string, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Name(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) Name() (string, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Name(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Nonces(&_Itokenwrappedbridgeupgradeable.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Nonces(&_Itokenwrappedbridgeupgradeable.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Symbol() (string, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Symbol(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) Symbol() (string, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Symbol(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Itokenwrappedbridgeupgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) TotalSupply() (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.TotalSupply(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.TotalSupply(&_Itokenwrappedbridgeupgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Approve(&_Itokenwrappedbridgeupgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Approve(&_Itokenwrappedbridgeupgradeable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Burn(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "burn", account, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Burn(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Burn(&_Itokenwrappedbridgeupgradeable.TransactOpts, account, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Burn(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Burn(&_Itokenwrappedbridgeupgradeable.TransactOpts, account, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "initialize", name, symbol, __decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Initialize(name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Initialize(&_Itokenwrappedbridgeupgradeable.TransactOpts, name, symbol, __decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x1624f6c6.
//
// Solidity: function initialize(string name, string symbol, uint8 __decimals) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Initialize(name string, symbol string, __decimals uint8) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Initialize(&_Itokenwrappedbridgeupgradeable.TransactOpts, name, symbol, __decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Mint(&_Itokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Mint(&_Itokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Permit(&_Itokenwrappedbridgeupgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Permit(&_Itokenwrappedbridgeupgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Transfer(&_Itokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.Transfer(&_Itokenwrappedbridgeupgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.TransferFrom(&_Itokenwrappedbridgeupgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Itokenwrappedbridgeupgradeable.Contract.TransferFrom(&_Itokenwrappedbridgeupgradeable.TransactOpts, from, to, value)
}

// ItokenwrappedbridgeupgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Itokenwrappedbridgeupgradeable contract.
type ItokenwrappedbridgeupgradeableApprovalIterator struct {
	Event *ItokenwrappedbridgeupgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ItokenwrappedbridgeupgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokenwrappedbridgeupgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ItokenwrappedbridgeupgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ItokenwrappedbridgeupgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokenwrappedbridgeupgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokenwrappedbridgeupgradeableApproval represents a Approval event raised by the Itokenwrappedbridgeupgradeable contract.
type ItokenwrappedbridgeupgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ItokenwrappedbridgeupgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Itokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeupgradeableApprovalIterator{contract: _Itokenwrappedbridgeupgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ItokenwrappedbridgeupgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Itokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokenwrappedbridgeupgradeableApproval)
				if err := _Itokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) ParseApproval(log types.Log) (*ItokenwrappedbridgeupgradeableApproval, error) {
	event := new(ItokenwrappedbridgeupgradeableApproval)
	if err := _Itokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ItokenwrappedbridgeupgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Itokenwrappedbridgeupgradeable contract.
type ItokenwrappedbridgeupgradeableTransferIterator struct {
	Event *ItokenwrappedbridgeupgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ItokenwrappedbridgeupgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ItokenwrappedbridgeupgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ItokenwrappedbridgeupgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ItokenwrappedbridgeupgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ItokenwrappedbridgeupgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ItokenwrappedbridgeupgradeableTransfer represents a Transfer event raised by the Itokenwrappedbridgeupgradeable contract.
type ItokenwrappedbridgeupgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ItokenwrappedbridgeupgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Itokenwrappedbridgeupgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ItokenwrappedbridgeupgradeableTransferIterator{contract: _Itokenwrappedbridgeupgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ItokenwrappedbridgeupgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Itokenwrappedbridgeupgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ItokenwrappedbridgeupgradeableTransfer)
				if err := _Itokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Itokenwrappedbridgeupgradeable *ItokenwrappedbridgeupgradeableFilterer) ParseTransfer(log types.Log) (*ItokenwrappedbridgeupgradeableTransfer, error) {
	event := new(ItokenwrappedbridgeupgradeableTransfer)
	if err := _Itokenwrappedbridgeupgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

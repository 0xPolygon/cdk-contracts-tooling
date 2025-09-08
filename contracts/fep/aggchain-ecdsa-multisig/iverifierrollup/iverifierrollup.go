// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iverifierrollup

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

// IverifierrollupMetaData contains all meta data concerning the Iverifierrollup contract.
var IverifierrollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"},{\"internalType\":\"uint256[1]\",\"name\":\"pubSignals\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IverifierrollupABI is the input ABI used to generate the binding from.
// Deprecated: Use IverifierrollupMetaData.ABI instead.
var IverifierrollupABI = IverifierrollupMetaData.ABI

// Iverifierrollup is an auto generated Go binding around an Ethereum contract.
type Iverifierrollup struct {
	IverifierrollupCaller     // Read-only binding to the contract
	IverifierrollupTransactor // Write-only binding to the contract
	IverifierrollupFilterer   // Log filterer for contract events
}

// IverifierrollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IverifierrollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IverifierrollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IverifierrollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IverifierrollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IverifierrollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IverifierrollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IverifierrollupSession struct {
	Contract     *Iverifierrollup  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IverifierrollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IverifierrollupCallerSession struct {
	Contract *IverifierrollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IverifierrollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IverifierrollupTransactorSession struct {
	Contract     *IverifierrollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IverifierrollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IverifierrollupRaw struct {
	Contract *Iverifierrollup // Generic contract binding to access the raw methods on
}

// IverifierrollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IverifierrollupCallerRaw struct {
	Contract *IverifierrollupCaller // Generic read-only contract binding to access the raw methods on
}

// IverifierrollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IverifierrollupTransactorRaw struct {
	Contract *IverifierrollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIverifierrollup creates a new instance of Iverifierrollup, bound to a specific deployed contract.
func NewIverifierrollup(address common.Address, backend bind.ContractBackend) (*Iverifierrollup, error) {
	contract, err := bindIverifierrollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iverifierrollup{IverifierrollupCaller: IverifierrollupCaller{contract: contract}, IverifierrollupTransactor: IverifierrollupTransactor{contract: contract}, IverifierrollupFilterer: IverifierrollupFilterer{contract: contract}}, nil
}

// NewIverifierrollupCaller creates a new read-only instance of Iverifierrollup, bound to a specific deployed contract.
func NewIverifierrollupCaller(address common.Address, caller bind.ContractCaller) (*IverifierrollupCaller, error) {
	contract, err := bindIverifierrollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IverifierrollupCaller{contract: contract}, nil
}

// NewIverifierrollupTransactor creates a new write-only instance of Iverifierrollup, bound to a specific deployed contract.
func NewIverifierrollupTransactor(address common.Address, transactor bind.ContractTransactor) (*IverifierrollupTransactor, error) {
	contract, err := bindIverifierrollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IverifierrollupTransactor{contract: contract}, nil
}

// NewIverifierrollupFilterer creates a new log filterer instance of Iverifierrollup, bound to a specific deployed contract.
func NewIverifierrollupFilterer(address common.Address, filterer bind.ContractFilterer) (*IverifierrollupFilterer, error) {
	contract, err := bindIverifierrollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IverifierrollupFilterer{contract: contract}, nil
}

// bindIverifierrollup binds a generic wrapper to an already deployed contract.
func bindIverifierrollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IverifierrollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iverifierrollup *IverifierrollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iverifierrollup.Contract.IverifierrollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iverifierrollup *IverifierrollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iverifierrollup.Contract.IverifierrollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iverifierrollup *IverifierrollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iverifierrollup.Contract.IverifierrollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iverifierrollup *IverifierrollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iverifierrollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iverifierrollup *IverifierrollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iverifierrollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iverifierrollup *IverifierrollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iverifierrollup.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) view returns(bool)
func (_Iverifierrollup *IverifierrollupCaller) VerifyProof(opts *bind.CallOpts, proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Iverifierrollup.contract.Call(opts, &out, "verifyProof", proof, pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) view returns(bool)
func (_Iverifierrollup *IverifierrollupSession) VerifyProof(proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	return _Iverifierrollup.Contract.VerifyProof(&_Iverifierrollup.CallOpts, proof, pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) view returns(bool)
func (_Iverifierrollup *IverifierrollupCallerSession) VerifyProof(proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	return _Iverifierrollup.Contract.VerifyProof(&_Iverifierrollup.CallOpts, proof, pubSignals)
}

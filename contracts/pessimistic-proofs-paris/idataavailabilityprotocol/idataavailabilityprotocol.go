// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idataavailabilityprotocol

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

// IdataavailabilityprotocolMetaData contains all meta data concerning the Idataavailabilityprotocol contract.
var IdataavailabilityprotocolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getProcotolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"verifyMessage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IdataavailabilityprotocolABI is the input ABI used to generate the binding from.
// Deprecated: Use IdataavailabilityprotocolMetaData.ABI instead.
var IdataavailabilityprotocolABI = IdataavailabilityprotocolMetaData.ABI

// Idataavailabilityprotocol is an auto generated Go binding around an Ethereum contract.
type Idataavailabilityprotocol struct {
	IdataavailabilityprotocolCaller     // Read-only binding to the contract
	IdataavailabilityprotocolTransactor // Write-only binding to the contract
	IdataavailabilityprotocolFilterer   // Log filterer for contract events
}

// IdataavailabilityprotocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdataavailabilityprotocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdataavailabilityprotocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdataavailabilityprotocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdataavailabilityprotocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdataavailabilityprotocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdataavailabilityprotocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdataavailabilityprotocolSession struct {
	Contract     *Idataavailabilityprotocol // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdataavailabilityprotocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdataavailabilityprotocolCallerSession struct {
	Contract *IdataavailabilityprotocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IdataavailabilityprotocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdataavailabilityprotocolTransactorSession struct {
	Contract     *IdataavailabilityprotocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IdataavailabilityprotocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdataavailabilityprotocolRaw struct {
	Contract *Idataavailabilityprotocol // Generic contract binding to access the raw methods on
}

// IdataavailabilityprotocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdataavailabilityprotocolCallerRaw struct {
	Contract *IdataavailabilityprotocolCaller // Generic read-only contract binding to access the raw methods on
}

// IdataavailabilityprotocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdataavailabilityprotocolTransactorRaw struct {
	Contract *IdataavailabilityprotocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdataavailabilityprotocol creates a new instance of Idataavailabilityprotocol, bound to a specific deployed contract.
func NewIdataavailabilityprotocol(address common.Address, backend bind.ContractBackend) (*Idataavailabilityprotocol, error) {
	contract, err := bindIdataavailabilityprotocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Idataavailabilityprotocol{IdataavailabilityprotocolCaller: IdataavailabilityprotocolCaller{contract: contract}, IdataavailabilityprotocolTransactor: IdataavailabilityprotocolTransactor{contract: contract}, IdataavailabilityprotocolFilterer: IdataavailabilityprotocolFilterer{contract: contract}}, nil
}

// NewIdataavailabilityprotocolCaller creates a new read-only instance of Idataavailabilityprotocol, bound to a specific deployed contract.
func NewIdataavailabilityprotocolCaller(address common.Address, caller bind.ContractCaller) (*IdataavailabilityprotocolCaller, error) {
	contract, err := bindIdataavailabilityprotocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdataavailabilityprotocolCaller{contract: contract}, nil
}

// NewIdataavailabilityprotocolTransactor creates a new write-only instance of Idataavailabilityprotocol, bound to a specific deployed contract.
func NewIdataavailabilityprotocolTransactor(address common.Address, transactor bind.ContractTransactor) (*IdataavailabilityprotocolTransactor, error) {
	contract, err := bindIdataavailabilityprotocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdataavailabilityprotocolTransactor{contract: contract}, nil
}

// NewIdataavailabilityprotocolFilterer creates a new log filterer instance of Idataavailabilityprotocol, bound to a specific deployed contract.
func NewIdataavailabilityprotocolFilterer(address common.Address, filterer bind.ContractFilterer) (*IdataavailabilityprotocolFilterer, error) {
	contract, err := bindIdataavailabilityprotocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdataavailabilityprotocolFilterer{contract: contract}, nil
}

// bindIdataavailabilityprotocol binds a generic wrapper to an already deployed contract.
func bindIdataavailabilityprotocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdataavailabilityprotocolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idataavailabilityprotocol.Contract.IdataavailabilityprotocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idataavailabilityprotocol.Contract.IdataavailabilityprotocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idataavailabilityprotocol.Contract.IdataavailabilityprotocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Idataavailabilityprotocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Idataavailabilityprotocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Idataavailabilityprotocol *IdataavailabilityprotocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Idataavailabilityprotocol.Contract.contract.Transact(opts, method, params...)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() view returns(string)
func (_Idataavailabilityprotocol *IdataavailabilityprotocolCaller) GetProcotolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Idataavailabilityprotocol.contract.Call(opts, &out, "getProcotolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() view returns(string)
func (_Idataavailabilityprotocol *IdataavailabilityprotocolSession) GetProcotolName() (string, error) {
	return _Idataavailabilityprotocol.Contract.GetProcotolName(&_Idataavailabilityprotocol.CallOpts)
}

// GetProcotolName is a free data retrieval call binding the contract method 0xe4f17120.
//
// Solidity: function getProcotolName() view returns(string)
func (_Idataavailabilityprotocol *IdataavailabilityprotocolCallerSession) GetProcotolName() (string, error) {
	return _Idataavailabilityprotocol.Contract.GetProcotolName(&_Idataavailabilityprotocol.CallOpts)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 hash, bytes dataAvailabilityMessage) view returns()
func (_Idataavailabilityprotocol *IdataavailabilityprotocolCaller) VerifyMessage(opts *bind.CallOpts, hash [32]byte, dataAvailabilityMessage []byte) error {
	var out []interface{}
	err := _Idataavailabilityprotocol.contract.Call(opts, &out, "verifyMessage", hash, dataAvailabilityMessage)

	if err != nil {
		return err
	}

	return err

}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 hash, bytes dataAvailabilityMessage) view returns()
func (_Idataavailabilityprotocol *IdataavailabilityprotocolSession) VerifyMessage(hash [32]byte, dataAvailabilityMessage []byte) error {
	return _Idataavailabilityprotocol.Contract.VerifyMessage(&_Idataavailabilityprotocol.CallOpts, hash, dataAvailabilityMessage)
}

// VerifyMessage is a free data retrieval call binding the contract method 0x3b51be4b.
//
// Solidity: function verifyMessage(bytes32 hash, bytes dataAvailabilityMessage) view returns()
func (_Idataavailabilityprotocol *IdataavailabilityprotocolCallerSession) VerifyMessage(hash [32]byte, dataAvailabilityMessage []byte) error {
	return _Idataavailabilityprotocol.Contract.VerifyMessage(&_Idataavailabilityprotocol.CallOpts, hash, dataAvailabilityMessage)
}

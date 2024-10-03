// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isp1verifierwithhash

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

// Isp1verifierwithhashMetaData contains all meta data concerning the Isp1verifierwithhash contract.
var Isp1verifierwithhashMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"VERIFIER_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Isp1verifierwithhashABI is the input ABI used to generate the binding from.
// Deprecated: Use Isp1verifierwithhashMetaData.ABI instead.
var Isp1verifierwithhashABI = Isp1verifierwithhashMetaData.ABI

// Isp1verifierwithhash is an auto generated Go binding around an Ethereum contract.
type Isp1verifierwithhash struct {
	Isp1verifierwithhashCaller     // Read-only binding to the contract
	Isp1verifierwithhashTransactor // Write-only binding to the contract
	Isp1verifierwithhashFilterer   // Log filterer for contract events
}

// Isp1verifierwithhashCaller is an auto generated read-only Go binding around an Ethereum contract.
type Isp1verifierwithhashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierwithhashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Isp1verifierwithhashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierwithhashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Isp1verifierwithhashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierwithhashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Isp1verifierwithhashSession struct {
	Contract     *Isp1verifierwithhash // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Isp1verifierwithhashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Isp1verifierwithhashCallerSession struct {
	Contract *Isp1verifierwithhashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// Isp1verifierwithhashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Isp1verifierwithhashTransactorSession struct {
	Contract     *Isp1verifierwithhashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// Isp1verifierwithhashRaw is an auto generated low-level Go binding around an Ethereum contract.
type Isp1verifierwithhashRaw struct {
	Contract *Isp1verifierwithhash // Generic contract binding to access the raw methods on
}

// Isp1verifierwithhashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Isp1verifierwithhashCallerRaw struct {
	Contract *Isp1verifierwithhashCaller // Generic read-only contract binding to access the raw methods on
}

// Isp1verifierwithhashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Isp1verifierwithhashTransactorRaw struct {
	Contract *Isp1verifierwithhashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIsp1verifierwithhash creates a new instance of Isp1verifierwithhash, bound to a specific deployed contract.
func NewIsp1verifierwithhash(address common.Address, backend bind.ContractBackend) (*Isp1verifierwithhash, error) {
	contract, err := bindIsp1verifierwithhash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierwithhash{Isp1verifierwithhashCaller: Isp1verifierwithhashCaller{contract: contract}, Isp1verifierwithhashTransactor: Isp1verifierwithhashTransactor{contract: contract}, Isp1verifierwithhashFilterer: Isp1verifierwithhashFilterer{contract: contract}}, nil
}

// NewIsp1verifierwithhashCaller creates a new read-only instance of Isp1verifierwithhash, bound to a specific deployed contract.
func NewIsp1verifierwithhashCaller(address common.Address, caller bind.ContractCaller) (*Isp1verifierwithhashCaller, error) {
	contract, err := bindIsp1verifierwithhash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierwithhashCaller{contract: contract}, nil
}

// NewIsp1verifierwithhashTransactor creates a new write-only instance of Isp1verifierwithhash, bound to a specific deployed contract.
func NewIsp1verifierwithhashTransactor(address common.Address, transactor bind.ContractTransactor) (*Isp1verifierwithhashTransactor, error) {
	contract, err := bindIsp1verifierwithhash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierwithhashTransactor{contract: contract}, nil
}

// NewIsp1verifierwithhashFilterer creates a new log filterer instance of Isp1verifierwithhash, bound to a specific deployed contract.
func NewIsp1verifierwithhashFilterer(address common.Address, filterer bind.ContractFilterer) (*Isp1verifierwithhashFilterer, error) {
	contract, err := bindIsp1verifierwithhash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierwithhashFilterer{contract: contract}, nil
}

// bindIsp1verifierwithhash binds a generic wrapper to an already deployed contract.
func bindIsp1verifierwithhash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Isp1verifierwithhashMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Isp1verifierwithhash *Isp1verifierwithhashRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Isp1verifierwithhash.Contract.Isp1verifierwithhashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Isp1verifierwithhash *Isp1verifierwithhashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Isp1verifierwithhash.Contract.Isp1verifierwithhashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Isp1verifierwithhash *Isp1verifierwithhashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Isp1verifierwithhash.Contract.Isp1verifierwithhashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Isp1verifierwithhash *Isp1verifierwithhashCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Isp1verifierwithhash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Isp1verifierwithhash *Isp1verifierwithhashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Isp1verifierwithhash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Isp1verifierwithhash *Isp1verifierwithhashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Isp1verifierwithhash.Contract.contract.Transact(opts, method, params...)
}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_Isp1verifierwithhash *Isp1verifierwithhashCaller) VERIFIERHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Isp1verifierwithhash.contract.Call(opts, &out, "VERIFIER_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_Isp1verifierwithhash *Isp1verifierwithhashSession) VERIFIERHASH() ([32]byte, error) {
	return _Isp1verifierwithhash.Contract.VERIFIERHASH(&_Isp1verifierwithhash.CallOpts)
}

// VERIFIERHASH is a free data retrieval call binding the contract method 0x2a510436.
//
// Solidity: function VERIFIER_HASH() pure returns(bytes32)
func (_Isp1verifierwithhash *Isp1verifierwithhashCallerSession) VERIFIERHASH() ([32]byte, error) {
	return _Isp1verifierwithhash.Contract.VERIFIERHASH(&_Isp1verifierwithhash.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifierwithhash *Isp1verifierwithhashCaller) VerifyProof(opts *bind.CallOpts, programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Isp1verifierwithhash.contract.Call(opts, &out, "verifyProof", programVKey, publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifierwithhash *Isp1verifierwithhashSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Isp1verifierwithhash.Contract.VerifyProof(&_Isp1verifierwithhash.CallOpts, programVKey, publicValues, proofBytes)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifierwithhash *Isp1verifierwithhashCallerSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Isp1verifierwithhash.Contract.VerifyProof(&_Isp1verifierwithhash.CallOpts, programVKey, publicValues, proofBytes)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isp1verifier

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

// Isp1verifierMetaData contains all meta data concerning the Isp1verifier contract.
var Isp1verifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Isp1verifierABI is the input ABI used to generate the binding from.
// Deprecated: Use Isp1verifierMetaData.ABI instead.
var Isp1verifierABI = Isp1verifierMetaData.ABI

// Isp1verifier is an auto generated Go binding around an Ethereum contract.
type Isp1verifier struct {
	Isp1verifierCaller     // Read-only binding to the contract
	Isp1verifierTransactor // Write-only binding to the contract
	Isp1verifierFilterer   // Log filterer for contract events
}

// Isp1verifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type Isp1verifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Isp1verifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Isp1verifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Isp1verifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Isp1verifierSession struct {
	Contract     *Isp1verifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Isp1verifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Isp1verifierCallerSession struct {
	Contract *Isp1verifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Isp1verifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Isp1verifierTransactorSession struct {
	Contract     *Isp1verifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Isp1verifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type Isp1verifierRaw struct {
	Contract *Isp1verifier // Generic contract binding to access the raw methods on
}

// Isp1verifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Isp1verifierCallerRaw struct {
	Contract *Isp1verifierCaller // Generic read-only contract binding to access the raw methods on
}

// Isp1verifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Isp1verifierTransactorRaw struct {
	Contract *Isp1verifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIsp1verifier creates a new instance of Isp1verifier, bound to a specific deployed contract.
func NewIsp1verifier(address common.Address, backend bind.ContractBackend) (*Isp1verifier, error) {
	contract, err := bindIsp1verifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Isp1verifier{Isp1verifierCaller: Isp1verifierCaller{contract: contract}, Isp1verifierTransactor: Isp1verifierTransactor{contract: contract}, Isp1verifierFilterer: Isp1verifierFilterer{contract: contract}}, nil
}

// NewIsp1verifierCaller creates a new read-only instance of Isp1verifier, bound to a specific deployed contract.
func NewIsp1verifierCaller(address common.Address, caller bind.ContractCaller) (*Isp1verifierCaller, error) {
	contract, err := bindIsp1verifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierCaller{contract: contract}, nil
}

// NewIsp1verifierTransactor creates a new write-only instance of Isp1verifier, bound to a specific deployed contract.
func NewIsp1verifierTransactor(address common.Address, transactor bind.ContractTransactor) (*Isp1verifierTransactor, error) {
	contract, err := bindIsp1verifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierTransactor{contract: contract}, nil
}

// NewIsp1verifierFilterer creates a new log filterer instance of Isp1verifier, bound to a specific deployed contract.
func NewIsp1verifierFilterer(address common.Address, filterer bind.ContractFilterer) (*Isp1verifierFilterer, error) {
	contract, err := bindIsp1verifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Isp1verifierFilterer{contract: contract}, nil
}

// bindIsp1verifier binds a generic wrapper to an already deployed contract.
func bindIsp1verifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Isp1verifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Isp1verifier *Isp1verifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Isp1verifier.Contract.Isp1verifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Isp1verifier *Isp1verifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Isp1verifier.Contract.Isp1verifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Isp1verifier *Isp1verifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Isp1verifier.Contract.Isp1verifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Isp1verifier *Isp1verifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Isp1verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Isp1verifier *Isp1verifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Isp1verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Isp1verifier *Isp1verifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Isp1verifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifier *Isp1verifierCaller) VerifyProof(opts *bind.CallOpts, programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Isp1verifier.contract.Call(opts, &out, "verifyProof", programVKey, publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifier *Isp1verifierSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Isp1verifier.Contract.VerifyProof(&_Isp1verifier.CallOpts, programVKey, publicValues, proofBytes)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) view returns()
func (_Isp1verifier *Isp1verifierCallerSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Isp1verifier.Contract.VerifyProof(&_Isp1verifier.CallOpts, programVKey, publicValues, proofBytes)
}

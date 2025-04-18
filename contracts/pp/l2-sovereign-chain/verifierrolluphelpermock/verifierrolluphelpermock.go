// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifierrolluphelpermock

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

// VerifierrolluphelpermockMetaData contains all meta data concerning the Verifierrolluphelpermock contract.
var VerifierrolluphelpermockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofBytes\",\"type\":\"bytes\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"},{\"internalType\":\"uint256[1]\",\"name\":\"pubSignals\",\"type\":\"uint256[1]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506102308061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c806341493c60146100385780639121da8a1461004f575b5f5ffd5b61004d6100463660046100be565b5050505050565b005b61006561005d366004610164565b600192915050565b604051901515815260200160405180910390f35b5f5f83601f840112610089575f5ffd5b50813567ffffffffffffffff8111156100a0575f5ffd5b6020830191508360208285010111156100b7575f5ffd5b9250929050565b5f5f5f5f5f606086880312156100d2575f5ffd5b85359450602086013567ffffffffffffffff8111156100ef575f5ffd5b6100fb88828901610079565b909550935050604086013567ffffffffffffffff81111561011a575f5ffd5b61012688828901610079565b969995985093965092949392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f6103208385031215610176575f5ffd5b610300830184811115610187575f5ffd5b8392508461031f850112610199575f5ffd5b6040516020810181811067ffffffffffffffff821117156101bc576101bc610137565b604052806103208601878111156101d1575f5ffd5b5b808410156101ed5783358252602093840193909101906101d2565b509396909550935050505056fea2646970667358221220d07f403c1747b906055846fb28b5169440a29119e2178f63b787c953f63c12b764736f6c634300081c0033",
}

// VerifierrolluphelpermockABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierrolluphelpermockMetaData.ABI instead.
var VerifierrolluphelpermockABI = VerifierrolluphelpermockMetaData.ABI

// VerifierrolluphelpermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierrolluphelpermockMetaData.Bin instead.
var VerifierrolluphelpermockBin = VerifierrolluphelpermockMetaData.Bin

// DeployVerifierrolluphelpermock deploys a new Ethereum contract, binding an instance of Verifierrolluphelpermock to it.
func DeployVerifierrolluphelpermock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifierrolluphelpermock, error) {
	parsed, err := VerifierrolluphelpermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierrolluphelpermockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifierrolluphelpermock{VerifierrolluphelpermockCaller: VerifierrolluphelpermockCaller{contract: contract}, VerifierrolluphelpermockTransactor: VerifierrolluphelpermockTransactor{contract: contract}, VerifierrolluphelpermockFilterer: VerifierrolluphelpermockFilterer{contract: contract}}, nil
}

// Verifierrolluphelpermock is an auto generated Go binding around an Ethereum contract.
type Verifierrolluphelpermock struct {
	VerifierrolluphelpermockCaller     // Read-only binding to the contract
	VerifierrolluphelpermockTransactor // Write-only binding to the contract
	VerifierrolluphelpermockFilterer   // Log filterer for contract events
}

// VerifierrolluphelpermockCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierrolluphelpermockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierrolluphelpermockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierrolluphelpermockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierrolluphelpermockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierrolluphelpermockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierrolluphelpermockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierrolluphelpermockSession struct {
	Contract     *Verifierrolluphelpermock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VerifierrolluphelpermockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierrolluphelpermockCallerSession struct {
	Contract *VerifierrolluphelpermockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// VerifierrolluphelpermockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierrolluphelpermockTransactorSession struct {
	Contract     *VerifierrolluphelpermockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// VerifierrolluphelpermockRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierrolluphelpermockRaw struct {
	Contract *Verifierrolluphelpermock // Generic contract binding to access the raw methods on
}

// VerifierrolluphelpermockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierrolluphelpermockCallerRaw struct {
	Contract *VerifierrolluphelpermockCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierrolluphelpermockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierrolluphelpermockTransactorRaw struct {
	Contract *VerifierrolluphelpermockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifierrolluphelpermock creates a new instance of Verifierrolluphelpermock, bound to a specific deployed contract.
func NewVerifierrolluphelpermock(address common.Address, backend bind.ContractBackend) (*Verifierrolluphelpermock, error) {
	contract, err := bindVerifierrolluphelpermock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifierrolluphelpermock{VerifierrolluphelpermockCaller: VerifierrolluphelpermockCaller{contract: contract}, VerifierrolluphelpermockTransactor: VerifierrolluphelpermockTransactor{contract: contract}, VerifierrolluphelpermockFilterer: VerifierrolluphelpermockFilterer{contract: contract}}, nil
}

// NewVerifierrolluphelpermockCaller creates a new read-only instance of Verifierrolluphelpermock, bound to a specific deployed contract.
func NewVerifierrolluphelpermockCaller(address common.Address, caller bind.ContractCaller) (*VerifierrolluphelpermockCaller, error) {
	contract, err := bindVerifierrolluphelpermock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierrolluphelpermockCaller{contract: contract}, nil
}

// NewVerifierrolluphelpermockTransactor creates a new write-only instance of Verifierrolluphelpermock, bound to a specific deployed contract.
func NewVerifierrolluphelpermockTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierrolluphelpermockTransactor, error) {
	contract, err := bindVerifierrolluphelpermock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierrolluphelpermockTransactor{contract: contract}, nil
}

// NewVerifierrolluphelpermockFilterer creates a new log filterer instance of Verifierrolluphelpermock, bound to a specific deployed contract.
func NewVerifierrolluphelpermockFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierrolluphelpermockFilterer, error) {
	contract, err := bindVerifierrolluphelpermock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierrolluphelpermockFilterer{contract: contract}, nil
}

// bindVerifierrolluphelpermock binds a generic wrapper to an already deployed contract.
func bindVerifierrolluphelpermock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierrolluphelpermockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierrolluphelpermock.Contract.VerifierrolluphelpermockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierrolluphelpermock.Contract.VerifierrolluphelpermockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierrolluphelpermock.Contract.VerifierrolluphelpermockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifierrolluphelpermock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifierrolluphelpermock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifierrolluphelpermock *VerifierrolluphelpermockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifierrolluphelpermock.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) pure returns()
func (_Verifierrolluphelpermock *VerifierrolluphelpermockCaller) VerifyProof(opts *bind.CallOpts, programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	var out []interface{}
	err := _Verifierrolluphelpermock.contract.Call(opts, &out, "verifyProof", programVKey, publicValues, proofBytes)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) pure returns()
func (_Verifierrolluphelpermock *VerifierrolluphelpermockSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Verifierrolluphelpermock.Contract.VerifyProof(&_Verifierrolluphelpermock.CallOpts, programVKey, publicValues, proofBytes)
}

// VerifyProof is a free data retrieval call binding the contract method 0x41493c60.
//
// Solidity: function verifyProof(bytes32 programVKey, bytes publicValues, bytes proofBytes) pure returns()
func (_Verifierrolluphelpermock *VerifierrolluphelpermockCallerSession) VerifyProof(programVKey [32]byte, publicValues []byte, proofBytes []byte) error {
	return _Verifierrolluphelpermock.Contract.VerifyProof(&_Verifierrolluphelpermock.CallOpts, programVKey, publicValues, proofBytes)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) pure returns(bool)
func (_Verifierrolluphelpermock *VerifierrolluphelpermockCaller) VerifyProof0(opts *bind.CallOpts, proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	var out []interface{}
	err := _Verifierrolluphelpermock.contract.Call(opts, &out, "verifyProof0", proof, pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof0 is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) pure returns(bool)
func (_Verifierrolluphelpermock *VerifierrolluphelpermockSession) VerifyProof0(proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	return _Verifierrolluphelpermock.Contract.VerifyProof0(&_Verifierrolluphelpermock.CallOpts, proof, pubSignals)
}

// VerifyProof0 is a free data retrieval call binding the contract method 0x9121da8a.
//
// Solidity: function verifyProof(bytes32[24] proof, uint256[1] pubSignals) pure returns(bool)
func (_Verifierrolluphelpermock *VerifierrolluphelpermockCallerSession) VerifyProof0(proof [24][32]byte, pubSignals [1]*big.Int) (bool, error) {
	return _Verifierrolluphelpermock.Contract.VerifyProof0(&_Verifierrolluphelpermock.CallOpts, proof, pubSignals)
}

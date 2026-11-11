// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositcontractbase

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

// DepositcontractbaseMetaData contains all meta data concerning the Depositcontractbase contract.
var DepositcontractbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506101298061001c5f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c80632dfdf0b51460345780635ca1e16514604e575b5f5ffd5b603c60205481565b60405190815260200160405180910390f35b603c6020545f90819081805b602081101560bd578083901c60011660010360975760915f8260208110608057608060c6565b0154855f9182526020526040902090565b935060a7565b5f84815260208390526040902093505b5f8281526020839052604090209150600101605a565b50919392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122034305dfe5b2d3603356536debdcedcb1c6b14af1c58db4f4127e90f6b101521264736f6c634300081c0033",
}

// DepositcontractbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositcontractbaseMetaData.ABI instead.
var DepositcontractbaseABI = DepositcontractbaseMetaData.ABI

// DepositcontractbaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositcontractbaseMetaData.Bin instead.
var DepositcontractbaseBin = DepositcontractbaseMetaData.Bin

// DeployDepositcontractbase deploys a new Ethereum contract, binding an instance of Depositcontractbase to it.
func DeployDepositcontractbase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Depositcontractbase, error) {
	parsed, err := DepositcontractbaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositcontractbaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Depositcontractbase{DepositcontractbaseCaller: DepositcontractbaseCaller{contract: contract}, DepositcontractbaseTransactor: DepositcontractbaseTransactor{contract: contract}, DepositcontractbaseFilterer: DepositcontractbaseFilterer{contract: contract}}, nil
}

// Depositcontractbase is an auto generated Go binding around an Ethereum contract.
type Depositcontractbase struct {
	DepositcontractbaseCaller     // Read-only binding to the contract
	DepositcontractbaseTransactor // Write-only binding to the contract
	DepositcontractbaseFilterer   // Log filterer for contract events
}

// DepositcontractbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositcontractbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositcontractbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositcontractbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositcontractbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositcontractbaseSession struct {
	Contract     *Depositcontractbase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DepositcontractbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositcontractbaseCallerSession struct {
	Contract *DepositcontractbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DepositcontractbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositcontractbaseTransactorSession struct {
	Contract     *DepositcontractbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DepositcontractbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositcontractbaseRaw struct {
	Contract *Depositcontractbase // Generic contract binding to access the raw methods on
}

// DepositcontractbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositcontractbaseCallerRaw struct {
	Contract *DepositcontractbaseCaller // Generic read-only contract binding to access the raw methods on
}

// DepositcontractbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositcontractbaseTransactorRaw struct {
	Contract *DepositcontractbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositcontractbase creates a new instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbase(address common.Address, backend bind.ContractBackend) (*Depositcontractbase, error) {
	contract, err := bindDepositcontractbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositcontractbase{DepositcontractbaseCaller: DepositcontractbaseCaller{contract: contract}, DepositcontractbaseTransactor: DepositcontractbaseTransactor{contract: contract}, DepositcontractbaseFilterer: DepositcontractbaseFilterer{contract: contract}}, nil
}

// NewDepositcontractbaseCaller creates a new read-only instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseCaller(address common.Address, caller bind.ContractCaller) (*DepositcontractbaseCaller, error) {
	contract, err := bindDepositcontractbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseCaller{contract: contract}, nil
}

// NewDepositcontractbaseTransactor creates a new write-only instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositcontractbaseTransactor, error) {
	contract, err := bindDepositcontractbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseTransactor{contract: contract}, nil
}

// NewDepositcontractbaseFilterer creates a new log filterer instance of Depositcontractbase, bound to a specific deployed contract.
func NewDepositcontractbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositcontractbaseFilterer, error) {
	contract, err := bindDepositcontractbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositcontractbaseFilterer{contract: contract}, nil
}

// bindDepositcontractbase binds a generic wrapper to an already deployed contract.
func bindDepositcontractbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositcontractbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbase *DepositcontractbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbase.Contract.DepositcontractbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbase *DepositcontractbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.DepositcontractbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbase *DepositcontractbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.DepositcontractbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositcontractbase *DepositcontractbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositcontractbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositcontractbase *DepositcontractbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositcontractbase *DepositcontractbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositcontractbase.Contract.contract.Transact(opts, method, params...)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbase.Contract.DepositCount(&_Depositcontractbase.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Depositcontractbase *DepositcontractbaseCallerSession) DepositCount() (*big.Int, error) {
	return _Depositcontractbase.Contract.DepositCount(&_Depositcontractbase.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Depositcontractbase.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbase.Contract.GetRoot(&_Depositcontractbase.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Depositcontractbase *DepositcontractbaseCallerSession) GetRoot() ([32]byte, error) {
	return _Depositcontractbase.Contract.GetRoot(&_Depositcontractbase.CallOpts)
}

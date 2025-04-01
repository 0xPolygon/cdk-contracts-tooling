// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package legacyzkevmstatevariables

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

// LegacyzkevmstatevariablesMetaData contains all meta data concerning the Legacyzkevmstatevariables contract.
var LegacyzkevmstatevariablesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600e575f80fd5b50603e80601a5f395ff3fe60806040525f80fdfea26469706673582212207d1cda64b0042e25c8af0410036be947607de834b6aa870346d43a504220558264736f6c63430008140033",
}

// LegacyzkevmstatevariablesABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyzkevmstatevariablesMetaData.ABI instead.
var LegacyzkevmstatevariablesABI = LegacyzkevmstatevariablesMetaData.ABI

// LegacyzkevmstatevariablesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyzkevmstatevariablesMetaData.Bin instead.
var LegacyzkevmstatevariablesBin = LegacyzkevmstatevariablesMetaData.Bin

// DeployLegacyzkevmstatevariables deploys a new Ethereum contract, binding an instance of Legacyzkevmstatevariables to it.
func DeployLegacyzkevmstatevariables(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Legacyzkevmstatevariables, error) {
	parsed, err := LegacyzkevmstatevariablesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyzkevmstatevariablesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Legacyzkevmstatevariables{LegacyzkevmstatevariablesCaller: LegacyzkevmstatevariablesCaller{contract: contract}, LegacyzkevmstatevariablesTransactor: LegacyzkevmstatevariablesTransactor{contract: contract}, LegacyzkevmstatevariablesFilterer: LegacyzkevmstatevariablesFilterer{contract: contract}}, nil
}

// Legacyzkevmstatevariables is an auto generated Go binding around an Ethereum contract.
type Legacyzkevmstatevariables struct {
	LegacyzkevmstatevariablesCaller     // Read-only binding to the contract
	LegacyzkevmstatevariablesTransactor // Write-only binding to the contract
	LegacyzkevmstatevariablesFilterer   // Log filterer for contract events
}

// LegacyzkevmstatevariablesCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyzkevmstatevariablesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyzkevmstatevariablesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyzkevmstatevariablesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyzkevmstatevariablesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyzkevmstatevariablesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyzkevmstatevariablesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyzkevmstatevariablesSession struct {
	Contract     *Legacyzkevmstatevariables // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LegacyzkevmstatevariablesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyzkevmstatevariablesCallerSession struct {
	Contract *LegacyzkevmstatevariablesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// LegacyzkevmstatevariablesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyzkevmstatevariablesTransactorSession struct {
	Contract     *LegacyzkevmstatevariablesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// LegacyzkevmstatevariablesRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyzkevmstatevariablesRaw struct {
	Contract *Legacyzkevmstatevariables // Generic contract binding to access the raw methods on
}

// LegacyzkevmstatevariablesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyzkevmstatevariablesCallerRaw struct {
	Contract *LegacyzkevmstatevariablesCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyzkevmstatevariablesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyzkevmstatevariablesTransactorRaw struct {
	Contract *LegacyzkevmstatevariablesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyzkevmstatevariables creates a new instance of Legacyzkevmstatevariables, bound to a specific deployed contract.
func NewLegacyzkevmstatevariables(address common.Address, backend bind.ContractBackend) (*Legacyzkevmstatevariables, error) {
	contract, err := bindLegacyzkevmstatevariables(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Legacyzkevmstatevariables{LegacyzkevmstatevariablesCaller: LegacyzkevmstatevariablesCaller{contract: contract}, LegacyzkevmstatevariablesTransactor: LegacyzkevmstatevariablesTransactor{contract: contract}, LegacyzkevmstatevariablesFilterer: LegacyzkevmstatevariablesFilterer{contract: contract}}, nil
}

// NewLegacyzkevmstatevariablesCaller creates a new read-only instance of Legacyzkevmstatevariables, bound to a specific deployed contract.
func NewLegacyzkevmstatevariablesCaller(address common.Address, caller bind.ContractCaller) (*LegacyzkevmstatevariablesCaller, error) {
	contract, err := bindLegacyzkevmstatevariables(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyzkevmstatevariablesCaller{contract: contract}, nil
}

// NewLegacyzkevmstatevariablesTransactor creates a new write-only instance of Legacyzkevmstatevariables, bound to a specific deployed contract.
func NewLegacyzkevmstatevariablesTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyzkevmstatevariablesTransactor, error) {
	contract, err := bindLegacyzkevmstatevariables(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyzkevmstatevariablesTransactor{contract: contract}, nil
}

// NewLegacyzkevmstatevariablesFilterer creates a new log filterer instance of Legacyzkevmstatevariables, bound to a specific deployed contract.
func NewLegacyzkevmstatevariablesFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyzkevmstatevariablesFilterer, error) {
	contract, err := bindLegacyzkevmstatevariables(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyzkevmstatevariablesFilterer{contract: contract}, nil
}

// bindLegacyzkevmstatevariables binds a generic wrapper to an already deployed contract.
func bindLegacyzkevmstatevariables(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyzkevmstatevariablesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyzkevmstatevariables.Contract.LegacyzkevmstatevariablesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyzkevmstatevariables.Contract.LegacyzkevmstatevariablesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyzkevmstatevariables.Contract.LegacyzkevmstatevariablesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Legacyzkevmstatevariables.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Legacyzkevmstatevariables.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Legacyzkevmstatevariables *LegacyzkevmstatevariablesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Legacyzkevmstatevariables.Contract.contract.Transact(opts, method, params...)
}

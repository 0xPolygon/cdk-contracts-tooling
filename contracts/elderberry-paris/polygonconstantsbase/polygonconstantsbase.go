// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonconstantsbase

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

// PolygonconstantsbaseMetaData contains all meta data concerning the Polygonconstantsbase contract.
var PolygonconstantsbaseMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212204d9eeb308b3d254686bb1c95d88bcd57329c13370974605492a15473f5975e1d64736f6c63430008140033",
}

// PolygonconstantsbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonconstantsbaseMetaData.ABI instead.
var PolygonconstantsbaseABI = PolygonconstantsbaseMetaData.ABI

// PolygonconstantsbaseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonconstantsbaseMetaData.Bin instead.
var PolygonconstantsbaseBin = PolygonconstantsbaseMetaData.Bin

// DeployPolygonconstantsbase deploys a new Ethereum contract, binding an instance of Polygonconstantsbase to it.
func DeployPolygonconstantsbase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Polygonconstantsbase, error) {
	parsed, err := PolygonconstantsbaseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonconstantsbaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonconstantsbase{PolygonconstantsbaseCaller: PolygonconstantsbaseCaller{contract: contract}, PolygonconstantsbaseTransactor: PolygonconstantsbaseTransactor{contract: contract}, PolygonconstantsbaseFilterer: PolygonconstantsbaseFilterer{contract: contract}}, nil
}

// Polygonconstantsbase is an auto generated Go binding around an Ethereum contract.
type Polygonconstantsbase struct {
	PolygonconstantsbaseCaller     // Read-only binding to the contract
	PolygonconstantsbaseTransactor // Write-only binding to the contract
	PolygonconstantsbaseFilterer   // Log filterer for contract events
}

// PolygonconstantsbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonconstantsbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconstantsbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonconstantsbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconstantsbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonconstantsbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconstantsbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonconstantsbaseSession struct {
	Contract     *Polygonconstantsbase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolygonconstantsbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonconstantsbaseCallerSession struct {
	Contract *PolygonconstantsbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolygonconstantsbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonconstantsbaseTransactorSession struct {
	Contract     *PolygonconstantsbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonconstantsbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonconstantsbaseRaw struct {
	Contract *Polygonconstantsbase // Generic contract binding to access the raw methods on
}

// PolygonconstantsbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonconstantsbaseCallerRaw struct {
	Contract *PolygonconstantsbaseCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonconstantsbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonconstantsbaseTransactorRaw struct {
	Contract *PolygonconstantsbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonconstantsbase creates a new instance of Polygonconstantsbase, bound to a specific deployed contract.
func NewPolygonconstantsbase(address common.Address, backend bind.ContractBackend) (*Polygonconstantsbase, error) {
	contract, err := bindPolygonconstantsbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonconstantsbase{PolygonconstantsbaseCaller: PolygonconstantsbaseCaller{contract: contract}, PolygonconstantsbaseTransactor: PolygonconstantsbaseTransactor{contract: contract}, PolygonconstantsbaseFilterer: PolygonconstantsbaseFilterer{contract: contract}}, nil
}

// NewPolygonconstantsbaseCaller creates a new read-only instance of Polygonconstantsbase, bound to a specific deployed contract.
func NewPolygonconstantsbaseCaller(address common.Address, caller bind.ContractCaller) (*PolygonconstantsbaseCaller, error) {
	contract, err := bindPolygonconstantsbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonconstantsbaseCaller{contract: contract}, nil
}

// NewPolygonconstantsbaseTransactor creates a new write-only instance of Polygonconstantsbase, bound to a specific deployed contract.
func NewPolygonconstantsbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonconstantsbaseTransactor, error) {
	contract, err := bindPolygonconstantsbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonconstantsbaseTransactor{contract: contract}, nil
}

// NewPolygonconstantsbaseFilterer creates a new log filterer instance of Polygonconstantsbase, bound to a specific deployed contract.
func NewPolygonconstantsbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonconstantsbaseFilterer, error) {
	contract, err := bindPolygonconstantsbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonconstantsbaseFilterer{contract: contract}, nil
}

// bindPolygonconstantsbase binds a generic wrapper to an already deployed contract.
func bindPolygonconstantsbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonconstantsbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonconstantsbase *PolygonconstantsbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonconstantsbase.Contract.PolygonconstantsbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonconstantsbase *PolygonconstantsbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonconstantsbase.Contract.PolygonconstantsbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonconstantsbase *PolygonconstantsbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonconstantsbase.Contract.PolygonconstantsbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonconstantsbase *PolygonconstantsbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonconstantsbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonconstantsbase *PolygonconstantsbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonconstantsbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonconstantsbase *PolygonconstantsbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonconstantsbase.Contract.contract.Transact(opts, method, params...)
}

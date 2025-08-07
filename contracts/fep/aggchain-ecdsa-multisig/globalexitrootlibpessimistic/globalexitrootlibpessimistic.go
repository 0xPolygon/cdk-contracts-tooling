// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootlibpessimistic

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

// GlobalexitrootlibpessimisticMetaData contains all meta data concerning the Globalexitrootlibpessimistic contract.
var GlobalexitrootlibpessimisticMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea264697066735822122050cb0c208a8dcef77ae341b6ce024141bd36e9240fb5cd2da04f88f016102d9e64736f6c63430008140033",
}

// GlobalexitrootlibpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalexitrootlibpessimisticMetaData.ABI instead.
var GlobalexitrootlibpessimisticABI = GlobalexitrootlibpessimisticMetaData.ABI

// GlobalexitrootlibpessimisticBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlobalexitrootlibpessimisticMetaData.Bin instead.
var GlobalexitrootlibpessimisticBin = GlobalexitrootlibpessimisticMetaData.Bin

// DeployGlobalexitrootlibpessimistic deploys a new Ethereum contract, binding an instance of Globalexitrootlibpessimistic to it.
func DeployGlobalexitrootlibpessimistic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Globalexitrootlibpessimistic, error) {
	parsed, err := GlobalexitrootlibpessimisticMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlobalexitrootlibpessimisticBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootlibpessimistic{GlobalexitrootlibpessimisticCaller: GlobalexitrootlibpessimisticCaller{contract: contract}, GlobalexitrootlibpessimisticTransactor: GlobalexitrootlibpessimisticTransactor{contract: contract}, GlobalexitrootlibpessimisticFilterer: GlobalexitrootlibpessimisticFilterer{contract: contract}}, nil
}

// Globalexitrootlibpessimistic is an auto generated Go binding around an Ethereum contract.
type Globalexitrootlibpessimistic struct {
	GlobalexitrootlibpessimisticCaller     // Read-only binding to the contract
	GlobalexitrootlibpessimisticTransactor // Write-only binding to the contract
	GlobalexitrootlibpessimisticFilterer   // Log filterer for contract events
}

// GlobalexitrootlibpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalexitrootlibpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalexitrootlibpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalexitrootlibpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalexitrootlibpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalexitrootlibpessimisticSession struct {
	Contract     *Globalexitrootlibpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GlobalexitrootlibpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalexitrootlibpessimisticCallerSession struct {
	Contract *GlobalexitrootlibpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// GlobalexitrootlibpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalexitrootlibpessimisticTransactorSession struct {
	Contract     *GlobalexitrootlibpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// GlobalexitrootlibpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalexitrootlibpessimisticRaw struct {
	Contract *Globalexitrootlibpessimistic // Generic contract binding to access the raw methods on
}

// GlobalexitrootlibpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalexitrootlibpessimisticCallerRaw struct {
	Contract *GlobalexitrootlibpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalexitrootlibpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalexitrootlibpessimisticTransactorRaw struct {
	Contract *GlobalexitrootlibpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootlibpessimistic creates a new instance of Globalexitrootlibpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootlibpessimistic(address common.Address, backend bind.ContractBackend) (*Globalexitrootlibpessimistic, error) {
	contract, err := bindGlobalexitrootlibpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootlibpessimistic{GlobalexitrootlibpessimisticCaller: GlobalexitrootlibpessimisticCaller{contract: contract}, GlobalexitrootlibpessimisticTransactor: GlobalexitrootlibpessimisticTransactor{contract: contract}, GlobalexitrootlibpessimisticFilterer: GlobalexitrootlibpessimisticFilterer{contract: contract}}, nil
}

// NewGlobalexitrootlibpessimisticCaller creates a new read-only instance of Globalexitrootlibpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootlibpessimisticCaller(address common.Address, caller bind.ContractCaller) (*GlobalexitrootlibpessimisticCaller, error) {
	contract, err := bindGlobalexitrootlibpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibpessimisticCaller{contract: contract}, nil
}

// NewGlobalexitrootlibpessimisticTransactor creates a new write-only instance of Globalexitrootlibpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootlibpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalexitrootlibpessimisticTransactor, error) {
	contract, err := bindGlobalexitrootlibpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibpessimisticTransactor{contract: contract}, nil
}

// NewGlobalexitrootlibpessimisticFilterer creates a new log filterer instance of Globalexitrootlibpessimistic, bound to a specific deployed contract.
func NewGlobalexitrootlibpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalexitrootlibpessimisticFilterer, error) {
	contract, err := bindGlobalexitrootlibpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalexitrootlibpessimisticFilterer{contract: contract}, nil
}

// bindGlobalexitrootlibpessimistic binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootlibpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlobalexitrootlibpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootlibpessimistic.Contract.GlobalexitrootlibpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootlibpessimistic.Contract.GlobalexitrootlibpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootlibpessimistic.Contract.GlobalexitrootlibpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootlibpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootlibpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootlibpessimistic *GlobalexitrootlibpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootlibpessimistic.Contract.contract.Transact(opts, method, params...)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package senddata

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

// SenddataMetaData contains all meta data concerning the Senddata contract.
var SenddataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610242806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638823d22414610030575b600080fd5b61004361003e3660046100e2565b610045565b005b8173ffffffffffffffffffffffffffffffffffffffff168160405161006a91906101dd565b6000604051808303816000865af19150503d80600081146100a7576040519150601f19603f3d011682016040523d82523d6000602084013e6100ac565b606091505b5050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156100f557600080fd5b823573ffffffffffffffffffffffffffffffffffffffff8116811461011957600080fd5b9150602083013567ffffffffffffffff8082111561013657600080fd5b818501915085601f83011261014a57600080fd5b81358181111561015c5761015c6100b3565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156101a2576101a26100b3565b816040528281528860208487010111156101bb57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000825160005b818110156101fe57602081860181015185830152016101e4565b50600092019182525091905056fea26469706673582212205b62ce6020e84114c136286355b9596c6757119ca9812138f947c7674205e31764736f6c63430008140033",
}

// SenddataABI is the input ABI used to generate the binding from.
// Deprecated: Use SenddataMetaData.ABI instead.
var SenddataABI = SenddataMetaData.ABI

// SenddataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenddataMetaData.Bin instead.
var SenddataBin = SenddataMetaData.Bin

// DeploySenddata deploys a new Ethereum contract, binding an instance of Senddata to it.
func DeploySenddata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Senddata, error) {
	parsed, err := SenddataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenddataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Senddata{SenddataCaller: SenddataCaller{contract: contract}, SenddataTransactor: SenddataTransactor{contract: contract}, SenddataFilterer: SenddataFilterer{contract: contract}}, nil
}

// Senddata is an auto generated Go binding around an Ethereum contract.
type Senddata struct {
	SenddataCaller     // Read-only binding to the contract
	SenddataTransactor // Write-only binding to the contract
	SenddataFilterer   // Log filterer for contract events
}

// SenddataCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenddataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenddataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenddataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenddataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenddataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenddataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenddataSession struct {
	Contract     *Senddata         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenddataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenddataCallerSession struct {
	Contract *SenddataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SenddataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenddataTransactorSession struct {
	Contract     *SenddataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SenddataRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenddataRaw struct {
	Contract *Senddata // Generic contract binding to access the raw methods on
}

// SenddataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenddataCallerRaw struct {
	Contract *SenddataCaller // Generic read-only contract binding to access the raw methods on
}

// SenddataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenddataTransactorRaw struct {
	Contract *SenddataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenddata creates a new instance of Senddata, bound to a specific deployed contract.
func NewSenddata(address common.Address, backend bind.ContractBackend) (*Senddata, error) {
	contract, err := bindSenddata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Senddata{SenddataCaller: SenddataCaller{contract: contract}, SenddataTransactor: SenddataTransactor{contract: contract}, SenddataFilterer: SenddataFilterer{contract: contract}}, nil
}

// NewSenddataCaller creates a new read-only instance of Senddata, bound to a specific deployed contract.
func NewSenddataCaller(address common.Address, caller bind.ContractCaller) (*SenddataCaller, error) {
	contract, err := bindSenddata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenddataCaller{contract: contract}, nil
}

// NewSenddataTransactor creates a new write-only instance of Senddata, bound to a specific deployed contract.
func NewSenddataTransactor(address common.Address, transactor bind.ContractTransactor) (*SenddataTransactor, error) {
	contract, err := bindSenddata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenddataTransactor{contract: contract}, nil
}

// NewSenddataFilterer creates a new log filterer instance of Senddata, bound to a specific deployed contract.
func NewSenddataFilterer(address common.Address, filterer bind.ContractFilterer) (*SenddataFilterer, error) {
	contract, err := bindSenddata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenddataFilterer{contract: contract}, nil
}

// bindSenddata binds a generic wrapper to an already deployed contract.
func bindSenddata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenddataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Senddata *SenddataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Senddata.Contract.SenddataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Senddata *SenddataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Senddata.Contract.SenddataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Senddata *SenddataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Senddata.Contract.SenddataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Senddata *SenddataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Senddata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Senddata *SenddataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Senddata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Senddata *SenddataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Senddata.Contract.contract.Transact(opts, method, params...)
}

// SendData is a paid mutator transaction binding the contract method 0x8823d224.
//
// Solidity: function sendData(address destination, bytes data) returns()
func (_Senddata *SenddataTransactor) SendData(opts *bind.TransactOpts, destination common.Address, data []byte) (*types.Transaction, error) {
	return _Senddata.contract.Transact(opts, "sendData", destination, data)
}

// SendData is a paid mutator transaction binding the contract method 0x8823d224.
//
// Solidity: function sendData(address destination, bytes data) returns()
func (_Senddata *SenddataSession) SendData(destination common.Address, data []byte) (*types.Transaction, error) {
	return _Senddata.Contract.SendData(&_Senddata.TransactOpts, destination, data)
}

// SendData is a paid mutator transaction binding the contract method 0x8823d224.
//
// Solidity: function sendData(address destination, bytes data) returns()
func (_Senddata *SenddataTransactorSession) SendData(destination common.Address, data []byte) (*types.Transaction, error) {
	return _Senddata.Contract.SendData(&_Senddata.TransactOpts, destination, data)
}

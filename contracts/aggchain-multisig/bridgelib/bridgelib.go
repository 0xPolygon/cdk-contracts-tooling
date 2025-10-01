// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgelib

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

// BridgelibMetaData contains all meta data concerning the Bridgelib contract.
var BridgelibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"safeSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedSpender\",\"type\":\"address\"}],\"name\":\"validateAndProcessPermit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506110eb8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610064575f3560e01c8063be3dcf621161004d578063be3dcf62146100b4578063c00f14ab146100d9578063cf825e55146100ec575f5ffd5b8063737ce34114610068578063a28fa4a314610091575b5f5ffd5b61007b610076366004610bfa565b6100ff565b6040516100889190610c68565b60405180910390f35b6100a461009f366004610c7a565b610213565b6040519015158152602001610088565b6100c76100c2366004610bfa565b6107c8565b60405160ff9091168152602001610088565b61007b6100e7366004610bfa565b6108b7565b61007b6100fa366004610bfa565b6108fc565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f95d89b410000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff8616916101809190610d1d565b5f60405180830381855afa9150503d805f81146101b8576040519150601f19603f3d011682016040523d82523d5f602084013e6101bd565b606091505b509150915081610202576040518060400160405280600981526020017f4e4f5f53594d424f4c000000000000000000000000000000000000000000000081525061020b565b61020b816109ff565b949350505050565b5f806102226004828789610d33565b61022b91610d5a565b90507f2afa5331000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008216016104d6575f5f5f5f5f5f5f8c8c600490809261029093929190610d33565b81019061029d9190610dce565b96509650965096509650965096508a73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610310576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8973ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614610375576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8e73ffffffffffffffffffffffffffffffffffffffff1663d505accf60e01b898989898989896040516024016103fb979695949392919073ffffffffffffffffffffffffffffffffffffffff97881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516104849190610d1d565b5f604051808303815f865af19150503d805f81146104bd576040519150601f19603f3d011682016040523d82523d5f602084013e6104c2565b606091505b50909a506107bf9950505050505050505050565b7f703450f4000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000082160161078d575f5f5f5f5f5f5f5f8d8d600490809261053a93929190610d33565b8101906105479190610e3a565b975097509750975097509750975097508b73ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16146105bc576040517f912ecce700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8a73ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610621576040517f750643af00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8f73ffffffffffffffffffffffffffffffffffffffff16638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016106b198979695949392919073ffffffffffffffffffffffffffffffffffffffff9889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161073a9190610d1d565b5f604051808303815f865af19150503d805f8114610773576040519150601f19603f3d011682016040523d82523d5f602084013e610778565b606091505b50909b506107bf9a5050505050505050505050565b6040517fe282c0ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f313ce5670000000000000000000000000000000000000000000000000000000017905290515f918291829173ffffffffffffffffffffffffffffffffffffffff8616916108489190610d1d565b5f60405180830381855afa9150503d805f8114610880576040519150601f19603f3d011682016040523d82523d5f602084013e610885565b606091505b5091509150818015610898575080516020145b6108a357601261020b565b8080602001905181019061020b9190610ebc565b60606108c2826108fc565b6108cb836100ff565b6108d4846107c8565b6040516020016108e693929190610ed7565b6040516020818303038152906040529050919050565b60408051600481526024810182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06fdde030000000000000000000000000000000000000000000000000000000017905290516060915f91829173ffffffffffffffffffffffffffffffffffffffff86169161097d9190610d1d565b5f60405180830381855afa9150503d805f81146109b5576040519150601f19603f3d011682016040523d82523d5f602084013e6109ba565b606091505b509150915081610202576040518060400160405280600781526020017f4e4f5f4e414d450000000000000000000000000000000000000000000000000081525061020b565b60606040825110610a245781806020019051810190610a1e9190610f3c565b92915050565b8151602003610b8d575f5b602081108015610a765750828181518110610a4c57610a4c61102c565b01602001517fff000000000000000000000000000000000000000000000000000000000000001615155b15610a8d5780610a8581611059565b915050610a2f565b805f03610acf57505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e4700000000000000000000000000006020820152919050565b5f8167ffffffffffffffff811115610ae957610ae9610f0f565b6040519080825280601f01601f191660200182016040528015610b13576020820181803683370190505b5090505f5b82811015610b8557848181518110610b3257610b3261102c565b602001015160f81c60f81b828281518110610b4f57610b4f61102c565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600101610b18565b509392505050565b505060408051808201909152601281527f4e4f545f56414c49445f454e434f44494e470000000000000000000000000000602082015290565b919050565b73ffffffffffffffffffffffffffffffffffffffff81168114610bec575f5ffd5b50565b8035610bc681610bcb565b5f60208284031215610c0a575f5ffd5b8135610c1581610bcb565b9392505050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610c156020830184610c1c565b5f5f5f5f5f60808688031215610c8e575f5ffd5b8535610c9981610bcb565b9450602086013567ffffffffffffffff811115610cb4575f5ffd5b8601601f81018813610cc4575f5ffd5b803567ffffffffffffffff811115610cda575f5ffd5b886020828401011115610ceb575f5ffd5b602091909101945092506040860135610d0381610bcb565b9150610d1160608701610bef565b90509295509295909350565b5f82518060208501845e5f920191825250919050565b5f5f85851115610d41575f5ffd5b83861115610d4d575f5ffd5b5050820193919092039150565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015610db9577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b60ff81168114610bec575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610de4575f5ffd5b8735610def81610bcb565b96506020880135610dff81610bcb565b955060408801359450606088013593506080880135610e1d81610dc0565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610e52575f5ffd5b8835610e5d81610bcb565b97506020890135610e6d81610bcb565b9650604089013595506060890135945060808901358015158114610e8f575f5ffd5b935060a0890135610e9f81610dc0565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610ecc575f5ffd5b8151610c1581610dc0565b606081525f610ee96060830186610c1c565b8281036020840152610efb8186610c1c565b91505060ff83166040830152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f60208284031215610f4c575f5ffd5b815167ffffffffffffffff811115610f62575f5ffd5b8201601f81018413610f72575f5ffd5b805167ffffffffffffffff811115610f8c57610f8c610f0f565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715610ff857610ff8610f0f565b60405281815282820160200186101561100f575f5ffd5b8160208401602083015e5f91810160200191909152949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110ae577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fea264697066735822122092a737e21b14e40d9caa84ea645db2319b10eb5bc4b8bf5c333a3e6631bb48df64736f6c634300081c0033",
}

// BridgelibABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgelibMetaData.ABI instead.
var BridgelibABI = BridgelibMetaData.ABI

// BridgelibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgelibMetaData.Bin instead.
var BridgelibBin = BridgelibMetaData.Bin

// DeployBridgelib deploys a new Ethereum contract, binding an instance of Bridgelib to it.
func DeployBridgelib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgelib, error) {
	parsed, err := BridgelibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgelibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgelib{BridgelibCaller: BridgelibCaller{contract: contract}, BridgelibTransactor: BridgelibTransactor{contract: contract}, BridgelibFilterer: BridgelibFilterer{contract: contract}}, nil
}

// Bridgelib is an auto generated Go binding around an Ethereum contract.
type Bridgelib struct {
	BridgelibCaller     // Read-only binding to the contract
	BridgelibTransactor // Write-only binding to the contract
	BridgelibFilterer   // Log filterer for contract events
}

// BridgelibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgelibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgelibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgelibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgelibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgelibSession struct {
	Contract     *Bridgelib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgelibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgelibCallerSession struct {
	Contract *BridgelibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgelibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgelibTransactorSession struct {
	Contract     *BridgelibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgelibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgelibRaw struct {
	Contract *Bridgelib // Generic contract binding to access the raw methods on
}

// BridgelibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgelibCallerRaw struct {
	Contract *BridgelibCaller // Generic read-only contract binding to access the raw methods on
}

// BridgelibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgelibTransactorRaw struct {
	Contract *BridgelibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgelib creates a new instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelib(address common.Address, backend bind.ContractBackend) (*Bridgelib, error) {
	contract, err := bindBridgelib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgelib{BridgelibCaller: BridgelibCaller{contract: contract}, BridgelibTransactor: BridgelibTransactor{contract: contract}, BridgelibFilterer: BridgelibFilterer{contract: contract}}, nil
}

// NewBridgelibCaller creates a new read-only instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibCaller(address common.Address, caller bind.ContractCaller) (*BridgelibCaller, error) {
	contract, err := bindBridgelib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgelibCaller{contract: contract}, nil
}

// NewBridgelibTransactor creates a new write-only instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgelibTransactor, error) {
	contract, err := bindBridgelib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgelibTransactor{contract: contract}, nil
}

// NewBridgelibFilterer creates a new log filterer instance of Bridgelib, bound to a specific deployed contract.
func NewBridgelibFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgelibFilterer, error) {
	contract, err := bindBridgelib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgelibFilterer{contract: contract}, nil
}

// bindBridgelib binds a generic wrapper to an already deployed contract.
func bindBridgelib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgelibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgelib *BridgelibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgelib.Contract.BridgelibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgelib *BridgelibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgelib.Contract.BridgelibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgelib *BridgelibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgelib.Contract.BridgelibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgelib *BridgelibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgelib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgelib *BridgelibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgelib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgelib *BridgelibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgelib.Contract.contract.Transact(opts, method, params...)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgelib.Contract.GetTokenMetadata(&_Bridgelib.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgelib *BridgelibCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgelib.Contract.GetTokenMetadata(&_Bridgelib.CallOpts, token)
}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibCaller) SafeDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibSession) SafeDecimals(token common.Address) (uint8, error) {
	return _Bridgelib.Contract.SafeDecimals(&_Bridgelib.CallOpts, token)
}

// SafeDecimals is a free data retrieval call binding the contract method 0xbe3dcf62.
//
// Solidity: function safeDecimals(address token) view returns(uint8)
func (_Bridgelib *BridgelibCallerSession) SafeDecimals(token common.Address) (uint8, error) {
	return _Bridgelib.Contract.SafeDecimals(&_Bridgelib.CallOpts, token)
}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibCaller) SafeName(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeName", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibSession) SafeName(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeName(&_Bridgelib.CallOpts, token)
}

// SafeName is a free data retrieval call binding the contract method 0xcf825e55.
//
// Solidity: function safeName(address token) view returns(string)
func (_Bridgelib *BridgelibCallerSession) SafeName(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeName(&_Bridgelib.CallOpts, token)
}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibCaller) SafeSymbol(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _Bridgelib.contract.Call(opts, &out, "safeSymbol", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibSession) SafeSymbol(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeSymbol(&_Bridgelib.CallOpts, token)
}

// SafeSymbol is a free data retrieval call binding the contract method 0x737ce341.
//
// Solidity: function safeSymbol(address token) view returns(string)
func (_Bridgelib *BridgelibCallerSession) SafeSymbol(token common.Address) (string, error) {
	return _Bridgelib.Contract.SafeSymbol(&_Bridgelib.CallOpts, token)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibTransactor) ValidateAndProcessPermit(opts *bind.TransactOpts, token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.contract.Transact(opts, "validateAndProcessPermit", token, permitData, expectedOwner, expectedSpender)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibSession) ValidateAndProcessPermit(token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.Contract.ValidateAndProcessPermit(&_Bridgelib.TransactOpts, token, permitData, expectedOwner, expectedSpender)
}

// ValidateAndProcessPermit is a paid mutator transaction binding the contract method 0xa28fa4a3.
//
// Solidity: function validateAndProcessPermit(address token, bytes permitData, address expectedOwner, address expectedSpender) returns(bool success)
func (_Bridgelib *BridgelibTransactorSession) ValidateAndProcessPermit(token common.Address, permitData []byte, expectedOwner common.Address, expectedSpender common.Address) (*types.Transaction, error) {
	return _Bridgelib.Contract.ValidateAndProcessPermit(&_Bridgelib.TransactOpts, token, permitData, expectedOwner, expectedSpender)
}

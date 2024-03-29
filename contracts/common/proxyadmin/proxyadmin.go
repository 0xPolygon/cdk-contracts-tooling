// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxyadmin

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

// ProxyadminMetaData contains all meta data concerning the Proxyadmin contract.
var ProxyadminMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6108658061007e6000396000f3fe60806040526004361061007b5760003560e01c80639623609d1161004e5780639623609d1461012b57806399a88ec41461013e578063f2fde38b1461015e578063f3b7dead1461017e57600080fd5b8063204e1c7a14610080578063715018a6146100c95780637eff275e146100e05780638da5cb5b14610100575b600080fd5b34801561008c57600080fd5b506100a061009b366004610608565b61019e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100d557600080fd5b506100de610255565b005b3480156100ec57600080fd5b506100de6100fb36600461062c565b610269565b34801561010c57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff166100a0565b6100de610139366004610694565b6102f7565b34801561014a57600080fd5b506100de61015936600461062c565b61038c565b34801561016a57600080fd5b506100de610179366004610608565b6103e8565b34801561018a57600080fd5b506100a0610199366004610608565b6104a4565b60008060008373ffffffffffffffffffffffffffffffffffffffff166040516101ea907f5c60da1b00000000000000000000000000000000000000000000000000000000815260040190565b600060405180830381855afa9150503d8060008114610225576040519150601f19603f3d011682016040523d82523d6000602084013e61022a565b606091505b50915091508161023957600080fd5b8080602001905181019061024d9190610788565b949350505050565b61025d6104f0565b6102676000610571565b565b6102716104f0565b6040517f8f28397000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8281166004830152831690638f283970906024015b600060405180830381600087803b1580156102db57600080fd5b505af11580156102ef573d6000803e3d6000fd5b505050505050565b6102ff6104f0565b6040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690634f1ef28690349061035590869086906004016107a5565b6000604051808303818588803b15801561036e57600080fd5b505af1158015610382573d6000803e3d6000fd5b5050505050505050565b6103946104f0565b6040517f3659cfe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8281166004830152831690633659cfe6906024016102c1565b6103f06104f0565b73ffffffffffffffffffffffffffffffffffffffff8116610498576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6104a181610571565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff166040516101ea907ff851a44000000000000000000000000000000000000000000000000000000000815260040190565b60005473ffffffffffffffffffffffffffffffffffffffff163314610267576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161048f565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff811681146104a157600080fd5b60006020828403121561061a57600080fd5b8135610625816105e6565b9392505050565b6000806040838503121561063f57600080fd5b823561064a816105e6565b9150602083013561065a816105e6565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156106a957600080fd5b83356106b4816105e6565b925060208401356106c4816105e6565b9150604084013567ffffffffffffffff808211156106e157600080fd5b818601915086601f8301126106f557600080fd5b81358181111561070757610707610665565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561074d5761074d610665565b8160405282815289602084870101111561076657600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60006020828403121561079a57600080fd5b8151610625816105e6565b73ffffffffffffffffffffffffffffffffffffffff8316815260006020604081840152835180604085015260005b818110156107ef578581018301518582016060015282016107d3565b5060006060828601015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010192505050939250505056fea2646970667358221220c9867ffac53151bdb1305d8f5e3e883cd83e5270c7ec09cdc24e837b2e65239064736f6c63430008140033",
}

// ProxyadminABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyadminMetaData.ABI instead.
var ProxyadminABI = ProxyadminMetaData.ABI

// ProxyadminBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyadminMetaData.Bin instead.
var ProxyadminBin = ProxyadminMetaData.Bin

// DeployProxyadmin deploys a new Ethereum contract, binding an instance of Proxyadmin to it.
func DeployProxyadmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proxyadmin, error) {
	parsed, err := ProxyadminMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyadminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxyadmin{ProxyadminCaller: ProxyadminCaller{contract: contract}, ProxyadminTransactor: ProxyadminTransactor{contract: contract}, ProxyadminFilterer: ProxyadminFilterer{contract: contract}}, nil
}

// Proxyadmin is an auto generated Go binding around an Ethereum contract.
type Proxyadmin struct {
	ProxyadminCaller     // Read-only binding to the contract
	ProxyadminTransactor // Write-only binding to the contract
	ProxyadminFilterer   // Log filterer for contract events
}

// ProxyadminCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyadminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyadminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyadminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyadminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyadminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyadminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyadminSession struct {
	Contract     *Proxyadmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyadminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyadminCallerSession struct {
	Contract *ProxyadminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProxyadminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyadminTransactorSession struct {
	Contract     *ProxyadminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyadminRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyadminRaw struct {
	Contract *Proxyadmin // Generic contract binding to access the raw methods on
}

// ProxyadminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyadminCallerRaw struct {
	Contract *ProxyadminCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyadminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyadminTransactorRaw struct {
	Contract *ProxyadminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyadmin creates a new instance of Proxyadmin, bound to a specific deployed contract.
func NewProxyadmin(address common.Address, backend bind.ContractBackend) (*Proxyadmin, error) {
	contract, err := bindProxyadmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxyadmin{ProxyadminCaller: ProxyadminCaller{contract: contract}, ProxyadminTransactor: ProxyadminTransactor{contract: contract}, ProxyadminFilterer: ProxyadminFilterer{contract: contract}}, nil
}

// NewProxyadminCaller creates a new read-only instance of Proxyadmin, bound to a specific deployed contract.
func NewProxyadminCaller(address common.Address, caller bind.ContractCaller) (*ProxyadminCaller, error) {
	contract, err := bindProxyadmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyadminCaller{contract: contract}, nil
}

// NewProxyadminTransactor creates a new write-only instance of Proxyadmin, bound to a specific deployed contract.
func NewProxyadminTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyadminTransactor, error) {
	contract, err := bindProxyadmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyadminTransactor{contract: contract}, nil
}

// NewProxyadminFilterer creates a new log filterer instance of Proxyadmin, bound to a specific deployed contract.
func NewProxyadminFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyadminFilterer, error) {
	contract, err := bindProxyadmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyadminFilterer{contract: contract}, nil
}

// bindProxyadmin binds a generic wrapper to an already deployed contract.
func bindProxyadmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyadminMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyadmin *ProxyadminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyadmin.Contract.ProxyadminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyadmin *ProxyadminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyadmin.Contract.ProxyadminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyadmin *ProxyadminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyadmin.Contract.ProxyadminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxyadmin *ProxyadminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxyadmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxyadmin *ProxyadminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyadmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxyadmin *ProxyadminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxyadmin.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Proxyadmin.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _Proxyadmin.Contract.GetProxyAdmin(&_Proxyadmin.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _Proxyadmin.Contract.GetProxyAdmin(&_Proxyadmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Proxyadmin.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _Proxyadmin.Contract.GetProxyImplementation(&_Proxyadmin.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_Proxyadmin *ProxyadminCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _Proxyadmin.Contract.GetProxyImplementation(&_Proxyadmin.CallOpts, proxy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyadmin *ProxyadminCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxyadmin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyadmin *ProxyadminSession) Owner() (common.Address, error) {
	return _Proxyadmin.Contract.Owner(&_Proxyadmin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proxyadmin *ProxyadminCallerSession) Owner() (common.Address, error) {
	return _Proxyadmin.Contract.Owner(&_Proxyadmin.CallOpts)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_Proxyadmin *ProxyadminTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _Proxyadmin.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_Proxyadmin *ProxyadminSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.ChangeProxyAdmin(&_Proxyadmin.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_Proxyadmin *ProxyadminTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.ChangeProxyAdmin(&_Proxyadmin.TransactOpts, proxy, newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proxyadmin *ProxyadminTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxyadmin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proxyadmin *ProxyadminSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proxyadmin.Contract.RenounceOwnership(&_Proxyadmin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proxyadmin *ProxyadminTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proxyadmin.Contract.RenounceOwnership(&_Proxyadmin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proxyadmin *ProxyadminTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proxyadmin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proxyadmin *ProxyadminSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.TransferOwnership(&_Proxyadmin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proxyadmin *ProxyadminTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.TransferOwnership(&_Proxyadmin.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_Proxyadmin *ProxyadminTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Proxyadmin.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_Proxyadmin *ProxyadminSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.Upgrade(&_Proxyadmin.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_Proxyadmin *ProxyadminTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _Proxyadmin.Contract.Upgrade(&_Proxyadmin.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_Proxyadmin *ProxyadminTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _Proxyadmin.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_Proxyadmin *ProxyadminSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _Proxyadmin.Contract.UpgradeAndCall(&_Proxyadmin.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_Proxyadmin *ProxyadminTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _Proxyadmin.Contract.UpgradeAndCall(&_Proxyadmin.TransactOpts, proxy, implementation, data)
}

// ProxyadminOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Proxyadmin contract.
type ProxyadminOwnershipTransferredIterator struct {
	Event *ProxyadminOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyadminOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyadminOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyadminOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyadminOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyadminOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyadminOwnershipTransferred represents a OwnershipTransferred event raised by the Proxyadmin contract.
type ProxyadminOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proxyadmin *ProxyadminFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyadminOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proxyadmin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyadminOwnershipTransferredIterator{contract: _Proxyadmin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proxyadmin *ProxyadminFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyadminOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proxyadmin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyadminOwnershipTransferred)
				if err := _Proxyadmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proxyadmin *ProxyadminFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyadminOwnershipTransferred, error) {
	event := new(ProxyadminOwnershipTransferred)
	if err := _Proxyadmin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenwrappedtransparentproxy

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

// TokenwrappedtransparentproxyMetaData contains all meta data concerning the Tokenwrappedtransparentproxy contract.
var TokenwrappedtransparentproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564",
}

// TokenwrappedtransparentproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenwrappedtransparentproxyMetaData.ABI instead.
var TokenwrappedtransparentproxyABI = TokenwrappedtransparentproxyMetaData.ABI

// TokenwrappedtransparentproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenwrappedtransparentproxyMetaData.Bin instead.
var TokenwrappedtransparentproxyBin = TokenwrappedtransparentproxyMetaData.Bin

// DeployTokenwrappedtransparentproxy deploys a new Ethereum contract, binding an instance of Tokenwrappedtransparentproxy to it.
func DeployTokenwrappedtransparentproxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenwrappedtransparentproxy, error) {
	parsed, err := TokenwrappedtransparentproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenwrappedtransparentproxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenwrappedtransparentproxy{TokenwrappedtransparentproxyCaller: TokenwrappedtransparentproxyCaller{contract: contract}, TokenwrappedtransparentproxyTransactor: TokenwrappedtransparentproxyTransactor{contract: contract}, TokenwrappedtransparentproxyFilterer: TokenwrappedtransparentproxyFilterer{contract: contract}}, nil
}

// Tokenwrappedtransparentproxy is an auto generated Go binding around an Ethereum contract.
type Tokenwrappedtransparentproxy struct {
	TokenwrappedtransparentproxyCaller     // Read-only binding to the contract
	TokenwrappedtransparentproxyTransactor // Write-only binding to the contract
	TokenwrappedtransparentproxyFilterer   // Log filterer for contract events
}

// TokenwrappedtransparentproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenwrappedtransparentproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedtransparentproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenwrappedtransparentproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedtransparentproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenwrappedtransparentproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenwrappedtransparentproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenwrappedtransparentproxySession struct {
	Contract     *Tokenwrappedtransparentproxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenwrappedtransparentproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenwrappedtransparentproxyCallerSession struct {
	Contract *TokenwrappedtransparentproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// TokenwrappedtransparentproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenwrappedtransparentproxyTransactorSession struct {
	Contract     *TokenwrappedtransparentproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// TokenwrappedtransparentproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenwrappedtransparentproxyRaw struct {
	Contract *Tokenwrappedtransparentproxy // Generic contract binding to access the raw methods on
}

// TokenwrappedtransparentproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenwrappedtransparentproxyCallerRaw struct {
	Contract *TokenwrappedtransparentproxyCaller // Generic read-only contract binding to access the raw methods on
}

// TokenwrappedtransparentproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenwrappedtransparentproxyTransactorRaw struct {
	Contract *TokenwrappedtransparentproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenwrappedtransparentproxy creates a new instance of Tokenwrappedtransparentproxy, bound to a specific deployed contract.
func NewTokenwrappedtransparentproxy(address common.Address, backend bind.ContractBackend) (*Tokenwrappedtransparentproxy, error) {
	contract, err := bindTokenwrappedtransparentproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenwrappedtransparentproxy{TokenwrappedtransparentproxyCaller: TokenwrappedtransparentproxyCaller{contract: contract}, TokenwrappedtransparentproxyTransactor: TokenwrappedtransparentproxyTransactor{contract: contract}, TokenwrappedtransparentproxyFilterer: TokenwrappedtransparentproxyFilterer{contract: contract}}, nil
}

// NewTokenwrappedtransparentproxyCaller creates a new read-only instance of Tokenwrappedtransparentproxy, bound to a specific deployed contract.
func NewTokenwrappedtransparentproxyCaller(address common.Address, caller bind.ContractCaller) (*TokenwrappedtransparentproxyCaller, error) {
	contract, err := bindTokenwrappedtransparentproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyCaller{contract: contract}, nil
}

// NewTokenwrappedtransparentproxyTransactor creates a new write-only instance of Tokenwrappedtransparentproxy, bound to a specific deployed contract.
func NewTokenwrappedtransparentproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenwrappedtransparentproxyTransactor, error) {
	contract, err := bindTokenwrappedtransparentproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyTransactor{contract: contract}, nil
}

// NewTokenwrappedtransparentproxyFilterer creates a new log filterer instance of Tokenwrappedtransparentproxy, bound to a specific deployed contract.
func NewTokenwrappedtransparentproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenwrappedtransparentproxyFilterer, error) {
	contract, err := bindTokenwrappedtransparentproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyFilterer{contract: contract}, nil
}

// bindTokenwrappedtransparentproxy binds a generic wrapper to an already deployed contract.
func bindTokenwrappedtransparentproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenwrappedtransparentproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenwrappedtransparentproxy.Contract.TokenwrappedtransparentproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.TokenwrappedtransparentproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.TokenwrappedtransparentproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenwrappedtransparentproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) Admin() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Admin(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) Admin() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Admin(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.ChangeAdmin(&_Tokenwrappedtransparentproxy.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.ChangeAdmin(&_Tokenwrappedtransparentproxy.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) Implementation() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Implementation(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) Implementation() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Implementation(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.UpgradeTo(&_Tokenwrappedtransparentproxy.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.UpgradeTo(&_Tokenwrappedtransparentproxy.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.UpgradeToAndCall(&_Tokenwrappedtransparentproxy.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.UpgradeToAndCall(&_Tokenwrappedtransparentproxy.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Fallback(&_Tokenwrappedtransparentproxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Fallback(&_Tokenwrappedtransparentproxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxySession) Receive() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Receive(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyTransactorSession) Receive() (*types.Transaction, error) {
	return _Tokenwrappedtransparentproxy.Contract.Receive(&_Tokenwrappedtransparentproxy.TransactOpts)
}

// TokenwrappedtransparentproxyAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyAdminChangedIterator struct {
	Event *TokenwrappedtransparentproxyAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenwrappedtransparentproxyAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedtransparentproxyAdminChanged)
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
		it.Event = new(TokenwrappedtransparentproxyAdminChanged)
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
func (it *TokenwrappedtransparentproxyAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedtransparentproxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedtransparentproxyAdminChanged represents a AdminChanged event raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TokenwrappedtransparentproxyAdminChangedIterator, error) {

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyAdminChangedIterator{contract: _Tokenwrappedtransparentproxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenwrappedtransparentproxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedtransparentproxyAdminChanged)
				if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) ParseAdminChanged(log types.Log) (*TokenwrappedtransparentproxyAdminChanged, error) {
	event := new(TokenwrappedtransparentproxyAdminChanged)
	if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenwrappedtransparentproxyBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyBeaconUpgradedIterator struct {
	Event *TokenwrappedtransparentproxyBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenwrappedtransparentproxyBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedtransparentproxyBeaconUpgraded)
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
		it.Event = new(TokenwrappedtransparentproxyBeaconUpgraded)
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
func (it *TokenwrappedtransparentproxyBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedtransparentproxyBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedtransparentproxyBeaconUpgraded represents a BeaconUpgraded event raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TokenwrappedtransparentproxyBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyBeaconUpgradedIterator{contract: _Tokenwrappedtransparentproxy.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TokenwrappedtransparentproxyBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedtransparentproxyBeaconUpgraded)
				if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) ParseBeaconUpgraded(log types.Log) (*TokenwrappedtransparentproxyBeaconUpgraded, error) {
	event := new(TokenwrappedtransparentproxyBeaconUpgraded)
	if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenwrappedtransparentproxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyUpgradedIterator struct {
	Event *TokenwrappedtransparentproxyUpgraded // Event containing the contract specifics and raw log

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
func (it *TokenwrappedtransparentproxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenwrappedtransparentproxyUpgraded)
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
		it.Event = new(TokenwrappedtransparentproxyUpgraded)
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
func (it *TokenwrappedtransparentproxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenwrappedtransparentproxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenwrappedtransparentproxyUpgraded represents a Upgraded event raised by the Tokenwrappedtransparentproxy contract.
type TokenwrappedtransparentproxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TokenwrappedtransparentproxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TokenwrappedtransparentproxyUpgradedIterator{contract: _Tokenwrappedtransparentproxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TokenwrappedtransparentproxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Tokenwrappedtransparentproxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenwrappedtransparentproxyUpgraded)
				if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Tokenwrappedtransparentproxy *TokenwrappedtransparentproxyFilterer) ParseUpgraded(log types.Log) (*TokenwrappedtransparentproxyUpgraded, error) {
	event := new(TokenwrappedtransparentproxyUpgraded)
	if err := _Tokenwrappedtransparentproxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootl2

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

// Polygonzkevmglobalexitrootl2MetaData contains all meta data concerning the Polygonzkevmglobalexitrootl2 contract.
var Polygonzkevmglobalexitrootl2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161023538038061023583398101604081905261002e9161003f565b6001600160a01b031660805261006c565b5f6020828403121561004f575f80fd5b81516001600160a01b0381168114610065575f80fd5b9392505050565b6080516101ab61008a5f395f818160a3015261010201526101ab5ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806301fd90441461004e578063257b36321461006a57806333d6247d14610089578063a3c573eb1461009e575b5f80fd5b61005760015481565b6040519081526020015b60405180910390f35b61005761007836600461015e565b5f6020819052908152604090205481565b61009c61009736600461015e565b6100ea565b005b6100c57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610061565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610159576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b5f6020828403121561016e575f80fd5b503591905056fea26469706673582212205108c6c4f924146b736832a1bdf696e20d900450207b7452462368d150f2c71c64736f6c63430008140033",
}

// Polygonzkevmglobalexitrootl2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonzkevmglobalexitrootl2MetaData.ABI instead.
var Polygonzkevmglobalexitrootl2ABI = Polygonzkevmglobalexitrootl2MetaData.ABI

// Polygonzkevmglobalexitrootl2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonzkevmglobalexitrootl2MetaData.Bin instead.
var Polygonzkevmglobalexitrootl2Bin = Polygonzkevmglobalexitrootl2MetaData.Bin

// DeployPolygonzkevmglobalexitrootl2 deploys a new Ethereum contract, binding an instance of Polygonzkevmglobalexitrootl2 to it.
func DeployPolygonzkevmglobalexitrootl2(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonzkevmglobalexitrootl2, error) {
	parsed, err := Polygonzkevmglobalexitrootl2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonzkevmglobalexitrootl2Bin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmglobalexitrootl2{Polygonzkevmglobalexitrootl2Caller: Polygonzkevmglobalexitrootl2Caller{contract: contract}, Polygonzkevmglobalexitrootl2Transactor: Polygonzkevmglobalexitrootl2Transactor{contract: contract}, Polygonzkevmglobalexitrootl2Filterer: Polygonzkevmglobalexitrootl2Filterer{contract: contract}}, nil
}

// Polygonzkevmglobalexitrootl2 is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2 struct {
	Polygonzkevmglobalexitrootl2Caller     // Read-only binding to the contract
	Polygonzkevmglobalexitrootl2Transactor // Write-only binding to the contract
	Polygonzkevmglobalexitrootl2Filterer   // Log filterer for contract events
}

// Polygonzkevmglobalexitrootl2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonzkevmglobalexitrootl2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonzkevmglobalexitrootl2Session struct {
	Contract     *Polygonzkevmglobalexitrootl2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootl2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonzkevmglobalexitrootl2CallerSession struct {
	Contract *Polygonzkevmglobalexitrootl2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// Polygonzkevmglobalexitrootl2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonzkevmglobalexitrootl2TransactorSession struct {
	Contract     *Polygonzkevmglobalexitrootl2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootl2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2Raw struct {
	Contract *Polygonzkevmglobalexitrootl2 // Generic contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootl2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2CallerRaw struct {
	Contract *Polygonzkevmglobalexitrootl2Caller // Generic read-only contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootl2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2TransactorRaw struct {
	Contract *Polygonzkevmglobalexitrootl2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootl2 creates a new instance of Polygonzkevmglobalexitrootl2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootl2, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2{Polygonzkevmglobalexitrootl2Caller: Polygonzkevmglobalexitrootl2Caller{contract: contract}, Polygonzkevmglobalexitrootl2Transactor: Polygonzkevmglobalexitrootl2Transactor{contract: contract}, Polygonzkevmglobalexitrootl2Filterer: Polygonzkevmglobalexitrootl2Filterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootl2Caller creates a new read-only instance of Polygonzkevmglobalexitrootl2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2Caller(address common.Address, caller bind.ContractCaller) (*Polygonzkevmglobalexitrootl2Caller, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2Caller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootl2Transactor creates a new write-only instance of Polygonzkevmglobalexitrootl2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2Transactor(address common.Address, transactor bind.ContractTransactor) (*Polygonzkevmglobalexitrootl2Transactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2Transactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootl2Filterer creates a new log filterer instance of Polygonzkevmglobalexitrootl2, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2Filterer(address common.Address, filterer bind.ContractFilterer) (*Polygonzkevmglobalexitrootl2Filterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2Filterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootl2 binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootl2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonzkevmglobalexitrootl2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootl2.Contract.Polygonzkevmglobalexitrootl2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.Polygonzkevmglobalexitrootl2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.Polygonzkevmglobalexitrootl2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootl2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Session) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootl2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2CallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootl2.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootl2.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootl2.CallOpts, arg0)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Session) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootl2.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootl2.CallOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootl2.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2 *Polygonzkevmglobalexitrootl2TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootl2.TransactOpts, newRoot)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmglobalexitrootl2mock

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

// Polygonzkevmglobalexitrootl2mockMetaData contains all meta data concerning the Polygonzkevmglobalexitrootl2mock contract.
var Polygonzkevmglobalexitrootl2mockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"setExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"setLastGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516102b93803806102b983398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161022761009260003960008181610100015261015f01526102276000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c806333d6247d1161005057806333d6247d146100c857806396e07459146100db578063a3c573eb146100fb57600080fd5b806301fd904414610077578063116c40c314610093578063257b3632146100a8575b600080fd5b61008060015481565b6040519081526020015b60405180910390f35b6100a66100a13660046101b6565b600155565b005b6100806100b63660046101b6565b60006020819052908152604090205481565b6100a66100d63660046101b6565b610147565b6100a66100e93660046101cf565b60009182526020829052604090912055565b6101227f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008a565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146100a1576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602082840312156101c857600080fd5b5035919050565b600080604083850312156101e257600080fd5b5050803592602090910135915056fea26469706673582212203dc131ea7b3c9a0e0d9f97c222b405f0086fc9beecf67bf79956cd48dcb6461d64736f6c63430008140033",
}

// Polygonzkevmglobalexitrootl2mockABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonzkevmglobalexitrootl2mockMetaData.ABI instead.
var Polygonzkevmglobalexitrootl2mockABI = Polygonzkevmglobalexitrootl2mockMetaData.ABI

// Polygonzkevmglobalexitrootl2mockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonzkevmglobalexitrootl2mockMetaData.Bin instead.
var Polygonzkevmglobalexitrootl2mockBin = Polygonzkevmglobalexitrootl2mockMetaData.Bin

// DeployPolygonzkevmglobalexitrootl2mock deploys a new Ethereum contract, binding an instance of Polygonzkevmglobalexitrootl2mock to it.
func DeployPolygonzkevmglobalexitrootl2mock(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonzkevmglobalexitrootl2mock, error) {
	parsed, err := Polygonzkevmglobalexitrootl2mockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonzkevmglobalexitrootl2mockBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmglobalexitrootl2mock{Polygonzkevmglobalexitrootl2mockCaller: Polygonzkevmglobalexitrootl2mockCaller{contract: contract}, Polygonzkevmglobalexitrootl2mockTransactor: Polygonzkevmglobalexitrootl2mockTransactor{contract: contract}, Polygonzkevmglobalexitrootl2mockFilterer: Polygonzkevmglobalexitrootl2mockFilterer{contract: contract}}, nil
}

// Polygonzkevmglobalexitrootl2mock is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mock struct {
	Polygonzkevmglobalexitrootl2mockCaller     // Read-only binding to the contract
	Polygonzkevmglobalexitrootl2mockTransactor // Write-only binding to the contract
	Polygonzkevmglobalexitrootl2mockFilterer   // Log filterer for contract events
}

// Polygonzkevmglobalexitrootl2mockCaller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2mockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2mockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonzkevmglobalexitrootl2mockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonzkevmglobalexitrootl2mockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonzkevmglobalexitrootl2mockSession struct {
	Contract     *Polygonzkevmglobalexitrootl2mock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootl2mockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonzkevmglobalexitrootl2mockCallerSession struct {
	Contract *Polygonzkevmglobalexitrootl2mockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// Polygonzkevmglobalexitrootl2mockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonzkevmglobalexitrootl2mockTransactorSession struct {
	Contract     *Polygonzkevmglobalexitrootl2mockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// Polygonzkevmglobalexitrootl2mockRaw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mockRaw struct {
	Contract *Polygonzkevmglobalexitrootl2mock // Generic contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootl2mockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mockCallerRaw struct {
	Contract *Polygonzkevmglobalexitrootl2mockCaller // Generic read-only contract binding to access the raw methods on
}

// Polygonzkevmglobalexitrootl2mockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonzkevmglobalexitrootl2mockTransactorRaw struct {
	Contract *Polygonzkevmglobalexitrootl2mockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmglobalexitrootl2mock creates a new instance of Polygonzkevmglobalexitrootl2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2mock(address common.Address, backend bind.ContractBackend) (*Polygonzkevmglobalexitrootl2mock, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2mock{Polygonzkevmglobalexitrootl2mockCaller: Polygonzkevmglobalexitrootl2mockCaller{contract: contract}, Polygonzkevmglobalexitrootl2mockTransactor: Polygonzkevmglobalexitrootl2mockTransactor{contract: contract}, Polygonzkevmglobalexitrootl2mockFilterer: Polygonzkevmglobalexitrootl2mockFilterer{contract: contract}}, nil
}

// NewPolygonzkevmglobalexitrootl2mockCaller creates a new read-only instance of Polygonzkevmglobalexitrootl2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2mockCaller(address common.Address, caller bind.ContractCaller) (*Polygonzkevmglobalexitrootl2mockCaller, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2mockCaller{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootl2mockTransactor creates a new write-only instance of Polygonzkevmglobalexitrootl2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2mockTransactor(address common.Address, transactor bind.ContractTransactor) (*Polygonzkevmglobalexitrootl2mockTransactor, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2mockTransactor{contract: contract}, nil
}

// NewPolygonzkevmglobalexitrootl2mockFilterer creates a new log filterer instance of Polygonzkevmglobalexitrootl2mock, bound to a specific deployed contract.
func NewPolygonzkevmglobalexitrootl2mockFilterer(address common.Address, filterer bind.ContractFilterer) (*Polygonzkevmglobalexitrootl2mockFilterer, error) {
	contract, err := bindPolygonzkevmglobalexitrootl2mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmglobalexitrootl2mockFilterer{contract: contract}, nil
}

// bindPolygonzkevmglobalexitrootl2mock binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmglobalexitrootl2mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonzkevmglobalexitrootl2mockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootl2mock.Contract.Polygonzkevmglobalexitrootl2mockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.Polygonzkevmglobalexitrootl2mockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.Polygonzkevmglobalexitrootl2mockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmglobalexitrootl2mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2mock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootl2mock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.BridgeAddress(&_Polygonzkevmglobalexitrootl2mock.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2mock.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootl2mock.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.GlobalExitRootMap(&_Polygonzkevmglobalexitrootl2mock.CallOpts, arg0)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmglobalexitrootl2mock.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootl2mock.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.LastRollupExitRoot(&_Polygonzkevmglobalexitrootl2mock.CallOpts)
}

// SetExitRoot is a paid mutator transaction binding the contract method 0x116c40c3.
//
// Solidity: function setExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactor) SetExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.contract.Transact(opts, "setExitRoot", newRoot)
}

// SetExitRoot is a paid mutator transaction binding the contract method 0x116c40c3.
//
// Solidity: function setExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) SetExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.SetExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, newRoot)
}

// SetExitRoot is a paid mutator transaction binding the contract method 0x116c40c3.
//
// Solidity: function setExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactorSession) SetExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.SetExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, newRoot)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x96e07459.
//
// Solidity: function setLastGlobalExitRoot(bytes32 globalExitRoot, uint256 blockNumber) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactor) SetLastGlobalExitRoot(opts *bind.TransactOpts, globalExitRoot [32]byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.contract.Transact(opts, "setLastGlobalExitRoot", globalExitRoot, blockNumber)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x96e07459.
//
// Solidity: function setLastGlobalExitRoot(bytes32 globalExitRoot, uint256 blockNumber) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) SetLastGlobalExitRoot(globalExitRoot [32]byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.SetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, globalExitRoot, blockNumber)
}

// SetLastGlobalExitRoot is a paid mutator transaction binding the contract method 0x96e07459.
//
// Solidity: function setLastGlobalExitRoot(bytes32 globalExitRoot, uint256 blockNumber) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactorSession) SetLastGlobalExitRoot(globalExitRoot [32]byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.SetLastGlobalExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, globalExitRoot, blockNumber)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Polygonzkevmglobalexitrootl2mock *Polygonzkevmglobalexitrootl2mockTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmglobalexitrootl2mock.Contract.UpdateExitRoot(&_Polygonzkevmglobalexitrootl2mock.TransactOpts, newRoot)
}

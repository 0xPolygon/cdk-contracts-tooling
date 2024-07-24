// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootnopush0

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

// Globalexitrootnopush0MetaData contains all meta data concerning the Globalexitrootnopush0 contract.
var Globalexitrootnopush0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"}],\"name\":\"UpdateGlobalExitRoot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastGlobalExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMainnetExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b506040516103f83803806103f883398101604081905261002f91610062565b6001600160a01b0391821660a05216608052610095565b80516001600160a01b038116811461005d57600080fd5b919050565b6000806040838503121561007557600080fd5b61007e83610046565b915061008c60208401610046565b90509250929050565b60805160a0516103316100c76000396000818160e901526101bd015260008181610135015261017401526103316000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806333d6247d1161005b57806333d6247d146100c75780633ed691ef146100dc5780635ec6a8df146100e4578063a3c573eb1461013057600080fd5b806301fd904414610082578063257b36321461009e578063319cf735146100be575b600080fd5b61008b60005481565b6040519081526020015b60405180910390f35b61008b6100ac3660046102e2565b60026020526000908152604090205481565b61008b60015481565b6100da6100d53660046102e2565b610157565b005b61008b6102a6565b61010b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610095565b61010b7f000000000000000000000000000000000000000000000000000000000000000081565b60005460015473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036101a65750600182905581610222565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633036101f0576000839055829150610222565b6040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051602080820184905281830185905282518083038401815260609092019092528051910120600090600081815260026020526040812054919250036102a05760008181526002602052604080822042905551849184917f61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce39190a35b50505050565b60006102dd600154600054604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b905090565b6000602082840312156102f457600080fd5b503591905056fea2646970667358221220bc23c6d5d3992802bdfd06ef45362230dcda7d33db81b1dc3ef40d86219e81c864736f6c63430008110033",
}

// Globalexitrootnopush0ABI is the input ABI used to generate the binding from.
// Deprecated: Use Globalexitrootnopush0MetaData.ABI instead.
var Globalexitrootnopush0ABI = Globalexitrootnopush0MetaData.ABI

// Globalexitrootnopush0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Globalexitrootnopush0MetaData.Bin instead.
var Globalexitrootnopush0Bin = Globalexitrootnopush0MetaData.Bin

// DeployGlobalexitrootnopush0 deploys a new Ethereum contract, binding an instance of Globalexitrootnopush0 to it.
func DeployGlobalexitrootnopush0(auth *bind.TransactOpts, backend bind.ContractBackend, _rollupManager common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Globalexitrootnopush0, error) {
	parsed, err := Globalexitrootnopush0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Globalexitrootnopush0Bin), backend, _rollupManager, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootnopush0{Globalexitrootnopush0Caller: Globalexitrootnopush0Caller{contract: contract}, Globalexitrootnopush0Transactor: Globalexitrootnopush0Transactor{contract: contract}, Globalexitrootnopush0Filterer: Globalexitrootnopush0Filterer{contract: contract}}, nil
}

// Globalexitrootnopush0 is an auto generated Go binding around an Ethereum contract.
type Globalexitrootnopush0 struct {
	Globalexitrootnopush0Caller     // Read-only binding to the contract
	Globalexitrootnopush0Transactor // Write-only binding to the contract
	Globalexitrootnopush0Filterer   // Log filterer for contract events
}

// Globalexitrootnopush0Caller is an auto generated read-only Go binding around an Ethereum contract.
type Globalexitrootnopush0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootnopush0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Globalexitrootnopush0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootnopush0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Globalexitrootnopush0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootnopush0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Globalexitrootnopush0Session struct {
	Contract     *Globalexitrootnopush0 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Globalexitrootnopush0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Globalexitrootnopush0CallerSession struct {
	Contract *Globalexitrootnopush0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Globalexitrootnopush0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Globalexitrootnopush0TransactorSession struct {
	Contract     *Globalexitrootnopush0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Globalexitrootnopush0Raw is an auto generated low-level Go binding around an Ethereum contract.
type Globalexitrootnopush0Raw struct {
	Contract *Globalexitrootnopush0 // Generic contract binding to access the raw methods on
}

// Globalexitrootnopush0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Globalexitrootnopush0CallerRaw struct {
	Contract *Globalexitrootnopush0Caller // Generic read-only contract binding to access the raw methods on
}

// Globalexitrootnopush0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Globalexitrootnopush0TransactorRaw struct {
	Contract *Globalexitrootnopush0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootnopush0 creates a new instance of Globalexitrootnopush0, bound to a specific deployed contract.
func NewGlobalexitrootnopush0(address common.Address, backend bind.ContractBackend) (*Globalexitrootnopush0, error) {
	contract, err := bindGlobalexitrootnopush0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootnopush0{Globalexitrootnopush0Caller: Globalexitrootnopush0Caller{contract: contract}, Globalexitrootnopush0Transactor: Globalexitrootnopush0Transactor{contract: contract}, Globalexitrootnopush0Filterer: Globalexitrootnopush0Filterer{contract: contract}}, nil
}

// NewGlobalexitrootnopush0Caller creates a new read-only instance of Globalexitrootnopush0, bound to a specific deployed contract.
func NewGlobalexitrootnopush0Caller(address common.Address, caller bind.ContractCaller) (*Globalexitrootnopush0Caller, error) {
	contract, err := bindGlobalexitrootnopush0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootnopush0Caller{contract: contract}, nil
}

// NewGlobalexitrootnopush0Transactor creates a new write-only instance of Globalexitrootnopush0, bound to a specific deployed contract.
func NewGlobalexitrootnopush0Transactor(address common.Address, transactor bind.ContractTransactor) (*Globalexitrootnopush0Transactor, error) {
	contract, err := bindGlobalexitrootnopush0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootnopush0Transactor{contract: contract}, nil
}

// NewGlobalexitrootnopush0Filterer creates a new log filterer instance of Globalexitrootnopush0, bound to a specific deployed contract.
func NewGlobalexitrootnopush0Filterer(address common.Address, filterer bind.ContractFilterer) (*Globalexitrootnopush0Filterer, error) {
	contract, err := bindGlobalexitrootnopush0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootnopush0Filterer{contract: contract}, nil
}

// bindGlobalexitrootnopush0 binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootnopush0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Globalexitrootnopush0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootnopush0 *Globalexitrootnopush0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootnopush0.Contract.Globalexitrootnopush0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootnopush0 *Globalexitrootnopush0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.Globalexitrootnopush0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootnopush0 *Globalexitrootnopush0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.Globalexitrootnopush0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootnopush0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootnopush0 *Globalexitrootnopush0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootnopush0 *Globalexitrootnopush0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) BridgeAddress() (common.Address, error) {
	return _Globalexitrootnopush0.Contract.BridgeAddress(&_Globalexitrootnopush0.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootnopush0.Contract.BridgeAddress(&_Globalexitrootnopush0.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) GetLastGlobalExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "getLastGlobalExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.GetLastGlobalExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// GetLastGlobalExitRoot is a free data retrieval call binding the contract method 0x3ed691ef.
//
// Solidity: function getLastGlobalExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) GetLastGlobalExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.GetLastGlobalExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootnopush0.Contract.GlobalExitRootMap(&_Globalexitrootnopush0.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootnopush0.Contract.GlobalExitRootMap(&_Globalexitrootnopush0.CallOpts, arg0)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) LastMainnetExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "lastMainnetExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) LastMainnetExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.LastMainnetExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// LastMainnetExitRoot is a free data retrieval call binding the contract method 0x319cf735.
//
// Solidity: function lastMainnetExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) LastMainnetExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.LastMainnetExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.LastRollupExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootnopush0.Contract.LastRollupExitRoot(&_Globalexitrootnopush0.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Caller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootnopush0.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) RollupManager() (common.Address, error) {
	return _Globalexitrootnopush0.Contract.RollupManager(&_Globalexitrootnopush0.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Globalexitrootnopush0 *Globalexitrootnopush0CallerSession) RollupManager() (common.Address, error) {
	return _Globalexitrootnopush0.Contract.RollupManager(&_Globalexitrootnopush0.CallOpts)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootnopush0 *Globalexitrootnopush0Transactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootnopush0.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootnopush0 *Globalexitrootnopush0Session) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.UpdateExitRoot(&_Globalexitrootnopush0.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootnopush0 *Globalexitrootnopush0TransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootnopush0.Contract.UpdateExitRoot(&_Globalexitrootnopush0.TransactOpts, newRoot)
}

// Globalexitrootnopush0UpdateGlobalExitRootIterator is returned from FilterUpdateGlobalExitRoot and is used to iterate over the raw logs and unpacked data for UpdateGlobalExitRoot events raised by the Globalexitrootnopush0 contract.
type Globalexitrootnopush0UpdateGlobalExitRootIterator struct {
	Event *Globalexitrootnopush0UpdateGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootnopush0UpdateGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootnopush0UpdateGlobalExitRoot)
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
		it.Event = new(Globalexitrootnopush0UpdateGlobalExitRoot)
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
func (it *Globalexitrootnopush0UpdateGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootnopush0UpdateGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootnopush0UpdateGlobalExitRoot represents a UpdateGlobalExitRoot event raised by the Globalexitrootnopush0 contract.
type Globalexitrootnopush0UpdateGlobalExitRoot struct {
	MainnetExitRoot [32]byte
	RollupExitRoot  [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateGlobalExitRoot is a free log retrieval operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Filterer) FilterUpdateGlobalExitRoot(opts *bind.FilterOpts, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (*Globalexitrootnopush0UpdateGlobalExitRootIterator, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Globalexitrootnopush0.contract.FilterLogs(opts, "UpdateGlobalExitRoot", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootnopush0UpdateGlobalExitRootIterator{contract: _Globalexitrootnopush0.contract, event: "UpdateGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchUpdateGlobalExitRoot is a free log subscription operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Filterer) WatchUpdateGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootnopush0UpdateGlobalExitRoot, mainnetExitRoot [][32]byte, rollupExitRoot [][32]byte) (event.Subscription, error) {

	var mainnetExitRootRule []interface{}
	for _, mainnetExitRootItem := range mainnetExitRoot {
		mainnetExitRootRule = append(mainnetExitRootRule, mainnetExitRootItem)
	}
	var rollupExitRootRule []interface{}
	for _, rollupExitRootItem := range rollupExitRoot {
		rollupExitRootRule = append(rollupExitRootRule, rollupExitRootItem)
	}

	logs, sub, err := _Globalexitrootnopush0.contract.WatchLogs(opts, "UpdateGlobalExitRoot", mainnetExitRootRule, rollupExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootnopush0UpdateGlobalExitRoot)
				if err := _Globalexitrootnopush0.contract.UnpackLog(event, "UpdateGlobalExitRoot", log); err != nil {
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

// ParseUpdateGlobalExitRoot is a log parse operation binding the contract event 0x61014378f82a0d809aefaf87a8ac9505b89c321808287a6e7810f29304c1fce3.
//
// Solidity: event UpdateGlobalExitRoot(bytes32 indexed mainnetExitRoot, bytes32 indexed rollupExitRoot)
func (_Globalexitrootnopush0 *Globalexitrootnopush0Filterer) ParseUpdateGlobalExitRoot(log types.Log) (*Globalexitrootnopush0UpdateGlobalExitRoot, error) {
	event := new(Globalexitrootnopush0UpdateGlobalExitRoot)
	if err := _Globalexitrootnopush0.contract.UnpackLog(event, "UpdateGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

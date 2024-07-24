// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pessimisticGlobalExitRoot

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

// PessimisticGlobalExitRootMetaData contains all meta data concerning the PessimisticGlobalExitRoot contract.
var PessimisticGlobalExitRootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516108be3803806108be83398101604081905261002e91610239565b6001600160a01b0381166080526100727f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c51765f8051602061089e8339815191526100ae565b6100895f8051602061089e833981519152806100ae565b6100a05f8051602061089e833981519152336100f8565b6100a861017c565b50610266565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b5f8281526034602090815260408083206001600160a01b038516845290915290205460ff16610178575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45b5050565b5f54610100900460ff16156101e75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610237575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610249575f80fd5b81516001600160a01b038116811461025f575f80fd5b9392505050565b6080516106196102855f395f81816101d4015261026f01526106195ff3fe608060405234801561000f575f80fd5b50600436106100c4575f3560e01c806336568abe1161007d578063a217fddf11610058578063a217fddf146101c8578063a3c573eb146101cf578063d547741f1461021b575f80fd5b806336568abe1461014d57806348da4bdd1461016057806391d1485414610173575f80fd5b8063257b3632116100ad578063257b3632146101065780632f2ff15d1461012557806333d6247d1461013a575f80fd5b806301fd9044146100c8578063248a9ca3146100e4575b5f80fd5b6100d160665481565b6040519081526020015b60405180910390f35b6100d16100f2366004610586565b5f9081526034602052604090206001015490565b6100d1610114366004610586565b60656020525f908152604090205481565b61013861013336600461059d565b61022e565b005b610138610148366004610586565b610257565b61013861015b36600461059d565b6102cb565b61013861016e366004610586565b610328565b6101b861018136600461059d565b5f91825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b60405190151581526020016100db565b6100d15f81565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100db565b61013861022936600461059d565b610379565b5f828152603460205260409020600101546102488161039d565b61025283836103aa565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102c6576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606655565b73ffffffffffffffffffffffffffffffffffffffff8116331461031a576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103248282610465565b5050565b7f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c51766103528161039d565b5f82815260656020526040812054900361032457505f908152606560205260409020429055565b5f828152603460205260409020600101546103938161039d565b6102528383610465565b6103a7813361051e565b50565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610324575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1615610324575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610324576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215610596575f80fd5b5035919050565b5f80604083850312156105ae575f80fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff811681146105d8575f80fd5b80915050925092905056fea2646970667358221220bea40990a781389a495e3c926a4a77689f1de82e6f907a45a06af8e13f589d7b64736f6c634300081400335ee7b1f4df16da4d137fdcc53b4b785d709ccbf7d129b16d96c5c3bedc99aef7",
}

// PessimisticGlobalExitRootABI is the input ABI used to generate the binding from.
// Deprecated: Use PessimisticGlobalExitRootMetaData.ABI instead.
var PessimisticGlobalExitRootABI = PessimisticGlobalExitRootMetaData.ABI

// PessimisticGlobalExitRootBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PessimisticGlobalExitRootMetaData.Bin instead.
var PessimisticGlobalExitRootBin = PessimisticGlobalExitRootMetaData.Bin

// DeployPessimisticGlobalExitRoot deploys a new Ethereum contract, binding an instance of PessimisticGlobalExitRoot to it.
func DeployPessimisticGlobalExitRoot(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *PessimisticGlobalExitRoot, error) {
	parsed, err := PessimisticGlobalExitRootMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PessimisticGlobalExitRootBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PessimisticGlobalExitRoot{PessimisticGlobalExitRootCaller: PessimisticGlobalExitRootCaller{contract: contract}, PessimisticGlobalExitRootTransactor: PessimisticGlobalExitRootTransactor{contract: contract}, PessimisticGlobalExitRootFilterer: PessimisticGlobalExitRootFilterer{contract: contract}}, nil
}

// PessimisticGlobalExitRoot is an auto generated Go binding around an Ethereum contract.
type PessimisticGlobalExitRoot struct {
	PessimisticGlobalExitRootCaller     // Read-only binding to the contract
	PessimisticGlobalExitRootTransactor // Write-only binding to the contract
	PessimisticGlobalExitRootFilterer   // Log filterer for contract events
}

// PessimisticGlobalExitRootCaller is an auto generated read-only Go binding around an Ethereum contract.
type PessimisticGlobalExitRootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticGlobalExitRootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PessimisticGlobalExitRootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticGlobalExitRootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PessimisticGlobalExitRootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticGlobalExitRootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PessimisticGlobalExitRootSession struct {
	Contract     *PessimisticGlobalExitRoot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PessimisticGlobalExitRootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PessimisticGlobalExitRootCallerSession struct {
	Contract *PessimisticGlobalExitRootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PessimisticGlobalExitRootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PessimisticGlobalExitRootTransactorSession struct {
	Contract     *PessimisticGlobalExitRootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PessimisticGlobalExitRootRaw is an auto generated low-level Go binding around an Ethereum contract.
type PessimisticGlobalExitRootRaw struct {
	Contract *PessimisticGlobalExitRoot // Generic contract binding to access the raw methods on
}

// PessimisticGlobalExitRootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PessimisticGlobalExitRootCallerRaw struct {
	Contract *PessimisticGlobalExitRootCaller // Generic read-only contract binding to access the raw methods on
}

// PessimisticGlobalExitRootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PessimisticGlobalExitRootTransactorRaw struct {
	Contract *PessimisticGlobalExitRootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPessimisticGlobalExitRoot creates a new instance of PessimisticGlobalExitRoot, bound to a specific deployed contract.
func NewPessimisticGlobalExitRoot(address common.Address, backend bind.ContractBackend) (*PessimisticGlobalExitRoot, error) {
	contract, err := bindPessimisticGlobalExitRoot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRoot{PessimisticGlobalExitRootCaller: PessimisticGlobalExitRootCaller{contract: contract}, PessimisticGlobalExitRootTransactor: PessimisticGlobalExitRootTransactor{contract: contract}, PessimisticGlobalExitRootFilterer: PessimisticGlobalExitRootFilterer{contract: contract}}, nil
}

// NewPessimisticGlobalExitRootCaller creates a new read-only instance of PessimisticGlobalExitRoot, bound to a specific deployed contract.
func NewPessimisticGlobalExitRootCaller(address common.Address, caller bind.ContractCaller) (*PessimisticGlobalExitRootCaller, error) {
	contract, err := bindPessimisticGlobalExitRoot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootCaller{contract: contract}, nil
}

// NewPessimisticGlobalExitRootTransactor creates a new write-only instance of PessimisticGlobalExitRoot, bound to a specific deployed contract.
func NewPessimisticGlobalExitRootTransactor(address common.Address, transactor bind.ContractTransactor) (*PessimisticGlobalExitRootTransactor, error) {
	contract, err := bindPessimisticGlobalExitRoot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootTransactor{contract: contract}, nil
}

// NewPessimisticGlobalExitRootFilterer creates a new log filterer instance of PessimisticGlobalExitRoot, bound to a specific deployed contract.
func NewPessimisticGlobalExitRootFilterer(address common.Address, filterer bind.ContractFilterer) (*PessimisticGlobalExitRootFilterer, error) {
	contract, err := bindPessimisticGlobalExitRoot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootFilterer{contract: contract}, nil
}

// bindPessimisticGlobalExitRoot binds a generic wrapper to an already deployed contract.
func bindPessimisticGlobalExitRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PessimisticGlobalExitRootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PessimisticGlobalExitRoot.Contract.PessimisticGlobalExitRootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.PessimisticGlobalExitRootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.PessimisticGlobalExitRootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PessimisticGlobalExitRoot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.DEFAULTADMINROLE(&_PessimisticGlobalExitRoot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.DEFAULTADMINROLE(&_PessimisticGlobalExitRoot.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) BridgeAddress() (common.Address, error) {
	return _PessimisticGlobalExitRoot.Contract.BridgeAddress(&_PessimisticGlobalExitRoot.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) BridgeAddress() (common.Address, error) {
	return _PessimisticGlobalExitRoot.Contract.BridgeAddress(&_PessimisticGlobalExitRoot.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.GetRoleAdmin(&_PessimisticGlobalExitRoot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.GetRoleAdmin(&_PessimisticGlobalExitRoot.CallOpts, role)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PessimisticGlobalExitRoot.Contract.GlobalExitRootMap(&_PessimisticGlobalExitRoot.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _PessimisticGlobalExitRoot.Contract.GlobalExitRootMap(&_PessimisticGlobalExitRoot.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PessimisticGlobalExitRoot.Contract.HasRole(&_PessimisticGlobalExitRoot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PessimisticGlobalExitRoot.Contract.HasRole(&_PessimisticGlobalExitRoot.CallOpts, role, account)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PessimisticGlobalExitRoot.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) LastRollupExitRoot() ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.LastRollupExitRoot(&_PessimisticGlobalExitRoot.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _PessimisticGlobalExitRoot.Contract.LastRollupExitRoot(&_PessimisticGlobalExitRoot.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.GrantRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.GrantRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.RenounceRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.RenounceRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.RevokeRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.RevokeRole(&_PessimisticGlobalExitRoot.TransactOpts, role, account)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.UpdateExitRoot(&_PessimisticGlobalExitRoot.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.UpdateExitRoot(&_PessimisticGlobalExitRoot.TransactOpts, newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.contract.Transact(opts, "updateGlobalExitRoot", _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootSession) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.UpdateGlobalExitRoot(&_PessimisticGlobalExitRoot.TransactOpts, _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootTransactorSession) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _PessimisticGlobalExitRoot.Contract.UpdateGlobalExitRoot(&_PessimisticGlobalExitRoot.TransactOpts, _newRoot)
}

// PessimisticGlobalExitRootInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootInitializedIterator struct {
	Event *PessimisticGlobalExitRootInitialized // Event containing the contract specifics and raw log

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
func (it *PessimisticGlobalExitRootInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticGlobalExitRootInitialized)
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
		it.Event = new(PessimisticGlobalExitRootInitialized)
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
func (it *PessimisticGlobalExitRootInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticGlobalExitRootInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticGlobalExitRootInitialized represents a Initialized event raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) FilterInitialized(opts *bind.FilterOpts) (*PessimisticGlobalExitRootInitializedIterator, error) {

	logs, sub, err := _PessimisticGlobalExitRoot.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootInitializedIterator{contract: _PessimisticGlobalExitRoot.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PessimisticGlobalExitRootInitialized) (event.Subscription, error) {

	logs, sub, err := _PessimisticGlobalExitRoot.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticGlobalExitRootInitialized)
				if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) ParseInitialized(log types.Log) (*PessimisticGlobalExitRootInitialized, error) {
	event := new(PessimisticGlobalExitRootInitialized)
	if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticGlobalExitRootRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleAdminChangedIterator struct {
	Event *PessimisticGlobalExitRootRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PessimisticGlobalExitRootRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticGlobalExitRootRoleAdminChanged)
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
		it.Event = new(PessimisticGlobalExitRootRoleAdminChanged)
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
func (it *PessimisticGlobalExitRootRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticGlobalExitRootRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticGlobalExitRootRoleAdminChanged represents a RoleAdminChanged event raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PessimisticGlobalExitRootRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootRoleAdminChangedIterator{contract: _PessimisticGlobalExitRoot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PessimisticGlobalExitRootRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticGlobalExitRootRoleAdminChanged)
				if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) ParseRoleAdminChanged(log types.Log) (*PessimisticGlobalExitRootRoleAdminChanged, error) {
	event := new(PessimisticGlobalExitRootRoleAdminChanged)
	if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticGlobalExitRootRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleGrantedIterator struct {
	Event *PessimisticGlobalExitRootRoleGranted // Event containing the contract specifics and raw log

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
func (it *PessimisticGlobalExitRootRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticGlobalExitRootRoleGranted)
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
		it.Event = new(PessimisticGlobalExitRootRoleGranted)
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
func (it *PessimisticGlobalExitRootRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticGlobalExitRootRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticGlobalExitRootRoleGranted represents a RoleGranted event raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PessimisticGlobalExitRootRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootRoleGrantedIterator{contract: _PessimisticGlobalExitRoot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PessimisticGlobalExitRootRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticGlobalExitRootRoleGranted)
				if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) ParseRoleGranted(log types.Log) (*PessimisticGlobalExitRootRoleGranted, error) {
	event := new(PessimisticGlobalExitRootRoleGranted)
	if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticGlobalExitRootRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleRevokedIterator struct {
	Event *PessimisticGlobalExitRootRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PessimisticGlobalExitRootRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticGlobalExitRootRoleRevoked)
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
		it.Event = new(PessimisticGlobalExitRootRoleRevoked)
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
func (it *PessimisticGlobalExitRootRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticGlobalExitRootRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticGlobalExitRootRoleRevoked represents a RoleRevoked event raised by the PessimisticGlobalExitRoot contract.
type PessimisticGlobalExitRootRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PessimisticGlobalExitRootRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticGlobalExitRootRoleRevokedIterator{contract: _PessimisticGlobalExitRoot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PessimisticGlobalExitRootRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PessimisticGlobalExitRoot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticGlobalExitRootRoleRevoked)
				if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PessimisticGlobalExitRoot *PessimisticGlobalExitRootFilterer) ParseRoleRevoked(log types.Log) (*PessimisticGlobalExitRootRoleRevoked, error) {
	event := new(PessimisticGlobalExitRootRoleRevoked)
	if err := _PessimisticGlobalExitRoot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

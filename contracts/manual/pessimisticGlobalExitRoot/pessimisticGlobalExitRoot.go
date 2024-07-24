// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pessimisticglobalexitroot

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

// PessimisticglobalexitrootMetaData contains all meta data concerning the Pessimisticglobalexitroot contract.
var PessimisticglobalexitrootMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516108be3803806108be83398101604081905261002e91610239565b6001600160a01b0381166080526100727f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c51765f8051602061089e8339815191526100ae565b6100895f8051602061089e833981519152806100ae565b6100a05f8051602061089e833981519152336100f8565b6100a861017c565b50610266565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b5f8281526034602090815260408083206001600160a01b038516845290915290205460ff16610178575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45b5050565b5f54610100900460ff16156101e75760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610237575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610249575f80fd5b81516001600160a01b038116811461025f575f80fd5b9392505050565b6080516106196102855f395f81816101d4015261026f01526106195ff3fe608060405234801561000f575f80fd5b50600436106100c4575f3560e01c806336568abe1161007d578063a217fddf11610058578063a217fddf146101c8578063a3c573eb146101cf578063d547741f1461021b575f80fd5b806336568abe1461014d57806348da4bdd1461016057806391d1485414610173575f80fd5b8063257b3632116100ad578063257b3632146101065780632f2ff15d1461012557806333d6247d1461013a575f80fd5b806301fd9044146100c8578063248a9ca3146100e4575b5f80fd5b6100d160665481565b6040519081526020015b60405180910390f35b6100d16100f2366004610586565b5f9081526034602052604090206001015490565b6100d1610114366004610586565b60656020525f908152604090205481565b61013861013336600461059d565b61022e565b005b610138610148366004610586565b610257565b61013861015b36600461059d565b6102cb565b61013861016e366004610586565b610328565b6101b861018136600461059d565b5f91825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b60405190151581526020016100db565b6100d15f81565b6101f67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100db565b61013861022936600461059d565b610379565b5f828152603460205260409020600101546102488161039d565b61025283836103aa565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102c6576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606655565b73ffffffffffffffffffffffffffffffffffffffff8116331461031a576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103248282610465565b5050565b7f7b95520991dfda409891be0afa2635b63540f92ee996fda0bf695a166e5c51766103528161039d565b5f82815260656020526040812054900361032457505f908152606560205260409020429055565b5f828152603460205260409020600101546103938161039d565b6102528383610465565b6103a7813361051e565b50565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610324575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1615610324575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610324576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60208284031215610596575f80fd5b5035919050565b5f80604083850312156105ae575f80fd5b82359150602083013573ffffffffffffffffffffffffffffffffffffffff811681146105d8575f80fd5b80915050925092905056fea2646970667358221220bea40990a781389a495e3c926a4a77689f1de82e6f907a45a06af8e13f589d7b64736f6c634300081400335ee7b1f4df16da4d137fdcc53b4b785d709ccbf7d129b16d96c5c3bedc99aef7",
}

// PessimisticglobalexitrootABI is the input ABI used to generate the binding from.
// Deprecated: Use PessimisticglobalexitrootMetaData.ABI instead.
var PessimisticglobalexitrootABI = PessimisticglobalexitrootMetaData.ABI

// PessimisticglobalexitrootBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PessimisticglobalexitrootMetaData.Bin instead.
var PessimisticglobalexitrootBin = PessimisticglobalexitrootMetaData.Bin

// DeployPessimisticglobalexitroot deploys a new Ethereum contract, binding an instance of Pessimisticglobalexitroot to it.
func DeployPessimisticglobalexitroot(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Pessimisticglobalexitroot, error) {
	parsed, err := PessimisticglobalexitrootMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PessimisticglobalexitrootBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pessimisticglobalexitroot{PessimisticglobalexitrootCaller: PessimisticglobalexitrootCaller{contract: contract}, PessimisticglobalexitrootTransactor: PessimisticglobalexitrootTransactor{contract: contract}, PessimisticglobalexitrootFilterer: PessimisticglobalexitrootFilterer{contract: contract}}, nil
}

// Pessimisticglobalexitroot is an auto generated Go binding around an Ethereum contract.
type Pessimisticglobalexitroot struct {
	PessimisticglobalexitrootCaller     // Read-only binding to the contract
	PessimisticglobalexitrootTransactor // Write-only binding to the contract
	PessimisticglobalexitrootFilterer   // Log filterer for contract events
}

// PessimisticglobalexitrootCaller is an auto generated read-only Go binding around an Ethereum contract.
type PessimisticglobalexitrootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticglobalexitrootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PessimisticglobalexitrootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticglobalexitrootFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PessimisticglobalexitrootFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PessimisticglobalexitrootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PessimisticglobalexitrootSession struct {
	Contract     *Pessimisticglobalexitroot // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PessimisticglobalexitrootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PessimisticglobalexitrootCallerSession struct {
	Contract *PessimisticglobalexitrootCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PessimisticglobalexitrootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PessimisticglobalexitrootTransactorSession struct {
	Contract     *PessimisticglobalexitrootTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PessimisticglobalexitrootRaw is an auto generated low-level Go binding around an Ethereum contract.
type PessimisticglobalexitrootRaw struct {
	Contract *Pessimisticglobalexitroot // Generic contract binding to access the raw methods on
}

// PessimisticglobalexitrootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PessimisticglobalexitrootCallerRaw struct {
	Contract *PessimisticglobalexitrootCaller // Generic read-only contract binding to access the raw methods on
}

// PessimisticglobalexitrootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PessimisticglobalexitrootTransactorRaw struct {
	Contract *PessimisticglobalexitrootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPessimisticglobalexitroot creates a new instance of Pessimisticglobalexitroot, bound to a specific deployed contract.
func NewPessimisticglobalexitroot(address common.Address, backend bind.ContractBackend) (*Pessimisticglobalexitroot, error) {
	contract, err := bindPessimisticglobalexitroot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pessimisticglobalexitroot{PessimisticglobalexitrootCaller: PessimisticglobalexitrootCaller{contract: contract}, PessimisticglobalexitrootTransactor: PessimisticglobalexitrootTransactor{contract: contract}, PessimisticglobalexitrootFilterer: PessimisticglobalexitrootFilterer{contract: contract}}, nil
}

// NewPessimisticglobalexitrootCaller creates a new read-only instance of Pessimisticglobalexitroot, bound to a specific deployed contract.
func NewPessimisticglobalexitrootCaller(address common.Address, caller bind.ContractCaller) (*PessimisticglobalexitrootCaller, error) {
	contract, err := bindPessimisticglobalexitroot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootCaller{contract: contract}, nil
}

// NewPessimisticglobalexitrootTransactor creates a new write-only instance of Pessimisticglobalexitroot, bound to a specific deployed contract.
func NewPessimisticglobalexitrootTransactor(address common.Address, transactor bind.ContractTransactor) (*PessimisticglobalexitrootTransactor, error) {
	contract, err := bindPessimisticglobalexitroot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootTransactor{contract: contract}, nil
}

// NewPessimisticglobalexitrootFilterer creates a new log filterer instance of Pessimisticglobalexitroot, bound to a specific deployed contract.
func NewPessimisticglobalexitrootFilterer(address common.Address, filterer bind.ContractFilterer) (*PessimisticglobalexitrootFilterer, error) {
	contract, err := bindPessimisticglobalexitroot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootFilterer{contract: contract}, nil
}

// bindPessimisticglobalexitroot binds a generic wrapper to an already deployed contract.
func bindPessimisticglobalexitroot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PessimisticglobalexitrootMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pessimisticglobalexitroot.Contract.PessimisticglobalexitrootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.PessimisticglobalexitrootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.PessimisticglobalexitrootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pessimisticglobalexitroot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.DEFAULTADMINROLE(&_Pessimisticglobalexitroot.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.DEFAULTADMINROLE(&_Pessimisticglobalexitroot.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) BridgeAddress() (common.Address, error) {
	return _Pessimisticglobalexitroot.Contract.BridgeAddress(&_Pessimisticglobalexitroot.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) BridgeAddress() (common.Address, error) {
	return _Pessimisticglobalexitroot.Contract.BridgeAddress(&_Pessimisticglobalexitroot.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.GetRoleAdmin(&_Pessimisticglobalexitroot.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.GetRoleAdmin(&_Pessimisticglobalexitroot.CallOpts, role)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Pessimisticglobalexitroot.Contract.GlobalExitRootMap(&_Pessimisticglobalexitroot.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Pessimisticglobalexitroot.Contract.GlobalExitRootMap(&_Pessimisticglobalexitroot.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pessimisticglobalexitroot.Contract.HasRole(&_Pessimisticglobalexitroot.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pessimisticglobalexitroot.Contract.HasRole(&_Pessimisticglobalexitroot.CallOpts, role, account)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pessimisticglobalexitroot.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) LastRollupExitRoot() ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.LastRollupExitRoot(&_Pessimisticglobalexitroot.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Pessimisticglobalexitroot.Contract.LastRollupExitRoot(&_Pessimisticglobalexitroot.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.GrantRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.GrantRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.RenounceRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.RenounceRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.RevokeRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.RevokeRole(&_Pessimisticglobalexitroot.TransactOpts, role, account)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.UpdateExitRoot(&_Pessimisticglobalexitroot.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.UpdateExitRoot(&_Pessimisticglobalexitroot.TransactOpts, newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.contract.Transact(opts, "updateGlobalExitRoot", _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootSession) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.UpdateGlobalExitRoot(&_Pessimisticglobalexitroot.TransactOpts, _newRoot)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x48da4bdd.
//
// Solidity: function updateGlobalExitRoot(bytes32 _newRoot) returns()
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootTransactorSession) UpdateGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Pessimisticglobalexitroot.Contract.UpdateGlobalExitRoot(&_Pessimisticglobalexitroot.TransactOpts, _newRoot)
}

// PessimisticglobalexitrootInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootInitializedIterator struct {
	Event *PessimisticglobalexitrootInitialized // Event containing the contract specifics and raw log

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
func (it *PessimisticglobalexitrootInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticglobalexitrootInitialized)
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
		it.Event = new(PessimisticglobalexitrootInitialized)
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
func (it *PessimisticglobalexitrootInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticglobalexitrootInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticglobalexitrootInitialized represents a Initialized event raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) FilterInitialized(opts *bind.FilterOpts) (*PessimisticglobalexitrootInitializedIterator, error) {

	logs, sub, err := _Pessimisticglobalexitroot.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootInitializedIterator{contract: _Pessimisticglobalexitroot.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PessimisticglobalexitrootInitialized) (event.Subscription, error) {

	logs, sub, err := _Pessimisticglobalexitroot.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticglobalexitrootInitialized)
				if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) ParseInitialized(log types.Log) (*PessimisticglobalexitrootInitialized, error) {
	event := new(PessimisticglobalexitrootInitialized)
	if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticglobalexitrootRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleAdminChangedIterator struct {
	Event *PessimisticglobalexitrootRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PessimisticglobalexitrootRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticglobalexitrootRoleAdminChanged)
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
		it.Event = new(PessimisticglobalexitrootRoleAdminChanged)
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
func (it *PessimisticglobalexitrootRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticglobalexitrootRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticglobalexitrootRoleAdminChanged represents a RoleAdminChanged event raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PessimisticglobalexitrootRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootRoleAdminChangedIterator{contract: _Pessimisticglobalexitroot.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PessimisticglobalexitrootRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticglobalexitrootRoleAdminChanged)
				if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) ParseRoleAdminChanged(log types.Log) (*PessimisticglobalexitrootRoleAdminChanged, error) {
	event := new(PessimisticglobalexitrootRoleAdminChanged)
	if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticglobalexitrootRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleGrantedIterator struct {
	Event *PessimisticglobalexitrootRoleGranted // Event containing the contract specifics and raw log

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
func (it *PessimisticglobalexitrootRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticglobalexitrootRoleGranted)
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
		it.Event = new(PessimisticglobalexitrootRoleGranted)
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
func (it *PessimisticglobalexitrootRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticglobalexitrootRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticglobalexitrootRoleGranted represents a RoleGranted event raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PessimisticglobalexitrootRoleGrantedIterator, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootRoleGrantedIterator{contract: _Pessimisticglobalexitroot.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PessimisticglobalexitrootRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticglobalexitrootRoleGranted)
				if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) ParseRoleGranted(log types.Log) (*PessimisticglobalexitrootRoleGranted, error) {
	event := new(PessimisticglobalexitrootRoleGranted)
	if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PessimisticglobalexitrootRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleRevokedIterator struct {
	Event *PessimisticglobalexitrootRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PessimisticglobalexitrootRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PessimisticglobalexitrootRoleRevoked)
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
		it.Event = new(PessimisticglobalexitrootRoleRevoked)
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
func (it *PessimisticglobalexitrootRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PessimisticglobalexitrootRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PessimisticglobalexitrootRoleRevoked represents a RoleRevoked event raised by the Pessimisticglobalexitroot contract.
type PessimisticglobalexitrootRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PessimisticglobalexitrootRoleRevokedIterator, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PessimisticglobalexitrootRoleRevokedIterator{contract: _Pessimisticglobalexitroot.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PessimisticglobalexitrootRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pessimisticglobalexitroot.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PessimisticglobalexitrootRoleRevoked)
				if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Pessimisticglobalexitroot *PessimisticglobalexitrootFilterer) ParseRoleRevoked(log types.Log) (*PessimisticglobalexitrootRoleRevoked, error) {
	event := new(PessimisticglobalexitrootRoleRevoked)
	if err := _Pessimisticglobalexitroot.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

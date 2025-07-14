// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggoraclemanager

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

// IaggoraclemanagerMetaData contains all meta data concerning the Iaggoraclemanager contract.
var IaggoraclemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProposedGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WasNotOracleMember\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"AddAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"consolidatedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"ConsolidatedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ProposedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleMemberRemoved\",\"type\":\"address\"}],\"name\":\"RemoveAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"UpdateQuorum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_PROPOSED_GER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"addOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"addressToLastProposedGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"aggOracleMembers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"getAggOracleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggOracleMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAggOracleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManagerL2Sovereign\",\"outputs\":[{\"internalType\":\"contractGlobalExitRootManagerL2SovereignChain\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_aggOracleMembers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_quorum\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"proposeGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMemberAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oracleMemberIndex\",\"type\":\"uint256\"}],\"name\":\"removeOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IaggoraclemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggoraclemanagerMetaData.ABI instead.
var IaggoraclemanagerABI = IaggoraclemanagerMetaData.ABI

// Iaggoraclemanager is an auto generated Go binding around an Ethereum contract.
type Iaggoraclemanager struct {
	IaggoraclemanagerCaller     // Read-only binding to the contract
	IaggoraclemanagerTransactor // Write-only binding to the contract
	IaggoraclemanagerFilterer   // Log filterer for contract events
}

// IaggoraclemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggoraclemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggoraclemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggoraclemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggoraclemanagerSession struct {
	Contract     *Iaggoraclemanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IaggoraclemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggoraclemanagerCallerSession struct {
	Contract *IaggoraclemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IaggoraclemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggoraclemanagerTransactorSession struct {
	Contract     *IaggoraclemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IaggoraclemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggoraclemanagerRaw struct {
	Contract *Iaggoraclemanager // Generic contract binding to access the raw methods on
}

// IaggoraclemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggoraclemanagerCallerRaw struct {
	Contract *IaggoraclemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IaggoraclemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggoraclemanagerTransactorRaw struct {
	Contract *IaggoraclemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggoraclemanager creates a new instance of Iaggoraclemanager, bound to a specific deployed contract.
func NewIaggoraclemanager(address common.Address, backend bind.ContractBackend) (*Iaggoraclemanager, error) {
	contract, err := bindIaggoraclemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggoraclemanager{IaggoraclemanagerCaller: IaggoraclemanagerCaller{contract: contract}, IaggoraclemanagerTransactor: IaggoraclemanagerTransactor{contract: contract}, IaggoraclemanagerFilterer: IaggoraclemanagerFilterer{contract: contract}}, nil
}

// NewIaggoraclemanagerCaller creates a new read-only instance of Iaggoraclemanager, bound to a specific deployed contract.
func NewIaggoraclemanagerCaller(address common.Address, caller bind.ContractCaller) (*IaggoraclemanagerCaller, error) {
	contract, err := bindIaggoraclemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerCaller{contract: contract}, nil
}

// NewIaggoraclemanagerTransactor creates a new write-only instance of Iaggoraclemanager, bound to a specific deployed contract.
func NewIaggoraclemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggoraclemanagerTransactor, error) {
	contract, err := bindIaggoraclemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerTransactor{contract: contract}, nil
}

// NewIaggoraclemanagerFilterer creates a new log filterer instance of Iaggoraclemanager, bound to a specific deployed contract.
func NewIaggoraclemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggoraclemanagerFilterer, error) {
	contract, err := bindIaggoraclemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerFilterer{contract: contract}, nil
}

// bindIaggoraclemanager binds a generic wrapper to an already deployed contract.
func bindIaggoraclemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggoraclemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclemanager *IaggoraclemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclemanager.Contract.IaggoraclemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclemanager *IaggoraclemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.IaggoraclemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclemanager *IaggoraclemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.IaggoraclemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclemanager *IaggoraclemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclemanager *IaggoraclemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclemanager *IaggoraclemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.contract.Transact(opts, method, params...)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) INITIALPROPOSEDGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "INITIAL_PROPOSED_GER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Iaggoraclemanager.Contract.INITIALPROPOSEDGER(&_Iaggoraclemanager.CallOpts)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Iaggoraclemanager.Contract.INITIALPROPOSEDGER(&_Iaggoraclemanager.CallOpts)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) AddressToLastProposedGER(opts *bind.CallOpts, oracleMember common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "addressToLastProposedGER", oracleMember)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerSession) AddressToLastProposedGER(oracleMember common.Address) ([32]byte, error) {
	return _Iaggoraclemanager.Contract.AddressToLastProposedGER(&_Iaggoraclemanager.CallOpts, oracleMember)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) AddressToLastProposedGER(oracleMember common.Address) ([32]byte, error) {
	return _Iaggoraclemanager.Contract.AddressToLastProposedGER(&_Iaggoraclemanager.CallOpts, oracleMember)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) AggOracleMembers(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "aggOracleMembers", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerSession) AggOracleMembers(index *big.Int) (common.Address, error) {
	return _Iaggoraclemanager.Contract.AggOracleMembers(&_Iaggoraclemanager.CallOpts, index)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) AggOracleMembers(index *big.Int) (common.Address, error) {
	return _Iaggoraclemanager.Contract.AggOracleMembers(&_Iaggoraclemanager.CallOpts, index)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) GetAggOracleMemberIndex(opts *bind.CallOpts, oracleMember common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "getAggOracleMemberIndex", oracleMember)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Iaggoraclemanager.Contract.GetAggOracleMemberIndex(&_Iaggoraclemanager.CallOpts, oracleMember)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Iaggoraclemanager.Contract.GetAggOracleMemberIndex(&_Iaggoraclemanager.CallOpts, oracleMember)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) GetAggOracleMembersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "getAggOracleMembersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Iaggoraclemanager.Contract.GetAggOracleMembersCount(&_Iaggoraclemanager.CallOpts)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Iaggoraclemanager.Contract.GetAggOracleMembersCount(&_Iaggoraclemanager.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclemanager *IaggoraclemanagerCaller) GetAllAggOracleMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "getAllAggOracleMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclemanager *IaggoraclemanagerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Iaggoraclemanager.Contract.GetAllAggOracleMembers(&_Iaggoraclemanager.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Iaggoraclemanager.Contract.GetAllAggOracleMembers(&_Iaggoraclemanager.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) GlobalExitRootManagerL2Sovereign(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "globalExitRootManagerL2Sovereign")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Iaggoraclemanager.Contract.GlobalExitRootManagerL2Sovereign(&_Iaggoraclemanager.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Iaggoraclemanager.Contract.GlobalExitRootManagerL2Sovereign(&_Iaggoraclemanager.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclemanager *IaggoraclemanagerCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Iaggoraclemanager.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclemanager *IaggoraclemanagerSession) Quorum() (uint64, error) {
	return _Iaggoraclemanager.Contract.Quorum(&_Iaggoraclemanager.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclemanager *IaggoraclemanagerCallerSession) Quorum() (uint64, error) {
	return _Iaggoraclemanager.Contract.Quorum(&_Iaggoraclemanager.CallOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.AcceptGlobalExitRootUpdater(&_Iaggoraclemanager.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.AcceptGlobalExitRootUpdater(&_Iaggoraclemanager.TransactOpts)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) AddOracleMember(opts *bind.TransactOpts, newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "addOracleMember", newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.AddOracleMember(&_Iaggoraclemanager.TransactOpts, newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.AddOracleMember(&_Iaggoraclemanager.TransactOpts, newOracleMember)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "initialize", _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.Initialize(&_Iaggoraclemanager.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.Initialize(&_Iaggoraclemanager.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) ProposeGlobalExitRoot(opts *bind.TransactOpts, proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "proposeGlobalExitRoot", proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.ProposeGlobalExitRoot(&_Iaggoraclemanager.TransactOpts, proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.ProposeGlobalExitRoot(&_Iaggoraclemanager.TransactOpts, proposedGlobalExitRoot)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) RemoveOracleMember(opts *bind.TransactOpts, oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "removeOracleMember", oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.RemoveOracleMember(&_Iaggoraclemanager.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.RemoveOracleMember(&_Iaggoraclemanager.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.TransferGlobalExitRootUpdater(&_Iaggoraclemanager.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.TransferGlobalExitRootUpdater(&_Iaggoraclemanager.TransactOpts, _newGlobalExitRootUpdater)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactor) UpdateQuorum(opts *bind.TransactOpts, newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.contract.Transact(opts, "updateQuorum", newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.UpdateQuorum(&_Iaggoraclemanager.TransactOpts, newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclemanager *IaggoraclemanagerTransactorSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclemanager.Contract.UpdateQuorum(&_Iaggoraclemanager.TransactOpts, newQuorum)
}

// IaggoraclemanagerAddAggOracleMemberIterator is returned from FilterAddAggOracleMember and is used to iterate over the raw logs and unpacked data for AddAggOracleMember events raised by the Iaggoraclemanager contract.
type IaggoraclemanagerAddAggOracleMemberIterator struct {
	Event *IaggoraclemanagerAddAggOracleMember // Event containing the contract specifics and raw log

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
func (it *IaggoraclemanagerAddAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclemanagerAddAggOracleMember)
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
		it.Event = new(IaggoraclemanagerAddAggOracleMember)
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
func (it *IaggoraclemanagerAddAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclemanagerAddAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclemanagerAddAggOracleMember represents a AddAggOracleMember event raised by the Iaggoraclemanager contract.
type IaggoraclemanagerAddAggOracleMember struct {
	NewOracleMember common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggOracleMember is a free log retrieval operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) FilterAddAggOracleMember(opts *bind.FilterOpts) (*IaggoraclemanagerAddAggOracleMemberIterator, error) {

	logs, sub, err := _Iaggoraclemanager.contract.FilterLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerAddAggOracleMemberIterator{contract: _Iaggoraclemanager.contract, event: "AddAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchAddAggOracleMember is a free log subscription operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) WatchAddAggOracleMember(opts *bind.WatchOpts, sink chan<- *IaggoraclemanagerAddAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclemanager.contract.WatchLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclemanagerAddAggOracleMember)
				if err := _Iaggoraclemanager.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
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

// ParseAddAggOracleMember is a log parse operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) ParseAddAggOracleMember(log types.Log) (*IaggoraclemanagerAddAggOracleMember, error) {
	event := new(IaggoraclemanagerAddAggOracleMember)
	if err := _Iaggoraclemanager.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclemanagerConsolidatedGlobalExitRootIterator is returned from FilterConsolidatedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ConsolidatedGlobalExitRoot events raised by the Iaggoraclemanager contract.
type IaggoraclemanagerConsolidatedGlobalExitRootIterator struct {
	Event *IaggoraclemanagerConsolidatedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *IaggoraclemanagerConsolidatedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclemanagerConsolidatedGlobalExitRoot)
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
		it.Event = new(IaggoraclemanagerConsolidatedGlobalExitRoot)
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
func (it *IaggoraclemanagerConsolidatedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclemanagerConsolidatedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclemanagerConsolidatedGlobalExitRoot represents a ConsolidatedGlobalExitRoot event raised by the Iaggoraclemanager contract.
type IaggoraclemanagerConsolidatedGlobalExitRoot struct {
	ConsolidatedGlobalExitRoot [32]byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterConsolidatedGlobalExitRoot is a free log retrieval operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) FilterConsolidatedGlobalExitRoot(opts *bind.FilterOpts) (*IaggoraclemanagerConsolidatedGlobalExitRootIterator, error) {

	logs, sub, err := _Iaggoraclemanager.contract.FilterLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerConsolidatedGlobalExitRootIterator{contract: _Iaggoraclemanager.contract, event: "ConsolidatedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchConsolidatedGlobalExitRoot is a free log subscription operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) WatchConsolidatedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *IaggoraclemanagerConsolidatedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclemanager.contract.WatchLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclemanagerConsolidatedGlobalExitRoot)
				if err := _Iaggoraclemanager.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
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

// ParseConsolidatedGlobalExitRoot is a log parse operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) ParseConsolidatedGlobalExitRoot(log types.Log) (*IaggoraclemanagerConsolidatedGlobalExitRoot, error) {
	event := new(IaggoraclemanagerConsolidatedGlobalExitRoot)
	if err := _Iaggoraclemanager.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclemanagerProposedGlobalExitRootIterator is returned from FilterProposedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ProposedGlobalExitRoot events raised by the Iaggoraclemanager contract.
type IaggoraclemanagerProposedGlobalExitRootIterator struct {
	Event *IaggoraclemanagerProposedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *IaggoraclemanagerProposedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclemanagerProposedGlobalExitRoot)
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
		it.Event = new(IaggoraclemanagerProposedGlobalExitRoot)
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
func (it *IaggoraclemanagerProposedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclemanagerProposedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclemanagerProposedGlobalExitRoot represents a ProposedGlobalExitRoot event raised by the Iaggoraclemanager contract.
type IaggoraclemanagerProposedGlobalExitRoot struct {
	ProposedGlobalExitRoot [32]byte
	Proposer               common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProposedGlobalExitRoot is a free log retrieval operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) FilterProposedGlobalExitRoot(opts *bind.FilterOpts) (*IaggoraclemanagerProposedGlobalExitRootIterator, error) {

	logs, sub, err := _Iaggoraclemanager.contract.FilterLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerProposedGlobalExitRootIterator{contract: _Iaggoraclemanager.contract, event: "ProposedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchProposedGlobalExitRoot is a free log subscription operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) WatchProposedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *IaggoraclemanagerProposedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclemanager.contract.WatchLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclemanagerProposedGlobalExitRoot)
				if err := _Iaggoraclemanager.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
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

// ParseProposedGlobalExitRoot is a log parse operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) ParseProposedGlobalExitRoot(log types.Log) (*IaggoraclemanagerProposedGlobalExitRoot, error) {
	event := new(IaggoraclemanagerProposedGlobalExitRoot)
	if err := _Iaggoraclemanager.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclemanagerRemoveAggOracleMemberIterator is returned from FilterRemoveAggOracleMember and is used to iterate over the raw logs and unpacked data for RemoveAggOracleMember events raised by the Iaggoraclemanager contract.
type IaggoraclemanagerRemoveAggOracleMemberIterator struct {
	Event *IaggoraclemanagerRemoveAggOracleMember // Event containing the contract specifics and raw log

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
func (it *IaggoraclemanagerRemoveAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclemanagerRemoveAggOracleMember)
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
		it.Event = new(IaggoraclemanagerRemoveAggOracleMember)
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
func (it *IaggoraclemanagerRemoveAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclemanagerRemoveAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclemanagerRemoveAggOracleMember represents a RemoveAggOracleMember event raised by the Iaggoraclemanager contract.
type IaggoraclemanagerRemoveAggOracleMember struct {
	OracleMemberRemoved common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemoveAggOracleMember is a free log retrieval operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) FilterRemoveAggOracleMember(opts *bind.FilterOpts) (*IaggoraclemanagerRemoveAggOracleMemberIterator, error) {

	logs, sub, err := _Iaggoraclemanager.contract.FilterLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerRemoveAggOracleMemberIterator{contract: _Iaggoraclemanager.contract, event: "RemoveAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchRemoveAggOracleMember is a free log subscription operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) WatchRemoveAggOracleMember(opts *bind.WatchOpts, sink chan<- *IaggoraclemanagerRemoveAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclemanager.contract.WatchLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclemanagerRemoveAggOracleMember)
				if err := _Iaggoraclemanager.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
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

// ParseRemoveAggOracleMember is a log parse operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) ParseRemoveAggOracleMember(log types.Log) (*IaggoraclemanagerRemoveAggOracleMember, error) {
	event := new(IaggoraclemanagerRemoveAggOracleMember)
	if err := _Iaggoraclemanager.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclemanagerUpdateQuorumIterator is returned from FilterUpdateQuorum and is used to iterate over the raw logs and unpacked data for UpdateQuorum events raised by the Iaggoraclemanager contract.
type IaggoraclemanagerUpdateQuorumIterator struct {
	Event *IaggoraclemanagerUpdateQuorum // Event containing the contract specifics and raw log

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
func (it *IaggoraclemanagerUpdateQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclemanagerUpdateQuorum)
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
		it.Event = new(IaggoraclemanagerUpdateQuorum)
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
func (it *IaggoraclemanagerUpdateQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclemanagerUpdateQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclemanagerUpdateQuorum represents a UpdateQuorum event raised by the Iaggoraclemanager contract.
type IaggoraclemanagerUpdateQuorum struct {
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuorum is a free log retrieval operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) FilterUpdateQuorum(opts *bind.FilterOpts) (*IaggoraclemanagerUpdateQuorumIterator, error) {

	logs, sub, err := _Iaggoraclemanager.contract.FilterLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return &IaggoraclemanagerUpdateQuorumIterator{contract: _Iaggoraclemanager.contract, event: "UpdateQuorum", logs: logs, sub: sub}, nil
}

// WatchUpdateQuorum is a free log subscription operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) WatchUpdateQuorum(opts *bind.WatchOpts, sink chan<- *IaggoraclemanagerUpdateQuorum) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclemanager.contract.WatchLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclemanagerUpdateQuorum)
				if err := _Iaggoraclemanager.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
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

// ParseUpdateQuorum is a log parse operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Iaggoraclemanager *IaggoraclemanagerFilterer) ParseUpdateQuorum(log types.Log) (*IaggoraclemanagerUpdateQuorum, error) {
	event := new(IaggoraclemanagerUpdateQuorum)
	if err := _Iaggoraclemanager.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

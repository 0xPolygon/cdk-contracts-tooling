// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iaggoraclecommittee

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

// IaggoraclecommitteeMetaData contains all meta data concerning the Iaggoraclecommittee contract.
var IaggoraclecommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyVotedForThisGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProposedGER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOracleMember\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OracleMemberNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeGreaterThanAggOracleMembers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuorumNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WasNotOracleMember\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"AddAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"consolidatedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"ConsolidatedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"ProposedGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracleMemberRemoved\",\"type\":\"address\"}],\"name\":\"RemoveAggOracleMember\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"UpdateQuorum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INITIAL_PROPOSED_GER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracleMember\",\"type\":\"address\"}],\"name\":\"addOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"addressToLastProposedGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"aggOracleMembers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"consolidateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMember\",\"type\":\"address\"}],\"name\":\"getAggOracleMemberIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggOracleMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAggOracleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManagerL2Sovereign\",\"outputs\":[{\"internalType\":\"contractIAgglayerGERL2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_aggOracleMembers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_quorum\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"proposeGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleMemberAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oracleMemberIndex\",\"type\":\"uint256\"}],\"name\":\"removeOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newQuorum\",\"type\":\"uint64\"}],\"name\":\"updateQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IaggoraclecommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use IaggoraclecommitteeMetaData.ABI instead.
var IaggoraclecommitteeABI = IaggoraclecommitteeMetaData.ABI

// Iaggoraclecommittee is an auto generated Go binding around an Ethereum contract.
type Iaggoraclecommittee struct {
	IaggoraclecommitteeCaller     // Read-only binding to the contract
	IaggoraclecommitteeTransactor // Write-only binding to the contract
	IaggoraclecommitteeFilterer   // Log filterer for contract events
}

// IaggoraclecommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IaggoraclecommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IaggoraclecommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IaggoraclecommitteeSession struct {
	Contract     *Iaggoraclecommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IaggoraclecommitteeCallerSession struct {
	Contract *IaggoraclecommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IaggoraclecommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IaggoraclecommitteeTransactorSession struct {
	Contract     *IaggoraclecommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IaggoraclecommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IaggoraclecommitteeRaw struct {
	Contract *Iaggoraclecommittee // Generic contract binding to access the raw methods on
}

// IaggoraclecommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IaggoraclecommitteeCallerRaw struct {
	Contract *IaggoraclecommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// IaggoraclecommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IaggoraclecommitteeTransactorRaw struct {
	Contract *IaggoraclecommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIaggoraclecommittee creates a new instance of Iaggoraclecommittee, bound to a specific deployed contract.
func NewIaggoraclecommittee(address common.Address, backend bind.ContractBackend) (*Iaggoraclecommittee, error) {
	contract, err := bindIaggoraclecommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iaggoraclecommittee{IaggoraclecommitteeCaller: IaggoraclecommitteeCaller{contract: contract}, IaggoraclecommitteeTransactor: IaggoraclecommitteeTransactor{contract: contract}, IaggoraclecommitteeFilterer: IaggoraclecommitteeFilterer{contract: contract}}, nil
}

// NewIaggoraclecommitteeCaller creates a new read-only instance of Iaggoraclecommittee, bound to a specific deployed contract.
func NewIaggoraclecommitteeCaller(address common.Address, caller bind.ContractCaller) (*IaggoraclecommitteeCaller, error) {
	contract, err := bindIaggoraclecommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeCaller{contract: contract}, nil
}

// NewIaggoraclecommitteeTransactor creates a new write-only instance of Iaggoraclecommittee, bound to a specific deployed contract.
func NewIaggoraclecommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*IaggoraclecommitteeTransactor, error) {
	contract, err := bindIaggoraclecommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeTransactor{contract: contract}, nil
}

// NewIaggoraclecommitteeFilterer creates a new log filterer instance of Iaggoraclecommittee, bound to a specific deployed contract.
func NewIaggoraclecommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*IaggoraclecommitteeFilterer, error) {
	contract, err := bindIaggoraclecommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeFilterer{contract: contract}, nil
}

// bindIaggoraclecommittee binds a generic wrapper to an already deployed contract.
func bindIaggoraclecommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IaggoraclecommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommittee *IaggoraclecommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommittee.Contract.IaggoraclecommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommittee *IaggoraclecommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.IaggoraclecommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommittee *IaggoraclecommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.IaggoraclecommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iaggoraclecommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.contract.Transact(opts, method, params...)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) INITIALPROPOSEDGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "INITIAL_PROPOSED_GER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Iaggoraclecommittee.Contract.INITIALPROPOSEDGER(&_Iaggoraclecommittee.CallOpts)
}

// INITIALPROPOSEDGER is a free data retrieval call binding the contract method 0x512bab0d.
//
// Solidity: function INITIAL_PROPOSED_GER() view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) INITIALPROPOSEDGER() ([32]byte, error) {
	return _Iaggoraclecommittee.Contract.INITIALPROPOSEDGER(&_Iaggoraclecommittee.CallOpts)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) AddressToLastProposedGER(opts *bind.CallOpts, oracleMember common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "addressToLastProposedGER", oracleMember)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) AddressToLastProposedGER(oracleMember common.Address) ([32]byte, error) {
	return _Iaggoraclecommittee.Contract.AddressToLastProposedGER(&_Iaggoraclecommittee.CallOpts, oracleMember)
}

// AddressToLastProposedGER is a free data retrieval call binding the contract method 0x38fd8a03.
//
// Solidity: function addressToLastProposedGER(address oracleMember) view returns(bytes32)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) AddressToLastProposedGER(oracleMember common.Address) ([32]byte, error) {
	return _Iaggoraclecommittee.Contract.AddressToLastProposedGER(&_Iaggoraclecommittee.CallOpts, oracleMember)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) AggOracleMembers(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "aggOracleMembers", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) AggOracleMembers(index *big.Int) (common.Address, error) {
	return _Iaggoraclecommittee.Contract.AggOracleMembers(&_Iaggoraclecommittee.CallOpts, index)
}

// AggOracleMembers is a free data retrieval call binding the contract method 0x1ffbaa74.
//
// Solidity: function aggOracleMembers(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) AggOracleMembers(index *big.Int) (common.Address, error) {
	return _Iaggoraclecommittee.Contract.AggOracleMembers(&_Iaggoraclecommittee.CallOpts, index)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetAggOracleMemberIndex(opts *bind.CallOpts, oracleMember common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getAggOracleMemberIndex", oracleMember)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetAggOracleMemberIndex(&_Iaggoraclecommittee.CallOpts, oracleMember)
}

// GetAggOracleMemberIndex is a free data retrieval call binding the contract method 0x72312af8.
//
// Solidity: function getAggOracleMemberIndex(address oracleMember) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetAggOracleMemberIndex(oracleMember common.Address) (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetAggOracleMemberIndex(&_Iaggoraclecommittee.CallOpts, oracleMember)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetAggOracleMembersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getAggOracleMembersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetAggOracleMembersCount(&_Iaggoraclecommittee.CallOpts)
}

// GetAggOracleMembersCount is a free data retrieval call binding the contract method 0x017e37d6.
//
// Solidity: function getAggOracleMembersCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetAggOracleMembersCount() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetAggOracleMembersCount(&_Iaggoraclecommittee.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetAllAggOracleMembers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getAllAggOracleMembers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Iaggoraclecommittee.Contract.GetAllAggOracleMembers(&_Iaggoraclecommittee.CallOpts)
}

// GetAllAggOracleMembers is a free data retrieval call binding the contract method 0x703d8d42.
//
// Solidity: function getAllAggOracleMembers() view returns(address[])
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetAllAggOracleMembers() ([]common.Address, error) {
	return _Iaggoraclecommittee.Contract.GetAllAggOracleMembers(&_Iaggoraclecommittee.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GlobalExitRootManagerL2Sovereign(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "globalExitRootManagerL2Sovereign")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.GlobalExitRootManagerL2Sovereign(&_Iaggoraclecommittee.CallOpts)
}

// GlobalExitRootManagerL2Sovereign is a free data retrieval call binding the contract method 0xa7a90e2a.
//
// Solidity: function globalExitRootManagerL2Sovereign() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GlobalExitRootManagerL2Sovereign() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.GlobalExitRootManagerL2Sovereign(&_Iaggoraclecommittee.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) Quorum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Quorum() (uint64, error) {
	return _Iaggoraclecommittee.Contract.Quorum(&_Iaggoraclecommittee.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint64)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) Quorum() (uint64, error) {
	return _Iaggoraclecommittee.Contract.Quorum(&_Iaggoraclecommittee.CallOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.AcceptGlobalExitRootUpdater(&_Iaggoraclecommittee.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.AcceptGlobalExitRootUpdater(&_Iaggoraclecommittee.TransactOpts)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) AddOracleMember(opts *bind.TransactOpts, newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "addOracleMember", newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.AddOracleMember(&_Iaggoraclecommittee.TransactOpts, newOracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xb164e437.
//
// Solidity: function addOracleMember(address newOracleMember) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) AddOracleMember(newOracleMember common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.AddOracleMember(&_Iaggoraclecommittee.TransactOpts, newOracleMember)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) ConsolidateGlobalExitRoot(opts *bind.TransactOpts, globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "consolidateGlobalExitRoot", globalExitRoot)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) ConsolidateGlobalExitRoot(globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.ConsolidateGlobalExitRoot(&_Iaggoraclecommittee.TransactOpts, globalExitRoot)
}

// ConsolidateGlobalExitRoot is a paid mutator transaction binding the contract method 0x9915c0d2.
//
// Solidity: function consolidateGlobalExitRoot(bytes32 globalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) ConsolidateGlobalExitRoot(globalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.ConsolidateGlobalExitRoot(&_Iaggoraclecommittee.TransactOpts, globalExitRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "initialize", _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.Initialize(&_Iaggoraclecommittee.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x145d1ed6.
//
// Solidity: function initialize(address _owner, address[] _aggOracleMembers, uint64 _quorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) Initialize(_owner common.Address, _aggOracleMembers []common.Address, _quorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.Initialize(&_Iaggoraclecommittee.TransactOpts, _owner, _aggOracleMembers, _quorum)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) ProposeGlobalExitRoot(opts *bind.TransactOpts, proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "proposeGlobalExitRoot", proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.ProposeGlobalExitRoot(&_Iaggoraclecommittee.TransactOpts, proposedGlobalExitRoot)
}

// ProposeGlobalExitRoot is a paid mutator transaction binding the contract method 0x542b56b3.
//
// Solidity: function proposeGlobalExitRoot(bytes32 proposedGlobalExitRoot) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) ProposeGlobalExitRoot(proposedGlobalExitRoot [32]byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.ProposeGlobalExitRoot(&_Iaggoraclecommittee.TransactOpts, proposedGlobalExitRoot)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) RemoveOracleMember(opts *bind.TransactOpts, oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "removeOracleMember", oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.RemoveOracleMember(&_Iaggoraclecommittee.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x53985e5a.
//
// Solidity: function removeOracleMember(address oracleMemberAddress, uint256 oracleMemberIndex) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) RemoveOracleMember(oracleMemberAddress common.Address, oracleMemberIndex *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.RemoveOracleMember(&_Iaggoraclecommittee.TransactOpts, oracleMemberAddress, oracleMemberIndex)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.TransferGlobalExitRootUpdater(&_Iaggoraclecommittee.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.TransferGlobalExitRootUpdater(&_Iaggoraclecommittee.TransactOpts, _newGlobalExitRootUpdater)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) UpdateQuorum(opts *bind.TransactOpts, newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "updateQuorum", newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateQuorum(&_Iaggoraclecommittee.TransactOpts, newQuorum)
}

// UpdateQuorum is a paid mutator transaction binding the contract method 0x29218b61.
//
// Solidity: function updateQuorum(uint64 newQuorum) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) UpdateQuorum(newQuorum uint64) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateQuorum(&_Iaggoraclecommittee.TransactOpts, newQuorum)
}

// IaggoraclecommitteeAddAggOracleMemberIterator is returned from FilterAddAggOracleMember and is used to iterate over the raw logs and unpacked data for AddAggOracleMember events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeAddAggOracleMemberIterator struct {
	Event *IaggoraclecommitteeAddAggOracleMember // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeAddAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeAddAggOracleMember)
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
		it.Event = new(IaggoraclecommitteeAddAggOracleMember)
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
func (it *IaggoraclecommitteeAddAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeAddAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeAddAggOracleMember represents a AddAggOracleMember event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeAddAggOracleMember struct {
	NewOracleMember common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggOracleMember is a free log retrieval operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterAddAggOracleMember(opts *bind.FilterOpts) (*IaggoraclecommitteeAddAggOracleMemberIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeAddAggOracleMemberIterator{contract: _Iaggoraclecommittee.contract, event: "AddAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchAddAggOracleMember is a free log subscription operation binding the contract event 0x59f8c2dd0bdc5787ef574566b9edcd7324a62b79ac952930f02679a1b25bbf6f.
//
// Solidity: event AddAggOracleMember(address newOracleMember)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchAddAggOracleMember(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeAddAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "AddAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeAddAggOracleMember)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
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
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseAddAggOracleMember(log types.Log) (*IaggoraclecommitteeAddAggOracleMember, error) {
	event := new(IaggoraclecommitteeAddAggOracleMember)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "AddAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeConsolidatedGlobalExitRootIterator is returned from FilterConsolidatedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ConsolidatedGlobalExitRoot events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeConsolidatedGlobalExitRootIterator struct {
	Event *IaggoraclecommitteeConsolidatedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeConsolidatedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeConsolidatedGlobalExitRoot)
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
		it.Event = new(IaggoraclecommitteeConsolidatedGlobalExitRoot)
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
func (it *IaggoraclecommitteeConsolidatedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeConsolidatedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeConsolidatedGlobalExitRoot represents a ConsolidatedGlobalExitRoot event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeConsolidatedGlobalExitRoot struct {
	ConsolidatedGlobalExitRoot [32]byte
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterConsolidatedGlobalExitRoot is a free log retrieval operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterConsolidatedGlobalExitRoot(opts *bind.FilterOpts) (*IaggoraclecommitteeConsolidatedGlobalExitRootIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeConsolidatedGlobalExitRootIterator{contract: _Iaggoraclecommittee.contract, event: "ConsolidatedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchConsolidatedGlobalExitRoot is a free log subscription operation binding the contract event 0x1676a1944900a9899dac554c2da10ba62637fe8d08375ffc3fb8494c92918d98.
//
// Solidity: event ConsolidatedGlobalExitRoot(bytes32 consolidatedGlobalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchConsolidatedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeConsolidatedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "ConsolidatedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeConsolidatedGlobalExitRoot)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
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
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseConsolidatedGlobalExitRoot(log types.Log) (*IaggoraclecommitteeConsolidatedGlobalExitRoot, error) {
	event := new(IaggoraclecommitteeConsolidatedGlobalExitRoot)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ConsolidatedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeProposedGlobalExitRootIterator is returned from FilterProposedGlobalExitRoot and is used to iterate over the raw logs and unpacked data for ProposedGlobalExitRoot events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeProposedGlobalExitRootIterator struct {
	Event *IaggoraclecommitteeProposedGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeProposedGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeProposedGlobalExitRoot)
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
		it.Event = new(IaggoraclecommitteeProposedGlobalExitRoot)
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
func (it *IaggoraclecommitteeProposedGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeProposedGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeProposedGlobalExitRoot represents a ProposedGlobalExitRoot event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeProposedGlobalExitRoot struct {
	ProposedGlobalExitRoot [32]byte
	Proposer               common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProposedGlobalExitRoot is a free log retrieval operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterProposedGlobalExitRoot(opts *bind.FilterOpts) (*IaggoraclecommitteeProposedGlobalExitRootIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeProposedGlobalExitRootIterator{contract: _Iaggoraclecommittee.contract, event: "ProposedGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchProposedGlobalExitRoot is a free log subscription operation binding the contract event 0x9f298266555f7b22d45b8ea34a9c99cba101d1ccbc3ff6dc4c392b7c915dc3b5.
//
// Solidity: event ProposedGlobalExitRoot(bytes32 proposedGlobalExitRoot, address proposer)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchProposedGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeProposedGlobalExitRoot) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "ProposedGlobalExitRoot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeProposedGlobalExitRoot)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
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
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseProposedGlobalExitRoot(log types.Log) (*IaggoraclecommitteeProposedGlobalExitRoot, error) {
	event := new(IaggoraclecommitteeProposedGlobalExitRoot)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ProposedGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeRemoveAggOracleMemberIterator is returned from FilterRemoveAggOracleMember and is used to iterate over the raw logs and unpacked data for RemoveAggOracleMember events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeRemoveAggOracleMemberIterator struct {
	Event *IaggoraclecommitteeRemoveAggOracleMember // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeRemoveAggOracleMemberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeRemoveAggOracleMember)
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
		it.Event = new(IaggoraclecommitteeRemoveAggOracleMember)
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
func (it *IaggoraclecommitteeRemoveAggOracleMemberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeRemoveAggOracleMemberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeRemoveAggOracleMember represents a RemoveAggOracleMember event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeRemoveAggOracleMember struct {
	OracleMemberRemoved common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemoveAggOracleMember is a free log retrieval operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterRemoveAggOracleMember(opts *bind.FilterOpts) (*IaggoraclecommitteeRemoveAggOracleMemberIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeRemoveAggOracleMemberIterator{contract: _Iaggoraclecommittee.contract, event: "RemoveAggOracleMember", logs: logs, sub: sub}, nil
}

// WatchRemoveAggOracleMember is a free log subscription operation binding the contract event 0x396186e0feff55f3e52260aee4ec024318786c7f5b8b01639b1d42c24b59eafc.
//
// Solidity: event RemoveAggOracleMember(address oracleMemberRemoved)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchRemoveAggOracleMember(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeRemoveAggOracleMember) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "RemoveAggOracleMember")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeRemoveAggOracleMember)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
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
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseRemoveAggOracleMember(log types.Log) (*IaggoraclecommitteeRemoveAggOracleMember, error) {
	event := new(IaggoraclecommitteeRemoveAggOracleMember)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "RemoveAggOracleMember", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeUpdateQuorumIterator is returned from FilterUpdateQuorum and is used to iterate over the raw logs and unpacked data for UpdateQuorum events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeUpdateQuorumIterator struct {
	Event *IaggoraclecommitteeUpdateQuorum // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeUpdateQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeUpdateQuorum)
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
		it.Event = new(IaggoraclecommitteeUpdateQuorum)
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
func (it *IaggoraclecommitteeUpdateQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeUpdateQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeUpdateQuorum represents a UpdateQuorum event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeUpdateQuorum struct {
	NewQuorum uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateQuorum is a free log retrieval operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterUpdateQuorum(opts *bind.FilterOpts) (*IaggoraclecommitteeUpdateQuorumIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeUpdateQuorumIterator{contract: _Iaggoraclecommittee.contract, event: "UpdateQuorum", logs: logs, sub: sub}, nil
}

// WatchUpdateQuorum is a free log subscription operation binding the contract event 0xb600f3cf7f38a4b49bb0c75f722ef69f7e3e39ef3bb4aa8207fd86e724a23249.
//
// Solidity: event UpdateQuorum(uint64 newQuorum)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchUpdateQuorum(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeUpdateQuorum) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "UpdateQuorum")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeUpdateQuorum)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
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
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseUpdateQuorum(log types.Log) (*IaggoraclecommitteeUpdateQuorum, error) {
	event := new(IaggoraclecommitteeUpdateQuorum)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "UpdateQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

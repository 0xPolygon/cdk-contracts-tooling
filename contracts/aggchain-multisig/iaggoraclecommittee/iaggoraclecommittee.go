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

// IAggOracleCommitteeRemoveValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggOracleCommitteeRemoveValidatorInfo struct {
	Addr  common.Address
	Index *big.Int
}

// IAggOracleCommitteeValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IAggOracleCommitteeValidatorInfo struct {
	Addr common.Address
	Url  string
}

// IaggoraclecommitteeMetaData contains all meta data concerning the Iaggoraclecommittee contract.
var IaggoraclecommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AgglayerGERL2CannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERAlreadyInjected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GERCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignaturesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggOracleProposerCanCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RenounceOwnershipDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeGreaterThanValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeThisContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorsNotOrdered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"GERMarkedAsInjected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposer\",\"type\":\"address\"}],\"name\":\"UpdateAggOracleProposer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ValidatorsAndThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agglayerGERL2\",\"outputs\":[{\"internalType\":\"contractIAgglayerGERL2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newProposer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"_newValidators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"injectGER\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"injectedGERs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProposer\",\"type\":\"address\"}],\"name\":\"updateAggOracleProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structIAggOracleCommittee.RemoveValidatorInfo[]\",\"name\":\"_validatorsToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIAggOracleCommittee.ValidatorInfo[]\",\"name\":\"_validatorsToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateValidatorsAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"validatorInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"verifyGER\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) AgglayerGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "agglayerGERL2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) AgglayerGERL2() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.AgglayerGERL2(&_Iaggoraclecommittee.CallOpts)
}

// AgglayerGERL2 is a free data retrieval call binding the contract method 0x56eb7637.
//
// Solidity: function agglayerGERL2() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) AgglayerGERL2() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.AgglayerGERL2(&_Iaggoraclecommittee.CallOpts)
}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetValidatorIndex(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getValidatorIndex", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorIndex(&_Iaggoraclecommittee.CallOpts, _validator)
}

// GetValidatorIndex is a free data retrieval call binding the contract method 0x1fcc449e.
//
// Solidity: function getValidatorIndex(address _validator) view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetValidatorIndex(_validator common.Address) (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorIndex(&_Iaggoraclecommittee.CallOpts, _validator)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetValidatorsCount() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorsCount(&_Iaggoraclecommittee.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetValidatorsCount() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorsCount(&_Iaggoraclecommittee.CallOpts)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) GetValidatorsInfo(opts *bind.CallOpts) ([]IAggOracleCommitteeValidatorInfo, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "getValidatorsInfo")

	if err != nil {
		return *new([]IAggOracleCommitteeValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAggOracleCommitteeValidatorInfo)).(*[]IAggOracleCommitteeValidatorInfo)

	return out0, err

}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) GetValidatorsInfo() ([]IAggOracleCommitteeValidatorInfo, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorsInfo(&_Iaggoraclecommittee.CallOpts)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x8452230e.
//
// Solidity: function getValidatorsInfo() view returns((address,string)[])
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) GetValidatorsInfo() ([]IAggOracleCommitteeValidatorInfo, error) {
	return _Iaggoraclecommittee.Contract.GetValidatorsInfo(&_Iaggoraclecommittee.CallOpts)
}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 globalExitRoot) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) InjectedGERs(opts *bind.CallOpts, globalExitRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "injectedGERs", globalExitRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 globalExitRoot) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) InjectedGERs(globalExitRoot [32]byte) (bool, error) {
	return _Iaggoraclecommittee.Contract.InjectedGERs(&_Iaggoraclecommittee.CallOpts, globalExitRoot)
}

// InjectedGERs is a free data retrieval call binding the contract method 0xe905f142.
//
// Solidity: function injectedGERs(bytes32 globalExitRoot) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) InjectedGERs(globalExitRoot [32]byte) (bool, error) {
	return _Iaggoraclecommittee.Contract.InjectedGERs(&_Iaggoraclecommittee.CallOpts, globalExitRoot)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) IsValidator(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "isValidator", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) IsValidator(_validator common.Address) (bool, error) {
	return _Iaggoraclecommittee.Contract.IsValidator(&_Iaggoraclecommittee.CallOpts, _validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validator) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) IsValidator(_validator common.Address) (bool, error) {
	return _Iaggoraclecommittee.Contract.IsValidator(&_Iaggoraclecommittee.CallOpts, _validator)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Proposer() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.Proposer(&_Iaggoraclecommittee.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) Proposer() (common.Address, error) {
	return _Iaggoraclecommittee.Contract.Proposer(&_Iaggoraclecommittee.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Threshold() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.Threshold(&_Iaggoraclecommittee.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) Threshold() (*big.Int, error) {
	return _Iaggoraclecommittee.Contract.Threshold(&_Iaggoraclecommittee.CallOpts)
}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address validator) view returns(bool exists, string url)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) ValidatorInfo(opts *bind.CallOpts, validator common.Address) (struct {
	Exists bool
	Url    string
}, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "validatorInfo", validator)

	outstruct := new(struct {
		Exists bool
		Url    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Url = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address validator) view returns(bool exists, string url)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) ValidatorInfo(validator common.Address) (struct {
	Exists bool
	Url    string
}, error) {
	return _Iaggoraclecommittee.Contract.ValidatorInfo(&_Iaggoraclecommittee.CallOpts, validator)
}

// ValidatorInfo is a free data retrieval call binding the contract method 0x4f1811dd.
//
// Solidity: function validatorInfo(address validator) view returns(bool exists, string url)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) ValidatorInfo(validator common.Address) (struct {
	Exists bool
	Url    string
}, error) {
	return _Iaggoraclecommittee.Contract.ValidatorInfo(&_Iaggoraclecommittee.CallOpts, validator)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) Validators(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "validators", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Validators(index *big.Int) (common.Address, error) {
	return _Iaggoraclecommittee.Contract.Validators(&_Iaggoraclecommittee.CallOpts, index)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 index) view returns(address)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) Validators(index *big.Int) (common.Address, error) {
	return _Iaggoraclecommittee.Contract.Validators(&_Iaggoraclecommittee.CallOpts, index)
}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) VerifyGER(opts *bind.CallOpts, _globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "verifyGER", _globalExitRoot, _signatures)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) VerifyGER(_globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	return _Iaggoraclecommittee.Contract.VerifyGER(&_Iaggoraclecommittee.CallOpts, _globalExitRoot, _signatures)
}

// VerifyGER is a free data retrieval call binding the contract method 0x677f3950.
//
// Solidity: function verifyGER(bytes32 _globalExitRoot, bytes _signatures) view returns(bool)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) VerifyGER(_globalExitRoot [32]byte, _signatures []byte) (bool, error) {
	return _Iaggoraclecommittee.Contract.VerifyGER(&_Iaggoraclecommittee.CallOpts, _globalExitRoot, _signatures)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iaggoraclecommittee *IaggoraclecommitteeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Iaggoraclecommittee.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Version() (string, error) {
	return _Iaggoraclecommittee.Contract.Version(&_Iaggoraclecommittee.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Iaggoraclecommittee *IaggoraclecommitteeCallerSession) Version() (string, error) {
	return _Iaggoraclecommittee.Contract.Version(&_Iaggoraclecommittee.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "initialize", _owner, _newProposer, _newValidators, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) Initialize(_owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.Initialize(&_Iaggoraclecommittee.TransactOpts, _owner, _newProposer, _newValidators, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x98e04bf8.
//
// Solidity: function initialize(address _owner, address _newProposer, (address,string)[] _newValidators, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) Initialize(_owner common.Address, _newProposer common.Address, _newValidators []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.Initialize(&_Iaggoraclecommittee.TransactOpts, _owner, _newProposer, _newValidators, _newThreshold)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) InjectGER(opts *bind.TransactOpts, _globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "injectGER", _globalExitRoot, _signatures)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) InjectGER(_globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.InjectGER(&_Iaggoraclecommittee.TransactOpts, _globalExitRoot, _signatures)
}

// InjectGER is a paid mutator transaction binding the contract method 0xa31dee4a.
//
// Solidity: function injectGER(bytes32 _globalExitRoot, bytes _signatures) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) InjectGER(_globalExitRoot [32]byte, _signatures []byte) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.InjectGER(&_Iaggoraclecommittee.TransactOpts, _globalExitRoot, _signatures)
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

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) UpdateAggOracleProposer(opts *bind.TransactOpts, _newProposer common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "updateAggOracleProposer", _newProposer)
}

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) UpdateAggOracleProposer(_newProposer common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateAggOracleProposer(&_Iaggoraclecommittee.TransactOpts, _newProposer)
}

// UpdateAggOracleProposer is a paid mutator transaction binding the contract method 0xd6ad3f8f.
//
// Solidity: function updateAggOracleProposer(address _newProposer) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) UpdateAggOracleProposer(_newProposer common.Address) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateAggOracleProposer(&_Iaggoraclecommittee.TransactOpts, _newProposer)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactor) UpdateValidatorsAndThreshold(opts *bind.TransactOpts, _validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.contract.Transact(opts, "updateValidatorsAndThreshold", _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeSession) UpdateValidatorsAndThreshold(_validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateValidatorsAndThreshold(&_Iaggoraclecommittee.TransactOpts, _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// UpdateValidatorsAndThreshold is a paid mutator transaction binding the contract method 0xfedaa2a7.
//
// Solidity: function updateValidatorsAndThreshold((address,uint256)[] _validatorsToRemove, (address,string)[] _validatorsToAdd, uint256 _newThreshold) returns()
func (_Iaggoraclecommittee *IaggoraclecommitteeTransactorSession) UpdateValidatorsAndThreshold(_validatorsToRemove []IAggOracleCommitteeRemoveValidatorInfo, _validatorsToAdd []IAggOracleCommitteeValidatorInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Iaggoraclecommittee.Contract.UpdateValidatorsAndThreshold(&_Iaggoraclecommittee.TransactOpts, _validatorsToRemove, _validatorsToAdd, _newThreshold)
}

// IaggoraclecommitteeGERInjectedIterator is returned from FilterGERInjected and is used to iterate over the raw logs and unpacked data for GERInjected events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeGERInjectedIterator struct {
	Event *IaggoraclecommitteeGERInjected // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeGERInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeGERInjected)
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
		it.Event = new(IaggoraclecommitteeGERInjected)
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
func (it *IaggoraclecommitteeGERInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeGERInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeGERInjected represents a GERInjected event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeGERInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERInjected is a free log retrieval operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterGERInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*IaggoraclecommitteeGERInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeGERInjectedIterator{contract: _Iaggoraclecommittee.contract, event: "GERInjected", logs: logs, sub: sub}, nil
}

// WatchGERInjected is a free log subscription operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchGERInjected(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeGERInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "GERInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeGERInjected)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "GERInjected", log); err != nil {
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

// ParseGERInjected is a log parse operation binding the contract event 0x45b91a7a09575a05b9dd504dff707cdeb7eda5f9d514491caf5951551b020e1e.
//
// Solidity: event GERInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseGERInjected(log types.Log) (*IaggoraclecommitteeGERInjected, error) {
	event := new(IaggoraclecommitteeGERInjected)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "GERInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeGERMarkedAsInjectedIterator is returned from FilterGERMarkedAsInjected and is used to iterate over the raw logs and unpacked data for GERMarkedAsInjected events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeGERMarkedAsInjectedIterator struct {
	Event *IaggoraclecommitteeGERMarkedAsInjected // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeGERMarkedAsInjectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeGERMarkedAsInjected)
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
		it.Event = new(IaggoraclecommitteeGERMarkedAsInjected)
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
func (it *IaggoraclecommitteeGERMarkedAsInjectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeGERMarkedAsInjectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeGERMarkedAsInjected represents a GERMarkedAsInjected event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeGERMarkedAsInjected struct {
	GlobalExitRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGERMarkedAsInjected is a free log retrieval operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterGERMarkedAsInjected(opts *bind.FilterOpts, globalExitRoot [][32]byte) (*IaggoraclecommitteeGERMarkedAsInjectedIterator, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeGERMarkedAsInjectedIterator{contract: _Iaggoraclecommittee.contract, event: "GERMarkedAsInjected", logs: logs, sub: sub}, nil
}

// WatchGERMarkedAsInjected is a free log subscription operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchGERMarkedAsInjected(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeGERMarkedAsInjected, globalExitRoot [][32]byte) (event.Subscription, error) {

	var globalExitRootRule []interface{}
	for _, globalExitRootItem := range globalExitRoot {
		globalExitRootRule = append(globalExitRootRule, globalExitRootItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "GERMarkedAsInjected", globalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeGERMarkedAsInjected)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
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

// ParseGERMarkedAsInjected is a log parse operation binding the contract event 0xcafdbc8c7e71f2990afc9e787ff729a197666534c9a49959c1373c819b708a8f.
//
// Solidity: event GERMarkedAsInjected(bytes32 indexed globalExitRoot)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseGERMarkedAsInjected(log types.Log) (*IaggoraclecommitteeGERMarkedAsInjected, error) {
	event := new(IaggoraclecommitteeGERMarkedAsInjected)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "GERMarkedAsInjected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeUpdateAggOracleProposerIterator is returned from FilterUpdateAggOracleProposer and is used to iterate over the raw logs and unpacked data for UpdateAggOracleProposer events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeUpdateAggOracleProposerIterator struct {
	Event *IaggoraclecommitteeUpdateAggOracleProposer // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeUpdateAggOracleProposerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeUpdateAggOracleProposer)
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
		it.Event = new(IaggoraclecommitteeUpdateAggOracleProposer)
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
func (it *IaggoraclecommitteeUpdateAggOracleProposerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeUpdateAggOracleProposerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeUpdateAggOracleProposer represents a UpdateAggOracleProposer event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeUpdateAggOracleProposer struct {
	NewProposer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggOracleProposer is a free log retrieval operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterUpdateAggOracleProposer(opts *bind.FilterOpts, newProposer []common.Address) (*IaggoraclecommitteeUpdateAggOracleProposerIterator, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeUpdateAggOracleProposerIterator{contract: _Iaggoraclecommittee.contract, event: "UpdateAggOracleProposer", logs: logs, sub: sub}, nil
}

// WatchUpdateAggOracleProposer is a free log subscription operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchUpdateAggOracleProposer(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeUpdateAggOracleProposer, newProposer []common.Address) (event.Subscription, error) {

	var newProposerRule []interface{}
	for _, newProposerItem := range newProposer {
		newProposerRule = append(newProposerRule, newProposerItem)
	}

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "UpdateAggOracleProposer", newProposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeUpdateAggOracleProposer)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
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

// ParseUpdateAggOracleProposer is a log parse operation binding the contract event 0x93c536fad4968838f70b6a5841cc12ffbd1a657048ef449f715f0d9ebdbeffca.
//
// Solidity: event UpdateAggOracleProposer(address indexed newProposer)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseUpdateAggOracleProposer(log types.Log) (*IaggoraclecommitteeUpdateAggOracleProposer, error) {
	event := new(IaggoraclecommitteeUpdateAggOracleProposer)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "UpdateAggOracleProposer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator is returned from FilterValidatorsAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsAndThresholdUpdated events raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator struct {
	Event *IaggoraclecommitteeValidatorsAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IaggoraclecommitteeValidatorsAndThresholdUpdated)
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
		it.Event = new(IaggoraclecommitteeValidatorsAndThresholdUpdated)
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
func (it *IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IaggoraclecommitteeValidatorsAndThresholdUpdated represents a ValidatorsAndThresholdUpdated event raised by the Iaggoraclecommittee contract.
type IaggoraclecommitteeValidatorsAndThresholdUpdated struct {
	Validators []common.Address
	Threshold  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsAndThresholdUpdated is a free log retrieval operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) FilterValidatorsAndThresholdUpdated(opts *bind.FilterOpts) (*IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.FilterLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &IaggoraclecommitteeValidatorsAndThresholdUpdatedIterator{contract: _Iaggoraclecommittee.contract, event: "ValidatorsAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsAndThresholdUpdated is a free log subscription operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) WatchValidatorsAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *IaggoraclecommitteeValidatorsAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Iaggoraclecommittee.contract.WatchLogs(opts, "ValidatorsAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IaggoraclecommitteeValidatorsAndThresholdUpdated)
				if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
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

// ParseValidatorsAndThresholdUpdated is a log parse operation binding the contract event 0x50051c796f4289791fedb6e1f365497b6527df6f6f19527d566e3c9791af5460.
//
// Solidity: event ValidatorsAndThresholdUpdated(address[] validators, uint256 threshold)
func (_Iaggoraclecommittee *IaggoraclecommitteeFilterer) ParseValidatorsAndThresholdUpdated(log types.Log) (*IaggoraclecommitteeValidatorsAndThresholdUpdated, error) {
	event := new(IaggoraclecommitteeValidatorsAndThresholdUpdated)
	if err := _Iaggoraclecommittee.contract.UnpackLog(event, "ValidatorsAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

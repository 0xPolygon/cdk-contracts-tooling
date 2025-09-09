// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonconsensusbasepessimistic

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

// PolygonconsensusbasepessimisticMetaData contains all meta data concerning the Polygonconsensusbasepessimistic contract.
var PolygonconsensusbasepessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AdminCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManagerPessimistic\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonconsensusbasepessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonconsensusbasepessimisticMetaData.ABI instead.
var PolygonconsensusbasepessimisticABI = PolygonconsensusbasepessimisticMetaData.ABI

// Polygonconsensusbasepessimistic is an auto generated Go binding around an Ethereum contract.
type Polygonconsensusbasepessimistic struct {
	PolygonconsensusbasepessimisticCaller     // Read-only binding to the contract
	PolygonconsensusbasepessimisticTransactor // Write-only binding to the contract
	PolygonconsensusbasepessimisticFilterer   // Log filterer for contract events
}

// PolygonconsensusbasepessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonconsensusbasepessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconsensusbasepessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonconsensusbasepessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconsensusbasepessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonconsensusbasepessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonconsensusbasepessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonconsensusbasepessimisticSession struct {
	Contract     *Polygonconsensusbasepessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolygonconsensusbasepessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonconsensusbasepessimisticCallerSession struct {
	Contract *PolygonconsensusbasepessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PolygonconsensusbasepessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonconsensusbasepessimisticTransactorSession struct {
	Contract     *PolygonconsensusbasepessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PolygonconsensusbasepessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonconsensusbasepessimisticRaw struct {
	Contract *Polygonconsensusbasepessimistic // Generic contract binding to access the raw methods on
}

// PolygonconsensusbasepessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonconsensusbasepessimisticCallerRaw struct {
	Contract *PolygonconsensusbasepessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonconsensusbasepessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonconsensusbasepessimisticTransactorRaw struct {
	Contract *PolygonconsensusbasepessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonconsensusbasepessimistic creates a new instance of Polygonconsensusbasepessimistic, bound to a specific deployed contract.
func NewPolygonconsensusbasepessimistic(address common.Address, backend bind.ContractBackend) (*Polygonconsensusbasepessimistic, error) {
	contract, err := bindPolygonconsensusbasepessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonconsensusbasepessimistic{PolygonconsensusbasepessimisticCaller: PolygonconsensusbasepessimisticCaller{contract: contract}, PolygonconsensusbasepessimisticTransactor: PolygonconsensusbasepessimisticTransactor{contract: contract}, PolygonconsensusbasepessimisticFilterer: PolygonconsensusbasepessimisticFilterer{contract: contract}}, nil
}

// NewPolygonconsensusbasepessimisticCaller creates a new read-only instance of Polygonconsensusbasepessimistic, bound to a specific deployed contract.
func NewPolygonconsensusbasepessimisticCaller(address common.Address, caller bind.ContractCaller) (*PolygonconsensusbasepessimisticCaller, error) {
	contract, err := bindPolygonconsensusbasepessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticCaller{contract: contract}, nil
}

// NewPolygonconsensusbasepessimisticTransactor creates a new write-only instance of Polygonconsensusbasepessimistic, bound to a specific deployed contract.
func NewPolygonconsensusbasepessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonconsensusbasepessimisticTransactor, error) {
	contract, err := bindPolygonconsensusbasepessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticTransactor{contract: contract}, nil
}

// NewPolygonconsensusbasepessimisticFilterer creates a new log filterer instance of Polygonconsensusbasepessimistic, bound to a specific deployed contract.
func NewPolygonconsensusbasepessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonconsensusbasepessimisticFilterer, error) {
	contract, err := bindPolygonconsensusbasepessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticFilterer{contract: contract}, nil
}

// bindPolygonconsensusbasepessimistic binds a generic wrapper to an already deployed contract.
func bindPolygonconsensusbasepessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonconsensusbasepessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonconsensusbasepessimistic.Contract.PolygonconsensusbasepessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.PolygonconsensusbasepessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.PolygonconsensusbasepessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonconsensusbasepessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) Admin() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.Admin(&_Polygonconsensusbasepessimistic.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) Admin() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.Admin(&_Polygonconsensusbasepessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) BridgeAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.BridgeAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.BridgeAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForceBatchAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForceBatchAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForceBatchTimeout(&_Polygonconsensusbasepessimistic.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForceBatchTimeout(&_Polygonconsensusbasepessimistic.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForcedBatches(&_Polygonconsensusbasepessimistic.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonconsensusbasepessimistic.Contract.ForcedBatches(&_Polygonconsensusbasepessimistic.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) GasTokenAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.GasTokenAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.GasTokenAddress(&_Polygonconsensusbasepessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) GasTokenNetwork() (uint32, error) {
	return _Polygonconsensusbasepessimistic.Contract.GasTokenNetwork(&_Polygonconsensusbasepessimistic.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonconsensusbasepessimistic.Contract.GasTokenNetwork(&_Polygonconsensusbasepessimistic.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.GlobalExitRootManager(&_Polygonconsensusbasepessimistic.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.GlobalExitRootManager(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastAccInputHash(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastAccInputHash(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) LastForceBatch() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastForceBatch(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastForceBatch(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastForceBatchSequenced(&_Polygonconsensusbasepessimistic.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonconsensusbasepessimistic.Contract.LastForceBatchSequenced(&_Polygonconsensusbasepessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) NetworkName() (string, error) {
	return _Polygonconsensusbasepessimistic.Contract.NetworkName(&_Polygonconsensusbasepessimistic.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) NetworkName() (string, error) {
	return _Polygonconsensusbasepessimistic.Contract.NetworkName(&_Polygonconsensusbasepessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) PendingAdmin() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.PendingAdmin(&_Polygonconsensusbasepessimistic.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.PendingAdmin(&_Polygonconsensusbasepessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) Pol() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.Pol(&_Polygonconsensusbasepessimistic.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) Pol() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.Pol(&_Polygonconsensusbasepessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) RollupManager() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.RollupManager(&_Polygonconsensusbasepessimistic.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) RollupManager() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.RollupManager(&_Polygonconsensusbasepessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) TrustedSequencer() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.TrustedSequencer(&_Polygonconsensusbasepessimistic.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonconsensusbasepessimistic.Contract.TrustedSequencer(&_Polygonconsensusbasepessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonconsensusbasepessimistic.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) TrustedSequencerURL() (string, error) {
	return _Polygonconsensusbasepessimistic.Contract.TrustedSequencerURL(&_Polygonconsensusbasepessimistic.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonconsensusbasepessimistic.Contract.TrustedSequencerURL(&_Polygonconsensusbasepessimistic.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.AcceptAdminRole(&_Polygonconsensusbasepessimistic.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.AcceptAdminRole(&_Polygonconsensusbasepessimistic.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.contract.Transact(opts, "initialize", _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) Initialize(_admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.Initialize(&_Polygonconsensusbasepessimistic.TransactOpts, _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 , address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorSession) Initialize(_admin common.Address, sequencer common.Address, arg2 uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.Initialize(&_Polygonconsensusbasepessimistic.TransactOpts, _admin, sequencer, arg2, _gasTokenAddress, sequencerURL, _networkName)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.SetTrustedSequencer(&_Polygonconsensusbasepessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.SetTrustedSequencer(&_Polygonconsensusbasepessimistic.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.SetTrustedSequencerURL(&_Polygonconsensusbasepessimistic.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.SetTrustedSequencerURL(&_Polygonconsensusbasepessimistic.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.TransferAdminRole(&_Polygonconsensusbasepessimistic.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonconsensusbasepessimistic.Contract.TransferAdminRole(&_Polygonconsensusbasepessimistic.TransactOpts, newPendingAdmin)
}

// PolygonconsensusbasepessimisticAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticAcceptAdminRoleIterator struct {
	Event *PolygonconsensusbasepessimisticAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonconsensusbasepessimisticAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonconsensusbasepessimisticAcceptAdminRole)
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
		it.Event = new(PolygonconsensusbasepessimisticAcceptAdminRole)
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
func (it *PolygonconsensusbasepessimisticAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonconsensusbasepessimisticAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonconsensusbasepessimisticAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonconsensusbasepessimisticAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticAcceptAdminRoleIterator{contract: _Polygonconsensusbasepessimistic.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonconsensusbasepessimisticAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonconsensusbasepessimisticAcceptAdminRole)
				if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonconsensusbasepessimisticAcceptAdminRole, error) {
	event := new(PolygonconsensusbasepessimisticAcceptAdminRole)
	if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonconsensusbasepessimisticInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticInitializedIterator struct {
	Event *PolygonconsensusbasepessimisticInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonconsensusbasepessimisticInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonconsensusbasepessimisticInitialized)
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
		it.Event = new(PolygonconsensusbasepessimisticInitialized)
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
func (it *PolygonconsensusbasepessimisticInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonconsensusbasepessimisticInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonconsensusbasepessimisticInitialized represents a Initialized event raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonconsensusbasepessimisticInitializedIterator, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticInitializedIterator{contract: _Polygonconsensusbasepessimistic.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonconsensusbasepessimisticInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonconsensusbasepessimisticInitialized)
				if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) ParseInitialized(log types.Log) (*PolygonconsensusbasepessimisticInitialized, error) {
	event := new(PolygonconsensusbasepessimisticInitialized)
	if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonconsensusbasepessimisticSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticSetTrustedSequencerIterator struct {
	Event *PolygonconsensusbasepessimisticSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonconsensusbasepessimisticSetTrustedSequencer)
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
		it.Event = new(PolygonconsensusbasepessimisticSetTrustedSequencer)
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
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonconsensusbasepessimisticSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonconsensusbasepessimisticSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticSetTrustedSequencerIterator{contract: _Polygonconsensusbasepessimistic.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonconsensusbasepessimisticSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonconsensusbasepessimisticSetTrustedSequencer)
				if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonconsensusbasepessimisticSetTrustedSequencer, error) {
	event := new(PolygonconsensusbasepessimisticSetTrustedSequencer)
	if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator struct {
	Event *PolygonconsensusbasepessimisticSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonconsensusbasepessimisticSetTrustedSequencerURL)
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
		it.Event = new(PolygonconsensusbasepessimisticSetTrustedSequencerURL)
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
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonconsensusbasepessimisticSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticSetTrustedSequencerURLIterator{contract: _Polygonconsensusbasepessimistic.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonconsensusbasepessimisticSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonconsensusbasepessimisticSetTrustedSequencerURL)
				if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonconsensusbasepessimisticSetTrustedSequencerURL, error) {
	event := new(PolygonconsensusbasepessimisticSetTrustedSequencerURL)
	if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonconsensusbasepessimisticTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticTransferAdminRoleIterator struct {
	Event *PolygonconsensusbasepessimisticTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonconsensusbasepessimisticTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonconsensusbasepessimisticTransferAdminRole)
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
		it.Event = new(PolygonconsensusbasepessimisticTransferAdminRole)
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
func (it *PolygonconsensusbasepessimisticTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonconsensusbasepessimisticTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonconsensusbasepessimisticTransferAdminRole represents a TransferAdminRole event raised by the Polygonconsensusbasepessimistic contract.
type PolygonconsensusbasepessimisticTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonconsensusbasepessimisticTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonconsensusbasepessimisticTransferAdminRoleIterator{contract: _Polygonconsensusbasepessimistic.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonconsensusbasepessimisticTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonconsensusbasepessimistic.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonconsensusbasepessimisticTransferAdminRole)
				if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonconsensusbasepessimistic *PolygonconsensusbasepessimisticFilterer) ParseTransferAdminRole(log types.Log) (*PolygonconsensusbasepessimisticTransferAdminRole, error) {
	event := new(PolygonconsensusbasepessimisticTransferAdminRole)
	if err := _Polygonconsensusbasepessimistic.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

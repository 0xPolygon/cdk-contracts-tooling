// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggchainbase

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

// AggchainbaseMetaData contains all meta data concerning the Aggchainbase contract.
var AggchainbaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainRouteAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainRouteNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainVKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadySet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useDefaultGateway\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultGatewayFlag\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addAggchainRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"aggchainVKeyRoutes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateAggchainRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AggchainbaseABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainbaseMetaData.ABI instead.
var AggchainbaseABI = AggchainbaseMetaData.ABI

// Aggchainbase is an auto generated Go binding around an Ethereum contract.
type Aggchainbase struct {
	AggchainbaseCaller     // Read-only binding to the contract
	AggchainbaseTransactor // Write-only binding to the contract
	AggchainbaseFilterer   // Log filterer for contract events
}

// AggchainbaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainbaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainbaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainbaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainbaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainbaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainbaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainbaseSession struct {
	Contract     *Aggchainbase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggchainbaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainbaseCallerSession struct {
	Contract *AggchainbaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AggchainbaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainbaseTransactorSession struct {
	Contract     *AggchainbaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AggchainbaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainbaseRaw struct {
	Contract *Aggchainbase // Generic contract binding to access the raw methods on
}

// AggchainbaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainbaseCallerRaw struct {
	Contract *AggchainbaseCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainbaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainbaseTransactorRaw struct {
	Contract *AggchainbaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainbase creates a new instance of Aggchainbase, bound to a specific deployed contract.
func NewAggchainbase(address common.Address, backend bind.ContractBackend) (*Aggchainbase, error) {
	contract, err := bindAggchainbase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggchainbase{AggchainbaseCaller: AggchainbaseCaller{contract: contract}, AggchainbaseTransactor: AggchainbaseTransactor{contract: contract}, AggchainbaseFilterer: AggchainbaseFilterer{contract: contract}}, nil
}

// NewAggchainbaseCaller creates a new read-only instance of Aggchainbase, bound to a specific deployed contract.
func NewAggchainbaseCaller(address common.Address, caller bind.ContractCaller) (*AggchainbaseCaller, error) {
	contract, err := bindAggchainbase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainbaseCaller{contract: contract}, nil
}

// NewAggchainbaseTransactor creates a new write-only instance of Aggchainbase, bound to a specific deployed contract.
func NewAggchainbaseTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainbaseTransactor, error) {
	contract, err := bindAggchainbase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainbaseTransactor{contract: contract}, nil
}

// NewAggchainbaseFilterer creates a new log filterer instance of Aggchainbase, bound to a specific deployed contract.
func NewAggchainbaseFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainbaseFilterer, error) {
	contract, err := bindAggchainbase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainbaseFilterer{contract: contract}, nil
}

// bindAggchainbase binds a generic wrapper to an already deployed contract.
func bindAggchainbase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainbaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainbase *AggchainbaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainbase.Contract.AggchainbaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainbase *AggchainbaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainbase.Contract.AggchainbaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainbase *AggchainbaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainbase.Contract.AggchainbaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainbase *AggchainbaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainbase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainbase *AggchainbaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainbase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainbase *AggchainbaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainbase.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainbase *AggchainbaseCaller) AGGCHAINTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainbase *AggchainbaseSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainbase.Contract.AGGCHAINTYPE(&_Aggchainbase.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(uint32)
func (_Aggchainbase *AggchainbaseCallerSession) AGGCHAINTYPE() (uint32, error) {
	return _Aggchainbase.Contract.AGGCHAINTYPE(&_Aggchainbase.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainbase *AggchainbaseSession) Admin() (common.Address, error) {
	return _Aggchainbase.Contract.Admin(&_Aggchainbase.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) Admin() (common.Address, error) {
	return _Aggchainbase.Contract.Admin(&_Aggchainbase.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainbase *AggchainbaseSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainbase.Contract.AggLayerGateway(&_Aggchainbase.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainbase.Contract.AggLayerGateway(&_Aggchainbase.CallOpts)
}

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainbase *AggchainbaseCaller) AggchainVKeyRoutes(opts *bind.CallOpts, aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "aggchainVKeyRoutes", aggchainVKeySelector)

	outstruct := new(struct {
		AggchainVKey [32]byte
		Frozen       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AggchainVKey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Frozen = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainbase *AggchainbaseSession) AggchainVKeyRoutes(aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	return _Aggchainbase.Contract.AggchainVKeyRoutes(&_Aggchainbase.CallOpts, aggchainVKeySelector)
}

// AggchainVKeyRoutes is a free data retrieval call binding the contract method 0x69d8f277.
//
// Solidity: function aggchainVKeyRoutes(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey, bool frozen)
func (_Aggchainbase *AggchainbaseCallerSession) AggchainVKeyRoutes(aggchainVKeySelector [4]byte) (struct {
	AggchainVKey [32]byte
	Frozen       bool
}, error) {
	return _Aggchainbase.Contract.AggchainVKeyRoutes(&_Aggchainbase.CallOpts, aggchainVKeySelector)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainbase *AggchainbaseSession) BridgeAddress() (common.Address, error) {
	return _Aggchainbase.Contract.BridgeAddress(&_Aggchainbase.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) BridgeAddress() (common.Address, error) {
	return _Aggchainbase.Contract.BridgeAddress(&_Aggchainbase.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainbase *AggchainbaseSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainbase.Contract.ForceBatchAddress(&_Aggchainbase.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainbase.Contract.ForceBatchAddress(&_Aggchainbase.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainbase *AggchainbaseCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainbase *AggchainbaseSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainbase.Contract.ForceBatchTimeout(&_Aggchainbase.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainbase *AggchainbaseCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainbase.Contract.ForceBatchTimeout(&_Aggchainbase.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainbase *AggchainbaseCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainbase *AggchainbaseSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainbase.Contract.ForcedBatches(&_Aggchainbase.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainbase *AggchainbaseCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainbase.Contract.ForcedBatches(&_Aggchainbase.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainbase *AggchainbaseSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainbase.Contract.GasTokenAddress(&_Aggchainbase.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainbase.Contract.GasTokenAddress(&_Aggchainbase.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainbase *AggchainbaseCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainbase *AggchainbaseSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainbase.Contract.GasTokenNetwork(&_Aggchainbase.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainbase *AggchainbaseCallerSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainbase.Contract.GasTokenNetwork(&_Aggchainbase.CallOpts)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainbase *AggchainbaseCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainSelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "getAggchainVKey", aggchainSelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainbase *AggchainbaseSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainbase.Contract.GetAggchainVKey(&_Aggchainbase.CallOpts, aggchainSelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainSelector) view returns(bytes32 aggchainVKey)
func (_Aggchainbase *AggchainbaseCallerSession) GetAggchainVKey(aggchainSelector [4]byte) ([32]byte, error) {
	return _Aggchainbase.Contract.GetAggchainVKey(&_Aggchainbase.CallOpts, aggchainSelector)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainbase *AggchainbaseSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainbase.Contract.GlobalExitRootManager(&_Aggchainbase.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainbase.Contract.GlobalExitRootManager(&_Aggchainbase.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainbase *AggchainbaseCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainbase *AggchainbaseSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainbase.Contract.Initialize0(&_Aggchainbase.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainbase *AggchainbaseCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainbase.Contract.Initialize0(&_Aggchainbase.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainbase *AggchainbaseCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainbase *AggchainbaseSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainbase.Contract.LastAccInputHash(&_Aggchainbase.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainbase *AggchainbaseCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainbase.Contract.LastAccInputHash(&_Aggchainbase.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainbase *AggchainbaseCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainbase *AggchainbaseSession) LastForceBatch() (uint64, error) {
	return _Aggchainbase.Contract.LastForceBatch(&_Aggchainbase.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainbase *AggchainbaseCallerSession) LastForceBatch() (uint64, error) {
	return _Aggchainbase.Contract.LastForceBatch(&_Aggchainbase.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainbase *AggchainbaseCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainbase *AggchainbaseSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainbase.Contract.LastForceBatchSequenced(&_Aggchainbase.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainbase *AggchainbaseCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainbase.Contract.LastForceBatchSequenced(&_Aggchainbase.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainbase *AggchainbaseCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainbase *AggchainbaseSession) NetworkName() (string, error) {
	return _Aggchainbase.Contract.NetworkName(&_Aggchainbase.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainbase *AggchainbaseCallerSession) NetworkName() (string, error) {
	return _Aggchainbase.Contract.NetworkName(&_Aggchainbase.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainbase *AggchainbaseSession) PendingAdmin() (common.Address, error) {
	return _Aggchainbase.Contract.PendingAdmin(&_Aggchainbase.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) PendingAdmin() (common.Address, error) {
	return _Aggchainbase.Contract.PendingAdmin(&_Aggchainbase.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainbase *AggchainbaseSession) Pol() (common.Address, error) {
	return _Aggchainbase.Contract.Pol(&_Aggchainbase.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) Pol() (common.Address, error) {
	return _Aggchainbase.Contract.Pol(&_Aggchainbase.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainbase *AggchainbaseSession) RollupManager() (common.Address, error) {
	return _Aggchainbase.Contract.RollupManager(&_Aggchainbase.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) RollupManager() (common.Address, error) {
	return _Aggchainbase.Contract.RollupManager(&_Aggchainbase.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainbase *AggchainbaseCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainbase *AggchainbaseSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainbase.Contract.TrustedSequencer(&_Aggchainbase.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainbase *AggchainbaseCallerSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainbase.Contract.TrustedSequencer(&_Aggchainbase.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainbase *AggchainbaseCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainbase *AggchainbaseSession) TrustedSequencerURL() (string, error) {
	return _Aggchainbase.Contract.TrustedSequencerURL(&_Aggchainbase.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainbase *AggchainbaseCallerSession) TrustedSequencerURL() (string, error) {
	return _Aggchainbase.Contract.TrustedSequencerURL(&_Aggchainbase.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainbase *AggchainbaseCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainbase.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainbase *AggchainbaseSession) UseDefaultGateway() (bool, error) {
	return _Aggchainbase.Contract.UseDefaultGateway(&_Aggchainbase.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainbase *AggchainbaseCallerSession) UseDefaultGateway() (bool, error) {
	return _Aggchainbase.Contract.UseDefaultGateway(&_Aggchainbase.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainbase *AggchainbaseTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainbase *AggchainbaseSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainbase.Contract.AcceptAdminRole(&_Aggchainbase.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainbase *AggchainbaseTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainbase.Contract.AcceptAdminRole(&_Aggchainbase.TransactOpts)
}

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainbase *AggchainbaseTransactor) AddAggchainRoute(opts *bind.TransactOpts, aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "addAggchainRoute", aggchainSelector, aggchainVKey)
}

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainbase *AggchainbaseSession) AddAggchainRoute(aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.AddAggchainRoute(&_Aggchainbase.TransactOpts, aggchainSelector, aggchainVKey)
}

// AddAggchainRoute is a paid mutator transaction binding the contract method 0xdf4970a8.
//
// Solidity: function addAggchainRoute(bytes4 aggchainSelector, bytes32 aggchainVKey) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) AddAggchainRoute(aggchainSelector [4]byte, aggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.AddAggchainRoute(&_Aggchainbase.TransactOpts, aggchainSelector, aggchainVKey)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainbase.Contract.DisableUseDefaultGatewayFlag(&_Aggchainbase.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainbase.Contract.DisableUseDefaultGatewayFlag(&_Aggchainbase.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainbase.Contract.EnableUseDefaultGatewayFlag(&_Aggchainbase.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainbase *AggchainbaseTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainbase.Contract.EnableUseDefaultGatewayFlag(&_Aggchainbase.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainbase *AggchainbaseTransactor) Initialize(opts *bind.TransactOpts, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "initialize", initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainbase *AggchainbaseSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.Initialize(&_Aggchainbase.TransactOpts, initializeBytesCustomChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesCustomChain) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) Initialize(initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.Initialize(&_Aggchainbase.TransactOpts, initializeBytesCustomChain)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainbase *AggchainbaseTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainbase *AggchainbaseSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainbase.Contract.SetTrustedSequencer(&_Aggchainbase.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainbase.Contract.SetTrustedSequencer(&_Aggchainbase.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainbase *AggchainbaseTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainbase *AggchainbaseSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainbase.Contract.SetTrustedSequencerURL(&_Aggchainbase.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainbase.Contract.SetTrustedSequencerURL(&_Aggchainbase.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainbase *AggchainbaseTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainbase *AggchainbaseSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainbase.Contract.TransferAdminRole(&_Aggchainbase.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainbase.Contract.TransferAdminRole(&_Aggchainbase.TransactOpts, newPendingAdmin)
}

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainbase *AggchainbaseTransactor) UpdateAggchainRoute(opts *bind.TransactOpts, aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.contract.Transact(opts, "updateAggchainRoute", aggchainSelector, updatedAggchainVKey)
}

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainbase *AggchainbaseSession) UpdateAggchainRoute(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.UpdateAggchainRoute(&_Aggchainbase.TransactOpts, aggchainSelector, updatedAggchainVKey)
}

// UpdateAggchainRoute is a paid mutator transaction binding the contract method 0xab27f141.
//
// Solidity: function updateAggchainRoute(bytes4 aggchainSelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainbase *AggchainbaseTransactorSession) UpdateAggchainRoute(aggchainSelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainbase.Contract.UpdateAggchainRoute(&_Aggchainbase.TransactOpts, aggchainSelector, updatedAggchainVKey)
}

// AggchainbaseAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Aggchainbase contract.
type AggchainbaseAcceptAdminRoleIterator struct {
	Event *AggchainbaseAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainbaseAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseAcceptAdminRole)
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
		it.Event = new(AggchainbaseAcceptAdminRole)
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
func (it *AggchainbaseAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseAcceptAdminRole represents a AcceptAdminRole event raised by the Aggchainbase contract.
type AggchainbaseAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainbase *AggchainbaseFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainbaseAcceptAdminRoleIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseAcceptAdminRoleIterator{contract: _Aggchainbase.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainbase *AggchainbaseFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainbaseAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseAcceptAdminRole)
				if err := _Aggchainbase.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Aggchainbase *AggchainbaseFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainbaseAcceptAdminRole, error) {
	event := new(AggchainbaseAcceptAdminRole)
	if err := _Aggchainbase.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Aggchainbase contract.
type AggchainbaseAddAggchainVKeyIterator struct {
	Event *AggchainbaseAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainbaseAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseAddAggchainVKey)
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
		it.Event = new(AggchainbaseAddAggchainVKey)
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
func (it *AggchainbaseAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseAddAggchainVKey represents a AddAggchainVKey event raised by the Aggchainbase contract.
type AggchainbaseAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainbaseAddAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseAddAggchainVKeyIterator{contract: _Aggchainbase.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainbaseAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseAddAggchainVKey)
				if err := _Aggchainbase.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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

// ParseAddAggchainVKey is a log parse operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainbaseAddAggchainVKey, error) {
	event := new(AggchainbaseAddAggchainVKey)
	if err := _Aggchainbase.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggchainbase contract.
type AggchainbaseInitializedIterator struct {
	Event *AggchainbaseInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainbaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseInitialized)
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
		it.Event = new(AggchainbaseInitialized)
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
func (it *AggchainbaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseInitialized represents a Initialized event raised by the Aggchainbase contract.
type AggchainbaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainbase *AggchainbaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainbaseInitializedIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseInitializedIterator{contract: _Aggchainbase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainbase *AggchainbaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainbaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseInitialized)
				if err := _Aggchainbase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggchainbase *AggchainbaseFilterer) ParseInitialized(log types.Log) (*AggchainbaseInitialized, error) {
	event := new(AggchainbaseInitialized)
	if err := _Aggchainbase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Aggchainbase contract.
type AggchainbaseSetTrustedSequencerIterator struct {
	Event *AggchainbaseSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainbaseSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseSetTrustedSequencer)
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
		it.Event = new(AggchainbaseSetTrustedSequencer)
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
func (it *AggchainbaseSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseSetTrustedSequencer represents a SetTrustedSequencer event raised by the Aggchainbase contract.
type AggchainbaseSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainbase *AggchainbaseFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainbaseSetTrustedSequencerIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseSetTrustedSequencerIterator{contract: _Aggchainbase.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainbase *AggchainbaseFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainbaseSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseSetTrustedSequencer)
				if err := _Aggchainbase.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Aggchainbase *AggchainbaseFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainbaseSetTrustedSequencer, error) {
	event := new(AggchainbaseSetTrustedSequencer)
	if err := _Aggchainbase.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Aggchainbase contract.
type AggchainbaseSetTrustedSequencerURLIterator struct {
	Event *AggchainbaseSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainbaseSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseSetTrustedSequencerURL)
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
		it.Event = new(AggchainbaseSetTrustedSequencerURL)
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
func (it *AggchainbaseSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Aggchainbase contract.
type AggchainbaseSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainbase *AggchainbaseFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainbaseSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseSetTrustedSequencerURLIterator{contract: _Aggchainbase.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainbase *AggchainbaseFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainbaseSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseSetTrustedSequencerURL)
				if err := _Aggchainbase.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Aggchainbase *AggchainbaseFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainbaseSetTrustedSequencerURL, error) {
	event := new(AggchainbaseSetTrustedSequencerURL)
	if err := _Aggchainbase.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Aggchainbase contract.
type AggchainbaseTransferAdminRoleIterator struct {
	Event *AggchainbaseTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainbaseTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseTransferAdminRole)
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
		it.Event = new(AggchainbaseTransferAdminRole)
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
func (it *AggchainbaseTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseTransferAdminRole represents a TransferAdminRole event raised by the Aggchainbase contract.
type AggchainbaseTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainbase *AggchainbaseFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainbaseTransferAdminRoleIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseTransferAdminRoleIterator{contract: _Aggchainbase.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainbase *AggchainbaseFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainbaseTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseTransferAdminRole)
				if err := _Aggchainbase.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Aggchainbase *AggchainbaseFilterer) ParseTransferAdminRole(log types.Log) (*AggchainbaseTransferAdminRole, error) {
	event := new(AggchainbaseTransferAdminRole)
	if err := _Aggchainbase.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Aggchainbase contract.
type AggchainbaseUpdateAggchainVKeyIterator struct {
	Event *AggchainbaseUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainbaseUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseUpdateAggchainVKey)
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
		it.Event = new(AggchainbaseUpdateAggchainVKey)
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
func (it *AggchainbaseUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Aggchainbase contract.
type AggchainbaseUpdateAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainbaseUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseUpdateAggchainVKeyIterator{contract: _Aggchainbase.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainbaseUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseUpdateAggchainVKey)
				if err := _Aggchainbase.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x2ae81c794f274593f3cd22db9ad3af161ff9d93d15922b35629a6ffd9be7243f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainbase *AggchainbaseFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainbaseUpdateAggchainVKey, error) {
	event := new(AggchainbaseUpdateAggchainVKey)
	if err := _Aggchainbase.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainbaseUpdateUseDefaultGatewayFlagIterator is returned from FilterUpdateUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultGatewayFlag events raised by the Aggchainbase contract.
type AggchainbaseUpdateUseDefaultGatewayFlagIterator struct {
	Event *AggchainbaseUpdateUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainbaseUpdateUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainbaseUpdateUseDefaultGatewayFlag)
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
		it.Event = new(AggchainbaseUpdateUseDefaultGatewayFlag)
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
func (it *AggchainbaseUpdateUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainbaseUpdateUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainbaseUpdateUseDefaultGatewayFlag represents a UpdateUseDefaultGatewayFlag event raised by the Aggchainbase contract.
type AggchainbaseUpdateUseDefaultGatewayFlag struct {
	UseDefaultGateway bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainbase *AggchainbaseFilterer) FilterUpdateUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainbaseUpdateUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainbase.contract.FilterLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainbaseUpdateUseDefaultGatewayFlagIterator{contract: _Aggchainbase.contract, event: "UpdateUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainbase *AggchainbaseFilterer) WatchUpdateUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainbaseUpdateUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainbase.contract.WatchLogs(opts, "UpdateUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainbaseUpdateUseDefaultGatewayFlag)
				if err := _Aggchainbase.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
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

// ParseUpdateUseDefaultGatewayFlag is a log parse operation binding the contract event 0xd4305f3587ac0e4b2b8b4e5b2072599efd931f81aba626cdbef4b040078b5070.
//
// Solidity: event UpdateUseDefaultGatewayFlag(bool useDefaultGateway)
func (_Aggchainbase *AggchainbaseFilterer) ParseUpdateUseDefaultGatewayFlag(log types.Log) (*AggchainbaseUpdateUseDefaultGatewayFlag, error) {
	event := new(AggchainbaseUpdateUseDefaultGatewayFlag)
	if err := _Aggchainbase.contract.UnpackLog(event, "UpdateUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

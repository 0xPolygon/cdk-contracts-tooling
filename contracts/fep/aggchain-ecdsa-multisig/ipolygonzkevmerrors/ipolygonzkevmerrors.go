// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmerrors

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

// IpolygonzkevmerrorsMetaData contains all meta data concerning the Ipolygonzkevmerrors contract.
var IpolygonzkevmerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"}]",
}

// IpolygonzkevmerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonzkevmerrorsMetaData.ABI instead.
var IpolygonzkevmerrorsABI = IpolygonzkevmerrorsMetaData.ABI

// Ipolygonzkevmerrors is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmerrors struct {
	IpolygonzkevmerrorsCaller     // Read-only binding to the contract
	IpolygonzkevmerrorsTransactor // Write-only binding to the contract
	IpolygonzkevmerrorsFilterer   // Log filterer for contract events
}

// IpolygonzkevmerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonzkevmerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonzkevmerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonzkevmerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonzkevmerrorsSession struct {
	Contract     *Ipolygonzkevmerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IpolygonzkevmerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonzkevmerrorsCallerSession struct {
	Contract *IpolygonzkevmerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IpolygonzkevmerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonzkevmerrorsTransactorSession struct {
	Contract     *IpolygonzkevmerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IpolygonzkevmerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonzkevmerrorsRaw struct {
	Contract *Ipolygonzkevmerrors // Generic contract binding to access the raw methods on
}

// IpolygonzkevmerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonzkevmerrorsCallerRaw struct {
	Contract *IpolygonzkevmerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonzkevmerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonzkevmerrorsTransactorRaw struct {
	Contract *IpolygonzkevmerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmerrors creates a new instance of Ipolygonzkevmerrors, bound to a specific deployed contract.
func NewIpolygonzkevmerrors(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmerrors, error) {
	contract, err := bindIpolygonzkevmerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmerrors{IpolygonzkevmerrorsCaller: IpolygonzkevmerrorsCaller{contract: contract}, IpolygonzkevmerrorsTransactor: IpolygonzkevmerrorsTransactor{contract: contract}, IpolygonzkevmerrorsFilterer: IpolygonzkevmerrorsFilterer{contract: contract}}, nil
}

// NewIpolygonzkevmerrorsCaller creates a new read-only instance of Ipolygonzkevmerrors, bound to a specific deployed contract.
func NewIpolygonzkevmerrorsCaller(address common.Address, caller bind.ContractCaller) (*IpolygonzkevmerrorsCaller, error) {
	contract, err := bindIpolygonzkevmerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmerrorsCaller{contract: contract}, nil
}

// NewIpolygonzkevmerrorsTransactor creates a new write-only instance of Ipolygonzkevmerrors, bound to a specific deployed contract.
func NewIpolygonzkevmerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonzkevmerrorsTransactor, error) {
	contract, err := bindIpolygonzkevmerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmerrorsTransactor{contract: contract}, nil
}

// NewIpolygonzkevmerrorsFilterer creates a new log filterer instance of Ipolygonzkevmerrors, bound to a specific deployed contract.
func NewIpolygonzkevmerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonzkevmerrorsFilterer, error) {
	contract, err := bindIpolygonzkevmerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmerrorsFilterer{contract: contract}, nil
}

// bindIpolygonzkevmerrors binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonzkevmerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmerrors.Contract.IpolygonzkevmerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmerrors.Contract.IpolygonzkevmerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmerrors.Contract.IpolygonzkevmerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmerrors *IpolygonzkevmerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmerrors.Contract.contract.Transact(opts, method, params...)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmetrogerrors

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

// IpolygonzkevmetrogerrorsMetaData contains all meta data concerning the Ipolygonzkevmetrogerrors contract.
var IpolygonzkevmetrogerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"}]",
}

// IpolygonzkevmetrogerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonzkevmetrogerrorsMetaData.ABI instead.
var IpolygonzkevmetrogerrorsABI = IpolygonzkevmetrogerrorsMetaData.ABI

// Ipolygonzkevmetrogerrors is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmetrogerrors struct {
	IpolygonzkevmetrogerrorsCaller     // Read-only binding to the contract
	IpolygonzkevmetrogerrorsTransactor // Write-only binding to the contract
	IpolygonzkevmetrogerrorsFilterer   // Log filterer for contract events
}

// IpolygonzkevmetrogerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonzkevmetrogerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmetrogerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonzkevmetrogerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmetrogerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonzkevmetrogerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmetrogerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonzkevmetrogerrorsSession struct {
	Contract     *Ipolygonzkevmetrogerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IpolygonzkevmetrogerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonzkevmetrogerrorsCallerSession struct {
	Contract *IpolygonzkevmetrogerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IpolygonzkevmetrogerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonzkevmetrogerrorsTransactorSession struct {
	Contract     *IpolygonzkevmetrogerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IpolygonzkevmetrogerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonzkevmetrogerrorsRaw struct {
	Contract *Ipolygonzkevmetrogerrors // Generic contract binding to access the raw methods on
}

// IpolygonzkevmetrogerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonzkevmetrogerrorsCallerRaw struct {
	Contract *IpolygonzkevmetrogerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonzkevmetrogerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonzkevmetrogerrorsTransactorRaw struct {
	Contract *IpolygonzkevmetrogerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmetrogerrors creates a new instance of Ipolygonzkevmetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmetrogerrors(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmetrogerrors, error) {
	contract, err := bindIpolygonzkevmetrogerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmetrogerrors{IpolygonzkevmetrogerrorsCaller: IpolygonzkevmetrogerrorsCaller{contract: contract}, IpolygonzkevmetrogerrorsTransactor: IpolygonzkevmetrogerrorsTransactor{contract: contract}, IpolygonzkevmetrogerrorsFilterer: IpolygonzkevmetrogerrorsFilterer{contract: contract}}, nil
}

// NewIpolygonzkevmetrogerrorsCaller creates a new read-only instance of Ipolygonzkevmetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmetrogerrorsCaller(address common.Address, caller bind.ContractCaller) (*IpolygonzkevmetrogerrorsCaller, error) {
	contract, err := bindIpolygonzkevmetrogerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmetrogerrorsCaller{contract: contract}, nil
}

// NewIpolygonzkevmetrogerrorsTransactor creates a new write-only instance of Ipolygonzkevmetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmetrogerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonzkevmetrogerrorsTransactor, error) {
	contract, err := bindIpolygonzkevmetrogerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmetrogerrorsTransactor{contract: contract}, nil
}

// NewIpolygonzkevmetrogerrorsFilterer creates a new log filterer instance of Ipolygonzkevmetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmetrogerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonzkevmetrogerrorsFilterer, error) {
	contract, err := bindIpolygonzkevmetrogerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmetrogerrorsFilterer{contract: contract}, nil
}

// bindIpolygonzkevmetrogerrors binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmetrogerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonzkevmetrogerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmetrogerrors.Contract.IpolygonzkevmetrogerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmetrogerrors.Contract.IpolygonzkevmetrogerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmetrogerrors.Contract.IpolygonzkevmetrogerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmetrogerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmetrogerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmetrogerrors *IpolygonzkevmetrogerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmetrogerrors.Contract.contract.Transact(opts, method, params...)
}

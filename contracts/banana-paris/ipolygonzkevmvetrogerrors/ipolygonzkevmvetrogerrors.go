// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonzkevmvetrogerrors

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

// IpolygonzkevmvetrogerrorsMetaData contains all meta data concerning the Ipolygonzkevmvetrogerrors contract.
var IpolygonzkevmvetrogerrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoRootIndexInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"}]",
}

// IpolygonzkevmvetrogerrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonzkevmvetrogerrorsMetaData.ABI instead.
var IpolygonzkevmvetrogerrorsABI = IpolygonzkevmvetrogerrorsMetaData.ABI

// Ipolygonzkevmvetrogerrors is an auto generated Go binding around an Ethereum contract.
type Ipolygonzkevmvetrogerrors struct {
	IpolygonzkevmvetrogerrorsCaller     // Read-only binding to the contract
	IpolygonzkevmvetrogerrorsTransactor // Write-only binding to the contract
	IpolygonzkevmvetrogerrorsFilterer   // Log filterer for contract events
}

// IpolygonzkevmvetrogerrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonzkevmvetrogerrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmvetrogerrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonzkevmvetrogerrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmvetrogerrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonzkevmvetrogerrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonzkevmvetrogerrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonzkevmvetrogerrorsSession struct {
	Contract     *Ipolygonzkevmvetrogerrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IpolygonzkevmvetrogerrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonzkevmvetrogerrorsCallerSession struct {
	Contract *IpolygonzkevmvetrogerrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IpolygonzkevmvetrogerrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonzkevmvetrogerrorsTransactorSession struct {
	Contract     *IpolygonzkevmvetrogerrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IpolygonzkevmvetrogerrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonzkevmvetrogerrorsRaw struct {
	Contract *Ipolygonzkevmvetrogerrors // Generic contract binding to access the raw methods on
}

// IpolygonzkevmvetrogerrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonzkevmvetrogerrorsCallerRaw struct {
	Contract *IpolygonzkevmvetrogerrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonzkevmvetrogerrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonzkevmvetrogerrorsTransactorRaw struct {
	Contract *IpolygonzkevmvetrogerrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonzkevmvetrogerrors creates a new instance of Ipolygonzkevmvetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmvetrogerrors(address common.Address, backend bind.ContractBackend) (*Ipolygonzkevmvetrogerrors, error) {
	contract, err := bindIpolygonzkevmvetrogerrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonzkevmvetrogerrors{IpolygonzkevmvetrogerrorsCaller: IpolygonzkevmvetrogerrorsCaller{contract: contract}, IpolygonzkevmvetrogerrorsTransactor: IpolygonzkevmvetrogerrorsTransactor{contract: contract}, IpolygonzkevmvetrogerrorsFilterer: IpolygonzkevmvetrogerrorsFilterer{contract: contract}}, nil
}

// NewIpolygonzkevmvetrogerrorsCaller creates a new read-only instance of Ipolygonzkevmvetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmvetrogerrorsCaller(address common.Address, caller bind.ContractCaller) (*IpolygonzkevmvetrogerrorsCaller, error) {
	contract, err := bindIpolygonzkevmvetrogerrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmvetrogerrorsCaller{contract: contract}, nil
}

// NewIpolygonzkevmvetrogerrorsTransactor creates a new write-only instance of Ipolygonzkevmvetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmvetrogerrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonzkevmvetrogerrorsTransactor, error) {
	contract, err := bindIpolygonzkevmvetrogerrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmvetrogerrorsTransactor{contract: contract}, nil
}

// NewIpolygonzkevmvetrogerrorsFilterer creates a new log filterer instance of Ipolygonzkevmvetrogerrors, bound to a specific deployed contract.
func NewIpolygonzkevmvetrogerrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonzkevmvetrogerrorsFilterer, error) {
	contract, err := bindIpolygonzkevmvetrogerrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonzkevmvetrogerrorsFilterer{contract: contract}, nil
}

// bindIpolygonzkevmvetrogerrors binds a generic wrapper to an already deployed contract.
func bindIpolygonzkevmvetrogerrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonzkevmvetrogerrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmvetrogerrors.Contract.IpolygonzkevmvetrogerrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmvetrogerrors.Contract.IpolygonzkevmvetrogerrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmvetrogerrors.Contract.IpolygonzkevmvetrogerrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonzkevmvetrogerrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonzkevmvetrogerrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonzkevmvetrogerrors *IpolygonzkevmvetrogerrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonzkevmvetrogerrors.Contract.contract.Transact(opts, method, params...)
}

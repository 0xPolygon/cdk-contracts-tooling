// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonrollupmanagerprevious

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

// IpolygonrollupmanagerpreviousMetaData contains all meta data concerning the Ipolygonrollupmanagerprevious contract.
var IpolygonrollupmanagerpreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"}]",
}

// IpolygonrollupmanagerpreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonrollupmanagerpreviousMetaData.ABI instead.
var IpolygonrollupmanagerpreviousABI = IpolygonrollupmanagerpreviousMetaData.ABI

// Ipolygonrollupmanagerprevious is an auto generated Go binding around an Ethereum contract.
type Ipolygonrollupmanagerprevious struct {
	IpolygonrollupmanagerpreviousCaller     // Read-only binding to the contract
	IpolygonrollupmanagerpreviousTransactor // Write-only binding to the contract
	IpolygonrollupmanagerpreviousFilterer   // Log filterer for contract events
}

// IpolygonrollupmanagerpreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonrollupmanagerpreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonrollupmanagerpreviousSession struct {
	Contract     *Ipolygonrollupmanagerprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerpreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonrollupmanagerpreviousCallerSession struct {
	Contract *IpolygonrollupmanagerpreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// IpolygonrollupmanagerpreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonrollupmanagerpreviousTransactorSession struct {
	Contract     *IpolygonrollupmanagerpreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerpreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonrollupmanagerpreviousRaw struct {
	Contract *Ipolygonrollupmanagerprevious // Generic contract binding to access the raw methods on
}

// IpolygonrollupmanagerpreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpreviousCallerRaw struct {
	Contract *IpolygonrollupmanagerpreviousCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonrollupmanagerpreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpreviousTransactorRaw struct {
	Contract *IpolygonrollupmanagerpreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonrollupmanagerprevious creates a new instance of Ipolygonrollupmanagerprevious, bound to a specific deployed contract.
func NewIpolygonrollupmanagerprevious(address common.Address, backend bind.ContractBackend) (*Ipolygonrollupmanagerprevious, error) {
	contract, err := bindIpolygonrollupmanagerprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonrollupmanagerprevious{IpolygonrollupmanagerpreviousCaller: IpolygonrollupmanagerpreviousCaller{contract: contract}, IpolygonrollupmanagerpreviousTransactor: IpolygonrollupmanagerpreviousTransactor{contract: contract}, IpolygonrollupmanagerpreviousFilterer: IpolygonrollupmanagerpreviousFilterer{contract: contract}}, nil
}

// NewIpolygonrollupmanagerpreviousCaller creates a new read-only instance of Ipolygonrollupmanagerprevious, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpreviousCaller(address common.Address, caller bind.ContractCaller) (*IpolygonrollupmanagerpreviousCaller, error) {
	contract, err := bindIpolygonrollupmanagerprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpreviousCaller{contract: contract}, nil
}

// NewIpolygonrollupmanagerpreviousTransactor creates a new write-only instance of Ipolygonrollupmanagerprevious, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonrollupmanagerpreviousTransactor, error) {
	contract, err := bindIpolygonrollupmanagerprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpreviousTransactor{contract: contract}, nil
}

// NewIpolygonrollupmanagerpreviousFilterer creates a new log filterer instance of Ipolygonrollupmanagerprevious, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonrollupmanagerpreviousFilterer, error) {
	contract, err := bindIpolygonrollupmanagerprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpreviousFilterer{contract: contract}, nil
}

// bindIpolygonrollupmanagerprevious binds a generic wrapper to an already deployed contract.
func bindIpolygonrollupmanagerprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonrollupmanagerpreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanagerprevious.Contract.IpolygonrollupmanagerpreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerprevious.Contract.IpolygonrollupmanagerpreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerprevious.Contract.IpolygonrollupmanagerpreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanagerprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanagerprevious *IpolygonrollupmanagerpreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerprevious.Contract.contract.Transact(opts, method, params...)
}

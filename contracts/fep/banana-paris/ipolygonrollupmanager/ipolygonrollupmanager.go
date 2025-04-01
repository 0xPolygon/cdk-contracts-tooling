// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonrollupmanager

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

// IpolygonrollupmanagerMetaData contains all meta data concerning the Ipolygonrollupmanager contract.
var IpolygonrollupmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"}]",
}

// IpolygonrollupmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonrollupmanagerMetaData.ABI instead.
var IpolygonrollupmanagerABI = IpolygonrollupmanagerMetaData.ABI

// Ipolygonrollupmanager is an auto generated Go binding around an Ethereum contract.
type Ipolygonrollupmanager struct {
	IpolygonrollupmanagerCaller     // Read-only binding to the contract
	IpolygonrollupmanagerTransactor // Write-only binding to the contract
	IpolygonrollupmanagerFilterer   // Log filterer for contract events
}

// IpolygonrollupmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonrollupmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonrollupmanagerSession struct {
	Contract     *Ipolygonrollupmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonrollupmanagerCallerSession struct {
	Contract *IpolygonrollupmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IpolygonrollupmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonrollupmanagerTransactorSession struct {
	Contract     *IpolygonrollupmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonrollupmanagerRaw struct {
	Contract *Ipolygonrollupmanager // Generic contract binding to access the raw methods on
}

// IpolygonrollupmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerCallerRaw struct {
	Contract *IpolygonrollupmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonrollupmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerTransactorRaw struct {
	Contract *IpolygonrollupmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonrollupmanager creates a new instance of Ipolygonrollupmanager, bound to a specific deployed contract.
func NewIpolygonrollupmanager(address common.Address, backend bind.ContractBackend) (*Ipolygonrollupmanager, error) {
	contract, err := bindIpolygonrollupmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonrollupmanager{IpolygonrollupmanagerCaller: IpolygonrollupmanagerCaller{contract: contract}, IpolygonrollupmanagerTransactor: IpolygonrollupmanagerTransactor{contract: contract}, IpolygonrollupmanagerFilterer: IpolygonrollupmanagerFilterer{contract: contract}}, nil
}

// NewIpolygonrollupmanagerCaller creates a new read-only instance of Ipolygonrollupmanager, bound to a specific deployed contract.
func NewIpolygonrollupmanagerCaller(address common.Address, caller bind.ContractCaller) (*IpolygonrollupmanagerCaller, error) {
	contract, err := bindIpolygonrollupmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerCaller{contract: contract}, nil
}

// NewIpolygonrollupmanagerTransactor creates a new write-only instance of Ipolygonrollupmanager, bound to a specific deployed contract.
func NewIpolygonrollupmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonrollupmanagerTransactor, error) {
	contract, err := bindIpolygonrollupmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerTransactor{contract: contract}, nil
}

// NewIpolygonrollupmanagerFilterer creates a new log filterer instance of Ipolygonrollupmanager, bound to a specific deployed contract.
func NewIpolygonrollupmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonrollupmanagerFilterer, error) {
	contract, err := bindIpolygonrollupmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerFilterer{contract: contract}, nil
}

// bindIpolygonrollupmanager binds a generic wrapper to an already deployed contract.
func bindIpolygonrollupmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonrollupmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanager.Contract.IpolygonrollupmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.IpolygonrollupmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.IpolygonrollupmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.contract.Transact(opts, method, params...)
}

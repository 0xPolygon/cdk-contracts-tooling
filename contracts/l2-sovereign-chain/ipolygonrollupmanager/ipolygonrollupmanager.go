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
	ABI: "[{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyChainsWithPessimisticProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesCustomChain\",\"type\":\"bytes\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"selectedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"customChainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.ActivateEmergencyState(&_Ipolygonrollupmanager.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.ActivateEmergencyState(&_Ipolygonrollupmanager.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.AddExistingRollup(&_Ipolygonrollupmanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.AddExistingRollup(&_Ipolygonrollupmanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.AddNewRollupType(&_Ipolygonrollupmanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.AddNewRollupType(&_Ipolygonrollupmanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// CalculateRewardPerBatch is a paid mutator transaction binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) CalculateRewardPerBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "calculateRewardPerBatch")
}

// CalculateRewardPerBatch is a paid mutator transaction binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) CalculateRewardPerBatch() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.CalculateRewardPerBatch(&_Ipolygonrollupmanager.TransactOpts)
}

// CalculateRewardPerBatch is a paid mutator transaction binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) CalculateRewardPerBatch() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.CalculateRewardPerBatch(&_Ipolygonrollupmanager.TransactOpts)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0xc5b4fdb6.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName, bytes initializeBytesCustomChain) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName, initializeBytesCustomChain)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0xc5b4fdb6.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName, bytes initializeBytesCustomChain) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.CreateNewRollup(&_Ipolygonrollupmanager.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName, initializeBytesCustomChain)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0xc5b4fdb6.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName, bytes initializeBytesCustomChain) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string, initializeBytesCustomChain []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.CreateNewRollup(&_Ipolygonrollupmanager.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName, initializeBytesCustomChain)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.DeactivateEmergencyState(&_Ipolygonrollupmanager.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.DeactivateEmergencyState(&_Ipolygonrollupmanager.TransactOpts)
}

// GetBatchFee is a paid mutator transaction binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetBatchFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getBatchFee")
}

// GetBatchFee is a paid mutator transaction binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetBatchFee() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetBatchFee(&_Ipolygonrollupmanager.TransactOpts)
}

// GetBatchFee is a paid mutator transaction binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetBatchFee() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetBatchFee(&_Ipolygonrollupmanager.TransactOpts)
}

// GetForcedBatchFee is a paid mutator transaction binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetForcedBatchFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getForcedBatchFee")
}

// GetForcedBatchFee is a paid mutator transaction binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetForcedBatchFee() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetForcedBatchFee(&_Ipolygonrollupmanager.TransactOpts)
}

// GetForcedBatchFee is a paid mutator transaction binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() returns(uint256)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetForcedBatchFee() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetForcedBatchFee(&_Ipolygonrollupmanager.TransactOpts)
}

// GetInputPessimisticBytes is a paid mutator transaction binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes customChainData) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetInputPessimisticBytes(opts *bind.TransactOpts, rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getInputPessimisticBytes", rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, customChainData)
}

// GetInputPessimisticBytes is a paid mutator transaction binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes customChainData) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetInputPessimisticBytes(&_Ipolygonrollupmanager.TransactOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, customChainData)
}

// GetInputPessimisticBytes is a paid mutator transaction binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes customChainData) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetInputPessimisticBytes(&_Ipolygonrollupmanager.TransactOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, customChainData)
}

// GetInputSnarkBytes is a paid mutator transaction binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetInputSnarkBytes(opts *bind.TransactOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a paid mutator transaction binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetInputSnarkBytes(&_Ipolygonrollupmanager.TransactOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a paid mutator transaction binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) returns(bytes)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetInputSnarkBytes(&_Ipolygonrollupmanager.TransactOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a paid mutator transaction binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetLastVerifiedBatch(opts *bind.TransactOpts, rollupID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getLastVerifiedBatch", rollupID)
}

// GetLastVerifiedBatch is a paid mutator transaction binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetLastVerifiedBatch(rollupID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetLastVerifiedBatch(&_Ipolygonrollupmanager.TransactOpts, rollupID)
}

// GetLastVerifiedBatch is a paid mutator transaction binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetLastVerifiedBatch(rollupID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetLastVerifiedBatch(&_Ipolygonrollupmanager.TransactOpts, rollupID)
}

// GetRollupBatchNumToStateRoot is a paid mutator transaction binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetRollupBatchNumToStateRoot(opts *bind.TransactOpts, rollupID uint32, batchNum uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getRollupBatchNumToStateRoot", rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a paid mutator transaction binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetRollupBatchNumToStateRoot(&_Ipolygonrollupmanager.TransactOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a paid mutator transaction binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetRollupBatchNumToStateRoot(&_Ipolygonrollupmanager.TransactOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a paid mutator transaction binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) GetRollupExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "getRollupExitRoot")
}

// GetRollupExitRoot is a paid mutator transaction binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) GetRollupExitRoot() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetRollupExitRoot(&_Ipolygonrollupmanager.TransactOpts)
}

// GetRollupExitRoot is a paid mutator transaction binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() returns(bytes32)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) GetRollupExitRoot() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.GetRollupExitRoot(&_Ipolygonrollupmanager.TransactOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) LastDeactivatedEmergencyStateTimestamp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "lastDeactivatedEmergencyStateTimestamp")
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) LastDeactivatedEmergencyStateTimestamp() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Ipolygonrollupmanager.TransactOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a paid mutator transaction binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) LastDeactivatedEmergencyStateTimestamp() (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Ipolygonrollupmanager.TransactOpts)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.ObsoleteRollupType(&_Ipolygonrollupmanager.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.ObsoleteRollupType(&_Ipolygonrollupmanager.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.OnSequenceBatches(&_Ipolygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.OnSequenceBatches(&_Ipolygonrollupmanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.RollbackBatches(&_Ipolygonrollupmanager.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.RollbackBatches(&_Ipolygonrollupmanager.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.SetBatchFee(&_Ipolygonrollupmanager.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.SetBatchFee(&_Ipolygonrollupmanager.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.UpdateRollup(&_Ipolygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.UpdateRollup(&_Ipolygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.UpdateRollupByRollupAdmin(&_Ipolygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.UpdateRollupByRollupAdmin(&_Ipolygonrollupmanager.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.VerifyBatchesTrustedAggregator(&_Ipolygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.VerifyBatchesTrustedAggregator(&_Ipolygonrollupmanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes customChainData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, customChainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes customChainData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.VerifyPessimisticTrustedAggregator(&_Ipolygonrollupmanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, customChainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes customChainData) returns()
func (_Ipolygonrollupmanager *IpolygonrollupmanagerTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, customChainData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanager.Contract.VerifyPessimisticTrustedAggregator(&_Ipolygonrollupmanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, customChainData)
}

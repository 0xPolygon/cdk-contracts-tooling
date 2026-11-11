// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ipolygonrollupmanagerpessimistic

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

// IpolygonrollupmanagerpessimisticMetaData contains all meta data concerning the Ipolygonrollupmanagerpessimistic contract.
var IpolygonrollupmanagerpessimisticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyChainsWithPessimisticProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIPolygonRollupManagerPessimistic.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManagerPessimistic.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"selectedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IpolygonrollupmanagerpessimisticABI is the input ABI used to generate the binding from.
// Deprecated: Use IpolygonrollupmanagerpessimisticMetaData.ABI instead.
var IpolygonrollupmanagerpessimisticABI = IpolygonrollupmanagerpessimisticMetaData.ABI

// Ipolygonrollupmanagerpessimistic is an auto generated Go binding around an Ethereum contract.
type Ipolygonrollupmanagerpessimistic struct {
	IpolygonrollupmanagerpessimisticCaller     // Read-only binding to the contract
	IpolygonrollupmanagerpessimisticTransactor // Write-only binding to the contract
	IpolygonrollupmanagerpessimisticFilterer   // Log filterer for contract events
}

// IpolygonrollupmanagerpessimisticCaller is an auto generated read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpessimisticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpessimisticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpessimisticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpessimisticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IpolygonrollupmanagerpessimisticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IpolygonrollupmanagerpessimisticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IpolygonrollupmanagerpessimisticSession struct {
	Contract     *Ipolygonrollupmanagerpessimistic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerpessimisticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IpolygonrollupmanagerpessimisticCallerSession struct {
	Contract *IpolygonrollupmanagerpessimisticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// IpolygonrollupmanagerpessimisticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IpolygonrollupmanagerpessimisticTransactorSession struct {
	Contract     *IpolygonrollupmanagerpessimisticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// IpolygonrollupmanagerpessimisticRaw is an auto generated low-level Go binding around an Ethereum contract.
type IpolygonrollupmanagerpessimisticRaw struct {
	Contract *Ipolygonrollupmanagerpessimistic // Generic contract binding to access the raw methods on
}

// IpolygonrollupmanagerpessimisticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpessimisticCallerRaw struct {
	Contract *IpolygonrollupmanagerpessimisticCaller // Generic read-only contract binding to access the raw methods on
}

// IpolygonrollupmanagerpessimisticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IpolygonrollupmanagerpessimisticTransactorRaw struct {
	Contract *IpolygonrollupmanagerpessimisticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIpolygonrollupmanagerpessimistic creates a new instance of Ipolygonrollupmanagerpessimistic, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpessimistic(address common.Address, backend bind.ContractBackend) (*Ipolygonrollupmanagerpessimistic, error) {
	contract, err := bindIpolygonrollupmanagerpessimistic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ipolygonrollupmanagerpessimistic{IpolygonrollupmanagerpessimisticCaller: IpolygonrollupmanagerpessimisticCaller{contract: contract}, IpolygonrollupmanagerpessimisticTransactor: IpolygonrollupmanagerpessimisticTransactor{contract: contract}, IpolygonrollupmanagerpessimisticFilterer: IpolygonrollupmanagerpessimisticFilterer{contract: contract}}, nil
}

// NewIpolygonrollupmanagerpessimisticCaller creates a new read-only instance of Ipolygonrollupmanagerpessimistic, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpessimisticCaller(address common.Address, caller bind.ContractCaller) (*IpolygonrollupmanagerpessimisticCaller, error) {
	contract, err := bindIpolygonrollupmanagerpessimistic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpessimisticCaller{contract: contract}, nil
}

// NewIpolygonrollupmanagerpessimisticTransactor creates a new write-only instance of Ipolygonrollupmanagerpessimistic, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpessimisticTransactor(address common.Address, transactor bind.ContractTransactor) (*IpolygonrollupmanagerpessimisticTransactor, error) {
	contract, err := bindIpolygonrollupmanagerpessimistic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpessimisticTransactor{contract: contract}, nil
}

// NewIpolygonrollupmanagerpessimisticFilterer creates a new log filterer instance of Ipolygonrollupmanagerpessimistic, bound to a specific deployed contract.
func NewIpolygonrollupmanagerpessimisticFilterer(address common.Address, filterer bind.ContractFilterer) (*IpolygonrollupmanagerpessimisticFilterer, error) {
	contract, err := bindIpolygonrollupmanagerpessimistic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IpolygonrollupmanagerpessimisticFilterer{contract: contract}, nil
}

// bindIpolygonrollupmanagerpessimistic binds a generic wrapper to an already deployed contract.
func bindIpolygonrollupmanagerpessimistic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IpolygonrollupmanagerpessimisticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanagerpessimistic.Contract.IpolygonrollupmanagerpessimisticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.IpolygonrollupmanagerpessimisticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.IpolygonrollupmanagerpessimisticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ipolygonrollupmanagerpessimistic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.contract.Transact(opts, method, params...)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.CalculateRewardPerBatch(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.CalculateRewardPerBatch(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetBatchFee() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetBatchFee(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetBatchFee() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetBatchFee(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetForcedBatchFee() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetForcedBatchFee(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetForcedBatchFee(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetInputPessimisticBytes(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetInputPessimisticBytes(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetInputSnarkBytes(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetInputSnarkBytes(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetLastVerifiedBatch(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetLastVerifiedBatch(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetRollupBatchNumToStateRoot(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetRollupBatchNumToStateRoot(&_Ipolygonrollupmanagerpessimistic.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) GetRollupExitRoot() ([32]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetRollupExitRoot(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.GetRollupExitRoot(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Ipolygonrollupmanagerpessimistic.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.LastDeactivatedEmergencyStateTimestamp(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.LastDeactivatedEmergencyStateTimestamp(&_Ipolygonrollupmanagerpessimistic.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.ActivateEmergencyState(&_Ipolygonrollupmanagerpessimistic.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.ActivateEmergencyState(&_Ipolygonrollupmanagerpessimistic.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.AddExistingRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.AddExistingRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.AddNewRollupType(&_Ipolygonrollupmanagerpessimistic.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.AddNewRollupType(&_Ipolygonrollupmanagerpessimistic.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.CreateNewRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.CreateNewRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.DeactivateEmergencyState(&_Ipolygonrollupmanagerpessimistic.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.DeactivateEmergencyState(&_Ipolygonrollupmanagerpessimistic.TransactOpts)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.ObsoleteRollupType(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.ObsoleteRollupType(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.OnSequenceBatches(&_Ipolygonrollupmanagerpessimistic.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.OnSequenceBatches(&_Ipolygonrollupmanagerpessimistic.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.RollbackBatches(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.RollbackBatches(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.SetBatchFee(&_Ipolygonrollupmanagerpessimistic.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.SetBatchFee(&_Ipolygonrollupmanagerpessimistic.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.UpdateRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.UpdateRollup(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.UpdateRollupByRollupAdmin(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.UpdateRollupByRollupAdmin(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.VerifyBatchesTrustedAggregator(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.VerifyBatchesTrustedAggregator(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.VerifyPessimisticTrustedAggregator(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Ipolygonrollupmanagerpessimistic *IpolygonrollupmanagerpessimisticTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Ipolygonrollupmanagerpessimistic.Contract.VerifyPessimisticTrustedAggregator(&_Ipolygonrollupmanagerpessimistic.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iagglayermanager

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

// IagglayermanagerMetaData contains all meta data concerning the Iagglayermanager contract.
var IagglayermanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewRollupTypeMustBePessimisticOrALGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"selectedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IagglayermanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IagglayermanagerMetaData.ABI instead.
var IagglayermanagerABI = IagglayermanagerMetaData.ABI

// Iagglayermanager is an auto generated Go binding around an Ethereum contract.
type Iagglayermanager struct {
	IagglayermanagerCaller     // Read-only binding to the contract
	IagglayermanagerTransactor // Write-only binding to the contract
	IagglayermanagerFilterer   // Log filterer for contract events
}

// IagglayermanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IagglayermanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayermanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IagglayermanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayermanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IagglayermanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IagglayermanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IagglayermanagerSession struct {
	Contract     *Iagglayermanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IagglayermanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IagglayermanagerCallerSession struct {
	Contract *IagglayermanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IagglayermanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IagglayermanagerTransactorSession struct {
	Contract     *IagglayermanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IagglayermanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IagglayermanagerRaw struct {
	Contract *Iagglayermanager // Generic contract binding to access the raw methods on
}

// IagglayermanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IagglayermanagerCallerRaw struct {
	Contract *IagglayermanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IagglayermanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IagglayermanagerTransactorRaw struct {
	Contract *IagglayermanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIagglayermanager creates a new instance of Iagglayermanager, bound to a specific deployed contract.
func NewIagglayermanager(address common.Address, backend bind.ContractBackend) (*Iagglayermanager, error) {
	contract, err := bindIagglayermanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iagglayermanager{IagglayermanagerCaller: IagglayermanagerCaller{contract: contract}, IagglayermanagerTransactor: IagglayermanagerTransactor{contract: contract}, IagglayermanagerFilterer: IagglayermanagerFilterer{contract: contract}}, nil
}

// NewIagglayermanagerCaller creates a new read-only instance of Iagglayermanager, bound to a specific deployed contract.
func NewIagglayermanagerCaller(address common.Address, caller bind.ContractCaller) (*IagglayermanagerCaller, error) {
	contract, err := bindIagglayermanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayermanagerCaller{contract: contract}, nil
}

// NewIagglayermanagerTransactor creates a new write-only instance of Iagglayermanager, bound to a specific deployed contract.
func NewIagglayermanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IagglayermanagerTransactor, error) {
	contract, err := bindIagglayermanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IagglayermanagerTransactor{contract: contract}, nil
}

// NewIagglayermanagerFilterer creates a new log filterer instance of Iagglayermanager, bound to a specific deployed contract.
func NewIagglayermanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IagglayermanagerFilterer, error) {
	contract, err := bindIagglayermanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IagglayermanagerFilterer{contract: contract}, nil
}

// bindIagglayermanager binds a generic wrapper to an already deployed contract.
func bindIagglayermanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IagglayermanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayermanager *IagglayermanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayermanager.Contract.IagglayermanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayermanager *IagglayermanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.IagglayermanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayermanager *IagglayermanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.IagglayermanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iagglayermanager *IagglayermanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iagglayermanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iagglayermanager *IagglayermanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iagglayermanager *IagglayermanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.contract.Transact(opts, method, params...)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Iagglayermanager.Contract.CalculateRewardPerBatch(&_Iagglayermanager.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Iagglayermanager.Contract.CalculateRewardPerBatch(&_Iagglayermanager.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerSession) GetBatchFee() (*big.Int, error) {
	return _Iagglayermanager.Contract.GetBatchFee(&_Iagglayermanager.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetBatchFee() (*big.Int, error) {
	return _Iagglayermanager.Contract.GetBatchFee(&_Iagglayermanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Iagglayermanager.Contract.GetForcedBatchFee(&_Iagglayermanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Iagglayermanager.Contract.GetForcedBatchFee(&_Iagglayermanager.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Iagglayermanager.Contract.GetInputPessimisticBytes(&_Iagglayermanager.CallOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 selectedGlobalExitRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetInputPessimisticBytes(rollupID uint32, selectedGlobalExitRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Iagglayermanager.Contract.GetInputPessimisticBytes(&_Iagglayermanager.CallOpts, rollupID, selectedGlobalExitRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Iagglayermanager.Contract.GetInputSnarkBytes(&_Iagglayermanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Iagglayermanager.Contract.GetInputSnarkBytes(&_Iagglayermanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Iagglayermanager *IagglayermanagerCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Iagglayermanager *IagglayermanagerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Iagglayermanager.Contract.GetLastVerifiedBatch(&_Iagglayermanager.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Iagglayermanager.Contract.GetLastVerifiedBatch(&_Iagglayermanager.CallOpts, rollupID)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Iagglayermanager.Contract.GetRollupBatchNumToStateRoot(&_Iagglayermanager.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Iagglayermanager.Contract.GetRollupBatchNumToStateRoot(&_Iagglayermanager.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Iagglayermanager.Contract.GetRollupExitRoot(&_Iagglayermanager.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Iagglayermanager *IagglayermanagerCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Iagglayermanager.Contract.GetRollupExitRoot(&_Iagglayermanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Iagglayermanager *IagglayermanagerCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Iagglayermanager.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Iagglayermanager *IagglayermanagerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Iagglayermanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Iagglayermanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Iagglayermanager *IagglayermanagerCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Iagglayermanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Iagglayermanager.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayermanager.Contract.ActivateEmergencyState(&_Iagglayermanager.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayermanager.Contract.ActivateEmergencyState(&_Iagglayermanager.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Iagglayermanager *IagglayermanagerSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AddExistingRollup(&_Iagglayermanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AddExistingRollup(&_Iagglayermanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Iagglayermanager *IagglayermanagerSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AddNewRollupType(&_Iagglayermanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 initRoot, string description, bytes32 programVKey) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, initRoot [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AddNewRollupType(&_Iagglayermanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, initRoot, description, programVKey)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Iagglayermanager *IagglayermanagerSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AttachAggchainToAL(&_Iagglayermanager.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.AttachAggchainToAL(&_Iagglayermanager.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayermanager.Contract.DeactivateEmergencyState(&_Iagglayermanager.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Iagglayermanager.Contract.DeactivateEmergencyState(&_Iagglayermanager.TransactOpts)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.ObsoleteRollupType(&_Iagglayermanager.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.ObsoleteRollupType(&_Iagglayermanager.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Iagglayermanager *IagglayermanagerTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Iagglayermanager *IagglayermanagerSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.OnSequenceBatches(&_Iagglayermanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Iagglayermanager *IagglayermanagerTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.OnSequenceBatches(&_Iagglayermanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Iagglayermanager *IagglayermanagerSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.RollbackBatches(&_Iagglayermanager.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.RollbackBatches(&_Iagglayermanager.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Iagglayermanager *IagglayermanagerSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.SetBatchFee(&_Iagglayermanager.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.SetBatchFee(&_Iagglayermanager.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Iagglayermanager *IagglayermanagerSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.UpdateRollup(&_Iagglayermanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.UpdateRollup(&_Iagglayermanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.UpdateRollupByRollupAdmin(&_Iagglayermanager.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.UpdateRollupByRollupAdmin(&_Iagglayermanager.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Iagglayermanager *IagglayermanagerSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.VerifyBatchesTrustedAggregator(&_Iagglayermanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.VerifyBatchesTrustedAggregator(&_Iagglayermanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Iagglayermanager *IagglayermanagerTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Iagglayermanager *IagglayermanagerSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.VerifyPessimisticTrustedAggregator(&_Iagglayermanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Iagglayermanager *IagglayermanagerTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Iagglayermanager.Contract.VerifyPessimisticTrustedAggregator(&_Iagglayermanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

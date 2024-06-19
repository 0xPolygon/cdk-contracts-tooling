// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmtestnetclearstorage

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

// PolygonZkEVMBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	Timestamp          uint64
	MinForcedTimestamp uint64
}

// PolygonZkEVMForcedBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMForcedBatchData struct {
	Transactions       []byte
	GlobalExitRoot     [32]byte
	MinForcedTimestamp uint64
}

// PolygonZkEVMInitializePackedParameters is an auto generated low-level Go binding around an user-defined struct.
type PolygonZkEVMInitializePackedParameters struct {
	Admin                    common.Address
	TrustedSequencer         common.Address
	PendingStateTimeout      uint64
	TrustedAggregator        common.Address
	TrustedAggregatorTimeout uint64
}

// PolygonzkevmtestnetclearstorageMetaData contains all meta data concerning the Polygonzkevmtestnetclearstorage contract.
var PolygonzkevmtestnetclearstorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyUpdated\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateZkEVMVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61014060405234801562000011575f80fd5b5060405162005fda38038062005fda8339810160408190526200003491620000a2565b6001600160a01b0395861660c05293851660805291841660a05290921660e0526001600160401b039182166101005216610120526200012a565b6001600160a01b038116811462000083575f80fd5b50565b80516001600160401b03811681146200009d575f80fd5b919050565b5f805f805f8060c08789031215620000b8575f80fd5b8651620000c5816200006e565b6020880151909650620000d8816200006e565b6040880151909550620000eb816200006e565b6060880151909450620000fe816200006e565b92506200010e6080880162000086565b91506200011e60a0880162000086565b90509295509295509295565b60805160a05160c05160e0516101005161012051615dea620001f05f395f81816106f801528181610e49015261319301525f81816108640152610e1f01525f818161082a01528181611dd6015281816138240152614c7601525f81816109cf01528181610fbb01528181611186015281816119e3015281816121a401528181613a05015261474e01525f8181610a7c015281816140b7015261450501525f818161091f01528181611da401528181612689015281816139da01526141a30152615dea5ff3fe608060405234801561000f575f80fd5b50600436106103bf575f3560e01c8063837a4738116101f5578063c754c7ed11610114578063e7a7ed02116100a9578063f14916d611610079578063f14916d614610ade578063f2fde38b14610af1578063f851a44014610b04578063f8b823e414610b24575f80fd5b8063e7a7ed0214610a47578063e8bf92ed14610a77578063eaeb077b14610a9e578063ed6b010414610ab1575f80fd5b8063d2e129f9116100e4578063d2e129f9146109f1578063d8d1091b14610a04578063d939b31514610a17578063dbc1697614610a3f575f80fd5b8063c754c7ed1461095c578063c89e42df14610988578063cfa8ed471461099b578063d02103ca146109ca575f80fd5b8063a3c573eb1161018a578063b4d63f581161015a578063b4d63f58146108b4578063b6b0b0971461091a578063ba58ae3914610941578063c0ed84e014610954575f80fd5b8063a3c573eb14610825578063ada8f9191461084c578063adc879e91461085f578063afd23cbe14610886575f80fd5b806399f5634e116101c557806399f5634e146107e45780639aa972a3146107ec5780639c9f3dfe146107ff578063a066215c14610812575f80fd5b8063837a47381461071a578063841b24d71461078e5780638c3d7301146107be5780638da5cb5b146107c6575f80fd5b80634a910e6a116102e1578063621dd41111610276578063715018a611610246578063715018a6146106c45780637215541a146106cc5780637fcb3653146106df578063831c7ead146106f3575f80fd5b8063621dd4111461062857806369bd38a01461063b5780636b8616ce146106925780636ff512cc146106b1575f80fd5b806354fd4d50116102b157806354fd4d50146105fc5780635e9145c9146106055780635ec91958146106185780636046916914610620575f80fd5b80634a910e6a146105a15780634e487706146105b45780635392c5e0146105c7578063542028d5146105f4575f80fd5b80632b0006fa11610357578063423fa85611610327578063423fa856146105255780634560526714610545578063458c04771461056d5780634a1a89a714610581575f80fd5b80632b0006fa146104d95780632c1f816a146104ec578063383b3be8146104ff578063394218e914610512575f80fd5b806319d8ac611161039257806319d8ac6114610441578063220d789914610455578063267822471461046857806329878983146104ad575f80fd5b80630a0d9fbe146103c3578063107bf28c146103fa57806315064c961461040f5780631816b7e51461042c575b5f80fd5b606f546103dc90610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b610402610b2d565b6040516103f19190615290565b606f5461041c9060ff1681565b60405190151581526020016103f1565b61043f61043a3660046152a9565b610bb9565b005b6073546103dc9067ffffffffffffffff1681565b6104026104633660046152e1565b610cd1565b607b546104889073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103f1565b6074546104889068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b61043f6104e7366004615341565b610ea7565b61043f6104fa3660046153a5565b611071565b61041c61050d36600461541a565b611279565b61043f61052036600461541a565b6112ce565b6073546103dc9068010000000000000000900467ffffffffffffffff1681565b6073546103dc90700100000000000000000000000000000000900467ffffffffffffffff1681565b6079546103dc9067ffffffffffffffff1681565b6079546103dc9068010000000000000000900467ffffffffffffffff1681565b61043f6105af36600461541a565b611452565b61043f6105c236600461541a565b611505565b6105e66105d536600461541a565b60756020525f908152604090205481565b6040519081526020016103f1565b610402611689565b6105e6607c5481565b61043f61061336600461549e565b611696565b61043f611e8f565b6105e6611f8e565b61043f610636366004615341565b611fa3565b607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c01000000000006978000000000000000000000000000000000000000001790555f607c819055607d819055607e55005b6105e66106a036600461541a565b60716020525f908152604090205481565b61043f6106bf3660046154ee565b612321565b61043f6123f6565b61043f6106da36600461541a565b612409565b6074546103dc9067ffffffffffffffff1681565b6103dc7f000000000000000000000000000000000000000000000000000000000000000081565b610762610728366004615507565b60786020525f908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff95861681529490931660208501529183015260608201526080016103f1565b6079546103dc907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61043f612576565b60335473ffffffffffffffffffffffffffffffffffffffff16610488565b6105e6612642565b61043f6107fa3660046153a5565b612795565b61043f61080d36600461541a565b612845565b61043f61082036600461541a565b6129c1565b6104887f000000000000000000000000000000000000000000000000000000000000000081565b61043f61085a3660046154ee565b612ac7565b6103dc7f000000000000000000000000000000000000000000000000000000000000000081565b606f546108a1906901000000000000000000900461ffff1681565b60405161ffff90911681526020016103f1565b6108f46108c236600461541a565b60726020525f90815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff92831660208501529116908201526060016103f1565b6104887f000000000000000000000000000000000000000000000000000000000000000081565b61041c61094f366004615507565b612b8b565b6103dc612c13565b607b546103dc9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b61043f6109963660046155f2565b612c66565b606f54610488906b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6104887f000000000000000000000000000000000000000000000000000000000000000081565b61043f6109ff366004615662565b612cf3565b61043f610a1236600461570d565b613236565b6079546103dc90700100000000000000000000000000000000900467ffffffffffffffff1681565b61043f6137d1565b6073546103dc907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6104887f000000000000000000000000000000000000000000000000000000000000000081565b61043f610aac36600461574c565b6138a5565b607b5461041c907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b61043f610aec3660046154ee565b613c95565b61043f610aff3660046154ee565b613d67565b607a546104889073ffffffffffffffffffffffffffffffffffffffff1681565b6105e660705481565b60778054610b3a90615794565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6690615794565b8015610bb15780601f10610b8857610100808354040283529160200191610bb1565b820191905f5260205f20905b815481529060010190602001808311610b9457829003601f168201915b505050505081565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610c0a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff161080610c2357506103ff8161ffff16115b15610c5a576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086165f818152607260205260408082205493881682529020546060929115801590610d04575081155b15610d3b576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610d72576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d7b84612b8b565b610db1576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314610f04576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f12868686868686613e1b565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f9081526075602052604090208390556079541615610f8c57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b158015611011575f80fd5b505af1158015611023573d5f803e3d5ffd5b505060405184815233925067ffffffffffffffff871691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe906020015b60405180910390a3505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1633146110ce576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110dd878787878787876141d6565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607560205260409020839055607954161561115757607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b1580156111dc575f80fd5b505af11580156111ee573d5f803e3d5ffd5b50506079805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051828152339067ffffffffffffffff8616907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a350505050505050565b60795467ffffffffffffffff8281165f90815260786020526040812054909242926112bc9270010000000000000000000000000000000090920481169116615812565b67ffffffffffffffff16111592915050565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461131f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611366576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166113d55760795467ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116908216106113d5576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001610cc6565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1633146114f957606f5460ff16156114ba576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114c381611279565b6114f9576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61150281614605565b50565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611556576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff8216111561159d576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1661160857607b5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610611608576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001610cc6565b60768054610b3a90615794565b606f5460ff16156116d3576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f546b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611733576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f81900361176e576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156117aa576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000820481165f81815260726020526040812054838516949293700100000000000000000000000000000000909304909216919082905b86811015611bf2575f8a8a838181106118105761181061583a565b90506020028101906118229190615867565b61182b906158a3565b8051805160209091012060608201519192509067ffffffffffffffff16156119a057856118578161592d565b9650505f81836020015184606001516040516020016118ae93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f90815260719093529120549091508114611936576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8088165f90815260716020526040808220919091556060850151908501519082169116101561199a576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611ada565b602082015115801590611a64575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303815f875af1158015611a3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a629190615953565b155b15611a9b576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516201d4c01015611ada576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8767ffffffffffffffff16826040015167ffffffffffffffff161080611b0d575042826040015167ffffffffffffffff16115b15611b44576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c016040516020818303038152906040528051906020012094508160400151975050508080611bea9061596a565b9150506117f5565b50611bfd8685615812565b60735490945067ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169084161115611c66576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f611c7182856159a1565b611c859067ffffffffffffffff16886159c2565b604080516060810182528581524267ffffffffffffffff908116602080840191825260738054680100000000000000009081900485168688019081528d86165f8181526072909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c8416911617930292909217905590915082811690851614611d7a57607380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b611dcc333083607054611d8d91906159d5565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190614812565b611dd46148f4565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611e39575f80fd5b505af1158015611e4b573d5f803e3d5ffd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce91505f90a250505050505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611ee0576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16611f3c576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f905f90a1565b5f6070546064611f9e91906159d5565b905090565b606f5460ff1615611fe0576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff8581165f90815260726020526040902060010154429261202c92780100000000000000000000000000000000000000000000000090910481169116615812565b67ffffffffffffffff16111561206e576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e861207b86866159a1565b67ffffffffffffffff1611156120bd576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120cb868686868686613e1b565b6120d4846149a3565b607954700100000000000000000000000000000000900467ffffffffffffffff165f0361221557607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8681169182179092555f908152607560205260409020839055607954161561217557607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b1580156121fa575f80fd5b505af115801561220c573d5f803e3d5ffd5b505050506122e3565b61221d6148f4565b6079805467ffffffffffffffff16905f6122368361592d565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815287831660208083019182528284018981526060840189815260795487165f908152607890935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051828152339067ffffffffffffffff8616907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596690602001611061565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612372576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b01000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610cc6565b6123fe614b7d565b6124075f614bfe565b565b60335473ffffffffffffffffffffffffffffffffffffffff16331461256e575f612431612c13565b90508067ffffffffffffffff168267ffffffffffffffff1611612480576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000909104811690831611806124c5575067ffffffffffffffff8083165f9081526072602052604090206001015416155b156124fc576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8083165f90815260726020526040902060010154429161252a9162093a809116615812565b67ffffffffffffffff16111561256c576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b611502614c74565b607b5473ffffffffffffffffffffffffffffffffffffffff1633146125c7576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b54607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156126ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126f29190615953565b90505f6126fd612c13565b60735467ffffffffffffffff68010000000000000000820481169161275591700100000000000000000000000000000000820481169178010000000000000000000000000000000000000000000000009004166159a1565b61275f9190615812565b61276991906159a1565b67ffffffffffffffff169050805f03612784575f9250505090565b61278e8183615a19565b9250505090565b606f5460ff16156127d2576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6127e1878787878787876141d6565b67ffffffffffffffff84165f908152607560209081526040918290205482519081529081018490527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a161283c614c74565b50505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612896576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156128dd576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166129445760795467ffffffffffffffff700100000000000000000000000000000000909104811690821610612944576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607980547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001610cc6565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612a12576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff161115612a59576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001610cc6565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612b18576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610cc6565b5f67ffffffff0000000167ffffffffffffffff8316108015612bc2575067ffffffff00000001604083901c67ffffffffffffffff16105b8015612be3575067ffffffff00000001608083901c67ffffffffffffffff16105b8015612bfa575067ffffffff0000000160c083901c105b15612c0757506001919050565b505f919050565b919050565b6079545f9067ffffffffffffffff1615612c55575060795467ffffffffffffffff9081165f908152607860205260409020546801000000000000000090041690565b5060745467ffffffffffffffff1690565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612cb7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076612cc38282615a79565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610cc69190615290565b5f54610100900460ff1615808015612d1157505f54600160ff909116105b80612d2a5750303b158015612d2a57505f5460ff166001145b612dbb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015612e17575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b612e2460208801886154ee565b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055612e7960408801602089016154ee565b606f805473ffffffffffffffffffffffffffffffffffffffff929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff909216919091179055612ede60808801606089016154ee565b6074805473ffffffffffffffffffffffffffffffffffffffff9290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790555f805260756020527ff9e3fbf150b7a0077118526f473c53cb4734f166167e2c6213e3567dd390b4ad8690556076612f688682615a79565b506077612f758582615a79565b5062093a80612f8a6060890160408a0161541a565b67ffffffffffffffff161115612fcc576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612fdc606088016040890161541a565b6079805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a8061303e60a0890160808a0161541a565b67ffffffffffffffff161115613080576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61309060a088016080890161541a565b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607055606f80547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c010000000000069780000000000000000000000000000000000000000017905561316f614cf7565b7fed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd65f7f000000000000000000000000000000000000000000000000000000000000000085856040516131c49493929190615bd8565b60405180910390a1801561283c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1615613293576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff16156132d0576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f81900361330b576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115613347576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff78010000000000000000000000000000000000000000000000008204811691613392918491700100000000000000000000000000000000900416615c0f565b11156133ca576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff68010000000000000000820481165f8181526072602052604081205491937001000000000000000000000000000000009004909216915b8481101561366f575f8787838181106134285761342861583a565b905060200281019061343a9190615c22565b61344390615c54565b90508361344f8161592d565b825180516020918201208185015160408087015190519499509194505f936134b09386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260719093529120549091508114613538576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f9081526071602052604081205561355c6001896159c2565b84036135cb5742607b60149054906101000a900467ffffffffffffffff1684604001516135899190615812565b67ffffffffffffffff1611156135cb576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c0160405160208183030381529060405280519060200120945050505080806136679061596a565b91505061340d565b5061367a8484615812565b6073805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589165f818152607290935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613822576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dbc169766040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613887575f80fd5b505af1158015613899573d5f803e3d5ffd5b50505050612407614d96565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1615613902576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff161561393f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613948611f8e565b905081811115613984576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113888311156139c0576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613a0273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084614812565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a6c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a909190615953565b60738054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16906018613aca8361592d565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550508484604051613b01929190615ccd565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291815281516020928301206073547801000000000000000000000000000000000000000000000000900467ffffffffffffffff165f9081526071909352912055323303613c2f576073546040805183815233602082015260609181018290525f91810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2613c8e565b607360189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93182338888604051613c859493929190615cdc565b60405180910390a25b5050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613ce6576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607480547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001610cc6565b613d6f614b7d565b73ffffffffffffffffffffffffffffffffffffffff8116613e12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401612db2565b61150281614bfe565b5f80613e25612c13565b905067ffffffffffffffff881615613ef45760795467ffffffffffffffff9081169089161115613e81576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8089165f908152607860205260409020600281015481549094509091898116680100000000000000009092041614613eee576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613f94565b67ffffffffffffffff87165f90815260756020526040902054915081613f46576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff161115613f94576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168667ffffffffffffffff1611613fe1576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613fef8888888689610cd1565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516140239190615d11565b602060405180830381855afa15801561403e573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906140619190615953565b61406b9190615d22565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a916140ed91899190600401615d35565b602060405180830381865afa158015614108573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061412c9190615d6f565b614162576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6141ca33614170858b6159a1565b67ffffffffffffffff16614182612642565b61418c91906159d5565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190614e24565b50505050505050505050565b5f67ffffffffffffffff8816156142a25760795467ffffffffffffffff9081169089161115614231576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f90815260786020526040902060028101548154909288811668010000000000000000909204161461429c576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061433d565b5067ffffffffffffffff85165f90815260756020526040902054806142f3576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745467ffffffffffffffff908116908716111561433d576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff908116908816118061436f57508767ffffffffffffffff168767ffffffffffffffff1611155b80614396575060795467ffffffffffffffff68010000000000000000909104811690881611155b156143cd576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f9081526078602052604090205468010000000000000000900481169086161461442f576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61443d8787878588610cd1565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516144719190615d11565b602060405180830381855afa15801561448c573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906144af9190615953565b6144b99190615d22565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a9161453b91889190600401615d35565b602060405180830381865afa158015614556573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061457a9190615d6f565b6145b0576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f908152607860205260409020600201548590036141ca576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff68010000000000000000909104811690821611158061463f575060795467ffffffffffffffff908116908216115b15614676576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f81815260786020908152604080832080546074805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260759094529382902092909255607980547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d906024015f604051808303815f87803b1580156147a4575f80fd5b505af11580156147b6573d5f803e3d5ffd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e846002015460405161480591815260200190565b60405180910390a3505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526148ee9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614e7f565b50505050565b60795467ffffffffffffffff680100000000000000008204811691161115612407576079545f9061493c9068010000000000000000900467ffffffffffffffff166001615812565b905061494781611279565b15611502576079545f9060029061496990849067ffffffffffffffff166159a1565b6149739190615d8e565b61497d9083615812565b905061498881611279565b1561499a5761499681614605565b5050565b61499682614605565b5f6149ac612c13565b9050815f806149bb84846159a1565b606f5467ffffffffffffffff91821692505f916149de91610100900416426159c2565b90505b8467ffffffffffffffff168467ffffffffffffffff1614614a685767ffffffffffffffff8085165f9081526072602052604090206001810154909116821015614a4657600181015468010000000000000000900467ffffffffffffffff169450614a62565b614a5086866159a1565b67ffffffffffffffff16935050614a68565b506149e1565b5f614a7384846159c2565b905083811015614aca57808403600c8111614a8e5780614a91565b600c5b9050806103e80a81606f60099054906101000a900461ffff1661ffff160a6070540281614ac057614ac06159ec565b0460705550614b39565b838103600c8111614adb5780614ade565b600c5b90505f816103e80a82606f60099054906101000a900461ffff1661ffff160a670de0b6b3a76400000281614b1457614b146159ec565b04905080607054670de0b6b3a76400000281614b3257614b326159ec565b0460705550505b683635c9adc5dea000006070541115614b5e57683635c9adc5dea0000060705561283c565b633b9aca00607054101561283c57633b9aca0060705550505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314612407576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401612db2565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614cd9575f80fd5b505af1158015614ceb573d5f803e3d5ffd5b50505050612407614f8a565b5f54610100900460ff16614d8d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401612db2565b61240733614bfe565b606f5460ff16614dd2576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052614e7a9084907fa9059cbb000000000000000000000000000000000000000000000000000000009060640161486c565b505050565b5f614ee0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661501c9092919063ffffffff16565b805190915015614e7a5780806020019051810190614efe9190615d6f565b614e7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401612db2565b606f5460ff1615614fc7576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b606061502a84845f85615032565b949350505050565b6060824710156150c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401612db2565b5f808673ffffffffffffffffffffffffffffffffffffffff1685876040516150ec9190615d11565b5f6040518083038185875af1925050503d805f8114615126576040519150601f19603f3d011682016040523d82523d5f602084013e61512b565b606091505b509150915061513c87838387615147565b979650505050505050565b606083156151dc5782515f036151d55773ffffffffffffffffffffffffffffffffffffffff85163b6151d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401612db2565b508161502a565b61502a83838151156151f15781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612db29190615290565b5f5b8381101561523f578181015183820152602001615227565b50505f910152565b5f815180845261525e816020860160208601615225565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6152a26020830184615247565b9392505050565b5f602082840312156152b9575f80fd5b813561ffff811681146152a2575f80fd5b803567ffffffffffffffff81168114612c0e575f80fd5b5f805f805f60a086880312156152f5575f80fd5b6152fe866152ca565b945061530c602087016152ca565b94979496505050506040830135926060810135926080909101359150565b80610300810183101561533b575f80fd5b92915050565b5f805f805f806103a08789031215615357575f80fd5b615360876152ca565b955061536e602088016152ca565b945061537c604088016152ca565b935060608701359250608087013591506153998860a0890161532a565b90509295509295509295565b5f805f805f805f6103c0888a0312156153bc575f80fd5b6153c5886152ca565b96506153d3602089016152ca565b95506153e1604089016152ca565b94506153ef606089016152ca565b93506080880135925060a0880135915061540c8960c08a0161532a565b905092959891949750929550565b5f6020828403121561542a575f80fd5b6152a2826152ca565b5f8083601f840112615443575f80fd5b50813567ffffffffffffffff81111561545a575f80fd5b6020830191508360208260051b8501011115615474575f80fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612c0e575f80fd5b5f805f604084860312156154b0575f80fd5b833567ffffffffffffffff8111156154c6575f80fd5b6154d286828701615433565b90945092506154e590506020850161547b565b90509250925092565b5f602082840312156154fe575f80fd5b6152a28261547b565b5f60208284031215615517575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f83011261555a575f80fd5b813567ffffffffffffffff808211156155755761557561551e565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156155bb576155bb61551e565b816040528381528660208588010111156155d3575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f60208284031215615602575f80fd5b813567ffffffffffffffff811115615618575f80fd5b61502a8482850161554b565b5f8083601f840112615634575f80fd5b50813567ffffffffffffffff81111561564b575f80fd5b602083019150836020828501011115615474575f80fd5b5f805f805f80868803610120811215615679575f80fd5b60a0811215615686575f80fd5b5086955060a0870135945060c087013567ffffffffffffffff808211156156ab575f80fd5b6156b78a838b0161554b565b955060e08901359150808211156156cc575f80fd5b6156d88a838b0161554b565b94506101008901359150808211156156ee575f80fd5b506156fb89828a01615624565b979a9699509497509295939492505050565b5f806020838503121561571e575f80fd5b823567ffffffffffffffff811115615734575f80fd5b61574085828601615433565b90969095509350505050565b5f805f6040848603121561575e575f80fd5b833567ffffffffffffffff811115615774575f80fd5b61578086828701615624565b909790965060209590950135949350505050565b600181811c908216806157a857607f821691505b6020821081036157df577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff818116838216019080821115615833576158336157e5565b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112615899575f80fd5b9190910192915050565b5f608082360312156158b3575f80fd5b6040516080810167ffffffffffffffff82821081831117156158d7576158d761551e565b8160405284359150808211156158eb575f80fd5b506158f83682860161554b565b82525060208301356020820152615911604084016152ca565b6040820152615922606084016152ca565b606082015292915050565b5f67ffffffffffffffff808316818103615949576159496157e5565b6001019392505050565b5f60208284031215615963575f80fd5b5051919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361599a5761599a6157e5565b5060010190565b67ffffffffffffffff828116828216039080821115615833576158336157e5565b8181038181111561533b5761533b6157e5565b808202811582820484141761533b5761533b6157e5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615a2757615a276159ec565b500490565b601f821115614e7a575f81815260208120601f850160051c81016020861015615a525750805b601f850160051c820191505b81811015615a7157828155600101615a5e565b505050505050565b815167ffffffffffffffff811115615a9357615a9361551e565b615aa781615aa18454615794565b84615a2c565b602080601f831160018114615af9575f8415615ac35750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555615a71565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015615b4557888601518255948401946001909101908401615b26565b5085821015615b8157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f67ffffffffffffffff808716835280861660208401525060606040830152615c05606083018486615b91565b9695505050505050565b8082018082111561533b5761533b6157e5565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615899575f80fd5b5f60608236031215615c64575f80fd5b6040516060810167ffffffffffffffff8282108183111715615c8857615c8861551e565b816040528435915080821115615c9c575f80fd5b50615ca93682860161554b565b82525060208301356020820152615cc2604084016152ca565b604082015292915050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201525f615c05606083018486615b91565b5f8251615899818460208701615225565b5f82615d3057615d306159ec565b500690565b6103208101610300808584378201835f5b6001811015615d65578151835260209283019290910190600101615d46565b5050509392505050565b5f60208284031215615d7f575f80fd5b815180151581146152a2575f80fd5b5f67ffffffffffffffff80841680615da857615da86159ec565b9216919091049291505056fea264697066735822122047cda75802a72f404a35b0e79fd74633f78e7724362c61c458b0a0e3d80d6d9364736f6c63430008140033",
}

// PolygonzkevmtestnetclearstorageABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmtestnetclearstorageMetaData.ABI instead.
var PolygonzkevmtestnetclearstorageABI = PolygonzkevmtestnetclearstorageMetaData.ABI

// PolygonzkevmtestnetclearstorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmtestnetclearstorageMetaData.Bin instead.
var PolygonzkevmtestnetclearstorageBin = PolygonzkevmtestnetclearstorageMetaData.Bin

// DeployPolygonzkevmtestnetclearstorage deploys a new Ethereum contract, binding an instance of Polygonzkevmtestnetclearstorage to it.
func DeployPolygonzkevmtestnetclearstorage(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Polygonzkevmtestnetclearstorage, error) {
	parsed, err := PolygonzkevmtestnetclearstorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmtestnetclearstorageBin), backend, _globalExitRootManager, _matic, _rollupVerifier, _bridgeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmtestnetclearstorage{PolygonzkevmtestnetclearstorageCaller: PolygonzkevmtestnetclearstorageCaller{contract: contract}, PolygonzkevmtestnetclearstorageTransactor: PolygonzkevmtestnetclearstorageTransactor{contract: contract}, PolygonzkevmtestnetclearstorageFilterer: PolygonzkevmtestnetclearstorageFilterer{contract: contract}}, nil
}

// Polygonzkevmtestnetclearstorage is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmtestnetclearstorage struct {
	PolygonzkevmtestnetclearstorageCaller     // Read-only binding to the contract
	PolygonzkevmtestnetclearstorageTransactor // Write-only binding to the contract
	PolygonzkevmtestnetclearstorageFilterer   // Log filterer for contract events
}

// PolygonzkevmtestnetclearstorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmtestnetclearstorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmtestnetclearstorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmtestnetclearstorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmtestnetclearstorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmtestnetclearstorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmtestnetclearstorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmtestnetclearstorageSession struct {
	Contract     *Polygonzkevmtestnetclearstorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolygonzkevmtestnetclearstorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmtestnetclearstorageCallerSession struct {
	Contract *PolygonzkevmtestnetclearstorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PolygonzkevmtestnetclearstorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmtestnetclearstorageTransactorSession struct {
	Contract     *PolygonzkevmtestnetclearstorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PolygonzkevmtestnetclearstorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmtestnetclearstorageRaw struct {
	Contract *Polygonzkevmtestnetclearstorage // Generic contract binding to access the raw methods on
}

// PolygonzkevmtestnetclearstorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmtestnetclearstorageCallerRaw struct {
	Contract *PolygonzkevmtestnetclearstorageCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmtestnetclearstorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmtestnetclearstorageTransactorRaw struct {
	Contract *PolygonzkevmtestnetclearstorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmtestnetclearstorage creates a new instance of Polygonzkevmtestnetclearstorage, bound to a specific deployed contract.
func NewPolygonzkevmtestnetclearstorage(address common.Address, backend bind.ContractBackend) (*Polygonzkevmtestnetclearstorage, error) {
	contract, err := bindPolygonzkevmtestnetclearstorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmtestnetclearstorage{PolygonzkevmtestnetclearstorageCaller: PolygonzkevmtestnetclearstorageCaller{contract: contract}, PolygonzkevmtestnetclearstorageTransactor: PolygonzkevmtestnetclearstorageTransactor{contract: contract}, PolygonzkevmtestnetclearstorageFilterer: PolygonzkevmtestnetclearstorageFilterer{contract: contract}}, nil
}

// NewPolygonzkevmtestnetclearstorageCaller creates a new read-only instance of Polygonzkevmtestnetclearstorage, bound to a specific deployed contract.
func NewPolygonzkevmtestnetclearstorageCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmtestnetclearstorageCaller, error) {
	contract, err := bindPolygonzkevmtestnetclearstorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageCaller{contract: contract}, nil
}

// NewPolygonzkevmtestnetclearstorageTransactor creates a new write-only instance of Polygonzkevmtestnetclearstorage, bound to a specific deployed contract.
func NewPolygonzkevmtestnetclearstorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmtestnetclearstorageTransactor, error) {
	contract, err := bindPolygonzkevmtestnetclearstorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageTransactor{contract: contract}, nil
}

// NewPolygonzkevmtestnetclearstorageFilterer creates a new log filterer instance of Polygonzkevmtestnetclearstorage, bound to a specific deployed contract.
func NewPolygonzkevmtestnetclearstorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmtestnetclearstorageFilterer, error) {
	contract, err := bindPolygonzkevmtestnetclearstorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageFilterer{contract: contract}, nil
}

// bindPolygonzkevmtestnetclearstorage binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmtestnetclearstorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmtestnetclearstorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmtestnetclearstorage.Contract.PolygonzkevmtestnetclearstorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PolygonzkevmtestnetclearstorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PolygonzkevmtestnetclearstorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmtestnetclearstorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) Admin() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Admin(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Admin(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BatchNumToStateRoot(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BatchNumToStateRoot(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BridgeAddress(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.BridgeAddress(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.CalculateRewardPerBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.CalculateRewardPerBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ChainID() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ChainID(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) ChainID() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ChainID(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.CheckStateRootInsidePrime(&_Polygonzkevmtestnetclearstorage.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.CheckStateRootInsidePrime(&_Polygonzkevmtestnetclearstorage.CallOpts, newStateRoot)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForceBatchTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForceBatchTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForcedBatches(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForcedBatches(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ForkID() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForkID(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) ForkID() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForkID(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetForcedBatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetForcedBatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetInputSnarkBytes(&_Polygonzkevmtestnetclearstorage.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetInputSnarkBytes(&_Polygonzkevmtestnetclearstorage.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetLastVerifiedBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GetLastVerifiedBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GlobalExitRootManager(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.GlobalExitRootManager(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsEmergencyState(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsEmergencyState(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsForcedBatchDisallowed(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsForcedBatchDisallowed(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsPendingStateConsolidable(&_Polygonzkevmtestnetclearstorage.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.IsPendingStateConsolidable(&_Polygonzkevmtestnetclearstorage.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastBatchSequenced(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastBatchSequenced(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastForceBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastForceBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastForceBatchSequenced(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastForceBatchSequenced(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastPendingState() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastPendingState(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastPendingState() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastPendingState(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastPendingStateConsolidated(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastPendingStateConsolidated(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastTimestamp(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastTimestamp(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastVerifiedBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.LastVerifiedBatch(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) Matic() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Matic(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) Matic() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Matic(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.MultiplierBatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.MultiplierBatchFee(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) NetworkName() (string, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.NetworkName(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) NetworkName() (string, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.NetworkName(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) Owner() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Owner(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) Owner() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Owner(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingAdmin(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingAdmin(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingStateTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingStateTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "pendingStateTransitions", arg0)

	outstruct := new(struct {
		Timestamp         uint64
		LastVerifiedBatch uint64
		ExitRoot          [32]byte
		StateRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExitRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.StateRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingStateTransitions(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.PendingStateTransitions(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.RollupVerifier(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.RollupVerifier(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "sequencedBatches", arg0)

	outstruct := new(struct {
		AccInputHash               [32]byte
		SequencedTimestamp         uint64
		PreviousLastBatchSequenced uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccInputHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.SequencedTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PreviousLastBatchSequenced = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequencedBatches(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequencedBatches(&_Polygonzkevmtestnetclearstorage.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedAggregator(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedAggregator(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedAggregatorTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedAggregatorTimeout(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedSequencer(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedSequencer(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedSequencerURL(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TrustedSequencerURL(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatchTimeTarget(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatchTimeTarget(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmtestnetclearstorage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) Version() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Version(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageCallerSession) Version() (*big.Int, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Version(&_Polygonzkevmtestnetclearstorage.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.AcceptAdminRole(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.AcceptAdminRole(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ActivateEmergencyState(&_Polygonzkevmtestnetclearstorage.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ActivateEmergencyState(&_Polygonzkevmtestnetclearstorage.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ActivateForceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ActivateForceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ClearStorage is a paid mutator transaction binding the contract method 0x69bd38a0.
//
// Solidity: function clearStorage() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ClearStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "clearStorage")
}

// ClearStorage is a paid mutator transaction binding the contract method 0x69bd38a0.
//
// Solidity: function clearStorage() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ClearStorage() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ClearStorage(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ClearStorage is a paid mutator transaction binding the contract method 0x69bd38a0.
//
// Solidity: function clearStorage() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ClearStorage() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ClearStorage(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ConsolidatePendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ConsolidatePendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.DeactivateEmergencyState(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.DeactivateEmergencyState(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForceBatch(&_Polygonzkevmtestnetclearstorage.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ForceBatch(&_Polygonzkevmtestnetclearstorage.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) Initialize(opts *bind.TransactOpts, initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Initialize(&_Polygonzkevmtestnetclearstorage.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.Initialize(&_Polygonzkevmtestnetclearstorage.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.OverridePendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.OverridePendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ProveNonDeterministicPendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.ProveNonDeterministicPendingState(&_Polygonzkevmtestnetclearstorage.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.RenounceOwnership(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.RenounceOwnership(&_Polygonzkevmtestnetclearstorage.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SequenceBatches(batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequenceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SequenceBatches(batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequenceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, batches, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequenceForceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SequenceForceBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetForceBatchTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetForceBatchTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetMultiplierBatchFee(&_Polygonzkevmtestnetclearstorage.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetMultiplierBatchFee(&_Polygonzkevmtestnetclearstorage.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetPendingStateTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetPendingStateTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedAggregator(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedAggregator(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedSequencer(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedSequencer(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedSequencerURL(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetTrustedSequencerURL(&_Polygonzkevmtestnetclearstorage.TransactOpts, newTrustedSequencerURL)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevmtestnetclearstorage.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevmtestnetclearstorage.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TransferAdminRole(&_Polygonzkevmtestnetclearstorage.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TransferAdminRole(&_Polygonzkevmtestnetclearstorage.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TransferOwnership(&_Polygonzkevmtestnetclearstorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.TransferOwnership(&_Polygonzkevmtestnetclearstorage.TransactOpts, newOwner)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatches(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageTransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmtestnetclearstorage.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevmtestnetclearstorage.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator struct {
	Event *PolygonzkevmtestnetclearstorageAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageAcceptAdminRole)
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
		it.Event = new(PolygonzkevmtestnetclearstorageAcceptAdminRole)
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
func (it *PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageAcceptAdminRoleIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageAcceptAdminRole)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmtestnetclearstorageAcceptAdminRole, error) {
	event := new(PolygonzkevmtestnetclearstorageAcceptAdminRole)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageActivateForceBatchesIterator struct {
	Event *PolygonzkevmtestnetclearstorageActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageActivateForceBatches)
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
		it.Event = new(PolygonzkevmtestnetclearstorageActivateForceBatches)
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
func (it *PolygonzkevmtestnetclearstorageActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageActivateForceBatches represents a ActivateForceBatches event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageActivateForceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageActivateForceBatchesIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageActivateForceBatches)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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

// ParseActivateForceBatches is a log parse operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseActivateForceBatches(log types.Log) (*PolygonzkevmtestnetclearstorageActivateForceBatches, error) {
	event := new(PolygonzkevmtestnetclearstorageActivateForceBatches)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator struct {
	Event *PolygonzkevmtestnetclearstorageConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageConsolidatePendingState)
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
		it.Event = new(PolygonzkevmtestnetclearstorageConsolidatePendingState)
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
func (it *PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageConsolidatePendingStateIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageConsolidatePendingState)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonzkevmtestnetclearstorageConsolidatePendingState, error) {
	event := new(PolygonzkevmtestnetclearstorageConsolidatePendingState)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator struct {
	Event *PolygonzkevmtestnetclearstorageEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageEmergencyStateActivated)
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
		it.Event = new(PolygonzkevmtestnetclearstorageEmergencyStateActivated)
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
func (it *PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageEmergencyStateActivatedIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageEmergencyStateActivated)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonzkevmtestnetclearstorageEmergencyStateActivated, error) {
	event := new(PolygonzkevmtestnetclearstorageEmergencyStateActivated)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator struct {
	Event *PolygonzkevmtestnetclearstorageEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageEmergencyStateDeactivated)
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
		it.Event = new(PolygonzkevmtestnetclearstorageEmergencyStateDeactivated)
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
func (it *PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageEmergencyStateDeactivatedIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageEmergencyStateDeactivated)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonzkevmtestnetclearstorageEmergencyStateDeactivated, error) {
	event := new(PolygonzkevmtestnetclearstorageEmergencyStateDeactivated)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageForceBatchIterator struct {
	Event *PolygonzkevmtestnetclearstorageForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageForceBatch)
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
		it.Event = new(PolygonzkevmtestnetclearstorageForceBatch)
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
func (it *PolygonzkevmtestnetclearstorageForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageForceBatch represents a ForceBatch event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmtestnetclearstorageForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageForceBatchIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageForceBatch)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmtestnetclearstorageForceBatch, error) {
	event := new(PolygonzkevmtestnetclearstorageForceBatch)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageInitializedIterator struct {
	Event *PolygonzkevmtestnetclearstorageInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageInitialized)
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
		it.Event = new(PolygonzkevmtestnetclearstorageInitialized)
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
func (it *PolygonzkevmtestnetclearstorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageInitialized represents a Initialized event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageInitializedIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageInitialized)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseInitialized(log types.Log) (*PolygonzkevmtestnetclearstorageInitialized, error) {
	event := new(PolygonzkevmtestnetclearstorageInitialized)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageOverridePendingStateIterator struct {
	Event *PolygonzkevmtestnetclearstorageOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageOverridePendingState)
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
		it.Event = new(PolygonzkevmtestnetclearstorageOverridePendingState)
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
func (it *PolygonzkevmtestnetclearstorageOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageOverridePendingState represents a OverridePendingState event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageOverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmtestnetclearstorageOverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageOverridePendingStateIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageOverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageOverridePendingState)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseOverridePendingState(log types.Log) (*PolygonzkevmtestnetclearstorageOverridePendingState, error) {
	event := new(PolygonzkevmtestnetclearstorageOverridePendingState)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageOwnershipTransferredIterator struct {
	Event *PolygonzkevmtestnetclearstorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageOwnershipTransferred)
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
		it.Event = new(PolygonzkevmtestnetclearstorageOwnershipTransferred)
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
func (it *PolygonzkevmtestnetclearstorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageOwnershipTransferred represents a OwnershipTransferred event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonzkevmtestnetclearstorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageOwnershipTransferredIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageOwnershipTransferred)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonzkevmtestnetclearstorageOwnershipTransferred, error) {
	event := new(PolygonzkevmtestnetclearstorageOwnershipTransferred)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator struct {
	Event *PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState)
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
		it.Event = new(PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState)
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
func (it *PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageProveNonDeterministicPendingStateIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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

// ParseProveNonDeterministicPendingState is a log parse operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState, error) {
	event := new(PolygonzkevmtestnetclearstorageProveNonDeterministicPendingState)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSequenceBatchesIterator struct {
	Event *PolygonzkevmtestnetclearstorageSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSequenceBatches)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSequenceBatches)
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
func (it *PolygonzkevmtestnetclearstorageSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSequenceBatches represents a SequenceBatches event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmtestnetclearstorageSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSequenceBatchesIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSequenceBatches)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmtestnetclearstorageSequenceBatches, error) {
	event := new(PolygonzkevmtestnetclearstorageSequenceBatches)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator struct {
	Event *PolygonzkevmtestnetclearstorageSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSequenceForceBatches)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSequenceForceBatches)
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
func (it *PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSequenceForceBatchesIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSequenceForceBatches)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmtestnetclearstorageSequenceForceBatches, error) {
	event := new(PolygonzkevmtestnetclearstorageSequenceForceBatches)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetForceBatchTimeout)
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
func (it *PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetForceBatchTimeoutIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetForceBatchTimeout)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmtestnetclearstorageSetForceBatchTimeout, error) {
	event := new(PolygonzkevmtestnetclearstorageSetForceBatchTimeout)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetMultiplierBatchFee)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetMultiplierBatchFee)
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
func (it *PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetMultiplierBatchFeeIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetMultiplierBatchFee)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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

// ParseSetMultiplierBatchFee is a log parse operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonzkevmtestnetclearstorageSetMultiplierBatchFee, error) {
	event := new(PolygonzkevmtestnetclearstorageSetMultiplierBatchFee)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetPendingStateTimeout)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetPendingStateTimeout)
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
func (it *PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetPendingStateTimeoutIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetPendingStateTimeout)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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

// ParseSetPendingStateTimeout is a log parse operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonzkevmtestnetclearstorageSetPendingStateTimeout, error) {
	event := new(PolygonzkevmtestnetclearstorageSetPendingStateTimeout)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedAggregator)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedAggregator)
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
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetTrustedAggregatorIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetTrustedAggregator)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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

// ParseSetTrustedAggregator is a log parse operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonzkevmtestnetclearstorageSetTrustedAggregator, error) {
	event := new(PolygonzkevmtestnetclearstorageSetTrustedAggregator)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout)
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
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeoutIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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

// ParseSetTrustedAggregatorTimeout is a log parse operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout, error) {
	event := new(PolygonzkevmtestnetclearstorageSetTrustedAggregatorTimeout)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedSequencer)
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
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetTrustedSequencerIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetTrustedSequencer)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmtestnetclearstorageSetTrustedSequencer, error) {
	event := new(PolygonzkevmtestnetclearstorageSetTrustedSequencer)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetTrustedSequencerURL)
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
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetTrustedSequencerURLIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetTrustedSequencerURL)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmtestnetclearstorageSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmtestnetclearstorageSetTrustedSequencerURL)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget)
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
func (it *PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTargetIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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

// ParseSetVerifyBatchTimeTarget is a log parse operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget, error) {
	event := new(PolygonzkevmtestnetclearstorageSetVerifyBatchTimeTarget)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageTransferAdminRoleIterator struct {
	Event *PolygonzkevmtestnetclearstorageTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageTransferAdminRole)
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
		it.Event = new(PolygonzkevmtestnetclearstorageTransferAdminRole)
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
func (it *PolygonzkevmtestnetclearstorageTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageTransferAdminRoleIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageTransferAdminRole)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmtestnetclearstorageTransferAdminRole, error) {
	event := new(PolygonzkevmtestnetclearstorageTransferAdminRole)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator is returned from FilterUpdateZkEVMVersion and is used to iterate over the raw logs and unpacked data for UpdateZkEVMVersion events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator struct {
	Event *PolygonzkevmtestnetclearstorageUpdateZkEVMVersion // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageUpdateZkEVMVersion)
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
		it.Event = new(PolygonzkevmtestnetclearstorageUpdateZkEVMVersion)
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
func (it *PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageUpdateZkEVMVersion represents a UpdateZkEVMVersion event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageUpdateZkEVMVersion struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateZkEVMVersion is a free log retrieval operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterUpdateZkEVMVersion(opts *bind.FilterOpts) (*PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageUpdateZkEVMVersionIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "UpdateZkEVMVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateZkEVMVersion is a free log subscription operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchUpdateZkEVMVersion(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageUpdateZkEVMVersion) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageUpdateZkEVMVersion)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
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

// ParseUpdateZkEVMVersion is a log parse operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseUpdateZkEVMVersion(log types.Log) (*PolygonzkevmtestnetclearstorageUpdateZkEVMVersion, error) {
	event := new(PolygonzkevmtestnetclearstorageUpdateZkEVMVersion)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageVerifyBatchesIterator struct {
	Event *PolygonzkevmtestnetclearstorageVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageVerifyBatches)
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
		it.Event = new(PolygonzkevmtestnetclearstorageVerifyBatches)
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
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageVerifyBatches represents a VerifyBatches event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmtestnetclearstorageVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageVerifyBatchesIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageVerifyBatches)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmtestnetclearstorageVerifyBatches, error) {
	event := new(PolygonzkevmtestnetclearstorageVerifyBatches)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator)
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
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonzkevmtestnetclearstorage contract.
type PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregatorIterator{contract: _Polygonzkevmtestnetclearstorage.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmtestnetclearstorage.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator)
				if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmtestnetclearstorage *PolygonzkevmtestnetclearstorageFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonzkevmtestnetclearstorageVerifyBatchesTrustedAggregator)
	if err := _Polygonzkevmtestnetclearstorage.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmmock

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

// PolygonzkevmmockMetaData contains all meta data concerning the Polygonzkevmmock contract.
var PolygonzkevmmockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_matic\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"_rollupVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_forkID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ActivateForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"UpdateZkEVMVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequencedBatchNum\",\"type\":\"uint64\"}],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"batchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currentAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"sequencerAddress\",\"type\":\"address\"}],\"name\":\"calculateAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"maticAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forkID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getNextSnarkInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"trustedAggregatorTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.InitializePackedParameters\",\"name\":\"initializePackedParameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"genesisRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isForcedBatchDisallowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingState\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPendingStateConsolidated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"matic\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingStateTransitions\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupVerifier\",\"outputs\":[{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"globalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"minForcedTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structPolygonZkEVM.ForcedBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"sequencedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"setNetworkName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_numBatch\",\"type\":\"uint64\"}],\"name\":\"setSequencedBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputData\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"}],\"name\":\"setSequencedBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"setStateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"setTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_numBatch\",\"type\":\"uint64\"}],\"name\":\"setVerifiedBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"proofB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"proofC\",\"type\":\"uint256[2]\"}],\"name\":\"trustedVerifyBatchesMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newLastVerifiedBatch\",\"type\":\"uint64\"}],\"name\":\"updateBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b5060405162007064380380620070648339810160408190526200003591620000a5565b6001600160a01b0395861660c05293851660805291841660a05290921660e0526001600160401b0391821661010052166101205262000131565b6001600160a01b03811681146200008557600080fd5b50565b80516001600160401b0381168114620000a057600080fd5b919050565b60008060008060008060c08789031215620000bf57600080fd5b8651620000cc816200006f565b6020880151909650620000df816200006f565b6040880151909550620000f2816200006f565b606088015190945062000105816200006f565b9250620001156080880162000088565b91506200012560a0880162000088565b90509295509295509295565b60805160a05160c05160e0516101005161012051616e5e620002066000396000818161083001528181611a880152613ec80152600081816109c30152611a5e01526000818161098901528181612a2d015281816145610152615a90015260008181610b550152818161173801528181611bfb01528181611dcc0152818161263501528181612e06015281816147d901526155df015260008181610c1501528181614f3a0152615392015260008181610a92015281816129fb0152818161333a015281816147ad01526150280152616e5e6000f3fe608060405234801561001057600080fd5b506004361061048d5760003560e01c80638c3d73011161026b578063c754c7ed11610150578063e7a7ed02116100c8578063f14916d611610097578063f851a4401161007c578063f851a44014610c9d578063f8b823e414610cbd578063fe16564f14610cc657600080fd5b8063f14916d614610c77578063f2fde38b14610c8a57600080fd5b8063e7a7ed0214610be0578063e8bf92ed14610c10578063eaeb077b14610c37578063ed6b010414610c4a57600080fd5b8063d2e129f91161011f578063d939b31511610104578063d939b31514610b9d578063dbc1697614610bc5578063e0d1744114610bcd57600080fd5b8063d2e129f914610b77578063d8d1091b14610b8a57600080fd5b8063c754c7ed14610ae2578063c89e42df14610b0e578063cfa8ed4714610b21578063d02103ca14610b5057600080fd5b8063ada8f919116101e3578063b4f77ea9116101b2578063ba58ae3911610197578063ba58ae3914610ab4578063c0cad30214610ac7578063c0ed84e014610ada57600080fd5b8063b4f77ea914610a7a578063b6b0b09714610a8d57600080fd5b8063ada8f919146109ab578063adc879e9146109be578063afd23cbe146109e5578063b4d63f5814610a1357600080fd5b80639aa972a31161023a5780639c9f3dfe1161021f5780639c9f3dfe1461095e578063a066215c14610971578063a3c573eb1461098457600080fd5b80639aa972a3146109385780639b7967601461094b57600080fd5b80638c3d7301146108f75780638da5cb5b146108ff57806396dc3d391461091d57806399f5634e1461093057600080fd5b80634a1a89a711610391578063621dd411116103095780637215541a116102d8578063831c7ead116102bd578063831c7ead1461082b578063837a473814610852578063841b24d7146108c757600080fd5b80637215541a146108045780637fcb36531461081757600080fd5b8063621dd411146107b65780636b8616ce146107c95780636ff512cc146107e9578063715018a6146107fc57600080fd5b8063542028d5116103605780635e9145c9116103455780635e9145c9146107935780635ec91958146107a657806360469169146107ae57600080fd5b8063542028d5146106ea578063574f649e146106f257600080fd5b80634a1a89a7146106845780634a910e6a146106a45780634e487706146106b75780635392c5e0146106ca57600080fd5b80632678224711610424578063383b3be8116103f3578063423fa856116103d8578063423fa856146106285780634560526714610648578063458c04771461067057600080fd5b8063383b3be814610602578063394218e91461061557600080fd5b8063267822471461056b57806329878983146105b05780632b0006fa146105dc5780632c1f816a146105ef57600080fd5b806315064c961161046057806315064c96146105145780631816b7e51461053157806319d8ac6114610544578063220d78991461055857600080fd5b80630a0d9fbe146104925780630eaa86ea146104c9578063107bf28c146104ea57806310a01a72146104ff575b600080fd5b606f546104ab90610100900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6104dc6104d7366004616066565b610cd9565b6040519081526020016104c0565b6104f26111b1565b6040516104c09190616129565b61051261050d36600461615a565b61123f565b005b606f546105219060ff1681565b60405190151581526020016104c0565b61051261053f3660046161ed565b6117f7565b6073546104ab9067ffffffffffffffff1681565b6104f2610566366004616211565b61190f565b607b5461058b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016104c0565b60745461058b9068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6105126105ea366004616270565b611ae6565b6105126105fd3660046162d8565b611cb6565b610521610610366004616352565b611ec4565b610512610623366004616352565b611f1a565b6073546104ab9068010000000000000000900467ffffffffffffffff1681565b6073546104ab90700100000000000000000000000000000000900467ffffffffffffffff1681565b6079546104ab9067ffffffffffffffff1681565b6079546104ab9068010000000000000000900467ffffffffffffffff1681565b6105126106b2366004616352565b61209e565b6105126106c5366004616352565b612151565b6104dc6106d8366004616352565b60756020526000908152604090205481565b6104f26122d5565b6104dc61070036600461646b565b835160209485012060408051808701979097528681019190915260608087019490945260c09290921b7fffffffffffffffff00000000000000000000000000000000000000000000000016608086015290911b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888401528051808403607c018152609c9093019052815191012090565b6105126107a1366004616529565b6122e2565b610512612aec565b6104dc612bec565b6105126107c4366004616270565b612c02565b6104dc6107d7366004616352565b60716020526000908152604090205481565b6105126107f736600461657d565b612f8a565b61051261305f565b610512610812366004616352565b613073565b6074546104ab9067ffffffffffffffff1681565b6104ab7f000000000000000000000000000000000000000000000000000000000000000081565b61089b610860366004616598565b60786020526000908152604090208054600182015460029092015467ffffffffffffffff808316936801000000000000000090930416919084565b6040805167ffffffffffffffff95861681529490931660208501529183015260608201526080016104c0565b6079546104ab907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6105126131e3565b60335473ffffffffffffffffffffffffffffffffffffffff1661058b565b61051261092b366004616352565b6132af565b6104dc6132f2565b6105126109463660046162d8565b61344b565b610512610959366004616352565b6134fc565b61051261096c366004616352565b61354b565b61051261097f366004616352565b6136c7565b61058b7f000000000000000000000000000000000000000000000000000000000000000081565b6105126109b936600461657d565b6137cd565b6104ab7f000000000000000000000000000000000000000000000000000000000000000081565b606f54610a00906901000000000000000000900461ffff1681565b60405161ffff90911681526020016104c0565b610a54610a21366004616352565b6072602052600090815260409020805460019091015467ffffffffffffffff808216916801000000000000000090041683565b6040805193845267ffffffffffffffff92831660208501529116908201526060016104c0565b610512610a88366004616352565b613891565b61058b7f000000000000000000000000000000000000000000000000000000000000000081565b610521610ac2366004616598565b6138a2565b610512610ad53660046165b1565b61392c565b6104ab613944565b607b546104ab9074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b610512610b1c3660046165b1565b613999565b606f5461058b906b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b61058b7f000000000000000000000000000000000000000000000000000000000000000081565b610512610b85366004616628565b613a26565b610512610b983660046166db565b613f6c565b6079546104ab90700100000000000000000000000000000000900467ffffffffffffffff1681565b61051261450e565b610512610bdb36600461671d565b6145e7565b6073546104ab907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61058b7f000000000000000000000000000000000000000000000000000000000000000081565b610512610c4536600461676a565b614677565b607b54610521907c0100000000000000000000000000000000000000000000000000000000900460ff1681565b610512610c8536600461657d565b614a6d565b610512610c9836600461657d565b614b3f565b607a5461058b9073ffffffffffffffffffffffffffffffffffffffff1681565b6104dc60705481565b610512610cd43660046167b6565b614bf3565b6000806000610ce6613944565b905067ffffffffffffffff881615610eb35760795467ffffffffffffffff9081169089161115610dc3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a2070656e60448201527f64696e6753746174654e756d206d757374206265206c657373206f722065717560648201527f616c207468616e206c61737450656e64696e6753746174650000000000000000608482015260a4015b60405180910390fd5b67ffffffffffffffff8089166000908152607860205260409020600281015481549094509091898116680100000000000000009092041614610ead576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d4261746368206d757374206d61746368207468652070656e64696e6760648201527f2073746174652062617463680000000000000000000000000000000000000000608482015260a401610dba565b5061104c565b67ffffffffffffffff8716600090815260756020526040902054915081610f82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d426174636820737461746520726f6f7420646f6573206e6f7420657860648201527f6973740000000000000000000000000000000000000000000000000000000000608482015260a401610dba565b8067ffffffffffffffff168767ffffffffffffffff16111561104c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605d60248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d4261746368206d757374206265206c657373206f7220657175616c2060648201527f7468616e2063757272656e744c61737456657269666965644261746368000000608482015260a401610dba565b8067ffffffffffffffff168667ffffffffffffffff1611611115576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605760248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a2066696e60448201527f616c4e65774261746368206d75737420626520626967676572207468616e206360648201527f757272656e744c61737456657269666965644261746368000000000000000000608482015260a401610dba565b6000611124888888868961190f565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161115991906167e2565b602060405180830381855afa158015611176573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061119991906167fe565b6111a39190616846565b9a9950505050505050505050565b607780546111be9061685a565b80601f01602080910402602001604051908101604052809291908181526020018280546111ea9061685a565b80156112375780601f1061120c57610100808354040283529160200191611237565b820191906000526020600020905b81548152906001019060200180831161121a57829003601f168201915b505050505081565b611247614c16565b600080611252613944565b905067ffffffffffffffff8a161561141a5760795467ffffffffffffffff908116908b16111561132a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a2070656e60448201527f64696e6753746174654e756d206d757374206265206c657373206f722065717560648201527f616c207468616e206c61737450656e64696e6753746174650000000000000000608482015260a401610dba565b67ffffffffffffffff808b1660009081526078602052604090206002810154815490945090918b8116680100000000000000009092041614611414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d4261746368206d757374206d61746368207468652070656e64696e6760648201527f2073746174652062617463680000000000000000000000000000000000000000608482015260a401610dba565b506115b3565b67ffffffffffffffff89166000908152607560205260409020549150816114e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d426174636820737461746520726f6f7420646f6573206e6f7420657860648201527f6973740000000000000000000000000000000000000000000000000000000000608482015260a401610dba565b8067ffffffffffffffff168967ffffffffffffffff1611156115b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605d60248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a20696e6960448201527f744e756d4261746368206d757374206265206c657373206f7220657175616c2060648201527f7468616e2063757272656e744c61737456657269666965644261746368000000608482015260a401610dba565b8067ffffffffffffffff168867ffffffffffffffff161161167c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605760248201527f506f6c79676f6e5a6b45564d3a3a766572696679426174636865733a2066696e60448201527f616c4e65774261746368206d75737420626520626967676572207468616e206360648201527f757272656e744c61737456657269666965644261746368000000000000000000608482015260a401610dba565b600061168b8a8a8a868b61190f565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff8c81169182179092556000908152607560205260409020899055607954919250161561170957607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018990527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b15801561179157600080fd5b505af11580156117a5573d6000803e3d6000fd5b505060405189815233925067ffffffffffffffff8c1691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe9060200160405180910390a35050505050505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611848576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88161ffff16108061186157506103ff8161ffff16115b15611898576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffff0000ffffffffffffffffff16690100000000000000000061ffff8416908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a150565b67ffffffffffffffff8086166000818152607260205260408082205493881682529020546060929115801590611943575081155b1561197a576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806119b1576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119ba846138a2565b6119f0576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b166020820152603481019690965260548601929092527fffffffffffffffff00000000000000000000000000000000000000000000000060c098891b811660748701527f0000000000000000000000000000000000000000000000000000000000000000891b8116607c8701527f0000000000000000000000000000000000000000000000000000000000000000891b81166084870152608c86019490945260ac85015260cc840194909452509290931b90911660ec830152805180830360d401815260f4909201905290565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611b43576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b51868686868686614c97565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86811691821790925560009081526075602052604090208390556079541615611bcc57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b158015611c5457600080fd5b505af1158015611c68573d6000803e3d6000fd5b505060405184815233925067ffffffffffffffff871691507fcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe906020015b60405180910390a3505050505050565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff163314611d13576040517fbbcbbc0500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d228787878787878761505b565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86811691821790925560009081526075602052604090208390556079541615611d9d57607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b158015611e2557600080fd5b505af1158015611e39573d6000803e3d6000fd5b50506079805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a800000000000000000000000000000000000000000000000001790555050604051828152339067ffffffffffffffff8616907fcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf729060200160405180910390a350505050505050565b60795467ffffffffffffffff82811660009081526078602052604081205490924292611f0892700100000000000000000000000000000000909204811691166168dc565b67ffffffffffffffff16111592915050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314611f6b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611fb2576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff166120215760795467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821610612021576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001611904565b60745468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16331461214557606f5460ff1615612106576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61210f81611ec4565b612145576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61214e81615495565b50565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146121a2576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156121e9576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1661225457607b5467ffffffffffffffff74010000000000000000000000000000000000000000909104811690821610612254576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b90602001611904565b607680546111be9061685a565b606f5460ff161561231f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f546b010000000000000000000000900473ffffffffffffffffffffffffffffffffffffffff16331461237f576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160008190036123bb576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156123f7576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff6801000000000000000082048116600081815260726020526040812054838516949293700100000000000000000000000000000000909304909216919082905b868110156128475760008a8a8381811061245f5761245f616904565b90506020028101906124719190616933565b61247a90616967565b8051805160209091012060608201519192509067ffffffffffffffff16156125f257856124a6816169f4565b965050600081836020015184606001516040516020016124fe93929190928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a16600090815260719093529120549091508114612587576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80881660009081526071602052604080822091909155606085015190850151908216911610156125ec576040517f7f7ab87200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061272f565b6020820151158015906126b9575060208201516040517f257b363200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063257b3632906024016020604051808303816000875af1158015612693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b791906167fe565b155b156126f0576040517f73bd668d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8151516201d4c0101561272f576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8767ffffffffffffffff16826040015167ffffffffffffffff161080612762575042826040015167ffffffffffffffff16115b15612799576040517fea82791600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602082810151604080850151815193840189905290830184905260608084019290925260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c01604051602081830303815290604052805190602001209450816040015197505050808061283f90616a1b565b915050612443565b5061285286856168dc565b60735490945067ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690841611156128bb576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006128c78285616a53565b6128db9067ffffffffffffffff1688616a74565b604080516060810182528581524267ffffffffffffffff908116602080840191825260738054680100000000000000009081900485168688019081528d861660008181526072909552979093209551865592516001909501805492519585167fffffffffffffffffffffffffffffffff000000000000000000000000000000009384161795851684029590951790945583548c84169116179302929092179055909150828116908516146129d157607380547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8716021790555b612a233330836070546129e49190616a87565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169291906156a8565b612a2b61578a565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612a9357600080fd5b505af1158015612aa7573d6000803e3d6000fd5b505060405167ffffffffffffffff881692507f303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce9150600090a250505050505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612b3d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16612b99576040517ff6ba91a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690556040517f854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f90600090a1565b60006070546064612bfd9190616a87565b905090565b606f5460ff1615612c3f576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff8581166000908152607260205260409020600101544292612c8c927801000000000000000000000000000000000000000000000000909104811691166168dc565b67ffffffffffffffff161115612cce576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8612cdb8686616a53565b67ffffffffffffffff161115612d1d576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612d2b868686868686614c97565b612d3484615837565b607954700100000000000000000000000000000000900467ffffffffffffffff16600003612e7c57607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff86811691821790925560009081526075602052604090208390556079541615612dd757607980547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690555b6040517f33d6247d000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b158015612e5f57600080fd5b505af1158015612e73573d6000803e3d6000fd5b50505050612f4c565b612e8461578a565b6079805467ffffffffffffffff16906000612e9e836169f4565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815287831660208083019182528284018981526060840189815260795487166000908152607890935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b604051828152339067ffffffffffffffff8616907f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596690602001611ca6565b607a5473ffffffffffffffffffffffffffffffffffffffff163314612fdb576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fff0000000000000000000000000000000000000000ffffffffffffffffffffff166b01000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001611904565b613067614c16565b6130716000615a17565b565b60335473ffffffffffffffffffffffffffffffffffffffff1633146131db57600061309c613944565b90508067ffffffffffffffff168267ffffffffffffffff16116130eb576040517f812a372d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff6801000000000000000090910481169083161180613131575067ffffffffffffffff80831660009081526072602052604090206001015416155b15613168576040517f98c5c01400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff80831660009081526072602052604090206001015442916131979162093a8091166168dc565b67ffffffffffffffff1611156131d9576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b61214e615a8e565b607b5473ffffffffffffffffffffffffffffffffffffffff163314613234576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b54607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b6132b7614c16565b607480547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff92909216919091179055565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015613381573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133a591906167fe565b905060006133b1613944565b60735467ffffffffffffffff6801000000000000000082048116916134099170010000000000000000000000000000000082048116917801000000000000000000000000000000000000000000000000900416616a53565b61341391906168dc565b61341d9190616a53565b67ffffffffffffffff1690508060000361343a5760009250505090565b6134448183616a9e565b9250505090565b606f5460ff1615613488576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6134978787878787878761505b565b67ffffffffffffffff84166000908152607560209081526040918290205482519081529081018490527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a16134f3615a8e565b50505050505050565b613504614c16565b6073805467ffffffffffffffff90921668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff909216919091179055565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461359c576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff821611156135e3576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1661364a5760795467ffffffffffffffff70010000000000000000000000000000000090910481169082161061364a576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607980547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001611904565b607a5473ffffffffffffffffffffffffffffffffffffffff163314613718576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808167ffffffffffffffff16111561375f576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff1661010067ffffffffffffffff8416908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001611904565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461381e576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001611904565b613899614c16565b61214e81615837565b600067ffffffff0000000167ffffffffffffffff83161080156138da575067ffffffff00000001604083901c67ffffffffffffffff16105b80156138fb575067ffffffff00000001608083901c67ffffffffffffffff16105b8015613912575067ffffffff0000000160c083901c105b1561391f57506001919050565b506000919050565b919050565b613934614c16565b60776139408282616b00565b5050565b60795460009067ffffffffffffffff1615613988575060795467ffffffffffffffff9081166000908152607860205260409020546801000000000000000090041690565b5060745467ffffffffffffffff1690565b607a5473ffffffffffffffffffffffffffffffffffffffff1633146139ea576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60766139f68282616b00565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516119049190616129565b600054610100900460ff1615808015613a465750600054600160ff909116105b80613a605750303b158015613a60575060005460ff166001145b613aec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610dba565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015613b4a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b613b57602088018861657d565b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055613bac604088016020890161657d565b606f805473ffffffffffffffffffffffffffffffffffffffff929092166b010000000000000000000000027fff0000000000000000000000000000000000000000ffffffffffffffffffffff909216919091179055613c11608088016060890161657d565b6074805473ffffffffffffffffffffffffffffffffffffffff9290921668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790556000805260756020527ff9e3fbf150b7a0077118526f473c53cb4734f166167e2c6213e3567dd390b4ad8690556076613c9c8682616b00565b506077613ca98582616b00565b5062093a80613cbe6060890160408a01616352565b67ffffffffffffffff161115613d00576040517fcc96507000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613d106060880160408901616352565b6079805467ffffffffffffffff92909216700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905562093a80613d7260a0890160808a01616352565b67ffffffffffffffff161115613db4576040517f1d06e87900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613dc460a0880160808901616352565b6079805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff939093169290920291909117905567016345785d8a0000607055606f80547fffffffffffffffffffffffffffffffffffffffffff00000000000000000000ff166a03ea000000000000070800179055607b80547fffffff000000000000000000ffffffffffffffffffffffffffffffffffffffff167c0100000000000697800000000000000000000000000000000000000000179055613ea3615b16565b7fed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd660007f00000000000000000000000000000000000000000000000000000000000000008585604051613ef99493929190616c63565b60405180910390a180156134f357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff1615613fc9576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1615614006576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000819003614042576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561407e576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff780100000000000000000000000000000000000000000000000082048116916140c9918491700100000000000000000000000000000000900416616c9b565b1115614101576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60735467ffffffffffffffff680100000000000000008204811660008181526072602052604081205491937001000000000000000000000000000000009004909216915b848110156143ab57600087878381811061416157614161616904565b90506020028101906141739190616cae565b61417c90616ce2565b905083614188816169f4565b825180516020918201208185015160408087015190519499509194506000936141ea9386939101928352602083019190915260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016604082015260480190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8916600090815260719093529120549091508114614273576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8616600090815260716020526040812055614298600189616a74565b84036143075742607b60149054906101000a900467ffffffffffffffff1684604001516142c591906168dc565b67ffffffffffffffff161115614307576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020838101516040805192830188905282018490526060808301919091524260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016608083015233901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088820152609c0160405160208183030381529060405280519060200120945050505080806143a390616a1b565b915050614145565b506143b684846168dc565b6073805467ffffffffffffffff4281167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217808455604080516060810182528781526020808201958652680100000000000000009384900485168284019081528589166000818152607290935284832093518455965160019390930180549151871686027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921693871693909317179091558554938916700100000000000000000000000000000000027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff938602939093167fffffffffffffffff00000000000000000000000000000000ffffffffffffffff90941693909317919091179093559151929550917f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a49190a2505050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff16331461455f576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dbc169766040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156145c757600080fd5b505af11580156145db573d6000803e3d6000fd5b50505050613071615bb6565b6145ef614c16565b6040805160608101825293845267ffffffffffffffff92831660208086019182529284168583019081529584166000908152607290935291209251835551600190920180549351821668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009094169290911691909117919091179055565b607b547c0100000000000000000000000000000000000000000000000000000000900460ff16156146d4576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f5460ff1615614711576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061471b612bec565b905081811115614757576040517f4732fdb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388831115614793576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6147d573ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846156a8565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614842573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061486691906167fe565b60738054919250780100000000000000000000000000000000000000000000000090910467ffffffffffffffff169060186148a0836169f4565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505084846040516148d7929190616d5e565b60408051918290038220602083015281018290527fffffffffffffffff0000000000000000000000000000000000000000000000004260c01b166060820152606801604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815291815281516020928301206073547801000000000000000000000000000000000000000000000000900467ffffffffffffffff1660009081526071909352912055323303614a0757607354604080518381523360208201526060918101829052600091810191909152780100000000000000000000000000000000000000000000000090910467ffffffffffffffff16907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319060800160405180910390a2614a66565b607360189054906101000a900467ffffffffffffffff1667ffffffffffffffff167ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93182338888604051614a5d9493929190616d6e565b60405180910390a25b5050505050565b607a5473ffffffffffffffffffffffffffffffffffffffff163314614abe576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607480547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca90602001611904565b614b47614c16565b73ffffffffffffffffffffffffffffffffffffffff8116614bea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610dba565b61214e81615a17565b614bfb614c16565b67ffffffffffffffff16600090815260756020526040902055565b60335473ffffffffffffffffffffffffffffffffffffffff163314613071576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610dba565b600080614ca2613944565b905067ffffffffffffffff881615614d725760795467ffffffffffffffff9081169089161115614cfe576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8089166000908152607860205260409020600281015481549094509091898116680100000000000000009092041614614d6c576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50614e13565b67ffffffffffffffff8716600090815260756020526040902054915081614dc5576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff161115614e13576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168667ffffffffffffffff1611614e60576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000614e6f888888868961190f565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051614ea491906167e2565b602060405180830381855afa158015614ec1573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190614ee491906167fe565b614eee9190616846565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a91614f7091899190600401616da4565b602060405180830381865afa158015614f8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614fb19190616ddf565b614fe7576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61504f33614ff5858b616a53565b67ffffffffffffffff166150076132f2565b6150119190616a87565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190615c45565b50505050505050505050565b600067ffffffffffffffff8816156151295760795467ffffffffffffffff90811690891611156150b7576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088166000908152607860205260409020600281015481549092888116680100000000000000009092041614615123576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506151c5565b5067ffffffffffffffff85166000908152607560205260409020548061517b576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745467ffffffffffffffff90811690871611156151c5576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff90811690881611806151f757508767ffffffffffffffff168767ffffffffffffffff1611155b8061521e575060795467ffffffffffffffff68010000000000000000909104811690881611155b15615255576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781166000908152607860205260409020546801000000000000000090048116908616146152b8576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006152c7878787858861190f565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516152fc91906167e2565b602060405180830381855afa158015615319573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061533c91906167fe565b6153469190616846565b6040805160208101825282815290517f9121da8a00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691639121da8a916153c891889190600401616da4565b602060405180830381865afa1580156153e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906154099190616ddf565b61543f576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff891660009081526078602052604090206002015485900361504f576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795467ffffffffffffffff6801000000000000000090910481169082161115806154cf575060795467ffffffffffffffff908116908216115b15615506576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff818116600081815260786020908152604080832080546074805468010000000000000000928390049098167fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090981688179055600282015487865260759094529382902092909255607980547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff169390940292909217909255600182015490517f33d6247d00000000000000000000000000000000000000000000000000000000815260048101919091529091907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906333d6247d90602401600060405180830381600087803b15801561563857600080fd5b505af115801561564c573d6000803e3d6000fd5b505050508267ffffffffffffffff168167ffffffffffffffff167f328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e846002015460405161569b91815260200190565b60405180910390a3505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526157849085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152615ca0565b50505050565b60795467ffffffffffffffff680100000000000000008204811691161115613071576079546000906157d39068010000000000000000900467ffffffffffffffff1660016168dc565b90506157de81611ec4565b1561214e5760795460009060029061580190849067ffffffffffffffff16616a53565b61580b9190616e01565b61581590836168dc565b905061582081611ec4565b1561582e5761394081615495565b61394082615495565b6000615841613944565b9050816000806158518484616a53565b606f5467ffffffffffffffff91821692506000916158759161010090041642616a74565b90505b8467ffffffffffffffff168467ffffffffffffffff16146159005767ffffffffffffffff808516600090815260726020526040902060018101549091168210156158de57600181015468010000000000000000900467ffffffffffffffff1694506158fa565b6158e88686616a53565b67ffffffffffffffff16935050615900565b50615878565b600061590c8484616a74565b90508381101561596357808403600c8111615927578061592a565b600c5b9050806103e80a81606f60099054906101000a900461ffff1661ffff160a607054028161595957615959616817565b04607055506159d3565b838103600c81116159745780615977565b600c5b90506000816103e80a82606f60099054906101000a900461ffff1661ffff160a670de0b6b3a764000002816159ae576159ae616817565b04905080607054670de0b6b3a764000002816159cc576159cc616817565b0460705550505b683635c9adc5dea0000060705411156159f857683635c9adc5dea000006070556134f3565b633b9aca0060705410156134f357633b9aca0060705550505050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615af657600080fd5b505af1158015615b0a573d6000803e3d6000fd5b50505050613071615dac565b600054610100900460ff16615bad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610dba565b61307133615a17565b606f5460ff16615bf2576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052615c9b9084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401615702565b505050565b6000615d02826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16615e3f9092919063ffffffff16565b805190915015615c9b5780806020019051810190615d209190616ddf565b615c9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610dba565b606f5460ff1615615de9576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b6060615e4e8484600085615e56565b949350505050565b606082471015615ee8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610dba565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051615f1191906167e2565b60006040518083038185875af1925050503d8060008114615f4e576040519150601f19603f3d011682016040523d82523d6000602084013e615f53565b606091505b5091509150615f6487838387615f6f565b979650505050505050565b60608315616005578251600003615ffe5773ffffffffffffffffffffffffffffffffffffffff85163b615ffe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610dba565b5081615e4e565b615e4e838381511561601a5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dba9190616129565b803567ffffffffffffffff8116811461392757600080fd5b600080600080600060a0868803121561607e57600080fd5b6160878661604e565b94506160956020870161604e565b93506160a36040870161604e565b94979396509394606081013594506080013592915050565b60005b838110156160d65781810151838201526020016160be565b50506000910152565b600081518084526160f78160208601602086016160bb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061613c60208301846160df565b9392505050565b806040810183101561615457600080fd5b92915050565b6000806000806000806000806101a0898b03121561617757600080fd5b6161808961604e565b975061618e60208a0161604e565b965061619c60408a0161604e565b955060608901359450608089013593506161b98a60a08b01616143565b925061016089018a8111156161cd57600080fd5b60e08a0192506161dd8b82616143565b9150509295985092959890939650565b6000602082840312156161ff57600080fd5b813561ffff8116811461613c57600080fd5b600080600080600060a0868803121561622957600080fd5b6162328661604e565b94506162406020870161604e565b94979496505050506040830135926060810135926080909101359150565b80610300810183101561615457600080fd5b6000806000806000806103a0878903121561628a57600080fd5b6162938761604e565b95506162a16020880161604e565b94506162af6040880161604e565b935060608701359250608087013591506162cc8860a0890161625e565b90509295509295509295565b60008060008060008060006103c0888a0312156162f457600080fd5b6162fd8861604e565b965061630b6020890161604e565b95506163196040890161604e565b94506163276060890161604e565b93506080880135925060a088013591506163448960c08a0161625e565b905092959891949750929550565b60006020828403121561636457600080fd5b61613c8261604e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126163ad57600080fd5b813567ffffffffffffffff808211156163c8576163c861636d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561640e5761640e61636d565b8160405283815286602085880101111561642757600080fd5b836020870160208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461392757600080fd5b600080600080600060a0868803121561648357600080fd5b85359450602086013567ffffffffffffffff8111156164a157600080fd5b6164ad8882890161639c565b945050604086013592506164c36060870161604e565b91506164d160808701616447565b90509295509295909350565b60008083601f8401126164ef57600080fd5b50813567ffffffffffffffff81111561650757600080fd5b6020830191508360208260051b850101111561652257600080fd5b9250929050565b60008060006040848603121561653e57600080fd5b833567ffffffffffffffff81111561655557600080fd5b616561868287016164dd565b9094509250616574905060208501616447565b90509250925092565b60006020828403121561658f57600080fd5b61613c82616447565b6000602082840312156165aa57600080fd5b5035919050565b6000602082840312156165c357600080fd5b813567ffffffffffffffff8111156165da57600080fd5b615e4e8482850161639c565b60008083601f8401126165f857600080fd5b50813567ffffffffffffffff81111561661057600080fd5b60208301915083602082850101111561652257600080fd5b60008060008060008086880361012081121561664357600080fd5b60a081121561665157600080fd5b5086955060a0870135945060c087013567ffffffffffffffff8082111561667757600080fd5b6166838a838b0161639c565b955060e089013591508082111561669957600080fd5b6166a58a838b0161639c565b94506101008901359150808211156166bc57600080fd5b506166c989828a016165e6565b979a9699509497509295939492505050565b600080602083850312156166ee57600080fd5b823567ffffffffffffffff81111561670557600080fd5b616711858286016164dd565b90969095509350505050565b6000806000806080858703121561673357600080fd5b61673c8561604e565b9350602085013592506167516040860161604e565b915061675f6060860161604e565b905092959194509250565b60008060006040848603121561677f57600080fd5b833567ffffffffffffffff81111561679657600080fd5b6167a2868287016165e6565b909790965060209590950135949350505050565b600080604083850312156167c957600080fd5b823591506167d96020840161604e565b90509250929050565b600082516167f48184602087016160bb565b9190910192915050565b60006020828403121561681057600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261685557616855616817565b500690565b600181811c9082168061686e57607f821691505b6020821081036168a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8181168382160190808211156168fd576168fd6168ad565b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818336030181126167f457600080fd5b60006080823603121561697957600080fd5b6040516080810167ffffffffffffffff828210818311171561699d5761699d61636d565b8160405284359150808211156169b257600080fd5b506169bf3682860161639c565b825250602083013560208201526169d86040840161604e565b60408201526169e96060840161604e565b606082015292915050565b600067ffffffffffffffff808316818103616a1157616a116168ad565b6001019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203616a4c57616a4c6168ad565b5060010190565b67ffffffffffffffff8281168282160390808211156168fd576168fd6168ad565b81810381811115616154576161546168ad565b8082028115828204841417616154576161546168ad565b600082616aad57616aad616817565b500490565b601f821115615c9b57600081815260208120601f850160051c81016020861015616ad95750805b601f850160051c820191505b81811015616af857828155600101616ae5565b505050505050565b815167ffffffffffffffff811115616b1a57616b1a61636d565b616b2e81616b28845461685a565b84616ab2565b602080601f831160018114616b815760008415616b4b5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555616af8565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015616bce57888601518255948401946001909101908401616baf565b5085821015616c0a57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600067ffffffffffffffff808716835280861660208401525060606040830152616c91606083018486616c1a565b9695505050505050565b80820180821115616154576161546168ad565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa18336030181126167f457600080fd5b600060608236031215616cf457600080fd5b6040516060810167ffffffffffffffff8282108183111715616d1857616d1861636d565b816040528435915080821115616d2d57600080fd5b50616d3a3682860161639c565b82525060208301356020820152616d536040840161604e565b604082015292915050565b8183823760009101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000616c91606083018486616c1a565b61032081016103008085843782018360005b6001811015616dd5578151835260209283019290910190600101616db6565b5050509392505050565b600060208284031215616df157600080fd5b8151801515811461613c57600080fd5b600067ffffffffffffffff80841680616e1c57616e1c616817565b9216919091049291505056fea2646970667358221220686527cac84d98a7b3630b988cb7efce8cacdf200640ffe3e0ae6e4b6af7579264736f6c63430008140033",
}

// PolygonzkevmmockABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmmockMetaData.ABI instead.
var PolygonzkevmmockABI = PolygonzkevmmockMetaData.ABI

// PolygonzkevmmockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmmockMetaData.Bin instead.
var PolygonzkevmmockBin = PolygonzkevmmockMetaData.Bin

// DeployPolygonzkevmmock deploys a new Ethereum contract, binding an instance of Polygonzkevmmock to it.
func DeployPolygonzkevmmock(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _matic common.Address, _rollupVerifier common.Address, _bridgeAddress common.Address, _chainID uint64, _forkID uint64) (common.Address, *types.Transaction, *Polygonzkevmmock, error) {
	parsed, err := PolygonzkevmmockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmmockBin), backend, _globalExitRootManager, _matic, _rollupVerifier, _bridgeAddress, _chainID, _forkID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmmock{PolygonzkevmmockCaller: PolygonzkevmmockCaller{contract: contract}, PolygonzkevmmockTransactor: PolygonzkevmmockTransactor{contract: contract}, PolygonzkevmmockFilterer: PolygonzkevmmockFilterer{contract: contract}}, nil
}

// Polygonzkevmmock is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmmock struct {
	PolygonzkevmmockCaller     // Read-only binding to the contract
	PolygonzkevmmockTransactor // Write-only binding to the contract
	PolygonzkevmmockFilterer   // Log filterer for contract events
}

// PolygonzkevmmockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmmockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmmockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmmockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmmockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmmockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmmockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmmockSession struct {
	Contract     *Polygonzkevmmock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonzkevmmockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmmockCallerSession struct {
	Contract *PolygonzkevmmockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PolygonzkevmmockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmmockTransactorSession struct {
	Contract     *PolygonzkevmmockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PolygonzkevmmockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmmockRaw struct {
	Contract *Polygonzkevmmock // Generic contract binding to access the raw methods on
}

// PolygonzkevmmockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmmockCallerRaw struct {
	Contract *PolygonzkevmmockCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmmockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmmockTransactorRaw struct {
	Contract *PolygonzkevmmockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmmock creates a new instance of Polygonzkevmmock, bound to a specific deployed contract.
func NewPolygonzkevmmock(address common.Address, backend bind.ContractBackend) (*Polygonzkevmmock, error) {
	contract, err := bindPolygonzkevmmock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmmock{PolygonzkevmmockCaller: PolygonzkevmmockCaller{contract: contract}, PolygonzkevmmockTransactor: PolygonzkevmmockTransactor{contract: contract}, PolygonzkevmmockFilterer: PolygonzkevmmockFilterer{contract: contract}}, nil
}

// NewPolygonzkevmmockCaller creates a new read-only instance of Polygonzkevmmock, bound to a specific deployed contract.
func NewPolygonzkevmmockCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmmockCaller, error) {
	contract, err := bindPolygonzkevmmock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockCaller{contract: contract}, nil
}

// NewPolygonzkevmmockTransactor creates a new write-only instance of Polygonzkevmmock, bound to a specific deployed contract.
func NewPolygonzkevmmockTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmmockTransactor, error) {
	contract, err := bindPolygonzkevmmock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockTransactor{contract: contract}, nil
}

// NewPolygonzkevmmockFilterer creates a new log filterer instance of Polygonzkevmmock, bound to a specific deployed contract.
func NewPolygonzkevmmockFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmmockFilterer, error) {
	contract, err := bindPolygonzkevmmock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockFilterer{contract: contract}, nil
}

// bindPolygonzkevmmock binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmmock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmmockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmmock *PolygonzkevmmockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmmock.Contract.PolygonzkevmmockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmmock *PolygonzkevmmockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.PolygonzkevmmockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmmock *PolygonzkevmmockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.PolygonzkevmmockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmmock *PolygonzkevmmockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmmock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmmock *PolygonzkevmmockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmmock *PolygonzkevmmockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) Admin() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Admin(&_Polygonzkevmmock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) Admin() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Admin(&_Polygonzkevmmock.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) BatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "batchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.BatchFee(&_Polygonzkevmmock.CallOpts)
}

// BatchFee is a free data retrieval call binding the contract method 0xf8b823e4.
//
// Solidity: function batchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) BatchFee() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.BatchFee(&_Polygonzkevmmock.CallOpts)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) BatchNumToStateRoot(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "batchNumToStateRoot", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.BatchNumToStateRoot(&_Polygonzkevmmock.CallOpts, arg0)
}

// BatchNumToStateRoot is a free data retrieval call binding the contract method 0x5392c5e0.
//
// Solidity: function batchNumToStateRoot(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) BatchNumToStateRoot(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.BatchNumToStateRoot(&_Polygonzkevmmock.CallOpts, arg0)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmmock.Contract.BridgeAddress(&_Polygonzkevmmock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonzkevmmock.Contract.BridgeAddress(&_Polygonzkevmmock.CallOpts)
}

// CalculateAccInputHash is a free data retrieval call binding the contract method 0x574f649e.
//
// Solidity: function calculateAccInputHash(bytes32 currentAccInputHash, bytes transactions, bytes32 globalExitRoot, uint64 timestamp, address sequencerAddress) pure returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) CalculateAccInputHash(opts *bind.CallOpts, currentAccInputHash [32]byte, transactions []byte, globalExitRoot [32]byte, timestamp uint64, sequencerAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "calculateAccInputHash", currentAccInputHash, transactions, globalExitRoot, timestamp, sequencerAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateAccInputHash is a free data retrieval call binding the contract method 0x574f649e.
//
// Solidity: function calculateAccInputHash(bytes32 currentAccInputHash, bytes transactions, bytes32 globalExitRoot, uint64 timestamp, address sequencerAddress) pure returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockSession) CalculateAccInputHash(currentAccInputHash [32]byte, transactions []byte, globalExitRoot [32]byte, timestamp uint64, sequencerAddress common.Address) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.CalculateAccInputHash(&_Polygonzkevmmock.CallOpts, currentAccInputHash, transactions, globalExitRoot, timestamp, sequencerAddress)
}

// CalculateAccInputHash is a free data retrieval call binding the contract method 0x574f649e.
//
// Solidity: function calculateAccInputHash(bytes32 currentAccInputHash, bytes transactions, bytes32 globalExitRoot, uint64 timestamp, address sequencerAddress) pure returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) CalculateAccInputHash(currentAccInputHash [32]byte, transactions []byte, globalExitRoot [32]byte, timestamp uint64, sequencerAddress common.Address) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.CalculateAccInputHash(&_Polygonzkevmmock.CallOpts, currentAccInputHash, transactions, globalExitRoot, timestamp, sequencerAddress)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.CalculateRewardPerBatch(&_Polygonzkevmmock.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.CalculateRewardPerBatch(&_Polygonzkevmmock.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) ChainID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) ChainID() (uint64, error) {
	return _Polygonzkevmmock.Contract.ChainID(&_Polygonzkevmmock.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) ChainID() (uint64, error) {
	return _Polygonzkevmmock.Contract.ChainID(&_Polygonzkevmmock.CallOpts)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) CheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevmmock.Contract.CheckStateRootInsidePrime(&_Polygonzkevmmock.CallOpts, newStateRoot)
}

// CheckStateRootInsidePrime is a free data retrieval call binding the contract method 0xba58ae39.
//
// Solidity: function checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) CheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonzkevmmock.Contract.CheckStateRootInsidePrime(&_Polygonzkevmmock.CallOpts, newStateRoot)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.ForceBatchTimeout(&_Polygonzkevmmock.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.ForceBatchTimeout(&_Polygonzkevmmock.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.ForcedBatches(&_Polygonzkevmmock.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonzkevmmock.Contract.ForcedBatches(&_Polygonzkevmmock.CallOpts, arg0)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) ForkID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "forkID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) ForkID() (uint64, error) {
	return _Polygonzkevmmock.Contract.ForkID(&_Polygonzkevmmock.CallOpts)
}

// ForkID is a free data retrieval call binding the contract method 0x831c7ead.
//
// Solidity: function forkID() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) ForkID() (uint64, error) {
	return _Polygonzkevmmock.Contract.ForkID(&_Polygonzkevmmock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.GetForcedBatchFee(&_Polygonzkevmmock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonzkevmmock.Contract.GetForcedBatchFee(&_Polygonzkevmmock.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) GetInputSnarkBytes(opts *bind.CallOpts, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "getInputSnarkBytes", initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmmock *PolygonzkevmmockSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevmmock.Contract.GetInputSnarkBytes(&_Polygonzkevmmock.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x220d7899.
//
// Solidity: function getInputSnarkBytes(uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) GetInputSnarkBytes(initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonzkevmmock.Contract.GetInputSnarkBytes(&_Polygonzkevmmock.CallOpts, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) GetLastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "getLastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.GetLastVerifiedBatch(&_Polygonzkevmmock.CallOpts)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0xc0ed84e0.
//
// Solidity: function getLastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) GetLastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.GetLastVerifiedBatch(&_Polygonzkevmmock.CallOpts)
}

// GetNextSnarkInput is a free data retrieval call binding the contract method 0x0eaa86ea.
//
// Solidity: function getNextSnarkInput(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot) view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) GetNextSnarkInput(opts *bind.CallOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "getNextSnarkInput", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextSnarkInput is a free data retrieval call binding the contract method 0x0eaa86ea.
//
// Solidity: function getNextSnarkInput(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot) view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockSession) GetNextSnarkInput(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte) (*big.Int, error) {
	return _Polygonzkevmmock.Contract.GetNextSnarkInput(&_Polygonzkevmmock.CallOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot)
}

// GetNextSnarkInput is a free data retrieval call binding the contract method 0x0eaa86ea.
//
// Solidity: function getNextSnarkInput(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot) view returns(uint256)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) GetNextSnarkInput(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte) (*big.Int, error) {
	return _Polygonzkevmmock.Contract.GetNextSnarkInput(&_Polygonzkevmmock.CallOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmmock.Contract.GlobalExitRootManager(&_Polygonzkevmmock.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmmock.Contract.GlobalExitRootManager(&_Polygonzkevmmock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmmock.Contract.IsEmergencyState(&_Polygonzkevmmock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmmock.Contract.IsEmergencyState(&_Polygonzkevmmock.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) IsForcedBatchDisallowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "isForcedBatchDisallowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevmmock.Contract.IsForcedBatchDisallowed(&_Polygonzkevmmock.CallOpts)
}

// IsForcedBatchDisallowed is a free data retrieval call binding the contract method 0xed6b0104.
//
// Solidity: function isForcedBatchDisallowed() view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) IsForcedBatchDisallowed() (bool, error) {
	return _Polygonzkevmmock.Contract.IsForcedBatchDisallowed(&_Polygonzkevmmock.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) IsPendingStateConsolidable(opts *bind.CallOpts, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "isPendingStateConsolidable", pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevmmock.Contract.IsPendingStateConsolidable(&_Polygonzkevmmock.CallOpts, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x383b3be8.
//
// Solidity: function isPendingStateConsolidable(uint64 pendingStateNum) view returns(bool)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) IsPendingStateConsolidable(pendingStateNum uint64) (bool, error) {
	return _Polygonzkevmmock.Contract.IsPendingStateConsolidable(&_Polygonzkevmmock.CallOpts, pendingStateNum)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastBatchSequenced(&_Polygonzkevmmock.CallOpts)
}

// LastBatchSequenced is a free data retrieval call binding the contract method 0x423fa856.
//
// Solidity: function lastBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastBatchSequenced() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastBatchSequenced(&_Polygonzkevmmock.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastForceBatch(&_Polygonzkevmmock.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastForceBatch(&_Polygonzkevmmock.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastForceBatchSequenced(&_Polygonzkevmmock.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastForceBatchSequenced(&_Polygonzkevmmock.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastPendingState(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastPendingState")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastPendingState() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastPendingState(&_Polygonzkevmmock.CallOpts)
}

// LastPendingState is a free data retrieval call binding the contract method 0x458c0477.
//
// Solidity: function lastPendingState() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastPendingState() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastPendingState(&_Polygonzkevmmock.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastPendingStateConsolidated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastPendingStateConsolidated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastPendingStateConsolidated(&_Polygonzkevmmock.CallOpts)
}

// LastPendingStateConsolidated is a free data retrieval call binding the contract method 0x4a1a89a7.
//
// Solidity: function lastPendingStateConsolidated() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastPendingStateConsolidated() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastPendingStateConsolidated(&_Polygonzkevmmock.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastTimestamp(&_Polygonzkevmmock.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastTimestamp() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastTimestamp(&_Polygonzkevmmock.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) LastVerifiedBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "lastVerifiedBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastVerifiedBatch(&_Polygonzkevmmock.CallOpts)
}

// LastVerifiedBatch is a free data retrieval call binding the contract method 0x7fcb3653.
//
// Solidity: function lastVerifiedBatch() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) LastVerifiedBatch() (uint64, error) {
	return _Polygonzkevmmock.Contract.LastVerifiedBatch(&_Polygonzkevmmock.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) Matic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "matic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) Matic() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Matic(&_Polygonzkevmmock.CallOpts)
}

// Matic is a free data retrieval call binding the contract method 0xb6b0b097.
//
// Solidity: function matic() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) Matic() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Matic(&_Polygonzkevmmock.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmmock *PolygonzkevmmockSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevmmock.Contract.MultiplierBatchFee(&_Polygonzkevmmock.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonzkevmmock.Contract.MultiplierBatchFee(&_Polygonzkevmmock.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockSession) NetworkName() (string, error) {
	return _Polygonzkevmmock.Contract.NetworkName(&_Polygonzkevmmock.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) NetworkName() (string, error) {
	return _Polygonzkevmmock.Contract.NetworkName(&_Polygonzkevmmock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) Owner() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Owner(&_Polygonzkevmmock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) Owner() (common.Address, error) {
	return _Polygonzkevmmock.Contract.Owner(&_Polygonzkevmmock.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmmock.Contract.PendingAdmin(&_Polygonzkevmmock.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonzkevmmock.Contract.PendingAdmin(&_Polygonzkevmmock.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.PendingStateTimeout(&_Polygonzkevmmock.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.PendingStateTimeout(&_Polygonzkevmmock.CallOpts)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) PendingStateTransitions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "pendingStateTransitions", arg0)

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
func (_Polygonzkevmmock *PolygonzkevmmockSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevmmock.Contract.PendingStateTransitions(&_Polygonzkevmmock.CallOpts, arg0)
}

// PendingStateTransitions is a free data retrieval call binding the contract method 0x837a4738.
//
// Solidity: function pendingStateTransitions(uint256 ) view returns(uint64 timestamp, uint64 lastVerifiedBatch, bytes32 exitRoot, bytes32 stateRoot)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) PendingStateTransitions(arg0 *big.Int) (struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}, error) {
	return _Polygonzkevmmock.Contract.PendingStateTransitions(&_Polygonzkevmmock.CallOpts, arg0)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) RollupVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "rollupVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevmmock.Contract.RollupVerifier(&_Polygonzkevmmock.CallOpts)
}

// RollupVerifier is a free data retrieval call binding the contract method 0xe8bf92ed.
//
// Solidity: function rollupVerifier() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) RollupVerifier() (common.Address, error) {
	return _Polygonzkevmmock.Contract.RollupVerifier(&_Polygonzkevmmock.CallOpts)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) SequencedBatches(opts *bind.CallOpts, arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "sequencedBatches", arg0)

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
func (_Polygonzkevmmock *PolygonzkevmmockSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevmmock.Contract.SequencedBatches(&_Polygonzkevmmock.CallOpts, arg0)
}

// SequencedBatches is a free data retrieval call binding the contract method 0xb4d63f58.
//
// Solidity: function sequencedBatches(uint64 ) view returns(bytes32 accInputHash, uint64 sequencedTimestamp, uint64 previousLastBatchSequenced)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) SequencedBatches(arg0 uint64) (struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}, error) {
	return _Polygonzkevmmock.Contract.SequencedBatches(&_Polygonzkevmmock.CallOpts, arg0)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) TrustedAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "trustedAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevmmock.Contract.TrustedAggregator(&_Polygonzkevmmock.CallOpts)
}

// TrustedAggregator is a free data retrieval call binding the contract method 0x29878983.
//
// Solidity: function trustedAggregator() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) TrustedAggregator() (common.Address, error) {
	return _Polygonzkevmmock.Contract.TrustedAggregator(&_Polygonzkevmmock.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.TrustedAggregatorTimeout(&_Polygonzkevmmock.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonzkevmmock.Contract.TrustedAggregatorTimeout(&_Polygonzkevmmock.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmmock.Contract.TrustedSequencer(&_Polygonzkevmmock.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonzkevmmock.Contract.TrustedSequencer(&_Polygonzkevmmock.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmmock.Contract.TrustedSequencerURL(&_Polygonzkevmmock.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonzkevmmock.Contract.TrustedSequencerURL(&_Polygonzkevmmock.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonzkevmmock.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevmmock.Contract.VerifyBatchTimeTarget(&_Polygonzkevmmock.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonzkevmmock *PolygonzkevmmockCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonzkevmmock.Contract.VerifyBatchTimeTarget(&_Polygonzkevmmock.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.AcceptAdminRole(&_Polygonzkevmmock.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.AcceptAdminRole(&_Polygonzkevmmock.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) ActivateEmergencyState(opts *bind.TransactOpts, sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "activateEmergencyState", sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ActivateEmergencyState(&_Polygonzkevmmock.TransactOpts, sequencedBatchNum)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x7215541a.
//
// Solidity: function activateEmergencyState(uint64 sequencedBatchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) ActivateEmergencyState(sequencedBatchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ActivateEmergencyState(&_Polygonzkevmmock.TransactOpts, sequencedBatchNum)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) ActivateForceBatches(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "activateForceBatches")
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ActivateForceBatches(&_Polygonzkevmmock.TransactOpts)
}

// ActivateForceBatches is a paid mutator transaction binding the contract method 0x5ec91958.
//
// Solidity: function activateForceBatches() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) ActivateForceBatches() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ActivateForceBatches(&_Polygonzkevmmock.TransactOpts)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) ConsolidatePendingState(opts *bind.TransactOpts, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "consolidatePendingState", pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ConsolidatePendingState(&_Polygonzkevmmock.TransactOpts, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x4a910e6a.
//
// Solidity: function consolidatePendingState(uint64 pendingStateNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) ConsolidatePendingState(pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ConsolidatePendingState(&_Polygonzkevmmock.TransactOpts, pendingStateNum)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.DeactivateEmergencyState(&_Polygonzkevmmock.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.DeactivateEmergencyState(&_Polygonzkevmmock.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "forceBatch", transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ForceBatch(&_Polygonzkevmmock.TransactOpts, transactions, maticAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 maticAmount) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) ForceBatch(transactions []byte, maticAmount *big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ForceBatch(&_Polygonzkevmmock.TransactOpts, transactions, maticAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) Initialize(opts *bind.TransactOpts, initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "initialize", initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.Initialize(&_Polygonzkevmmock.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2e129f9.
//
// Solidity: function initialize((address,address,uint64,address,uint64) initializePackedParameters, bytes32 genesisRoot, string _trustedSequencerURL, string _networkName, string _version) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) Initialize(initializePackedParameters PolygonZkEVMInitializePackedParameters, genesisRoot [32]byte, _trustedSequencerURL string, _networkName string, _version string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.Initialize(&_Polygonzkevmmock.TransactOpts, initializePackedParameters, genesisRoot, _trustedSequencerURL, _networkName, _version)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) OverridePendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "overridePendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.OverridePendingState(&_Polygonzkevmmock.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x2c1f816a.
//
// Solidity: function overridePendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) OverridePendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.OverridePendingState(&_Polygonzkevmmock.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "proveNonDeterministicPendingState", initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ProveNonDeterministicPendingState(&_Polygonzkevmmock.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x9aa972a3.
//
// Solidity: function proveNonDeterministicPendingState(uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) ProveNonDeterministicPendingState(initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.ProveNonDeterministicPendingState(&_Polygonzkevmmock.TransactOpts, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.RenounceOwnership(&_Polygonzkevmmock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.RenounceOwnership(&_Polygonzkevmmock.TransactOpts)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "sequenceBatches", batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SequenceBatches(batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SequenceBatches(&_Polygonzkevmmock.TransactOpts, batches, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0x5e9145c9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,uint64)[] batches, address l2Coinbase) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SequenceBatches(batches []PolygonZkEVMBatchData, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SequenceBatches(&_Polygonzkevmmock.TransactOpts, batches, l2Coinbase)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SequenceForceBatches(&_Polygonzkevmmock.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0xd8d1091b.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64)[] batches) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SequenceForceBatches(batches []PolygonZkEVMForcedBatchData) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SequenceForceBatches(&_Polygonzkevmmock.TransactOpts, batches)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetForceBatchTimeout(&_Polygonzkevmmock.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetForceBatchTimeout(&_Polygonzkevmmock.TransactOpts, newforceBatchTimeout)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetMultiplierBatchFee(&_Polygonzkevmmock.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetMultiplierBatchFee(&_Polygonzkevmmock.TransactOpts, newMultiplierBatchFee)
}

// SetNetworkName is a paid mutator transaction binding the contract method 0xc0cad302.
//
// Solidity: function setNetworkName(string _networkName) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetNetworkName(opts *bind.TransactOpts, _networkName string) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setNetworkName", _networkName)
}

// SetNetworkName is a paid mutator transaction binding the contract method 0xc0cad302.
//
// Solidity: function setNetworkName(string _networkName) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetNetworkName(_networkName string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetNetworkName(&_Polygonzkevmmock.TransactOpts, _networkName)
}

// SetNetworkName is a paid mutator transaction binding the contract method 0xc0cad302.
//
// Solidity: function setNetworkName(string _networkName) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetNetworkName(_networkName string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetNetworkName(&_Polygonzkevmmock.TransactOpts, _networkName)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetPendingStateTimeout(&_Polygonzkevmmock.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetPendingStateTimeout(&_Polygonzkevmmock.TransactOpts, newPendingStateTimeout)
}

// SetSequencedBatch is a paid mutator transaction binding the contract method 0x9b796760.
//
// Solidity: function setSequencedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetSequencedBatch(opts *bind.TransactOpts, _numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setSequencedBatch", _numBatch)
}

// SetSequencedBatch is a paid mutator transaction binding the contract method 0x9b796760.
//
// Solidity: function setSequencedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetSequencedBatch(_numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetSequencedBatch(&_Polygonzkevmmock.TransactOpts, _numBatch)
}

// SetSequencedBatch is a paid mutator transaction binding the contract method 0x9b796760.
//
// Solidity: function setSequencedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetSequencedBatch(_numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetSequencedBatch(&_Polygonzkevmmock.TransactOpts, _numBatch)
}

// SetSequencedBatches is a paid mutator transaction binding the contract method 0xe0d17441.
//
// Solidity: function setSequencedBatches(uint64 batchNum, bytes32 accInputData, uint64 timestamp, uint64 lastPendingStateConsolidated) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetSequencedBatches(opts *bind.TransactOpts, batchNum uint64, accInputData [32]byte, timestamp uint64, lastPendingStateConsolidated uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setSequencedBatches", batchNum, accInputData, timestamp, lastPendingStateConsolidated)
}

// SetSequencedBatches is a paid mutator transaction binding the contract method 0xe0d17441.
//
// Solidity: function setSequencedBatches(uint64 batchNum, bytes32 accInputData, uint64 timestamp, uint64 lastPendingStateConsolidated) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetSequencedBatches(batchNum uint64, accInputData [32]byte, timestamp uint64, lastPendingStateConsolidated uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetSequencedBatches(&_Polygonzkevmmock.TransactOpts, batchNum, accInputData, timestamp, lastPendingStateConsolidated)
}

// SetSequencedBatches is a paid mutator transaction binding the contract method 0xe0d17441.
//
// Solidity: function setSequencedBatches(uint64 batchNum, bytes32 accInputData, uint64 timestamp, uint64 lastPendingStateConsolidated) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetSequencedBatches(batchNum uint64, accInputData [32]byte, timestamp uint64, lastPendingStateConsolidated uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetSequencedBatches(&_Polygonzkevmmock.TransactOpts, batchNum, accInputData, timestamp, lastPendingStateConsolidated)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xfe16564f.
//
// Solidity: function setStateRoot(bytes32 newStateRoot, uint64 batchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetStateRoot(opts *bind.TransactOpts, newStateRoot [32]byte, batchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setStateRoot", newStateRoot, batchNum)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xfe16564f.
//
// Solidity: function setStateRoot(bytes32 newStateRoot, uint64 batchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetStateRoot(newStateRoot [32]byte, batchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetStateRoot(&_Polygonzkevmmock.TransactOpts, newStateRoot, batchNum)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xfe16564f.
//
// Solidity: function setStateRoot(bytes32 newStateRoot, uint64 batchNum) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetStateRoot(newStateRoot [32]byte, batchNum uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetStateRoot(&_Polygonzkevmmock.TransactOpts, newStateRoot, batchNum)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetTrustedAggregator(opts *bind.TransactOpts, newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setTrustedAggregator", newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedAggregator(&_Polygonzkevmmock.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregator is a paid mutator transaction binding the contract method 0xf14916d6.
//
// Solidity: function setTrustedAggregator(address newTrustedAggregator) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetTrustedAggregator(newTrustedAggregator common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedAggregator(&_Polygonzkevmmock.TransactOpts, newTrustedAggregator)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevmmock.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedAggregatorTimeout(&_Polygonzkevmmock.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedSequencer(&_Polygonzkevmmock.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedSequencer(&_Polygonzkevmmock.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedSequencerURL(&_Polygonzkevmmock.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetTrustedSequencerURL(&_Polygonzkevmmock.TransactOpts, newTrustedSequencerURL)
}

// SetVerifiedBatch is a paid mutator transaction binding the contract method 0x96dc3d39.
//
// Solidity: function setVerifiedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetVerifiedBatch(opts *bind.TransactOpts, _numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setVerifiedBatch", _numBatch)
}

// SetVerifiedBatch is a paid mutator transaction binding the contract method 0x96dc3d39.
//
// Solidity: function setVerifiedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetVerifiedBatch(_numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetVerifiedBatch(&_Polygonzkevmmock.TransactOpts, _numBatch)
}

// SetVerifiedBatch is a paid mutator transaction binding the contract method 0x96dc3d39.
//
// Solidity: function setVerifiedBatch(uint64 _numBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetVerifiedBatch(_numBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetVerifiedBatch(&_Polygonzkevmmock.TransactOpts, _numBatch)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevmmock.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.SetVerifyBatchTimeTarget(&_Polygonzkevmmock.TransactOpts, newVerifyBatchTimeTarget)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TransferAdminRole(&_Polygonzkevmmock.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TransferAdminRole(&_Polygonzkevmmock.TransactOpts, newPendingAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TransferOwnership(&_Polygonzkevmmock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TransferOwnership(&_Polygonzkevmmock.TransactOpts, newOwner)
}

// TrustedVerifyBatchesMock is a paid mutator transaction binding the contract method 0x10a01a72.
//
// Solidity: function trustedVerifyBatchesMock(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) TrustedVerifyBatchesMock(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "trustedVerifyBatchesMock", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proofA, proofB, proofC)
}

// TrustedVerifyBatchesMock is a paid mutator transaction binding the contract method 0x10a01a72.
//
// Solidity: function trustedVerifyBatchesMock(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) TrustedVerifyBatchesMock(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TrustedVerifyBatchesMock(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proofA, proofB, proofC)
}

// TrustedVerifyBatchesMock is a paid mutator transaction binding the contract method 0x10a01a72.
//
// Solidity: function trustedVerifyBatchesMock(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, uint256[2] proofA, uint256[2][2] proofB, uint256[2] proofC) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) TrustedVerifyBatchesMock(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proofA [2]*big.Int, proofB [2][2]*big.Int, proofC [2]*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.TrustedVerifyBatchesMock(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proofA, proofB, proofC)
}

// UpdateBatchFee is a paid mutator transaction binding the contract method 0xb4f77ea9.
//
// Solidity: function updateBatchFee(uint64 newLastVerifiedBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) UpdateBatchFee(opts *bind.TransactOpts, newLastVerifiedBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "updateBatchFee", newLastVerifiedBatch)
}

// UpdateBatchFee is a paid mutator transaction binding the contract method 0xb4f77ea9.
//
// Solidity: function updateBatchFee(uint64 newLastVerifiedBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) UpdateBatchFee(newLastVerifiedBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.UpdateBatchFee(&_Polygonzkevmmock.TransactOpts, newLastVerifiedBatch)
}

// UpdateBatchFee is a paid mutator transaction binding the contract method 0xb4f77ea9.
//
// Solidity: function updateBatchFee(uint64 newLastVerifiedBatch) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) UpdateBatchFee(newLastVerifiedBatch uint64) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.UpdateBatchFee(&_Polygonzkevmmock.TransactOpts, newLastVerifiedBatch)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) VerifyBatches(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "verifyBatches", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.VerifyBatches(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x621dd411.
//
// Solidity: function verifyBatches(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) VerifyBatches(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.VerifyBatches(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.contract.Transact(opts, "verifyBatchesTrustedAggregator", pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x2b0006fa.
//
// Solidity: function verifyBatchesTrustedAggregator(uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonzkevmmock *PolygonzkevmmockTransactorSession) VerifyBatchesTrustedAggregator(pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmmock.Contract.VerifyBatchesTrustedAggregator(&_Polygonzkevmmock.TransactOpts, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// PolygonzkevmmockAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockAcceptAdminRoleIterator struct {
	Event *PolygonzkevmmockAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockAcceptAdminRole)
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
		it.Event = new(PolygonzkevmmockAcceptAdminRole)
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
func (it *PolygonzkevmmockAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonzkevmmockAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockAcceptAdminRoleIterator{contract: _Polygonzkevmmock.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockAcceptAdminRole)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonzkevmmockAcceptAdminRole, error) {
	event := new(PolygonzkevmmockAcceptAdminRole)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockActivateForceBatchesIterator is returned from FilterActivateForceBatches and is used to iterate over the raw logs and unpacked data for ActivateForceBatches events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockActivateForceBatchesIterator struct {
	Event *PolygonzkevmmockActivateForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockActivateForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockActivateForceBatches)
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
		it.Event = new(PolygonzkevmmockActivateForceBatches)
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
func (it *PolygonzkevmmockActivateForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockActivateForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockActivateForceBatches represents a ActivateForceBatches event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockActivateForceBatches struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterActivateForceBatches is a free log retrieval operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterActivateForceBatches(opts *bind.FilterOpts) (*PolygonzkevmmockActivateForceBatchesIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockActivateForceBatchesIterator{contract: _Polygonzkevmmock.contract, event: "ActivateForceBatches", logs: logs, sub: sub}, nil
}

// WatchActivateForceBatches is a free log subscription operation binding the contract event 0x854dd6ce5a1445c4c54388b21cffd11cf5bba1b9e763aec48ce3da75d617412f.
//
// Solidity: event ActivateForceBatches()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchActivateForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockActivateForceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "ActivateForceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockActivateForceBatches)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseActivateForceBatches(log types.Log) (*PolygonzkevmmockActivateForceBatches, error) {
	event := new(PolygonzkevmmockActivateForceBatches)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "ActivateForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockConsolidatePendingStateIterator struct {
	Event *PolygonzkevmmockConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockConsolidatePendingState)
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
		it.Event = new(PolygonzkevmmockConsolidatePendingState)
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
func (it *PolygonzkevmmockConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockConsolidatePendingState struct {
	NumBatch        uint64
	StateRoot       [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, numBatch []uint64, pendingStateNum []uint64) (*PolygonzkevmmockConsolidatePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockConsolidatePendingStateIterator{contract: _Polygonzkevmmock.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x328d3c6c0fd6f1be0515e422f2d87e59f25922cbc2233568515a0c4bc3f8510e.
//
// Solidity: event ConsolidatePendingState(uint64 indexed numBatch, bytes32 stateRoot, uint64 indexed pendingStateNum)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockConsolidatePendingState, numBatch []uint64, pendingStateNum []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var pendingStateNumRule []interface{}
	for _, pendingStateNumItem := range pendingStateNum {
		pendingStateNumRule = append(pendingStateNumRule, pendingStateNumItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "ConsolidatePendingState", numBatchRule, pendingStateNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockConsolidatePendingState)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonzkevmmockConsolidatePendingState, error) {
	event := new(PolygonzkevmmockConsolidatePendingState)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockEmergencyStateActivatedIterator struct {
	Event *PolygonzkevmmockEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockEmergencyStateActivated)
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
		it.Event = new(PolygonzkevmmockEmergencyStateActivated)
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
func (it *PolygonzkevmmockEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonzkevmmockEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockEmergencyStateActivatedIterator{contract: _Polygonzkevmmock.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockEmergencyStateActivated)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonzkevmmockEmergencyStateActivated, error) {
	event := new(PolygonzkevmmockEmergencyStateActivated)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockEmergencyStateDeactivatedIterator struct {
	Event *PolygonzkevmmockEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockEmergencyStateDeactivated)
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
		it.Event = new(PolygonzkevmmockEmergencyStateDeactivated)
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
func (it *PolygonzkevmmockEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonzkevmmockEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockEmergencyStateDeactivatedIterator{contract: _Polygonzkevmmock.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockEmergencyStateDeactivated)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonzkevmmockEmergencyStateDeactivated, error) {
	event := new(PolygonzkevmmockEmergencyStateDeactivated)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockForceBatchIterator struct {
	Event *PolygonzkevmmockForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockForceBatch)
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
		it.Event = new(PolygonzkevmmockForceBatch)
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
func (it *PolygonzkevmmockForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockForceBatch represents a ForceBatch event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonzkevmmockForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockForceBatchIterator{contract: _Polygonzkevmmock.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockForceBatch)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseForceBatch(log types.Log) (*PolygonzkevmmockForceBatch, error) {
	event := new(PolygonzkevmmockForceBatch)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockInitializedIterator struct {
	Event *PolygonzkevmmockInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockInitialized)
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
		it.Event = new(PolygonzkevmmockInitialized)
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
func (it *PolygonzkevmmockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockInitialized represents a Initialized event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmmockInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockInitializedIterator{contract: _Polygonzkevmmock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockInitialized)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseInitialized(log types.Log) (*PolygonzkevmmockInitialized, error) {
	event := new(PolygonzkevmmockInitialized)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockOverridePendingStateIterator struct {
	Event *PolygonzkevmmockOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockOverridePendingState)
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
		it.Event = new(PolygonzkevmmockOverridePendingState)
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
func (it *PolygonzkevmmockOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockOverridePendingState represents a OverridePendingState event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockOverridePendingState struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterOverridePendingState(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmmockOverridePendingStateIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockOverridePendingStateIterator{contract: _Polygonzkevmmock.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0xcc1b5520188bf1dd3e63f98164b577c4d75c11a619ddea692112f0d1aec4cf72.
//
// Solidity: event OverridePendingState(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockOverridePendingState, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "OverridePendingState", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockOverridePendingState)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseOverridePendingState(log types.Log) (*PolygonzkevmmockOverridePendingState, error) {
	event := new(PolygonzkevmmockOverridePendingState)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockOwnershipTransferredIterator struct {
	Event *PolygonzkevmmockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockOwnershipTransferred)
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
		it.Event = new(PolygonzkevmmockOwnershipTransferred)
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
func (it *PolygonzkevmmockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockOwnershipTransferred represents a OwnershipTransferred event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonzkevmmockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockOwnershipTransferredIterator{contract: _Polygonzkevmmock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockOwnershipTransferred)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonzkevmmockOwnershipTransferred, error) {
	event := new(PolygonzkevmmockOwnershipTransferred)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockProveNonDeterministicPendingStateIterator struct {
	Event *PolygonzkevmmockProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockProveNonDeterministicPendingState)
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
		it.Event = new(PolygonzkevmmockProveNonDeterministicPendingState)
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
func (it *PolygonzkevmmockProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonzkevmmockProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockProveNonDeterministicPendingStateIterator{contract: _Polygonzkevmmock.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockProveNonDeterministicPendingState)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonzkevmmockProveNonDeterministicPendingState, error) {
	event := new(PolygonzkevmmockProveNonDeterministicPendingState)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSequenceBatchesIterator struct {
	Event *PolygonzkevmmockSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSequenceBatches)
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
		it.Event = new(PolygonzkevmmockSequenceBatches)
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
func (it *PolygonzkevmmockSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSequenceBatches represents a SequenceBatches event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSequenceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmmockSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSequenceBatchesIterator{contract: _Polygonzkevmmock.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x303446e6a8cb73c83dff421c0b1d5e5ce0719dab1bff13660fc254e58cc17fce.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSequenceBatches)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSequenceBatches(log types.Log) (*PolygonzkevmmockSequenceBatches, error) {
	event := new(PolygonzkevmmockSequenceBatches)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSequenceForceBatchesIterator struct {
	Event *PolygonzkevmmockSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSequenceForceBatches)
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
		it.Event = new(PolygonzkevmmockSequenceForceBatches)
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
func (it *PolygonzkevmmockSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonzkevmmockSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSequenceForceBatchesIterator{contract: _Polygonzkevmmock.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSequenceForceBatches)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonzkevmmockSequenceForceBatches, error) {
	event := new(PolygonzkevmmockSequenceForceBatches)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetForceBatchTimeoutIterator struct {
	Event *PolygonzkevmmockSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetForceBatchTimeout)
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
		it.Event = new(PolygonzkevmmockSetForceBatchTimeout)
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
func (it *PolygonzkevmmockSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonzkevmmockSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetForceBatchTimeoutIterator{contract: _Polygonzkevmmock.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetForceBatchTimeout)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonzkevmmockSetForceBatchTimeout, error) {
	event := new(PolygonzkevmmockSetForceBatchTimeout)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetMultiplierBatchFeeIterator struct {
	Event *PolygonzkevmmockSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetMultiplierBatchFee)
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
		it.Event = new(PolygonzkevmmockSetMultiplierBatchFee)
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
func (it *PolygonzkevmmockSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonzkevmmockSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetMultiplierBatchFeeIterator{contract: _Polygonzkevmmock.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetMultiplierBatchFee)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonzkevmmockSetMultiplierBatchFee, error) {
	event := new(PolygonzkevmmockSetMultiplierBatchFee)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetPendingStateTimeoutIterator struct {
	Event *PolygonzkevmmockSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetPendingStateTimeout)
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
		it.Event = new(PolygonzkevmmockSetPendingStateTimeout)
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
func (it *PolygonzkevmmockSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonzkevmmockSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetPendingStateTimeoutIterator{contract: _Polygonzkevmmock.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetPendingStateTimeout)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonzkevmmockSetPendingStateTimeout, error) {
	event := new(PolygonzkevmmockSetPendingStateTimeout)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedAggregatorIterator struct {
	Event *PolygonzkevmmockSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetTrustedAggregator)
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
		it.Event = new(PolygonzkevmmockSetTrustedAggregator)
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
func (it *PolygonzkevmmockSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonzkevmmockSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetTrustedAggregatorIterator{contract: _Polygonzkevmmock.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetTrustedAggregator)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonzkevmmockSetTrustedAggregator, error) {
	event := new(PolygonzkevmmockSetTrustedAggregator)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonzkevmmockSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonzkevmmockSetTrustedAggregatorTimeout)
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
func (it *PolygonzkevmmockSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonzkevmmockSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetTrustedAggregatorTimeoutIterator{contract: _Polygonzkevmmock.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetTrustedAggregatorTimeout)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonzkevmmockSetTrustedAggregatorTimeout, error) {
	event := new(PolygonzkevmmockSetTrustedAggregatorTimeout)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedSequencerIterator struct {
	Event *PolygonzkevmmockSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetTrustedSequencer)
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
		it.Event = new(PolygonzkevmmockSetTrustedSequencer)
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
func (it *PolygonzkevmmockSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonzkevmmockSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetTrustedSequencerIterator{contract: _Polygonzkevmmock.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetTrustedSequencer)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonzkevmmockSetTrustedSequencer, error) {
	event := new(PolygonzkevmmockSetTrustedSequencer)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedSequencerURLIterator struct {
	Event *PolygonzkevmmockSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetTrustedSequencerURL)
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
		it.Event = new(PolygonzkevmmockSetTrustedSequencerURL)
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
func (it *PolygonzkevmmockSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonzkevmmockSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetTrustedSequencerURLIterator{contract: _Polygonzkevmmock.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetTrustedSequencerURL)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonzkevmmockSetTrustedSequencerURL, error) {
	event := new(PolygonzkevmmockSetTrustedSequencerURL)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonzkevmmockSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonzkevmmockSetVerifyBatchTimeTarget)
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
func (it *PolygonzkevmmockSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonzkevmmockSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockSetVerifyBatchTimeTargetIterator{contract: _Polygonzkevmmock.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockSetVerifyBatchTimeTarget)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonzkevmmockSetVerifyBatchTimeTarget, error) {
	event := new(PolygonzkevmmockSetVerifyBatchTimeTarget)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockTransferAdminRoleIterator struct {
	Event *PolygonzkevmmockTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockTransferAdminRole)
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
		it.Event = new(PolygonzkevmmockTransferAdminRole)
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
func (it *PolygonzkevmmockTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockTransferAdminRole represents a TransferAdminRole event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonzkevmmockTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockTransferAdminRoleIterator{contract: _Polygonzkevmmock.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockTransferAdminRole)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseTransferAdminRole(log types.Log) (*PolygonzkevmmockTransferAdminRole, error) {
	event := new(PolygonzkevmmockTransferAdminRole)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockUpdateZkEVMVersionIterator is returned from FilterUpdateZkEVMVersion and is used to iterate over the raw logs and unpacked data for UpdateZkEVMVersion events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockUpdateZkEVMVersionIterator struct {
	Event *PolygonzkevmmockUpdateZkEVMVersion // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockUpdateZkEVMVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockUpdateZkEVMVersion)
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
		it.Event = new(PolygonzkevmmockUpdateZkEVMVersion)
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
func (it *PolygonzkevmmockUpdateZkEVMVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockUpdateZkEVMVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockUpdateZkEVMVersion represents a UpdateZkEVMVersion event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockUpdateZkEVMVersion struct {
	NumBatch uint64
	ForkID   uint64
	Version  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateZkEVMVersion is a free log retrieval operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterUpdateZkEVMVersion(opts *bind.FilterOpts) (*PolygonzkevmmockUpdateZkEVMVersionIterator, error) {

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockUpdateZkEVMVersionIterator{contract: _Polygonzkevmmock.contract, event: "UpdateZkEVMVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateZkEVMVersion is a free log subscription operation binding the contract event 0xed7be53c9f1a96a481223b15568a5b1a475e01a74b347d6ca187c8bf0c078cd6.
//
// Solidity: event UpdateZkEVMVersion(uint64 numBatch, uint64 forkID, string version)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchUpdateZkEVMVersion(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockUpdateZkEVMVersion) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "UpdateZkEVMVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockUpdateZkEVMVersion)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseUpdateZkEVMVersion(log types.Log) (*PolygonzkevmmockUpdateZkEVMVersion, error) {
	event := new(PolygonzkevmmockUpdateZkEVMVersion)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "UpdateZkEVMVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockVerifyBatchesIterator struct {
	Event *PolygonzkevmmockVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockVerifyBatches)
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
		it.Event = new(PolygonzkevmmockVerifyBatches)
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
func (it *PolygonzkevmmockVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockVerifyBatches represents a VerifyBatches event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmmockVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockVerifyBatchesIterator{contract: _Polygonzkevmmock.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockVerifyBatches)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseVerifyBatches(log types.Log) (*PolygonzkevmmockVerifyBatches, error) {
	event := new(PolygonzkevmmockVerifyBatches)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonzkevmmock contract.
type PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonzkevmmockVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmmockVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonzkevmmockVerifyBatchesTrustedAggregator)
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
func (it *PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmmockVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonzkevmmock contract.
type PolygonzkevmmockVerifyBatchesTrustedAggregator struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmmockVerifyBatchesTrustedAggregatorIterator{contract: _Polygonzkevmmock.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xcb339b570a7f0b25afa7333371ff11192092a0aeace12b671f4c212f2815c6fe.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonzkevmmockVerifyBatchesTrustedAggregator, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonzkevmmock.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmmockVerifyBatchesTrustedAggregator)
				if err := _Polygonzkevmmock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Polygonzkevmmock *PolygonzkevmmockFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonzkevmmockVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonzkevmmockVerifyBatchesTrustedAggregator)
	if err := _Polygonzkevmmock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

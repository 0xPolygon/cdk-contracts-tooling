// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanagerprevious

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

// LegacyZKEVMStateVariablesPendingState is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesPendingState struct {
	Timestamp         uint64
	LastVerifiedBatch uint64
	ExitRoot          [32]byte
	StateRoot         [32]byte
}

// LegacyZKEVMStateVariablesSequencedBatchData is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesSequencedBatchData struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}

// PolygonrollupmanagerpreviousMetaData contains all meta data concerning the Polygonrollupmanagerprevious contract.
var PolygonrollupmanagerpreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupPendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structLegacyZKEVMStateVariables.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051616e1f380380616e1f83398101604081905261002e9161012b565b6001600160a01b0380841660805282811660c052811660a05261004f610057565b505050610175565b5f54610100900460ff16156100c25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610112575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610128575f5ffd5b50565b5f5f5f6060848603121561013d575f5ffd5b835161014881610114565b602085015190935061015981610114565b604085015190925061016a81610114565b809150509250925092565b60805160a05160c051616c486101d75f395f8181610ad5015281816126720152613fe201525f81816108620152818161307e01526143a701525f8181610a2601528181610dab01528181610fc101528181611e6e01526142860152616c485ff3fe608060405234801561000f575f5ffd5b5060043610610351575f3560e01c80638bd4f071116101be578063c4c928c2116100fe578063dde0ff771161009e578063e46761c411610079578063e46761c414610ad0578063f34eb8eb14610af7578063f4e9267514610b0a578063f9c4c2ae14610b1a575f5ffd5b8063dde0ff7714610a8a578063dfdb8c5e14610aaa578063e0bfd3d214610abd575f5ffd5b8063d5073f6f116100d9578063d5073f6f14610a48578063d547741f14610a5b578063d939b31514610a6e578063dbc1697614610a82575f5ffd5b8063c4c928c2146109e9578063ceee281d146109fc578063d02103ca14610a21575f5ffd5b8063a066215c11610169578063a3c573eb11610144578063a3c573eb1461085d578063afd23cbe146108a9578063b99d0ad7146108de578063c1acbc34146109c1575f5ffd5b8063a066215c1461083b578063a217fddf1461084e578063a2967d9914610855575f5ffd5b806399f5634e1161019957806399f5634e1461080d5780639a908e73146108155780639c9f3dfe14610828575f5ffd5b80638bd4f071146107a25780638fd88cc2146107b557806391d14854146107c8575f5ffd5b80632f2ff15d1161029457806365c0504d116102345780637975fcfe1161020f5780637975fcfe1461071a5780637fb6e76a1461073a578063841b24d71461075f57806387c20c011461078f575f5ffd5b806365c0504d146105ea5780637222020f146106f4578063727885e914610707575f5ffd5b8063394218e91161026f578063394218e914610586578063477fa2701461059957806355a71ee0146105a157806360469169146105e2575f5ffd5b80632f2ff15d1461054c57806330c27dde1461055f57806336568abe14610573575f5ffd5b806315064c96116102ff5780631816b7e5116102da5780631816b7e51461044a5780632072f6c51461045d578063248a9ca3146104655780632528016914610495575f5ffd5b806315064c96146104055780631608859c146104125780631796a1ae14610425575f5ffd5b806311f6b2871161032f57806311f6b287146103ca57806312b86e19146103dd5780631489ed10146103f2575f5ffd5b8063066ec01214610355578063080b3111146103875780630a0d9fbe146103aa575b5f5ffd5b6084546103699067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b61039a6103953660046156ac565b610c80565b604051901515815260200161037e565b6085546103699068010000000000000000900467ffffffffffffffff1681565b6103696103d83660046156dd565b610ca7565b6103f06103eb366004615707565b610cc4565b005b6103f06104003660046157ad565b610eda565b606f5461039a9060ff1681565b6103f06104203660046156ac565b6110af565b607e546104359063ffffffff1681565b60405163ffffffff909116815260200161037e565b6103f0610458366004615825565b611184565b6103f061127d565b61048761047336600461584d565b5f9081526034602052604090206001015490565b60405190815260200161037e565b6105186104a33660046156ac565b60408051606080820183525f808352602080840182905292840181905263ffffffff9590951685526081825282852067ffffffffffffffff9485168652600301825293829020825194850183528054855260010154808416918501919091526801000000000000000090049091169082015290565b604080518251815260208084015167ffffffffffffffff90811691830191909152928201519092169082015260600161037e565b6103f061055a366004615864565b611387565b6087546103699067ffffffffffffffff1681565b6103f0610581366004615864565b6113ab565b6103f0610594366004615892565b611408565b608654610487565b6104876105af3660046156ac565b63ffffffff82165f90815260816020908152604080832067ffffffffffffffff8516845260020190915290205492915050565b61048761151e565b61069d6105f83660046156dd565b607f6020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff918216929182169167ffffffffffffffff740100000000000000000000000000000000000000008204169160ff7c010000000000000000000000000000000000000000000000000000000083048116927d0100000000000000000000000000000000000000000000000000000000009004169086565b6040805173ffffffffffffffffffffffffffffffffffffffff978816815296909516602087015267ffffffffffffffff9093169385019390935260ff166060840152901515608083015260a082015260c00161037e565b6103f06107023660046156dd565b611533565b6103f061071536600461598b565b6116a9565b61072d610728366004615a46565b611c48565b60405161037e9190615aed565b610435610748366004615892565b60836020525f908152604090205463ffffffff1681565b608454610369907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6103f061079d3660046157ad565b611c78565b6103f06107b0366004615707565b61205e565b6103f06107c3366004615aff565b61212a565b61039a6107d6366004615864565b5f91825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61048761262b565b610369610823366004615b1b565b612731565b6103f0610836366004615892565b612978565b6103f0610849366004615892565b612a5c565b6104875f81565b610487612b42565b6108847f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161037e565b6085546108cb90700100000000000000000000000000000000900461ffff1681565b60405161ffff909116815260200161037e565b6109766108ec3660046156ac565b60408051608080820183525f8083526020808401829052838501829052606093840182905263ffffffff9690961681526081865283812067ffffffffffffffff958616825260040186528390208351918201845280548086168352680100000000000000009004909416948101949094526001830154918401919091526002909101549082015290565b60405161037e91905f60808201905067ffffffffffffffff835116825267ffffffffffffffff6020840151166020830152604083015160408301526060830151606083015292915050565b60845461036990700100000000000000000000000000000000900467ffffffffffffffff1681565b6103f06109f7366004615b43565b612ee6565b610435610a0a366004615bb2565b60826020525f908152604090205463ffffffff1681565b6108847f000000000000000000000000000000000000000000000000000000000000000081565b6103f0610a5636600461584d565b612f21565b6103f0610a69366004615864565b612fd1565b6085546103699067ffffffffffffffff1681565b6103f0612ff5565b6084546103699068010000000000000000900467ffffffffffffffff1681565b6103f0610ab8366004615bcd565b6130f9565b6103f0610acb366004615c07565b6132cc565b6108847f000000000000000000000000000000000000000000000000000000000000000081565b6103f0610b05366004615c74565b613421565b6080546104359063ffffffff1681565b610bf3610b283660046156dd565b60816020525f90815260409020805460018201546005830154600684015460079094015473ffffffffffffffffffffffffffffffffffffffff80851695740100000000000000000000000000000000000000009586900467ffffffffffffffff908116969286169592909204821693928282169268010000000000000000808404821693700100000000000000000000000000000000808204841694780100000000000000000000000000000000000000000000000090920484169380831693830416910460ff168c565b6040805173ffffffffffffffffffffffffffffffffffffffff9d8e16815267ffffffffffffffff9c8d1660208201529c909a16998c019990995296891660608b015260808a019590955292871660a089015290861660c0880152851660e0870152841661010086015283166101208501529190911661014083015260ff166101608201526101800161037e565b63ffffffff82165f908152608160205260408120610c9e9083613660565b90505b92915050565b63ffffffff81165f908152608160205260408120610ca1906136a4565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4610cee81613739565b63ffffffff89165f908152608160205260409020610d12818a8a8a8a8a8a8a613743565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f908152600284016020526040902086905560058301879055905470010000000000000000000000000000000090041615610da9576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d610ded612b42565b6040518263ffffffff1660e01b8152600401610e0b91815260200190565b5f604051808303815f87803b158015610e22575f5ffd5b505af1158015610e34573d5f5f3e3d5ffd5b50506084805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a8000000000000000000000000000000000000000000000000017905550506040805167ffffffffffffffff881681526020810186905290810186905233606082015263ffffffff8b16907f3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e9060800160405180910390a250505050505050505050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4610f0481613739565b63ffffffff89165f908152608160205260409020610f28818a8a8a8a8a8a8a613c07565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8a81169182029290921783555f908152600284016020526040902087905560058301889055905470010000000000000000000000000000000090041615610fbf576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d611003612b42565b6040518263ffffffff1660e01b815260040161102191815260200190565b5f604051808303815f87803b158015611038575f5ffd5b505af115801561104a573d5f5f3e3d5ffd5b50506040805167ffffffffffffffff8b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b63ffffffff82165f9081526081602090815260408083203384527fc17b14a573f65366cdad721c7c0a0f76536bb4a86b935cdac44610e4f010b52a9092529091205460ff1661117557606f5460ff1615611135576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61113f8183613660565b611175576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61117f8183614128565b505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db16111ae81613739565b6103e88261ffff1610806111c757506103ff8261ffff16115b156111fe576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000061ffff8516908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a15050565b335f9081527f8875b94af5657a2903def9906d67a3f42d8a836d24b5602c00f00fc855339fcd602052604090205460ff1661137d57608454700100000000000000000000000000000000900467ffffffffffffffff1615806113175750608454429061130b9062093a8090700100000000000000000000000000000000900467ffffffffffffffff16615d28565b67ffffffffffffffff16115b806113465750608754429061133a9062093a809067ffffffffffffffff16615d28565b67ffffffffffffffff16115b1561137d576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113856143a5565b565b5f828152603460205260409020600101546113a181613739565b61117f8383614428565b73ffffffffffffffffffffffffffffffffffffffff811633146113fa576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61140482826144e3565b5050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db161143281613739565b606f5460ff166114a15760845467ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116908316106114a1576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6084805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190602001611271565b5f608654606461152e9190615d48565b905090565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd61155d81613739565b63ffffffff8216158061157b5750607e5463ffffffff908116908316115b156115b2576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff161515900361162a576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167d01000000000000000000000000000000000000000000000000000000000017905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd086116d381613739565b63ffffffff881615806116f15750607e5463ffffffff908116908916115b15611728576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff16151590036117a0576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff67ffffffffffffffff891611156117e8576040517f4c753f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88165f9081526083602052604090205463ffffffff161561183e576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608080545f919082906118569063ffffffff16615d5f565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f808252602082019283905293945073ffffffffffffffffffffffffffffffffffffffff9092169130916118ae90615675565b6118ba93929190615d83565b604051809103905ff0801580156118d3573d5f5f3e3d5ffd5b5090508160835f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508160825f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8463ffffffff1663ffffffff1681526020019081526020015f20905081815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360010160149054906101000a900467ffffffffffffffff168160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a815f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508360020154816002015f5f67ffffffffffffffff1681526020019081526020015f20819055508b63ffffffff168160070160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083600101601c9054906101000a900460ff168160070160106101000a81548160ff021916908360ff1602179055508263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418d848e8c604051611ba9949392919063ffffffff94909416845273ffffffffffffffffffffffffffffffffffffffff928316602085015267ffffffffffffffff91909116604084015216606082015260800190565b60405180910390a26040517f7125702200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831690637125702290611c0d908d908d9088908e908e908e90600401615dd6565b5f604051808303815f87803b158015611c24575f5ffd5b505af1158015611c36573d5f5f3e3d5ffd5b50505050505050505050505050505050565b63ffffffff86165f908152608160205260409020606090611c6d90878787878761459c565b979650505050505050565b606f5460ff1615611cb5576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f90815260816020908152604080832060845467ffffffffffffffff8a811686526003830190945291909320600101544292611d17927801000000000000000000000000000000000000000000000000900481169116615d28565b67ffffffffffffffff161115611d59576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8611d668888615e67565b67ffffffffffffffff161115611da8576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611db88189898989898989613c07565b611dc281876146ed565b60855467ffffffffffffffff165f03611f00576006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f908152600284016020526040902086905560058301879055905470010000000000000000000000000000000090041615611e6c576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d611eb0612b42565b6040518263ffffffff1660e01b8152600401611ece91815260200190565b5f604051808303815f87803b158015611ee5575f5ffd5b505af1158015611ef7573d5f5f3e3d5ffd5b50505050611ffe565b611f09816148d3565b600681018054700100000000000000000000000000000000900467ffffffffffffffff16906010611f3983615e87565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815289831660208083019182528284018b8152606084018b81526006890154700100000000000000000000000000000000900487165f90815260048a0190935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b6040805167ffffffffffffffff8816815260208101869052908101869052339063ffffffff8b16907faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b49060600160405180910390a3505050505050505050565b606f5460ff161561209b576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f9081526081602052604090206120bf8189898989898989613743565b67ffffffffffffffff87165f9081526004820160209081526040918290206002015482519081529081018590527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a161211f6143a5565b505050505050505050565b335f9081527ff14f5a8ad59d90593602e905b358229bff5ceea677d5bf0f5a1701793550a9a6602052604090205460ff1615801561220057503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121e79190615eaa565b73ffffffffffffffffffffffffffffffffffffffff1614155b15612237576040517f1a06d0fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526082602052604081205463ffffffff169081900361229c576040517f74a086a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff81165f908152608160205260409020600681015467ffffffffffffffff908116908416811115806122ef5750600682015467ffffffffffffffff680100000000000000009091048116908516105b15612326576040517fcb23ebdf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805b8467ffffffffffffffff168167ffffffffffffffff16146123f85767ffffffffffffffff8082165f90815260038501602052604090206001015468010000000000000000900481169086168110156123ac576040517f9753965f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff9091165f908152600384016020526040812090815560010180547fffffffffffffffffffffffffffffffff00000000000000000000000000000000169055612328565b6006830180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff87161790556124388583615e67565b608480545f9061245390849067ffffffffffffffff16615e67565b82546101009290920a67ffffffffffffffff818102199093169183160217909155600685015470010000000000000000000000000000000090041615905061252f575f61249f846136a4565b60068501549091506124c79068010000000000000000900467ffffffffffffffff1682615e67565b608480546008906124ef90849068010000000000000000900467ffffffffffffffff16615e67565b825467ffffffffffffffff9182166101009390930a928302919092021990911617905550506006830180546fffffffffffffffffffffffffffffffff1690555b67ffffffffffffffff85165f818152600385016020526040908190205490517f669adece0000000000000000000000000000000000000000000000000000000081526004810192909252602482015273ffffffffffffffffffffffffffffffffffffffff87169063669adece906044015f604051808303815f87803b1580156125b6575f5ffd5b505af11580156125c8573d5f5f3e3d5ffd5b5050505067ffffffffffffffff85165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a3505050505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156126b7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906126db9190615ec5565b6084549091505f906127059067ffffffffffffffff68010000000000000000820481169116615e67565b67ffffffffffffffff169050805f03612720575f9250505090565b61272a8183615f09565b9250505090565b606f545f9060ff1615612770576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f9081526082602052604081205463ffffffff16908190036127bf576040517f71653c1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8367ffffffffffffffff165f03612802576040517f2590ccf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff81165f908152608160205260408120608480549192879261283390849067ffffffffffffffff16615d28565b82546101009290920a67ffffffffffffffff81810219909316918316021790915560068301541690505f6128678783615d28565b60068401805467ffffffffffffffff8084167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217909255604080516060810182528a815242841660208083019182528886168385019081525f95865260038b019091529290932090518155915160019290920180549151841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921692909316919091171790559050612929836148d3565b60405167ffffffffffffffff8216815263ffffffff8516907f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a259060200160405180910390a29695505050505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db16129a281613739565b606f5460ff166129f55760855467ffffffffffffffff908116908316106129f5576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff84169081179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590602001611271565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1612a8681613739565b620151808267ffffffffffffffff161115612acd576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890602001611271565b6080545f9063ffffffff16808203612b5b57505f919050565b5f8167ffffffffffffffff811115612b7557612b756158ab565b604051908082528060200260200182016040528015612b9e578160200160208202803683370190505b5090505f5b82811015612bfb5760815f612bb9836001615f1c565b63ffffffff1663ffffffff1681526020019081526020015f2060050154828281518110612be857612be8615f2f565b6020908102919091010152600101612ba3565b505f60205b83600114612e17575f612c14600286615f5c565b612c1f600287615f09565b612c299190615f1c565b90505f8167ffffffffffffffff811115612c4557612c456158ab565b604051908082528060200260200182016040528015612c6e578160200160208202803683370190505b5090505f5b82811015612dc757612c86600184615f6f565b81148015612c9e5750612c9a600288615f5c565b6001145b15612d1b5785612caf826002615d48565b81518110612cbf57612cbf615f2f565b602002602001015185604051602001612ce2929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110612d0a57612d0a615f2f565b602002602001018181525050612dbf565b85612d27826002615d48565b81518110612d3757612d37615f2f565b602002602001015186826002612d4d9190615d48565b612d58906001615f1c565b81518110612d6857612d68615f2f565b6020026020010151604051602001612d8a929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110612db257612db2615f2f565b6020026020010181815250505b600101612c73565b508094508195508384604051602001612dea929190918252602082015260400190565b6040516020818303038152906040528051906020012093508280612e0d90615f82565b9350505050612c00565b5f835f81518110612e2a57612e2a615f2f565b602002602001015190505f5f90505b82811015612edc576040805160208101849052908101859052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083018790529082018690529250606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209350600101612e39565b5095945050505050565b7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac612f1081613739565b612f1b8484846149d0565b50505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb612f4b81613739565b683635c9adc5dea00000821180612f655750633b9aca0082105b15612f9c576040517f8586952500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b290602001611271565b5f82815260346020526040902060010154612feb81613739565b61117f83836144e3565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f461301f81613739565b608780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff16179055604080517fdbc1697600000000000000000000000000000000000000000000000000000000815290517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169163dbc16976916004808301925f92919082900301818387803b1580156130d8575f5ffd5b505af11580156130ea573d5f5f3e3d5ffd5b505050506130f6614e55565b50565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015613159573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061317d9190615eaa565b73ffffffffffffffffffffffffffffffffffffffff16146131ca576040517f696072e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526082602090815260408083205463ffffffff16835260819091529020600681015467ffffffffffffffff808216680100000000000000009092041614613257576040517fcc862d4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600781015463ffffffff83166801000000000000000090910467ffffffffffffffff16106132b1576040517f3e37e23300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080515f81526020810190915261117f90849084906149d0565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e6132f681613739565b67ffffffffffffffff84165f9081526083602052604090205463ffffffff161561334c576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff67ffffffffffffffff85161115613394576040517f4c753f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff87165f9081526082602052604090205463ffffffff16156133f6576040517fd409b93000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6134048888888887614ee3565b5f8080526002909101602052604090209390935550505050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59061344b81613739565b607e80545f919082906134639063ffffffff16615d5f565b91906101000a81548163ffffffff021916908363ffffffff160217905590506040518060c001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018767ffffffffffffffff1681526020018660ff1681526020015f1515815260200185815250607f5f8363ffffffff1663ffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101601c6101000a81548160ff021916908360ff160217905550608082015181600101601d6101000a81548160ff02191690831515021790555060a082015181600201559050508063ffffffff167fa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b5289898989898960405161364e96959493929190615fb6565b60405180910390a25050505050505050565b60855467ffffffffffffffff8281165f90815260048501602052604081205490924292613691929181169116615d28565b67ffffffffffffffff1611159392505050565b60068101545f90700100000000000000000000000000000000900467ffffffffffffffff16156137165750600681015467ffffffffffffffff70010000000000000000000000000000000090910481165f90815260049092016020526040909120546801000000000000000090041690565b506006015468010000000000000000900467ffffffffffffffff1690565b919050565b6130f68133615162565b60078801545f9067ffffffffffffffff9081169087161015613791576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88161561387457600689015467ffffffffffffffff70010000000000000000000000000000000090910481169089161115613801576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f90815260048a016020526040902060028101548154909288811668010000000000000000909204161461386e576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5061391f565b5067ffffffffffffffff85165f908152600289016020526040902054806138c7576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff680100000000000000009091048116908716111561391f576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff7001000000000000000000000000000000009091048116908816118061396757508767ffffffffffffffff168767ffffffffffffffff1611155b806139a05750600689015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690881611155b156139d7576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f90815260048b016020526040902054680100000000000000009004811690861614613a3b576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613a4a8a888888868961459c565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051613a7e919061602b565b602060405180830381855afa158015613a99573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613abc9190615ec5565b613ac69190615f5c565b60018c01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a91613b2e91889190600401616041565b602060405180830381865afa158015613b49573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b6d919061607d565b613ba3576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f90815260048c016020526040902060020154859003613bfa576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b5f5f613c128a6136a4565b60078b015490915067ffffffffffffffff9081169089161015613c61576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff891615613d465760068a015467ffffffffffffffff7001000000000000000000000000000000009091048116908a161115613cd1576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff808a165f90815260048c01602052604090206002810154815490945090918a8116680100000000000000009092041614613d40576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613de8565b67ffffffffffffffff88165f90815260028b016020526040902054915081613d9a576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168867ffffffffffffffff161115613de8576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff1611613e35576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613e448b8a8a8a878b61459c565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051613e78919061602b565b602060405180830381855afa158015613e93573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613eb69190615ec5565b613ec09190615f5c565b60018d01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a91613f2891899190600401616041565b602060405180830381865afa158015613f43573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f67919061607d565b613f9d576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613fa8848b615e67565b9050614009878267ffffffffffffffff16613fc161262b565b613fcb9190615d48565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691906151ca565b80608460088282829054906101000a900467ffffffffffffffff1661402e9190615d28565b82546101009290920a67ffffffffffffffff818102199093169183160217909155608480547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff16700100000000000000000000000000000000428416021790558e546040517f32c2d153000000000000000000000000000000000000000000000000000000008152918d166004830152602482018b905233604483015273ffffffffffffffffffffffffffffffffffffffff1691506332c2d153906064015f604051808303815f87803b158015614103575f5ffd5b505af1158015614115573d5f5f3e3d5ffd5b5050505050505050505050505050505050565b600682015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690821611158061418a5750600682015467ffffffffffffffff7001000000000000000000000000000000009091048116908216115b156141c1576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f818152600485016020908152604080832080546006890180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000092839004909816918202979097178755600280830154828752908a01909452919093209190915560018201546005870155835477ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000909302929092179092557f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d6142c8612b42565b6040518263ffffffff1660e01b81526004016142e691815260200190565b5f604051808303815f87803b1580156142fd575f5ffd5b505af115801561430f573d5f5f3e3d5ffd5b5050855473ffffffffffffffffffffffffffffffffffffffff165f908152608260209081526040918290205460028701546001880154845167ffffffffffffffff898116825294810192909252818501529188166060830152915163ffffffff90921693507f581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b925081900360800190a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561440a575f5ffd5b505af115801561441c573d5f5f3e3d5ffd5b50505050611385615257565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611404575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1615611404575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b67ffffffffffffffff8086165f81815260038901602052604080822054938816825290205460609291158015906145d1575081155b15614608576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8061463f576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614648846152e9565b61467e576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3385838a8c5f0160149054906101000a900467ffffffffffffffff168d60010160149054906101000a900467ffffffffffffffff1689878d8f6040516020016146d09a9998979695949392919061609c565b604051602081830303815290604052925050509695505050505050565b5f6146f7836136a4565b9050815f806147068484615e67565b60855467ffffffffffffffff91821692505f91614730916801000000000000000090041642615f6f565b90505b8467ffffffffffffffff168467ffffffffffffffff16146147bc5767ffffffffffffffff8085165f9081526003890160205260409020600181015490911682101561479a57600181015468010000000000000000900467ffffffffffffffff1694506147b6565b6147a48686615e67565b67ffffffffffffffff169350506147bc565b50614733565b5f6147c78484615f6f565b90508381101561481e57808403600c81116147e257806147e5565b600c5b9050806103e80a81608560109054906101000a900461ffff1661ffff160a608654028161481457614814615edc565b046086555061488d565b838103600c811161482f5780614832565b600c5b90505f816103e80a82608560109054906101000a900461ffff1661ffff160a670de0b6b3a7640000028161486857614868615edc565b04905080608654670de0b6b3a7640000028161488657614886615edc565b0460865550505b683635c9adc5dea0000060865411156148b257683635c9adc5dea000006086556148c9565b633b9aca0060865410156148c957633b9aca006086555b5050505050505050565b600681015467ffffffffffffffff7801000000000000000000000000000000000000000000000000820481167001000000000000000000000000000000009092041611156130f65760068101545f90614953907801000000000000000000000000000000000000000000000000900467ffffffffffffffff166001615d28565b905061495f8282613660565b156114045760068201545f90600290614997908490700100000000000000000000000000000000900467ffffffffffffffff16615e67565b6149a191906161a5565b6149ab9083615d28565b90506149b78382613660565b156149c65761117f8382614128565b61117f8383614128565b63ffffffff821615806149ee5750607e5463ffffffff908116908316115b15614a25576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f9081526082602052604081205463ffffffff1690819003614a8a576040517f74a086a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8181165f9081526081602052604090206007810154909185166801000000000000000090910467ffffffffffffffff1603614af6576040517f4f61d51900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff1615159003614b6e576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018101546007830154700100000000000000000000000000000000900460ff9081167c01000000000000000000000000000000000000000000000000000000009092041614614bea576040517fb541abe200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018181018054918401805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff000000000000000000000000000000000000000084168117825591547fffffffff00000000000000000000000000000000000000000000000000000000909316909117740100000000000000000000000000000000000000009283900467ffffffffffffffff9081169093021790556007830180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff1663ffffffff88166801000000000000000002179055600683015478010000000000000000000000000000000000000000000000008104821670010000000000000000000000000000000090910490911614614d39576040517f9d59507b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f614d4384610ca7565b6007840180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831617905582546040517f4f1ef28600000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff89811692634f1ef28692614dd392169089906004016161d4565b5f604051808303815f87803b158015614dea575f5ffd5b505af1158015614dfc573d5f5f3e3d5ffd5b50506040805163ffffffff8a8116825267ffffffffffffffff86166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff16614e91576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b608080545f9182918290614efc9063ffffffff16615d5f565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff16021790555060815f8263ffffffff1663ffffffff1681526020019081526020015f20915086825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550848260010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555085826001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083825f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550828260070160106101000a81548160ff021916908360ff1602179055508063ffffffff167fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850868987875f60405161515095949392919067ffffffffffffffff958616815273ffffffffffffffffffffffffffffffffffffffff949094166020850152918416604084015260ff166060830152909116608082015260a00190565b60405180910390a25095945050505050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611404576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261117f90849061536c565b606f5460ff1615615294576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b5f67ffffffff0000000167ffffffffffffffff8316108015615320575067ffffffff00000001604083901c67ffffffffffffffff16105b8015615341575067ffffffff00000001608083901c67ffffffffffffffff16105b8015615358575067ffffffff0000000160c083901c105b1561536557506001919050565b505f919050565b5f6153cd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661547c9092919063ffffffff16565b80519091501561117f57808060200190518101906153eb919061607d565b61117f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b606061548a84845f85615492565b949350505050565b606082471015615524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401615473565b5f5f8673ffffffffffffffffffffffffffffffffffffffff16858760405161554c919061602b565b5f6040518083038185875af1925050503d805f8114615586576040519150601f19603f3d011682016040523d82523d5f602084013e61558b565b606091505b5091509150611c6d878383876060831561562c5782515f036156255773ffffffffffffffffffffffffffffffffffffffff85163b615625576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401615473565b508161548a565b61548a83838151156156415781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016154739190615aed565b610a108061620383390190565b803563ffffffff81168114613734575f5ffd5b803567ffffffffffffffff81168114613734575f5ffd5b5f5f604083850312156156bd575f5ffd5b6156c683615682565b91506156d460208401615695565b90509250929050565b5f602082840312156156ed575f5ffd5b610c9e82615682565b806103008101831015610ca1575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b03121561571f575f5ffd5b61572889615682565b975061573660208a01615695565b965061574460408a01615695565b955061575260608a01615695565b945061576060808a01615695565b935060a0890135925060c0890135915061577d8a60e08b016156f6565b90509295985092959890939650565b73ffffffffffffffffffffffffffffffffffffffff811681146130f6575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b0312156157c5575f5ffd5b6157ce89615682565b97506157dc60208a01615695565b96506157ea60408a01615695565b95506157f860608a01615695565b94506080890135935060a0890135925060c08901356158168161578c565b915061577d8a60e08b016156f6565b5f60208284031215615835575f5ffd5b813561ffff81168114615846575f5ffd5b9392505050565b5f6020828403121561585d575f5ffd5b5035919050565b5f5f60408385031215615875575f5ffd5b8235915060208301356158878161578c565b809150509250929050565b5f602082840312156158a2575f5ffd5b610c9e82615695565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f67ffffffffffffffff8411156158f2576158f26158ab565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff8211171561593f5761593f6158ab565b604052838152905080828401851015615956575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f83011261597c575f5ffd5b610c9e838335602085016158d8565b5f5f5f5f5f5f5f60e0888a0312156159a1575f5ffd5b6159aa88615682565b96506159b860208901615695565b955060408801356159c88161578c565b945060608801356159d88161578c565b935060808801356159e88161578c565b925060a088013567ffffffffffffffff811115615a03575f5ffd5b615a0f8a828b0161596d565b92505060c088013567ffffffffffffffff811115615a2b575f5ffd5b615a378a828b0161596d565b91505092959891949750929550565b5f5f5f5f5f5f60c08789031215615a5b575f5ffd5b615a6487615682565b9550615a7260208801615695565b9450615a8060408801615695565b959894975094956060810135955060808101359460a0909101359350915050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610c9e6020830184615aa1565b5f5f60408385031215615b10575f5ffd5b82356156c68161578c565b5f5f60408385031215615b2c575f5ffd5b615b3583615695565b946020939093013593505050565b5f5f5f60608486031215615b55575f5ffd5b8335615b608161578c565b9250615b6e60208501615682565b9150604084013567ffffffffffffffff811115615b89575f5ffd5b8401601f81018613615b99575f5ffd5b615ba8868235602084016158d8565b9150509250925092565b5f60208284031215615bc2575f5ffd5b81356158468161578c565b5f5f60408385031215615bde575f5ffd5b8235615be98161578c565b91506156d460208401615682565b803560ff81168114613734575f5ffd5b5f5f5f5f5f5f60c08789031215615c1c575f5ffd5b8635615c278161578c565b95506020870135615c378161578c565b9450615c4560408801615695565b9350615c5360608801615695565b925060808701359150615c6860a08801615bf7565b90509295509295509295565b5f5f5f5f5f5f60c08789031215615c89575f5ffd5b8635615c948161578c565b95506020870135615ca48161578c565b9450615cb260408801615695565b9350615cc060608801615bf7565b92506080870135915060a087013567ffffffffffffffff811115615ce2575f5ffd5b615cee89828a0161596d565b9150509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8181168382160190811115610ca157610ca1615cfb565b8082028115828204841417610ca157610ca1615cfb565b5f63ffffffff821663ffffffff8103615d7a57615d7a615cfb565b60010192915050565b73ffffffffffffffffffffffffffffffffffffffff8416815273ffffffffffffffffffffffffffffffffffffffff83166020820152606060408201525f615dcd6060830184615aa1565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015263ffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015260c060808201525f615e4860c0830185615aa1565b82810360a0840152615e5a8185615aa1565b9998505050505050505050565b67ffffffffffffffff8281168282160390811115610ca157610ca1615cfb565b5f67ffffffffffffffff821667ffffffffffffffff8103615d7a57615d7a615cfb565b5f60208284031215615eba575f5ffd5b81516158468161578c565b5f60208284031215615ed5575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615f1757615f17615edc565b500490565b80820180821115610ca157610ca1615cfb565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82615f6a57615f6a615edc565b500690565b81810381811115610ca157610ca1615cfb565b5f81615f9057615f90615cfb565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015267ffffffffffffffff8516604082015260ff8416606082015282608082015260c060a08201525f61601f60c0830184615aa1565b98975050505050505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b6001811015616073578151835260209283019290910190600101616054565b5050509392505050565b5f6020828403121561608d575f5ffd5b81518015158114615846575f5ffd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008b60601b1681528960148201528860348201527fffffffffffffffff0000000000000000000000000000000000000000000000008860c01b1660548201527fffffffffffffffff0000000000000000000000000000000000000000000000008760c01b16605c8201527fffffffffffffffff0000000000000000000000000000000000000000000000008660c01b16606482015284606c82015283608c8201528260ac82015261619460cc82018360c01b7fffffffffffffffff000000000000000000000000000000000000000000000000169052565b60d4019a9950505050505050505050565b5f67ffffffffffffffff8316806161be576161be615edc565b8067ffffffffffffffff84160491505092915050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f61548a6040830184615aa156fe60a0604052604051610a10380380610a1083398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206109f05f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206109f05f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516105cb6104255f395f601001526105cb5ff3fe608060405261000c61000e565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633036100a7575f357fffffffff00000000000000000000000000000000000000000000000000000000167f4f1ef286000000000000000000000000000000000000000000000000000000001461009f5761009d6100ab565b565b61009d6100bb565b61009d5b61009d6100b66100e9565b61012d565b5f806100ca3660048184610410565b8101906100d79190610464565b915091506100e5828261014b565b5050565b5f6101287f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e808015610147573d5ff35b3d5ffd5b610154826101b2565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101aa576101a58282610285565b505050565b6100e5610304565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361021f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516102ae919061057f565b5f60405180830381855af49150503d805f81146102e6576040519150601f19603f3d011682016040523d82523d5f602084013e6102eb565b606091505b50915091506102fb85838361033c565b95945050505050565b341561009d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826103515761034c826103ce565b6103c7565b8151158015610375575073ffffffffffffffffffffffffffffffffffffffff84163b155b156103c4576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610216565b50805b9392505050565b8051156103de5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f8585111561041e575f5ffd5b8386111561042a575f5ffd5b5050820193919092039150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215610475575f5ffd5b823573ffffffffffffffffffffffffffffffffffffffff81168114610498575f5ffd5b9150602083013567ffffffffffffffff8111156104b3575f5ffd5b8301601f810185136104c3575f5ffd5b803567ffffffffffffffff8111156104dd576104dd610437565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561054957610549610437565b604052818152828201602001871015610560575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220d77806b34ab0065d93e4bcbbd2cd8eea87814e7eac6721c2b3b7f2dd0abe324964736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103a2646970667358221220560407977c80fec68763929e4026bacb0dc8996cb028ff9b5654381b430fb10964736f6c634300081c0033",
}

// PolygonrollupmanagerpreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagerpreviousMetaData.ABI instead.
var PolygonrollupmanagerpreviousABI = PolygonrollupmanagerpreviousMetaData.ABI

// PolygonrollupmanagerpreviousBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagerpreviousMetaData.Bin instead.
var PolygonrollupmanagerpreviousBin = PolygonrollupmanagerpreviousMetaData.Bin

// DeployPolygonrollupmanagerprevious deploys a new Ethereum contract, binding an instance of Polygonrollupmanagerprevious to it.
func DeployPolygonrollupmanagerprevious(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagerprevious, error) {
	parsed, err := PolygonrollupmanagerpreviousMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagerpreviousBin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanagerprevious{PolygonrollupmanagerpreviousCaller: PolygonrollupmanagerpreviousCaller{contract: contract}, PolygonrollupmanagerpreviousTransactor: PolygonrollupmanagerpreviousTransactor{contract: contract}, PolygonrollupmanagerpreviousFilterer: PolygonrollupmanagerpreviousFilterer{contract: contract}}, nil
}

// Polygonrollupmanagerprevious is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanagerprevious struct {
	PolygonrollupmanagerpreviousCaller     // Read-only binding to the contract
	PolygonrollupmanagerpreviousTransactor // Write-only binding to the contract
	PolygonrollupmanagerpreviousFilterer   // Log filterer for contract events
}

// PolygonrollupmanagerpreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupmanagerpreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerpreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupmanagerpreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerpreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupmanagerpreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagerpreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupmanagerpreviousSession struct {
	Contract     *Polygonrollupmanagerprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonrollupmanagerpreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupmanagerpreviousCallerSession struct {
	Contract *PolygonrollupmanagerpreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// PolygonrollupmanagerpreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupmanagerpreviousTransactorSession struct {
	Contract     *PolygonrollupmanagerpreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// PolygonrollupmanagerpreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupmanagerpreviousRaw struct {
	Contract *Polygonrollupmanagerprevious // Generic contract binding to access the raw methods on
}

// PolygonrollupmanagerpreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupmanagerpreviousCallerRaw struct {
	Contract *PolygonrollupmanagerpreviousCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupmanagerpreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupmanagerpreviousTransactorRaw struct {
	Contract *PolygonrollupmanagerpreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanagerprevious creates a new instance of Polygonrollupmanagerprevious, bound to a specific deployed contract.
func NewPolygonrollupmanagerprevious(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanagerprevious, error) {
	contract, err := bindPolygonrollupmanagerprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerprevious{PolygonrollupmanagerpreviousCaller: PolygonrollupmanagerpreviousCaller{contract: contract}, PolygonrollupmanagerpreviousTransactor: PolygonrollupmanagerpreviousTransactor{contract: contract}, PolygonrollupmanagerpreviousFilterer: PolygonrollupmanagerpreviousFilterer{contract: contract}}, nil
}

// NewPolygonrollupmanagerpreviousCaller creates a new read-only instance of Polygonrollupmanagerprevious, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupmanagerpreviousCaller, error) {
	contract, err := bindPolygonrollupmanagerprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousCaller{contract: contract}, nil
}

// NewPolygonrollupmanagerpreviousTransactor creates a new write-only instance of Polygonrollupmanagerprevious, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupmanagerpreviousTransactor, error) {
	contract, err := bindPolygonrollupmanagerprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousTransactor{contract: contract}, nil
}

// NewPolygonrollupmanagerpreviousFilterer creates a new log filterer instance of Polygonrollupmanagerprevious, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupmanagerpreviousFilterer, error) {
	contract, err := bindPolygonrollupmanagerprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousFilterer{contract: contract}, nil
}

// bindPolygonrollupmanagerprevious binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanagerprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupmanagerpreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagerprevious.Contract.PolygonrollupmanagerpreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.PolygonrollupmanagerpreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.PolygonrollupmanagerpreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagerprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagerprevious.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagerprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.BridgeAddress(&_Polygonrollupmanagerprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.BridgeAddress(&_Polygonrollupmanagerprevious.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagerprevious.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagerprevious.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.ChainIDToRollupID(&_Polygonrollupmanagerprevious.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.ChainIDToRollupID(&_Polygonrollupmanagerprevious.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.GetBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.GetBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.GetForcedBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerprevious.Contract.GetForcedBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetInputSnarkBytes(&_Polygonrollupmanagerprevious.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetInputSnarkBytes(&_Polygonrollupmanagerprevious.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagerprevious.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagerprevious.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRoleAdmin(&_Polygonrollupmanagerprevious.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRoleAdmin(&_Polygonrollupmanagerprevious.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupExitRoot(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupExitRoot(&_Polygonrollupmanagerprevious.CallOpts)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetRollupPendingStateTransitions(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getRollupPendingStateTransitions", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesPendingState)).(*LegacyZKEVMStateVariablesPendingState)

	return out0, err

}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagerprevious.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagerprevious.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.GlobalExitRootManager(&_Polygonrollupmanagerprevious.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.GlobalExitRootManager(&_Polygonrollupmanagerprevious.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.HasRole(&_Polygonrollupmanagerprevious.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.HasRole(&_Polygonrollupmanagerprevious.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.IsEmergencyState(&_Polygonrollupmanagerprevious.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.IsEmergencyState(&_Polygonrollupmanagerprevious.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) IsPendingStateConsolidable(opts *bind.CallOpts, rollupID uint32, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "isPendingStateConsolidable", rollupID, pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagerprevious.CallOpts, rollupID, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagerprevious.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagerprevious.CallOpts, rollupID, pendingStateNum)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.LastAggregationTimestamp(&_Polygonrollupmanagerprevious.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.LastAggregationTimestamp(&_Polygonrollupmanagerprevious.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagerprevious.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagerprevious.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagerprevious.Contract.MultiplierBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagerprevious.Contract.MultiplierBatchFee(&_Polygonrollupmanagerprevious.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.PendingStateTimeout(&_Polygonrollupmanagerprevious.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.PendingStateTimeout(&_Polygonrollupmanagerprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.Pol(&_Polygonrollupmanagerprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagerprevious.Contract.Pol(&_Polygonrollupmanagerprevious.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupAddressToID(&_Polygonrollupmanagerprevious.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupAddressToID(&_Polygonrollupmanagerprevious.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupCount(&_Polygonrollupmanagerprevious.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupCount(&_Polygonrollupmanagerprevious.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	outstruct := new(struct {
		RollupContract                 common.Address
		ChainID                        uint64
		Verifier                       common.Address
		ForkID                         uint64
		LastLocalExitRoot              [32]byte
		LastBatchSequenced             uint64
		LastVerifiedBatch              uint64
		LastPendingState               uint64
		LastPendingStateConsolidated   uint64
		LastVerifiedBatchBeforeUpgrade uint64
		RollupTypeID                   uint64
		RollupCompatibilityID          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RollupContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainID = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Verifier = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ForkID = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.LastLocalExitRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.LastBatchSequenced = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatch = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.LastPendingState = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.LastPendingStateConsolidated = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatchBeforeUpgrade = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.RollupTypeID = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.RollupCompatibilityID = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollupIDToRollupData(rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupIDToRollupData(&_Polygonrollupmanagerprevious.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) RollupIDToRollupData(rollupID uint32) (struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastPendingState               uint64
	LastPendingStateConsolidated   uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupCompatibilityID          uint8
}, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupIDToRollupData(&_Polygonrollupmanagerprevious.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupTypeCount(&_Polygonrollupmanagerprevious.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupTypeCount(&_Polygonrollupmanagerprevious.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

	outstruct := new(struct {
		ConsensusImplementation common.Address
		Verifier                common.Address
		ForkID                  uint64
		RollupCompatibilityID   uint8
		Obsolete                bool
		Genesis                 [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusImplementation = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ForkID = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RollupCompatibilityID = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Obsolete = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Genesis = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupTypeMap(&_Polygonrollupmanagerprevious.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagerprevious.Contract.RollupTypeMap(&_Polygonrollupmanagerprevious.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TotalSequencedBatches(&_Polygonrollupmanagerprevious.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TotalSequencedBatches(&_Polygonrollupmanagerprevious.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TotalVerifiedBatches(&_Polygonrollupmanagerprevious.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TotalVerifiedBatches(&_Polygonrollupmanagerprevious.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagerprevious.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagerprevious.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerprevious.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagerprevious.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagerprevious.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ActivateEmergencyState(&_Polygonrollupmanagerprevious.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ActivateEmergencyState(&_Polygonrollupmanagerprevious.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.AddExistingRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.AddExistingRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.AddNewRollupType(&_Polygonrollupmanagerprevious.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.AddNewRollupType(&_Polygonrollupmanagerprevious.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) ConsolidatePendingState(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "consolidatePendingState", rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ConsolidatePendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ConsolidatePendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.CreateNewRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.CreateNewRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.DeactivateEmergencyState(&_Polygonrollupmanagerprevious.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.DeactivateEmergencyState(&_Polygonrollupmanagerprevious.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.GrantRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.GrantRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ObsoleteRollupType(&_Polygonrollupmanagerprevious.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ObsoleteRollupType(&_Polygonrollupmanagerprevious.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.OnSequenceBatches(&_Polygonrollupmanagerprevious.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.OnSequenceBatches(&_Polygonrollupmanagerprevious.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) OverridePendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "overridePendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.OverridePendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.OverridePendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "proveNonDeterministicPendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RenounceRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RenounceRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RevokeRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RevokeRole(&_Polygonrollupmanagerprevious.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RollbackBatches(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.RollbackBatches(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetBatchFee(&_Polygonrollupmanagerprevious.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetBatchFee(&_Polygonrollupmanagerprevious.TransactOpts, newBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagerprevious.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagerprevious.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetPendingStateTimeout(&_Polygonrollupmanagerprevious.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetPendingStateTimeout(&_Polygonrollupmanagerprevious.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagerprevious.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagerprevious.TransactOpts, newTrustedAggregatorTimeout)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagerprevious.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagerprevious.TransactOpts, newVerifyBatchTimeTarget)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.UpdateRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.UpdateRollup(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagerprevious.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) VerifyBatches(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "verifyBatches", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatches(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatches(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerprevious.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagerprevious.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// PolygonrollupmanagerpreviousAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousAddExistingRollupIterator struct {
	Event *PolygonrollupmanagerpreviousAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousAddExistingRollup)
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
		it.Event = new(PolygonrollupmanagerpreviousAddExistingRollup)
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
func (it *PolygonrollupmanagerpreviousAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousAddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousAddExistingRollup struct {
	RollupID                       uint32
	ForkID                         uint64
	RollupAddress                  common.Address
	ChainID                        uint64
	RollupCompatibilityID          uint8
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterAddExistingRollup is a free log retrieval operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousAddExistingRollupIterator{contract: _Polygonrollupmanagerprevious.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousAddExistingRollup)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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

// ParseAddExistingRollup is a log parse operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseAddExistingRollup(log types.Log) (*PolygonrollupmanagerpreviousAddExistingRollup, error) {
	event := new(PolygonrollupmanagerpreviousAddExistingRollup)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousAddNewRollupTypeIterator struct {
	Event *PolygonrollupmanagerpreviousAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousAddNewRollupType)
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
		it.Event = new(PolygonrollupmanagerpreviousAddNewRollupType)
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
func (it *PolygonrollupmanagerpreviousAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousAddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousAddNewRollupType struct {
	RollupTypeID            uint32
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Genesis                 [32]byte
	Description             string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAddNewRollupType is a free log retrieval operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagerpreviousAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousAddNewRollupTypeIterator{contract: _Polygonrollupmanagerprevious.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousAddNewRollupType)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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

// ParseAddNewRollupType is a log parse operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagerpreviousAddNewRollupType, error) {
	event := new(PolygonrollupmanagerpreviousAddNewRollupType)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousConsolidatePendingStateIterator struct {
	Event *PolygonrollupmanagerpreviousConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousConsolidatePendingState)
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
		it.Event = new(PolygonrollupmanagerpreviousConsolidatePendingState)
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
func (it *PolygonrollupmanagerpreviousConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousConsolidatePendingState struct {
	RollupID        uint32
	NumBatch        uint64
	StateRoot       [32]byte
	ExitRoot        [32]byte
	PendingStateNum uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterConsolidatePendingState is a free log retrieval operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousConsolidatePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousConsolidatePendingStateIterator{contract: _Polygonrollupmanagerprevious.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousConsolidatePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousConsolidatePendingState)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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

// ParseConsolidatePendingState is a log parse operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonrollupmanagerpreviousConsolidatePendingState, error) {
	event := new(PolygonrollupmanagerpreviousConsolidatePendingState)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousCreateNewRollupIterator struct {
	Event *PolygonrollupmanagerpreviousCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousCreateNewRollup)
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
		it.Event = new(PolygonrollupmanagerpreviousCreateNewRollup)
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
func (it *PolygonrollupmanagerpreviousCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousCreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousCreateNewRollup struct {
	RollupID        uint32
	RollupTypeID    uint32
	RollupAddress   common.Address
	ChainID         uint64
	GasTokenAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateNewRollup is a free log retrieval operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousCreateNewRollupIterator{contract: _Polygonrollupmanagerprevious.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousCreateNewRollup)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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

// ParseCreateNewRollup is a log parse operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseCreateNewRollup(log types.Log) (*PolygonrollupmanagerpreviousCreateNewRollup, error) {
	event := new(PolygonrollupmanagerpreviousCreateNewRollup)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousEmergencyStateActivatedIterator struct {
	Event *PolygonrollupmanagerpreviousEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousEmergencyStateActivated)
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
		it.Event = new(PolygonrollupmanagerpreviousEmergencyStateActivated)
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
func (it *PolygonrollupmanagerpreviousEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousEmergencyStateActivatedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousEmergencyStateActivated)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonrollupmanagerpreviousEmergencyStateActivated, error) {
	event := new(PolygonrollupmanagerpreviousEmergencyStateActivated)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator struct {
	Event *PolygonrollupmanagerpreviousEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousEmergencyStateDeactivated)
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
		it.Event = new(PolygonrollupmanagerpreviousEmergencyStateDeactivated)
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
func (it *PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousEmergencyStateDeactivatedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousEmergencyStateDeactivated)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonrollupmanagerpreviousEmergencyStateDeactivated, error) {
	event := new(PolygonrollupmanagerpreviousEmergencyStateDeactivated)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousInitializedIterator struct {
	Event *PolygonrollupmanagerpreviousInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousInitialized)
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
		it.Event = new(PolygonrollupmanagerpreviousInitialized)
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
func (it *PolygonrollupmanagerpreviousInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousInitialized represents a Initialized event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousInitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousInitializedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousInitialized)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseInitialized(log types.Log) (*PolygonrollupmanagerpreviousInitialized, error) {
	event := new(PolygonrollupmanagerpreviousInitialized)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousObsoleteRollupTypeIterator struct {
	Event *PolygonrollupmanagerpreviousObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousObsoleteRollupType)
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
		it.Event = new(PolygonrollupmanagerpreviousObsoleteRollupType)
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
func (it *PolygonrollupmanagerpreviousObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagerpreviousObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousObsoleteRollupTypeIterator{contract: _Polygonrollupmanagerprevious.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousObsoleteRollupType)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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

// ParseObsoleteRollupType is a log parse operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseObsoleteRollupType(log types.Log) (*PolygonrollupmanagerpreviousObsoleteRollupType, error) {
	event := new(PolygonrollupmanagerpreviousObsoleteRollupType)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousOnSequenceBatchesIterator struct {
	Event *PolygonrollupmanagerpreviousOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousOnSequenceBatches)
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
		it.Event = new(PolygonrollupmanagerpreviousOnSequenceBatches)
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
func (it *PolygonrollupmanagerpreviousOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousOnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousOnSequenceBatchesIterator{contract: _Polygonrollupmanagerprevious.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousOnSequenceBatches)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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

// ParseOnSequenceBatches is a log parse operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseOnSequenceBatches(log types.Log) (*PolygonrollupmanagerpreviousOnSequenceBatches, error) {
	event := new(PolygonrollupmanagerpreviousOnSequenceBatches)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousOverridePendingStateIterator struct {
	Event *PolygonrollupmanagerpreviousOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousOverridePendingState)
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
		it.Event = new(PolygonrollupmanagerpreviousOverridePendingState)
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
func (it *PolygonrollupmanagerpreviousOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousOverridePendingState represents a OverridePendingState event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousOverridePendingState struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOverridePendingState is a free log retrieval operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterOverridePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousOverridePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousOverridePendingStateIterator{contract: _Polygonrollupmanagerprevious.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousOverridePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousOverridePendingState)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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

// ParseOverridePendingState is a log parse operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseOverridePendingState(log types.Log) (*PolygonrollupmanagerpreviousOverridePendingState, error) {
	event := new(PolygonrollupmanagerpreviousOverridePendingState)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator struct {
	Event *PolygonrollupmanagerpreviousProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousProveNonDeterministicPendingState)
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
		it.Event = new(PolygonrollupmanagerpreviousProveNonDeterministicPendingState)
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
func (it *PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousProveNonDeterministicPendingStateIterator{contract: _Polygonrollupmanagerprevious.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousProveNonDeterministicPendingState)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonrollupmanagerpreviousProveNonDeterministicPendingState, error) {
	event := new(PolygonrollupmanagerpreviousProveNonDeterministicPendingState)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleAdminChangedIterator struct {
	Event *PolygonrollupmanagerpreviousRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousRoleAdminChanged)
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
		it.Event = new(PolygonrollupmanagerpreviousRoleAdminChanged)
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
func (it *PolygonrollupmanagerpreviousRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonrollupmanagerpreviousRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousRoleAdminChangedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousRoleAdminChanged)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonrollupmanagerpreviousRoleAdminChanged, error) {
	event := new(PolygonrollupmanagerpreviousRoleAdminChanged)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleGrantedIterator struct {
	Event *PolygonrollupmanagerpreviousRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousRoleGranted)
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
		it.Event = new(PolygonrollupmanagerpreviousRoleGranted)
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
func (it *PolygonrollupmanagerpreviousRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousRoleGranted represents a RoleGranted event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagerpreviousRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousRoleGrantedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousRoleGranted)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseRoleGranted(log types.Log) (*PolygonrollupmanagerpreviousRoleGranted, error) {
	event := new(PolygonrollupmanagerpreviousRoleGranted)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleRevokedIterator struct {
	Event *PolygonrollupmanagerpreviousRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousRoleRevoked)
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
		it.Event = new(PolygonrollupmanagerpreviousRoleRevoked)
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
func (it *PolygonrollupmanagerpreviousRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousRoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagerpreviousRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousRoleRevokedIterator{contract: _Polygonrollupmanagerprevious.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousRoleRevoked)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseRoleRevoked(log types.Log) (*PolygonrollupmanagerpreviousRoleRevoked, error) {
	event := new(PolygonrollupmanagerpreviousRoleRevoked)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRollbackBatchesIterator struct {
	Event *PolygonrollupmanagerpreviousRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousRollbackBatches)
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
		it.Event = new(PolygonrollupmanagerpreviousRollbackBatches)
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
func (it *PolygonrollupmanagerpreviousRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousRollbackBatches represents a RollbackBatches event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*PolygonrollupmanagerpreviousRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousRollbackBatchesIterator{contract: _Polygonrollupmanagerprevious.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousRollbackBatches)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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

// ParseRollbackBatches is a log parse operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseRollbackBatches(log types.Log) (*PolygonrollupmanagerpreviousRollbackBatches, error) {
	event := new(PolygonrollupmanagerpreviousRollbackBatches)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetBatchFeeIterator struct {
	Event *PolygonrollupmanagerpreviousSetBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetBatchFee)
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
		it.Event = new(PolygonrollupmanagerpreviousSetBatchFee)
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
func (it *PolygonrollupmanagerpreviousSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetBatchFeeIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetBatchFee)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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

// ParseSetBatchFee is a log parse operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetBatchFee(log types.Log) (*PolygonrollupmanagerpreviousSetBatchFee, error) {
	event := new(PolygonrollupmanagerpreviousSetBatchFee)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator struct {
	Event *PolygonrollupmanagerpreviousSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetMultiplierBatchFee)
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
		it.Event = new(PolygonrollupmanagerpreviousSetMultiplierBatchFee)
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
func (it *PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetMultiplierBatchFeeIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetMultiplierBatchFee)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonrollupmanagerpreviousSetMultiplierBatchFee, error) {
	event := new(PolygonrollupmanagerpreviousSetMultiplierBatchFee)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator struct {
	Event *PolygonrollupmanagerpreviousSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetPendingStateTimeout)
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
		it.Event = new(PolygonrollupmanagerpreviousSetPendingStateTimeout)
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
func (it *PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetPendingStateTimeoutIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetPendingStateTimeout)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonrollupmanagerpreviousSetPendingStateTimeout, error) {
	event := new(PolygonrollupmanagerpreviousSetPendingStateTimeout)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagerpreviousSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagerpreviousSetTrustedAggregator)
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
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetTrustedAggregatorIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetTrustedAggregator)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonrollupmanagerpreviousSetTrustedAggregator, error) {
	event := new(PolygonrollupmanagerpreviousSetTrustedAggregator)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout)
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
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetTrustedAggregatorTimeoutIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout, error) {
	event := new(PolygonrollupmanagerpreviousSetTrustedAggregatorTimeout)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget)
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
func (it *PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousSetVerifyBatchTimeTargetIterator{contract: _Polygonrollupmanagerprevious.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget, error) {
	event := new(PolygonrollupmanagerpreviousSetVerifyBatchTimeTarget)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousUpdateRollupIterator struct {
	Event *PolygonrollupmanagerpreviousUpdateRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousUpdateRollup)
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
		it.Event = new(PolygonrollupmanagerpreviousUpdateRollup)
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
func (it *PolygonrollupmanagerpreviousUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousUpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagerpreviousUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousUpdateRollupIterator{contract: _Polygonrollupmanagerprevious.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousUpdateRollup)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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

// ParseUpdateRollup is a log parse operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseUpdateRollup(log types.Log) (*PolygonrollupmanagerpreviousUpdateRollup, error) {
	event := new(PolygonrollupmanagerpreviousUpdateRollup)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousVerifyBatchesIterator struct {
	Event *PolygonrollupmanagerpreviousVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousVerifyBatches)
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
		it.Event = new(PolygonrollupmanagerpreviousVerifyBatches)
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
func (it *PolygonrollupmanagerpreviousVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousVerifyBatches represents a VerifyBatches event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousVerifyBatches struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterVerifyBatches(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagerpreviousVerifyBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousVerifyBatchesIterator{contract: _Polygonrollupmanagerprevious.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousVerifyBatches, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousVerifyBatches)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupmanagerpreviousVerifyBatches, error) {
	event := new(PolygonrollupmanagerpreviousVerifyBatches)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator)
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
func (it *PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanagerprevious contract.
type PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator struct {
	RollupID   uint32
	NumBatch   uint64
	StateRoot  [32]byte
	ExitRoot   [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatchesTrustedAggregator is a free log retrieval operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanagerprevious.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerprevious.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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

// ParseVerifyBatchesTrustedAggregator is a log parse operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerprevious *PolygonrollupmanagerpreviousFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonrollupmanagerpreviousVerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanagerprevious.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

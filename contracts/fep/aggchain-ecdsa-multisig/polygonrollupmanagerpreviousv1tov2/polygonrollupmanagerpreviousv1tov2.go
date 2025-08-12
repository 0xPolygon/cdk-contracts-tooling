// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanagerpreviousv1tov2

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

// Polygonrollupmanagerpreviousv1tov2MetaData contains all meta data concerning the Polygonrollupmanagerpreviousv1tov2 contract.
var Polygonrollupmanagerpreviousv1tov2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupPendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structLegacyZKEVMStateVariables.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"},{\"internalType\":\"contractPolygonZkEVMExistentEtrog\",\"name\":\"polygonZkEVM\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"zkEVMVerifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMForkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMChainID\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051616fdf380380616fdf83398101604081905261002e9161012b565b6001600160a01b0380841660805282811660c052811660a05261004f610057565b505050610175565b5f54610100900460ff16156100c25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610112575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610128575f5ffd5b50565b5f5f5f6060848603121561013d575f5ffd5b835161014881610114565b602085015190935061015981610114565b604085015190925061016a81610114565b809150509250925092565b60805160a05160c051616e086101d75f395f8181610ab7015281816127d6015261469b01525f8181610857015281816135e70152614a6001525f8181610a1b015281816114580152818161166e015281816124d3015261493f0152616e085ff3fe608060405234801561000f575f5ffd5b5060043610610346575f3560e01c8063841b24d7116101be578063c1acbc34116100fe578063dbc169761161009e578063e46761c411610079578063e46761c414610ab2578063f34eb8eb14610ad9578063f4e9267514610aec578063f9c4c2ae14610afc575f5ffd5b8063dbc1697614610a77578063dde0ff7714610a7f578063e0bfd3d214610a9f575f5ffd5b8063d02103ca116100d9578063d02103ca14610a16578063d5073f6f14610a3d578063d547741f14610a50578063d939b31514610a63575f5ffd5b8063c1acbc34146109b6578063c4c928c2146109de578063ceee281d146109f1575f5ffd5b80639c9f3dfe11610169578063a2967d9911610144578063a2967d991461084a578063a3c573eb14610852578063afd23cbe1461089e578063b99d0ad7146108d3575f5ffd5b80639c9f3dfe1461081d578063a066215c14610830578063a217fddf14610843575f5ffd5b806391d148541161019957806391d14854146107bd57806399f5634e146108025780639a908e731461080a575f5ffd5b8063841b24d71461076757806387c20c01146107975780638bd4f071146107aa575f5ffd5b8063252801691161028957806355a71ee0116102345780637222020f1161020f5780637222020f146106fc578063727885e91461070f5780637975fcfe146107225780637fb6e76a14610742575f5ffd5b806355a71ee0146105a957806360469169146105ea57806365c0504d146105f2575f5ffd5b806336568abe1161026457806336568abe1461057b578063394218e91461058e578063477fa270146105a1575f5ffd5b8063252801691461049d5780632f2ff15d1461055457806330c27dde14610567575f5ffd5b80631489ed10116102f45780631796a1ae116102cf5780631796a1ae1461042d5780631816b7e5146104525780632072f6c514610465578063248a9ca31461046d575f5ffd5b80631489ed10146103fa57806315064c961461040d5780631608859c1461041a575f5ffd5b80630a0d9fbe116103245780630a0d9fbe146103b457806311f6b287146103d457806312b86e19146103e7575f5ffd5b80630645af091461034a578063066ec0121461035f578063080b311114610391575b5f5ffd5b61035d61035836600461566a565b610c62565b005b6084546103739067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6103a461039f36600461573d565b61132d565b6040519015158152602001610388565b6085546103739068010000000000000000900467ffffffffffffffff1681565b6103736103e236600461576e565b611354565b61035d6103f5366004615798565b611371565b61035d61040836600461581d565b611587565b606f546103a49060ff1681565b61035d61042836600461573d565b61175c565b607e5461043d9063ffffffff1681565b60405163ffffffff9091168152602001610388565b61035d610460366004615895565b611831565b61035d61192a565b61048f61047b3660046158bd565b5f9081526034602052604090206001015490565b604051908152602001610388565b6105206104ab36600461573d565b60408051606080820183525f808352602080840182905292840181905263ffffffff9590951685526081825282852067ffffffffffffffff9485168652600301825293829020825194850183528054855260010154808416918501919091526801000000000000000090049091169082015290565b604080518251815260208084015167ffffffffffffffff908116918301919091529282015190921690820152606001610388565b61035d6105623660046158d4565b611a34565b6087546103739067ffffffffffffffff1681565b61035d6105893660046158d4565b611a58565b61035d61059c366004615902565b611ab5565b60865461048f565b61048f6105b736600461573d565b63ffffffff82165f90815260816020908152604080832067ffffffffffffffff8516845260020190915290205492915050565b61048f611bcb565b6106a561060036600461576e565b607f6020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff918216929182169167ffffffffffffffff740100000000000000000000000000000000000000008204169160ff7c010000000000000000000000000000000000000000000000000000000083048116927d0100000000000000000000000000000000000000000000000000000000009004169086565b6040805173ffffffffffffffffffffffffffffffffffffffff978816815296909516602087015267ffffffffffffffff9093169385019390935260ff166060840152901515608083015260a082015260c001610388565b61035d61070a36600461576e565b611be0565b61035d61071d366004615a10565b611d56565b610735610730366004615acb565b6122ad565b6040516103889190615b72565b61043d610750366004615902565b60836020525f908152604090205463ffffffff1681565b608454610373907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b61035d6107a536600461581d565b6122dd565b61035d6107b8366004615798565b6126c3565b6103a46107cb3660046158d4565b5f91825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61048f61278f565b610373610818366004615b84565b612895565b61035d61082b366004615902565b612adc565b61035d61083e366004615902565b612bc0565b61048f5f81565b61048f612ca6565b6108797f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610388565b6085546108c090700100000000000000000000000000000000900461ffff1681565b60405161ffff9091168152602001610388565b61096b6108e136600461573d565b60408051608080820183525f8083526020808401829052838501829052606093840182905263ffffffff9690961681526081865283812067ffffffffffffffff958616825260040186528390208351918201845280548086168352680100000000000000009004909416948101949094526001830154918401919091526002909101549082015290565b60405161038891905f60808201905067ffffffffffffffff835116825267ffffffffffffffff6020840151166020830152604083015160408301526060830151606083015292915050565b60845461037390700100000000000000000000000000000000900467ffffffffffffffff1681565b61035d6109ec366004615bac565b61304a565b61043d6109ff366004615c38565b60826020525f908152604090205463ffffffff1681565b6108797f000000000000000000000000000000000000000000000000000000000000000081565b61035d610a4b3660046158bd565b61348a565b61035d610a5e3660046158d4565b61353a565b6085546103739067ffffffffffffffff1681565b61035d61355e565b6084546103739068010000000000000000900467ffffffffffffffff1681565b61035d610aad366004615c63565b613662565b6108797f000000000000000000000000000000000000000000000000000000000000000081565b61035d610ae7366004615cd0565b613770565b60805461043d9063ffffffff1681565b610bd5610b0a36600461576e565b60816020525f90815260409020805460018201546005830154600684015460079094015473ffffffffffffffffffffffffffffffffffffffff80851695740100000000000000000000000000000000000000009586900467ffffffffffffffff908116969286169592909204821693928282169268010000000000000000808404821693700100000000000000000000000000000000808204841694780100000000000000000000000000000000000000000000000090920484169380831693830416910460ff168c565b6040805173ffffffffffffffffffffffffffffffffffffffff9d8e16815267ffffffffffffffff9c8d1660208201529c909a16998c019990995296891660608b015260808a019590955292871660a089015290861660c0880152851660e0870152841661010086015283166101208501529190911661014083015260ff1661016082015261018001610388565b5f54600290610100900460ff16158015610c8257505f5460ff8083169116105b610d13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f805461010060ff84167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090921691909117179055608580546084805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8e8116919091029190911790915567016345785d8a00006086558c167fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116176907080000000000000000177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff167103ea00000000000000000000000000000000179055610e1a6139af565b610e447f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd48c613a45565b610e4e5f88613a45565b610e787fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59088613a45565b610ea27f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e88613a45565b610ecc7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac88613a45565b610ef67fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd89613a45565b610f207fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd0889613a45565b610f4a7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f489613a45565b610f747fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db189613a45565b610fbe7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd47f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0613a4f565b610fe87f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f089613a45565b6110127f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb89613a45565b61105c7f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e7f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff285951613a4f565b6110867f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e87613a45565b6110b07f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff28595187613a45565b60735460745467ffffffffffffffff6801000000000000000090920482169116808214611109576040517f5c998a8600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61112d888888885f60745f9054906101000a900467ffffffffffffffff16613a99565b67ffffffffffffffff8381165f81815260756020908152604080832054600287018352818420558885168084526072808452828520600389018552948390208554815560018087018054919092018054918a167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008084168217835593547fffffffffffffffffffffffffffffffff000000000000000000000000000000009384169091176801000000000000000091829004909b1681029a909a17905560068a01805490911690931797870297909717909155600787018054909616909417909455607a54606f5493909152905492517f5d6717a500000000000000000000000000000000000000000000000000000000815293945073ffffffffffffffffffffffffffffffffffffffff8c811694635d6717a59461128c94938316936b01000000000000000000000090049092169160769160779190600401615e2c565b5f604051808303815f87803b1580156112a3575f5ffd5b505af11580156112b5573d5f5f3e3d5ffd5b50505f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055505060405160ff851681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498935060200191506113189050565b60405180910390a15050505050505050505050565b63ffffffff82165f90815260816020526040812061134b9083613d19565b90505b92915050565b63ffffffff81165f90815260816020526040812061134e90613d5d565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd461139b81613df2565b63ffffffff89165f9081526081602052604090206113bf818a8a8a8a8a8a8a613dfc565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f908152600284016020526040902086905560058301879055905470010000000000000000000000000000000090041615611456576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d61149a612ca6565b6040518263ffffffff1660e01b81526004016114b891815260200190565b5f604051808303815f87803b1580156114cf575f5ffd5b505af11580156114e1573d5f5f3e3d5ffd5b50506084805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a8000000000000000000000000000000000000000000000000017905550506040805167ffffffffffffffff881681526020810186905290810186905233606082015263ffffffff8b16907f3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e9060800160405180910390a250505050505050505050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd46115b181613df2565b63ffffffff89165f9081526081602052604090206115d5818a8a8a8a8a8a8a6142c0565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8a81169182029290921783555f90815260028401602052604090208790556005830188905590547001000000000000000000000000000000009004161561166c576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d6116b0612ca6565b6040518263ffffffff1660e01b81526004016116ce91815260200190565b5f604051808303815f87803b1580156116e5575f5ffd5b505af11580156116f7573d5f5f3e3d5ffd5b50506040805167ffffffffffffffff8b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b63ffffffff82165f9081526081602090815260408083203384527fc17b14a573f65366cdad721c7c0a0f76536bb4a86b935cdac44610e4f010b52a9092529091205460ff1661182257606f5460ff16156117e2576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117ec8183613d19565b611822576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61182c81836147e1565b505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db161185b81613df2565b6103e88261ffff16108061187457506103ff8261ffff16115b156118ab576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000061ffff8516908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a15050565b335f9081527f8875b94af5657a2903def9906d67a3f42d8a836d24b5602c00f00fc855339fcd602052604090205460ff16611a2a57608454700100000000000000000000000000000000900467ffffffffffffffff1615806119c4575060845442906119b89062093a8090700100000000000000000000000000000000900467ffffffffffffffff16615ec8565b67ffffffffffffffff16115b806119f3575060875442906119e79062093a809067ffffffffffffffff16615ec8565b67ffffffffffffffff16115b15611a2a576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a32614a5e565b565b5f82815260346020526040902060010154611a4e81613df2565b61182c8383614ae1565b73ffffffffffffffffffffffffffffffffffffffff81163314611aa7576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ab18282614b9c565b5050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1611adf81613df2565b606f5460ff16611b4e5760845467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690831610611b4e576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6084805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a19060200161191e565b5f6086546064611bdb9190615ee8565b905090565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd611c0a81613df2565b63ffffffff82161580611c285750607e5463ffffffff908116908316115b15611c5f576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff1615159003611cd7576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167d01000000000000000000000000000000000000000000000000000000000017905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08611d8081613df2565b63ffffffff88161580611d9e5750607e5463ffffffff908116908916115b15611dd5576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff1615159003611e4d576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88165f9081526083602052604090205463ffffffff1615611ea3576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608080545f91908290611ebb9063ffffffff16615eff565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f808252602082019283905293945073ffffffffffffffffffffffffffffffffffffffff909216913091611f1390615625565b611f1f93929190615f23565b604051809103905ff080158015611f38573d5f5f3e3d5ffd5b5090508160835f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508160825f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8463ffffffff1663ffffffff1681526020019081526020015f20905081815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360010160149054906101000a900467ffffffffffffffff168160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a815f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508360020154816002015f5f67ffffffffffffffff1681526020019081526020015f20819055508b63ffffffff168160070160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083600101601c9054906101000a900460ff168160070160106101000a81548160ff021916908360ff1602179055508263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418d848e8c60405161220e949392919063ffffffff94909416845273ffffffffffffffffffffffffffffffffffffffff928316602085015267ffffffffffffffff91909116604084015216606082015260800190565b60405180910390a26040517f7125702200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831690637125702290612272908d908d9088908e908e908e90600401615f76565b5f604051808303815f87803b158015612289575f5ffd5b505af115801561229b573d5f5f3e3d5ffd5b50505050505050505050505050505050565b63ffffffff86165f9081526081602052604090206060906122d2908787878787614c55565b979650505050505050565b606f5460ff161561231a576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f90815260816020908152604080832060845467ffffffffffffffff8a81168652600383019094529190932060010154429261237c927801000000000000000000000000000000000000000000000000900481169116615ec8565b67ffffffffffffffff1611156123be576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e86123cb8888616007565b67ffffffffffffffff16111561240d576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61241d81898989898989896142c0565b6124278187614da6565b60855467ffffffffffffffff165f03612565576006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f9081526002840160205260409020869055600583018790559054700100000000000000000000000000000000900416156124d1576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d612515612ca6565b6040518263ffffffff1660e01b815260040161253391815260200190565b5f604051808303815f87803b15801561254a575f5ffd5b505af115801561255c573d5f5f3e3d5ffd5b50505050612663565b61256e81614f8c565b600681018054700100000000000000000000000000000000900467ffffffffffffffff1690601061259e83616027565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815289831660208083019182528284018b8152606084018b81526006890154700100000000000000000000000000000000900487165f90815260048a0190935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b6040805167ffffffffffffffff8816815260208101869052908101869052339063ffffffff8b16907faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b49060600160405180910390a3505050505050505050565b606f5460ff1615612700576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f9081526081602052604090206127248189898989898989613dfc565b67ffffffffffffffff87165f9081526004820160209081526040918290206002015482519081529081018590527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a1612784614a5e565b505050505050505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561281b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061283f919061604a565b6084549091505f906128699067ffffffffffffffff68010000000000000000820481169116616007565b67ffffffffffffffff169050805f03612884575f9250505090565b61288e818361608e565b9250505090565b606f545f9060ff16156128d4576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f9081526082602052604081205463ffffffff1690819003612923576040517f71653c1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8367ffffffffffffffff165f03612966576040517f2590ccf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff81165f908152608160205260408120608480549192879261299790849067ffffffffffffffff16615ec8565b82546101009290920a67ffffffffffffffff81810219909316918316021790915560068301541690505f6129cb8783615ec8565b60068401805467ffffffffffffffff8084167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217909255604080516060810182528a815242841660208083019182528886168385019081525f95865260038b019091529290932090518155915160019290920180549151841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090921692909316919091171790559050612a8d83614f8c565b60405167ffffffffffffffff8216815263ffffffff8516907f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a259060200160405180910390a29695505050505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1612b0681613df2565b606f5460ff16612b595760855467ffffffffffffffff90811690831610612b59576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff84169081179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c759060200161191e565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1612bea81613df2565b620151808267ffffffffffffffff161115612c31576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c289060200161191e565b6080545f9063ffffffff16808203612cbf57505f919050565b5f8167ffffffffffffffff811115612cd957612cd961591b565b604051908082528060200260200182016040528015612d02578160200160208202803683370190505b5090505f5b82811015612d5f5760815f612d1d8360016160a1565b63ffffffff1663ffffffff1681526020019081526020015f2060050154828281518110612d4c57612d4c6160b4565b6020908102919091010152600101612d07565b505f60205b83600114612f7b575f612d786002866160e1565b612d8360028761608e565b612d8d91906160a1565b90505f8167ffffffffffffffff811115612da957612da961591b565b604051908082528060200260200182016040528015612dd2578160200160208202803683370190505b5090505f5b82811015612f2b57612dea6001846160f4565b81148015612e025750612dfe6002886160e1565b6001145b15612e7f5785612e13826002615ee8565b81518110612e2357612e236160b4565b602002602001015185604051602001612e46929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110612e6e57612e6e6160b4565b602002602001018181525050612f23565b85612e8b826002615ee8565b81518110612e9b57612e9b6160b4565b602002602001015186826002612eb19190615ee8565b612ebc9060016160a1565b81518110612ecc57612ecc6160b4565b6020026020010151604051602001612eee929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110612f1657612f166160b4565b6020026020010181815250505b600101612dd7565b508094508195508384604051602001612f4e929190918252602082015260400190565b6040516020818303038152906040528051906020012093508280612f7190616107565b9350505050612d64565b5f835f81518110612f8e57612f8e6160b4565b602002602001015190505f5f90505b82811015613040576040805160208101849052908101859052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083018790529082018690529250606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209350600101612f9d565b5095945050505050565b7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac61307481613df2565b63ffffffff841615806130925750607e5463ffffffff908116908516115b156130c9576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff85165f9081526082602052604081205463ffffffff169081900361312e576040517f74a086a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8181165f9081526081602052604090206007810154909187166801000000000000000090910467ffffffffffffffff160361319a576040517f4f61d51900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff86165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff1615159003613212576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018101546007830154700100000000000000000000000000000000900460ff9081167c0100000000000000000000000000000000000000000000000000000000909204161461328e576040517fb541abe200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001808201805491840180547fffffffffffffffffffffffff0000000000000000000000000000000000000000811673ffffffffffffffffffffffffffffffffffffffff9094169384178255915467ffffffffffffffff740100000000000000000000000000000000000000009182900416027fffffffff000000000000000000000000000000000000000000000000000000009092169092171790556007820180546801000000000000000063ffffffff8a16027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790555f61337484611354565b6007840180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831617905582546040517f4f1ef28600000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff8b811692634f1ef286926134069216908b908b9060040161613b565b5f604051808303815f87803b15801561341d575f5ffd5b505af115801561342f573d5f5f3e3d5ffd5b50506040805163ffffffff8c8116825267ffffffffffffffff86166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a2505050505050505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb6134b481613df2565b683635c9adc5dea000008211806134ce5750633b9aca0082105b15613505576040517f8586952500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200161191e565b5f8281526034602052604090206001015461355481613df2565b61182c8383614b9c565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f461358881613df2565b608780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff16179055604080517fdbc1697600000000000000000000000000000000000000000000000000000000815290517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169163dbc16976916004808301925f92919082900301818387803b158015613641575f5ffd5b505af1158015613653573d5f5f3e3d5ffd5b5050505061365f615089565b50565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e61368c81613df2565b67ffffffffffffffff84165f9081526083602052604090205463ffffffff16156136e2576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff87165f9081526082602052604090205463ffffffff1615613744576040517fd409b93000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f61375388888888875f613a99565b5f8080526002909101602052604090209390935550505050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59061379a81613df2565b607e80545f919082906137b29063ffffffff16615eff565b91906101000a81548163ffffffff021916908363ffffffff160217905590506040518060c001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018767ffffffffffffffff1681526020018660ff1681526020015f1515815260200185815250607f5f8363ffffffff1663ffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101601c6101000a81548160ff021916908360ff160217905550608082015181600101601d6101000a81548160ff02191690831515021790555060a082015181600201559050508063ffffffff167fa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b5289898989898960405161399d969594939291906161a4565b60405180910390a25050505050505050565b5f54610100900460ff16611a32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610d0a565b611ab18282614ae1565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b608080545f9182918290613ab29063ffffffff16615eff565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff16021790555060815f8263ffffffff1663ffffffff1681526020019081526020015f20915087825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550858260010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555086826001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084825f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550838260070160106101000a81548160ff021916908360ff1602179055508063ffffffff167fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850878a888888604051613d0695949392919067ffffffffffffffff958616815273ffffffffffffffffffffffffffffffffffffffff949094166020850152918416604084015260ff166060830152909116608082015260a00190565b60405180910390a2509695505050505050565b60855467ffffffffffffffff8281165f90815260048501602052604081205490924292613d4a929181169116615ec8565b67ffffffffffffffff1611159392505050565b60068101545f90700100000000000000000000000000000000900467ffffffffffffffff1615613dcf5750600681015467ffffffffffffffff70010000000000000000000000000000000090910481165f90815260049092016020526040909120546801000000000000000090041690565b506006015468010000000000000000900467ffffffffffffffff1690565b919050565b61365f8133615117565b60078801545f9067ffffffffffffffff9081169087161015613e4a576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff881615613f2d57600689015467ffffffffffffffff70010000000000000000000000000000000090910481169089161115613eba576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f90815260048a0160205260409020600281015481549092888116680100000000000000009092041614613f27576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50613fd8565b5067ffffffffffffffff85165f90815260028901602052604090205480613f80576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff6801000000000000000090910481169087161115613fd8576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff7001000000000000000000000000000000009091048116908816118061402057508767ffffffffffffffff168767ffffffffffffffff1611155b806140595750600689015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690881611155b15614090576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f90815260048b0160205260409020546801000000000000000090048116908616146140f4576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6141038a8888888689614c55565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516141379190616219565b602060405180830381855afa158015614152573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190614175919061604a565b61417f91906160e1565b60018c01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a916141e79188919060040161622f565b602060405180830381865afa158015614202573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614226919061626b565b61425c576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f90815260048c0160205260409020600201548590036142b3576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b5f5f6142cb8a613d5d565b60078b015490915067ffffffffffffffff908116908916101561431a576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8916156143ff5760068a015467ffffffffffffffff7001000000000000000000000000000000009091048116908a16111561438a576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff808a165f90815260048c01602052604090206002810154815490945090918a81166801000000000000000090920416146143f9576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506144a1565b67ffffffffffffffff88165f90815260028b016020526040902054915081614453576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168867ffffffffffffffff1611156144a1576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff16116144ee576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6144fd8b8a8a8a878b614c55565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516145319190616219565b602060405180830381855afa15801561454c573d5f5f3e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061456f919061604a565b61457991906160e1565b60018d01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a916145e19189919060040161622f565b602060405180830381865afa1580156145fc573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614620919061626b565b614656576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f614661848b616007565b90506146c2878267ffffffffffffffff1661467a61278f565b6146849190615ee8565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016919061517f565b80608460088282829054906101000a900467ffffffffffffffff166146e79190615ec8565b82546101009290920a67ffffffffffffffff818102199093169183160217909155608480547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff16700100000000000000000000000000000000428416021790558e546040517f32c2d153000000000000000000000000000000000000000000000000000000008152918d166004830152602482018b905233604483015273ffffffffffffffffffffffffffffffffffffffff1691506332c2d153906064015f604051808303815f87803b1580156147bc575f5ffd5b505af11580156147ce573d5f5f3e3d5ffd5b5050505050505050505050505050505050565b600682015467ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169082161115806148435750600682015467ffffffffffffffff7001000000000000000000000000000000009091048116908216115b1561487a576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f818152600485016020908152604080832080546006890180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000092839004909816918202979097178755600280830154828752908a01909452919093209190915560018201546005870155835477ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000909302929092179092557f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d614981612ca6565b6040518263ffffffff1660e01b815260040161499f91815260200190565b5f604051808303815f87803b1580156149b6575f5ffd5b505af11580156149c8573d5f5f3e3d5ffd5b5050855473ffffffffffffffffffffffffffffffffffffffff165f908152608260209081526040918290205460028701546001880154845167ffffffffffffffff898116825294810192909252818501529188166060830152915163ffffffff90921693507f581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b925081900360800190a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b158015614ac3575f5ffd5b505af1158015614ad5573d5f5f3e3d5ffd5b50505050611a3261520c565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611ab1575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1615611ab1575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b67ffffffffffffffff8086165f8181526003890160205260408082205493881682529020546060929115801590614c8a575081155b15614cc1576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80614cf8576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614d018461529e565b614d37576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3385838a8c5f0160149054906101000a900467ffffffffffffffff168d60010160149054906101000a900467ffffffffffffffff1689878d8f604051602001614d899a9998979695949392919061628a565b604051602081830303815290604052925050509695505050505050565b5f614db083613d5d565b9050815f80614dbf8484616007565b60855467ffffffffffffffff91821692505f91614de99168010000000000000000900416426160f4565b90505b8467ffffffffffffffff168467ffffffffffffffff1614614e755767ffffffffffffffff8085165f90815260038901602052604090206001810154909116821015614e5357600181015468010000000000000000900467ffffffffffffffff169450614e6f565b614e5d8686616007565b67ffffffffffffffff16935050614e75565b50614dec565b5f614e8084846160f4565b905083811015614ed757808403600c8111614e9b5780614e9e565b600c5b9050806103e80a81608560109054906101000a900461ffff1661ffff160a6086540281614ecd57614ecd616061565b0460865550614f46565b838103600c8111614ee85780614eeb565b600c5b90505f816103e80a82608560109054906101000a900461ffff1661ffff160a670de0b6b3a76400000281614f2157614f21616061565b04905080608654670de0b6b3a76400000281614f3f57614f3f616061565b0460865550505b683635c9adc5dea000006086541115614f6b57683635c9adc5dea00000608655614f82565b633b9aca006086541015614f8257633b9aca006086555b5050505050505050565b600681015467ffffffffffffffff78010000000000000000000000000000000000000000000000008204811670010000000000000000000000000000000090920416111561365f5760068101545f9061500c907801000000000000000000000000000000000000000000000000900467ffffffffffffffff166001615ec8565b90506150188282613d19565b15611ab15760068201545f90600290615050908490700100000000000000000000000000000000900467ffffffffffffffff16616007565b61505a9190616393565b6150649083615ec8565b90506150708382613d19565b1561507f5761182c83826147e1565b61182c83836147e1565b606f5460ff166150c5576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16611ab1576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261182c908490615321565b606f5460ff1615615249576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b5f67ffffffff0000000167ffffffffffffffff83161080156152d5575067ffffffff00000001604083901c67ffffffffffffffff16105b80156152f6575067ffffffff00000001608083901c67ffffffffffffffff16105b801561530d575067ffffffff0000000160c083901c105b1561531a57506001919050565b505f919050565b5f615382826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661542c9092919063ffffffff16565b80519091501561182c57808060200190518101906153a0919061626b565b61182c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610d0a565b606061543a84845f85615442565b949350505050565b6060824710156154d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610d0a565b5f5f8673ffffffffffffffffffffffffffffffffffffffff1685876040516154fc9190616219565b5f6040518083038185875af1925050503d805f8114615536576040519150601f19603f3d011682016040523d82523d5f602084013e61553b565b606091505b50915091506122d287838387606083156155dc5782515f036155d55773ffffffffffffffffffffffffffffffffffffffff85163b6155d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610d0a565b508161543a565b61543a83838151156155f15781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0a9190615b72565b610a10806163c383390190565b73ffffffffffffffffffffffffffffffffffffffff8116811461365f575f5ffd5b803567ffffffffffffffff81168114613ded575f5ffd5b5f5f5f5f5f5f5f5f5f5f6101408b8d031215615684575f5ffd5b8a3561568f81615632565b995061569d60208c01615653565b98506156ab60408c01615653565b975060608b01356156bb81615632565b965060808b01356156cb81615632565b955060a08b01356156db81615632565b945060c08b01356156eb81615632565b935060e08b01356156fb81615632565b925061570a6101008c01615653565b91506157196101208c01615653565b90509295989b9194979a5092959850565b803563ffffffff81168114613ded575f5ffd5b5f5f6040838503121561574e575f5ffd5b6157578361572a565b915061576560208401615653565b90509250929050565b5f6020828403121561577e575f5ffd5b61134b8261572a565b80610300810183101561134e575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b0312156157b0575f5ffd5b6157b98961572a565b97506157c760208a01615653565b96506157d560408a01615653565b95506157e360608a01615653565b94506157f160808a01615653565b935060a0890135925060c0890135915061580e8a60e08b01615787565b90509295985092959890939650565b5f5f5f5f5f5f5f5f6103e0898b031215615835575f5ffd5b61583e8961572a565b975061584c60208a01615653565b965061585a60408a01615653565b955061586860608a01615653565b94506080890135935060a0890135925060c089013561588681615632565b915061580e8a60e08b01615787565b5f602082840312156158a5575f5ffd5b813561ffff811681146158b6575f5ffd5b9392505050565b5f602082840312156158cd575f5ffd5b5035919050565b5f5f604083850312156158e5575f5ffd5b8235915060208301356158f781615632565b809150509250929050565b5f60208284031215615912575f5ffd5b61134b82615653565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112615957575f5ffd5b813567ffffffffffffffff8111156159715761597161591b565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff821117156159dd576159dd61591b565b6040528181528382016020018510156159f4575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f5f5f60e0888a031215615a26575f5ffd5b615a2f8861572a565b9650615a3d60208901615653565b95506040880135615a4d81615632565b94506060880135615a5d81615632565b93506080880135615a6d81615632565b925060a088013567ffffffffffffffff811115615a88575f5ffd5b615a948a828b01615948565b92505060c088013567ffffffffffffffff811115615ab0575f5ffd5b615abc8a828b01615948565b91505092959891949750929550565b5f5f5f5f5f5f60c08789031215615ae0575f5ffd5b615ae98761572a565b9550615af760208801615653565b9450615b0560408801615653565b959894975094956060810135955060808101359460a0909101359350915050565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61134b6020830184615b26565b5f5f60408385031215615b95575f5ffd5b615b9e83615653565b946020939093013593505050565b5f5f5f5f60608587031215615bbf575f5ffd5b8435615bca81615632565b9350615bd86020860161572a565b9250604085013567ffffffffffffffff811115615bf3575f5ffd5b8501601f81018713615c03575f5ffd5b803567ffffffffffffffff811115615c19575f5ffd5b876020828401011115615c2a575f5ffd5b949793965060200194505050565b5f60208284031215615c48575f5ffd5b81356158b681615632565b803560ff81168114613ded575f5ffd5b5f5f5f5f5f5f60c08789031215615c78575f5ffd5b8635615c8381615632565b95506020870135615c9381615632565b9450615ca160408801615653565b9350615caf60608801615653565b925060808701359150615cc460a08801615c53565b90509295509295509295565b5f5f5f5f5f5f60c08789031215615ce5575f5ffd5b8635615cf081615632565b95506020870135615d0081615632565b9450615d0e60408801615653565b9350615d1c60608801615c53565b92506080870135915060a087013567ffffffffffffffff811115615d3e575f5ffd5b615d4a89828a01615948565b9150509295509295509295565b80545f90600181811c90821680615d6f57607f821691505b602082108103615da6577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b81865260208601818015615dc15760018114615df557615e21565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008516825283151560051b82019550615e21565b5f878152602090205f5b85811015615e1b57815484820152600190910190602001615dff565b83019650505b505050505092915050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015260a060408201525f615e7660a0830186615d57565b8281036060840152615e888186615d57565b9150508260808301529695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff818116838216019081111561134e5761134e615e9b565b808202811582820484141761134e5761134e615e9b565b5f63ffffffff821663ffffffff8103615f1a57615f1a615e9b565b60010192915050565b73ffffffffffffffffffffffffffffffffffffffff8416815273ffffffffffffffffffffffffffffffffffffffff83166020820152606060408201525f615f6d6060830184615b26565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015263ffffffff8516604082015273ffffffffffffffffffffffffffffffffffffffff8416606082015260c060808201525f615fe860c0830185615b26565b82810360a0840152615ffa8185615b26565b9998505050505050505050565b67ffffffffffffffff828116828216039081111561134e5761134e615e9b565b5f67ffffffffffffffff821667ffffffffffffffff8103615f1a57615f1a615e9b565b5f6020828403121561605a575f5ffd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f8261609c5761609c616061565b500490565b8082018082111561134e5761134e615e9b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f826160ef576160ef616061565b500690565b8181038181111561134e5761134e615e9b565b5f8161611557616115615e9b565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b73ffffffffffffffffffffffffffffffffffffffff8416815260406020820152816040820152818360608301375f818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b73ffffffffffffffffffffffffffffffffffffffff8716815273ffffffffffffffffffffffffffffffffffffffff8616602082015267ffffffffffffffff8516604082015260ff8416606082015282608082015260c060a08201525f61620d60c0830184615b26565b98975050505050505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b6001811015616261578151835260209283019290910190600101616242565b5050509392505050565b5f6020828403121561627b575f5ffd5b815180151581146158b6575f5ffd5b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008b60601b1681528960148201528860348201527fffffffffffffffff0000000000000000000000000000000000000000000000008860c01b1660548201527fffffffffffffffff0000000000000000000000000000000000000000000000008760c01b16605c8201527fffffffffffffffff0000000000000000000000000000000000000000000000008660c01b16606482015284606c82015283608c8201528260ac82015261638260cc82018360c01b7fffffffffffffffff000000000000000000000000000000000000000000000000169052565b60d4019a9950505050505050505050565b5f67ffffffffffffffff8316806163ac576163ac616061565b8067ffffffffffffffff8416049150509291505056fe60a0604052604051610a10380380610a1083398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206109f05f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206109f05f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516105cb6104255f395f601001526105cb5ff3fe608060405261000c61000e565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633036100a7575f357fffffffff00000000000000000000000000000000000000000000000000000000167f4f1ef286000000000000000000000000000000000000000000000000000000001461009f5761009d6100ab565b565b61009d6100bb565b61009d5b61009d6100b66100e9565b61012d565b5f806100ca3660048184610410565b8101906100d79190610464565b915091506100e5828261014b565b5050565b5f6101287f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f5f375f5f365f845af43d5f5f3e808015610147573d5ff35b3d5ffd5b610154826101b2565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101aa576101a58282610285565b505050565b6100e5610304565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361021f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff16846040516102ae919061057f565b5f60405180830381855af49150503d805f81146102e6576040519150601f19603f3d011682016040523d82523d5f602084013e6102eb565b606091505b50915091506102fb85838361033c565b95945050505050565b341561009d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826103515761034c826103ce565b6103c7565b8151158015610375575073ffffffffffffffffffffffffffffffffffffffff84163b155b156103c4576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610216565b50805b9392505050565b8051156103de5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f8585111561041e575f5ffd5b8386111561042a575f5ffd5b5050820193919092039150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f5f60408385031215610475575f5ffd5b823573ffffffffffffffffffffffffffffffffffffffff81168114610498575f5ffd5b9150602083013567ffffffffffffffff8111156104b3575f5ffd5b8301601f810185136104c3575f5ffd5b803567ffffffffffffffff8111156104dd576104dd610437565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561054957610549610437565b604052818152828201602001871015610560575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220d77806b34ab0065d93e4bcbbd2cd8eea87814e7eac6721c2b3b7f2dd0abe324964736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103a26469706673582212201af9796ebf1648c81919d57b237256e9d498e5e681414e93aab29d180be2377d64736f6c634300081c0033",
}

// Polygonrollupmanagerpreviousv1tov2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Polygonrollupmanagerpreviousv1tov2MetaData.ABI instead.
var Polygonrollupmanagerpreviousv1tov2ABI = Polygonrollupmanagerpreviousv1tov2MetaData.ABI

// Polygonrollupmanagerpreviousv1tov2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Polygonrollupmanagerpreviousv1tov2MetaData.Bin instead.
var Polygonrollupmanagerpreviousv1tov2Bin = Polygonrollupmanagerpreviousv1tov2MetaData.Bin

// DeployPolygonrollupmanagerpreviousv1tov2 deploys a new Ethereum contract, binding an instance of Polygonrollupmanagerpreviousv1tov2 to it.
func DeployPolygonrollupmanagerpreviousv1tov2(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagerpreviousv1tov2, error) {
	parsed, err := Polygonrollupmanagerpreviousv1tov2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Polygonrollupmanagerpreviousv1tov2Bin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanagerpreviousv1tov2{Polygonrollupmanagerpreviousv1tov2Caller: Polygonrollupmanagerpreviousv1tov2Caller{contract: contract}, Polygonrollupmanagerpreviousv1tov2Transactor: Polygonrollupmanagerpreviousv1tov2Transactor{contract: contract}, Polygonrollupmanagerpreviousv1tov2Filterer: Polygonrollupmanagerpreviousv1tov2Filterer{contract: contract}}, nil
}

// Polygonrollupmanagerpreviousv1tov2 is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2 struct {
	Polygonrollupmanagerpreviousv1tov2Caller     // Read-only binding to the contract
	Polygonrollupmanagerpreviousv1tov2Transactor // Write-only binding to the contract
	Polygonrollupmanagerpreviousv1tov2Filterer   // Log filterer for contract events
}

// Polygonrollupmanagerpreviousv1tov2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonrollupmanagerpreviousv1tov2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonrollupmanagerpreviousv1tov2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Polygonrollupmanagerpreviousv1tov2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Polygonrollupmanagerpreviousv1tov2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Polygonrollupmanagerpreviousv1tov2Session struct {
	Contract     *Polygonrollupmanagerpreviousv1tov2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// Polygonrollupmanagerpreviousv1tov2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Polygonrollupmanagerpreviousv1tov2CallerSession struct {
	Contract *Polygonrollupmanagerpreviousv1tov2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// Polygonrollupmanagerpreviousv1tov2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Polygonrollupmanagerpreviousv1tov2TransactorSession struct {
	Contract     *Polygonrollupmanagerpreviousv1tov2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// Polygonrollupmanagerpreviousv1tov2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2Raw struct {
	Contract *Polygonrollupmanagerpreviousv1tov2 // Generic contract binding to access the raw methods on
}

// Polygonrollupmanagerpreviousv1tov2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2CallerRaw struct {
	Contract *Polygonrollupmanagerpreviousv1tov2Caller // Generic read-only contract binding to access the raw methods on
}

// Polygonrollupmanagerpreviousv1tov2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Polygonrollupmanagerpreviousv1tov2TransactorRaw struct {
	Contract *Polygonrollupmanagerpreviousv1tov2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanagerpreviousv1tov2 creates a new instance of Polygonrollupmanagerpreviousv1tov2, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousv1tov2(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanagerpreviousv1tov2, error) {
	contract, err := bindPolygonrollupmanagerpreviousv1tov2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2{Polygonrollupmanagerpreviousv1tov2Caller: Polygonrollupmanagerpreviousv1tov2Caller{contract: contract}, Polygonrollupmanagerpreviousv1tov2Transactor: Polygonrollupmanagerpreviousv1tov2Transactor{contract: contract}, Polygonrollupmanagerpreviousv1tov2Filterer: Polygonrollupmanagerpreviousv1tov2Filterer{contract: contract}}, nil
}

// NewPolygonrollupmanagerpreviousv1tov2Caller creates a new read-only instance of Polygonrollupmanagerpreviousv1tov2, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousv1tov2Caller(address common.Address, caller bind.ContractCaller) (*Polygonrollupmanagerpreviousv1tov2Caller, error) {
	contract, err := bindPolygonrollupmanagerpreviousv1tov2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2Caller{contract: contract}, nil
}

// NewPolygonrollupmanagerpreviousv1tov2Transactor creates a new write-only instance of Polygonrollupmanagerpreviousv1tov2, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousv1tov2Transactor(address common.Address, transactor bind.ContractTransactor) (*Polygonrollupmanagerpreviousv1tov2Transactor, error) {
	contract, err := bindPolygonrollupmanagerpreviousv1tov2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2Transactor{contract: contract}, nil
}

// NewPolygonrollupmanagerpreviousv1tov2Filterer creates a new log filterer instance of Polygonrollupmanagerpreviousv1tov2, bound to a specific deployed contract.
func NewPolygonrollupmanagerpreviousv1tov2Filterer(address common.Address, filterer bind.ContractFilterer) (*Polygonrollupmanagerpreviousv1tov2Filterer, error) {
	contract, err := bindPolygonrollupmanagerpreviousv1tov2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2Filterer{contract: contract}, nil
}

// bindPolygonrollupmanagerpreviousv1tov2 binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanagerpreviousv1tov2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Polygonrollupmanagerpreviousv1tov2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Polygonrollupmanagerpreviousv1tov2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Polygonrollupmanagerpreviousv1tov2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Polygonrollupmanagerpreviousv1tov2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.BridgeAddress(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.BridgeAddress(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ChainIDToRollupID(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ChainIDToRollupID(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetForcedBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetForcedBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetInputSnarkBytes(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetInputSnarkBytes(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRoleAdmin(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRoleAdmin(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupExitRoot(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupExitRoot(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetRollupPendingStateTransitions(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getRollupPendingStateTransitions", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesPendingState)).(*LegacyZKEVMStateVariablesPendingState)

	return out0, err

}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GlobalExitRootManager(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GlobalExitRootManager(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.HasRole(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.HasRole(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.IsEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.IsEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) IsPendingStateConsolidable(opts *bind.CallOpts, rollupID uint32, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "isPendingStateConsolidable", rollupID, pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID, pendingStateNum)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.LastAggregationTimestamp(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.LastAggregationTimestamp(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.MultiplierBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.MultiplierBatchFee(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.PendingStateTimeout(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.PendingStateTimeout(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) Pol() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Pol(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Pol(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupAddressToID(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupAddressToID(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RollupCount() (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupCount(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupCount(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupIDToRollupData(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupIDToRollupData(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupTypeCount(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupTypeCount(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupTypeMap(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RollupTypeMap(&_Polygonrollupmanagerpreviousv1tov2.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TotalSequencedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TotalSequencedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TotalVerifiedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TotalVerifiedBatches(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Caller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagerpreviousv1tov2.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2CallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagerpreviousv1tov2.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ActivateEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ActivateEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.AddExistingRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.AddExistingRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.AddNewRollupType(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.AddNewRollupType(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) ConsolidatePendingState(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "consolidatePendingState", rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ConsolidatePendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ConsolidatePendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.CreateNewRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.CreateNewRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.DeactivateEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.DeactivateEmergencyState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GrantRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.GrantRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) Initialize(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "initialize", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Initialize(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.Initialize(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ObsoleteRollupType(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ObsoleteRollupType(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.OnSequenceBatches(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.OnSequenceBatches(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) OverridePendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "overridePendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.OverridePendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.OverridePendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "proveNonDeterministicPendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RenounceRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RenounceRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RevokeRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.RevokeRole(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, role, account)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetBatchFee(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetBatchFee(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetPendingStateTimeout(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetPendingStateTimeout(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newTrustedAggregatorTimeout)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, newVerifyBatchTimeTarget)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.UpdateRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.UpdateRollup(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) VerifyBatches(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "verifyBatches", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatches(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatches(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Transactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Session) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2TransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagerpreviousv1tov2.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagerpreviousv1tov2.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2AddExistingRollup // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2AddExistingRollup)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2AddExistingRollup)
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
func (it *Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2AddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2AddExistingRollup struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2AddExistingRollupIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2AddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2AddExistingRollup)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseAddExistingRollup(log types.Log) (*Polygonrollupmanagerpreviousv1tov2AddExistingRollup, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2AddExistingRollup)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2AddNewRollupType // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2AddNewRollupType)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2AddNewRollupType)
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
func (it *Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2AddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2AddNewRollupType struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2AddNewRollupTypeIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2AddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2AddNewRollupType)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseAddNewRollupType(log types.Log) (*Polygonrollupmanagerpreviousv1tov2AddNewRollupType, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2AddNewRollupType)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState)
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
func (it *Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterConsolidatePendingState(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2ConsolidatePendingStateIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseConsolidatePendingState(log types.Log) (*Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2ConsolidatePendingState)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2CreateNewRollup // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2CreateNewRollup)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2CreateNewRollup)
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
func (it *Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2CreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2CreateNewRollup struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2CreateNewRollupIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2CreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2CreateNewRollup)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseCreateNewRollup(log types.Log) (*Polygonrollupmanagerpreviousv1tov2CreateNewRollup, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2CreateNewRollup)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated)
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
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2EmergencyStateActivatedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseEmergencyStateActivated(log types.Log) (*Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2EmergencyStateActivated)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated)
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
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivatedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseEmergencyStateDeactivated(log types.Log) (*Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2EmergencyStateDeactivated)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2InitializedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2Initialized // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2Initialized)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2Initialized)
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
func (it *Polygonrollupmanagerpreviousv1tov2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2Initialized represents a Initialized event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterInitialized(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2InitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2InitializedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2Initialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2Initialized)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseInitialized(log types.Log) (*Polygonrollupmanagerpreviousv1tov2Initialized, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2Initialized)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType)
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
func (it *Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2ObsoleteRollupTypeIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseObsoleteRollupType(log types.Log) (*Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2ObsoleteRollupType)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2OnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2OnSequenceBatches)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2OnSequenceBatches)
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
func (it *Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2OnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2OnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2OnSequenceBatchesIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2OnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2OnSequenceBatches)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseOnSequenceBatches(log types.Log) (*Polygonrollupmanagerpreviousv1tov2OnSequenceBatches, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2OnSequenceBatches)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2OverridePendingState // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2OverridePendingState)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2OverridePendingState)
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
func (it *Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2OverridePendingState represents a OverridePendingState event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2OverridePendingState struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterOverridePendingState(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2OverridePendingStateIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2OverridePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2OverridePendingState)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseOverridePendingState(log types.Log) (*Polygonrollupmanagerpreviousv1tov2OverridePendingState, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2OverridePendingState)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState)
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
func (it *Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingStateIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseProveNonDeterministicPendingState(log types.Log) (*Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2ProveNonDeterministicPendingState)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleAdminChanged)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleAdminChanged)
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
func (it *Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2RoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2RoleAdminChangedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2RoleAdminChanged)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseRoleAdminChanged(log types.Log) (*Polygonrollupmanagerpreviousv1tov2RoleAdminChanged, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2RoleAdminChanged)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2RoleGranted // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleGranted)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleGranted)
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
func (it *Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2RoleGranted represents a RoleGranted event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2RoleGrantedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2RoleGranted)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseRoleGranted(log types.Log) (*Polygonrollupmanagerpreviousv1tov2RoleGranted, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2RoleGranted)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleRevoked)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2RoleRevoked)
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
func (it *Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2RoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2RoleRevokedIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2RoleRevoked)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseRoleRevoked(log types.Log) (*Polygonrollupmanagerpreviousv1tov2RoleRevoked, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2RoleRevoked)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetBatchFee // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetBatchFee)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetBatchFee)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetBatchFee(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetBatchFeeIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetBatchFee)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetBatchFee(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetBatchFee, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetBatchFee)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFeeIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetMultiplierBatchFee(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetMultiplierBatchFee)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeoutIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetPendingStateTimeout(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetPendingStateTimeout)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetTrustedAggregator(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregator)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeoutIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetTrustedAggregatorTimeout)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget)
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
func (it *Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTargetIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2SetVerifyBatchTimeTarget)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2UpdateRollup // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2UpdateRollup)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2UpdateRollup)
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
func (it *Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2UpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2UpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2UpdateRollupIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2UpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2UpdateRollup)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseUpdateRollup(log types.Log) (*Polygonrollupmanagerpreviousv1tov2UpdateRollup, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2UpdateRollup)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2VerifyBatches // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2VerifyBatches)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2VerifyBatches)
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
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2VerifyBatches represents a VerifyBatches event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2VerifyBatches struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterVerifyBatches(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2VerifyBatchesIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2VerifyBatches, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2VerifyBatches)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseVerifyBatches(log types.Log) (*Polygonrollupmanagerpreviousv1tov2VerifyBatches, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2VerifyBatches)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator struct {
	Event *Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator)
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
		it.Event = new(Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator)
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
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanagerpreviousv1tov2 contract.
type Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator struct {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanagerpreviousv1tov2.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagerpreviousv1tov2.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagerpreviousv1tov2 *Polygonrollupmanagerpreviousv1tov2Filterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator, error) {
	event := new(Polygonrollupmanagerpreviousv1tov2VerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanagerpreviousv1tov2.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

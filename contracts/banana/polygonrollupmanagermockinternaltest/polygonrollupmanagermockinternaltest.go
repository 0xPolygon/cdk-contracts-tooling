// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanagermockinternaltest

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

// PolygonrollupmanagermockinternaltestMetaData contains all meta data concerning the Polygonrollupmanagermockinternaltest contract.
var PolygonrollupmanagermockinternaltestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchToRollback\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupPendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structLegacyZKEVMStateVariables.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"},{\"internalType\":\"contractPolygonZkEVMExistentEtrog\",\"name\":\"polygonZkEVM\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"zkEVMVerifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMForkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMChainID\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"batchToRollback\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801562000010575f80fd5b50604051620077563803806200775683398101604081905262000033916200013c565b6001600160a01b0380841660805280831660c052811660a0528282826200005962000065565b5050505050506200018d565b5f54610100900460ff1615620000d15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000122575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811462000139575f80fd5b50565b5f805f606084860312156200014f575f80fd5b83516200015c8162000124565b60208501519093506200016f8162000124565b6040850151909250620001828162000124565b809150509250925092565b60805160a05160c051617566620001f05f395f8181610bd001528181612c1c015261474101525f818161093e015281816136950152614b1201525f8181610b0c0152818161139e015281816115be015281816124ba01526149ec01526175665ff3fe608060405234801562000010575f80fd5b5060043610620003a4575f3560e01c806387c20c0111620001ef578063c4c928c21162000113578063dde0ff7711620000ab578063e46761c41162000083578063e46761c41462000bca578063f34eb8eb1462000bf2578063f4e926751462000c09578063f9c4c2ae1462000c1a575f80fd5b8063dde0ff771462000b7b578063dfdb8c5e1462000b9c578063e0bfd3d21462000bb3575f80fd5b8063d5073f6f11620000eb578063d5073f6f1462000b2e578063d547741f1462000b45578063d939b3151462000b5c578063dbc169761462000b71575f80fd5b8063c4c928c21462000ac7578063ceee281d1462000ade578063d02103ca1462000b06575f80fd5b8063a066215c1162000187578063a3c573eb116200015f578063a3c573eb1462000938578063afd23cbe1462000986578063b99d0ad714620009bd578063c1acbc341462000a9e575f80fd5b8063a066215c146200090f578063a217fddf1462000926578063a2967d99146200092e575f80fd5b806391d1485411620001c757806391d14854146200088f57806399f5634e14620008d75780639a908e7314620008e15780639c9f3dfe14620008f8575f80fd5b806387c20c01146200084a5780638bd4f07114620008615780638fd88cc21462000878575f80fd5b80632528016911620002d757806360469169116200026f578063727885e91162000247578063727885e914620007b45780637975fcfe14620007cb5780637fb6e76a14620007f1578063841b24d71462000819575f80fd5b806360469169146200068557806365c0504d146200068f5780637222020f146200079d575f80fd5b806336568abe11620002af57806336568abe146200060a578063394218e91462000621578063477fa270146200063857806355a71ee01462000641575f80fd5b80632528016914620005235780632f2ff15d14620005de57806330c27dde14620005f5575f80fd5b80631489ed10116200034b5780631796a1ae11620003235780631796a1ae14620004a75780631816b7e514620004ce5780632072f6c514620004e5578063248a9ca314620004ef575f80fd5b80631489ed10146200046b57806315064c9614620004825780631608859c1462000490575f80fd5b80630a0d9fbe116200037f5780630a0d9fbe146200041c57806311f6b287146200043d57806312b86e191462000454575f80fd5b80630645af0914620003a8578063066ec01214620003c1578063080b311114620003f4575b5f80fd5b620003bf620003b936600462005f05565b62000d84565b005b608454620003d69067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6200040b6200040536600462005fee565b6200126a565b6040519015158152602001620003eb565b608554620003d69068010000000000000000900467ffffffffffffffff1681565b620003d66200044e36600462006024565b62001293565b620003bf6200046536600462006052565b620012b2565b620003bf6200047c366004620060e4565b620014d2565b606f546200040b9060ff1681565b620003bf620004a136600462005fee565b620016b1565b607e54620004b89063ffffffff1681565b60405163ffffffff9091168152602001620003eb565b620003bf620004df36600462006169565b6200178d565b620003bf6200188a565b620005146200050036600462006193565b5f9081526034602052604090206001015490565b604051908152602001620003eb565b620005a96200053436600462005fee565b60408051606080820183525f808352602080840182905292840181905263ffffffff9590951685526081825282852067ffffffffffffffff9485168652600301825293829020825194850183528054855260010154808416918501919091526801000000000000000090049091169082015290565b604080518251815260208084015167ffffffffffffffff908116918301919091529282015190921690820152606001620003eb565b620003bf620005ef366004620061ab565b6200199e565b608754620003d69067ffffffffffffffff1681565b620003bf6200061b366004620061ab565b620019c6565b620003bf62000632366004620061dc565b62001a26565b60865462000514565b620005146200065236600462005fee565b63ffffffff82165f90815260816020908152604080832067ffffffffffffffff8516845260020190915290205492915050565b6200051462001b41565b62000745620006a036600462006024565b607f6020525f908152604090208054600182015460029092015473ffffffffffffffffffffffffffffffffffffffff918216929182169167ffffffffffffffff740100000000000000000000000000000000000000008204169160ff7c010000000000000000000000000000000000000000000000000000000083048116927d0100000000000000000000000000000000000000000000000000000000009004169086565b6040805173ffffffffffffffffffffffffffffffffffffffff978816815296909516602087015267ffffffffffffffff9093169385019390935260ff166060840152901515608083015260a082015260c001620003eb565b620003bf620007ae36600462006024565b62001b58565b620003bf620007c5366004620062de565b62001cd3565b620007e2620007dc366004620063a5565b62002285565b604051620003eb919062006476565b620004b862000802366004620061dc565b60836020525f908152604090205463ffffffff1681565b608454620003d6907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b620003bf6200085b366004620060e4565b620022b7565b620003bf6200087236600462006052565b620026b4565b620003bf620008893660046200648a565b62002785565b6200040b620008a0366004620061ab565b5f91825260346020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6200051462002bd5565b620003d6620008f2366004620064a9565b62002ce3565b620003bf62000909366004620061dc565b62002f33565b620003bf62000920366004620061dc565b6200301c565b620005145f81565b6200051462003106565b620009607f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001620003eb565b608554620009a990700100000000000000000000000000000000900461ffff1681565b60405161ffff9091168152602001620003eb565b62000a58620009ce36600462005fee565b60408051608080820183525f8083526020808401829052838501829052606093840182905263ffffffff9690961681526081865283812067ffffffffffffffff958616825260040186528390208351918201845280548086168352680100000000000000009004909416948101949094526001830154918401919091526002909101549082015290565b604051620003eb91905f60808201905067ffffffffffffffff80845116835280602085015116602084015250604083015160408301526060830151606083015292915050565b608454620003d690700100000000000000000000000000000000900467ffffffffffffffff1681565b620003bf62000ad8366004620064d4565b620034ee565b620004b862000aef3660046200654c565b60826020525f908152604090205463ffffffff1681565b620009607f000000000000000000000000000000000000000000000000000000000000000081565b620003bf62000b3f36600462006193565b6200352d565b620003bf62000b56366004620061ab565b620035e2565b608554620003d69067ffffffffffffffff1681565b620003bf6200360a565b608454620003d69068010000000000000000900467ffffffffffffffff1681565b620003bf62000bad3660046200656a565b62003714565b620003bf62000bc4366004620065aa565b620038f0565b620009607f000000000000000000000000000000000000000000000000000000000000000081565b620003bf62000c0336600462006622565b62003a4c565b608054620004b89063ffffffff1681565b62000cf662000c2b36600462006024565b60816020525f90815260409020805460018201546005830154600684015460079094015473ffffffffffffffffffffffffffffffffffffffff80851695740100000000000000000000000000000000000000009586900467ffffffffffffffff908116969286169592909204821693928282169268010000000000000000808404821693700100000000000000000000000000000000808204841694780100000000000000000000000000000000000000000000000090920484169380831693830416910460ff168c565b6040805173ffffffffffffffffffffffffffffffffffffffff9d8e16815267ffffffffffffffff9c8d1660208201529c909a16998c019990995296891660608b015260808a019590955292871660a089015290861660c0880152851660e0870152841661010086015283166101208501529190911661014083015260ff1661016082015261018001620003eb565b5f54600290610100900460ff1615801562000da557505f5460ff8083169116105b62000e37576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f805461010060ff84167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090921691909117179055608580546084805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8e8116919091029190911790915567016345785d8a00006086558c167fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116176907080000000000000000177fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff167103ea0000000000000000000000000000000017905562000f4062003c91565b62000f6c7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd48c62003d29565b62000f785f8862003d29565b62000fa47fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f5908862003d29565b62000fd07f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e8862003d29565b62000ffc7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac8862003d29565b620010287fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd8962003d29565b620010547fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd088962003d29565b620010807f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f48962003d29565b620010ac7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db18962003d29565b620010f87f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd47f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f062003d35565b620011247f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f08962003d29565b620011507f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb8962003d29565b6200119c7f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e7f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff28595162003d35565b620011c87f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e8762003d29565b620011f47f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff2859518762003d29565b620012005f3362003d29565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050505050505050565b63ffffffff82165f9081526081602052604081206200128a908362003d7f565b90505b92915050565b63ffffffff81165f9081526081602052604081206200128d9062003dc5565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4620012de8162003e5b565b63ffffffff89165f90815260816020526040902062001304818a8a8a8a8a8a8a62003e67565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f9081526002840160205260409020869055600583018790559054700100000000000000000000000000000000900416156200139c576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d620013e262003106565b6040518263ffffffff1660e01b81526004016200140191815260200190565b5f604051808303815f87803b15801562001419575f80fd5b505af11580156200142c573d5f803e3d5ffd5b50506084805477ffffffffffffffffffffffffffffffffffffffffffffffff167a093a8000000000000000000000000000000000000000000000000017905550506040805167ffffffffffffffff881681526020810186905290810186905233606082015263ffffffff8b16907f3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e9060800160405180910390a250505050505050505050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4620014fe8162003e5b565b63ffffffff89165f90815260816020526040902062001524818a8a8a8a8a8a8a62004346565b6006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8a81169182029290921783555f908152600284016020526040902087905560058301889055905470010000000000000000000000000000000090041615620015bc576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d6200160262003106565b6040518263ffffffff1660e01b81526004016200162191815260200190565b5f604051808303815f87803b15801562001639575f80fd5b505af11580156200164c573d5f803e3d5ffd5b50506040805167ffffffffffffffff8b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b63ffffffff82165f9081526081602090815260408083203384527fc17b14a573f65366cdad721c7c0a0f76536bb4a86b935cdac44610e4f010b52a9092529091205460ff166200177c57606f5460ff161562001739576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62001745818362003d7f565b6200177c576040517f0ce9e4a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200178881836200488c565b505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1620017b98162003e5b565b6103e88261ffff161080620017d357506103ff8261ffff16115b156200180b576040517f4c2533c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffff0000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000061ffff8516908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a15050565b335f9081527f8875b94af5657a2903def9906d67a3f42d8a836d24b5602c00f00fc855339fcd602052604090205460ff166200199257608454700100000000000000000000000000000000900467ffffffffffffffff16158062001928575060845442906200191c9062093a8090700100000000000000000000000000000000900467ffffffffffffffff16620066e2565b67ffffffffffffffff16115b806200195a575060875442906200194e9062093a809067ffffffffffffffff16620066e2565b67ffffffffffffffff16115b1562001992576040517fd257555a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6200199c62004b10565b565b5f82815260346020526040902060010154620019ba8162003e5b565b62001788838362004b97565b73ffffffffffffffffffffffffffffffffffffffff8116331462001a16576040517f5a568e6800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62001a22828262004c53565b5050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db162001a528162003e5b565b606f5460ff1662001ac35760845467ffffffffffffffff780100000000000000000000000000000000000000000000000090910481169083161062001ac3576040517f401636df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6084805477ffffffffffffffffffffffffffffffffffffffffffffffff16780100000000000000000000000000000000000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1906020016200187e565b5f608654606462001b5391906200670d565b905090565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd62001b848162003e5b565b63ffffffff8216158062001ba35750607e5463ffffffff908116908316115b1562001bdb576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff161515900362001c54576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001810180547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167d01000000000000000000000000000000000000000000000000000000000017905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd0862001cff8162003e5b565b63ffffffff8816158062001d1e5750607e5463ffffffff908116908916115b1562001d56576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff161515900362001dcf576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff67ffffffffffffffff8916111562001e18576040517f4c753f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88165f9081526083602052604090205463ffffffff161562001e6f576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608080545f9190829062001e899063ffffffff1662006727565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f808252602082019283905293945073ffffffffffffffffffffffffffffffffffffffff90921691309162001ee39062005ebd565b62001ef1939291906200674c565b604051809103905ff08015801562001f0b573d5f803e3d5ffd5b5090508160835f8c67ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508160825f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8463ffffffff1663ffffffff1681526020019081526020015f20905081815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360010160149054906101000a900467ffffffffffffffff168160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a815f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508360020154816002015f8067ffffffffffffffff1681526020019081526020015f20819055508b63ffffffff168160070160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083600101601c9054906101000a900460ff168160070160106101000a81548160ff021916908360ff1602179055508263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418d848e8c604051620021e2949392919063ffffffff94909416845273ffffffffffffffffffffffffffffffffffffffff928316602085015267ffffffffffffffff91909116604084015216606082015260800190565b60405180910390a26040517f7125702200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83169063712570229062002248908d908d9088908e908e908e906004016200678f565b5f604051808303815f87803b15801562002260575f80fd5b505af115801562002273573d5f803e3d5ffd5b50505050505050505050505050505050565b63ffffffff86165f908152608160205260409020606090620022ac90878787878762004d0d565b979650505050505050565b606f5460ff1615620022f5576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f90815260816020908152604080832060845467ffffffffffffffff8a81168652600383019094529190932060010154429262002359927801000000000000000000000000000000000000000000000000900481169116620066e2565b67ffffffffffffffff1611156200239c576040517f8a0704d300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8620023ab8888620067fe565b67ffffffffffffffff161115620023ee576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62002400818989898989898962004346565b6200240c818762004ed4565b60855467ffffffffffffffff165f0362002552576006810180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8981169182029290921783555f908152600284016020526040902086905560058301879055905470010000000000000000000000000000000090041615620024b8576006810180546fffffffffffffffffffffffffffffffff1690555b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d620024fe62003106565b6040518263ffffffff1660e01b81526004016200251d91815260200190565b5f604051808303815f87803b15801562002535575f80fd5b505af115801562002548573d5f803e3d5ffd5b5050505062002654565b6200255d81620050db565b600681018054700100000000000000000000000000000000900467ffffffffffffffff169060106200258f8362006822565b825467ffffffffffffffff9182166101009390930a92830292820219169190911790915560408051608081018252428316815289831660208083019182528284018b8152606084018b81526006890154700100000000000000000000000000000000900487165f90815260048a0190935294909120925183549251861668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009093169516949094171781559151600183015551600290910155505b6040805167ffffffffffffffff8816815260208101869052908101869052339063ffffffff8b16907faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b49060600160405180910390a3505050505050505050565b606f5460ff1615620026f2576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff88165f90815260816020526040902062002718818989898989898962003e67565b67ffffffffffffffff87165f9081526004820160209081526040918290206002015482519081529081018590527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a16200277a62004b10565b505050505050505050565b335f9081527ff14f5a8ad59d90593602e905b358229bff5ceea677d5bf0f5a1701793550a9a6602052604090205460ff161580156200286057503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b81526004016020604051808303815f875af115801562002821573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002847919062006841565b73ffffffffffffffffffffffffffffffffffffffff1614155b1562002898576040517f1a06d0fe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526082602052604081205463ffffffff1690819003620028fe576040517f74a086a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff81165f908152608160205260409020600681015467ffffffffffffffff90811690841681111580620029525750600682015467ffffffffffffffff680100000000000000009091048116908516105b156200298a576040517fcb23ebdf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805b8467ffffffffffffffff168167ffffffffffffffff161462002a5f5767ffffffffffffffff8082165f908152600385016020526040902060010154680100000000000000009004811690861681101562002a12576040517f9753965f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff9091165f908152600384016020526040812090815560010180547fffffffffffffffffffffffffffffffff000000000000000000000000000000001690556200298c565b6006830180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff871617905562002aa18583620067fe565b608480545f9062002abe90849067ffffffffffffffff16620067fe565b82546101009290920a67ffffffffffffffff81810219909316918316021790915586165f818152600386016020526040908190205490517f669adece0000000000000000000000000000000000000000000000000000000081526004810192909252602482015273ffffffffffffffffffffffffffffffffffffffff8816915063669adece906044015f604051808303815f87803b15801562002b5f575f80fd5b505af115801562002b72573d5f803e3d5ffd5b5050505067ffffffffffffffff85165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a3505050505050565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801562002c62573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062002c8891906200685f565b6084549091505f9062002cb49067ffffffffffffffff68010000000000000000820481169116620067fe565b67ffffffffffffffff169050805f0362002cd0575f9250505090565b62002cdc8183620068a4565b9250505090565b606f545f9060ff161562002d23576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f9081526082602052604081205463ffffffff169081900362002d73576040517f71653c1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8367ffffffffffffffff165f0362002db7576040517f2590ccf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff81165f908152608160205260408120608480549192879262002dea90849067ffffffffffffffff16620066e2565b82546101009290920a67ffffffffffffffff81810219909316918316021790915560068301541690505f62002e208783620066e2565b60068401805467ffffffffffffffff8084167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009092168217909255604080516060810182528a815242841660208083019182528886168385019081525f95865260038b019091529290932090518155915160019290920180549151841668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009092169290931691909117179055905062002ee483620050db565b60405167ffffffffffffffff8216815263ffffffff8516907f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a259060200160405180910390a29695505050505050565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db162002f5f8162003e5b565b606f5460ff1662002fb45760855467ffffffffffffffff9081169083161062002fb4576040517f48a05a9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff84169081179091556040519081527fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75906020016200187e565b7fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db1620030488162003e5b565b620151808267ffffffffffffffff16111562003090576040517fe067dfe800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608580547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8516908102919091179091556040519081527f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28906020016200187e565b6080545f9063ffffffff168082036200312057505f919050565b5f8167ffffffffffffffff8111156200313d576200313d620061f8565b60405190808252806020026020018201604052801562003167578160200160208202803683370190505b5090505f5b82811015620031d75760815f62003185836001620068ba565b63ffffffff1663ffffffff1681526020019081526020015f2060050154828281518110620031b757620031b7620068d0565b602090810291909101015280620031ce81620068fd565b9150506200316c565b505f60205b8360011462003430575f620031f360028662006937565b62003200600287620068a4565b6200320c9190620068ba565b90505f8167ffffffffffffffff8111156200322b576200322b620061f8565b60405190808252806020026020018201604052801562003255578160200160208202803683370190505b5090505f5b82811015620033dc57620032706001846200694d565b811480156200328b57506200328760028862006937565b6001145b156200331357856200329f8260026200670d565b81518110620032b257620032b2620068d0565b602002602001015185604051602001620032d6929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110620033015762003301620068d0565b602002602001018181525050620033c7565b85620033218260026200670d565b81518110620033345762003334620068d0565b6020026020010151868260026200334c91906200670d565b62003359906001620068ba565b815181106200336c576200336c620068d0565b60200260200101516040516020016200338f929190918252602082015260400190565b60405160208183030381529060405280519060200120828281518110620033ba57620033ba620068d0565b6020026020010181815250505b80620033d381620068fd565b9150506200325a565b50809450819550838460405160200162003400929190918252602082015260400190565b6040516020818303038152906040528051906020012093508280620034259062006963565b9350505050620031dc565b5f835f81518110620034465762003446620068d0565b602002602001015190505f5b82811015620034e4576040805160208101849052908101859052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012090830187905290820186905292506060016040516020818303038152906040528051906020012093508080620034db90620068fd565b91505062003452565b5095945050505050565b7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac6200351a8162003e5b565b62003527848484620051eb565b50505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb620035598162003e5b565b683635c9adc5dea00000821180620035745750633b9aca0082105b15620035ac576040517f8586952500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2906020016200187e565b5f82815260346020526040902060010154620035fe8162003e5b565b62001788838362004c53565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f4620036368162003e5b565b608780547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff16179055604080517fdbc1697600000000000000000000000000000000000000000000000000000000815290517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169163dbc16976916004808301925f92919082900301818387803b158015620036f0575f80fd5b505af115801562003703573d5f803e3d5ffd5b50505050620037116200567d565b50565b3373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b81526004016020604051808303815f875af115801562003776573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200379c919062006841565b73ffffffffffffffffffffffffffffffffffffffff1614620037ea576040517f696072e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82165f9081526082602090815260408083205463ffffffff16835260819091529020600681015467ffffffffffffffff80821668010000000000000000909204161462003878576040517fcc862d4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600781015463ffffffff83166801000000000000000090910467ffffffffffffffff1610620038d3576040517f3e37e23300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080515f815260208101909152620017889084908490620051eb565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e6200391c8162003e5b565b67ffffffffffffffff84165f9081526083602052604090205463ffffffff161562003973576040517f6f91fc1200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff67ffffffffffffffff85161115620039bc576040517f4c753f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff87165f9081526082602052604090205463ffffffff161562003a1f576040517fd409b93000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f62003a2f88888888876200570c565b5f8080526002909101602052604090209390935550505050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59062003a788162003e5b565b607e80545f9190829062003a929063ffffffff1662006727565b91906101000a81548163ffffffff021916908363ffffffff160217905590506040518060c001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018767ffffffffffffffff1681526020018660ff1681526020015f1515815260200185815250607f5f8363ffffffff1663ffffffff1681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101601c6101000a81548160ff021916908360ff160217905550608082015181600101601d6101000a81548160ff02191690831515021790555060a082015181600201559050508063ffffffff167fa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b5289898989898960405162003c7f969594939291906200699a565b60405180910390a25050505050505050565b5f54610100900460ff166200199c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840162000e2e565b62001a22828262004b97565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b60855467ffffffffffffffff8281165f9081526004850160205260408120549092429262003db2929181169116620066e2565b67ffffffffffffffff1611159392505050565b60068101545f90700100000000000000000000000000000000900467ffffffffffffffff161562003e385750600681015467ffffffffffffffff70010000000000000000000000000000000090910481165f90815260049092016020526040909120546801000000000000000090041690565b506006015468010000000000000000900467ffffffffffffffff1690565b919050565b6200371181336200598e565b60078801545f9067ffffffffffffffff908116908716101562003eb6576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff88161562003f9d57600689015467ffffffffffffffff7001000000000000000000000000000000009091048116908916111562003f28576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5067ffffffffffffffff8088165f90815260048a016020526040902060028101548154909288811668010000000000000000909204161462003f96576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b506200404a565b5067ffffffffffffffff85165f9081526002890160205260409020548062003ff1576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff68010000000000000000909104811690871611156200404a576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600689015467ffffffffffffffff700100000000000000000000000000000000909104811690881611806200409357508767ffffffffffffffff168767ffffffffffffffff1611155b80620040cd5750600689015467ffffffffffffffff7801000000000000000000000000000000000000000000000000909104811690881611155b1562004105576040517fbfa7079f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8781165f90815260048b0160205260409020546801000000000000000090048116908616146200416a576040517f32a2a77f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6200417b8a888888868962004d0d565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051620041b19190620069ff565b602060405180830381855afa158015620041cd573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190620041f291906200685f565b620041fe919062006937565b60018c01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a91620042689188919060040162006a1c565b602060405180830381865afa15801562004284573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620042aa919062006a58565b620042e1576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff89165f90815260048c01602052604090206002015485900362004339576040517fa47276bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b5f80620043538a62003dc5565b60078b015490915067ffffffffffffffff9081169089161015620043a3576040517fead1340b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8916156200448c5760068a015467ffffffffffffffff7001000000000000000000000000000000009091048116908a16111562004415576040517fbb14c20500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff808a165f90815260048c01602052604090206002810154815490945090918a811668010000000000000000909204161462004485576040517f2bd2e3e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5062004530565b67ffffffffffffffff88165f90815260028b016020526040902054915081620044e1576040517f4997b98600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168867ffffffffffffffff16111562004530576040517f1e56e9e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff168767ffffffffffffffff16116200457e576040517fb9b18f5700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6200458f8b8a8a8a878b62004d0d565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001600283604051620045c59190620069ff565b602060405180830381855afa158015620045e1573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906200460691906200685f565b62004612919062006937565b60018d01546040805160208101825283815290517f9121da8a00000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff90911691639121da8a916200467c9189919060040162006a1c565b602060405180830381865afa15801562004698573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620046be919062006a58565b620046f5576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f62004702848b620067fe565b905062004769878267ffffffffffffffff166200471e62002bd5565b6200472a91906200670d565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169190620059f7565b80608460088282829054906101000a900467ffffffffffffffff16620047909190620066e2565b82546101009290920a67ffffffffffffffff818102199093169183160217909155608480547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff16700100000000000000000000000000000000428416021790558e546040517f32c2d153000000000000000000000000000000000000000000000000000000008152918d166004830152602482018b905233604483015273ffffffffffffffffffffffffffffffffffffffff1691506332c2d153906064015f604051808303815f87803b15801562004866575f80fd5b505af115801562004879573d5f803e3d5ffd5b5050505050505050505050505050505050565b600682015467ffffffffffffffff78010000000000000000000000000000000000000000000000009091048116908216111580620048ef5750600682015467ffffffffffffffff7001000000000000000000000000000000009091048116908216115b1562004927576040517fd086b70b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8181165f818152600485016020908152604080832080546006890180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000092839004909816918202979097178755600280830154828752908a01909452919093209190915560018201546005870155835477ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000909302929092179092557f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166333d6247d62004a3062003106565b6040518263ffffffff1660e01b815260040162004a4f91815260200190565b5f604051808303815f87803b15801562004a67575f80fd5b505af115801562004a7a573d5f803e3d5ffd5b5050855473ffffffffffffffffffffffffffffffffffffffff165f908152608260209081526040918290205460028701546001880154845167ffffffffffffffff898116825294810192909252818501529188166060830152915163ffffffff90921693507f581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b925081900360800190a250505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b15801562004b76575f80fd5b505af115801562004b89573d5f803e3d5ffd5b505050506200199c62005a86565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1662001a22575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161562001a22575f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b67ffffffffffffffff8086165f818152600389016020526040808220549388168252902054606092911580159062004d43575081155b1562004d7b576040517f6818c29e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8062004db3576040517f66385b5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62004dbe8462005b19565b62004df5576040517f176b913c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b885460018a01546040517fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003360601b16602082015260348101889052605481018590527fffffffffffffffff00000000000000000000000000000000000000000000000060c08c811b821660748401527401000000000000000000000000000000000000000094859004811b8216607c84015293909204831b82166084820152608c810187905260ac810184905260cc81018990529189901b1660ec82015260f401604051602081830303815290604052925050509695505050505050565b5f62004ee08362003dc5565b9050815f8062004ef18484620067fe565b60855467ffffffffffffffff91821692505f9162004f1d9168010000000000000000900416426200694d565b90505b8467ffffffffffffffff168467ffffffffffffffff161462004fb05767ffffffffffffffff8085165f9081526003890160205260409020600181015490911682101562004f8a57600181015468010000000000000000900467ffffffffffffffff16945062004fa9565b62004f968686620067fe565b67ffffffffffffffff1693505062004fb0565b5062004f20565b5f62004fbd84846200694d565b9050838110156200501b57808403600c811162004fdb578062004fde565b600c5b9050806103e80a81608560109054906101000a900461ffff1661ffff160a608654028162005010576200501062006877565b046086555062005092565b838103600c81116200502e578062005031565b600c5b90505f816103e80a82608560109054906101000a900461ffff1661ffff160a670de0b6b3a764000002816200506a576200506a62006877565b04905080608654670de0b6b3a764000002816200508b576200508b62006877565b0460865550505b683635c9adc5dea000006086541115620050b957683635c9adc5dea00000608655620050d1565b633b9aca006086541015620050d157633b9aca006086555b5050505050505050565b600681015467ffffffffffffffff780100000000000000000000000000000000000000000000000082048116700100000000000000000000000000000000909204161115620037115760068101545f906200515e907801000000000000000000000000000000000000000000000000900467ffffffffffffffff166001620066e2565b90506200516c828262003d7f565b1562001a225760068201545f90600290620051a7908490700100000000000000000000000000000000900467ffffffffffffffff16620067fe565b620051b3919062006a79565b620051bf9083620066e2565b9050620051cd838262003d7f565b15620051df576200178883826200488c565b6200178883836200488c565b63ffffffff821615806200520a5750607e5463ffffffff908116908316115b1562005242576040517f7512e5cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f9081526082602052604081205463ffffffff1690819003620052a8576040517f74a086a300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8181165f9081526081602052604090206007810154909185166801000000000000000090910467ffffffffffffffff160362005315576040517f4f61d51900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001808201547d010000000000000000000000000000000000000000000000000000000000900460ff16151590036200538e576040517f3b8d3d9900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018101546007830154700100000000000000000000000000000000900460ff9081167c010000000000000000000000000000000000000000000000000000000090920416146200540b576040517fb541abe200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018181018054918401805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff000000000000000000000000000000000000000084168117825591547fffffffff00000000000000000000000000000000000000000000000000000000909316909117740100000000000000000000000000000000000000009283900467ffffffffffffffff9081169093021790556007830180547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff1663ffffffff881668010000000000000000021790556006830154780100000000000000000000000000000000000000000000000081048216700100000000000000000000000000000000909104909116146200555b576040517f9d59507b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f620055678462001293565b6007840180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831617905582546040517f4f1ef28600000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff89811692634f1ef28692620055f9921690899060040162006aa2565b5f604051808303815f87803b15801562005611575f80fd5b505af115801562005624573d5f803e3d5ffd5b50506040805163ffffffff8a8116825267ffffffffffffffff86166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff16620056ba576040517f5386698100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b608080545f9182918290620057279063ffffffff1662006727565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f8667ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff16021790555060815f8263ffffffff1663ffffffff1681526020019081526020015f20915086825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550848260010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555085826001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083825f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550828260070160106101000a81548160ff021916908360ff1602179055508063ffffffff167fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850868987875f6040516200597c95949392919067ffffffffffffffff958616815273ffffffffffffffffffffffffffffffffffffffff949094166020850152918416604084015260ff166060830152909116608082015260a00190565b60405180910390a25095945050505050565b5f82815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1662001a22576040517fec2b7c3e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526200178890849062005ba0565b606f5460ff161562005ac4576040517f2f0047fc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b5f67ffffffff0000000167ffffffffffffffff831610801562005b51575067ffffffff00000001604083901c67ffffffffffffffff16105b801562005b73575067ffffffff00000001608083901c67ffffffffffffffff16105b801562005b8b575067ffffffff0000000160c083901c105b1562005b9957506001919050565b505f919050565b5f62005c03826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1662005cb29092919063ffffffff16565b80519091501562001788578080602001905181019062005c24919062006a58565b62001788576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840162000e2e565b606062005cc284845f8562005cca565b949350505050565b60608247101562005d5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840162000e2e565b5f808673ffffffffffffffffffffffffffffffffffffffff16858760405162005d889190620069ff565b5f6040518083038185875af1925050503d805f811462005dc4576040519150601f19603f3d011682016040523d82523d5f602084013e62005dc9565b606091505b5091509150620022ac878383876060831562005e705782515f0362005e685773ffffffffffffffffffffffffffffffffffffffff85163b62005e68576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000e2e565b508162005cc2565b62005cc2838381511562005e875781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000e2e919062006476565b610a5e8062006ad383390190565b73ffffffffffffffffffffffffffffffffffffffff8116811462003711575f80fd5b803567ffffffffffffffff8116811462003e56575f80fd5b5f805f805f805f805f806101408b8d03121562005f20575f80fd5b8a3562005f2d8162005ecb565b995062005f3d60208c0162005eed565b985062005f4d60408c0162005eed565b975060608b013562005f5f8162005ecb565b965060808b013562005f718162005ecb565b955060a08b013562005f838162005ecb565b945060c08b013562005f958162005ecb565b935060e08b013562005fa78162005ecb565b925062005fb86101008c0162005eed565b915062005fc96101208c0162005eed565b90509295989b9194979a5092959850565b803563ffffffff8116811462003e56575f80fd5b5f806040838503121562006000575f80fd5b6200600b8362005fda565b91506200601b6020840162005eed565b90509250929050565b5f6020828403121562006035575f80fd5b6200128a8262005fda565b8061030081018310156200128d575f80fd5b5f805f805f805f806103e0898b0312156200606b575f80fd5b620060768962005fda565b97506200608660208a0162005eed565b96506200609660408a0162005eed565b9550620060a660608a0162005eed565b9450620060b660808a0162005eed565b935060a0890135925060c08901359150620060d58a60e08b0162006040565b90509295985092959890939650565b5f805f805f805f806103e0898b031215620060fd575f80fd5b620061088962005fda565b97506200611860208a0162005eed565b96506200612860408a0162005eed565b95506200613860608a0162005eed565b94506080890135935060a0890135925060c0890135620061588162005ecb565b9150620060d58a60e08b0162006040565b5f602082840312156200617a575f80fd5b813561ffff811681146200618c575f80fd5b9392505050565b5f60208284031215620061a4575f80fd5b5035919050565b5f8060408385031215620061bd575f80fd5b823591506020830135620061d18162005ecb565b809150509250929050565b5f60208284031215620061ed575f80fd5b6200128a8262005eed565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f67ffffffffffffffff80841115620062425762006242620061f8565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156200628b576200628b620061f8565b81604052809350858152868686011115620062a4575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112620062cd575f80fd5b6200128a8383356020850162006225565b5f805f805f805f60e0888a031215620062f5575f80fd5b620063008862005fda565b9650620063106020890162005eed565b95506040880135620063228162005ecb565b94506060880135620063348162005ecb565b93506080880135620063468162005ecb565b925060a088013567ffffffffffffffff8082111562006363575f80fd5b620063718b838c01620062bd565b935060c08a013591508082111562006387575f80fd5b50620063968a828b01620062bd565b91505092959891949750929550565b5f805f805f8060c08789031215620063bb575f80fd5b620063c68762005fda565b9550620063d66020880162005eed565b9450620063e66040880162005eed565b9350606087013592506080870135915060a087013590509295509295509295565b5f5b838110156200642357818101518382015260200162006409565b50505f910152565b5f81518084526200644481602086016020860162006407565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f6200128a60208301846200642b565b5f80604083850312156200649c575f80fd5b82356200600b8162005ecb565b5f8060408385031215620064bb575f80fd5b620064c68362005eed565b946020939093013593505050565b5f805f60608486031215620064e7575f80fd5b8335620064f48162005ecb565b9250620065046020850162005fda565b9150604084013567ffffffffffffffff81111562006520575f80fd5b8401601f8101861362006531575f80fd5b620065428682356020840162006225565b9150509250925092565b5f602082840312156200655d575f80fd5b81356200618c8162005ecb565b5f80604083850312156200657c575f80fd5b8235620065898162005ecb565b91506200601b6020840162005fda565b803560ff8116811462003e56575f80fd5b5f805f805f8060c08789031215620065c0575f80fd5b8635620065cd8162005ecb565b95506020870135620065df8162005ecb565b9450620065ef6040880162005eed565b9350620065ff6060880162005eed565b9250608087013591506200661660a0880162006599565b90509295509295509295565b5f805f805f8060c0878903121562006638575f80fd5b8635620066458162005ecb565b95506020870135620066578162005ecb565b9450620066676040880162005eed565b9350620066776060880162006599565b92506080870135915060a087013567ffffffffffffffff8111156200669a575f80fd5b620066a889828a01620062bd565b9150509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff818116838216019080821115620067065762006706620066b5565b5092915050565b80820281158282048414176200128d576200128d620066b5565b5f63ffffffff808316818103620067425762006742620066b5565b6001019392505050565b5f73ffffffffffffffffffffffffffffffffffffffff8086168352808516602084015250606060408301526200678660608301846200642b565b95945050505050565b5f73ffffffffffffffffffffffffffffffffffffffff8089168352808816602084015263ffffffff8716604084015280861660608401525060c06080830152620067dd60c08301856200642b565b82810360a0840152620067f181856200642b565b9998505050505050505050565b67ffffffffffffffff828116828216039080821115620067065762006706620066b5565b5f67ffffffffffffffff808316818103620067425762006742620066b5565b5f6020828403121562006852575f80fd5b81516200618c8162005ecb565b5f6020828403121562006870575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82620068b557620068b562006877565b500490565b808201808211156200128d576200128d620066b5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620069305762006930620066b5565b5060010190565b5f8262006948576200694862006877565b500690565b818103818111156200128d576200128d620066b5565b5f81620069745762006974620066b5565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b5f73ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525067ffffffffffffffff8616604083015260ff8516606083015283608083015260c060a0830152620069f360c08301846200642b565b98975050505050505050565b5f825162006a1281846020870162006407565b9190910192915050565b6103208101610300808584378201835f5b600181101562006a4e57815183526020928301929091019060010162006a2d565b5050509392505050565b5f6020828403121562006a69575f80fd5b815180151581146200618c575f80fd5b5f67ffffffffffffffff8084168062006a965762006a9662006877565b92169190910492915050565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201525f62005cc260408301846200642b56fe60a060405260405162000a5e38038062000a5e833981016040819052620000269162000375565b828162000034828262000060565b50506001600160a01b038216608052620000576200005160805190565b620000c5565b5050506200046c565b6200006b8262000136565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115620000b757620000b28282620001b5565b505050565b620000c16200022e565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f620001065f8051602062000a3e833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1620001338162000250565b50565b806001600160a01b03163b5f036200017157604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f80846001600160a01b031684604051620001d391906200044f565b5f60405180830381855af49150503d805f81146200020d576040519150601f19603f3d011682016040523d82523d5f602084013e62000212565b606091505b5090925090506200022585838362000291565b95945050505050565b34156200024e5760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b0381166200027b57604051633173bdd160e11b81525f600482015260240162000168565b805f8051602062000a3e83398151915262000194565b606082620002aa57620002a482620002f7565b620002f0565b8151158015620002c257506001600160a01b0384163b155b15620002ed57604051639996b31560e01b81526001600160a01b038516600482015260240162000168565b50805b9392505050565b805115620003085780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811462000338575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200036d57818101518382015260200162000353565b50505f910152565b5f805f6060848603121562000388575f80fd5b620003938462000321565b9250620003a36020850162000321565b60408501519092506001600160401b0380821115620003c0575f80fd5b818601915086601f830112620003d4575f80fd5b815181811115620003e957620003e96200033d565b604051601f8201601f19908116603f011681019083821181831017156200041457620004146200033d565b816040528281528960208487010111156200042d575f80fd5b6200044083602083016020880162000351565b80955050505050509250925092565b5f82516200046281846020870162000351565b9190910192915050565b6080516105ba620004845f395f601001526105ba5ff3fe608060405261000c61000e565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633036100a7575f357fffffffff00000000000000000000000000000000000000000000000000000000167f4f1ef286000000000000000000000000000000000000000000000000000000001461009f5761009d6100ab565b565b61009d6100bb565b61009d5b61009d6100b66100e9565b61012d565b5f806100ca3660048184610410565b8101906100d79190610464565b915091506100e5828261014b565b5050565b5f6101287f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b365f80375f80365f845af43d5f803e808015610147573d5ff35b3d5ffd5b610154826101b2565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101aa576101a58282610285565b505050565b6100e5610304565b8073ffffffffffffffffffffffffffffffffffffffff163b5f0361021f576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60605f808473ffffffffffffffffffffffffffffffffffffffff16846040516102ae9190610558565b5f60405180830381855af49150503d805f81146102e6576040519150601f19603f3d011682016040523d82523d5f602084013e6102eb565b606091505b50915091506102fb85838361033c565b95945050505050565b341561009d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826103515761034c826103ce565b6103c7565b8151158015610375575073ffffffffffffffffffffffffffffffffffffffff84163b155b156103c4576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610216565b50805b9392505050565b8051156103de5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f808585111561041e575f80fd5b8386111561042a575f80fd5b5050820193919092039150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610475575f80fd5b823573ffffffffffffffffffffffffffffffffffffffff81168114610498575f80fd5b9150602083013567ffffffffffffffff808211156104b4575f80fd5b818501915085601f8301126104c7575f80fd5b8135818111156104d9576104d9610437565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561051f5761051f610437565b81604052828152886020848701011115610537575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f82515f5b81811015610577576020818601810151858301520161055d565b505f92019182525091905056fea26469706673582212200ca61bd1e45d482203caba1d216b11bb6992f1ce0f6427bfe86e65b2f53457a264736f6c63430008140033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103a26469706673582212207bf6ae57411231dfe038cbed0bac99071e73b9b073610c2ef54cdedf17b8244064736f6c63430008140033",
}

// PolygonrollupmanagermockinternaltestABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagermockinternaltestMetaData.ABI instead.
var PolygonrollupmanagermockinternaltestABI = PolygonrollupmanagermockinternaltestMetaData.ABI

// PolygonrollupmanagermockinternaltestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagermockinternaltestMetaData.Bin instead.
var PolygonrollupmanagermockinternaltestBin = PolygonrollupmanagermockinternaltestMetaData.Bin

// DeployPolygonrollupmanagermockinternaltest deploys a new Ethereum contract, binding an instance of Polygonrollupmanagermockinternaltest to it.
func DeployPolygonrollupmanagermockinternaltest(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagermockinternaltest, error) {
	parsed, err := PolygonrollupmanagermockinternaltestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagermockinternaltestBin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanagermockinternaltest{PolygonrollupmanagermockinternaltestCaller: PolygonrollupmanagermockinternaltestCaller{contract: contract}, PolygonrollupmanagermockinternaltestTransactor: PolygonrollupmanagermockinternaltestTransactor{contract: contract}, PolygonrollupmanagermockinternaltestFilterer: PolygonrollupmanagermockinternaltestFilterer{contract: contract}}, nil
}

// Polygonrollupmanagermockinternaltest is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanagermockinternaltest struct {
	PolygonrollupmanagermockinternaltestCaller     // Read-only binding to the contract
	PolygonrollupmanagermockinternaltestTransactor // Write-only binding to the contract
	PolygonrollupmanagermockinternaltestFilterer   // Log filterer for contract events
}

// PolygonrollupmanagermockinternaltestCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockinternaltestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockinternaltestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockinternaltestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockinternaltestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupmanagermockinternaltestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockinternaltestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupmanagermockinternaltestSession struct {
	Contract     *Polygonrollupmanagermockinternaltest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PolygonrollupmanagermockinternaltestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupmanagermockinternaltestCallerSession struct {
	Contract *PolygonrollupmanagermockinternaltestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// PolygonrollupmanagermockinternaltestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupmanagermockinternaltestTransactorSession struct {
	Contract     *PolygonrollupmanagermockinternaltestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// PolygonrollupmanagermockinternaltestRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupmanagermockinternaltestRaw struct {
	Contract *Polygonrollupmanagermockinternaltest // Generic contract binding to access the raw methods on
}

// PolygonrollupmanagermockinternaltestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockinternaltestCallerRaw struct {
	Contract *PolygonrollupmanagermockinternaltestCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupmanagermockinternaltestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockinternaltestTransactorRaw struct {
	Contract *PolygonrollupmanagermockinternaltestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanagermockinternaltest creates a new instance of Polygonrollupmanagermockinternaltest, bound to a specific deployed contract.
func NewPolygonrollupmanagermockinternaltest(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanagermockinternaltest, error) {
	contract, err := bindPolygonrollupmanagermockinternaltest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagermockinternaltest{PolygonrollupmanagermockinternaltestCaller: PolygonrollupmanagermockinternaltestCaller{contract: contract}, PolygonrollupmanagermockinternaltestTransactor: PolygonrollupmanagermockinternaltestTransactor{contract: contract}, PolygonrollupmanagermockinternaltestFilterer: PolygonrollupmanagermockinternaltestFilterer{contract: contract}}, nil
}

// NewPolygonrollupmanagermockinternaltestCaller creates a new read-only instance of Polygonrollupmanagermockinternaltest, bound to a specific deployed contract.
func NewPolygonrollupmanagermockinternaltestCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupmanagermockinternaltestCaller, error) {
	contract, err := bindPolygonrollupmanagermockinternaltest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestCaller{contract: contract}, nil
}

// NewPolygonrollupmanagermockinternaltestTransactor creates a new write-only instance of Polygonrollupmanagermockinternaltest, bound to a specific deployed contract.
func NewPolygonrollupmanagermockinternaltestTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupmanagermockinternaltestTransactor, error) {
	contract, err := bindPolygonrollupmanagermockinternaltest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestTransactor{contract: contract}, nil
}

// NewPolygonrollupmanagermockinternaltestFilterer creates a new log filterer instance of Polygonrollupmanagermockinternaltest, bound to a specific deployed contract.
func NewPolygonrollupmanagermockinternaltestFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupmanagermockinternaltestFilterer, error) {
	contract, err := bindPolygonrollupmanagermockinternaltest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestFilterer{contract: contract}, nil
}

// bindPolygonrollupmanagermockinternaltest binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanagermockinternaltest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupmanagermockinternaltestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagermockinternaltest.Contract.PolygonrollupmanagermockinternaltestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.PolygonrollupmanagermockinternaltestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.PolygonrollupmanagermockinternaltestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagermockinternaltest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.BridgeAddress(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.BridgeAddress(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ChainIDToRollupID(&_Polygonrollupmanagermockinternaltest.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ChainIDToRollupID(&_Polygonrollupmanagermockinternaltest.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetForcedBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetForcedBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetInputSnarkBytes(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetInputSnarkBytes(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRoleAdmin(&_Polygonrollupmanagermockinternaltest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRoleAdmin(&_Polygonrollupmanagermockinternaltest.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupExitRoot(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupExitRoot(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetRollupPendingStateTransitions(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getRollupPendingStateTransitions", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesPendingState)).(*LegacyZKEVMStateVariablesPendingState)

	return out0, err

}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GlobalExitRootManager(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GlobalExitRootManager(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.HasRole(&_Polygonrollupmanagermockinternaltest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.HasRole(&_Polygonrollupmanagermockinternaltest.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.IsEmergencyState(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.IsEmergencyState(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) IsPendingStateConsolidable(opts *bind.CallOpts, rollupID uint32, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "isPendingStateConsolidable", rollupID, pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID, pendingStateNum)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.LastAggregationTimestamp(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.LastAggregationTimestamp(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.MultiplierBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.MultiplierBatchFee(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.PendingStateTimeout(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.PendingStateTimeout(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.Pol(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.Pol(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupAddressToID(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupAddressToID(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupCount(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupCount(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagermockinternaltest.Contract.RollupIDToRollupData(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagermockinternaltest.Contract.RollupIDToRollupData(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupTypeCount(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupTypeCount(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupTypeMap(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollupTypeMap(&_Polygonrollupmanagermockinternaltest.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TotalSequencedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TotalSequencedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TotalVerifiedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TotalVerifiedBatches(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermockinternaltest.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagermockinternaltest.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ActivateEmergencyState(&_Polygonrollupmanagermockinternaltest.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ActivateEmergencyState(&_Polygonrollupmanagermockinternaltest.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.AddExistingRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.AddExistingRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.AddNewRollupType(&_Polygonrollupmanagermockinternaltest.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.AddNewRollupType(&_Polygonrollupmanagermockinternaltest.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) ConsolidatePendingState(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "consolidatePendingState", rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ConsolidatePendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ConsolidatePendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.CreateNewRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.CreateNewRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.DeactivateEmergencyState(&_Polygonrollupmanagermockinternaltest.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.DeactivateEmergencyState(&_Polygonrollupmanagermockinternaltest.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GrantRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.GrantRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) Initialize(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "initialize", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.Initialize(&_Polygonrollupmanagermockinternaltest.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.Initialize(&_Polygonrollupmanagermockinternaltest.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ObsoleteRollupType(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ObsoleteRollupType(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.OnSequenceBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.OnSequenceBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) OverridePendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "overridePendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.OverridePendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.OverridePendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "proveNonDeterministicPendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RenounceRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RenounceRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RevokeRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RevokeRole(&_Polygonrollupmanagermockinternaltest.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 batchToRollback) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, batchToRollback uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "rollbackBatches", rollupContract, batchToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 batchToRollback) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) RollbackBatches(rollupContract common.Address, batchToRollback uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollbackBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, batchToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 batchToRollback) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) RollbackBatches(rollupContract common.Address, batchToRollback uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.RollbackBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, batchToRollback)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetBatchFee(&_Polygonrollupmanagermockinternaltest.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetBatchFee(&_Polygonrollupmanagermockinternaltest.TransactOpts, newBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagermockinternaltest.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagermockinternaltest.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetPendingStateTimeout(&_Polygonrollupmanagermockinternaltest.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetPendingStateTimeout(&_Polygonrollupmanagermockinternaltest.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagermockinternaltest.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagermockinternaltest.TransactOpts, newTrustedAggregatorTimeout)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagermockinternaltest.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagermockinternaltest.TransactOpts, newVerifyBatchTimeTarget)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.UpdateRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.UpdateRollup(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) VerifyBatches(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "verifyBatches", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatches(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermockinternaltest.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagermockinternaltest.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// PolygonrollupmanagermockinternaltestAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestAddExistingRollupIterator struct {
	Event *PolygonrollupmanagermockinternaltestAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestAddExistingRollup)
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
		it.Event = new(PolygonrollupmanagermockinternaltestAddExistingRollup)
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
func (it *PolygonrollupmanagermockinternaltestAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestAddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestAddExistingRollup struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestAddExistingRollupIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestAddExistingRollup)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseAddExistingRollup(log types.Log) (*PolygonrollupmanagermockinternaltestAddExistingRollup, error) {
	event := new(PolygonrollupmanagermockinternaltestAddExistingRollup)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator struct {
	Event *PolygonrollupmanagermockinternaltestAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestAddNewRollupType)
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
		it.Event = new(PolygonrollupmanagermockinternaltestAddNewRollupType)
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
func (it *PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestAddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestAddNewRollupType struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestAddNewRollupTypeIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestAddNewRollupType)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagermockinternaltestAddNewRollupType, error) {
	event := new(PolygonrollupmanagermockinternaltestAddNewRollupType)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator struct {
	Event *PolygonrollupmanagermockinternaltestConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestConsolidatePendingState)
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
		it.Event = new(PolygonrollupmanagermockinternaltestConsolidatePendingState)
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
func (it *PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestConsolidatePendingState struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestConsolidatePendingStateIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestConsolidatePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestConsolidatePendingState)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonrollupmanagermockinternaltestConsolidatePendingState, error) {
	event := new(PolygonrollupmanagermockinternaltestConsolidatePendingState)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestCreateNewRollupIterator struct {
	Event *PolygonrollupmanagermockinternaltestCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestCreateNewRollup)
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
		it.Event = new(PolygonrollupmanagermockinternaltestCreateNewRollup)
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
func (it *PolygonrollupmanagermockinternaltestCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestCreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestCreateNewRollup struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestCreateNewRollupIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestCreateNewRollup)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseCreateNewRollup(log types.Log) (*PolygonrollupmanagermockinternaltestCreateNewRollup, error) {
	event := new(PolygonrollupmanagermockinternaltestCreateNewRollup)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator struct {
	Event *PolygonrollupmanagermockinternaltestEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestEmergencyStateActivated)
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
		it.Event = new(PolygonrollupmanagermockinternaltestEmergencyStateActivated)
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
func (it *PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestEmergencyStateActivatedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestEmergencyStateActivated)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonrollupmanagermockinternaltestEmergencyStateActivated, error) {
	event := new(PolygonrollupmanagermockinternaltestEmergencyStateActivated)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator struct {
	Event *PolygonrollupmanagermockinternaltestEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestEmergencyStateDeactivated)
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
		it.Event = new(PolygonrollupmanagermockinternaltestEmergencyStateDeactivated)
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
func (it *PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestEmergencyStateDeactivatedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestEmergencyStateDeactivated)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonrollupmanagermockinternaltestEmergencyStateDeactivated, error) {
	event := new(PolygonrollupmanagermockinternaltestEmergencyStateDeactivated)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestInitializedIterator struct {
	Event *PolygonrollupmanagermockinternaltestInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestInitialized)
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
		it.Event = new(PolygonrollupmanagermockinternaltestInitialized)
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
func (it *PolygonrollupmanagermockinternaltestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestInitialized represents a Initialized event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestInitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestInitializedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestInitialized)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseInitialized(log types.Log) (*PolygonrollupmanagermockinternaltestInitialized, error) {
	event := new(PolygonrollupmanagermockinternaltestInitialized)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator struct {
	Event *PolygonrollupmanagermockinternaltestObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestObsoleteRollupType)
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
		it.Event = new(PolygonrollupmanagermockinternaltestObsoleteRollupType)
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
func (it *PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestObsoleteRollupTypeIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestObsoleteRollupType)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseObsoleteRollupType(log types.Log) (*PolygonrollupmanagermockinternaltestObsoleteRollupType, error) {
	event := new(PolygonrollupmanagermockinternaltestObsoleteRollupType)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator struct {
	Event *PolygonrollupmanagermockinternaltestOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestOnSequenceBatches)
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
		it.Event = new(PolygonrollupmanagermockinternaltestOnSequenceBatches)
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
func (it *PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestOnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestOnSequenceBatchesIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestOnSequenceBatches)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseOnSequenceBatches(log types.Log) (*PolygonrollupmanagermockinternaltestOnSequenceBatches, error) {
	event := new(PolygonrollupmanagermockinternaltestOnSequenceBatches)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestOverridePendingStateIterator struct {
	Event *PolygonrollupmanagermockinternaltestOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestOverridePendingState)
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
		it.Event = new(PolygonrollupmanagermockinternaltestOverridePendingState)
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
func (it *PolygonrollupmanagermockinternaltestOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestOverridePendingState represents a OverridePendingState event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestOverridePendingState struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterOverridePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestOverridePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestOverridePendingStateIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestOverridePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestOverridePendingState)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseOverridePendingState(log types.Log) (*PolygonrollupmanagermockinternaltestOverridePendingState, error) {
	event := new(PolygonrollupmanagermockinternaltestOverridePendingState)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator struct {
	Event *PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState)
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
		it.Event = new(PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState)
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
func (it *PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestProveNonDeterministicPendingStateIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState, error) {
	event := new(PolygonrollupmanagermockinternaltestProveNonDeterministicPendingState)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleAdminChangedIterator struct {
	Event *PolygonrollupmanagermockinternaltestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestRoleAdminChanged)
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
		it.Event = new(PolygonrollupmanagermockinternaltestRoleAdminChanged)
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
func (it *PolygonrollupmanagermockinternaltestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonrollupmanagermockinternaltestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestRoleAdminChangedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestRoleAdminChanged)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonrollupmanagermockinternaltestRoleAdminChanged, error) {
	event := new(PolygonrollupmanagermockinternaltestRoleAdminChanged)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleGrantedIterator struct {
	Event *PolygonrollupmanagermockinternaltestRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestRoleGranted)
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
		it.Event = new(PolygonrollupmanagermockinternaltestRoleGranted)
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
func (it *PolygonrollupmanagermockinternaltestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestRoleGranted represents a RoleGranted event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagermockinternaltestRoleGrantedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestRoleGrantedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestRoleGranted)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseRoleGranted(log types.Log) (*PolygonrollupmanagermockinternaltestRoleGranted, error) {
	event := new(PolygonrollupmanagermockinternaltestRoleGranted)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleRevokedIterator struct {
	Event *PolygonrollupmanagermockinternaltestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestRoleRevoked)
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
		it.Event = new(PolygonrollupmanagermockinternaltestRoleRevoked)
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
func (it *PolygonrollupmanagermockinternaltestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestRoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagermockinternaltestRoleRevokedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestRoleRevokedIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestRoleRevoked)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseRoleRevoked(log types.Log) (*PolygonrollupmanagermockinternaltestRoleRevoked, error) {
	event := new(PolygonrollupmanagermockinternaltestRoleRevoked)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRollbackBatchesIterator struct {
	Event *PolygonrollupmanagermockinternaltestRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestRollbackBatches)
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
		it.Event = new(PolygonrollupmanagermockinternaltestRollbackBatches)
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
func (it *PolygonrollupmanagermockinternaltestRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestRollbackBatches represents a RollbackBatches event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestRollbackBatches struct {
	RollupID               uint32
	BatchToRollback        uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed batchToRollback, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, batchToRollback []uint64) (*PolygonrollupmanagermockinternaltestRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var batchToRollbackRule []interface{}
	for _, batchToRollbackItem := range batchToRollback {
		batchToRollbackRule = append(batchToRollbackRule, batchToRollbackItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, batchToRollbackRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestRollbackBatchesIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed batchToRollback, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestRollbackBatches, rollupID []uint32, batchToRollback []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var batchToRollbackRule []interface{}
	for _, batchToRollbackItem := range batchToRollback {
		batchToRollbackRule = append(batchToRollbackRule, batchToRollbackItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, batchToRollbackRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestRollbackBatches)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed batchToRollback, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseRollbackBatches(log types.Log) (*PolygonrollupmanagermockinternaltestRollbackBatches, error) {
	event := new(PolygonrollupmanagermockinternaltestRollbackBatches)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetBatchFeeIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetBatchFee)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetBatchFee)
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
func (it *PolygonrollupmanagermockinternaltestSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetBatchFeeIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetBatchFee)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetBatchFee(log types.Log) (*PolygonrollupmanagermockinternaltestSetBatchFee, error) {
	event := new(PolygonrollupmanagermockinternaltestSetBatchFee)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetMultiplierBatchFee)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetMultiplierBatchFee)
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
func (it *PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetMultiplierBatchFeeIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetMultiplierBatchFee)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonrollupmanagermockinternaltestSetMultiplierBatchFee, error) {
	event := new(PolygonrollupmanagermockinternaltestSetMultiplierBatchFee)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetPendingStateTimeout)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetPendingStateTimeout)
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
func (it *PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetPendingStateTimeoutIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetPendingStateTimeout)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonrollupmanagermockinternaltestSetPendingStateTimeout, error) {
	event := new(PolygonrollupmanagermockinternaltestSetPendingStateTimeout)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetTrustedAggregator)
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
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetTrustedAggregatorIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetTrustedAggregator)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonrollupmanagermockinternaltestSetTrustedAggregator, error) {
	event := new(PolygonrollupmanagermockinternaltestSetTrustedAggregator)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout)
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
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeoutIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout, error) {
	event := new(PolygonrollupmanagermockinternaltestSetTrustedAggregatorTimeout)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget)
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
func (it *PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTargetIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget, error) {
	event := new(PolygonrollupmanagermockinternaltestSetVerifyBatchTimeTarget)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestUpdateRollupIterator struct {
	Event *PolygonrollupmanagermockinternaltestUpdateRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestUpdateRollup)
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
		it.Event = new(PolygonrollupmanagermockinternaltestUpdateRollup)
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
func (it *PolygonrollupmanagermockinternaltestUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestUpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockinternaltestUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestUpdateRollupIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestUpdateRollup)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseUpdateRollup(log types.Log) (*PolygonrollupmanagermockinternaltestUpdateRollup, error) {
	event := new(PolygonrollupmanagermockinternaltestUpdateRollup)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestVerifyBatchesIterator struct {
	Event *PolygonrollupmanagermockinternaltestVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestVerifyBatches)
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
		it.Event = new(PolygonrollupmanagermockinternaltestVerifyBatches)
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
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestVerifyBatches represents a VerifyBatches event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestVerifyBatches struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterVerifyBatches(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagermockinternaltestVerifyBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestVerifyBatchesIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestVerifyBatches, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestVerifyBatches)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupmanagermockinternaltestVerifyBatches, error) {
	event := new(PolygonrollupmanagermockinternaltestVerifyBatches)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator)
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
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanagermockinternaltest contract.
type PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator struct {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanagermockinternaltest.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermockinternaltest.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagermockinternaltest *PolygonrollupmanagermockinternaltestFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonrollupmanagermockinternaltestVerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanagermockinternaltest.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

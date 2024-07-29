// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanagermock

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

// PolygonrollupmanagermockMetaData contains all meta data concerning the Polygonrollupmanagermock contract.
var PolygonrollupmanagermockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"ConsolidatePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"OverridePendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storedStateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"provedStateRoot\",\"type\":\"bytes32\"}],\"name\":\"ProveNonDeterministicPendingState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"SetMultiplierBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"SetPendingStateTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"SetTrustedAggregatorTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"SetVerifyBatchTimeTarget\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"consolidatePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupPendingStateTransitions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structLegacyZKEVMStateVariables.PendingState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"},{\"internalType\":\"contractPolygonZkEVMExistentEtrog\",\"name\":\"polygonZkEVM\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"zkEVMVerifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMForkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"zkEVMChainID\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"}],\"name\":\"initializeMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"}],\"name\":\"isPendingStateConsolidable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiplierBatchFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"overridePendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStateTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"localExitRoots\",\"type\":\"bytes32[]\"}],\"name\":\"prepareMockCalculateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalPendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"proveNonDeterministicPendingState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"rollupCompatibilityID\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMultiplierBatchFee\",\"type\":\"uint16\"}],\"name\":\"setMultiplierBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newPendingStateTimeout\",\"type\":\"uint64\"}],\"name\":\"setPendingStateTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newTrustedAggregatorTimeout\",\"type\":\"uint64\"}],\"name\":\"setTrustedAggregatorTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newVerifyBatchTimeTarget\",\"type\":\"uint64\"}],\"name\":\"setVerifyBatchTimeTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedAggregatorTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyBatchTimeTarget\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200635b3803806200635b833981016040819052620000349162000141565b6001600160a01b0380841660805280831660c052811660a0528282826200005a62000066565b50505050505062000195565b600054610100900460ff1615620000d35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000126576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013e57600080fd5b50565b6000806000606084860312156200015757600080fd5b8351620001648162000128565b6020850151909350620001778162000128565b60408501519092506200018a8162000128565b809150509250925092565b60805160a05160c05161615e620001fd600039600081816109dc0152818161231c0152613bfb0152600081816107a201528181612e960152613ef50152600081816109360152818161132f015281816114df01528181611ff90152613de4015261615e6000f3fe60806040523480156200001157600080fd5b5060043610620002b65760003560e01c80630645af0914620002bb578063066ec01214620002d4578063080b311114620003005780630a0d9fbe14620003285780630e36f582146200034357806311f6b287146200035a57806312b86e1914620003715780631489ed10146200038857806315064c96146200039f5780631608859c14620003ad5780631796a1ae14620003c45780631816b7e514620003eb5780632072f6c51462000402578063248a9ca3146200040c5780632528016914620004325780632f2ff15d14620004e757806330c27dde14620004fe57806336568abe1462000512578063394218e91462000529578063477fa270146200054057806355a71ee0146200054957806360469169146200058d57806365c0504d14620005975780637222020f1462000646578063727885e9146200065d5780637975fcfe14620006745780637fb6e76a146200069a578063841b24d714620006c357806387c20c0114620006de5780638bd4f07114620006f55780638f698ec5146200070c57806391d14854146200072357806399f5634e146200073a5780639a908e7314620007445780639c9f3dfe146200075b578063a066215c1462000772578063a217fddf1462000789578063a2967d991462000792578063a3c573eb146200079c578063afd23cbe14620007d3578063b99d0ad714620007fd578063c1acbc3414620008d5578063c4c928c214620008f0578063ceee281d1462000907578063d02103ca1462000930578063d5073f6f1462000958578063d547741f146200096f578063d939b3151462000986578063dbc16976146200099a578063dde0ff7714620009a4578063e0bfd3d214620009bf578063e46761c414620009d6578063f34eb8eb14620009fe578063f4e926751462000a15578063f9c4c2ae1462000a26575b600080fd5b620002d2620002cc366004620048ac565b62000b3d565b005b608454620002e8906001600160401b031681565b604051620002f7919062004987565b60405180910390f35b6200031762000311366004620049b0565b62000f8b565b6040519015158152602001620002f7565b608554620002e890600160401b90046001600160401b031681565b620002d262000354366004620049e8565b62000fb5565b620002e86200036b36600462004a73565b62001263565b620002d26200038236600462004aa4565b62001283565b620002d26200039936600462004b3b565b62001433565b606f54620003179060ff1681565b620002d2620003be366004620049b0565b620015c3565b607e54620003d59063ffffffff1681565b60405163ffffffff9091168152602001620002f7565b620002d2620003fc36600462004bc5565b62001658565b620002d262001704565b620004236200041d36600462004bf2565b620017ca565b604051908152602001620002f7565b620004b362000443366004620049b0565b60408051606080820183526000808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b03908116918301919091529282015190921690820152606001620002f7565b620002d2620004f836600462004c0c565b620017df565b608754620002e8906001600160401b031681565b620002d26200052336600462004c0c565b62001801565b620002d26200053a36600462004c3f565b6200183b565b60865462000423565b620004236200055a366004620049b0565b63ffffffff821660009081526081602090815260408083206001600160401b038516845260020190915290205492915050565b62000423620018ea565b620005fc620005a836600462004a73565b607f602052600090815260409020805460018201546002909201546001600160a01b0391821692918216916001600160401b03600160a01b8204169160ff600160e01b8304811692600160e81b9004169086565b604080516001600160a01b0397881681529690951660208701526001600160401b039093169385019390935260ff166060840152901515608083015260a082015260c001620002f7565b620002d26200065736600462004a73565b62001902565b620002d26200066e36600462004d1c565b620019ed565b6200068b6200068536600462004de9565b62001e55565b604051620002f7919062004ea3565b620003d5620006ab36600462004c3f565b60836020526000908152604090205463ffffffff1681565b608454620002e890600160c01b90046001600160401b031681565b620002d2620006ef36600462004b3b565b62001e88565b620002d26200070636600462004aa4565b620021ac565b620002d26200071d36600462004eb8565b62002265565b620003176200073436600462004c0c565b620022ec565b6200042362002317565b620002e86200075536600462004f68565b62002403565b620002d26200076c36600462004c3f565b620025d0565b620002d26200078336600462004c3f565b62002673565b62000423600081565b6200042362002712565b620007c47f000000000000000000000000000000000000000000000000000000000000000081565b604051620002f7919062004f95565b608554620007e990600160801b900461ffff1681565b60405161ffff9091168152602001620002f7565b620008936200080e366004620049b0565b604080516080808201835260008083526020808401829052838501829052606093840182905263ffffffff969096168152608186528381206001600160401b03958616825260040186528390208351918201845280548086168352600160401b9004909416948101949094526001830154918401919091526002909101549082015290565b604051620002f7919081516001600160401b03908116825260208084015190911690820152604082810151908201526060918201519181019190915260800190565b608454620002e890600160801b90046001600160401b031681565b620002d26200090136600462004fa9565b62002ad4565b620003d56200091836600462005041565b60826020526000908152604090205463ffffffff1681565b620007c47f000000000000000000000000000000000000000000000000000000000000000081565b620002d26200096936600462004bf2565b62002da1565b620002d26200098036600462004c0c565b62002e2c565b608554620002e8906001600160401b031681565b620002d262002e4e565b608454620002e890600160401b90046001600160401b031681565b620002d2620009d036600462005073565b62002f0c565b620007c47f000000000000000000000000000000000000000000000000000000000000000081565b620002d262000a0f366004620050ef565b62002fd4565b608054620003d59063ffffffff1681565b62000abd62000a3736600462004a73565b608160205260009081526040902080546001820154600583015460068401546007909401546001600160a01b0380851695600160a01b958690046001600160401b039081169692861695929092048216939282821692600160401b808404821693600160801b808204841694600160c01b90920484169380831693830416910460ff168c565b604080516001600160a01b039d8e1681526001600160401b039c8d1660208201529c909a16998c019990995296891660608b015260808a019590955292871660a089015290861660c0880152851660e0870152841661010086015283166101208501529190911661014083015260ff1661016082015261018001620002f7565b600054600290610100900460ff1615801562000b60575060005460ff8083169116105b62000b885760405162461bcd60e51b815260040162000b7f9062005186565b60405180910390fd5b6000805461010060ff841661ffff199092169190911717905560858054608480546001600160c01b0316600160c01b6001600160401b038e8116919091029190911790915567016345785d8a00006086558c166001600160801b03199091161760e160431b1761ffff60801b19166101f560811b17905562000c09620031bf565b62000c24600080516020620061098339815191528c6200322c565b62000c316000886200322c565b62000c4c60008051602062005fe9833981519152886200322c565b62000c6760008051602062006069833981519152886200322c565b62000c8260008051602062005f89833981519152886200322c565b62000c9d60008051602062005fc9833981519152896200322c565b62000cb8600080516020620060e9833981519152896200322c565b62000cd360008051602062006009833981519152896200322c565b62000cee60008051602062006089833981519152896200322c565b62000d186000805160206200610983398151915260008051602062005f6983398151915262003238565b62000d3360008051602062005f69833981519152896200322c565b62000d4e60008051602062005fa9833981519152896200322c565b62000d78600080516020620060c9833981519152600080516020620060a983398151915262003238565b62000d93600080516020620060c9833981519152876200322c565b62000dae600080516020620060a9833981519152876200322c565b6073546074546001600160401b03600160401b9092048216911680821462000de957604051632e4cc54360e11b815260040160405180910390fd5b600062000e11888888886000607460009054906101000a90046001600160401b03166200328d565b6001600160401b03838116600081815260756020908152604080832054600287018352818420558885168084526072808452828520600389018552948390208554815560018087018054919092018054918a166001600160401b03198084168217835593546001600160801b0319938416909117600160401b91829004909b1681029a909a17905560068a01805490911690931797870297909717909155600787018054909616909417909455607a54606f549390915290549251635d6717a560e01b81529394506001600160a01b038c811694635d6717a59462000f0f9493831693600160581b9004909216916076916077919060040162005280565b600060405180830381600087803b15801562000f2a57600080fd5b505af115801562000f3f573d6000803e3d6000fd5b50506000805461ff0019169055505060405160ff85168152600080516020620060498339815191529350602001915062000f769050565b60405180910390a15050505050505050505050565b63ffffffff8216600090815260816020526040812062000fac9083620034bb565b90505b92915050565b600054600290610100900460ff1615801562000fd8575060005460ff8083169116105b62000ff75760405162461bcd60e51b815260040162000b7f9062005186565b6000805461010060ff841661ffff199092169190911717905560858054608480546001600160c01b0316600160c01b6001600160401b038a8116919091029190911790915567016345785d8a000060865588166001600160801b03199091161760e160431b1761ffff60801b19166101f560811b17905562001078620031bf565b6200109360008051602062006109833981519152886200322c565b620010a06000846200322c565b620010bb60008051602062005fe9833981519152846200322c565b620010d660008051602062006069833981519152846200322c565b620010f160008051602062005f89833981519152846200322c565b6200110c60008051602062005fc9833981519152856200322c565b62001127600080516020620060e9833981519152856200322c565b6200114260008051602062006009833981519152856200322c565b6200115d60008051602062006089833981519152856200322c565b620011876000805160206200610983398151915260008051602062005f6983398151915262003238565b620011a260008051602062005f69833981519152856200322c565b620011bd60008051602062005fa9833981519152856200322c565b620011e7600080516020620060c9833981519152600080516020620060a983398151915262003238565b62001202600080516020620060c9833981519152836200322c565b6200121d600080516020620060a9833981519152836200322c565b6200122a6000336200322c565b6000805461ff001916905560405160ff82168152600080516020620060498339815191529060200160405180910390a150505050505050565b63ffffffff8116600090815260816020526040812062000faf9062003500565b600080516020620061098339815191526200129e8162003571565b63ffffffff89166000908152608160205260409020620012c5818a8a8a8a8a8a8a6200357d565b600681018054600160401b600160801b031916600160401b6001600160401b0389811691820292909217835560009081526002840160205260409020869055600583018790559054600160801b900416156200132d576006810180546001600160801b031690555b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d6200136662002712565b6040518263ffffffff1660e01b81526004016200138591815260200190565b600060405180830381600087803b158015620013a057600080fd5b505af1158015620013b5573d6000803e3d6000fd5b5050608480546001600160c01b031661127560c71b1790555050604080516001600160401b03881681526020810186905290810186905233606082015263ffffffff8b16907f3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e9060800160405180910390a250505050505050505050565b600080516020620061098339815191526200144e8162003571565b63ffffffff8916600090815260816020526040902062001475818a8a8a8a8a8a8a62003905565b600681018054600160401b600160801b031916600160401b6001600160401b038a811691820292909217835560009081526002840160205260409020879055600583018890559054600160801b90041615620014dd576006810180546001600160801b031690555b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d6200151662002712565b6040518263ffffffff1660e01b81526004016200153591815260200190565b600060405180830381600087803b1580156200155057600080fd5b505af115801562001565573d6000803e3d6000fd5b50505050336001600160a01b03168a63ffffffff167fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d389888a604051620015af93929190620052d5565b60405180910390a350505050505050505050565b63ffffffff82166000908152608160205260409020620015f36000805160206200610983398151915233620022ec565b6200164757606f5460ff16156200161d57604051630bc011ff60e21b815260040160405180910390fd5b620016298183620034bb565b6200164757604051630674f25160e11b815260040160405180910390fd5b62001653818362003d01565b505050565b60008051602062006089833981519152620016738162003571565b6103e88261ffff1610806200168d57506103ff8261ffff16115b15620016ac57604051630984a67960e31b815260040160405180910390fd5b6085805461ffff60801b1916600160801b61ffff8516908102919091179091556040519081527f7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5906020015b60405180910390a15050565b6200171f600080516020620060c983398151915233620022ec565b620017be57608454600160801b90046001600160401b031615806200176f57506084544290620017649062093a8090600160801b90046001600160401b03166200530c565b6001600160401b0316115b806200179f57506087544290620017949062093a80906001600160401b03166200530c565b6001600160401b0316115b15620017be5760405163692baaad60e11b815260040160405180910390fd5b620017c862003ef3565b565b60009081526034602052604090206001015490565b620017ea82620017ca565b620017f58162003571565b62001653838362003f72565b6001600160a01b03811633146200182b57604051630b4ad1cd60e31b815260040160405180910390fd5b62001837828262003fde565b5050565b60008051602062006089833981519152620018568162003571565b606f5460ff1662001898576084546001600160401b03600160c01b909104811690831610620018985760405163401636df60e01b815260040160405180910390fd5b608480546001600160c01b0316600160c01b6001600160401b038516021790556040517f1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a190620016f890849062004987565b60006086546064620018fd919062005336565b905090565b60008051602062005fc98339815191526200191d8162003571565b63ffffffff821615806200193c5750607e5463ffffffff908116908316115b156200195b57604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82166000908152607f60205260409020600180820154600160e81b900460ff1615159003620019a257604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e4490600090a2505050565b600080516020620060e983398151915262001a088162003571565b63ffffffff8816158062001a275750607e5463ffffffff908116908916115b1562001a4657604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff88166000908152607f60205260409020600180820154600160e81b900460ff161515900362001a8d57604051633b8d3d9960e01b815260040160405180910390fd5b6001600160401b03881660009081526083602052604090205463ffffffff161562001acb576040516337c8fe0960e11b815260040160405180910390fd5b6080805460009190829062001ae69063ffffffff1662005350565b825463ffffffff8281166101009490940a9384029302191691909117909155825460408051600080825260208201928390529394506001600160a01b0390921691309162001b349062004870565b62001b429392919062005376565b604051809103906000f08015801562001b5f573d6000803e3d6000fd5b50905081608360008c6001600160401b03166001600160401b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055508160826000836001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055506000608160008463ffffffff1663ffffffff1681526020019081526020016000209050818160000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508360010160149054906101000a90046001600160401b03168160010160146101000a8154816001600160401b0302191690836001600160401b031602179055508360010160009054906101000a90046001600160a01b03168160010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508a8160000160146101000a8154816001600160401b0302191690836001600160401b031602179055508360020154816002016000806001600160401b03168152602001908152602001600020819055508b63ffffffff168160070160086101000a8154816001600160401b0302191690836001600160401b0316021790555083600101601c9054906101000a900460ff168160070160106101000a81548160ff021916908360ff1602179055508263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418d848e8c60405162001dd3949392919063ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b60405180910390a2604051633892b81160e11b81526001600160a01b0383169063712570229062001e13908d908d9088908e908e908e90600401620053ad565b600060405180830381600087803b15801562001e2e57600080fd5b505af115801562001e43573d6000803e3d6000fd5b50505050505050505050505050505050565b63ffffffff8616600090815260816020526040902060609062001e7d90878787878762004048565b979650505050505050565b606f5460ff161562001ead57604051630bc011ff60e21b815260040160405180910390fd5b63ffffffff881660009081526081602090815260408083206084546001600160401b038a81168652600383019094529190932060010154429262001efc92600160c01b9004811691166200530c565b6001600160401b0316111562001f2557604051638a0704d360e01b815260040160405180910390fd5b6103e862001f34888862005410565b6001600160401b0316111562001f5d57604051635acfba9d60e11b815260040160405180910390fd5b62001f6f818989898989898962003905565b62001f7b818762004183565b6085546001600160401b03166000036200208957600681018054600160401b600160801b031916600160401b6001600160401b0389811691820292909217835560009081526002840160205260409020869055600583018790559054600160801b9004161562001ff7576006810180546001600160801b031690555b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d6200203062002712565b6040518263ffffffff1660e01b81526004016200204f91815260200190565b600060405180830381600087803b1580156200206a57600080fd5b505af11580156200207f573d6000803e3d6000fd5b5050505062002153565b620020948162004380565b600681018054600160801b90046001600160401b0316906010620020b88362005433565b82546001600160401b039182166101009390930a92830292820219169190911790915560408051608081018252428316815289831660208083019182528284018b8152606084018b81526006890154600160801b90048716600090815260048a01909352949091209251835492518616600160401b026001600160801b03199093169516949094171781559151600183015551600290910155505b336001600160a01b03168963ffffffff167faac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b48887896040516200219993929190620052d5565b60405180910390a3505050505050505050565b606f5460ff1615620021d157604051630bc011ff60e21b815260040160405180910390fd5b63ffffffff88166000908152608160205260409020620021f881898989898989896200357d565b6001600160401b03871660009081526004820160209081526040918290206002015482519081529081018590527f1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010910160405180910390a16200225a62003ef3565b505050505050505050565b80516080805463ffffffff191663ffffffff90921691909117905560005b81518110156200183757818181518110620022a257620022a26200545a565b602002602001015160816000836001620022bd919062005470565b63ffffffff16815260208101919091526040016000206005015580620022e38162005486565b91505062002283565b60009182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040162002368919062004f95565b602060405180830381865afa15801562002386573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620023ac9190620054a2565b608454909150600090620023d3906001600160401b03600160401b82048116911662005410565b6001600160401b0316905080600003620023f05760009250505090565b620023fc8183620054d2565b9250505090565b606f5460009060ff16156200242b57604051630bc011ff60e21b815260040160405180910390fd5b3360009081526082602052604081205463ffffffff169081900362002463576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03166000036200248e57604051632590ccf960e01b815260040160405180910390fd5b63ffffffff811660009081526081602052604081206084805491928792620024c19084906001600160401b03166200530c565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690506000620024f787836200530c565b6006840180546001600160401b038084166001600160401b03199092168217909255604080516060810182528a81524284166020808301918252888616838501908152600095865260038b0190915292909320905181559151600192909201805491518416600160401b026001600160801b031990921692909316919091171790559050620025868362004380565b8363ffffffff167f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a2582604051620025be919062004987565b60405180910390a29695505050505050565b60008051602062006089833981519152620025eb8162003571565b606f5460ff1662002626576085546001600160401b0390811690831610620026265760405163048a05a960e41b815260040160405180910390fd5b608580546001600160401b0319166001600160401b0384161790556040517fc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c7590620016f890849062004987565b600080516020620060898339815191526200268e8162003571565b62015180826001600160401b03161115620026bc57604051631c0cfbfd60e31b815260040160405180910390fd5b60858054600160401b600160801b031916600160401b6001600160401b038516021790556040517f1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c2890620016f890849062004987565b60805460009063ffffffff168082036200272e57506000919050565b6000816001600160401b038111156200274b576200274b62004c5d565b60405190808252806020026020018201604052801562002775578160200160208202803683370190505b50905060005b82811015620027e857608160006200279583600162005470565b63ffffffff1663ffffffff16815260200190815260200160002060050154828281518110620027c857620027c86200545a565b602090810291909101015280620027df8162005486565b9150506200277b565b50600060205b8360011462002a2c57600062002806600286620054e9565b62002813600287620054d2565b6200281f919062005470565b90506000816001600160401b038111156200283e576200283e62004c5d565b60405190808252806020026020018201604052801562002868578160200160208202803683370190505b50905060005b82811015620029e0576200288460018462005500565b811480156200289f57506200289b600288620054e9565b6001145b156200291f5785620028b382600262005336565b81518110620028c657620028c66200545a565b602002602001015185604051602001620028e292919062005516565b604051602081830303815290604052805190602001208282815181106200290d576200290d6200545a565b602002602001018181525050620029cb565b856200292d82600262005336565b815181106200294057620029406200545a565b60200260200101518682600262002958919062005336565b6200296590600162005470565b815181106200297857620029786200545a565b60200260200101516040516020016200299392919062005516565b60405160208183030381529060405280519060200120828281518110620029be57620029be6200545a565b6020026020010181815250505b80620029d78162005486565b9150506200286e565b508094508195508384604051602001620029fc92919062005516565b604051602081830303815290604052805190602001209350828062002a219062005524565b9350505050620027ee565b60008360008151811062002a445762002a446200545a565b6020026020010151905060005b8281101562002aca57818460405160200162002a6f92919062005516565b604051602081830303815290604052805190602001209150838460405160200162002a9c92919062005516565b604051602081830303815290604052805190602001209350808062002ac19062005486565b91505062002a51565b5095945050505050565b60008051602062005f8983398151915262002aef8162003571565b63ffffffff8416158062002b0e5750607e5463ffffffff908116908516115b1562002b2d57604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b03851660009081526082602052604081205463ffffffff169081900362002b6e576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181166000908152608160205260409020600781015490918716600160401b9091046001600160401b03160362002bbd57604051634f61d51960e01b815260040160405180910390fd5b63ffffffff86166000908152607f60205260409020600180820154600160e81b900460ff161515900362002c0457604051633b8d3d9960e01b815260040160405180910390fd5b60018101546007830154600160801b900460ff908116600160e01b909204161462002c4257604051635aa0d5f160e11b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825591546001600160401b03600160a01b9182900416026001600160e01b0319909216909217179055600782018054600160401b63ffffffff8a1602600160401b600160801b0319909116179055600062002cc28462001263565b6007840180546001600160401b0319166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b038b811692634f1ef2869262002d179216908b908b906004016200553e565b600060405180830381600087803b15801562002d3257600080fd5b505af115801562002d47573d6000803e3d6000fd5b50506040805163ffffffff8c811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a2505050505050505050565b60008051602062005fa983398151915262002dbc8162003571565b683635c9adc5dea0000082118062002dd75750633b9aca0082105b1562002df657604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b290602001620016f8565b62002e3782620017ca565b62002e428162003571565b62001653838362003fde565b6000805160206200600983398151915262002e698162003571565b608780546001600160401b031916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc1697691600480830192600092919082900301818387803b15801562002ee657600080fd5b505af115801562002efb573d6000803e3d6000fd5b5050505062002f096200444b565b50565b6000805160206200606983398151915262002f278162003571565b6001600160401b03841660009081526083602052604090205463ffffffff161562002f65576040516337c8fe0960e11b815260040160405180910390fd5b6001600160a01b03871660009081526082602052604090205463ffffffff161562002fa357604051630d409b9360e41b815260040160405180910390fd5b600062002fb6888888888760006200328d565b60008080526002909101602052604090209390935550505050505050565b60008051602062005fe983398151915262002fef8162003571565b607e80546000919082906200300a9063ffffffff1662005350565b91906101000a81548163ffffffff021916908363ffffffff160217905590506040518060c00160405280896001600160a01b03168152602001886001600160a01b03168152602001876001600160401b031681526020018660ff16815260200160001515815260200185815250607f60008363ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060408201518160010160146101000a8154816001600160401b0302191690836001600160401b03160217905550606082015181600101601c6101000a81548160ff021916908360ff160217905550608082015181600101601d6101000a81548160ff02191690831515021790555060a082015181600201559050508063ffffffff167fa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52898989898989604051620031ad969594939291906200557e565b60405180910390a25050505050505050565b600054610100900460ff16620017c85760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840162000b7f565b62001837828262003f72565b60006200324583620017ca565b600084815260346020526040808220600101859055519192508391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6080805460009182918290620032a99063ffffffff1662005350565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060836000876001600160401b03166001600160401b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555080608260008a6001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff160217905550608160008263ffffffff1663ffffffff1681526020019081526020016000209150878260000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550858260010160146101000a8154816001600160401b0302191690836001600160401b03160217905550868260010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550848260000160146101000a8154816001600160401b0302191690836001600160401b03160217905550838260070160106101000a81548160ff021916908360ff1602179055508063ffffffff167fadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850878a888888604051620034a89594939291906001600160401b0395861681526001600160a01b03949094166020850152918416604084015260ff166060830152909116608082015260a00190565b60405180910390a2509695505050505050565b6085546001600160401b03828116600090815260048501602052604081205490924292620034ee9291811691166200530c565b6001600160401b031611159392505050565b6006810154600090600160801b90046001600160401b03161562003554575060068101546001600160401b03600160801b909104811660009081526004909201602052604090912054600160401b90041690565b5060060154600160401b90046001600160401b031690565b919050565b62002f098133620044a4565b60078801546000906001600160401b039081169087161015620035b35760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b03881615620036545760068901546001600160401b03600160801b90910481169089161115620035fd5760405163bb14c20560e01b815260040160405180910390fd5b506001600160401b03808816600090815260048a0160205260409020600281015481549092888116600160401b90920416146200364d57604051632bd2e3e760e01b815260040160405180910390fd5b50620036c9565b506001600160401b0385166000908152600289016020526040902054806200368f576040516324cbdcc360e11b815260040160405180910390fd5b60068901546001600160401b03600160401b90910481169087161115620036c957604051630f2b74f160e11b815260040160405180910390fd5b60068901546001600160401b03600160801b90910481169088161180620037025750876001600160401b0316876001600160401b031611155b8062003726575060068901546001600160401b03600160c01b909104811690881611155b15620037455760405163bfa7079f60e01b815260040160405180910390fd5b6001600160401b03878116600090815260048b016020526040902054600160401b90048116908616146200378c576040516332a2a77f60e01b815260040160405180910390fd5b60006200379e8a888888868962004048565b9050600060008051602062006029833981519152600283604051620037c49190620055cb565b602060405180830381855afa158015620037e2573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190620038079190620054a2565b620038139190620054e9565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a916200385791889190600401620055e9565b602060405180830381865afa15801562003875573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200389b919062005626565b620038b9576040516309bde33960e01b815260040160405180910390fd5b6001600160401b038916600090815260048c016020526040902060020154859003620038f85760405163a47276bd60e01b815260040160405180910390fd5b5050505050505050505050565b600080620039138a62003500565b60078b01549091506001600160401b039081169089161015620039495760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b03891615620039ec5760068a01546001600160401b03600160801b9091048116908a161115620039935760405163bb14c20560e01b815260040160405180910390fd5b6001600160401b03808a16600090815260048c01602052604090206002810154815490945090918a8116600160401b9092041614620039e557604051632bd2e3e760e01b815260040160405180910390fd5b5062003a5c565b6001600160401b038816600090815260028b01602052604090205491508162003a28576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b0316111562003a5c57604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b03161162003a8f5760405163b9b18f5760e01b815260040160405180910390fd5b600062003aa18b8a8a8a878b62004048565b905060006000805160206200602983398151915260028360405162003ac79190620055cb565b602060405180830381855afa15801562003ae5573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019062003b0a9190620054a2565b62003b169190620054e9565b60018d0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a9162003b5a91899190600401620055e9565b602060405180830381865afa15801562003b78573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062003b9e919062005626565b62003bbc576040516309bde33960e01b815260040160405180910390fd5b600062003bca848b62005410565b905062003c2387826001600160401b031662003be562002317565b62003bf1919062005336565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190620044ce565b80608460088282829054906101000a90046001600160401b031662003c4991906200530c565b82546101009290920a6001600160401b0381810219909316918316021790915560848054600160801b600160c01b031916600160801b428416021790558e546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d15390606401600060405180830381600087803b15801562003cd957600080fd5b505af115801562003cee573d6000803e3d6000fd5b5050505050505050505050505050505050565b60068201546001600160401b03600160c01b909104811690821611158062003d40575060068201546001600160401b03600160801b9091048116908216115b1562003d5f5760405163d086b70b60e01b815260040160405180910390fd5b6001600160401b03818116600081815260048501602090815260408083208054600689018054600160401b600160801b031916600160401b92839004909816918202979097178755600280830154828752908a0190945291909320919091556001820154600587015583546001600160c01b0316600160c01b909302929092179092557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d62003e1b62002712565b6040518263ffffffff1660e01b815260040162003e3a91815260200190565b600060405180830381600087803b15801562003e5557600080fd5b505af115801562003e6a573d6000803e3d6000fd5b505085546001600160a01b0316600090815260826020908152604091829020546002870154600188015484516001600160401b03898116825294810192909252818501529188166060830152915163ffffffff90921693507f581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b925081900360800190a250505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b8152600401600060405180830381600087803b15801562003f4f57600080fd5b505af115801562003f64573d6000803e3d6000fd5b50505050620017c862004522565b62003f7e8282620022ec565b620018375760008281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b62003fea8282620022ec565b15620018375760008281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6001600160401b03808616600081815260038901602052604080822054938816825290205460609291158015906200407e575081155b156200409d5760405163340c614f60e11b815260040160405180910390fd5b80620040bc576040516366385b5160e01b815260040160405180910390fd5b620040c7846200457f565b620040e5576040516305dae44f60e21b815260040160405180910390fd5b885460018a01546040516001600160601b03193360601b16602082015260348101889052605481018590526001600160c01b031960c08c811b82166074840152600160a01b94859004811b8216607c84015293909204831b82166084820152608c810187905260ac810184905260cc81018990529189901b1660ec82015260f401604051602081830303815290604052925050509695505050505050565b6000620041908362003500565b905081600080620041a2848462005410565b6085546001600160401b039182169250600091620041c991600160401b9004164262005500565b90505b846001600160401b0316846001600160401b03161462004253576001600160401b038085166000908152600389016020526040902060018101549091168210156200422e576001810154600160401b90046001600160401b031694506200424c565b6200423a868662005410565b6001600160401b031693505062004253565b50620041cc565b600062004261848462005500565b905083811015620042bf57808403600c81116200427f578062004282565b600c5b9050806103e80a81608560109054906101000a900461ffff1661ffff160a6086540281620042b457620042b4620054bc565b046086555062004337565b838103600c8111620042d25780620042d5565b600c5b90506000816103e80a82608560109054906101000a900461ffff1661ffff160a670de0b6b3a764000002816200430f576200430f620054bc565b04905080608654670de0b6b3a76400000281620043305762004330620054bc565b0460865550505b683635c9adc5dea0000060865411156200435e57683635c9adc5dea0000060865562004376565b633b9aca0060865410156200437657633b9aca006086555b5050505050505050565b60068101546001600160401b03600160c01b82048116600160801b90920416111562002f09576006810154600090620043cb90600160c01b90046001600160401b031660016200530c565b9050620043d98282620034bb565b156200183757600682015460009060029062004407908490600160801b90046001600160401b031662005410565b6200441391906200564a565b6200441f90836200530c565b90506200442d8382620034bb565b156200443f5762001653838262003d01565b62001653838362003d01565b606f5460ff166200446f57604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b620044b08282620022ec565b6200183757604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526200165390849062004605565b606f5460ff16156200454757604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b600067ffffffff000000016001600160401b038316108015620045b6575067ffffffff00000001604083901c6001600160401b0316105b8015620045d7575067ffffffff00000001608083901c6001600160401b0316105b8015620045ef575067ffffffff0000000160c083901c105b15620045fd57506001919050565b506000919050565b60006200465c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316620046de9092919063ffffffff16565b8051909150156200165357808060200190518101906200467d919062005626565b620016535760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000b7f565b6060620046ef8484600085620046f7565b949350505050565b6060824710156200475a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840162000b7f565b600080866001600160a01b03168587604051620047789190620055cb565b60006040518083038185875af1925050503d8060008114620047b7576040519150601f19603f3d011682016040523d82523d6000602084013e620047bc565b606091505b509150915062001e7d87838387606083156200483d57825160000362004835576001600160a01b0385163b620048355760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000b7f565b5081620046ef565b620046ef8383815115620048545781518083602001fd5b8060405162461bcd60e51b815260040162000b7f919062004ea3565b6108f5806200567483390190565b6001600160a01b038116811462002f0957600080fd5b80356001600160401b03811681146200356c57600080fd5b6000806000806000806000806000806101408b8d031215620048cd57600080fd5b8a35620048da816200487e565b9950620048ea60208c0162004894565b9850620048fa60408c0162004894565b975060608b01356200490c816200487e565b965060808b01356200491e816200487e565b955060a08b013562004930816200487e565b945060c08b013562004942816200487e565b935060e08b013562004954816200487e565b9250620049656101008c0162004894565b9150620049766101208c0162004894565b90509295989b9194979a5092959850565b6001600160401b0391909116815260200190565b803563ffffffff811681146200356c57600080fd5b60008060408385031215620049c457600080fd5b620049cf836200499b565b9150620049df6020840162004894565b90509250929050565b60008060008060008060c0878903121562004a0257600080fd5b863562004a0f816200487e565b955062004a1f6020880162004894565b945062004a2f6040880162004894565b9350606087013562004a41816200487e565b9250608087013562004a53816200487e565b915060a087013562004a65816200487e565b809150509295509295509295565b60006020828403121562004a8657600080fd5b62000fac826200499b565b80610300810183101562000faf57600080fd5b6000806000806000806000806103e0898b03121562004ac257600080fd5b62004acd896200499b565b975062004add60208a0162004894565b965062004aed60408a0162004894565b955062004afd60608a0162004894565b945062004b0d60808a0162004894565b935060a0890135925060c0890135915062004b2c8a60e08b0162004a91565b90509295985092959890939650565b6000806000806000806000806103e0898b03121562004b5957600080fd5b62004b64896200499b565b975062004b7460208a0162004894565b965062004b8460408a0162004894565b955062004b9460608a0162004894565b94506080890135935060a0890135925060c089013562004bb4816200487e565b915062004b2c8a60e08b0162004a91565b60006020828403121562004bd857600080fd5b813561ffff8116811462004beb57600080fd5b9392505050565b60006020828403121562004c0557600080fd5b5035919050565b6000806040838503121562004c2057600080fd5b82359150602083013562004c34816200487e565b809150509250929050565b60006020828403121562004c5257600080fd5b62000fac8262004894565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171562004c9e5762004c9e62004c5d565b604052919050565b600082601f83011262004cb857600080fd5b81356001600160401b0381111562004cd45762004cd462004c5d565b62004ce9601f8201601f191660200162004c73565b81815284602083860101111562004cff57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060e0888a03121562004d3857600080fd5b62004d43886200499b565b965062004d536020890162004894565b9550604088013562004d65816200487e565b9450606088013562004d77816200487e565b9350608088013562004d89816200487e565b925060a08801356001600160401b038082111562004da657600080fd5b62004db48b838c0162004ca6565b935060c08a013591508082111562004dcb57600080fd5b5062004dda8a828b0162004ca6565b91505092959891949750929550565b60008060008060008060c0878903121562004e0357600080fd5b62004e0e876200499b565b955062004e1e6020880162004894565b945062004e2e6040880162004894565b9350606087013592506080870135915060a087013590509295509295509295565b60005b8381101562004e6c57818101518382015260200162004e52565b50506000910152565b6000815180845262004e8f81602086016020860162004e4f565b601f01601f19169290920160200192915050565b60208152600062000fac602083018462004e75565b6000602080838503121562004ecc57600080fd5b82356001600160401b038082111562004ee457600080fd5b818501915085601f83011262004ef957600080fd5b81358181111562004f0e5762004f0e62004c5d565b8060051b915062004f2184830162004c73565b818152918301840191848101908884111562004f3c57600080fd5b938501935b8385101562004f5c5784358252938501939085019062004f41565b98975050505050505050565b6000806040838503121562004f7c57600080fd5b62004f878362004894565b946020939093013593505050565b6001600160a01b0391909116815260200190565b6000806000806060858703121562004fc057600080fd5b843562004fcd816200487e565b935062004fdd602086016200499b565b925060408501356001600160401b038082111562004ffa57600080fd5b818701915087601f8301126200500f57600080fd5b8135818111156200501f57600080fd5b8860208285010111156200503257600080fd5b95989497505060200194505050565b6000602082840312156200505457600080fd5b813562004beb816200487e565b803560ff811681146200356c57600080fd5b60008060008060008060c087890312156200508d57600080fd5b86356200509a816200487e565b95506020870135620050ac816200487e565b9450620050bc6040880162004894565b9350620050cc6060880162004894565b925060808701359150620050e360a0880162005061565b90509295509295509295565b60008060008060008060c087890312156200510957600080fd5b863562005116816200487e565b9550602087013562005128816200487e565b9450620051386040880162004894565b9350620051486060880162005061565b92506080870135915060a08701356001600160401b038111156200516b57600080fd5b6200517989828a0162004ca6565b9150509295509295509295565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b8054600090600181811c9080831680620051ef57607f831692505b602080841082036200521157634e487b7160e01b600052602260045260246000fd5b838852602088018280156200522f5760018114620052465762005273565b60ff198716825285151560051b8201975062005273565b60008981526020902060005b878110156200526d5781548482015290860190840162005252565b83019850505b5050505050505092915050565b6001600160a01b0386811682528516602082015260a060408201819052600090620052ae90830186620051d4565b8281036060840152620052c28186620051d4565b9150508260808301529695505050505050565b6001600160401b039390931683526020830191909152604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160190808211156200532f576200532f620052f6565b5092915050565b808202811582820484141762000faf5762000faf620052f6565b600063ffffffff8083168181036200536c576200536c620052f6565b6001019392505050565b6001600160a01b03848116825283166020820152606060408201819052600090620053a49083018462004e75565b95945050505050565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c060808201819052600090620053ef9083018562004e75565b82810360a084015262005403818562004e75565b9998505050505050505050565b6001600160401b038281168282160390808211156200532f576200532f620052f6565b60006001600160401b038281166002600160401b031981016200536c576200536c620052f6565b634e487b7160e01b600052603260045260246000fd5b8082018082111562000faf5762000faf620052f6565b6000600182016200549b576200549b620052f6565b5060010190565b600060208284031215620054b557600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b600082620054e457620054e4620054bc565b500490565b600082620054fb57620054fb620054bc565b500690565b8181038181111562000faf5762000faf620052f6565b918252602082015260400190565b600081620055365762005536620052f6565b506000190190565b6001600160a01b03841681526040602082018190528101829052818360608301376000818301606090810191909152601f909201601f1916010192915050565b6001600160a01b038781168252861660208201526001600160401b038516604082015260ff841660608201526080810183905260c060a0820181905260009062004f5c9083018462004e75565b60008251620055df81846020870162004e4f565b9190910192915050565b61032081016103008085843782018360005b60018110156200561c578151835260209283019290910190600101620055fb565b5050509392505050565b6000602082840312156200563957600080fd5b8151801515811462004beb57600080fd5b60006001600160401b0383811680620056675762005667620054bc565b9216919091049291505056fe60a0604052604051620008f5380380620008f58339810160408190526100249161035b565b82816100308282610058565b50506001600160a01b03821660805261005061004b60805190565b6100b7565b505050610447565b61006182610126565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156100ab576100a682826101a5565b505050565b6100b361021c565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f8600080516020620008d5833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a16101238161023d565b50565b806001600160a01b03163b60000361016157604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080846001600160a01b0316846040516101c2919061042b565b600060405180830381855af49150503d80600081146101fd576040519150601f19603f3d011682016040523d82523d6000602084013e610202565b606091505b50909250905061021385838361027d565b95945050505050565b341561023b5760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661026757604051633173bdd160e11b815260006004820152602401610158565b80600080516020620008d5833981519152610184565b6060826102925761028d826102dc565b6102d5565b81511580156102a957506001600160a01b0384163b155b156102d257604051639996b31560e01b81526001600160a01b0385166004820152602401610158565b50805b9392505050565b8051156102ec5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461031c57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561035257818101518382015260200161033a565b50506000910152565b60008060006060848603121561037057600080fd5b61037984610305565b925061038760208501610305565b60408501519092506001600160401b03808211156103a457600080fd5b818601915086601f8301126103b857600080fd5b8151818111156103ca576103ca610321565b604051601f8201601f19908116603f011681019083821181831017156103f2576103f2610321565b8160405282815289602084870101111561040b57600080fd5b61041c836020830160208801610337565b80955050505050509250925092565b6000825161043d818460208701610337565b9190910192915050565b608051610473620004626000396000601001526104736000f3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361006a576000356001600160e01b03191663278f794360e11b146100625761006061006e565b565b61006061007e565b6100605b6100606100796100ad565b6100d3565b60008061008e36600481846102cb565b81019061009b919061030b565b915091506100a982826100f7565b5050565b60006100ce60008051602061041e833981519152546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100f2573d6000f35b3d6000fd5b61010082610152565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561014a5761014582826101b7565b505050565b6100a961022d565b806001600160a01b03163b6000036101885780604051634c9c8ce360e01b815260040161017f91906103da565b60405180910390fd5b60008051602061041e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516101d491906103ee565b600060405180830381855af49150503d806000811461020f576040519150601f19603f3d011682016040523d82523d6000602084013e610214565b606091505b509150915061022485838361024c565b95945050505050565b34156100605760405163b398979f60e01b815260040160405180910390fd5b6060826102615761025c826102a2565b61029b565b815115801561027857506001600160a01b0384163b155b156102985783604051639996b31560e01b815260040161017f91906103da565b50805b9392505050565b8051156102b25780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b600080858511156102db57600080fd5b838611156102e857600080fd5b5050820193919092039150565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561031e57600080fd5b82356001600160a01b038116811461033557600080fd5b915060208301356001600160401b038082111561035157600080fd5b818501915085601f83011261036557600080fd5b813581811115610377576103776102f5565b604051601f8201601f19908116603f0116810190838211818310171561039f5761039f6102f5565b816040528281528860208487010111156103b857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6001600160a01b0391909116815260200190565b6000825160005b8181101561040f57602081860181015185830152016103f5565b50600092019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212208e78e901799caaaff866d77d874534e79db9f4bae5f48cfae79611891464d2c664736f6c63430008140033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610373cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f066156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fbab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bdac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59062ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f430644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000017f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024983dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68ea5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db19b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff285951141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545ea0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a26469706673582212202276f0dddf4576c256304656d9bb90d85a72af2bb0dd4e2714ceeaab8c671a2364736f6c63430008140033",
}

// PolygonrollupmanagermockABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagermockMetaData.ABI instead.
var PolygonrollupmanagermockABI = PolygonrollupmanagermockMetaData.ABI

// PolygonrollupmanagermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagermockMetaData.Bin instead.
var PolygonrollupmanagermockBin = PolygonrollupmanagermockMetaData.Bin

// DeployPolygonrollupmanagermock deploys a new Ethereum contract, binding an instance of Polygonrollupmanagermock to it.
func DeployPolygonrollupmanagermock(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagermock, error) {
	parsed, err := PolygonrollupmanagermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagermockBin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanagermock{PolygonrollupmanagermockCaller: PolygonrollupmanagermockCaller{contract: contract}, PolygonrollupmanagermockTransactor: PolygonrollupmanagermockTransactor{contract: contract}, PolygonrollupmanagermockFilterer: PolygonrollupmanagermockFilterer{contract: contract}}, nil
}

// Polygonrollupmanagermock is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanagermock struct {
	PolygonrollupmanagermockCaller     // Read-only binding to the contract
	PolygonrollupmanagermockTransactor // Write-only binding to the contract
	PolygonrollupmanagermockFilterer   // Log filterer for contract events
}

// PolygonrollupmanagermockCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupmanagermockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagermockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupmanagermockSession struct {
	Contract     *Polygonrollupmanagermock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PolygonrollupmanagermockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupmanagermockCallerSession struct {
	Contract *PolygonrollupmanagermockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PolygonrollupmanagermockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupmanagermockTransactorSession struct {
	Contract     *PolygonrollupmanagermockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PolygonrollupmanagermockRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupmanagermockRaw struct {
	Contract *Polygonrollupmanagermock // Generic contract binding to access the raw methods on
}

// PolygonrollupmanagermockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockCallerRaw struct {
	Contract *PolygonrollupmanagermockCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupmanagermockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupmanagermockTransactorRaw struct {
	Contract *PolygonrollupmanagermockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanagermock creates a new instance of Polygonrollupmanagermock, bound to a specific deployed contract.
func NewPolygonrollupmanagermock(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanagermock, error) {
	contract, err := bindPolygonrollupmanagermock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagermock{PolygonrollupmanagermockCaller: PolygonrollupmanagermockCaller{contract: contract}, PolygonrollupmanagermockTransactor: PolygonrollupmanagermockTransactor{contract: contract}, PolygonrollupmanagermockFilterer: PolygonrollupmanagermockFilterer{contract: contract}}, nil
}

// NewPolygonrollupmanagermockCaller creates a new read-only instance of Polygonrollupmanagermock, bound to a specific deployed contract.
func NewPolygonrollupmanagermockCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupmanagermockCaller, error) {
	contract, err := bindPolygonrollupmanagermock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockCaller{contract: contract}, nil
}

// NewPolygonrollupmanagermockTransactor creates a new write-only instance of Polygonrollupmanagermock, bound to a specific deployed contract.
func NewPolygonrollupmanagermockTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupmanagermockTransactor, error) {
	contract, err := bindPolygonrollupmanagermock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockTransactor{contract: contract}, nil
}

// NewPolygonrollupmanagermockFilterer creates a new log filterer instance of Polygonrollupmanagermock, bound to a specific deployed contract.
func NewPolygonrollupmanagermockFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupmanagermockFilterer, error) {
	contract, err := bindPolygonrollupmanagermock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockFilterer{contract: contract}, nil
}

// bindPolygonrollupmanagermock binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanagermock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupmanagermockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagermock.Contract.PolygonrollupmanagermockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.PolygonrollupmanagermockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.PolygonrollupmanagermockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagermock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagermock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagermock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.BridgeAddress(&_Polygonrollupmanagermock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.BridgeAddress(&_Polygonrollupmanagermock.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagermock.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagermock.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagermock.Contract.ChainIDToRollupID(&_Polygonrollupmanagermock.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagermock.Contract.ChainIDToRollupID(&_Polygonrollupmanagermock.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.GetBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.GetBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.GetForcedBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagermock.Contract.GetForcedBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetInputSnarkBytes(&_Polygonrollupmanagermock.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetInputSnarkBytes(&_Polygonrollupmanagermock.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagermock.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagermock.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRoleAdmin(&_Polygonrollupmanagermock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRoleAdmin(&_Polygonrollupmanagermock.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupExitRoot(&_Polygonrollupmanagermock.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupExitRoot(&_Polygonrollupmanagermock.CallOpts)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetRollupPendingStateTransitions(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getRollupPendingStateTransitions", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesPendingState), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesPendingState)).(*LegacyZKEVMStateVariablesPendingState)

	return out0, err

}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupPendingStateTransitions is a free data retrieval call binding the contract method 0xb99d0ad7.
//
// Solidity: function getRollupPendingStateTransitions(uint32 rollupID, uint64 batchNum) view returns((uint64,uint64,bytes32,bytes32))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetRollupPendingStateTransitions(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesPendingState, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupPendingStateTransitions(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagermock.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagermock.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.GlobalExitRootManager(&_Polygonrollupmanagermock.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.GlobalExitRootManager(&_Polygonrollupmanagermock.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagermock.Contract.HasRole(&_Polygonrollupmanagermock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagermock.Contract.HasRole(&_Polygonrollupmanagermock.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagermock.Contract.IsEmergencyState(&_Polygonrollupmanagermock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagermock.Contract.IsEmergencyState(&_Polygonrollupmanagermock.CallOpts)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) IsPendingStateConsolidable(opts *bind.CallOpts, rollupID uint32, pendingStateNum uint64) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "isPendingStateConsolidable", rollupID, pendingStateNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagermock.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagermock.CallOpts, rollupID, pendingStateNum)
}

// IsPendingStateConsolidable is a free data retrieval call binding the contract method 0x080b3111.
//
// Solidity: function isPendingStateConsolidable(uint32 rollupID, uint64 pendingStateNum) view returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) IsPendingStateConsolidable(rollupID uint32, pendingStateNum uint64) (bool, error) {
	return _Polygonrollupmanagermock.Contract.IsPendingStateConsolidable(&_Polygonrollupmanagermock.CallOpts, rollupID, pendingStateNum)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.LastAggregationTimestamp(&_Polygonrollupmanagermock.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.LastAggregationTimestamp(&_Polygonrollupmanagermock.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagermock.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagermock.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) MultiplierBatchFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "multiplierBatchFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagermock.Contract.MultiplierBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// MultiplierBatchFee is a free data retrieval call binding the contract method 0xafd23cbe.
//
// Solidity: function multiplierBatchFee() view returns(uint16)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) MultiplierBatchFee() (uint16, error) {
	return _Polygonrollupmanagermock.Contract.MultiplierBatchFee(&_Polygonrollupmanagermock.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) PendingStateTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "pendingStateTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.PendingStateTimeout(&_Polygonrollupmanagermock.CallOpts)
}

// PendingStateTimeout is a free data retrieval call binding the contract method 0xd939b315.
//
// Solidity: function pendingStateTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) PendingStateTimeout() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.PendingStateTimeout(&_Polygonrollupmanagermock.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.Pol(&_Polygonrollupmanagermock.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.Pol(&_Polygonrollupmanagermock.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupAddressToID(&_Polygonrollupmanagermock.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupAddressToID(&_Polygonrollupmanagermock.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupCount(&_Polygonrollupmanagermock.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupCount(&_Polygonrollupmanagermock.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupData(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastPendingState, uint64 lastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupCompatibilityID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupIDToRollupData(rollupID uint32) (struct {
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
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupData(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupTypeCount(&_Polygonrollupmanagermock.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagermock.Contract.RollupTypeCount(&_Polygonrollupmanagermock.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupTypeMap(&_Polygonrollupmanagermock.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bool obsolete, bytes32 genesis)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupCompatibilityID   uint8
	Obsolete                bool
	Genesis                 [32]byte
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupTypeMap(&_Polygonrollupmanagermock.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TotalSequencedBatches(&_Polygonrollupmanagermock.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TotalSequencedBatches(&_Polygonrollupmanagermock.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TotalVerifiedBatches(&_Polygonrollupmanagermock.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TotalVerifiedBatches(&_Polygonrollupmanagermock.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) TrustedAggregatorTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "trustedAggregatorTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagermock.CallOpts)
}

// TrustedAggregatorTimeout is a free data retrieval call binding the contract method 0x841b24d7.
//
// Solidity: function trustedAggregatorTimeout() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) TrustedAggregatorTimeout() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.TrustedAggregatorTimeout(&_Polygonrollupmanagermock.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) VerifyBatchTimeTarget(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "verifyBatchTimeTarget")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagermock.CallOpts)
}

// VerifyBatchTimeTarget is a free data retrieval call binding the contract method 0x0a0d9fbe.
//
// Solidity: function verifyBatchTimeTarget() view returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) VerifyBatchTimeTarget() (uint64, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatchTimeTarget(&_Polygonrollupmanagermock.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ActivateEmergencyState(&_Polygonrollupmanagermock.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ActivateEmergencyState(&_Polygonrollupmanagermock.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddExistingRollup(&_Polygonrollupmanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe0bfd3d2.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 genesis, uint8 rollupCompatibilityID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, genesis [32]byte, rollupCompatibilityID uint8) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddExistingRollup(&_Polygonrollupmanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, genesis, rollupCompatibilityID)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddNewRollupType(&_Polygonrollupmanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xf34eb8eb.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupCompatibilityID uint8, genesis [32]byte, description string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddNewRollupType(&_Polygonrollupmanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupCompatibilityID, genesis, description)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) ConsolidatePendingState(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "consolidatePendingState", rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ConsolidatePendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum)
}

// ConsolidatePendingState is a paid mutator transaction binding the contract method 0x1608859c.
//
// Solidity: function consolidatePendingState(uint32 rollupID, uint64 pendingStateNum) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) ConsolidatePendingState(rollupID uint32, pendingStateNum uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ConsolidatePendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.CreateNewRollup(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.CreateNewRollup(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.DeactivateEmergencyState(&_Polygonrollupmanagermock.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.DeactivateEmergencyState(&_Polygonrollupmanagermock.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.GrantRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.GrantRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) Initialize(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "initialize", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.Initialize(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address polygonZkEVM, address zkEVMVerifier, uint64 zkEVMForkID, uint64 zkEVMChainID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, polygonZkEVM common.Address, zkEVMVerifier common.Address, zkEVMForkID uint64, zkEVMChainID uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.Initialize(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, polygonZkEVM, zkEVMVerifier, zkEVMForkID, zkEVMChainID)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x0e36f582.
//
// Solidity: function initializeMock(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) InitializeMock(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "initializeMock", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x0e36f582.
//
// Solidity: function initializeMock(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) InitializeMock(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.InitializeMock(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x0e36f582.
//
// Solidity: function initializeMock(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) InitializeMock(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.InitializeMock(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ObsoleteRollupType(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ObsoleteRollupType(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.OnSequenceBatches(&_Polygonrollupmanagermock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.OnSequenceBatches(&_Polygonrollupmanagermock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) OverridePendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "overridePendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.OverridePendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// OverridePendingState is a paid mutator transaction binding the contract method 0x12b86e19.
//
// Solidity: function overridePendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) OverridePendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.OverridePendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) PrepareMockCalculateRoot(opts *bind.TransactOpts, localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "prepareMockCalculateRoot", localExitRoots)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) PrepareMockCalculateRoot(localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.PrepareMockCalculateRoot(&_Polygonrollupmanagermock.TransactOpts, localExitRoots)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) PrepareMockCalculateRoot(localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.PrepareMockCalculateRoot(&_Polygonrollupmanagermock.TransactOpts, localExitRoots)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) ProveNonDeterministicPendingState(opts *bind.TransactOpts, rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "proveNonDeterministicPendingState", rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// ProveNonDeterministicPendingState is a paid mutator transaction binding the contract method 0x8bd4f071.
//
// Solidity: function proveNonDeterministicPendingState(uint32 rollupID, uint64 initPendingStateNum, uint64 finalPendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) ProveNonDeterministicPendingState(rollupID uint32, initPendingStateNum uint64, finalPendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.ProveNonDeterministicPendingState(&_Polygonrollupmanagermock.TransactOpts, rollupID, initPendingStateNum, finalPendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RenounceRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RenounceRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RevokeRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RevokeRole(&_Polygonrollupmanagermock.TransactOpts, role, account)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetBatchFee(&_Polygonrollupmanagermock.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetBatchFee(&_Polygonrollupmanagermock.TransactOpts, newBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetMultiplierBatchFee(opts *bind.TransactOpts, newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setMultiplierBatchFee", newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagermock.TransactOpts, newMultiplierBatchFee)
}

// SetMultiplierBatchFee is a paid mutator transaction binding the contract method 0x1816b7e5.
//
// Solidity: function setMultiplierBatchFee(uint16 newMultiplierBatchFee) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetMultiplierBatchFee(newMultiplierBatchFee uint16) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetMultiplierBatchFee(&_Polygonrollupmanagermock.TransactOpts, newMultiplierBatchFee)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetPendingStateTimeout(opts *bind.TransactOpts, newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setPendingStateTimeout", newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetPendingStateTimeout(&_Polygonrollupmanagermock.TransactOpts, newPendingStateTimeout)
}

// SetPendingStateTimeout is a paid mutator transaction binding the contract method 0x9c9f3dfe.
//
// Solidity: function setPendingStateTimeout(uint64 newPendingStateTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetPendingStateTimeout(newPendingStateTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetPendingStateTimeout(&_Polygonrollupmanagermock.TransactOpts, newPendingStateTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetTrustedAggregatorTimeout(opts *bind.TransactOpts, newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setTrustedAggregatorTimeout", newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagermock.TransactOpts, newTrustedAggregatorTimeout)
}

// SetTrustedAggregatorTimeout is a paid mutator transaction binding the contract method 0x394218e9.
//
// Solidity: function setTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetTrustedAggregatorTimeout(newTrustedAggregatorTimeout uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetTrustedAggregatorTimeout(&_Polygonrollupmanagermock.TransactOpts, newTrustedAggregatorTimeout)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetVerifyBatchTimeTarget(opts *bind.TransactOpts, newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setVerifyBatchTimeTarget", newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagermock.TransactOpts, newVerifyBatchTimeTarget)
}

// SetVerifyBatchTimeTarget is a paid mutator transaction binding the contract method 0xa066215c.
//
// Solidity: function setVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetVerifyBatchTimeTarget(newVerifyBatchTimeTarget uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetVerifyBatchTimeTarget(&_Polygonrollupmanagermock.TransactOpts, newVerifyBatchTimeTarget)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.UpdateRollup(&_Polygonrollupmanagermock.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.UpdateRollup(&_Polygonrollupmanagermock.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) VerifyBatches(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "verifyBatches", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatches(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatches is a paid mutator transaction binding the contract method 0x87c20c01.
//
// Solidity: function verifyBatches(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) VerifyBatches(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatches(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// PolygonrollupmanagermockAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockAddExistingRollupIterator struct {
	Event *PolygonrollupmanagermockAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockAddExistingRollup)
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
		it.Event = new(PolygonrollupmanagermockAddExistingRollup)
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
func (it *PolygonrollupmanagermockAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockAddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockAddExistingRollup struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockAddExistingRollupIterator{contract: _Polygonrollupmanagermock.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xadfc7d56f7e39b08b321534f14bfb135ad27698f7d2f5ad0edc2356ea9a3f850.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupCompatibilityID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockAddExistingRollup)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseAddExistingRollup(log types.Log) (*PolygonrollupmanagermockAddExistingRollup, error) {
	event := new(PolygonrollupmanagermockAddExistingRollup)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockAddNewRollupTypeIterator struct {
	Event *PolygonrollupmanagermockAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockAddNewRollupType)
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
		it.Event = new(PolygonrollupmanagermockAddNewRollupType)
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
func (it *PolygonrollupmanagermockAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockAddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockAddNewRollupType struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagermockAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockAddNewRollupTypeIterator{contract: _Polygonrollupmanagermock.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0xa2970448b3bd66ba7e524e7b2a5b9cf94fa29e32488fb942afdfe70dd4b77b52.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupCompatibilityID, bytes32 genesis, string description)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockAddNewRollupType)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagermockAddNewRollupType, error) {
	event := new(PolygonrollupmanagermockAddNewRollupType)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockConsolidatePendingStateIterator is returned from FilterConsolidatePendingState and is used to iterate over the raw logs and unpacked data for ConsolidatePendingState events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockConsolidatePendingStateIterator struct {
	Event *PolygonrollupmanagermockConsolidatePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockConsolidatePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockConsolidatePendingState)
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
		it.Event = new(PolygonrollupmanagermockConsolidatePendingState)
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
func (it *PolygonrollupmanagermockConsolidatePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockConsolidatePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockConsolidatePendingState represents a ConsolidatePendingState event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockConsolidatePendingState struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterConsolidatePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockConsolidatePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockConsolidatePendingStateIterator{contract: _Polygonrollupmanagermock.contract, event: "ConsolidatePendingState", logs: logs, sub: sub}, nil
}

// WatchConsolidatePendingState is a free log subscription operation binding the contract event 0x581910eb7a27738945c2f00a91f2284b2d6de9d4e472b12f901c2b0df045e21b.
//
// Solidity: event ConsolidatePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, uint64 pendingStateNum)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchConsolidatePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockConsolidatePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "ConsolidatePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockConsolidatePendingState)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseConsolidatePendingState(log types.Log) (*PolygonrollupmanagermockConsolidatePendingState, error) {
	event := new(PolygonrollupmanagermockConsolidatePendingState)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ConsolidatePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockCreateNewRollupIterator struct {
	Event *PolygonrollupmanagermockCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockCreateNewRollup)
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
		it.Event = new(PolygonrollupmanagermockCreateNewRollup)
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
func (it *PolygonrollupmanagermockCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockCreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockCreateNewRollup struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockCreateNewRollupIterator{contract: _Polygonrollupmanagermock.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockCreateNewRollup)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseCreateNewRollup(log types.Log) (*PolygonrollupmanagermockCreateNewRollup, error) {
	event := new(PolygonrollupmanagermockCreateNewRollup)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockEmergencyStateActivatedIterator struct {
	Event *PolygonrollupmanagermockEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockEmergencyStateActivated)
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
		it.Event = new(PolygonrollupmanagermockEmergencyStateActivated)
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
func (it *PolygonrollupmanagermockEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonrollupmanagermockEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockEmergencyStateActivatedIterator{contract: _Polygonrollupmanagermock.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockEmergencyStateActivated)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonrollupmanagermockEmergencyStateActivated, error) {
	event := new(PolygonrollupmanagermockEmergencyStateActivated)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockEmergencyStateDeactivatedIterator struct {
	Event *PolygonrollupmanagermockEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockEmergencyStateDeactivated)
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
		it.Event = new(PolygonrollupmanagermockEmergencyStateDeactivated)
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
func (it *PolygonrollupmanagermockEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonrollupmanagermockEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockEmergencyStateDeactivatedIterator{contract: _Polygonrollupmanagermock.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockEmergencyStateDeactivated)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonrollupmanagermockEmergencyStateDeactivated, error) {
	event := new(PolygonrollupmanagermockEmergencyStateDeactivated)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockInitializedIterator struct {
	Event *PolygonrollupmanagermockInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockInitialized)
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
		it.Event = new(PolygonrollupmanagermockInitialized)
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
func (it *PolygonrollupmanagermockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockInitialized represents a Initialized event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupmanagermockInitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockInitializedIterator{contract: _Polygonrollupmanagermock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockInitialized)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseInitialized(log types.Log) (*PolygonrollupmanagermockInitialized, error) {
	event := new(PolygonrollupmanagermockInitialized)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockObsoleteRollupTypeIterator struct {
	Event *PolygonrollupmanagermockObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockObsoleteRollupType)
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
		it.Event = new(PolygonrollupmanagermockObsoleteRollupType)
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
func (it *PolygonrollupmanagermockObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagermockObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockObsoleteRollupTypeIterator{contract: _Polygonrollupmanagermock.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockObsoleteRollupType)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseObsoleteRollupType(log types.Log) (*PolygonrollupmanagermockObsoleteRollupType, error) {
	event := new(PolygonrollupmanagermockObsoleteRollupType)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockOnSequenceBatchesIterator struct {
	Event *PolygonrollupmanagermockOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockOnSequenceBatches)
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
		it.Event = new(PolygonrollupmanagermockOnSequenceBatches)
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
func (it *PolygonrollupmanagermockOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockOnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockOnSequenceBatchesIterator{contract: _Polygonrollupmanagermock.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockOnSequenceBatches)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseOnSequenceBatches(log types.Log) (*PolygonrollupmanagermockOnSequenceBatches, error) {
	event := new(PolygonrollupmanagermockOnSequenceBatches)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockOverridePendingStateIterator is returned from FilterOverridePendingState and is used to iterate over the raw logs and unpacked data for OverridePendingState events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockOverridePendingStateIterator struct {
	Event *PolygonrollupmanagermockOverridePendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockOverridePendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockOverridePendingState)
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
		it.Event = new(PolygonrollupmanagermockOverridePendingState)
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
func (it *PolygonrollupmanagermockOverridePendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockOverridePendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockOverridePendingState represents a OverridePendingState event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockOverridePendingState struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterOverridePendingState(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockOverridePendingStateIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockOverridePendingStateIterator{contract: _Polygonrollupmanagermock.contract, event: "OverridePendingState", logs: logs, sub: sub}, nil
}

// WatchOverridePendingState is a free log subscription operation binding the contract event 0x3182bd6e6f74fc1fdc88b60f3a4f4c7f79db6ae6f5b88a1b3f5a1e28ec210d5e.
//
// Solidity: event OverridePendingState(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address aggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchOverridePendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockOverridePendingState, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "OverridePendingState", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockOverridePendingState)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseOverridePendingState(log types.Log) (*PolygonrollupmanagermockOverridePendingState, error) {
	event := new(PolygonrollupmanagermockOverridePendingState)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "OverridePendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockProveNonDeterministicPendingStateIterator is returned from FilterProveNonDeterministicPendingState and is used to iterate over the raw logs and unpacked data for ProveNonDeterministicPendingState events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockProveNonDeterministicPendingStateIterator struct {
	Event *PolygonrollupmanagermockProveNonDeterministicPendingState // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockProveNonDeterministicPendingStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockProveNonDeterministicPendingState)
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
		it.Event = new(PolygonrollupmanagermockProveNonDeterministicPendingState)
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
func (it *PolygonrollupmanagermockProveNonDeterministicPendingStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockProveNonDeterministicPendingStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockProveNonDeterministicPendingState represents a ProveNonDeterministicPendingState event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockProveNonDeterministicPendingState struct {
	StoredStateRoot [32]byte
	ProvedStateRoot [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProveNonDeterministicPendingState is a free log retrieval operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterProveNonDeterministicPendingState(opts *bind.FilterOpts) (*PolygonrollupmanagermockProveNonDeterministicPendingStateIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockProveNonDeterministicPendingStateIterator{contract: _Polygonrollupmanagermock.contract, event: "ProveNonDeterministicPendingState", logs: logs, sub: sub}, nil
}

// WatchProveNonDeterministicPendingState is a free log subscription operation binding the contract event 0x1f44c21118c4603cfb4e1b621dbcfa2b73efcececee2b99b620b2953d33a7010.
//
// Solidity: event ProveNonDeterministicPendingState(bytes32 storedStateRoot, bytes32 provedStateRoot)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchProveNonDeterministicPendingState(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockProveNonDeterministicPendingState) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "ProveNonDeterministicPendingState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockProveNonDeterministicPendingState)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseProveNonDeterministicPendingState(log types.Log) (*PolygonrollupmanagermockProveNonDeterministicPendingState, error) {
	event := new(PolygonrollupmanagermockProveNonDeterministicPendingState)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "ProveNonDeterministicPendingState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleAdminChangedIterator struct {
	Event *PolygonrollupmanagermockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockRoleAdminChanged)
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
		it.Event = new(PolygonrollupmanagermockRoleAdminChanged)
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
func (it *PolygonrollupmanagermockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonrollupmanagermockRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockRoleAdminChangedIterator{contract: _Polygonrollupmanagermock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockRoleAdminChanged)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonrollupmanagermockRoleAdminChanged, error) {
	event := new(PolygonrollupmanagermockRoleAdminChanged)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleGrantedIterator struct {
	Event *PolygonrollupmanagermockRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockRoleGranted)
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
		it.Event = new(PolygonrollupmanagermockRoleGranted)
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
func (it *PolygonrollupmanagermockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockRoleGranted represents a RoleGranted event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagermockRoleGrantedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockRoleGrantedIterator{contract: _Polygonrollupmanagermock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockRoleGranted)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseRoleGranted(log types.Log) (*PolygonrollupmanagermockRoleGranted, error) {
	event := new(PolygonrollupmanagermockRoleGranted)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleRevokedIterator struct {
	Event *PolygonrollupmanagermockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockRoleRevoked)
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
		it.Event = new(PolygonrollupmanagermockRoleRevoked)
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
func (it *PolygonrollupmanagermockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockRoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagermockRoleRevokedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockRoleRevokedIterator{contract: _Polygonrollupmanagermock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockRoleRevoked)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseRoleRevoked(log types.Log) (*PolygonrollupmanagermockRoleRevoked, error) {
	event := new(PolygonrollupmanagermockRoleRevoked)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetBatchFeeIterator struct {
	Event *PolygonrollupmanagermockSetBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetBatchFee)
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
		it.Event = new(PolygonrollupmanagermockSetBatchFee)
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
func (it *PolygonrollupmanagermockSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetBatchFeeIterator{contract: _Polygonrollupmanagermock.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetBatchFee)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetBatchFee(log types.Log) (*PolygonrollupmanagermockSetBatchFee, error) {
	event := new(PolygonrollupmanagermockSetBatchFee)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetMultiplierBatchFeeIterator is returned from FilterSetMultiplierBatchFee and is used to iterate over the raw logs and unpacked data for SetMultiplierBatchFee events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetMultiplierBatchFeeIterator struct {
	Event *PolygonrollupmanagermockSetMultiplierBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetMultiplierBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetMultiplierBatchFee)
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
		it.Event = new(PolygonrollupmanagermockSetMultiplierBatchFee)
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
func (it *PolygonrollupmanagermockSetMultiplierBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetMultiplierBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetMultiplierBatchFee represents a SetMultiplierBatchFee event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetMultiplierBatchFee struct {
	NewMultiplierBatchFee uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetMultiplierBatchFee is a free log retrieval operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetMultiplierBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetMultiplierBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetMultiplierBatchFeeIterator{contract: _Polygonrollupmanagermock.contract, event: "SetMultiplierBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetMultiplierBatchFee is a free log subscription operation binding the contract event 0x7019933d795eba185c180209e8ae8bffbaa25bcef293364687702c31f4d302c5.
//
// Solidity: event SetMultiplierBatchFee(uint16 newMultiplierBatchFee)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetMultiplierBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetMultiplierBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetMultiplierBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetMultiplierBatchFee)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetMultiplierBatchFee(log types.Log) (*PolygonrollupmanagermockSetMultiplierBatchFee, error) {
	event := new(PolygonrollupmanagermockSetMultiplierBatchFee)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetMultiplierBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetPendingStateTimeoutIterator is returned from FilterSetPendingStateTimeout and is used to iterate over the raw logs and unpacked data for SetPendingStateTimeout events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetPendingStateTimeoutIterator struct {
	Event *PolygonrollupmanagermockSetPendingStateTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetPendingStateTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetPendingStateTimeout)
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
		it.Event = new(PolygonrollupmanagermockSetPendingStateTimeout)
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
func (it *PolygonrollupmanagermockSetPendingStateTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetPendingStateTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetPendingStateTimeout represents a SetPendingStateTimeout event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetPendingStateTimeout struct {
	NewPendingStateTimeout uint64
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetPendingStateTimeout is a free log retrieval operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetPendingStateTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetPendingStateTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetPendingStateTimeoutIterator{contract: _Polygonrollupmanagermock.contract, event: "SetPendingStateTimeout", logs: logs, sub: sub}, nil
}

// WatchSetPendingStateTimeout is a free log subscription operation binding the contract event 0xc4121f4e22c69632ebb7cf1f462be0511dc034f999b52013eddfb24aab765c75.
//
// Solidity: event SetPendingStateTimeout(uint64 newPendingStateTimeout)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetPendingStateTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetPendingStateTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetPendingStateTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetPendingStateTimeout)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetPendingStateTimeout(log types.Log) (*PolygonrollupmanagermockSetPendingStateTimeout, error) {
	event := new(PolygonrollupmanagermockSetPendingStateTimeout)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetPendingStateTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagermockSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagermockSetTrustedAggregator)
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
func (it *PolygonrollupmanagermockSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetTrustedAggregatorIterator{contract: _Polygonrollupmanagermock.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetTrustedAggregator)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonrollupmanagermockSetTrustedAggregator, error) {
	event := new(PolygonrollupmanagermockSetTrustedAggregator)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator is returned from FilterSetTrustedAggregatorTimeout and is used to iterate over the raw logs and unpacked data for SetTrustedAggregatorTimeout events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator struct {
	Event *PolygonrollupmanagermockSetTrustedAggregatorTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetTrustedAggregatorTimeout)
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
		it.Event = new(PolygonrollupmanagermockSetTrustedAggregatorTimeout)
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
func (it *PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetTrustedAggregatorTimeout represents a SetTrustedAggregatorTimeout event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetTrustedAggregatorTimeout struct {
	NewTrustedAggregatorTimeout uint64
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregatorTimeout is a free log retrieval operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetTrustedAggregatorTimeout(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetTrustedAggregatorTimeoutIterator{contract: _Polygonrollupmanagermock.contract, event: "SetTrustedAggregatorTimeout", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregatorTimeout is a free log subscription operation binding the contract event 0x1f4fa24c2e4bad19a7f3ec5c5485f70d46c798461c2e684f55bbd0fc661373a1.
//
// Solidity: event SetTrustedAggregatorTimeout(uint64 newTrustedAggregatorTimeout)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetTrustedAggregatorTimeout(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetTrustedAggregatorTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetTrustedAggregatorTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetTrustedAggregatorTimeout)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetTrustedAggregatorTimeout(log types.Log) (*PolygonrollupmanagermockSetTrustedAggregatorTimeout, error) {
	event := new(PolygonrollupmanagermockSetTrustedAggregatorTimeout)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetTrustedAggregatorTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator is returned from FilterSetVerifyBatchTimeTarget and is used to iterate over the raw logs and unpacked data for SetVerifyBatchTimeTarget events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator struct {
	Event *PolygonrollupmanagermockSetVerifyBatchTimeTarget // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockSetVerifyBatchTimeTarget)
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
		it.Event = new(PolygonrollupmanagermockSetVerifyBatchTimeTarget)
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
func (it *PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockSetVerifyBatchTimeTarget represents a SetVerifyBatchTimeTarget event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockSetVerifyBatchTimeTarget struct {
	NewVerifyBatchTimeTarget uint64
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetVerifyBatchTimeTarget is a free log retrieval operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterSetVerifyBatchTimeTarget(opts *bind.FilterOpts) (*PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockSetVerifyBatchTimeTargetIterator{contract: _Polygonrollupmanagermock.contract, event: "SetVerifyBatchTimeTarget", logs: logs, sub: sub}, nil
}

// WatchSetVerifyBatchTimeTarget is a free log subscription operation binding the contract event 0x1b023231a1ab6b5d93992f168fb44498e1a7e64cef58daff6f1c216de6a68c28.
//
// Solidity: event SetVerifyBatchTimeTarget(uint64 newVerifyBatchTimeTarget)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchSetVerifyBatchTimeTarget(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockSetVerifyBatchTimeTarget) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "SetVerifyBatchTimeTarget")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockSetVerifyBatchTimeTarget)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseSetVerifyBatchTimeTarget(log types.Log) (*PolygonrollupmanagermockSetVerifyBatchTimeTarget, error) {
	event := new(PolygonrollupmanagermockSetVerifyBatchTimeTarget)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "SetVerifyBatchTimeTarget", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockUpdateRollupIterator struct {
	Event *PolygonrollupmanagermockUpdateRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockUpdateRollup)
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
		it.Event = new(PolygonrollupmanagermockUpdateRollup)
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
func (it *PolygonrollupmanagermockUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockUpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockUpdateRollupIterator{contract: _Polygonrollupmanagermock.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockUpdateRollup)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseUpdateRollup(log types.Log) (*PolygonrollupmanagermockUpdateRollup, error) {
	event := new(PolygonrollupmanagermockUpdateRollup)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyBatchesIterator struct {
	Event *PolygonrollupmanagermockVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockVerifyBatches)
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
		it.Event = new(PolygonrollupmanagermockVerifyBatches)
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
func (it *PolygonrollupmanagermockVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockVerifyBatches represents a VerifyBatches event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyBatches struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterVerifyBatches(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagermockVerifyBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockVerifyBatchesIterator{contract: _Polygonrollupmanagermock.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0xaac1e7a157b259544ebacd6e8a82ae5d6c8f174e12aa48696277bcc9a661f0b4.
//
// Solidity: event VerifyBatches(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockVerifyBatches, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "VerifyBatches", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockVerifyBatches)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseVerifyBatches(log types.Log) (*PolygonrollupmanagermockVerifyBatches, error) {
	event := new(PolygonrollupmanagermockVerifyBatches)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagermockVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagermockVerifyBatchesTrustedAggregator)
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
func (it *PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyBatchesTrustedAggregator struct {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockVerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanagermock.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockVerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonrollupmanagermockVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonrollupmanagermockVerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

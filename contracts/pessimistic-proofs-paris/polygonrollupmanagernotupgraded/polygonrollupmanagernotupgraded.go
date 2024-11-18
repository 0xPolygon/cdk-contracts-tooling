// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonrollupmanagernotupgraded

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

// LegacyZKEVMStateVariablesSequencedBatchData is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesSequencedBatchData struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}

// PolygonRollupManagerRollupDataReturn is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupManagerRollupDataReturn struct {
	RollupContract                     common.Address
	ChainID                            uint64
	Verifier                           common.Address
	ForkID                             uint64
	LastLocalExitRoot                  [32]byte
	LastBatchSequenced                 uint64
	LastVerifiedBatch                  uint64
	LegacyLastPendingState             uint64
	LegacyLastPendingStateConsolidated uint64
	LastVerifiedBatchBeforeUpgrade     uint64
	RollupTypeID                       uint64
	RollupVerifierType                 uint8
}

// PolygonRollupManagerRollupDataReturnV2 is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupManagerRollupDataReturnV2 struct {
	RollupContract                 common.Address
	ChainID                        uint64
	Verifier                       common.Address
	ForkID                         uint64
	LastLocalExitRoot              [32]byte
	LastBatchSequenced             uint64
	LastVerifiedBatch              uint64
	LastVerifiedBatchBeforeUpgrade uint64
	RollupTypeID                   uint64
	RollupVerifierType             uint8
	LastPessimisticRoot            [32]byte
	ProgramVKey                    [32]byte
}

// PolygonrollupmanagernotupgradedMetaData contains all meta data concerning the Polygonrollupmanagernotupgraded contract.
var PolygonrollupmanagernotupgradedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyChainsWithPessimisticProofs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"}],\"name\":\"createNewRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_pendingStateTimeout\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_trustedAggregatorTimeout\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"},{\"internalType\":\"contractPolygonZkEVMExistentEtrog\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"contractIVerifierRollup\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200569438038062005694833981016040819052620000349162000141565b6001600160a01b0380841660805280831660c052811660a0528282826200005a62000066565b50505050505062000195565b600054610100900460ff1615620000d35760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000126576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200013e57600080fd5b50565b6000806000606084860312156200015757600080fd5b8351620001648162000128565b6020850151909350620001778162000128565b60408501519092506200018a8162000128565b809150509250925092565b60805160a05160c05161549e620001f6600039600081816107cc01528181611f030152612fe80152600081816106660152818161296501526130b301526000818161070f01528181610ba8015281816116280152611776015261549e6000f3fe60806040523480156200001157600080fd5b50600436106200024a5760003560e01c80630645af09146200024f578063066ec012146200026857806311f6b28714620002995780631489ed1014620002b057806315064c9614620002c75780631796a1ae14620002e65780632072f6c5146200030d578063248a9ca3146200031757806325280169146200033d5780632f2ff15d14620003f257806330c27dde146200040957806336568abe146200041d578063477fa270146200043457806355a71ee0146200043d57806360469169146200048157806365c0504d146200048b5780637222020f146200050c578063727885e9146200052357806374d9c244146200053a5780637975fcfe14620005605780637fb6e76a14620005865780638129fc1c14620005af5780638875f03c14620005b95780638fd88cc214620005d057806391d1485414620005e757806397bf07e814620005fe57806399f5634e14620006155780639a908e73146200061f5780639e36c5651462000636578063a217fddf146200064d578063a2967d991462000656578063a3c573eb1462000660578063abcb51981462000697578063c1acbc3414620006ae578063c4c928c214620006c9578063ceee281d14620006e0578063d02103ca1462000709578063d5073f6f1462000731578063d547741f1462000748578063d8905812146200075f578063dbc16976146200078a578063dde0ff771462000794578063dfdb8c5e14620007af578063e46761c414620007c6578063f4e9267514620007ee578063f9c4c2ae14620007ff575b600080fd5b620002666200026036600462003afc565b62000825565b005b6084546200027c906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b6200027c620002aa36600462003bf9565b62000a8d565b62000266620002c136600462003c1e565b62000ab3565b606f54620002d59060ff1681565b604051901515815260200162000290565b607e54620002f79063ffffffff1681565b60405163ffffffff909116815260200162000290565b6200026662000c7b565b6200032e6200032836600462003cbe565b62000d41565b60405190815260200162000290565b620003be6200034e36600462003cd8565b60408051606080820183526000808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b0390811691830191909152928201519092169082015260600162000290565b620002666200040336600462003d10565b62000d56565b6087546200027c906001600160401b031681565b620002666200042e36600462003d10565b62000d7d565b6086546200032e565b6200032e6200044e36600462003cd8565b63ffffffff821660009081526081602090815260408083206001600160401b038516845260020190915290205492915050565b6200032e62000db7565b620004f76200049c36600462003bf9565b607f6020526000908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b60405162000290979695949392919062003d89565b620002666200051d36600462003bf9565b62000dcf565b620002666200053436600462003e92565b62000eb6565b620005516200054b36600462003bf9565b6200134f565b60405162000290919062003f5f565b620005776200057136600462004048565b62001482565b60405162000290919062004102565b620002f76200059736600462004117565b60836020526000908152604090205463ffffffff1681565b62000266620014b5565b62000266620005ca36600462004135565b62001596565b62000266620005e1366004620041e3565b62001830565b620002d5620005f836600462003d10565b62001bba565b620002666200060f36600462004214565b62001be5565b6200032e62001efe565b6200027c620006303660046200429b565b62001fea565b6200057762000647366004620042c8565b620021e8565b6200032e600081565b6200032e6200221b565b620006887f000000000000000000000000000000000000000000000000000000000000000081565b60405162000290919062004304565b62000266620006a836600462004318565b620025dd565b6084546200027c90600160801b90046001600160401b031681565b62000266620006da366004620043ba565b6200283c565b620002f7620006f136600462004436565b60826020526000908152604090205463ffffffff1681565b620006887f000000000000000000000000000000000000000000000000000000000000000081565b620002666200074236600462003cbe565b6200286a565b620002666200075936600462003d10565b620028fb565b620005776040518060400160405280600b81526020016a70657373696d697374696360a81b81525081565b620002666200291d565b6084546200027c90600160401b90046001600160401b031681565b62000266620007c036600462004456565b620029db565b620006887f000000000000000000000000000000000000000000000000000000000000000081565b608054620002f79063ffffffff1681565b620008166200081036600462003bf9565b62002b31565b60405162000290919062004487565b600054600290610100900460ff1615801562000848575060005460ff8083169116105b620008705760405162461bcd60e51b8152600401620008679062004587565b60405180910390fd5b6000805461ffff191660ff83161761010017905567016345785d8a00006086556200089a62002c93565b620008b5600080516020620054498339815191528c62002d00565b620008c260008862002d00565b620008dd600080516020620053698339815191528862002d00565b620008f8600080516020620053c98339815191528862002d00565b62000913600080516020620053098339815191528862002d00565b6200092e600080516020620053498339815191528962002d00565b62000949600080516020620054298339815191528962002d00565b62000964600080516020620053898339815191528962002d00565b620009907fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db18962002d00565b620009ba60008051602062005449833981519152600080516020620052c983398151915262002d0c565b620009d5600080516020620052c98339815191528962002d00565b620009f0600080516020620053298339815191528962002d00565b62000a1a60008051602062005409833981519152600080516020620053e983398151915262002d0c565b62000a35600080516020620054098339815191528762002d00565b62000a50600080516020620053e98339815191528762002d00565b6000805461ff001916905560405160ff82168152600080516020620053a98339815191529060200160405180910390a15050505050505050505050565b63ffffffff8116600090815260816020526040812062000aad9062002d61565b92915050565b6000805160206200544983398151915262000ace8162002d78565b6001600160401b0388161562000af75760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89166000908152608160205260408120906007820154600160801b900460ff16600181111562000b305762000b3062003d50565b1462000b4f576040516390db0d0760e01b815260040160405180910390fd5b62000b608189898989898962002d84565b600681018054600160401b600160801b031916600160401b6001600160401b038a169081029190911790915560009081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d62000bdf6200221b565b6040518263ffffffff1660e01b815260040162000bfe91815260200190565b600060405180830381600087803b15801562000c1957600080fd5b505af115801562000c2e573d6000803e3d6000fd5b50505050336001600160a01b03168a63ffffffff16600080516020620052e983398151915289888a60405162000c6793929190620045d5565b60405180910390a350505050505050505050565b62000c96600080516020620054098339815191523362001bba565b62000d3557608454600160801b90046001600160401b0316158062000ce65750608454429062000cdb9062093a8090600160801b90046001600160401b03166200460c565b6001600160401b0316115b8062000d165750608754429062000d0b9062093a80906001600160401b03166200460c565b6001600160401b0316115b1562000d355760405163692baaad60e11b815260040160405180910390fd5b62000d3f620030b1565b565b60009081526034602052604090206001015490565b62000d618262000d41565b62000d6c8162002d78565b62000d78838362003130565b505050565b6001600160a01b038116331462000da757604051630b4ad1cd60e31b815260040160405180910390fd5b62000db382826200319c565b5050565b6000608654606462000dca919062004636565b905090565b6000805160206200534983398151915262000dea8162002d78565b63ffffffff8216158062000e095750607e5463ffffffff908116908316115b1562000e2857604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82166000908152607f602052604090206001810154600160e81b900460ff161562000e6b57604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e4490600090a2505050565b6000805160206200542983398151915262000ed18162002d78565b63ffffffff8816158062000ef05750607e5463ffffffff908116908916115b1562000f0f57604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff88166000908152607f602052604090206001810154600160e81b900460ff161562000f5257604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b038916111562000f8157604051634c753f5760e01b815260040160405180910390fd5b6001600160401b03881660009081526083602052604090205463ffffffff161562000fbf576040516337c8fe0960e11b815260040160405180910390fd5b6080805460009190829062000fda9063ffffffff1662004650565b825463ffffffff8281166101009490940a9384029302191691909117909155825460408051600080825260208201928390529394506001600160a01b03909216913091620010289062003ac0565b620010369392919062004676565b604051809103906000f08015801562001053573d6000803e3d6000fd5b50905081608360008c6001600160401b03166001600160401b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055508160826000836001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055506000608160008463ffffffff1663ffffffff1681526020019081526020016000209050818160000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508360010160149054906101000a90046001600160401b03168160010160146101000a8154816001600160401b0302191690836001600160401b031602179055508360010160009054906101000a90046001600160a01b03168160010160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508a8160000160146101000a8154816001600160401b0302191690836001600160401b031602179055508360020154816002016000806001600160401b03168152602001908152602001600020819055508b63ffffffff168160070160086101000a8154816001600160401b0302191690836001600160401b0316021790555083600101601c9054906101000a900460ff168160070160106101000a81548160ff0219169083600181111562001263576200126362003d50565b0217905550600384015460098201556040805163ffffffff8e811682526001600160a01b0385811660208401526001600160401b038f16838501528b1660608301529151918516917f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6419181900360800190a2604051633892b81160e11b81526001600160a01b038316906371257022906200130d908d908d9088908e908e908e90600401620046a4565b600060405180830381600087803b1580156200132857600080fd5b505af11580156200133d573d6000803e3d6000fd5b50505050505050505050505050505050565b604080516101808101825260008082526020808301828152838501838152606085018481526080860185815260a0870186815260c0880187815260e089018881526101008a018981526101208b018a81526101408c018b90526101608c018b905263ffffffff8e168b5260819099529a90982080546001600160a01b038082168c52600160a01b918290046001600160401b0390811690995260018084015491821690985204871690945260058401549092526006830154808616909152600160401b9081900485169091526007820154808516909652850490921690955292939091600160801b900460ff16908111156200144f576200144f62003d50565b9081600181111562001465576200146562003d50565b905250600881015461014083015260090154610160820152919050565b63ffffffff86166000908152608160205260409020606090620014aa90878787878762003206565b979650505050505050565b600054600390610100900460ff16158015620014d8575060005460ff8083169116105b620014f75760405162461bcd60e51b8152600401620008679062004587565b6000805461ffff191660ff831617610100179055604080518082018252600b81526a70657373696d697374696360a81b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2916200155b9162004102565b60405180910390a16000805461ff001916905560405160ff82168152600080516020620053a98339815191529060200160405180910390a150565b60008051602062005449833981519152620015b18162002d78565b63ffffffff8716600090815260816020526040902060016007820154600160801b900460ff166001811115620015eb57620015eb62003d50565b146200160a57604051633a64d97360e01b815260040160405180910390fd5b60405163ef4eeb3560e01b815263ffffffff881660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ef4eeb3590602401602060405180830381865afa15801562001678573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200169e919062004707565b905080620016bf5760405163a60721e160e01b815260040160405180910390fd5b6000620016d08a84848b8b62003341565b6001840154600985015460405163020a49e360e51b81529293506001600160a01b03909116916341493c6091620017109185908b908b9060040162004721565b60006040518083038186803b1580156200172957600080fd5b505afa1580156200173e573d6000803e3d6000fd5b505060848054600160801b600160c01b031916600160801b426001600160401b031602179055505060058301889055600883018790557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d620017ad6200221b565b6040518263ffffffff1660e01b8152600401620017cc91815260200190565b600060405180830381600087803b158015620017e757600080fd5b505af1158015620017fc573d6000803e3d6000fd5b505060405133925063ffffffff8d169150600080516020620052e98339815191529062000c679060009081908e90620045d5565b6200184b600080516020620053098339815191523362001bba565b158015620018d15750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b81526004016020604051808303816000875af11580156200189f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620018c5919062004770565b6001600160a01b031614155b15620018f057604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b03821660009081526082602052604081205463ffffffff169081900362001931576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81166000908152608160205260408120906007820154600160801b900460ff1660018111156200196a576200196a62003d50565b1462001989576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b0390811690841681111580620019c2575060068201546001600160401b03600160401b9091048116908516105b15620019e15760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b03161462001a7e576001600160401b038082166000908152600385016020526040902060010154600160401b9004811690861681101562001a4957604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091166000908152600384016020526040812090815560010180546001600160801b0319169055620019e3565b6006830180546001600160401b0319166001600160401b03871617905562001aa7858362004790565b6084805460009062001ac49084906001600160401b031662004790565b82546101009290920a6001600160401b03818102199093169183160217909155861660008181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece90604401600060405180830381600087803b15801562001b4257600080fd5b505af115801562001b57573d6000803e3d6000fd5b505050506001600160401b038516600081815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a3505050505050565b60009182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020620053c983398151915262001c008162002d78565b6001600160401b03851660009081526083602052604090205463ffffffff161562001c3e576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b038616111562001c6d57604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b03881660009081526082602052604090205463ffffffff161562001cab57604051630d409b9360e41b815260040160405180910390fd5b6080805460009190829062001cc69063ffffffff1662004650565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060836000886001600160401b03166001600160401b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff16021790555080608260008b6001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055506000608160008363ffffffff1663ffffffff1681526020019081526020016000209050898160000160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550878160010160146101000a8154816001600160401b0302191690836001600160401b03160217905550888160010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550868160000160146101000a8154816001600160401b0302191690836001600160401b03160217905550848160070160106101000a81548160ff0219169083600181111562001e5d5762001e5d62003d50565b0217905550600185600181111562001e795762001e7962003d50565b0362001e9357600981018490556005810186905562001ea7565b600080805260028201602052604090208690555b8163ffffffff167fd490df184152edba8455dac3228134939c71f8cb4c8f310c3145dec9037147ac898c8a8960008a60405162001eea96959493929190620047b3565b60405180910390a250505050505050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040162001f4f919062004304565b602060405180830381865afa15801562001f6d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062001f93919062004707565b60845490915060009062001fba906001600160401b03600160401b82048116911662004790565b6001600160401b031690508060000362001fd75760009250505090565b62001fe3818362004816565b9250505090565b606f5460009060ff16156200201257604051630bc011ff60e21b815260040160405180910390fd5b3360009081526082602052604081205463ffffffff16908190036200204a576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03166000036200207557604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81166000908152608160205260408120906007820154600160801b900460ff166001811115620020ae57620020ae62003d50565b14620020cd576040516390db0d0760e01b815260040160405180910390fd5b60848054869190600090620020ed9084906001600160401b03166200460c565b82546101009290920a6001600160401b03818102199093169183160217909155600683015416905060006200212387836200460c565b6006840180546001600160401b038381166001600160401b03199092168217909255604080516060810182528a81524284166020808301918252888616838501908152600086815260038c018352859020935184559151600193909301805492518716600160401b026001600160801b03199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b63ffffffff841660009081526081602052604090206060906200221090869086868662003341565b90505b949350505050565b60805460009063ffffffff168082036200223757506000919050565b6000816001600160401b0381111562002254576200225462003ddd565b6040519080825280602002602001820160405280156200227e578160200160208202803683370190505b50905060005b82811015620022f157608160006200229e8360016200482d565b63ffffffff1663ffffffff16815260200190815260200160002060050154828281518110620022d157620022d162004843565b602090810291909101015280620022e88162004859565b91505062002284565b50600060205b83600114620025355760006200230f60028662004875565b6200231c60028762004816565b6200232891906200482d565b90506000816001600160401b0381111562002347576200234762003ddd565b60405190808252806020026020018201604052801562002371578160200160208202803683370190505b50905060005b82811015620024e9576200238d6001846200488c565b81148015620023a85750620023a460028862004875565b6001145b15620024285785620023bc82600262004636565b81518110620023cf57620023cf62004843565b602002602001015185604051602001620023eb929190620048a2565b6040516020818303038152906040528051906020012082828151811062002416576200241662004843565b602002602001018181525050620024d4565b856200243682600262004636565b8151811062002449576200244962004843565b60200260200101518682600262002461919062004636565b6200246e9060016200482d565b8151811062002481576200248162004843565b60200260200101516040516020016200249c929190620048a2565b60405160208183030381529060405280519060200120828281518110620024c757620024c762004843565b6020026020010181815250505b80620024e08162004859565b91505062002377565b50809450819550838460405160200162002505929190620048a2565b60405160208183030381529060405280519060200120935082806200252a90620048b0565b9350505050620022f7565b6000836000815181106200254d576200254d62004843565b6020026020010151905060005b82811015620025d357818460405160200162002578929190620048a2565b6040516020818303038152906040528051906020012091508384604051602001620025a5929190620048a2565b6040516020818303038152906040528051906020012093508080620025ca9062004859565b9150506200255a565b5095945050505050565b60008051602062005369833981519152620025f88162002d78565b607e8054600091908290620026139063ffffffff1662004650565b91906101000a81548163ffffffff021916908363ffffffff1602179055905060018081111562002647576200264762003d50565b8660018111156200265c576200265c62003d50565b036200268857841562002682576040516363d722e760e01b815260040160405180910390fd5b620026a8565b8215620026a8576040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876001811115620026f557620026f562003d50565b81526000602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b03199091161782559185015160018281018054958801516001600160401b0316600160a01b026001600160e01b0319909616929094169190911793909317808355938501519093909260ff60e01b1990911690600160e01b908490811115620027a557620027a562003d50565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a9062002829908c908c908c908c908c908c908c90620048ca565b60405180910390a2505050505050505050565b60008051602062005309833981519152620028578162002d78565b620028648484846200342b565b50505050565b60008051602062005329833981519152620028858162002d78565b683635c9adc5dea00000821180620028a05750633b9aca0082105b15620028bf57604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b620029068262000d41565b620029118162002d78565b62000d7883836200319c565b60008051602062005389833981519152620029388162002d78565b608780546001600160401b031916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc1697691600480830192600092919082900301818387803b158015620029b557600080fd5b505af1158015620029ca573d6000803e3d6000fd5b50505050620029d862003706565b50565b336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b81526004016020604051808303816000875af115801562002a26573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002a4c919062004770565b6001600160a01b03161462002a745760405163696072e960e01b815260040160405180910390fd5b6001600160a01b03821660009081526082602090815260408083205463ffffffff1683526081909152902060068101546001600160401b03808216600160401b909204161462002ad75760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b03161062002b1357604051633e37e23360e01b815260040160405180910390fd5b60408051600081526020810190915262000d7890849084906200342b565b62002b976040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff821660009081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b039081169487019490945260018084015491821695870195909552048216606085015260058101546080850152600681015480831660a0860152600160401b808204841660c0870152600160801b808304851660e0880152600160c01b9092048416610100870152600783015480851661012088015290810490931661014086015290926101608501929190910460ff169081111562002c735762002c7362003d50565b9081600181111562002c895762002c8962003d50565b8152505050919050565b600054610100900460ff1662000d3f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840162000867565b62000db3828262003130565b600062002d198362000d41565b600084815260346020526040808220600101859055519192508391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b60060154600160401b90046001600160401b031690565b620029d881336200375f565b60008062002d928962002d61565b60078a01549091506001600160401b03908116908916101562002dc85760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b038816600090815260028a01602052604090205491508162002e04576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b0316111562002e3857604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b03161162002e6b5760405163b9b18f5760e01b815260040160405180910390fd5b600062002e7d8a8a8a8a878b62003206565b905060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405162002eb491906200492e565b602060405180830381855afa15801562002ed2573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019062002ef7919062004707565b62002f03919062004875565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a9162002f47918991906004016200494c565b602060405180830381865afa15801562002f65573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062002f8b919062004989565b62002fa9576040516309bde33960e01b815260040160405180910390fd5b600062002fb7848b62004790565b90506200301087826001600160401b031662002fd262001efe565b62002fde919062004636565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016919062003789565b80608460088282829054906101000a90046001600160401b03166200303691906200460c565b82546101009290920a6001600160401b0381810219909316918316021790915560848054600160801b600160c01b031916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064016200130d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200310d57600080fd5b505af115801562003122573d6000803e3d6000fd5b5050505062000d3f620037dd565b6200313c828262001bba565b62000db35760008281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b620031a8828262001bba565b1562000db35760008281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6001600160401b03808616600081815260038901602052604080822054938816825290205460609291158015906200323c575081155b156200325b5760405163340c614f60e11b815260040160405180910390fd5b806200327a576040516366385b5160e01b815260040160405180910390fd5b62003285846200383a565b620032a3576040516305dae44f60e21b815260040160405180910390fd5b885460018a01546040516001600160601b03193360601b16602082015260348101889052605481018590526001600160c01b031960c08c811b82166074840152600160a01b94859004811b8216607c84015293909204831b82166084820152608c810187905260ac810184905260cc81018990529189901b1660ec82015260f401604051602081830303815290604052925050509695505050505050565b606060008560000160009054906101000a90046001600160a01b03166001600160a01b031663ad1edf346040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200339b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620033c1919062004707565b60058701546008880154604080516020810193909352820152606081018790526001600160e01b031960e08a901b1660808201526084810182905260a4810186905260c4810185905290915060e40160405160208183030381529060405291505095945050505050565b63ffffffff821615806200344a5750607e5463ffffffff908116908316115b156200346957604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b03831660009081526082602052604081205463ffffffff1690819003620034aa576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181166000908152608160205260409020600781015490918516600160401b9091046001600160401b031603620034f957604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84166000908152607f602052604090206001810154600160e81b900460ff16156200353c57604051633b8d3d9960e01b815260040160405180910390fd5b600181810154600160e01b900460ff16908111156200355f576200355f62003d50565b6007830154600160801b900460ff16600181111562003582576200358262003d50565b14620035a157604051635aa0d5f160e11b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825591546001600160401b03600160a01b9182900416026001600160e01b031990921690921717905560038101546009830155600782018054600160401b63ffffffff881602600160401b600160801b031990911617905560006200362b8462000a8d565b6007840180546001600160401b0319166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef286926200367e9216908990600401620049ad565b600060405180830381600087803b1580156200369957600080fd5b505af1158015620036ae573d6000803e3d6000fd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff166200372a57604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b390600090a1565b6200376b828262001bba565b62000db357604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b17905262000d78908490620038c5565b606f5460ff16156200380257604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a549790600090a1565b600067ffffffff000000016001600160401b03831610801562003871575067ffffffff00000001604083901c6001600160401b0316105b801562003892575067ffffffff00000001608083901c6001600160401b0316105b8015620038aa575067ffffffff0000000160c083901c105b15620038b857506001919050565b506000919050565b919050565b60006200391c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166200399e9092919063ffffffff16565b80519091501562000d7857808060200190518101906200393d919062004989565b62000d785760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840162000867565b606062002213848460008585600080866001600160a01b03168587604051620039c891906200492e565b60006040518083038185875af1925050503d806000811462003a07576040519150601f19603f3d011682016040523d82523d6000602084013e62003a0c565b606091505b5091509150620014aa878383876060831562003a8d57825160000362003a85576001600160a01b0385163b62003a855760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640162000867565b508162002213565b62002213838381511562003aa45781518083602001fd5b8060405162461bcd60e51b815260040162000867919062004102565b6108f580620049d483390190565b6001600160a01b0381168114620029d857600080fd5b80356001600160401b0381168114620038c057600080fd5b6000806000806000806000806000806101408b8d03121562003b1d57600080fd5b8a3562003b2a8162003ace565b995062003b3a60208c0162003ae4565b985062003b4a60408c0162003ae4565b975060608b013562003b5c8162003ace565b965060808b013562003b6e8162003ace565b955060a08b013562003b808162003ace565b945060c08b013562003b928162003ace565b935060e08b013562003ba48162003ace565b925062003bb56101008c0162003ae4565b915062003bc66101208c0162003ae4565b90509295989b9194979a5092959850565b6001600160401b03169052565b803563ffffffff81168114620038c057600080fd5b60006020828403121562003c0c57600080fd5b62003c178262003be4565b9392505050565b6000806000806000806000806103e0808a8c03121562003c3d57600080fd5b62003c488a62003be4565b985062003c5860208b0162003ae4565b975062003c6860408b0162003ae4565b965062003c7860608b0162003ae4565b955060808a0135945060a08a0135935060c08a013562003c988162003ace565b92508981018b101562003caa57600080fd5b5060e0890190509295985092959890939650565b60006020828403121562003cd157600080fd5b5035919050565b6000806040838503121562003cec57600080fd5b62003cf78362003be4565b915062003d076020840162003ae4565b90509250929050565b6000806040838503121562003d2457600080fd5b82359150602083013562003d388162003ace565b809150509250929050565b6001600160a01b03169052565b634e487b7160e01b600052602160045260246000fd5b6002811062003d8557634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e0810162003dbf606083018762003d66565b931515608082015260a081019290925260c090910152949350505050565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b038084111562003e105762003e1062003ddd565b604051601f8501601f19908116603f0116810190828211818310171562003e3b5762003e3b62003ddd565b8160405280935085815286868601111562003e5557600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011262003e8157600080fd5b62003c178383356020850162003df3565b600080600080600080600060e0888a03121562003eae57600080fd5b62003eb98862003be4565b965062003ec96020890162003ae4565b9550604088013562003edb8162003ace565b9450606088013562003eed8162003ace565b9350608088013562003eff8162003ace565b925060a08801356001600160401b038082111562003f1c57600080fd5b62003f2a8b838c0162003e6f565b935060c08a013591508082111562003f4157600080fd5b5062003f508a828b0162003e6f565b91505092959891949750929550565b60006101808201905062003f7582845162003d43565b602083015162003f89602084018262003bd7565b50604083015162003f9e604084018262003d43565b50606083015162003fb3606084018262003bd7565b506080830151608083015260a083015162003fd260a084018262003bd7565b5060c083015162003fe760c084018262003bd7565b5060e083015162003ffc60e084018262003bd7565b5061010080840151620040128285018262003bd7565b505061012080840151620040298285018262003d66565b5050610140838101519083015261016092830151929091019190915290565b60008060008060008060c087890312156200406257600080fd5b6200406d8762003be4565b95506200407d6020880162003ae4565b94506200408d6040880162003ae4565b9350606087013592506080870135915060a087013590509295509295509295565b60005b83811015620040cb578181015183820152602001620040b1565b50506000910152565b60008151808452620040ee816020860160208601620040ae565b601f01601f19169290920160200192915050565b60208152600062003c176020830184620040d4565b6000602082840312156200412a57600080fd5b62003c178262003ae4565b60008060008060008060a087890312156200414f57600080fd5b6200415a8762003be4565b95506200416a6020880162003be4565b9450604087013593506060870135925060808701356001600160401b03808211156200419557600080fd5b818901915089601f830112620041aa57600080fd5b813581811115620041ba57600080fd5b8a6020828501011115620041cd57600080fd5b6020830194508093505050509295509295509295565b60008060408385031215620041f757600080fd5b823562003cf78162003ace565b803560028110620038c057600080fd5b600080600080600080600060e0888a0312156200423057600080fd5b87356200423d8162003ace565b965060208801356200424f8162003ace565b95506200425f6040890162003ae4565b94506200426f6060890162003ae4565b9350608088013592506200428660a0890162004204565b915060c0880135905092959891949750929550565b60008060408385031215620042af57600080fd5b620042ba8362003ae4565b946020939093013593505050565b60008060008060808587031215620042df57600080fd5b620042ea8562003be4565b966020860135965060408601359560600135945092505050565b6001600160a01b0391909116815260200190565b600080600080600080600060e0888a0312156200433457600080fd5b8735620043418162003ace565b96506020880135620043538162003ace565b9550620043636040890162003ae4565b9450620043736060890162004204565b93506080880135925060a08801356001600160401b038111156200439657600080fd5b620043a48a828b0162003e6f565b92505060c0880135905092959891949750929550565b600080600060608486031215620043d057600080fd5b8335620043dd8162003ace565b9250620043ed6020850162003be4565b915060408401356001600160401b038111156200440957600080fd5b8401601f810186136200441b57600080fd5b6200442c8682356020840162003df3565b9150509250925092565b6000602082840312156200444957600080fd5b813562003c178162003ace565b600080604083850312156200446a57600080fd5b8235620044778162003ace565b915062003d076020840162003be4565b6000610180820190506200449d82845162003d43565b6020830151620044b1602084018262003bd7565b506040830151620044c6604084018262003d43565b506060830151620044db606084018262003bd7565b506080830151608083015260a0830151620044fa60a084018262003bd7565b5060c08301516200450f60c084018262003bd7565b5060e08301516200452460e084018262003bd7565b50610100808401516200453a8285018262003bd7565b505061012080840151620045518285018262003bd7565b505061014080840151620045688285018262003bd7565b5050610160808401516200457f8285018262003d66565b505092915050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6001600160401b039390931683526020830191909152604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160190808211156200462f576200462f620045f6565b5092915050565b808202811582820484141762000aad5762000aad620045f6565b600063ffffffff8083168181036200466c576200466c620045f6565b6001019392505050565b6001600160a01b038481168252831660208201526060604082018190526000906200221090830184620040d4565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c060808201819052600090620046e690830185620040d4565b82810360a0840152620046fa8185620040d4565b9998505050505050505050565b6000602082840312156200471a57600080fd5b5051919050565b8481526060602082015260006200473c6060830186620040d4565b8281036040840152838152838560208301376000602085830101526020601f19601f86011682010191505095945050505050565b6000602082840312156200478357600080fd5b815162003c178162003ace565b6001600160401b038281168282160390808211156200462f576200462f620045f6565b6001600160401b0387811682526001600160a01b0387166020830152858116604083015260c0820190620047eb606084018762003d66565b93909316608082015260a00152949350505050565b634e487b7160e01b600052601260045260246000fd5b60008262004828576200482862004800565b500490565b8082018082111562000aad5762000aad620045f6565b634e487b7160e01b600052603260045260246000fd5b6000600182016200486e576200486e620045f6565b5060010190565b60008262004887576200488762004800565b500690565b8181038181111562000aad5762000aad620045f6565b918252602082015260400190565b600081620048c257620048c2620045f6565b506000190190565b6001600160a01b038881168252871660208201526001600160401b03861660408201526000620048fe606083018762003d66565b84608083015260e060a08301526200491a60e0830185620040d4565b90508260c083015298975050505050505050565b6000825162004942818460208701620040ae565b9190910192915050565b61032081016103008085843782018360005b60018110156200497f5781518352602092830192909101906001016200495e565b5050509392505050565b6000602082840312156200499c57600080fd5b8151801515811462003c1757600080fd5b6001600160a01b03831681526040602082018190526000906200221390830184620040d456fe60a0604052604051620008f5380380620008f58339810160408190526100249161035b565b82816100308282610058565b50506001600160a01b03821660805261005061004b60805190565b6100b7565b505050610447565b61006182610126565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156100ab576100a682826101a5565b505050565b6100b361021c565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f8600080516020620008d5833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a16101238161023d565b50565b806001600160a01b03163b60000361016157604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080846001600160a01b0316846040516101c2919061042b565b600060405180830381855af49150503d80600081146101fd576040519150601f19603f3d011682016040523d82523d6000602084013e610202565b606091505b50909250905061021385838361027d565b95945050505050565b341561023b5760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661026757604051633173bdd160e11b815260006004820152602401610158565b80600080516020620008d5833981519152610184565b6060826102925761028d826102dc565b6102d5565b81511580156102a957506001600160a01b0384163b155b156102d257604051639996b31560e01b81526001600160a01b0385166004820152602401610158565b50805b9392505050565b8051156102ec5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461031c57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561035257818101518382015260200161033a565b50506000910152565b60008060006060848603121561037057600080fd5b61037984610305565b925061038760208501610305565b60408501519092506001600160401b03808211156103a457600080fd5b818601915086601f8301126103b857600080fd5b8151818111156103ca576103ca610321565b604051601f8201601f19908116603f011681019083821181831017156103f2576103f2610321565b8160405282815289602084870101111561040b57600080fd5b61041c836020830160208801610337565b80955050505050509250925092565b6000825161043d818460208701610337565b9190910192915050565b608051610473620004626000396000601001526104736000f3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361006a576000356001600160e01b03191663278f794360e11b146100625761006061006e565b565b61006061007e565b6100605b6100606100796100ad565b6100d3565b60008061008e36600481846102cb565b81019061009b919061030b565b915091506100a982826100f7565b5050565b60006100ce60008051602061041e833981519152546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100f2573d6000f35b3d6000fd5b61010082610152565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561014a5761014582826101b7565b505050565b6100a961022d565b806001600160a01b03163b6000036101885780604051634c9c8ce360e01b815260040161017f91906103da565b60405180910390fd5b60008051602061041e83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516101d491906103ee565b600060405180830381855af49150503d806000811461020f576040519150601f19603f3d011682016040523d82523d6000602084013e610214565b606091505b509150915061022485838361024c565b95945050505050565b34156100605760405163b398979f60e01b815260040160405180910390fd5b6060826102615761025c826102a2565b61029b565b815115801561027857506001600160a01b0384163b155b156102985783604051639996b31560e01b815260040161017f91906103da565b50805b9392505050565b8051156102b25780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b600080858511156102db57600080fd5b838611156102e857600080fd5b5050820193919092039150565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561031e57600080fd5b82356001600160a01b038116811461033557600080fd5b915060208301356001600160401b038082111561035157600080fd5b818501915085601f83011261036557600080fd5b813581811115610377576103776102f5565b604051601f8201601f19908116603f0116810190838211818310171561039f5761039f6102f5565b816040528281528860208487010111156103b857600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6001600160a01b0391909116815260200190565b6000825160005b8181101561040f57602081860181015185830152016103f5565b50600092019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212208e78e901799caaaff866d77d874534e79db9f4bae5f48cfae79611891464d2c664736f6c63430008140033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610373cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0d1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d366156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fbab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bdac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59062ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f47f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024983dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff285951141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545ea0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a264697066735822122028ef20d1e3fb7f101901b760a1f491fff6a78fd352d5f674e604e91a2b9e9c6164736f6c63430008140033",
}

// PolygonrollupmanagernotupgradedABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagernotupgradedMetaData.ABI instead.
var PolygonrollupmanagernotupgradedABI = PolygonrollupmanagernotupgradedMetaData.ABI

// PolygonrollupmanagernotupgradedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagernotupgradedMetaData.Bin instead.
var PolygonrollupmanagernotupgradedBin = PolygonrollupmanagernotupgradedMetaData.Bin

// DeployPolygonrollupmanagernotupgraded deploys a new Ethereum contract, binding an instance of Polygonrollupmanagernotupgraded to it.
func DeployPolygonrollupmanagernotupgraded(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagernotupgraded, error) {
	parsed, err := PolygonrollupmanagernotupgradedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagernotupgradedBin), backend, _globalExitRootManager, _pol, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonrollupmanagernotupgraded{PolygonrollupmanagernotupgradedCaller: PolygonrollupmanagernotupgradedCaller{contract: contract}, PolygonrollupmanagernotupgradedTransactor: PolygonrollupmanagernotupgradedTransactor{contract: contract}, PolygonrollupmanagernotupgradedFilterer: PolygonrollupmanagernotupgradedFilterer{contract: contract}}, nil
}

// Polygonrollupmanagernotupgraded is an auto generated Go binding around an Ethereum contract.
type Polygonrollupmanagernotupgraded struct {
	PolygonrollupmanagernotupgradedCaller     // Read-only binding to the contract
	PolygonrollupmanagernotupgradedTransactor // Write-only binding to the contract
	PolygonrollupmanagernotupgradedFilterer   // Log filterer for contract events
}

// PolygonrollupmanagernotupgradedCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonrollupmanagernotupgradedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagernotupgradedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonrollupmanagernotupgradedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagernotupgradedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonrollupmanagernotupgradedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonrollupmanagernotupgradedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonrollupmanagernotupgradedSession struct {
	Contract     *Polygonrollupmanagernotupgraded // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PolygonrollupmanagernotupgradedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonrollupmanagernotupgradedCallerSession struct {
	Contract *PolygonrollupmanagernotupgradedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// PolygonrollupmanagernotupgradedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonrollupmanagernotupgradedTransactorSession struct {
	Contract     *PolygonrollupmanagernotupgradedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// PolygonrollupmanagernotupgradedRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonrollupmanagernotupgradedRaw struct {
	Contract *Polygonrollupmanagernotupgraded // Generic contract binding to access the raw methods on
}

// PolygonrollupmanagernotupgradedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonrollupmanagernotupgradedCallerRaw struct {
	Contract *PolygonrollupmanagernotupgradedCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonrollupmanagernotupgradedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonrollupmanagernotupgradedTransactorRaw struct {
	Contract *PolygonrollupmanagernotupgradedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonrollupmanagernotupgraded creates a new instance of Polygonrollupmanagernotupgraded, bound to a specific deployed contract.
func NewPolygonrollupmanagernotupgraded(address common.Address, backend bind.ContractBackend) (*Polygonrollupmanagernotupgraded, error) {
	contract, err := bindPolygonrollupmanagernotupgraded(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonrollupmanagernotupgraded{PolygonrollupmanagernotupgradedCaller: PolygonrollupmanagernotupgradedCaller{contract: contract}, PolygonrollupmanagernotupgradedTransactor: PolygonrollupmanagernotupgradedTransactor{contract: contract}, PolygonrollupmanagernotupgradedFilterer: PolygonrollupmanagernotupgradedFilterer{contract: contract}}, nil
}

// NewPolygonrollupmanagernotupgradedCaller creates a new read-only instance of Polygonrollupmanagernotupgraded, bound to a specific deployed contract.
func NewPolygonrollupmanagernotupgradedCaller(address common.Address, caller bind.ContractCaller) (*PolygonrollupmanagernotupgradedCaller, error) {
	contract, err := bindPolygonrollupmanagernotupgraded(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedCaller{contract: contract}, nil
}

// NewPolygonrollupmanagernotupgradedTransactor creates a new write-only instance of Polygonrollupmanagernotupgraded, bound to a specific deployed contract.
func NewPolygonrollupmanagernotupgradedTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonrollupmanagernotupgradedTransactor, error) {
	contract, err := bindPolygonrollupmanagernotupgraded(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedTransactor{contract: contract}, nil
}

// NewPolygonrollupmanagernotupgradedFilterer creates a new log filterer instance of Polygonrollupmanagernotupgraded, bound to a specific deployed contract.
func NewPolygonrollupmanagernotupgradedFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonrollupmanagernotupgradedFilterer, error) {
	contract, err := bindPolygonrollupmanagernotupgraded(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedFilterer{contract: contract}, nil
}

// bindPolygonrollupmanagernotupgraded binds a generic wrapper to an already deployed contract.
func bindPolygonrollupmanagernotupgraded(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonrollupmanagernotupgradedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagernotupgraded.Contract.PolygonrollupmanagernotupgradedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.PolygonrollupmanagernotupgradedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.PolygonrollupmanagernotupgradedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonrollupmanagernotupgraded.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.DEFAULTADMINROLE(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) ROLLUPMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "ROLLUP_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ROLLUPMANAGERVERSION(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ROLLUPMANAGERVERSION(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.BridgeAddress(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.BridgeAddress(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.CalculateRewardPerBatch(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ChainIDToRollupID(&_Polygonrollupmanagernotupgraded.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ChainIDToRollupID(&_Polygonrollupmanagernotupgraded.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetBatchFee(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetBatchFee(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetForcedBatchFee(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetForcedBatchFee(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0x9e36c565.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputSnarkBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputSnarkBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetLastVerifiedBatch(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRoleAdmin(&_Polygonrollupmanagernotupgraded.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRoleAdmin(&_Polygonrollupmanagernotupgraded.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupBatchNumToStateRoot(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupExitRoot(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupExitRoot(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetRollupSequencedBatches(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GlobalExitRootManager(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GlobalExitRootManager(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.HasRole(&_Polygonrollupmanagernotupgraded.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.HasRole(&_Polygonrollupmanagernotupgraded.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.IsEmergencyState(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.IsEmergencyState(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.LastAggregationTimestamp(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.LastAggregationTimestamp(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.LastDeactivatedEmergencyStateTimestamp(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Pol(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) Pol() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Pol(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupAddressToID(&_Polygonrollupmanagernotupgraded.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupAddressToID(&_Polygonrollupmanagernotupgraded.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupCount(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupCount() (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupCount(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	if err != nil {
		return *new(PolygonRollupManagerRollupDataReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(PolygonRollupManagerRollupDataReturn)).(*PolygonRollupManagerRollupDataReturn)

	return out0, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupIDToRollupData(rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupData(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupIDToRollupData(rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupData(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupIDToRollupDataV2(opts *bind.CallOpts, rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataV2", rollupID)

	if err != nil {
		return *new(PolygonRollupManagerRollupDataReturnV2), err
	}

	out0 := *abi.ConvertType(out[0], new(PolygonRollupManagerRollupDataReturnV2)).(*PolygonRollupManagerRollupDataReturnV2)

	return out0, err

}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupIDToRollupDataV2(rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataV2(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupIDToRollupDataV2(rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataV2(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupTypeCount(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupTypeCount() (uint32, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupTypeCount(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

	outstruct := new(struct {
		ConsensusImplementation common.Address
		Verifier                common.Address
		ForkID                  uint64
		RollupVerifierType      uint8
		Obsolete                bool
		Genesis                 [32]byte
		ProgramVKey             [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusImplementation = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Verifier = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ForkID = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RollupVerifierType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Obsolete = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Genesis = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.ProgramVKey = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupTypeMap(&_Polygonrollupmanagernotupgraded.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollupTypeMap(&_Polygonrollupmanagernotupgraded.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.TotalSequencedBatches(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.TotalSequencedBatches(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.TotalVerifiedBatches(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Polygonrollupmanagernotupgraded.Contract.TotalVerifiedBatches(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ActivateEmergencyState(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ActivateEmergencyState(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddExistingRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0x97bf07e8.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddExistingRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddNewRollupType(&_Polygonrollupmanagernotupgraded.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddNewRollupType(&_Polygonrollupmanagernotupgraded.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) CreateNewRollup(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "createNewRollup", rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.CreateNewRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// CreateNewRollup is a paid mutator transaction binding the contract method 0x727885e9.
//
// Solidity: function createNewRollup(uint32 rollupTypeID, uint64 chainID, address admin, address sequencer, address gasTokenAddress, string sequencerURL, string networkName) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) CreateNewRollup(rollupTypeID uint32, chainID uint64, admin common.Address, sequencer common.Address, gasTokenAddress common.Address, sequencerURL string, networkName string) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.CreateNewRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID, chainID, admin, sequencer, gasTokenAddress, sequencerURL, networkName)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.DeactivateEmergencyState(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.DeactivateEmergencyState(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GrantRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GrantRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address , address , uint64 , uint64 ) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) Initialize(opts *bind.TransactOpts, trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, arg6 common.Address, arg7 common.Address, arg8 uint64, arg9 uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "initialize", trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, arg6, arg7, arg8, arg9)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address , address , uint64 , uint64 ) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, arg6 common.Address, arg7 common.Address, arg8 uint64, arg9 uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize(&_Polygonrollupmanagernotupgraded.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, arg6, arg7, arg8, arg9)
}

// Initialize is a paid mutator transaction binding the contract method 0x0645af09.
//
// Solidity: function initialize(address trustedAggregator, uint64 _pendingStateTimeout, uint64 _trustedAggregatorTimeout, address admin, address timelock, address emergencyCouncil, address , address , uint64 , uint64 ) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) Initialize(trustedAggregator common.Address, _pendingStateTimeout uint64, _trustedAggregatorTimeout uint64, admin common.Address, timelock common.Address, emergencyCouncil common.Address, arg6 common.Address, arg7 common.Address, arg8 uint64, arg9 uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize(&_Polygonrollupmanagernotupgraded.TransactOpts, trustedAggregator, _pendingStateTimeout, _trustedAggregatorTimeout, admin, timelock, emergencyCouncil, arg6, arg7, arg8, arg9)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) Initialize0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "initialize0")
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Initialize0() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize0(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) Initialize0() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize0(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ObsoleteRollupType(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.ObsoleteRollupType(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.OnSequenceBatches(&_Polygonrollupmanagernotupgraded.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.OnSequenceBatches(&_Polygonrollupmanagernotupgraded.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RenounceRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RenounceRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RevokeRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RevokeRole(&_Polygonrollupmanagernotupgraded.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollbackBatches(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.RollbackBatches(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.SetBatchFee(&_Polygonrollupmanagernotupgraded.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.SetBatchFee(&_Polygonrollupmanagernotupgraded.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.UpdateRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.UpdateRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyBatchesTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x8875f03c.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof)
}

// PolygonrollupmanagernotupgradedAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedAddExistingRollupIterator struct {
	Event *PolygonrollupmanagernotupgradedAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedAddExistingRollup)
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
		it.Event = new(PolygonrollupmanagernotupgradedAddExistingRollup)
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
func (it *PolygonrollupmanagernotupgradedAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedAddExistingRollup represents a AddExistingRollup event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedAddExistingRollup struct {
	RollupID                       uint32
	ForkID                         uint64
	RollupAddress                  common.Address
	ChainID                        uint64
	RollupVerifierType             uint8
	LastVerifiedBatchBeforeUpgrade uint64
	ProgramVKey                    [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterAddExistingRollup is a free log retrieval operation binding the contract event 0xd490df184152edba8455dac3228134939c71f8cb4c8f310c3145dec9037147ac.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedAddExistingRollupIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0xd490df184152edba8455dac3228134939c71f8cb4c8f310c3145dec9037147ac.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedAddExistingRollup)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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

// ParseAddExistingRollup is a log parse operation binding the contract event 0xd490df184152edba8455dac3228134939c71f8cb4c8f310c3145dec9037147ac.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseAddExistingRollup(log types.Log) (*PolygonrollupmanagernotupgradedAddExistingRollup, error) {
	event := new(PolygonrollupmanagernotupgradedAddExistingRollup)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedAddNewRollupTypeIterator struct {
	Event *PolygonrollupmanagernotupgradedAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedAddNewRollupType)
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
		it.Event = new(PolygonrollupmanagernotupgradedAddNewRollupType)
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
func (it *PolygonrollupmanagernotupgradedAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedAddNewRollupType represents a AddNewRollupType event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedAddNewRollupType struct {
	RollupTypeID            uint32
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Genesis                 [32]byte
	Description             string
	ProgramVKey             [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAddNewRollupType is a free log retrieval operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagernotupgradedAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedAddNewRollupTypeIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedAddNewRollupType)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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

// ParseAddNewRollupType is a log parse operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagernotupgradedAddNewRollupType, error) {
	event := new(PolygonrollupmanagernotupgradedAddNewRollupType)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCreateNewRollupIterator struct {
	Event *PolygonrollupmanagernotupgradedCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedCreateNewRollup)
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
		it.Event = new(PolygonrollupmanagernotupgradedCreateNewRollup)
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
func (it *PolygonrollupmanagernotupgradedCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedCreateNewRollup represents a CreateNewRollup event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCreateNewRollup struct {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedCreateNewRollupIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedCreateNewRollup)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseCreateNewRollup(log types.Log) (*PolygonrollupmanagernotupgradedCreateNewRollup, error) {
	event := new(PolygonrollupmanagernotupgradedCreateNewRollup)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator struct {
	Event *PolygonrollupmanagernotupgradedEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedEmergencyStateActivated)
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
		it.Event = new(PolygonrollupmanagernotupgradedEmergencyStateActivated)
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
func (it *PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedEmergencyStateActivatedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedEmergencyStateActivated)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonrollupmanagernotupgradedEmergencyStateActivated, error) {
	event := new(PolygonrollupmanagernotupgradedEmergencyStateActivated)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator struct {
	Event *PolygonrollupmanagernotupgradedEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedEmergencyStateDeactivated)
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
		it.Event = new(PolygonrollupmanagernotupgradedEmergencyStateDeactivated)
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
func (it *PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedEmergencyStateDeactivatedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedEmergencyStateDeactivated)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonrollupmanagernotupgradedEmergencyStateDeactivated, error) {
	event := new(PolygonrollupmanagernotupgradedEmergencyStateDeactivated)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedInitializedIterator struct {
	Event *PolygonrollupmanagernotupgradedInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedInitialized)
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
		it.Event = new(PolygonrollupmanagernotupgradedInitialized)
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
func (it *PolygonrollupmanagernotupgradedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedInitialized represents a Initialized event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedInitializedIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedInitializedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedInitialized)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseInitialized(log types.Log) (*PolygonrollupmanagernotupgradedInitialized, error) {
	event := new(PolygonrollupmanagernotupgradedInitialized)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator struct {
	Event *PolygonrollupmanagernotupgradedObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedObsoleteRollupType)
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
		it.Event = new(PolygonrollupmanagernotupgradedObsoleteRollupType)
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
func (it *PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedObsoleteRollupType represents a ObsoleteRollupType event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedObsoleteRollupTypeIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedObsoleteRollupType)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseObsoleteRollupType(log types.Log) (*PolygonrollupmanagernotupgradedObsoleteRollupType, error) {
	event := new(PolygonrollupmanagernotupgradedObsoleteRollupType)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedOnSequenceBatchesIterator struct {
	Event *PolygonrollupmanagernotupgradedOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedOnSequenceBatches)
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
		it.Event = new(PolygonrollupmanagernotupgradedOnSequenceBatches)
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
func (it *PolygonrollupmanagernotupgradedOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedOnSequenceBatches represents a OnSequenceBatches event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedOnSequenceBatchesIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedOnSequenceBatches)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseOnSequenceBatches(log types.Log) (*PolygonrollupmanagernotupgradedOnSequenceBatches, error) {
	event := new(PolygonrollupmanagernotupgradedOnSequenceBatches)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleAdminChangedIterator struct {
	Event *PolygonrollupmanagernotupgradedRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedRoleAdminChanged)
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
		it.Event = new(PolygonrollupmanagernotupgradedRoleAdminChanged)
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
func (it *PolygonrollupmanagernotupgradedRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedRoleAdminChanged represents a RoleAdminChanged event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PolygonrollupmanagernotupgradedRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedRoleAdminChangedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedRoleAdminChanged)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseRoleAdminChanged(log types.Log) (*PolygonrollupmanagernotupgradedRoleAdminChanged, error) {
	event := new(PolygonrollupmanagernotupgradedRoleAdminChanged)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleGrantedIterator struct {
	Event *PolygonrollupmanagernotupgradedRoleGranted // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedRoleGranted)
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
		it.Event = new(PolygonrollupmanagernotupgradedRoleGranted)
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
func (it *PolygonrollupmanagernotupgradedRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedRoleGranted represents a RoleGranted event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagernotupgradedRoleGrantedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedRoleGrantedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedRoleGranted)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseRoleGranted(log types.Log) (*PolygonrollupmanagernotupgradedRoleGranted, error) {
	event := new(PolygonrollupmanagernotupgradedRoleGranted)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleRevokedIterator struct {
	Event *PolygonrollupmanagernotupgradedRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedRoleRevoked)
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
		it.Event = new(PolygonrollupmanagernotupgradedRoleRevoked)
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
func (it *PolygonrollupmanagernotupgradedRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedRoleRevoked represents a RoleRevoked event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PolygonrollupmanagernotupgradedRoleRevokedIterator, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedRoleRevokedIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedRoleRevoked)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseRoleRevoked(log types.Log) (*PolygonrollupmanagernotupgradedRoleRevoked, error) {
	event := new(PolygonrollupmanagernotupgradedRoleRevoked)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRollbackBatchesIterator struct {
	Event *PolygonrollupmanagernotupgradedRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedRollbackBatches)
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
		it.Event = new(PolygonrollupmanagernotupgradedRollbackBatches)
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
func (it *PolygonrollupmanagernotupgradedRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedRollbackBatches represents a RollbackBatches event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*PolygonrollupmanagernotupgradedRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedRollbackBatchesIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedRollbackBatches)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseRollbackBatches(log types.Log) (*PolygonrollupmanagernotupgradedRollbackBatches, error) {
	event := new(PolygonrollupmanagernotupgradedRollbackBatches)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedSetBatchFeeIterator struct {
	Event *PolygonrollupmanagernotupgradedSetBatchFee // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedSetBatchFee)
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
		it.Event = new(PolygonrollupmanagernotupgradedSetBatchFee)
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
func (it *PolygonrollupmanagernotupgradedSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedSetBatchFee represents a SetBatchFee event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedSetBatchFeeIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedSetBatchFeeIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedSetBatchFee)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseSetBatchFee(log types.Log) (*PolygonrollupmanagernotupgradedSetBatchFee, error) {
	event := new(PolygonrollupmanagernotupgradedSetBatchFee)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagernotupgradedSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedSetTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagernotupgradedSetTrustedAggregator)
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
func (it *PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedSetTrustedAggregator represents a SetTrustedAggregator event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedSetTrustedAggregatorIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedSetTrustedAggregator)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseSetTrustedAggregator(log types.Log) (*PolygonrollupmanagernotupgradedSetTrustedAggregator, error) {
	event := new(PolygonrollupmanagernotupgradedSetTrustedAggregator)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedUpdateRollupIterator struct {
	Event *PolygonrollupmanagernotupgradedUpdateRollup // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedUpdateRollup)
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
		it.Event = new(PolygonrollupmanagernotupgradedUpdateRollup)
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
func (it *PolygonrollupmanagernotupgradedUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedUpdateRollup represents a UpdateRollup event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedUpdateRollupIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedUpdateRollup)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseUpdateRollup(log types.Log) (*PolygonrollupmanagernotupgradedUpdateRollup, error) {
	event := new(PolygonrollupmanagernotupgradedUpdateRollup)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator is returned from FilterUpdateRollupManagerVersion and is used to iterate over the raw logs and unpacked data for UpdateRollupManagerVersion events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator struct {
	Event *PolygonrollupmanagernotupgradedUpdateRollupManagerVersion // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedUpdateRollupManagerVersion)
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
		it.Event = new(PolygonrollupmanagernotupgradedUpdateRollupManagerVersion)
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
func (it *PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedUpdateRollupManagerVersion represents a UpdateRollupManagerVersion event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedUpdateRollupManagerVersion struct {
	RollupManagerVersion string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupManagerVersion is a free log retrieval operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterUpdateRollupManagerVersion(opts *bind.FilterOpts) (*PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedUpdateRollupManagerVersionIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "UpdateRollupManagerVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupManagerVersion is a free log subscription operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchUpdateRollupManagerVersion(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedUpdateRollupManagerVersion) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedUpdateRollupManagerVersion)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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

// ParseUpdateRollupManagerVersion is a log parse operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseUpdateRollupManagerVersion(log types.Log) (*PolygonrollupmanagernotupgradedUpdateRollupManagerVersion, error) {
	event := new(PolygonrollupmanagernotupgradedUpdateRollupManagerVersion)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator struct {
	Event *PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator)
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
		it.Event = new(PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator)
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
func (it *PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator struct {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregatorIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator, error) {
	event := new(PolygonrollupmanagernotupgradedVerifyBatchesTrustedAggregator)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

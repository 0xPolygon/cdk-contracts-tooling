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

// PolygonrollupmanagermockMetaData contains all meta data concerning the Polygonrollupmanagermock contract.
var PolygonrollupmanagermockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"CreateNewAggchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"}],\"name\":\"VerifyPessimisticStateTransition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"exposed_checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"}],\"name\":\"initializeMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"localExitRoots\",\"type\":\"bytes32[]\"}],\"name\":\"prepareMockCalculateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataDeserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2Deserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"setRollupData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516161fe3803806161fe83398101604081905261002f9161019a565b838383836001600160a01b038416158061005057506001600160a01b038316155b8061006257506001600160a01b038216155b8061007457506001600160a01b038116155b15610092576040516342eda64b60e11b815260040160405180910390fd5b6001600160a01b0380851660805283811660c05282811660a052811660e0526100b96100c6565b50505050505050506101f6565b5f54610100900460ff16156101315760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610181575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610197575f5ffd5b50565b5f5f5f5f608085870312156101ad575f5ffd5b84516101b881610183565b60208601519094506101c981610183565b60408601519093506101da81610183565b60608601519092506101eb81610183565b939692955090935050565b60805160a05160c05160e051615f9d6102615f395f818161083c01526113b401525f8181610985015281816123f0015261375301525f81816107fd01528181612da4015261385101525f81816108c801528181610bca015281816112d301526114c10152615f9d5ff3fe608060405234801561000f575f5ffd5b506004361061033b575f3560e01c80638fd88cc2116101b3578063d02103ca116100f3578063dde0ff771161009e578063e4f3d8f911610079578063e4f3d8f9146109a7578063e80e503014610a57578063f4e9267514610a6a578063f9c4c2ae14610a7a575f5ffd5b8063dde0ff7714610953578063dfdb8c5e1461096d578063e46761c414610980575f5ffd5b8063d8905812116100ce578063d890581214610910578063dbc1697614610938578063dd0464b914610940575f5ffd5b8063d02103ca146108c3578063d5073f6f146108ea578063d547741f146108fd575f5ffd5b8063a2967d991161015e578063abcb519811610139578063abcb51981461085e578063c1acbc3414610871578063c4c928c21461088b578063ceee281d1461089e575f5ffd5b8063a2967d99146107f0578063a3c573eb146107f8578063ab0475cf14610837575f5ffd5b806399f5634e1161018e57806399f5634e146107ce5780639a908e73146107d6578063a217fddf146107e9575f5ffd5b80638fd88cc21461077057806391d148541461078357806397d289a3146107bb575f5ffd5b8063604691691161027e57806370603909116102295780637975fcfe116102045780637975fcfe146107105780637fb6e76a146107305780638129fc1c146107555780638f698ec51461075d575f5ffd5b8063706039091461062f5780637222020f146106dd57806374d9c244146106f0575f5ffd5b8063680658a111610259578063680658a1146105d85780636c0a51b7146106095780636c7668771461061c575f5ffd5b8063604691691461054257806362d87e661461054a57806365c0504d1461055d575f5ffd5b8063248a9ca3116102e957806330c27dde116102c457806330c27dde146104d457806336568abe146104e7578063477fa270146104fa57806355a71ee014610502575f5ffd5b8063248a9ca3146103e157806325280169146104115780632f2ff15d146104c1575f5ffd5b806315064c961161031957806315064c96146103975780631796a1ae146103b45780632072f6c5146103d9575f5ffd5b8063066ec0121461033f57806311f6b2871461036f5780631489ed1014610382575b5f5ffd5b608454610352906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61035261037d366004614589565b610a9a565b6103956103903660046145d3565b610ac9565b005b606f546103a49060ff1681565b6040519015158152602001610366565b607e546103c49063ffffffff1681565b60405163ffffffff9091168152602001610366565b610395610caa565b6104036103ef366004614661565b5f9081526034602052604090206001015490565b604051908152602001610366565b61048e61041f366004614678565b60408051606080820183525f808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b03908116918301919091529282015190921690820152606001610366565b6103956104cf3660046146a9565b610d7c565b608754610352906001600160401b031681565b6103956104f53660046146a9565b610da5565b608654610403565b610403610510366004614678565b63ffffffff82165f9081526081602090815260408083206001600160401b038516845260020190915290205492915050565b610403610ddc565b6103a4610558366004614661565b610df1565b6105c561056b366004614589565b607f6020525f908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b604051610366979695949392919061470b565b6103956105e636600461475d565b63ffffffff9092165f908152608160205260409020600581019190915560080155565b61039561061736600461478d565b610dfb565b61039561062a36600461482a565b6111e5565b6106c561063d366004614589565b63ffffffff165f9081526081602052604090208054600182015460058301546006840154600785015460088601546009909601546001600160a01b0380871698600160a01b978890046001600160401b03908116999288169890970487169680861695600160401b908190048216958281169591810490921693600160801b90920460ff1692565b6040516103669c9b9a999897969594939291906148ca565b6103956106eb366004614589565b6116bc565b6107036106fe366004614589565b6117ad565b604051610366919061495b565b61072361071e366004614a69565b6118fc565b6040516103669190614af2565b6103c461073e366004614b04565b60836020525f908152604090205463ffffffff1681565b61039561192c565b61039561076b366004614b61565b611a59565b61039561077e366004614c00565b611aca565b6103a46107913660046146a9565b5f9182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6103956107c9366004614c9b565b611e4f565b6104036123cf565b6103526107e4366004614cf4565b6124a8565b6104035f81565b61040361269e565b61081f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610366565b61081f7f000000000000000000000000000000000000000000000000000000000000000081565b61039561086c366004614d2a565b612964565b60845461035290600160801b90046001600160401b031681565b610395610899366004614dd0565b612c4f565b6103c46108ac366004614dfb565b60826020525f908152604090205463ffffffff1681565b61081f7f000000000000000000000000000000000000000000000000000000000000000081565b6103956108f8366004614661565b612c8a565b61039561090b3660046146a9565b612d28565b610723604051806040016040528060098152602001680616c2d76302e332e360bc1b81525081565b610395612d4c565b61072361094e366004614e16565b612e12565b60845461035290600160401b90046001600160401b031681565b61039561097b366004614e81565b612e39565b61081f7f000000000000000000000000000000000000000000000000000000000000000081565b610a3f6109b5366004614589565b63ffffffff165f90815260816020526040902080546001820154600583015460068401546007909401546001600160a01b0380851696600160a01b958690046001600160401b03908116979286169690950485169480831693600160401b808504831694600160801b808204851695600160c01b9092048516948085169493840416920460ff1690565b6040516103669c9b9a99989796959493929190614eab565b610395610a65366004614f51565b612ff2565b6080546103c49063ffffffff1681565b610a8d610a88366004614589565b6133a1565b6040516103669190614fd2565b63ffffffff81165f90815260816020526040812060060154600160401b90046001600160401b03165b92915050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4610af3816134f7565b6001600160401b03881615610b1b5760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610b5057610b506146d7565b14610b6e576040516390db0d0760e01b815260040160405180910390fd5b610b7d81898989898989613501565b6006810180546fffffffffffffffff00000000000000001916600160401b6001600160401b038a16908102919091179091555f9081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610bff61269e565b6040518263ffffffff1660e01b8152600401610c1d91815260200190565b5f604051808303815f87803b158015610c34575f5ffd5b505af1158015610c46573d5f5f3e3d5ffd5b5050604080516001600160401b038b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b335f9081527f8875b94af5657a2903def9906d67a3f42d8a836d24b5602c00f00fc855339fcd602052604090205460ff16610d7257608454600160801b90046001600160401b03161580610d2757506084544290610d1c9062093a8090600160801b90046001600160401b0316615116565b6001600160401b0316115b80610d5457506087544290610d499062093a80906001600160401b0316615116565b6001600160401b0316115b15610d725760405163692baaad60e11b815260040160405180910390fd5b610d7a61384f565b565b5f82815260346020526040902060010154610d96816134f7565b610da083836138c5565b505050565b6001600160a01b0381163314610dce57604051630b4ad1cd60e31b815260040160405180910390fd5b610dd88282613948565b5050565b5f6086546064610dec9190615135565b905090565b5f610ac3826139c9565b5f54600490610100900460ff16158015610e1b57505f5460ff8083169116105b610e835760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805461ffff191660ff83161761010017905567016345785d8a0000608655610eaa613a4e565b610ed47f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd486613ab8565b610ede5f84613ab8565b610f087fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59084613ab8565b610f327f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e84613ab8565b610f5c7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac84613ab8565b610f867fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd85613ab8565b610fb07fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd0885613ab8565b610fda7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f485613ab8565b6110047fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db185613ab8565b61104e7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd47f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0613ac2565b6110787f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f085613ab8565b6110a27f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb85613ab8565b6110ec7f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e7f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff285951613ac2565b6111167f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e83613ab8565b6111407f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff28595183613ab8565b61114a5f33613ab8565b60408051808201825260098152680616c2d76302e332e360bc1b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c29161119691614af2565b60405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd461120f816134f7565b611217613b0c565b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff16600281111561124c5761124c6146d7565b0361126a57604051635b6602b760e01b815260040160405180910390fd5b60016007820154600160801b900460ff16600281111561128c5761128c6146d7565b14801561129857508215155b156112b657604051630ded782d60e01b815260040160405180910390fd5b60405163ef4eeb3560e01b815263ffffffff8a1660048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ef4eeb3590602401602060405180830381865afa158015611320573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611344919061514c565b9050806113645760405163a60721e160e01b815260040160405180910390fd5b5f6113748c84848d8d8b8b613b79565b905060026007840154600160801b900460ff166002811115611398576113986146d7565b0361141e5760405163a48fd37760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a48fd377906113ed9084908c908c9060040161518b565b5f6040518083038186803b158015611403575f5ffd5b505afa158015611415573d5f5f3e3d5ffd5b50505050611486565b6001830154600984015460405163020a49e360e51b81526001600160a01b03909216916341493c60916114599185908d908d906004016151ba565b5f6040518083038186803b15801561146f575f5ffd5b505afa158015611481573d5f5f3e3d5ffd5b505050505b6084805467ffffffffffffffff60801b1916600160801b426001600160401b031602179055600583018054908b9055600884018054908b90557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d6114f661269e565b6040518263ffffffff1660e01b815260040161151491815260200190565b5f604051808303815f87803b15801561152b575f5ffd5b505af115801561153d573d5f5f3e3d5ffd5b505050505f8c9050336001600160a01b03168f63ffffffff167fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d35f5f5f1b856040516115a7939291906001600160401b039390931683526020830191909152604082015260600190565b60405180910390a3336001600160a01b03168f63ffffffff167fdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711848f87868b604051611615959493929190948552602085019390935260408401919091526060830152608082015260a00190565b60405180910390a360026007870154600160801b900460ff16600281111561163f5761163f6146d7565b036116a3578554604051639ee4afa360e01b81526001600160a01b0390911690639ee4afa390611675908c908c906004016151e5565b5f604051808303815f87803b15801561168c575f5ffd5b505af115801561169e573d5f5f3e3d5ffd5b505050505b5050505050506116b1613d4e565b505050505050505050565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd6116e6816134f7565b63ffffffff821615806117045750607e5463ffffffff908116908316115b1561172257604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001810154600160e81b900460ff161561176357604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b61181060408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082019081526020015f81526020015f81525090565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b0380821686526001600160401b03600160a01b928390048116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b90819004821660c0850152600783015480831660e086015290810490911661010084015261012083019060ff600160801b9091041660028111156118cc576118cc6146d7565b908160028111156118df576118df6146d7565b905250600881015461014083015260090154610160820152919050565b63ffffffff86165f908152608160205260409020606090611921908787878787613d78565b979650505050505050565b5f54600490610100900460ff1615801561194c57505f5460ff8083169116105b6119af5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610e7a565b5f805461ffff191660ff83161761010017905560408051808201825260098152680616c2d76302e332e360bc1b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c291611a0e91614af2565b60405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b80516080805463ffffffff191663ffffffff9092169190911790555f5b8151811015610dd857818181518110611a9157611a916151f8565b602002602001015160815f836001611aa9919061520c565b63ffffffff16815260208101919091526040015f2060050155600101611a76565b335f9081527ff14f5a8ad59d90593602e905b358229bff5ceea677d5bf0f5a1701793550a9a6602052604090205460ff16158015611b795750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b49573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b6d919061521f565b6001600160a01b031614155b15611b9757604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b0382165f9081526082602052604081205463ffffffff1690819003611bd6576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff166002811115611c0b57611c0b6146d7565b14611c29576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b0390811690841681111580611c61575060068201546001600160401b03600160401b9091048116908516105b15611c7f5760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b031614611d20576001600160401b038082165f908152600385016020526040902060010154600160401b90048116908616811015611ce457604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091165f908152600384016020526040812090815560010180546fffffffffffffffffffffffffffffffff19169055611c81565b60068301805467ffffffffffffffff19166001600160401b038716179055611d48858361523a565b608480545f90611d629084906001600160401b031661523a565b82546101009290920a6001600160401b0381810219909316918316021790915586165f8181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece906044015f604051808303815f87803b158015611ddb575f5ffd5b505af1158015611ded573d5f5f3e3d5ffd5b505050506001600160401b0385165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a3505050505050565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08611e79816134f7565b63ffffffff84161580611e975750607e5463ffffffff908116908516115b15611eb557604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615611ef657604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b0385161115611f2457604051634c753f5760e01b815260040160405180910390fd5b6001600160401b0384165f9081526083602052604090205463ffffffff1615611f60576040516337c8fe0960e11b815260040160405180910390fd5b608080545f91908290611f789063ffffffff16615259565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f80825260208201928390529394506001600160a01b03909216913091611fc390614569565b611fcf9392919061527d565b604051809103905ff080158015611fe8573d5f5f3e3d5ffd5b506001600160401b0387165f818152608360209081526040808320805463ffffffff1990811663ffffffff8a81169182179093556001600160a01b038816808752608286528487208054909316821790925585526081909352922080546001600160e01b031916909117600160a01b9093029290921782556007820180546fffffffffffffffff0000000000000000198116928c16600160401b02928317825560018801549495509293600160e01b900460ff1692909170ffffffffffffffffff00000000000000001990911660ff60801b1990911617600160801b8360028111156120d6576120d66146d7565b021790555060026001850154600160e01b900460ff1660028111156120fd576120fd6146d7565b03612229578263ffffffff167f144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b89848a88600101601c9054906101000a900460ff166002811115612150576121506146d7565b8b6040516121629594939291906152b6565b60405180910390a26040805163ffffffff8a811682526001600160a01b0385811660208401526001600160401b038b168385015260608301529151918516917f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6419181900360800190a260405163439fab9160e01b81526001600160a01b0383169063439fab91906121f7908990600401614af2565b5f604051808303815f87803b15801561220e575f5ffd5b505af1158015612220573d5f5f3e3d5ffd5b505050506123c5565b6001848101805491830180547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff8116600160a01b948590046001600160401b0316909402938417825591546001600160a01b03166001600160e01b03199092166001600160a01b0319909316929092171790556002808501545f8080529183016020908152604083209190915560038601546009840155875182918291829182916122db91908d018101908d01615348565b945094509450945094508763ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418e898f87604051612354949392919063ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b60405180910390a2604051633892b81160e11b81526001600160a01b0388169063712570229061239290889088908d908990899089906004016153e4565b5f604051808303815f87803b1580156123a9575f5ffd5b505af11580156123bb573d5f5f3e3d5ffd5b5050505050505050505b5050505050505050565b6040516370a0823160e01b81523060048201525f9081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015612435573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612459919061514c565b6084549091505f9061247d906001600160401b03600160401b82048116911661523a565b6001600160401b03169050805f03612497575f9250505090565b6124a18183615462565b9250505090565b606f545f9060ff16156124ce57604051630bc011ff60e21b815260040160405180910390fd5b335f9081526082602052604081205463ffffffff1690819003612504576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03165f0361252d57604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff166002811115612562576125626146d7565b14612580576040516390db0d0760e01b815260040160405180910390fd5b608480548691905f9061259d9084906001600160401b0316615116565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690505f6125d08783615116565b6006840180546001600160401b0383811667ffffffffffffffff199092168217909255604080516060810182528a815242841660208083019182528886168385019081525f86815260038c018352859020935184559151600193909301805492518716600160401b026fffffffffffffffffffffffffffffffff199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b6080545f9063ffffffff168082036126b757505f919050565b5f816001600160401b038111156126d0576126d0614b1d565b6040519080825280602002602001820160405280156126f9578160200160208202803683370190505b5090505f5b828110156127565760815f61271483600161520c565b63ffffffff1663ffffffff1681526020019081526020015f2060050154828281518110612743576127436151f8565b60209081029190910101526001016126fe565b505f60205b83600114612913575f61276f600286615475565b61277a600287615462565b612784919061520c565b90505f816001600160401b0381111561279f5761279f614b1d565b6040519080825280602002602001820160405280156127c8578160200160208202803683370190505b5090505f5b828110156128e3576127e0600184615488565b811480156127f857506127f4600288615475565b6001145b15612856576128338661280c836002615135565b8151811061281c5761281c6151f8565b6020026020010151865f9182526020526040902090565b828281518110612845576128456151f8565b6020026020010181815250506128db565b6128bc86612865836002615135565b81518110612875576128756151f8565b60200260200101518783600261288b9190615135565b61289690600161520c565b815181106128a6576128a66151f8565b60200260200101515f9182526020526040902090565b8282815181106128ce576128ce6151f8565b6020026020010181815250505b6001016127cd565b508094508195506128fd84855f9182526020526040902090565b9350826129098161549b565b935050505061275b565b5f835f81518110612926576129266151f8565b602002602001015190505f5f90505b8281101561295a575f9182526020849052604080832085845292209350600101612935565b5095945050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59061298e816134f7565b607e80545f919082906129a69063ffffffff16615259565b91906101000a81548163ffffffff021916908363ffffffff16021790559050600160028111156129d8576129d86146d7565b8660028111156129ea576129ea6146d7565b03612a13578415612a0e576040516363d722e760e01b815260040160405180910390fd5b612acd565b6002866002811115612a2757612a276146d7565b03612a7d576001600160a01b038816151580612a4b57506001600160401b03871615155b80612a5557508415155b80612a5f57508215155b15612a0e576040516363d722e760e01b815260040160405180910390fd5b5f866002811115612a9057612a906146d7565b03612ab4578215612a0e576040516363d722e760e01b815260040160405180910390fd5b6040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876002811115612b1757612b176146d7565b81525f602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b031990911617825591850151600182018054948701516001600160401b0316600160a01b026001600160e01b031990951691909316179290921780825592840151919260ff60e01b1916600160e01b836002811115612bba57612bba6146d7565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a90612c3c908c908c908c908c908c908c908c906154b0565b60405180910390a2505050505050505050565b7f66156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac612c79816134f7565b612c84848484613e7b565b50505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb612cb4816134f7565b683635c9adc5dea00000821180612cce5750633b9aca0082105b15612cec57604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b5f82815260346020526040902060010154612d42816134f7565b610da08383613948565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f4612d76816134f7565b6087805467ffffffffffffffff1916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc16976916004808301925f92919082900301818387803b158015612df1575f5ffd5b505af1158015612e03573d5f5f3e3d5ffd5b50505050612e0f6141bc565b50565b63ffffffff86165f9081526081602052604090206060906119219088908888888888613b79565b336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e7f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ea3919061521f565b6001600160a01b031614612eca5760405163696072e960e01b815260040160405180910390fd5b6001600160a01b0382165f9081526082602090815260408083205463ffffffff1683526081909152902060068101546001600160401b03808216600160401b9092041614612f2b5760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b031610612f6657604051633e37e23360e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020526040902060010154600160e01b900460ff166002811115612f9957612f996146d7565b6007820154600160801b900460ff166002811115612fb957612fb96146d7565b14612fd757604051635aa0d5f160e11b815260040160405180910390fd5b604080515f815260208101909152610da09084908490613e7b565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e61301c816134f7565b6001600160401b0386165f9081526083602052604090205463ffffffff1615613058576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b038716111561308657604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b0389165f9081526082602052604090205463ffffffff16156130c257604051630d409b9360e41b815260040160405180910390fd5b608080545f919082906130da9063ffffffff16615259565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f896001600160401b03166001600160401b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8363ffffffff1663ffffffff1681526020019081526020015f2090508a815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550888160010160146101000a8154816001600160401b0302191690836001600160401b0316021790555089816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555087815f0160146101000a8154816001600160401b0302191690836001600160401b03160217905550858160070160106101000a81548160ff02191690836002811115613261576132616146d7565b0217905550600186600281111561327a5761327a6146d7565b036132995760098101859055600881018490556005810187905561334a565b60028660028111156132ad576132ad6146d7565b0361330c576001600160a01b038a161515806132d157506001600160401b03891615155b806132db57508415155b156132f957604051636fc5557f60e01b815260040160405180910390fd5b600881018490556005810187905561334a565b8415158061331957508315155b1561333757604051636fc5557f60e01b815260040160405180910390fd5b5f80805260028201602052604090208790555b8163ffffffff167f4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e8a8d8b8a5f8b8b60405161338c9796959493929190615516565b60405180910390a25050505050505050505050565b61340560408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b03908116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b808204831660c0860152600160801b808304841660e0870152600160c01b909204831661010086015260078401548084166101208701529081049092166101408501526101608401910460ff1660028111156134da576134da6146d7565b908160028111156134ed576134ed6146d7565b8152505050919050565b612e0f8133614213565b5f5f61351f89600601546001600160401b03600160401b9091041690565b60078a01549091506001600160401b0390811690891610156135545760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b0388165f90815260028a01602052604090205491508161358e576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b031611156135c157604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b0316116135f35760405163b9b18f5760e01b815260040160405180910390fd5b5f6136028a8a8a8a878b613d78565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516136369190615571565b602060405180830381855afa158015613651573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613674919061514c565b61367e9190615475565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a916136c091899190600401615587565b602060405180830381865afa1580156136db573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136ff91906155c3565b61371c576040516309bde33960e01b815260040160405180910390fd5b5f613727848b61523a565b905061377a87826001600160401b031661373f6123cf565b6137499190615135565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190614255565b80608460088282829054906101000a90046001600160401b031661379e9190615116565b82546101009290920a6001600160401b038181021990931691831602179091556084805467ffffffffffffffff60801b1916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064015f604051808303815f87803b15801561382b575f5ffd5b505af115801561383d573d5f5f3e3d5ffd5b50505050505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156138a7575f5ffd5b505af11580156138b9573d5f5f3e3d5ffd5b50505050610d7a6142bc565b5f8281526034602090815260408083206001600160a01b038516845290915290205460ff16610dd8575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b5f8281526034602090815260408083206001600160a01b038516845290915290205460ff1615610dd8575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f67ffffffff000000016001600160401b0383161080156139fe575067ffffffff00000001604083901c6001600160401b0316105b8015613a1e575067ffffffff00000001608083901c6001600160401b0316105b8015613a35575067ffffffff0000000160c083901c105b15613a4257506001919050565b505f919050565b919050565b5f54610100900460ff16610d7a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610e7a565b610dd882826138c5565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005c15613b4c57604051633ee5aeb560e01b815260040160405180910390fd5b610d7a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005b90614317565b606060026007880154600160801b900460ff166002811115613b9d57613b9d6146d7565b03613c78578654604051631a957d9b60e21b81525f916001600160a01b031690636a55f66c90613bd390879087906004016151e5565b602060405180830381865afa158015613bee573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c12919061514c565b600589015460088a0154604080516020810193909352820152606081018990526001600160e01b031960e08c901b1660808201526084810182905260a4810188905260c4810187905290915060e401604051602081830303815290604052915050611921565b865460408051632b47b7cd60e21b815290515f926001600160a01b03169163ad1edf349160048083019260209291908290030181865afa158015613cbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ce2919061514c565b600589015460088a0154604080516020810193909352820152606081018990526001600160e01b031960e08c901b1660808201526084810182905260a4810188905260c4810187905290915060e401604051602081830303815290604052915050979650505050505050565b610d7a5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00613b73565b6001600160401b038086165f8181526003890160205260408082205493881682529020546060929115801590613dac575081155b15613dca5760405163340c614f60e11b815260040160405180910390fd5b80613de8576040516366385b5160e01b815260040160405180910390fd5b613df1846139c9565b613e0e576040516305dae44f60e21b815260040160405180910390fd5b3385838a8c5f0160149054906101000a90046001600160401b03168d60010160149054906101000a90046001600160401b031689878d8f604051602001613e5e9a999897969594939291906155e2565b604051602081830303815290604052925050509695505050505050565b63ffffffff82161580613e995750607e5463ffffffff908116908316115b15613eb757604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526082602052604081205463ffffffff1690819003613ef6576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181165f908152608160205260409020600781015490918516600160401b9091046001600160401b031603613f4357604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615613f8457604051633b8d3d9960e01b815260040160405180910390fd5b60026001820154600160e01b900460ff166002811115613fa657613fa66146d7565b14158015613ff257506001810154600160e01b900460ff166002811115613fcf57613fcf6146d7565b6007830154600160801b900460ff166002811115613fef57613fef6146d7565b14155b1561401057604051635aa0d5f160e11b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825582546001600160401b03600160a01b9182900416026001600160e01b03199091169093179290921790915560038201546009840155600783018054600160401b63ffffffff8916026fffffffffffffffff000000000000000019821681178355925460ff600160e01b909104169260ff60801b191670ffffffffffffffffff00000000000000001990911617600160801b8360028111156140da576140da6146d7565b02179055505f6140e984610a9a565b60078401805467ffffffffffffffff19166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef2869261413b9216908990600401615696565b5f604051808303815f87803b158015614152575f5ffd5b505af1158015614164573d5f5f3e3d5ffd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff166141df57604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f8281526034602090815260408083206001600160a01b038516845290915290205460ff16610dd857604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1663a9059cbb60e01b179052610da090849061431e565b606f5460ff16156142e057604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b80825d5050565b5f614372826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166143ef9092919063ffffffff16565b805190915015610da0578080602001905181019061439091906155c3565b610da05760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610e7a565b60606143fd84845f85614405565b949350505050565b6060824710156144665760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610e7a565b5f5f866001600160a01b031685876040516144819190615571565b5f6040518083038185875af1925050503d805f81146144bb576040519150601f19603f3d011682016040523d82523d5f602084013e6144c0565b606091505b5091509150611921878383876060831561453a5782515f03614533576001600160a01b0385163b6145335760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610e7a565b50816143fd565b6143fd838381511561454f5781518083602001fd5b8060405162461bcd60e51b8152600401610e7a9190614af2565b6108b0806156b883390190565b803563ffffffff81168114613a49575f5ffd5b5f60208284031215614599575f5ffd5b6145a282614576565b9392505050565b80356001600160401b0381168114613a49575f5ffd5b6001600160a01b0381168114612e0f575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b0312156145eb575f5ffd5b6145f489614576565b975061460260208a016145a9565b965061461060408a016145a9565b955061461e60608a016145a9565b94506080890135935060a0890135925060c089013561463c816145bf565b91506103e089018a101561464e575f5ffd5b60e0890190509295985092959890939650565b5f60208284031215614671575f5ffd5b5035919050565b5f5f60408385031215614689575f5ffd5b61469283614576565b91506146a0602084016145a9565b90509250929050565b5f5f604083850312156146ba575f5ffd5b8235915060208301356146cc816145bf565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b6003811061470757634e487b7160e01b5f52602160045260245ffd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e0810161473f60608301876146eb565b931515608082015260a081019290925260c090910152949350505050565b5f5f5f6060848603121561476f575f5ffd5b61477884614576565b95602085013595506040909401359392505050565b5f5f5f5f608085870312156147a0575f5ffd5b84356147ab816145bf565b935060208501356147bb816145bf565b925060408501356147cb816145bf565b915060608501356147db816145bf565b939692955090935050565b5f5f83601f8401126147f6575f5ffd5b5081356001600160401b0381111561480c575f5ffd5b602083019150836020828501011115614823575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b031215614841575f5ffd5b61484a89614576565b975061485860208a01614576565b9650604089013595506060890135945060808901356001600160401b03811115614880575f5ffd5b61488c8b828c016147e6565b90955093505060a08901356001600160401b038111156148aa575f5ffd5b6148b68b828c016147e6565b999c989b5096995094979396929594505050565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b03861661010083015261493f6101208301866146eb565b61014082019390935261016001529a9950505050505050505050565b81516001600160a01b031681526101808101602083015161498760208401826001600160401b03169052565b5060408301516149a260408401826001600160a01b03169052565b5060608301516149bd60608401826001600160401b03169052565b506080830151608083015260a08301516149e260a08401826001600160401b03169052565b5060c08301516149fd60c08401826001600160401b03169052565b5060e0830151614a1860e08401826001600160401b03169052565b50610100830151614a356101008401826001600160401b03169052565b50610120830151614a4a6101208401826146eb565b5061014083015161014083015261016083015161016083015292915050565b5f5f5f5f5f5f60c08789031215614a7e575f5ffd5b614a8787614576565b9550614a95602088016145a9565b9450614aa3604088016145a9565b959894975094956060810135955060808101359460a0909101359350915050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6145a26020830184614ac4565b5f60208284031215614b14575f5ffd5b6145a2826145a9565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715614b5957614b59614b1d565b604052919050565b5f60208284031215614b71575f5ffd5b81356001600160401b03811115614b86575f5ffd5b8201601f81018413614b96575f5ffd5b80356001600160401b03811115614baf57614baf614b1d565b8060051b614bbf60208201614b31565b91825260208184018101929081019087841115614bda575f5ffd5b6020850194505b8385101561192157843580835260209586019590935090910190614be1565b5f5f60408385031215614c11575f5ffd5b8235614692816145bf565b5f6001600160401b03821115614c3457614c34614b1d565b50601f01601f191660200190565b5f614c54614c4f84614c1c565b614b31565b9050828152838383011115614c67575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112614c8c575f5ffd5b6145a283833560208501614c42565b5f5f5f60608486031215614cad575f5ffd5b614cb684614576565b9250614cc4602085016145a9565b915060408401356001600160401b03811115614cde575f5ffd5b614cea86828701614c7d565b9150509250925092565b5f5f60408385031215614d05575f5ffd5b614d0e836145a9565b946020939093013593505050565b803560038110613a49575f5ffd5b5f5f5f5f5f5f5f60e0888a031215614d40575f5ffd5b8735614d4b816145bf565b96506020880135614d5b816145bf565b9550614d69604089016145a9565b9450614d7760608901614d1c565b93506080880135925060a08801356001600160401b03811115614d98575f5ffd5b8801601f81018a13614da8575f5ffd5b614db78a823560208401614c42565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614de2575f5ffd5b8335614ded816145bf565b9250614cc460208501614576565b5f60208284031215614e0b575f5ffd5b81356145a2816145bf565b5f5f5f5f5f5f60a08789031215614e2b575f5ffd5b614e3487614576565b955060208701359450604087013593506060870135925060808701356001600160401b03811115614e63575f5ffd5b614e6f89828a016147e6565b979a9699509497509295939492505050565b5f5f60408385031215614e92575f5ffd5b8235614e9d816145bf565b91506146a060208401614576565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b0386166101008301526001600160401b0385166101208301526001600160401b038416610140830152614f406101608301846146eb565b9d9c50505050505050505050505050565b5f5f5f5f5f5f5f5f610100898b031215614f69575f5ffd5b8835614f74816145bf565b97506020890135614f84816145bf565b9650614f9260408a016145a9565b9550614fa060608a016145a9565b945060808901359350614fb560a08a01614d1c565b979a969950949793969295929450505060c08201359160e0013590565b81516001600160a01b0316815261018081016020830151614ffe60208401826001600160401b03169052565b50604083015161501960408401826001600160a01b03169052565b50606083015161503460608401826001600160401b03169052565b506080830151608083015260a083015161505960a08401826001600160401b03169052565b5060c083015161507460c08401826001600160401b03169052565b5060e083015161508f60e08401826001600160401b03169052565b506101008301516150ac6101008401826001600160401b03169052565b506101208301516150c96101208401826001600160401b03169052565b506101408301516150e66101408401826001600160401b03169052565b506101608301516150fb6101608401826146eb565b5092915050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190811115610ac357610ac3615102565b8082028115828204841417610ac357610ac3615102565b5f6020828403121561515c575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f61519d6040830186614ac4565b82810360208401526151b0818587615163565b9695505050505050565b848152606060208201525f6151d26060830186614ac4565b8281036040840152611921818587615163565b602081525f6143fd602083018486615163565b634e487b7160e01b5f52603260045260245ffd5b80820180821115610ac357610ac3615102565b5f6020828403121561522f575f5ffd5b81516145a2816145bf565b6001600160401b038281168282160390811115610ac357610ac3615102565b5f63ffffffff821663ffffffff810361527457615274615102565b60010192915050565b6001600160a01b03841681526001600160a01b0383166020820152606060408201525f6152ad6060830184614ac4565b95945050505050565b63ffffffff861681526001600160a01b03851660208201526001600160401b038416604082015260ff8316606082015260a060808201525f61192160a0830184614ac4565b5f82601f83011261530a575f5ffd5b8151615318614c4f82614c1c565b81815284602083860101111561532c575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a0868803121561535c575f5ffd5b8551615367816145bf565b6020870151909550615378816145bf565b6040870151909450615389816145bf565b60608701519093506001600160401b038111156153a4575f5ffd5b6153b0888289016152fb565b92505060808601516001600160401b038111156153cb575f5ffd5b6153d7888289016152fb565b9150509295509295909350565b6001600160a01b03871681526001600160a01b038616602082015263ffffffff851660408201526001600160a01b038416606082015260c060808201525f61542f60c0830185614ac4565b82810360a08401526154418185614ac4565b9998505050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f826154705761547061544e565b500490565b5f826154835761548361544e565b500690565b81810381811115610ac357610ac3615102565b5f816154a9576154a9615102565b505f190190565b6001600160a01b03881681526001600160a01b03871660208201526001600160401b03861660408201526154e760608201866146eb565b83608082015260e060a08201525f61550260e0830185614ac4565b90508260c083015298975050505050505050565b6001600160401b0388811682526001600160a01b03881660208301528616604082015260e0810161554a60608301876146eb565b6001600160401b03851660808301528360a08301528260c083015298975050505050505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b60018110156155b957815183526020928301929091019060010161559a565b5050509392505050565b5f602082840312156155d3575f5ffd5b815180151581146145a2575f5ffd5b6bffffffffffffffffffffffff198b60601b1681528960148201528860348201526001600160401b0360c01b8860c01b1660548201526001600160401b0360c01b8760c01b16605c8201526001600160401b0360c01b8660c01b16606482015284606c82015283608c8201528260ac82015261568560cc82018360c01b7fffffffffffffffff000000000000000000000000000000000000000000000000169052565b60d4019a9950505050505050505050565b6001600160a01b0383168152604060208201525f6143fd6040830184614ac456fe60a06040526040516108b03803806108b083398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206108905f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206108905f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b60805161046b6104255f395f6010015261046b5ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610081575f357fffffffff000000000000000000000000000000000000000000000000000000001663278f794360e11b1461007957610077610085565b565b610077610095565b6100775b6100776100906100c3565b6100fa565b5f806100a43660048184610313565b8101906100b1919061034e565b915091506100bf8282610118565b5050565b5f6100f57f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e808015610114573d5ff35b3d5ffd5b61012182610172565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561016a5761016582826101fa565b505050565b6100bf61026c565b806001600160a01b03163b5f036101ac57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051610216919061041f565b5f60405180830381855af49150503d805f811461024e576040519150601f19603f3d011682016040523d82523d5f602084013e610253565b606091505b509150915061026385838361028b565b95945050505050565b34156100775760405163b398979f60e01b815260040160405180910390fd5b6060826102a05761029b826102ea565b6102e3565b81511580156102b757506001600160a01b0384163b155b156102e057604051639996b31560e01b81526001600160a01b03851660048201526024016101a3565b50805b9392505050565b8051156102fa5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f85851115610321575f5ffd5b8386111561032d575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561035f575f5ffd5b82356001600160a01b0381168114610375575f5ffd5b9150602083013567ffffffffffffffff811115610390575f5ffd5b8301601f810185136103a0575f5ffd5b803567ffffffffffffffff8111156103ba576103ba61033a565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156103e9576103e961033a565b604052818152828201602001871015610400575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220f0c40b8d0bc55a45013a71ee3fd7264ce873fcc561d04a41ba1f24adbcef958c64736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103a2646970667358221220b7a51ba6eebd3f683f2639c5683af571d124e75eaed64fb37dc7a1aef053b38664736f6c634300081c0033",
}

// PolygonrollupmanagermockABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagermockMetaData.ABI instead.
var PolygonrollupmanagermockABI = PolygonrollupmanagermockMetaData.ABI

// PolygonrollupmanagermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagermockMetaData.Bin instead.
var PolygonrollupmanagermockBin = PolygonrollupmanagermockMetaData.Bin

// DeployPolygonrollupmanagermock deploys a new Ethereum contract, binding an instance of Polygonrollupmanagermock to it.
func DeployPolygonrollupmanagermock(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagermock, error) {
	parsed, err := PolygonrollupmanagermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagermockBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _aggLayerGateway)
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

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) ROLLUPMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "ROLLUP_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Polygonrollupmanagermock.Contract.ROLLUPMANAGERVERSION(&_Polygonrollupmanagermock.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Polygonrollupmanagermock.Contract.ROLLUPMANAGERVERSION(&_Polygonrollupmanagermock.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AggLayerGateway() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.AggLayerGateway(&_Polygonrollupmanagermock.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) AggLayerGateway() (common.Address, error) {
	return _Polygonrollupmanagermock.Contract.AggLayerGateway(&_Polygonrollupmanagermock.CallOpts)
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

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) ExposedCheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "exposed_checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) ExposedCheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonrollupmanagermock.Contract.ExposedCheckStateRootInsidePrime(&_Polygonrollupmanagermock.CallOpts, newStateRoot)
}

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) ExposedCheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Polygonrollupmanagermock.Contract.ExposedCheckStateRootInsidePrime(&_Polygonrollupmanagermock.CallOpts, newStateRoot)
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

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagermock.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Polygonrollupmanagermock.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagermock.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
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
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	if err != nil {
		return *new(PolygonRollupManagerRollupDataReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(PolygonRollupManagerRollupDataReturn)).(*PolygonRollupManagerRollupDataReturn)

	return out0, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupIDToRollupData(rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupData(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupIDToRollupData(rollupID uint32) (PolygonRollupManagerRollupDataReturn, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupData(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 _legacyLastPendingState, uint64 _legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupIDToRollupDataDeserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupIDToRollupDataDeserialized", rollupID)

	outstruct := new(struct {
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
	outstruct.LegacyLastPendingState = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.LegacyLastPendingStateConsolidated = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.LastVerifiedBatchBeforeUpgrade = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.RollupTypeID = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.RollupVerifierType = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 _legacyLastPendingState, uint64 _legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataDeserialized(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 _legacyLastPendingState, uint64 _legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataDeserialized(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupIDToRollupDataV2(opts *bind.CallOpts, rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupIDToRollupDataV2", rollupID)

	if err != nil {
		return *new(PolygonRollupManagerRollupDataReturnV2), err
	}

	out0 := *abi.ConvertType(out[0], new(PolygonRollupManagerRollupDataReturnV2)).(*PolygonRollupManagerRollupDataReturnV2)

	return out0, err

}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupIDToRollupDataV2(rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataV2(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupIDToRollupDataV2(rollupID uint32) (PolygonRollupManagerRollupDataReturnV2, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataV2(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupIDToRollupDataV2Deserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupIDToRollupDataV2Deserialized", rollupID)

	outstruct := new(struct {
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
	outstruct.LastVerifiedBatchBeforeUpgrade = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.RollupTypeID = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.RollupVerifierType = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.LastPessimisticRoot = *abi.ConvertType(out[10], new([32]byte)).(*[32]byte)
	outstruct.ProgramVKey = *abi.ConvertType(out[11], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataV2Deserialized(&_Polygonrollupmanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupIDToRollupDataV2Deserialized(&_Polygonrollupmanagermock.CallOpts, rollupID)
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
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	var out []interface{}
	err := _Polygonrollupmanagermock.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Polygonrollupmanagermock.Contract.RollupTypeMap(&_Polygonrollupmanagermock.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
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

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddExistingRollup(&_Polygonrollupmanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddExistingRollup(&_Polygonrollupmanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddNewRollupType(&_Polygonrollupmanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AddNewRollupType(&_Polygonrollupmanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AttachAggchainToAL(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.AttachAggchainToAL(&_Polygonrollupmanagermock.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) Initialize() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.Initialize(&_Polygonrollupmanagermock.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) Initialize() (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.Initialize(&_Polygonrollupmanagermock.TransactOpts)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) InitializeMock(opts *bind.TransactOpts, trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "initializeMock", trustedAggregator, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) InitializeMock(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.InitializeMock(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) InitializeMock(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.InitializeMock(&_Polygonrollupmanagermock.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
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

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RollbackBatches(&_Polygonrollupmanagermock.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.RollbackBatches(&_Polygonrollupmanagermock.TransactOpts, rollupContract, targetBatch)
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

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) SetRollupData(opts *bind.TransactOpts, rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "setRollupData", rollupID, lastLocalExitRoot, lastPessimisticRoot)
}

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) SetRollupData(rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetRollupData(&_Polygonrollupmanagermock.TransactOpts, rollupID, lastLocalExitRoot, lastPessimisticRoot)
}

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) SetRollupData(rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.SetRollupData(&_Polygonrollupmanagermock.TransactOpts, rollupID, lastLocalExitRoot, lastPessimisticRoot)
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

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagermock.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.UpdateRollupByRollupAdmin(&_Polygonrollupmanagermock.TransactOpts, rollupContract, newRollupTypeID)
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

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagermock.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagermock *PolygonrollupmanagermockTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagermock.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagermock.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
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
	RollupVerifierType             uint8
	LastVerifiedBatchBeforeUpgrade uint64
	ProgramVKey                    [32]byte
	InitPessimisticRoot            [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterAddExistingRollup is a free log retrieval operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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

// ParseAddExistingRollup is a log parse operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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
	RollupVerifierType      uint8
	Genesis                 [32]byte
	Description             string
	ProgramVKey             [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAddNewRollupType is a free log retrieval operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
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

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
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

// ParseAddNewRollupType is a log parse operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseAddNewRollupType(log types.Log) (*PolygonrollupmanagermockAddNewRollupType, error) {
	event := new(PolygonrollupmanagermockAddNewRollupType)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagermockCreateNewAggchainIterator is returned from FilterCreateNewAggchain and is used to iterate over the raw logs and unpacked data for CreateNewAggchain events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockCreateNewAggchainIterator struct {
	Event *PolygonrollupmanagermockCreateNewAggchain // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockCreateNewAggchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockCreateNewAggchain)
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
		it.Event = new(PolygonrollupmanagermockCreateNewAggchain)
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
func (it *PolygonrollupmanagermockCreateNewAggchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockCreateNewAggchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockCreateNewAggchain represents a CreateNewAggchain event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockCreateNewAggchain struct {
	RollupID                uint32
	RollupTypeID            uint32
	RollupAddress           common.Address
	ChainID                 uint64
	RollupVerifierType      uint8
	InitializeBytesAggchain []byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterCreateNewAggchain is a free log retrieval operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterCreateNewAggchain(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagermockCreateNewAggchainIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockCreateNewAggchainIterator{contract: _Polygonrollupmanagermock.contract, event: "CreateNewAggchain", logs: logs, sub: sub}, nil
}

// WatchCreateNewAggchain is a free log subscription operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchCreateNewAggchain(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockCreateNewAggchain, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockCreateNewAggchain)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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

// ParseCreateNewAggchain is a log parse operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseCreateNewAggchain(log types.Log) (*PolygonrollupmanagermockCreateNewAggchain, error) {
	event := new(PolygonrollupmanagermockCreateNewAggchain)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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

// PolygonrollupmanagermockRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRollbackBatchesIterator struct {
	Event *PolygonrollupmanagermockRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockRollbackBatches)
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
		it.Event = new(PolygonrollupmanagermockRollbackBatches)
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
func (it *PolygonrollupmanagermockRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockRollbackBatches represents a RollbackBatches event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*PolygonrollupmanagermockRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockRollbackBatchesIterator{contract: _Polygonrollupmanagermock.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockRollbackBatches)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseRollbackBatches(log types.Log) (*PolygonrollupmanagermockRollbackBatches, error) {
	event := new(PolygonrollupmanagermockRollbackBatches)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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

// PolygonrollupmanagermockUpdateRollupManagerVersionIterator is returned from FilterUpdateRollupManagerVersion and is used to iterate over the raw logs and unpacked data for UpdateRollupManagerVersion events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockUpdateRollupManagerVersionIterator struct {
	Event *PolygonrollupmanagermockUpdateRollupManagerVersion // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockUpdateRollupManagerVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockUpdateRollupManagerVersion)
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
		it.Event = new(PolygonrollupmanagermockUpdateRollupManagerVersion)
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
func (it *PolygonrollupmanagermockUpdateRollupManagerVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockUpdateRollupManagerVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockUpdateRollupManagerVersion represents a UpdateRollupManagerVersion event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockUpdateRollupManagerVersion struct {
	RollupManagerVersion string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupManagerVersion is a free log retrieval operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterUpdateRollupManagerVersion(opts *bind.FilterOpts) (*PolygonrollupmanagermockUpdateRollupManagerVersionIterator, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockUpdateRollupManagerVersionIterator{contract: _Polygonrollupmanagermock.contract, event: "UpdateRollupManagerVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupManagerVersion is a free log subscription operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchUpdateRollupManagerVersion(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockUpdateRollupManagerVersion) (event.Subscription, error) {

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockUpdateRollupManagerVersion)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseUpdateRollupManagerVersion(log types.Log) (*PolygonrollupmanagermockUpdateRollupManagerVersion, error) {
	event := new(PolygonrollupmanagermockUpdateRollupManagerVersion)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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

// PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator is returned from FilterVerifyPessimisticStateTransition and is used to iterate over the raw logs and unpacked data for VerifyPessimisticStateTransition events raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator struct {
	Event *PolygonrollupmanagermockVerifyPessimisticStateTransition // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagermockVerifyPessimisticStateTransition)
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
		it.Event = new(PolygonrollupmanagermockVerifyPessimisticStateTransition)
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
func (it *PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagermockVerifyPessimisticStateTransition represents a VerifyPessimisticStateTransition event raised by the Polygonrollupmanagermock contract.
type PolygonrollupmanagermockVerifyPessimisticStateTransition struct {
	RollupID            uint32
	PrevPessimisticRoot [32]byte
	NewPessimisticRoot  [32]byte
	PrevLocalExitRoot   [32]byte
	NewLocalExitRoot    [32]byte
	L1InfoRoot          [32]byte
	TrustedAggregator   common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVerifyPessimisticStateTransition is a free log retrieval operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) FilterVerifyPessimisticStateTransition(opts *bind.FilterOpts, rollupID []uint32, trustedAggregator []common.Address) (*PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.FilterLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagermockVerifyPessimisticStateTransitionIterator{contract: _Polygonrollupmanagermock.contract, event: "VerifyPessimisticStateTransition", logs: logs, sub: sub}, nil
}

// WatchVerifyPessimisticStateTransition is a free log subscription operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) WatchVerifyPessimisticStateTransition(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagermockVerifyPessimisticStateTransition, rollupID []uint32, trustedAggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagermock.contract.WatchLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagermockVerifyPessimisticStateTransition)
				if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
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

// ParseVerifyPessimisticStateTransition is a log parse operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Polygonrollupmanagermock *PolygonrollupmanagermockFilterer) ParseVerifyPessimisticStateTransition(log types.Log) (*PolygonrollupmanagermockVerifyPessimisticStateTransition, error) {
	event := new(PolygonrollupmanagermockVerifyPessimisticStateTransition)
	if err := _Polygonrollupmanagermock.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

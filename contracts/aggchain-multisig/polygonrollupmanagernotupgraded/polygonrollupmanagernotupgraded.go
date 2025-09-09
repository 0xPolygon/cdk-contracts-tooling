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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewRollupTypeMustBePessimisticOrALGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"CompletedMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"CreateNewAggchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"InitMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"}],\"name\":\"VerifyPessimisticStateTransition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"isRollupMigrating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataDeserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2Deserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIPolygonRollupManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516161ce3803806161ce83398101604081905261002f9161019a565b838383836001600160a01b038416158061005057506001600160a01b038316155b8061006257506001600160a01b038216155b8061007457506001600160a01b038116155b15610092576040516342eda64b60e11b815260040160405180910390fd5b6001600160a01b0380851660805283811660c05282811660a052811660e0526100b96100c6565b50505050505050506101f6565b5f54610100900460ff16156101315760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610181575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610197575f5ffd5b50565b5f5f5f5f608085870312156101ad575f5ffd5b84516101b881610183565b60208601519094506101c981610183565b60408601519093506101da81610183565b60608601519092506101eb81610183565b939692955090935050565b60805160a05160c05160e051615f6d6102615f395f81816106ea0152610ee501525f818161083001528181611d83015261376101525f81816106b6015281816127e9015261385e01525f818161077601528181610a7601528181610d4b0152610ff10152615f6d5ff3fe608060405234801561000f575f5ffd5b506004361061025a575f3560e01c8063066ec0121461025e57806311f6b2871461028e5780631489ed10146102a157806315064c96146102b65780631796a1ae146102d35780632072f6c5146102f0578063248a9ca3146102f857806325280169146103195780632f2ff15d146103c957806330c27dde146103dc57806336568abe146103ef5780633a7094bd14610402578063477fa2701461042457806354fd4d501461042c57806355a71ee014610457578063604691691461049757806365c0504d1461049f5780636c7668771461051a578063706039091461052d5780637222020f146105db57806374d9c244146105ee5780637975fcfe1461060e5780637fb6e76a146106215780638129fc1c146106465780638fd88cc21461064e57806391d148541461066157806397d289a31461067457806399f5634e146106875780639a908e731461068f578063a217fddf146106a2578063a2967d99146106a9578063a3c573eb146106b1578063ab0475cf146106e5578063abcb51981461070c578063c1acbc341461071f578063c4c928c214610739578063ceee281d1461074c578063d02103ca14610771578063d5073f6f14610798578063d547741f146107ab578063d8905812146107be578063dbc16976146107e3578063dd0464b9146107eb578063dde0ff77146107fe578063dfdb8c5e14610818578063e46761c41461082b578063e4f3d8f914610852578063e764a37314610902578063e80e503014610915578063f4e9267514610928578063f8c8765e14610938578063f9c4c2ae1461094b575b5f5ffd5b608454610271906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61027161029c36600461445e565b61096b565b6102b46102af3660046144a8565b61098e565b005b606f546102c39060ff1681565b6040519015158152602001610285565b607e546102e39063ffffffff1681565b6040516102859190614536565b6102b4610b3b565b61030b610306366004614547565b610bf3565b604051908152602001610285565b61039661032736600461455e565b60408051606080820183525f808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b03908116918301919091529282015190921690820152606001610285565b6102b46103d736600461458f565b610c07565b608754610271906001600160401b031681565b6102b46103fd36600461458f565b610c28565b6102c361041036600461445e565b60886020525f908152604090205460ff1681565b60865461030b565b604080518082019091526006815265076312e302e360d41b60208201525b60405161028591906145eb565b61030b61046536600461455e565b63ffffffff82165f9081526081602090815260408083206001600160401b038516845260020190915290205492915050565b61030b610c5f565b6105076104ad36600461445e565b607f6020525f908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b604051610285979695949392919061463e565b6102b46105283660046146d4565b610c74565b6105c361053b36600461445e565b63ffffffff165f9081526081602052604090208054600182015460058301546006840154600785015460088601546009909601546001600160a01b0380871698600160a01b978890046001600160401b03908116999288169890970487169680861695600160401b908190048216958281169591810490921693600160801b90920460ff1692565b6040516102859c9b9a99989796959493929190614774565b6102b46105e936600461445e565b6111bd565b6106016105fc36600461445e565b61129b565b6040516102859190614806565b61044a61061c3660046148dd565b6113ea565b6102e361062f366004614938565b60836020525f908152604090205463ffffffff1681565b6102b461141a565b6102b461065c366004614951565b6114f3565b6102c361066f36600461458f565b611860565b6102b4610682366004614a30565b61188a565b61030b611d7f565b61027161069d366004614a89565b611e5b565b61030b5f81565b61030b612047565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6040516102859190614ab1565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6102b461071a366004614ad3565b6122f5565b60845461027190600160801b90046001600160401b031681565b6102b4610747366004614b79565b6125f6565b6102e361075a366004614ba4565b60826020525f908152604090205463ffffffff1681565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6102b46107a6366004614547565b6126fe565b6102b46107b936600461458f565b612789565b61044a60405180604001604052806006815260200165076312e302e360d41b81525081565b6102b46127a5565b61044a6107f9366004614bbf565b612857565b60845461027190600160401b90046001600160401b031681565b6102b4610826366004614c2a565b61287e565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6108ea61086036600461445e565b63ffffffff165f90815260816020526040902080546001820154600583015460068401546007909401546001600160a01b0380851696600160a01b958690046001600160401b03908116979286169690950485169480831693600160401b808504831694600160801b808204851695600160c01b9092048516948085169493840416920460ff1690565b6040516102859c9b9a99989796959493929190614c54565b6102b4610910366004614cf7565b612c02565b6102b4610923366004614d12565b612d6f565b6080546102e39063ffffffff1681565b6102b4610946366004614d93565b61319b565b61095e61095936600461445e565b6133ab565b6040516102859190614dec565b63ffffffff81165f90815260816020526040812061098890613501565b92915050565b5f516020615f185f395f51905f526109a581613518565b6001600160401b038816156109cd5760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610a0257610a0261460a565b14610a20576040516390db0d0760e01b815260040160405180910390fd5b610a2f81898989898989613522565b600681018054600160401b600160801b031916600160401b6001600160401b038a16908102919091179091555f9081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610aab612047565b6040518263ffffffff1660e01b8152600401610ac991815260200190565b5f604051808303815f87803b158015610ae0575f5ffd5b505af1158015610af2573d5f5f3e3d5ffd5b50505050336001600160a01b03168a63ffffffff165f516020615d785f395f51905f5289888a604051610b2793929190614ed5565b60405180910390a350505050505050505050565b610b525f516020615ed85f395f51905f5233611860565b610be957608454600160801b90046001600160401b03161580610b9e57506084544290610b939062093a8090600160801b90046001600160401b0316614f0a565b6001600160401b0316115b80610bcb57506087544290610bc09062093a80906001600160401b0316614f0a565b6001600160401b0316115b15610be95760405163692baaad60e11b815260040160405180910390fd5b610bf161385c565b565b5f9081526034602052604090206001015490565b610c1082610bf3565b610c1981613518565b610c2383836138d2565b505050565b6001600160a01b0381163314610c5157604051630b4ad1cd60e31b815260040160405180910390fd5b610c5b828261393a565b5050565b5f6086546064610c6f9190614f29565b905090565b5f516020615f185f395f51905f52610c8b81613518565b610c936139a0565b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610cc857610cc861460a565b03610ce657604051635b6602b760e01b815260040160405180910390fd5b60016007820154600160801b900460ff166002811115610d0857610d0861460a565b148015610d1457508215155b15610d3257604051630ded782d60e01b815260040160405180910390fd5b60405163ef4eeb3560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ef4eeb3590610d80908d90600401614536565b602060405180830381865afa158015610d9b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dbf9190614f40565b905080610ddf5760405163a60721e160e01b815260040160405180910390fd5b63ffffffff8b165f9081526088602052604090205460ff1615610e9557600582015480610e2957507f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d7575b898114610e49576040516306c6486360e11b815260040160405180910390fd5b5f6005840181905563ffffffff8d1680825260886020526040808320805460ff191690555190917f6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e91a2505b5f610ea58c84848d8d8b8b6139e7565b905060026007840154600160801b900460ff166002811115610ec957610ec961460a565b03610f4f5760405163a48fd37760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a48fd37790610f1e9084908c908c90600401614f7f565b5f6040518083038186803b158015610f34575f5ffd5b505afa158015610f46573d5f5f3e3d5ffd5b50505050610fb7565b6001830154600984015460405163020a49e360e51b81526001600160a01b03909216916341493c6091610f8a9185908d908d90600401614fae565b5f6040518083038186803b158015610fa0575f5ffd5b505afa158015610fb2573d5f5f3e3d5ffd5b505050505b60848054600160801b600160c01b031916600160801b426001600160401b031602179055600583018054908b9055600884018054908b90557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d611026612047565b6040518263ffffffff1660e01b815260040161104491815260200190565b5f604051808303815f87803b15801561105b575f5ffd5b505af115801561106d573d5f5f3e3d5ffd5b505050505f8c9050336001600160a01b03168f63ffffffff165f516020615d785f395f51905f525f5f5f1b856040516110a893929190614ed5565b60405180910390a3336001600160a01b03168f63ffffffff167fdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711848f87868b604051611116959493929190948552602085019390935260408401919091526060830152608082015260a00190565b60405180910390a360026007870154600160801b900460ff1660028111156111405761114061460a565b036111a4578554604051639ee4afa360e01b81526001600160a01b0390911690639ee4afa390611176908c908c90600401614fd9565b5f604051808303815f87803b15801561118d575f5ffd5b505af115801561119f573d5f5f3e3d5ffd5b505050505b5050505050506111b2613b6a565b505050505050505050565b5f516020615dd85f395f51905f526111d481613518565b63ffffffff821615806111f25750607e5463ffffffff908116908316115b1561121057604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001810154600160e81b900460ff161561125157604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b6112fe60408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082019081526020015f81526020015f81525090565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b0380821686526001600160401b03600160a01b928390048116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b90819004821660c0850152600783015480831660e086015290810490911661010084015261012083019060ff600160801b9091041660028111156113ba576113ba61460a565b908160028111156113cd576113cd61460a565b905250600881015461014083015260090154610160820152919050565b63ffffffff86165f90815260816020526040902060609061140f908787878787613b81565b979650505050505050565b5f54600590610100900460ff1615801561143a57505f5460ff8083169116105b61145f5760405162461bcd60e51b815260040161145690614fec565b60405180910390fd5b5f805461ffff191660ff8316176101001790556040805180820182526006815265076312e302e360d41b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2916114bb916145eb565b60405180910390a15f805461ff001916905560405160ff821681525f516020615e585f395f51905f529060200160405180910390a150565b6114fb6139a0565b6115125f516020615d985f395f51905f5233611860565b1580156115905750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015611560573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611584919061503a565b6001600160a01b031614155b156115ae57604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b0382165f9081526082602052604081205463ffffffff16908190036115ed576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff1660028111156116225761162261460a565b14611640576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b0390811690841681111580611678575060068201546001600160401b03600160401b9091048116908516105b156116965760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b03161461172e576001600160401b038082165f908152600385016020526040902060010154600160401b900481169086168110156116fb57604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091165f908152600384016020526040812090815560010180546001600160801b0319169055611698565b6006830180546001600160401b0319166001600160401b0387161790556117558583615055565b608480545f9061176f9084906001600160401b0316615055565b82546101009290920a6001600160401b0381810219909316918316021790915586165f8181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece906044015f604051808303815f87803b1580156117e8575f5ffd5b505af11580156117fa573d5f5f3e3d5ffd5b505050506001600160401b0385165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a350505050610c5b613b6a565b5f9182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f516020615ef85f395f51905f526118a181613518565b63ffffffff841615806118bf5750607e5463ffffffff908116908516115b156118dd57604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff161561191e57604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b038516111561194c57604051634c753f5760e01b815260040160405180910390fd5b6001600160401b0384165f9081526083602052604090205463ffffffff1615611988576040516337c8fe0960e11b815260040160405180910390fd5b608080545f919082906119a09063ffffffff16615074565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f80825260208201928390529394506001600160a01b039092169130916119eb90614431565b6119f793929190615098565b604051809103905ff080158015611a10573d5f5f3e3d5ffd5b506001600160401b0387165f818152608360209081526040808320805463ffffffff1990811663ffffffff8a81169182179093556001600160a01b038816808752608286528487208054909316821790925585526081909352922080546001600160e01b031916909117600160a01b909302929092178255600782018054600160401b600160801b03198116928c16600160401b02928317825560018801549495509293600160e01b900460ff16929091600160401b600160881b031990911660ff60801b1990911617600160801b836002811115611af157611af161460a565b021790555060026001850154600160e01b900460ff166002811115611b1857611b1861460a565b03611c34578263ffffffff167f144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b89848a88600101601c9054906101000a900460ff166002811115611b6b57611b6b61460a565b8b604051611b7d9594939291906150cc565b60405180910390a28263ffffffff165f516020615df85f395f51905f5289848a6001600160a01b03604051611bb59493929190615113565b60405180910390a25f86806020019051810190611bd2919061503a565b60405163b3a326f760e01b81529091506001600160a01b0384169063b3a326f790611c01908490600401614ab1565b5f604051808303815f87803b158015611c18575f5ffd5b505af1158015611c2a573d5f5f3e3d5ffd5b5050505050611d75565b600184810180549183018054600160a01b600160e01b03198116600160a01b948590046001600160401b0316909402938417825591546001600160a01b03166001600160e01b03199092166001600160a01b0319909316929092171790556002808501545f808052918301602090815260408320919091556003860154600984015587518291829182918291611cd191908d018101908d01615198565b945094509450945094508763ffffffff165f516020615df85f395f51905f528e898f87604051611d049493929190615113565b60405180910390a2604051633892b81160e11b81526001600160a01b03881690637125702290611d4290889088908d90899089908990600401615236565b5f604051808303815f87803b158015611d59575f5ffd5b505af1158015611d6b573d5f5f3e3d5ffd5b5050505050505050505b5050505050505050565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611dcd9190614ab1565b602060405180830381865afa158015611de8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e0c9190614f40565b6084549091505f90611e30906001600160401b03600160401b820481169116615055565b6001600160401b03169050805f03611e4a575f9250505090565b611e5481836152a8565b9250505090565b606f545f9060ff1615611e8157604051630bc011ff60e21b815260040160405180910390fd5b335f9081526082602052604081205463ffffffff1690819003611eb7576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03165f03611ee057604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff166002811115611f1557611f1561460a565b14611f33576040516390db0d0760e01b815260040160405180910390fd5b608480548691905f90611f509084906001600160401b0316614f0a565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690505f611f838783614f0a565b6006840180546001600160401b038381166001600160401b03199092168217909255604080516060810182528a815242841660208083019182528886168385019081525f86815260038c018352859020935184559151600193909301805492518716600160401b026001600160801b03199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b6080545f9063ffffffff1680820361206057505f919050565b5f816001600160401b038111156120795761207961496d565b6040519080825280602002602001820160405280156120a2578160200160208202803683370190505b5090505f5b828110156120ff5760815f6120bd8360016152bb565b63ffffffff1663ffffffff1681526020019081526020015f20600501548282815181106120ec576120ec6152ce565b60209081029190910101526001016120a7565b505f60205b836001146122a1575f6121186002866152e2565b6121236002876152a8565b61212d91906152bb565b90505f816001600160401b038111156121485761214861496d565b604051908082528060200260200182016040528015612171578160200160208202803683370190505b5090505f5b8281101561227a576121896001846152f5565b811480156121a1575061219d6002886152e2565b6001145b156121f6576121d3866121b5836002614f29565b815181106121c5576121c56152ce565b602002602001015186613cb5565b8282815181106121e5576121e56152ce565b602002602001018181525050612272565b61225386612205836002614f29565b81518110612215576122156152ce565b60200260200101518783600261222b9190614f29565b6122369060016152bb565b81518110612246576122466152ce565b6020026020010151613cb5565b828281518110612265576122656152ce565b6020026020010181815250505b600101612176565b5080945081955061228b8485613cb5565b93508261229781615308565b9350505050612104565b5f835f815181106122b4576122b46152ce565b602002602001015190505f5f90505b828110156122eb576122d58285613cb5565b91506122e18485613cb5565b93506001016122c3565b5095945050505050565b5f516020615e185f395f51905f5261230c81613518565b306001600160a01b038916036123355760405163325c055b60e21b815260040160405180910390fd5b607e80545f9190829061234d9063ffffffff16615074565b91906101000a81548163ffffffff021916908363ffffffff160217905590506001600281111561237f5761237f61460a565b8660028111156123915761239161460a565b036123ba5784156123b5576040516363d722e760e01b815260040160405180910390fd5b612474565b60028660028111156123ce576123ce61460a565b03612424576001600160a01b0388161515806123f257506001600160401b03871615155b806123fc57508415155b8061240657508215155b156123b5576040516363d722e760e01b815260040160405180910390fd5b5f8660028111156124375761243761460a565b0361245b5782156123b5576040516363d722e760e01b815260040160405180910390fd5b6040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b031681526020018760028111156124be576124be61460a565b81525f602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b031990911617825591850151600182018054948701516001600160401b0316600160a01b026001600160e01b031990951691909316179290921780825592840151919260ff60e01b1916600160e01b8360028111156125615761256161460a565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a906125e3908c908c908c908c908c908c908c9061531d565b60405180910390a2505050505050505050565b5f516020615d985f395f51905f5261260d81613518565b6001600160a01b0384165f9081526082602090815260408083205463ffffffff168084526081909252909120600263ffffffff86165f908152607f6020526040902060010154600160e01b900460ff16600281111561266e5761266e61460a565b141580156126cd575063ffffffff85165f908152607f6020526040902060010154600160e01b900460ff1660028111156126aa576126aa61460a565b6007820154600160801b900460ff1660028111156126ca576126ca61460a565b14155b156126eb57604051635aa0d5f160e11b815260040160405180910390fd5b6126f6868686613cc3565b505050505050565b5f516020615db85f395f51905f5261271581613518565b683635c9adc5dea0000082118061272f5750633b9aca0082105b1561274d57604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b61279282610bf3565b61279b81613518565b610c23838361393a565b5f516020615e385f395f51905f526127bc81613518565b608780546001600160401b031916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc16976916004808301925f92919082900301818387803b158015612836575f5ffd5b505af1158015612848573d5f5f3e3d5ffd5b50505050612854613f6a565b50565b63ffffffff86165f90815260816020526040902060609061140f90889088888888886139e7565b6001600160a01b0382165f9081526082602090815260408083205463ffffffff1683526081909152902060026007820154600160801b900460ff1660028111156128ca576128ca61460a565b03612a7357336001600160a01b0316836001600160a01b0316637388c4366040518163ffffffff1660e01b8152600401602060405180830381865afa158015612915573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612939919061503a565b6001600160a01b0316146129605760405163660a7ce560e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020908152604091829020548251636e7fbce960e01b815292516001600160a01b0390911692636e7fbce99260048083019391928290030181865afa1580156129b8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129dc919061537c565b6001600160f01b031916836001600160a01b0316636e7fbce96040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a22573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a46919061537c565b6001600160f01b03191614612a6e57604051635aa0d5f160e11b815260040160405180910390fd5b612b3b565b336001600160a01b0316836001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ab9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612add919061503a565b6001600160a01b031614612b045760405163696072e960e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612b3b5760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b031610612b7657604051633e37e23360e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020526040902060010154600160e01b900460ff166002811115612ba957612ba961460a565b6007820154600160801b900460ff166002811115612bc957612bc961460a565b14612be757604051635aa0d5f160e11b815260040160405180910390fd5b604080515f815260208101909152610c239084908490613cc3565b5f516020615d985f395f51905f52612c1981613518565b63ffffffff84165f908152608160205260408120906007820154600160801b900460ff166002811115612c4e57612c4e61460a565b14612c6c576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612ca35760405163664316a560e11b815260040160405180910390fd5b5f63ffffffff85165f908152607f6020526040902060010154600160e01b900460ff166002811115612cd757612cd761460a565b03612cf5576040516301ea4e5b60e01b815260040160405180910390fd5b63ffffffff85165f908152608860205260409020805460ff191660011790558054612d2a906001600160a01b03168585613cc3565b8463ffffffff167f3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de985604051612d609190614536565b60405180910390a25050505050565b5f516020615e785f395f51905f52612d8681613518565b6001600160401b0386165f9081526083602052604090205463ffffffff1615612dc2576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b0387161115612df057604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b0389165f9081526082602052604090205463ffffffff1615612e2c57604051630d409b9360e41b815260040160405180910390fd5b6001600160a01b038916301480612e4b57506001600160a01b0389163b155b15612e695760405163325c055b60e21b815260040160405180910390fd5b608080545f91908290612e819063ffffffff16615074565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f896001600160401b03166001600160401b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8363ffffffff1663ffffffff1681526020019081526020015f2090508a815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550888160010160146101000a8154816001600160401b0302191690836001600160401b0316021790555089816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555087815f0160146101000a8154816001600160401b0302191690836001600160401b03160217905550858160070160106101000a81548160ff021916908360028111156130085761300861460a565b021790555060018660028111156130215761302161460a565b0361306a576009810185905560088101849055600581018790556001600160a01b038a163b5f036130655760405163043103a360e21b815260040160405180910390fd5b613144565b600286600281111561307e5761307e61460a565b036130dd576001600160a01b038a161515806130a257506001600160401b03891615155b806130ac57508415155b156130ca57604051636fc5557f60e01b815260040160405180910390fd5b6008810184905560058101879055613144565b841515806130ea57508315155b1561310857604051636fc5557f60e01b815260040160405180910390fd5b5f80805260028201602052604081208890556001600160a01b038b163b90036131445760405163043103a360e21b815260040160405180910390fd5b8163ffffffff167f4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e8a8d8b8a5f8b8b60405161318697969594939291906153a3565b60405180910390a25050505050505050505050565b5f54600590610100900460ff161580156131bb57505f5460ff8083169116105b6131d75760405162461bcd60e51b815260040161145690614fec565b5f805461ffff191660ff83161761010017905567016345785d8a00006086556131fe613fc1565b6132155f516020615f185f395f51905f528661402b565b61321f5f8461402b565b6132365f516020615e185f395f51905f528461402b565b61324d5f516020615e785f395f51905f528461402b565b6132645f516020615d985f395f51905f528461402b565b61327b5f516020615dd85f395f51905f528561402b565b6132925f516020615ef85f395f51905f528561402b565b6132a95f516020615e385f395f51905f528561402b565b6132d37fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db18561402b565b6132f75f516020615f185f395f51905f525f516020615d585f395f51905f52614035565b61330e5f516020615d585f395f51905f528561402b565b6133255f516020615db85f395f51905f528561402b565b6133495f516020615ed85f395f51905f525f516020615e985f395f51905f52614035565b6133605f516020615ed85f395f51905f528361402b565b6133775f516020615e985f395f51905f528361402b565b5f805461ff001916905560405160ff821681525f516020615e585f395f51905f529060200160405180910390a15050505050565b61340f60408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b03908116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b808204831660c0860152600160801b808304841660e0870152600160c01b909204831661010086015260078401548084166101208701529081049092166101408501526101608401910460ff1660028111156134e4576134e461460a565b908160028111156134f7576134f761460a565b8152505050919050565b60060154600160401b90046001600160401b031690565b6128548133614086565b5f5f61352d89613501565b60078a01549091506001600160401b0390811690891610156135625760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b0388165f90815260028a01602052604090205491508161359c576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b031611156135cf57604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b0316116136015760405163b9b18f5760e01b815260040160405180910390fd5b5f6136108a8a8a8a878b613b81565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161364491906153fe565b602060405180830381855afa15801561365f573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906136829190614f40565b61368c91906152e2565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a916136ce91899190600401615414565b602060405180830381865afa1580156136e9573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061370d9190615450565b61372a576040516309bde33960e01b815260040160405180910390fd5b5f613735848b615055565b905061378887826001600160401b031661374d611d7f565b6137579190614f29565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691906140ad565b80608460088282829054906101000a90046001600160401b03166137ac9190614f0a565b82546101009290920a6001600160401b0381810219909316918316021790915560848054600160801b600160c01b031916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064015f604051808303815f87803b158015613838575f5ffd5b505af115801561384a573d5f5f3e3d5ffd5b50505050505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156138b4575f5ffd5b505af11580156138c6573d5f5f3e3d5ffd5b50505050610bf16140ff565b6138dc8282611860565b610c5b575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6139448282611860565b15610c5b575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f516020615eb85f395f51905f525c156139cd57604051633ee5aeb560e01b815260040160405180910390fd5b610bf160015f516020615eb85f395f51905f525b9061415a565b606060026007880154600160801b900460ff166002811115613a0b57613a0b61460a565b03613abd578654604051631a957d9b60e21b81525f916001600160a01b031690636a55f66c90613a419087908790600401614fd9565b602060405180830381865afa158015613a5c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a809190614f40565b600589015460088a0154604051929350613aa6928a908d9086908c908c9060200161546f565b60405160208183030381529060405291505061140f565b865460408051632b47b7cd60e21b815290515f926001600160a01b03169163ad1edf349160048083019260209291908290030181865afa158015613b03573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b279190614f40565b600589015460088a0154604051929350613b4d928a908d9086908c908c9060200161546f565b604051602081830303815290604052915050979650505050505050565b610bf15f5f516020615eb85f395f51905f526139e1565b6001600160401b038086165f8181526003890160205260408082205493881682529020546060929115801590613bb5575081155b15613bd35760405163340c614f60e11b815260040160405180910390fd5b80613bf1576040516366385b5160e01b815260040160405180910390fd5b613bfa84614161565b613c17576040516305dae44f60e21b815260040160405180910390fd5b885460018a01546040516001600160601b03193360601b16602082015260348101889052605481018590526001600160c01b031960c08c811b82166074840152600160a01b94859004811b8216607c84015293909204831b82166084820152608c810187905260ac810184905260cc81018990529189901b1660ec82015260f401604051602081830303815290604052925050509695505050505050565b5f9182526020526040902090565b63ffffffff82161580613ce15750607e5463ffffffff908116908316115b15613cff57604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526082602052604081205463ffffffff1690819003613d3e576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181165f908152608160205260409020600781015490918516600160401b9091046001600160401b031603613d8b57604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615613dcc57604051633b8d3d9960e01b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825582546001600160401b03600160a01b9182900416026001600160e01b03199091169093179290921790915560038201546009840155600783018054600160401b63ffffffff891602600160401b600160801b0319821681178355925460ff600160e01b909104169260ff60801b1916600160401b600160881b031990911617600160801b836002811115613e8957613e8961460a565b02179055505f613e988461096b565b6007840180546001600160401b0319166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef28692613ee992169089906004016154ac565b5f604051808303815f87803b158015613f00575f5ffd5b505af1158015613f12573d5f5f3e3d5ffd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff16613f8d57604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f54610100900460ff16610bf15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401611456565b610c5b82826138d2565b5f61403f83610bf3565b5f84815260346020526040808220600101859055519192508391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6140908282611860565b610c5b57604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610c239084906141e6565b606f5460ff161561412357604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b80825d5050565b5f67ffffffff000000016001600160401b038316108015614196575067ffffffff00000001604083901c6001600160401b0316105b80156141b6575067ffffffff00000001608083901c6001600160401b0316105b80156141cd575067ffffffff0000000160c083901c105b156141da57506001919050565b505f919050565b919050565b5f61423a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166142b79092919063ffffffff16565b805190915015610c2357808060200190518101906142589190615450565b610c235760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611456565b60606142c584845f856142cd565b949350505050565b60608247101561432e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611456565b5f5f866001600160a01b0316858760405161434991906153fe565b5f6040518083038185875af1925050503d805f8114614383576040519150601f19603f3d011682016040523d82523d5f602084013e614388565b606091505b509150915061140f87838387606083156144025782515f036143fb576001600160a01b0385163b6143fb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611456565b50816142c5565b6142c583838151156144175781518083602001fd5b8060405162461bcd60e51b815260040161145691906145eb565b610888806154d083390190565b6001600160401b03169052565b803563ffffffff811681146141e1575f5ffd5b5f6020828403121561446e575f5ffd5b6144778261444b565b9392505050565b80356001600160401b03811681146141e1575f5ffd5b6001600160a01b0381168114612854575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b0312156144c0575f5ffd5b6144c98961444b565b97506144d760208a0161447e565b96506144e560408a0161447e565b95506144f360608a0161447e565b94506080890135935060a0890135925060c089013561451181614494565b91506103e089018a1015614523575f5ffd5b60e0890190509295985092959890939650565b63ffffffff91909116815260200190565b5f60208284031215614557575f5ffd5b5035919050565b5f5f6040838503121561456f575f5ffd5b6145788361444b565b91506145866020840161447e565b90509250929050565b5f5f604083850312156145a0575f5ffd5b8235915060208301356145b281614494565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61447760208301846145bd565b6001600160a01b03169052565b634e487b7160e01b5f52602160045260245ffd5b6003811061463a57634e487b7160e01b5f52602160045260245ffd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e08101614672606083018761461e565b931515608082015260a081019290925260c090910152949350505050565b5f5f83601f8401126146a0575f5ffd5b5081356001600160401b038111156146b6575f5ffd5b6020830191508360208285010111156146cd575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b0312156146eb575f5ffd5b6146f48961444b565b975061470260208a0161444b565b9650604089013595506060890135945060808901356001600160401b0381111561472a575f5ffd5b6147368b828c01614690565b90955093505060a08901356001600160401b03811115614754575f5ffd5b6147608b828c01614690565b999c989b5096995094979396929594505050565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a9052881660a082015261018081016147c160c083018961443e565b6147ce60e083018861443e565b6147dc61010083018761443e565b6147ea61012083018661461e565b61014082019390935261016001529a9950505050505050505050565b5f610180820190506148198284516145fd565b602083015161482b602084018261443e565b50604083015161483e60408401826145fd565b506060830151614851606084018261443e565b506080830151608083015260a083015161486e60a084018261443e565b5060c083015161488160c084018261443e565b5060e083015161489460e084018261443e565b506101008301516148a961010084018261443e565b506101208301516148be61012084018261461e565b5061014083015161014083015261016083015161016083015292915050565b5f5f5f5f5f5f60c087890312156148f2575f5ffd5b6148fb8761444b565b95506149096020880161447e565b94506149176040880161447e565b959894975094956060810135955060808101359460a0909101359350915050565b5f60208284031215614948575f5ffd5b6144778261447e565b5f5f60408385031215614962575f5ffd5b823561457881614494565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156149a9576149a961496d565b604052919050565b5f6001600160401b038211156149c9576149c961496d565b50601f01601f191660200190565b5f6149e96149e4846149b1565b614981565b90508281528383830111156149fc575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112614a21575f5ffd5b614477838335602085016149d7565b5f5f5f60608486031215614a42575f5ffd5b614a4b8461444b565b9250614a596020850161447e565b915060408401356001600160401b03811115614a73575f5ffd5b614a7f86828701614a12565b9150509250925092565b5f5f60408385031215614a9a575f5ffd5b614aa38361447e565b946020939093013593505050565b6001600160a01b0391909116815260200190565b8035600381106141e1575f5ffd5b5f5f5f5f5f5f5f60e0888a031215614ae9575f5ffd5b8735614af481614494565b96506020880135614b0481614494565b9550614b126040890161447e565b9450614b2060608901614ac5565b93506080880135925060a08801356001600160401b03811115614b41575f5ffd5b8801601f81018a13614b51575f5ffd5b614b608a8235602084016149d7565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614b8b575f5ffd5b8335614b9681614494565b9250614a596020850161444b565b5f60208284031215614bb4575f5ffd5b813561447781614494565b5f5f5f5f5f5f60a08789031215614bd4575f5ffd5b614bdd8761444b565b955060208701359450604087013593506060870135925060808701356001600160401b03811115614c0c575f5ffd5b614c1889828a01614690565b979a9699509497509295939492505050565b5f5f60408385031215614c3b575f5ffd5b8235614c4681614494565b91506145866020840161444b565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a9052881660a08201526101808101614ca160c083018961443e565b614cae60e083018861443e565b614cbc61010083018761443e565b614cca61012083018661443e565b614cd861014083018561443e565b614ce661016083018461461e565b9d9c50505050505050505050505050565b5f5f5f60608486031215614d09575f5ffd5b614b968461444b565b5f5f5f5f5f5f5f5f610100898b031215614d2a575f5ffd5b8835614d3581614494565b97506020890135614d4581614494565b9650614d5360408a0161447e565b9550614d6160608a0161447e565b945060808901359350614d7660a08a01614ac5565b979a969950949793969295929450505060c08201359160e0013590565b5f5f5f5f60808587031215614da6575f5ffd5b8435614db181614494565b93506020850135614dc181614494565b92506040850135614dd181614494565b91506060850135614de181614494565b939692955090935050565b5f61018082019050614dff8284516145fd565b6020830151614e11602084018261443e565b506040830151614e2460408401826145fd565b506060830151614e37606084018261443e565b506080830151608083015260a0830151614e5460a084018261443e565b5060c0830151614e6760c084018261443e565b5060e0830151614e7a60e084018261443e565b50610100830151614e8f61010084018261443e565b50610120830151614ea461012084018261443e565b50610140830151614eb961014084018261443e565b50610160830151614ece61016084018261461e565b5092915050565b6001600160401b039390931683526020830191909152604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03818116838216019081111561098857610988614ef6565b808202811582820484141761098857610988614ef6565b5f60208284031215614f50575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f614f9160408301866145bd565b8281036020840152614fa4818587614f57565b9695505050505050565b848152606060208201525f614fc660608301866145bd565b828103604084015261140f818587614f57565b602081525f6142c5602083018486614f57565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f6020828403121561504a575f5ffd5b815161447781614494565b6001600160401b03828116828216039081111561098857610988614ef6565b5f63ffffffff821663ffffffff810361508f5761508f614ef6565b60010192915050565b6001600160a01b038481168252831660208201526060604082018190525f906150c3908301846145bd565b95945050505050565b63ffffffff861681526001600160a01b03851660208201526001600160401b038416604082015260ff8316606082015260a0608082018190525f9061140f908301846145bd565b63ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b5f82601f83011261515a575f5ffd5b81516151686149e4826149b1565b81815284602083860101111561517c575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a086880312156151ac575f5ffd5b85516151b781614494565b60208701519095506151c881614494565b60408701519094506151d981614494565b60608701519093506001600160401b038111156151f4575f5ffd5b6152008882890161514b565b608088015190935090506001600160401b0381111561521d575f5ffd5b6152298882890161514b565b9150509295509295909350565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c0608082018190525f90615275908301856145bd565b82810360a084015261528781856145bd565b9998505050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f826152b6576152b6615294565b500490565b8082018082111561098857610988614ef6565b634e487b7160e01b5f52603260045260245ffd5b5f826152f0576152f0615294565b500690565b8181038181111561098857610988614ef6565b5f8161531657615316614ef6565b505f190190565b6001600160a01b038881168252871660208201526001600160401b038616604082015261534d606082018661461e565b83608082015260e060a08201525f61536860e08301856145bd565b90508260c083015298975050505050505050565b5f6020828403121561538c575f5ffd5b81516001600160f01b031981168114614477575f5ffd5b6001600160401b0388811682526001600160a01b03881660208301528616604082015260e081016153d7606083018761461e565b6001600160401b0394909416608082015260a081019290925260c090910152949350505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b6001811015615446578151835260209283019290910190600101615427565b5050509392505050565b5f60208284031215615460575f5ffd5b81518015158114614477575f5ffd5b9687526020870195909552604086019390935260e09190911b6001600160e01b03191660608501526064840152608483015260a482015260c40190565b6001600160a01b03831681526040602082018190525f906142c5908301846145bd56fe60a060405260405161088838038061088883398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206108685f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206108685f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516104436104255f395f601001526104435ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610069575f356001600160e01b03191663278f794360e11b146100615761005f61006d565b565b61005f61007d565b61005f5b61005f6100786100ab565b6100cf565b5f8061008c36600481846102ba565b81019061009991906102f5565b915091506100a782826100ed565b5050565b5f6100ca5f5160206103ee5f395f51905f52546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156100e9573d5ff35b3d5ffd5b6100f682610147565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561013f5761013a82826101aa565b505050565b6100a761021c565b806001600160a01b03163b5f0361017c5780604051634c9c8ce360e01b815260040161017391906103c3565b60405180910390fd5b5f5160206103ee5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516101c691906103d7565b5f60405180830381855af49150503d805f81146101fe576040519150601f19603f3d011682016040523d82523d5f602084013e610203565b606091505b509150915061021385838361023b565b95945050505050565b341561005f5760405163b398979f60e01b815260040160405180910390fd5b6060826102505761024b82610291565b61028a565b815115801561026757506001600160a01b0384163b155b156102875783604051639996b31560e01b815260040161017391906103c3565b50805b9392505050565b8051156102a15780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f858511156102c8575f5ffd5b838611156102d4575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610306575f5ffd5b82356001600160a01b038116811461031c575f5ffd5b915060208301356001600160401b03811115610336575f5ffd5b8301601f81018513610346575f5ffd5b80356001600160401b0381111561035f5761035f6102e1565b604051601f8201601f19908116603f011681016001600160401b038111828210171561038d5761038d6102e1565b6040528181528282016020018710156103a4575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b6001600160a01b0391909116815260200190565b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212209a5a4dfebf1519ea255e9040aa4f94d958c02e32bb27ddaba04dc327d126376264736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610373cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0d1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d366156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fbab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641ac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59062ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f47f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024983dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff2859519b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545ea0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a2646970667358221220fc5a050820926cdfe347a93e866cbbde27bcbc1d4bf7952b34d9dcad76955d0264736f6c634300081c0033",
}

// PolygonrollupmanagernotupgradedABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonrollupmanagernotupgradedMetaData.ABI instead.
var PolygonrollupmanagernotupgradedABI = PolygonrollupmanagernotupgradedMetaData.ABI

// PolygonrollupmanagernotupgradedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonrollupmanagernotupgradedMetaData.Bin instead.
var PolygonrollupmanagernotupgradedBin = PolygonrollupmanagernotupgradedMetaData.Bin

// DeployPolygonrollupmanagernotupgraded deploys a new Ethereum contract, binding an instance of Polygonrollupmanagernotupgraded to it.
func DeployPolygonrollupmanagernotupgraded(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Polygonrollupmanagernotupgraded, error) {
	parsed, err := PolygonrollupmanagernotupgradedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonrollupmanagernotupgradedBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _aggLayerGateway)
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

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) AggLayerGateway() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AggLayerGateway(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) AggLayerGateway() (common.Address, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AggLayerGateway(&_Polygonrollupmanagernotupgraded.CallOpts)
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

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Polygonrollupmanagernotupgraded.Contract.GetInputPessimisticBytes(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
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

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) IsRollupMigrating(opts *bind.CallOpts, rollupID uint32) (bool, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "isRollupMigrating", rollupID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.IsRollupMigrating(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Polygonrollupmanagernotupgraded.Contract.IsRollupMigrating(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
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

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupIDToRollupDataDeserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataDeserialized", rollupID)

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
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataDeserialized(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataDeserialized(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
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

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) RollupIDToRollupDataV2Deserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataV2Deserialized", rollupID)

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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataV2Deserialized(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Polygonrollupmanagernotupgraded.Contract.RollupIDToRollupDataV2Deserialized(&_Polygonrollupmanagernotupgraded.CallOpts, rollupID)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonrollupmanagernotupgraded.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Version() (string, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Version(&_Polygonrollupmanagernotupgraded.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedCallerSession) Version() (string, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Version(&_Polygonrollupmanagernotupgraded.CallOpts)
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

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddExistingRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AddExistingRollup(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
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

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AttachAggchainToAL(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.AttachAggchainToAL(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
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

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) InitMigration(opts *bind.TransactOpts, rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "initMigration", rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.InitMigration(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.InitMigration(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Initialize() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) Initialize() (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize(&_Polygonrollupmanagernotupgraded.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) Initialize0(opts *bind.TransactOpts, trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "initialize0", trustedAggregator, admin, timelock, emergencyCouncil)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) Initialize0(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize0(&_Polygonrollupmanagernotupgraded.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) Initialize0(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.Initialize0(&_Polygonrollupmanagernotupgraded.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
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

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Polygonrollupmanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Polygonrollupmanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
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
	InitPessimisticRoot            [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterAddExistingRollup is a free log retrieval operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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

// ParseAddExistingRollup is a log parse operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
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

// PolygonrollupmanagernotupgradedCompletedMigrationIterator is returned from FilterCompletedMigration and is used to iterate over the raw logs and unpacked data for CompletedMigration events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCompletedMigrationIterator struct {
	Event *PolygonrollupmanagernotupgradedCompletedMigration // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedCompletedMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedCompletedMigration)
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
		it.Event = new(PolygonrollupmanagernotupgradedCompletedMigration)
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
func (it *PolygonrollupmanagernotupgradedCompletedMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedCompletedMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedCompletedMigration represents a CompletedMigration event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCompletedMigration struct {
	RollupID uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompletedMigration is a free log retrieval operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterCompletedMigration(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedCompletedMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedCompletedMigrationIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "CompletedMigration", logs: logs, sub: sub}, nil
}

// WatchCompletedMigration is a free log subscription operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchCompletedMigration(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedCompletedMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedCompletedMigration)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
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

// ParseCompletedMigration is a log parse operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseCompletedMigration(log types.Log) (*PolygonrollupmanagernotupgradedCompletedMigration, error) {
	event := new(PolygonrollupmanagernotupgradedCompletedMigration)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonrollupmanagernotupgradedCreateNewAggchainIterator is returned from FilterCreateNewAggchain and is used to iterate over the raw logs and unpacked data for CreateNewAggchain events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCreateNewAggchainIterator struct {
	Event *PolygonrollupmanagernotupgradedCreateNewAggchain // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedCreateNewAggchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedCreateNewAggchain)
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
		it.Event = new(PolygonrollupmanagernotupgradedCreateNewAggchain)
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
func (it *PolygonrollupmanagernotupgradedCreateNewAggchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedCreateNewAggchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedCreateNewAggchain represents a CreateNewAggchain event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedCreateNewAggchain struct {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterCreateNewAggchain(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedCreateNewAggchainIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedCreateNewAggchainIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "CreateNewAggchain", logs: logs, sub: sub}, nil
}

// WatchCreateNewAggchain is a free log subscription operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchCreateNewAggchain(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedCreateNewAggchain, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedCreateNewAggchain)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseCreateNewAggchain(log types.Log) (*PolygonrollupmanagernotupgradedCreateNewAggchain, error) {
	event := new(PolygonrollupmanagernotupgradedCreateNewAggchain)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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

// PolygonrollupmanagernotupgradedInitMigrationIterator is returned from FilterInitMigration and is used to iterate over the raw logs and unpacked data for InitMigration events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedInitMigrationIterator struct {
	Event *PolygonrollupmanagernotupgradedInitMigration // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedInitMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedInitMigration)
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
		it.Event = new(PolygonrollupmanagernotupgradedInitMigration)
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
func (it *PolygonrollupmanagernotupgradedInitMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedInitMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedInitMigration represents a InitMigration event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedInitMigration struct {
	RollupID        uint32
	NewRollupTypeID uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitMigration is a free log retrieval operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterInitMigration(opts *bind.FilterOpts, rollupID []uint32) (*PolygonrollupmanagernotupgradedInitMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedInitMigrationIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "InitMigration", logs: logs, sub: sub}, nil
}

// WatchInitMigration is a free log subscription operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchInitMigration(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedInitMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedInitMigration)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "InitMigration", log); err != nil {
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

// ParseInitMigration is a log parse operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseInitMigration(log types.Log) (*PolygonrollupmanagernotupgradedInitMigration, error) {
	event := new(PolygonrollupmanagernotupgradedInitMigration)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "InitMigration", log); err != nil {
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

// PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator is returned from FilterVerifyPessimisticStateTransition and is used to iterate over the raw logs and unpacked data for VerifyPessimisticStateTransition events raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator struct {
	Event *PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition // Event containing the contract specifics and raw log

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
func (it *PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition)
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
		it.Event = new(PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition)
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
func (it *PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition represents a VerifyPessimisticStateTransition event raised by the Polygonrollupmanagernotupgraded contract.
type PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition struct {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) FilterVerifyPessimisticStateTransition(opts *bind.FilterOpts, rollupID []uint32, trustedAggregator []common.Address) (*PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.FilterLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonrollupmanagernotupgradedVerifyPessimisticStateTransitionIterator{contract: _Polygonrollupmanagernotupgraded.contract, event: "VerifyPessimisticStateTransition", logs: logs, sub: sub}, nil
}

// WatchVerifyPessimisticStateTransition is a free log subscription operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) WatchVerifyPessimisticStateTransition(opts *bind.WatchOpts, sink chan<- *PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition, rollupID []uint32, trustedAggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Polygonrollupmanagernotupgraded.contract.WatchLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition)
				if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
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
func (_Polygonrollupmanagernotupgraded *PolygonrollupmanagernotupgradedFilterer) ParseVerifyPessimisticStateTransition(log types.Log) (*PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition, error) {
	event := new(PolygonrollupmanagernotupgradedVerifyPessimisticStateTransition)
	if err := _Polygonrollupmanagernotupgraded.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayermanagernotupgraded

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

// AgglayerManagerRollupDataReturn is an auto generated low-level Go binding around an user-defined struct.
type AgglayerManagerRollupDataReturn struct {
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

// AgglayerManagerRollupDataReturnV2 is an auto generated low-level Go binding around an user-defined struct.
type AgglayerManagerRollupDataReturnV2 struct {
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

// LegacyZKEVMStateVariablesSequencedBatchData is an auto generated low-level Go binding around an user-defined struct.
type LegacyZKEVMStateVariablesSequencedBatchData struct {
	AccInputHash               [32]byte
	SequencedTimestamp         uint64
	PreviousLastBatchSequenced uint64
}

// AgglayermanagernotupgradedMetaData contains all meta data concerning the Agglayermanagernotupgraded contract.
var AgglayermanagernotupgradedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewRollupTypeMustBePessimisticOrALGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"CompletedMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"CreateNewAggchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"InitMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"}],\"name\":\"VerifyPessimisticStateTransition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"isRollupMigrating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structAgglayerManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataDeserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structAgglayerManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2Deserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b506040516161a43803806161a483398101604081905261002f9161019a565b838383836001600160a01b038416158061005057506001600160a01b038316155b8061006257506001600160a01b038216155b8061007457506001600160a01b038116155b15610092576040516342eda64b60e11b815260040160405180910390fd5b6001600160a01b0380851660805283811660c05282811660a052811660e0526100b96100c6565b50505050505050506101f6565b5f54610100900460ff16156101315760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610181575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610197575f5ffd5b50565b5f5f5f5f608085870312156101ad575f5ffd5b84516101b881610183565b60208601519094506101c981610183565b60408601519093506101da81610183565b60608601519092506101eb81610183565b939692955090935050565b60805160a05160c05160e051615f436102615f395f81816106ea0152610ebb01525f818161083001528181611d59015261373701525f81816106b6015281816127bf015261383401525f818161077601528181610a7601528181610d4b0152610fc70152615f435ff3fe608060405234801561000f575f5ffd5b506004361061025a575f3560e01c8063066ec0121461025e57806311f6b2871461028e5780631489ed10146102a157806315064c96146102b65780631796a1ae146102d35780632072f6c5146102f0578063248a9ca3146102f857806325280169146103195780632f2ff15d146103c957806330c27dde146103dc57806336568abe146103ef5780633a7094bd14610402578063477fa2701461042457806354fd4d501461042c57806355a71ee014610457578063604691691461049757806365c0504d1461049f5780636c7668771461051a578063706039091461052d5780637222020f146105db57806374d9c244146105ee5780637975fcfe1461060e5780637fb6e76a146106215780638129fc1c146106465780638fd88cc21461064e57806391d148541461066157806397d289a31461067457806399f5634e146106875780639a908e731461068f578063a217fddf146106a2578063a2967d99146106a9578063a3c573eb146106b1578063ab0475cf146106e5578063abcb51981461070c578063c1acbc341461071f578063c4c928c214610739578063ceee281d1461074c578063d02103ca14610771578063d5073f6f14610798578063d547741f146107ab578063d8905812146107be578063dbc16976146107e3578063dd0464b9146107eb578063dde0ff77146107fe578063dfdb8c5e14610818578063e46761c41461082b578063e4f3d8f914610852578063e764a37314610902578063e80e503014610915578063f4e9267514610928578063f8c8765e14610938578063f9c4c2ae1461094b575b5f5ffd5b608454610271906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b61027161029c366004614434565b61096b565b6102b46102af36600461447e565b61098e565b005b606f546102c39060ff1681565b6040519015158152602001610285565b607e546102e39063ffffffff1681565b604051610285919061450c565b6102b4610b3b565b61030b61030636600461451d565b610bf3565b604051908152602001610285565b610396610327366004614534565b60408051606080820183525f808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b03908116918301919091529282015190921690820152606001610285565b6102b46103d7366004614565565b610c07565b608754610271906001600160401b031681565b6102b46103fd366004614565565b610c28565b6102c3610410366004614434565b60886020525f908152604090205460ff1681565b60865461030b565b604080518082019091526006815265076312e302e360d41b60208201525b60405161028591906145c1565b61030b610465366004614534565b63ffffffff82165f9081526081602090815260408083206001600160401b038516845260020190915290205492915050565b61030b610c5f565b6105076104ad366004614434565b607f6020525f908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b6040516102859796959493929190614614565b6102b46105283660046146aa565b610c74565b6105c361053b366004614434565b63ffffffff165f9081526081602052604090208054600182015460058301546006840154600785015460088601546009909601546001600160a01b0380871698600160a01b978890046001600160401b03908116999288169890970487169680861695600160401b908190048216958281169591810490921693600160801b90920460ff1692565b6040516102859c9b9a9998979695949392919061474a565b6102b46105e9366004614434565b611193565b6106016105fc366004614434565b611271565b60405161028591906147dc565b61044a61061c3660046148b3565b6113c0565b6102e361062f36600461490e565b60836020525f908152604090205463ffffffff1681565b6102b46113f0565b6102b461065c366004614927565b6114c9565b6102c361066f366004614565565b611836565b6102b4610682366004614a06565b611860565b61030b611d55565b61027161069d366004614a5f565b611e31565b61030b5f81565b61030b61201d565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6040516102859190614a87565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6102b461071a366004614aa9565b6122cb565b60845461027190600160801b90046001600160401b031681565b6102b4610747366004614b4f565b6125cc565b6102e361075a366004614b7a565b60826020525f908152604090205463ffffffff1681565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6102b46107a636600461451d565b6126d4565b6102b46107b9366004614565565b61275f565b61044a60405180604001604052806006815260200165076312e302e360d41b81525081565b6102b461277b565b61044a6107f9366004614b95565b61282d565b60845461027190600160401b90046001600160401b031681565b6102b4610826366004614c00565b612854565b6106d87f000000000000000000000000000000000000000000000000000000000000000081565b6108ea610860366004614434565b63ffffffff165f90815260816020526040902080546001820154600583015460068401546007909401546001600160a01b0380851696600160a01b958690046001600160401b03908116979286169690950485169480831693600160401b808504831694600160801b808204851695600160c01b9092048516948085169493840416920460ff1690565b6040516102859c9b9a99989796959493929190614c2a565b6102b4610910366004614ccd565b612bd8565b6102b4610923366004614ce8565b612d45565b6080546102e39063ffffffff1681565b6102b4610946366004614d69565b613171565b61095e610959366004614434565b613381565b6040516102859190614dc2565b63ffffffff81165f908152608160205260408120610988906134d7565b92915050565b5f516020615eee5f395f51905f526109a5816134ee565b6001600160401b038816156109cd5760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610a0257610a026145e0565b14610a20576040516390db0d0760e01b815260040160405180910390fd5b610a2f818989898989896134f8565b600681018054600160401b600160801b031916600160401b6001600160401b038a16908102919091179091555f9081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610aab61201d565b6040518263ffffffff1660e01b8152600401610ac991815260200190565b5f604051808303815f87803b158015610ae0575f5ffd5b505af1158015610af2573d5f5f3e3d5ffd5b50505050336001600160a01b03168a63ffffffff165f516020615d4e5f395f51905f5289888a604051610b2793929190614eab565b60405180910390a350505050505050505050565b610b525f516020615eae5f395f51905f5233611836565b610be957608454600160801b90046001600160401b03161580610b9e57506084544290610b939062093a8090600160801b90046001600160401b0316614ee0565b6001600160401b0316115b80610bcb57506087544290610bc09062093a80906001600160401b0316614ee0565b6001600160401b0316115b15610be95760405163692baaad60e11b815260040160405180910390fd5b610bf1613832565b565b5f9081526034602052604090206001015490565b610c1082610bf3565b610c19816134ee565b610c2383836138a8565b505050565b6001600160a01b0381163314610c5157604051630b4ad1cd60e31b815260040160405180910390fd5b610c5b8282613910565b5050565b5f6086546064610c6f9190614eff565b905090565b5f516020615eee5f395f51905f52610c8b816134ee565b610c93613976565b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610cc857610cc86145e0565b03610ce657604051635b6602b760e01b815260040160405180910390fd5b60016007820154600160801b900460ff166002811115610d0857610d086145e0565b148015610d1457508215155b15610d3257604051630ded782d60e01b815260040160405180910390fd5b60405163ef4eeb3560e01b81525f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063ef4eeb3590610d80908d9060040161450c565b602060405180830381865afa158015610d9b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dbf9190614f16565b905080610ddf5760405163a60721e160e01b815260040160405180910390fd5b63ffffffff8b165f9081526088602052604090205460ff1615610e6b5781600501548914610e20576040516306c6486360e11b815260040160405180910390fd5b5f6005830181905563ffffffff8c1680825260886020526040808320805460ff191690555190917f6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e91a25b5f610e7b8c84848d8d8b8b6139bd565b905060026007840154600160801b900460ff166002811115610e9f57610e9f6145e0565b03610f255760405163a48fd37760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a48fd37790610ef49084908c908c90600401614f55565b5f6040518083038186803b158015610f0a575f5ffd5b505afa158015610f1c573d5f5f3e3d5ffd5b50505050610f8d565b6001830154600984015460405163020a49e360e51b81526001600160a01b03909216916341493c6091610f609185908d908d90600401614f84565b5f6040518083038186803b158015610f76575f5ffd5b505afa158015610f88573d5f5f3e3d5ffd5b505050505b60848054600160801b600160c01b031916600160801b426001600160401b031602179055600583018054908b9055600884018054908b90557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610ffc61201d565b6040518263ffffffff1660e01b815260040161101a91815260200190565b5f604051808303815f87803b158015611031575f5ffd5b505af1158015611043573d5f5f3e3d5ffd5b505050505f8c9050336001600160a01b03168f63ffffffff165f516020615d4e5f395f51905f525f5f5f1b8560405161107e93929190614eab565b60405180910390a3336001600160a01b03168f63ffffffff167fdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711848f87868b6040516110ec959493929190948552602085019390935260408401919091526060830152608082015260a00190565b60405180910390a360026007870154600160801b900460ff166002811115611116576111166145e0565b0361117a578554604051639ee4afa360e01b81526001600160a01b0390911690639ee4afa39061114c908c908c90600401614faf565b5f604051808303815f87803b158015611163575f5ffd5b505af1158015611175573d5f5f3e3d5ffd5b505050505b505050505050611188613b40565b505050505050505050565b5f516020615dae5f395f51905f526111aa816134ee565b63ffffffff821615806111c85750607e5463ffffffff908116908316115b156111e657604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001810154600160e81b900460ff161561122757604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b6112d460408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082019081526020015f81526020015f81525090565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b0380821686526001600160401b03600160a01b928390048116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b90819004821660c0850152600783015480831660e086015290810490911661010084015261012083019060ff600160801b909104166002811115611390576113906145e0565b908160028111156113a3576113a36145e0565b905250600881015461014083015260090154610160820152919050565b63ffffffff86165f9081526081602052604090206060906113e5908787878787613b57565b979650505050505050565b5f54600590610100900460ff1615801561141057505f5460ff8083169116105b6114355760405162461bcd60e51b815260040161142c90614fc2565b60405180910390fd5b5f805461ffff191660ff8316176101001790556040805180820182526006815265076312e302e360d41b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c291611491916145c1565b60405180910390a15f805461ff001916905560405160ff821681525f516020615e2e5f395f51905f529060200160405180910390a150565b6114d1613976565b6114e85f516020615d6e5f395f51905f5233611836565b1580156115665750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015611536573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061155a9190615010565b6001600160a01b031614155b1561158457604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b0382165f9081526082602052604081205463ffffffff16908190036115c3576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff1660028111156115f8576115f86145e0565b14611616576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b039081169084168111158061164e575060068201546001600160401b03600160401b9091048116908516105b1561166c5760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b031614611704576001600160401b038082165f908152600385016020526040902060010154600160401b900481169086168110156116d157604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091165f908152600384016020526040812090815560010180546001600160801b031916905561166e565b6006830180546001600160401b0319166001600160401b03871617905561172b858361502b565b608480545f906117459084906001600160401b031661502b565b82546101009290920a6001600160401b0381810219909316918316021790915586165f8181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece906044015f604051808303815f87803b1580156117be575f5ffd5b505af11580156117d0573d5f5f3e3d5ffd5b505050506001600160401b0385165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a350505050610c5b613b40565b5f9182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f516020615ece5f395f51905f52611877816134ee565b63ffffffff841615806118955750607e5463ffffffff908116908516115b156118b357604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff16156118f457604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b038516111561192257604051634c753f5760e01b815260040160405180910390fd5b6001600160401b0384165f9081526083602052604090205463ffffffff161561195e576040516337c8fe0960e11b815260040160405180910390fd5b608080545f919082906119769063ffffffff1661504a565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f80825260208201928390529394506001600160a01b039092169130916119c190614407565b6119cd9392919061506e565b604051809103905ff0801580156119e6573d5f5f3e3d5ffd5b506001600160401b0387165f818152608360209081526040808320805463ffffffff1990811663ffffffff8a81169182179093556001600160a01b038816808752608286528487208054909316821790925585526081909352922080546001600160e01b031916909117600160a01b909302929092178255600782018054600160401b600160801b03198116928c16600160401b02928317825560018801549495509293600160e01b900460ff16929091600160401b600160881b031990911660ff60801b1990911617600160801b836002811115611ac757611ac76145e0565b021790555060026001850154600160e01b900460ff166002811115611aee57611aee6145e0565b03611c0a578263ffffffff167f144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b89848a88600101601c9054906101000a900460ff166002811115611b4157611b416145e0565b8b604051611b539594939291906150a2565b60405180910390a28263ffffffff165f516020615dce5f395f51905f5289848a6001600160a01b03604051611b8b94939291906150e9565b60405180910390a25f86806020019051810190611ba89190615010565b60405163b3a326f760e01b81529091506001600160a01b0384169063b3a326f790611bd7908490600401614a87565b5f604051808303815f87803b158015611bee575f5ffd5b505af1158015611c00573d5f5f3e3d5ffd5b5050505050611d4b565b600184810180549183018054600160a01b600160e01b03198116600160a01b948590046001600160401b0316909402938417825591546001600160a01b03166001600160e01b03199092166001600160a01b0319909316929092171790556002808501545f808052918301602090815260408320919091556003860154600984015587518291829182918291611ca791908d018101908d0161516e565b945094509450945094508763ffffffff165f516020615dce5f395f51905f528e898f87604051611cda94939291906150e9565b60405180910390a2604051633892b81160e11b81526001600160a01b03881690637125702290611d1890889088908d9089908990899060040161520c565b5f604051808303815f87803b158015611d2f575f5ffd5b505af1158015611d41573d5f5f3e3d5ffd5b5050505050505050505b5050505050505050565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611da39190614a87565b602060405180830381865afa158015611dbe573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611de29190614f16565b6084549091505f90611e06906001600160401b03600160401b82048116911661502b565b6001600160401b03169050805f03611e20575f9250505090565b611e2a818361527e565b9250505090565b606f545f9060ff1615611e5757604051630bc011ff60e21b815260040160405180910390fd5b335f9081526082602052604081205463ffffffff1690819003611e8d576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03165f03611eb657604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff166002811115611eeb57611eeb6145e0565b14611f09576040516390db0d0760e01b815260040160405180910390fd5b608480548691905f90611f269084906001600160401b0316614ee0565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690505f611f598783614ee0565b6006840180546001600160401b038381166001600160401b03199092168217909255604080516060810182528a815242841660208083019182528886168385019081525f86815260038c018352859020935184559151600193909301805492518716600160401b026001600160801b03199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b6080545f9063ffffffff1680820361203657505f919050565b5f816001600160401b0381111561204f5761204f614943565b604051908082528060200260200182016040528015612078578160200160208202803683370190505b5090505f5b828110156120d55760815f612093836001615291565b63ffffffff1663ffffffff1681526020019081526020015f20600501548282815181106120c2576120c26152a4565b602090810291909101015260010161207d565b505f60205b83600114612277575f6120ee6002866152b8565b6120f960028761527e565b6121039190615291565b90505f816001600160401b0381111561211e5761211e614943565b604051908082528060200260200182016040528015612147578160200160208202803683370190505b5090505f5b828110156122505761215f6001846152cb565b8114801561217757506121736002886152b8565b6001145b156121cc576121a98661218b836002614eff565b8151811061219b5761219b6152a4565b602002602001015186613c8b565b8282815181106121bb576121bb6152a4565b602002602001018181525050612248565b612229866121db836002614eff565b815181106121eb576121eb6152a4565b6020026020010151878360026122019190614eff565b61220c906001615291565b8151811061221c5761221c6152a4565b6020026020010151613c8b565b82828151811061223b5761223b6152a4565b6020026020010181815250505b60010161214c565b508094508195506122618485613c8b565b93508261226d816152de565b93505050506120da565b5f835f8151811061228a5761228a6152a4565b602002602001015190505f5f90505b828110156122c1576122ab8285613c8b565b91506122b78485613c8b565b9350600101612299565b5095945050505050565b5f516020615dee5f395f51905f526122e2816134ee565b306001600160a01b0389160361230b5760405163325c055b60e21b815260040160405180910390fd5b607e80545f919082906123239063ffffffff1661504a565b91906101000a81548163ffffffff021916908363ffffffff1602179055905060016002811115612355576123556145e0565b866002811115612367576123676145e0565b0361239057841561238b576040516363d722e760e01b815260040160405180910390fd5b61244a565b60028660028111156123a4576123a46145e0565b036123fa576001600160a01b0388161515806123c857506001600160401b03871615155b806123d257508415155b806123dc57508215155b1561238b576040516363d722e760e01b815260040160405180910390fd5b5f86600281111561240d5761240d6145e0565b0361243157821561238b576040516363d722e760e01b815260040160405180910390fd5b6040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876002811115612494576124946145e0565b81525f602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b031990911617825591850151600182018054948701516001600160401b0316600160a01b026001600160e01b031990951691909316179290921780825592840151919260ff60e01b1916600160e01b836002811115612537576125376145e0565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a906125b9908c908c908c908c908c908c908c906152f3565b60405180910390a2505050505050505050565b5f516020615d6e5f395f51905f526125e3816134ee565b6001600160a01b0384165f9081526082602090815260408083205463ffffffff168084526081909252909120600263ffffffff86165f908152607f6020526040902060010154600160e01b900460ff166002811115612644576126446145e0565b141580156126a3575063ffffffff85165f908152607f6020526040902060010154600160e01b900460ff166002811115612680576126806145e0565b6007820154600160801b900460ff1660028111156126a0576126a06145e0565b14155b156126c157604051635aa0d5f160e11b815260040160405180910390fd5b6126cc868686613c99565b505050505050565b5f516020615d8e5f395f51905f526126eb816134ee565b683635c9adc5dea000008211806127055750633b9aca0082105b1561272357604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b61276882610bf3565b612771816134ee565b610c238383613910565b5f516020615e0e5f395f51905f52612792816134ee565b608780546001600160401b031916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc16976916004808301925f92919082900301818387803b15801561280c575f5ffd5b505af115801561281e573d5f5f3e3d5ffd5b5050505061282a613f40565b50565b63ffffffff86165f9081526081602052604090206060906113e590889088888888886139bd565b6001600160a01b0382165f9081526082602090815260408083205463ffffffff1683526081909152902060026007820154600160801b900460ff1660028111156128a0576128a06145e0565b03612a4957336001600160a01b0316836001600160a01b0316637388c4366040518163ffffffff1660e01b8152600401602060405180830381865afa1580156128eb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061290f9190615010565b6001600160a01b0316146129365760405163660a7ce560e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020908152604091829020548251636e7fbce960e01b815292516001600160a01b0390911692636e7fbce99260048083019391928290030181865afa15801561298e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129b29190615352565b6001600160f01b031916836001600160a01b0316636e7fbce96040518163ffffffff1660e01b8152600401602060405180830381865afa1580156129f8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a1c9190615352565b6001600160f01b03191614612a4457604051635aa0d5f160e11b815260040160405180910390fd5b612b11565b336001600160a01b0316836001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a8f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ab39190615010565b6001600160a01b031614612ada5760405163696072e960e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612b115760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b031610612b4c57604051633e37e23360e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020526040902060010154600160e01b900460ff166002811115612b7f57612b7f6145e0565b6007820154600160801b900460ff166002811115612b9f57612b9f6145e0565b14612bbd57604051635aa0d5f160e11b815260040160405180910390fd5b604080515f815260208101909152610c239084908490613c99565b5f516020615d6e5f395f51905f52612bef816134ee565b63ffffffff84165f908152608160205260408120906007820154600160801b900460ff166002811115612c2457612c246145e0565b14612c42576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612c795760405163664316a560e11b815260040160405180910390fd5b5f63ffffffff85165f908152607f6020526040902060010154600160e01b900460ff166002811115612cad57612cad6145e0565b03612ccb576040516301ea4e5b60e01b815260040160405180910390fd5b63ffffffff85165f908152608860205260409020805460ff191660011790558054612d00906001600160a01b03168585613c99565b8463ffffffff167f3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de985604051612d36919061450c565b60405180910390a25050505050565b5f516020615e4e5f395f51905f52612d5c816134ee565b6001600160401b0386165f9081526083602052604090205463ffffffff1615612d98576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b0387161115612dc657604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b0389165f9081526082602052604090205463ffffffff1615612e0257604051630d409b9360e41b815260040160405180910390fd5b6001600160a01b038916301480612e2157506001600160a01b0389163b155b15612e3f5760405163325c055b60e21b815260040160405180910390fd5b608080545f91908290612e579063ffffffff1661504a565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f896001600160401b03166001600160401b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8363ffffffff1663ffffffff1681526020019081526020015f2090508a815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550888160010160146101000a8154816001600160401b0302191690836001600160401b0316021790555089816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555087815f0160146101000a8154816001600160401b0302191690836001600160401b03160217905550858160070160106101000a81548160ff02191690836002811115612fde57612fde6145e0565b02179055506001866002811115612ff757612ff76145e0565b03613040576009810185905560088101849055600581018790556001600160a01b038a163b5f0361303b5760405163043103a360e21b815260040160405180910390fd5b61311a565b6002866002811115613054576130546145e0565b036130b3576001600160a01b038a1615158061307857506001600160401b03891615155b8061308257508415155b156130a057604051636fc5557f60e01b815260040160405180910390fd5b600881018490556005810187905561311a565b841515806130c057508315155b156130de57604051636fc5557f60e01b815260040160405180910390fd5b5f80805260028201602052604081208890556001600160a01b038b163b900361311a5760405163043103a360e21b815260040160405180910390fd5b8163ffffffff167f4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e8a8d8b8a5f8b8b60405161315c9796959493929190615379565b60405180910390a25050505050505050505050565b5f54600590610100900460ff1615801561319157505f5460ff8083169116105b6131ad5760405162461bcd60e51b815260040161142c90614fc2565b5f805461ffff191660ff83161761010017905567016345785d8a00006086556131d4613f97565b6131eb5f516020615eee5f395f51905f5286614001565b6131f55f84614001565b61320c5f516020615dee5f395f51905f5284614001565b6132235f516020615e4e5f395f51905f5284614001565b61323a5f516020615d6e5f395f51905f5284614001565b6132515f516020615dae5f395f51905f5285614001565b6132685f516020615ece5f395f51905f5285614001565b61327f5f516020615e0e5f395f51905f5285614001565b6132a97fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db185614001565b6132cd5f516020615eee5f395f51905f525f516020615d2e5f395f51905f5261400b565b6132e45f516020615d2e5f395f51905f5285614001565b6132fb5f516020615d8e5f395f51905f5285614001565b61331f5f516020615eae5f395f51905f525f516020615e6e5f395f51905f5261400b565b6133365f516020615eae5f395f51905f5283614001565b61334d5f516020615e6e5f395f51905f5283614001565b5f805461ff001916905560405160ff821681525f516020615e2e5f395f51905f529060200160405180910390a15050505050565b6133e560408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b03908116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b808204831660c0860152600160801b808304841660e0870152600160c01b909204831661010086015260078401548084166101208701529081049092166101408501526101608401910460ff1660028111156134ba576134ba6145e0565b908160028111156134cd576134cd6145e0565b8152505050919050565b60060154600160401b90046001600160401b031690565b61282a813361405c565b5f5f613503896134d7565b60078a01549091506001600160401b0390811690891610156135385760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b0388165f90815260028a016020526040902054915081613572576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b031611156135a557604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b0316116135d75760405163b9b18f5760e01b815260040160405180910390fd5b5f6135e68a8a8a8a878b613b57565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000160028360405161361a91906153d4565b602060405180830381855afa158015613635573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906136589190614f16565b61366291906152b8565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a916136a4918991906004016153ea565b602060405180830381865afa1580156136bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136e39190615426565b613700576040516309bde33960e01b815260040160405180910390fd5b5f61370b848b61502b565b905061375e87826001600160401b0316613723611d55565b61372d9190614eff565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190614083565b80608460088282829054906101000a90046001600160401b03166137829190614ee0565b82546101009290920a6001600160401b0381810219909316918316021790915560848054600160801b600160c01b031916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064015f604051808303815f87803b15801561380e575f5ffd5b505af1158015613820573d5f5f3e3d5ffd5b50505050505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b15801561388a575f5ffd5b505af115801561389c573d5f5f3e3d5ffd5b50505050610bf16140d5565b6138b28282611836565b610c5b575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b61391a8282611836565b15610c5b575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f516020615e8e5f395f51905f525c156139a357604051633ee5aeb560e01b815260040160405180910390fd5b610bf160015f516020615e8e5f395f51905f525b90614130565b606060026007880154600160801b900460ff1660028111156139e1576139e16145e0565b03613a93578654604051631a957d9b60e21b81525f916001600160a01b031690636a55f66c90613a179087908790600401614faf565b602060405180830381865afa158015613a32573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a569190614f16565b600589015460088a0154604051929350613a7c928a908d9086908c908c90602001615445565b6040516020818303038152906040529150506113e5565b865460408051632b47b7cd60e21b815290515f926001600160a01b03169163ad1edf349160048083019260209291908290030181865afa158015613ad9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613afd9190614f16565b600589015460088a0154604051929350613b23928a908d9086908c908c90602001615445565b604051602081830303815290604052915050979650505050505050565b610bf15f5f516020615e8e5f395f51905f526139b7565b6001600160401b038086165f8181526003890160205260408082205493881682529020546060929115801590613b8b575081155b15613ba95760405163340c614f60e11b815260040160405180910390fd5b80613bc7576040516366385b5160e01b815260040160405180910390fd5b613bd084614137565b613bed576040516305dae44f60e21b815260040160405180910390fd5b885460018a01546040516001600160601b03193360601b16602082015260348101889052605481018590526001600160c01b031960c08c811b82166074840152600160a01b94859004811b8216607c84015293909204831b82166084820152608c810187905260ac810184905260cc81018990529189901b1660ec82015260f401604051602081830303815290604052925050509695505050505050565b5f9182526020526040902090565b63ffffffff82161580613cb75750607e5463ffffffff908116908316115b15613cd557604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526082602052604081205463ffffffff1690819003613d14576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181165f908152608160205260409020600781015490918516600160401b9091046001600160401b031603613d6157604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615613da257604051633b8d3d9960e01b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825582546001600160401b03600160a01b9182900416026001600160e01b03199091169093179290921790915560038201546009840155600783018054600160401b63ffffffff891602600160401b600160801b0319821681178355925460ff600160e01b909104169260ff60801b1916600160401b600160881b031990911617600160801b836002811115613e5f57613e5f6145e0565b02179055505f613e6e8461096b565b6007840180546001600160401b0319166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef28692613ebf9216908990600401615482565b5f604051808303815f87803b158015613ed6575f5ffd5b505af1158015613ee8573d5f5f3e3d5ffd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff16613f6357604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b5f54610100900460ff16610bf15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840161142c565b610c5b82826138a8565b5f61401583610bf3565b5f84815260346020526040808220600101859055519192508391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6140668282611836565b610c5b57604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610c239084906141bc565b606f5460ff16156140f957604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b80825d5050565b5f67ffffffff000000016001600160401b03831610801561416c575067ffffffff00000001604083901c6001600160401b0316105b801561418c575067ffffffff00000001608083901c6001600160401b0316105b80156141a3575067ffffffff0000000160c083901c105b156141b057506001919050565b505f919050565b919050565b5f614210826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661428d9092919063ffffffff16565b805190915015610c23578080602001905181019061422e9190615426565b610c235760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161142c565b606061429b84845f856142a3565b949350505050565b6060824710156143045760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161142c565b5f5f866001600160a01b0316858760405161431f91906153d4565b5f6040518083038185875af1925050503d805f8114614359576040519150601f19603f3d011682016040523d82523d5f602084013e61435e565b606091505b50915091506113e587838387606083156143d85782515f036143d1576001600160a01b0385163b6143d15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161142c565b508161429b565b61429b83838151156143ed5781518083602001fd5b8060405162461bcd60e51b815260040161142c91906145c1565b610888806154a683390190565b6001600160401b03169052565b803563ffffffff811681146141b7575f5ffd5b5f60208284031215614444575f5ffd5b61444d82614421565b9392505050565b80356001600160401b03811681146141b7575f5ffd5b6001600160a01b038116811461282a575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b031215614496575f5ffd5b61449f89614421565b97506144ad60208a01614454565b96506144bb60408a01614454565b95506144c960608a01614454565b94506080890135935060a0890135925060c08901356144e78161446a565b91506103e089018a10156144f9575f5ffd5b60e0890190509295985092959890939650565b63ffffffff91909116815260200190565b5f6020828403121561452d575f5ffd5b5035919050565b5f5f60408385031215614545575f5ffd5b61454e83614421565b915061455c60208401614454565b90509250929050565b5f5f60408385031215614576575f5ffd5b8235915060208301356145888161446a565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61444d6020830184614593565b6001600160a01b03169052565b634e487b7160e01b5f52602160045260245ffd5b6003811061461057634e487b7160e01b5f52602160045260245ffd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e0810161464860608301876145f4565b931515608082015260a081019290925260c090910152949350505050565b5f5f83601f840112614676575f5ffd5b5081356001600160401b0381111561468c575f5ffd5b6020830191508360208285010111156146a3575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b0312156146c1575f5ffd5b6146ca89614421565b97506146d860208a01614421565b9650604089013595506060890135945060808901356001600160401b03811115614700575f5ffd5b61470c8b828c01614666565b90955093505060a08901356001600160401b0381111561472a575f5ffd5b6147368b828c01614666565b999c989b5096995094979396929594505050565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a9052881660a0820152610180810161479760c0830189614414565b6147a460e0830188614414565b6147b2610100830187614414565b6147c06101208301866145f4565b61014082019390935261016001529a9950505050505050505050565b5f610180820190506147ef8284516145d3565b60208301516148016020840182614414565b50604083015161481460408401826145d3565b5060608301516148276060840182614414565b506080830151608083015260a083015161484460a0840182614414565b5060c083015161485760c0840182614414565b5060e083015161486a60e0840182614414565b5061010083015161487f610100840182614414565b506101208301516148946101208401826145f4565b5061014083015161014083015261016083015161016083015292915050565b5f5f5f5f5f5f60c087890312156148c8575f5ffd5b6148d187614421565b95506148df60208801614454565b94506148ed60408801614454565b959894975094956060810135955060808101359460a0909101359350915050565b5f6020828403121561491e575f5ffd5b61444d82614454565b5f5f60408385031215614938575f5ffd5b823561454e8161446a565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561497f5761497f614943565b604052919050565b5f6001600160401b0382111561499f5761499f614943565b50601f01601f191660200190565b5f6149bf6149ba84614987565b614957565b90508281528383830111156149d2575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126149f7575f5ffd5b61444d838335602085016149ad565b5f5f5f60608486031215614a18575f5ffd5b614a2184614421565b9250614a2f60208501614454565b915060408401356001600160401b03811115614a49575f5ffd5b614a55868287016149e8565b9150509250925092565b5f5f60408385031215614a70575f5ffd5b614a7983614454565b946020939093013593505050565b6001600160a01b0391909116815260200190565b8035600381106141b7575f5ffd5b5f5f5f5f5f5f5f60e0888a031215614abf575f5ffd5b8735614aca8161446a565b96506020880135614ada8161446a565b9550614ae860408901614454565b9450614af660608901614a9b565b93506080880135925060a08801356001600160401b03811115614b17575f5ffd5b8801601f81018a13614b27575f5ffd5b614b368a8235602084016149ad565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614b61575f5ffd5b8335614b6c8161446a565b9250614a2f60208501614421565b5f60208284031215614b8a575f5ffd5b813561444d8161446a565b5f5f5f5f5f5f60a08789031215614baa575f5ffd5b614bb387614421565b955060208701359450604087013593506060870135925060808701356001600160401b03811115614be2575f5ffd5b614bee89828a01614666565b979a9699509497509295939492505050565b5f5f60408385031215614c11575f5ffd5b8235614c1c8161446a565b915061455c60208401614421565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a9052881660a08201526101808101614c7760c0830189614414565b614c8460e0830188614414565b614c92610100830187614414565b614ca0610120830186614414565b614cae610140830185614414565b614cbc6101608301846145f4565b9d9c50505050505050505050505050565b5f5f5f60608486031215614cdf575f5ffd5b614b6c84614421565b5f5f5f5f5f5f5f5f610100898b031215614d00575f5ffd5b8835614d0b8161446a565b97506020890135614d1b8161446a565b9650614d2960408a01614454565b9550614d3760608a01614454565b945060808901359350614d4c60a08a01614a9b565b979a969950949793969295929450505060c08201359160e0013590565b5f5f5f5f60808587031215614d7c575f5ffd5b8435614d878161446a565b93506020850135614d978161446a565b92506040850135614da78161446a565b91506060850135614db78161446a565b939692955090935050565b5f61018082019050614dd58284516145d3565b6020830151614de76020840182614414565b506040830151614dfa60408401826145d3565b506060830151614e0d6060840182614414565b506080830151608083015260a0830151614e2a60a0840182614414565b5060c0830151614e3d60c0840182614414565b5060e0830151614e5060e0840182614414565b50610100830151614e65610100840182614414565b50610120830151614e7a610120840182614414565b50610140830151614e8f610140840182614414565b50610160830151614ea46101608401826145f4565b5092915050565b6001600160401b039390931683526020830191909152604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03818116838216019081111561098857610988614ecc565b808202811582820484141761098857610988614ecc565b5f60208284031215614f26575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f614f676040830186614593565b8281036020840152614f7a818587614f2d565b9695505050505050565b848152606060208201525f614f9c6060830186614593565b82810360408401526113e5818587614f2d565b602081525f61429b602083018486614f2d565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f60208284031215615020575f5ffd5b815161444d8161446a565b6001600160401b03828116828216039081111561098857610988614ecc565b5f63ffffffff821663ffffffff810361506557615065614ecc565b60010192915050565b6001600160a01b038481168252831660208201526060604082018190525f9061509990830184614593565b95945050505050565b63ffffffff861681526001600160a01b03851660208201526001600160401b038416604082015260ff8316606082015260a0608082018190525f906113e590830184614593565b63ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b5f82601f830112615130575f5ffd5b815161513e6149ba82614987565b818152846020838601011115615152575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215615182575f5ffd5b855161518d8161446a565b602087015190955061519e8161446a565b60408701519094506151af8161446a565b60608701519093506001600160401b038111156151ca575f5ffd5b6151d688828901615121565b608088015190935090506001600160401b038111156151f3575f5ffd5b6151ff88828901615121565b9150509295509295909350565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c0608082018190525f9061524b90830185614593565b82810360a084015261525d8185614593565b9998505050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f8261528c5761528c61526a565b500490565b8082018082111561098857610988614ecc565b634e487b7160e01b5f52603260045260245ffd5b5f826152c6576152c661526a565b500690565b8181038181111561098857610988614ecc565b5f816152ec576152ec614ecc565b505f190190565b6001600160a01b038881168252871660208201526001600160401b038616604082015261532360608201866145f4565b83608082015260e060a08201525f61533e60e0830185614593565b90508260c083015298975050505050505050565b5f60208284031215615362575f5ffd5b81516001600160f01b03198116811461444d575f5ffd5b6001600160401b0388811682526001600160a01b03881660208301528616604082015260e081016153ad60608301876145f4565b6001600160401b0394909416608082015260a081019290925260c090910152949350505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b600181101561541c5781518352602092830192909101906001016153fd565b5050509392505050565b5f60208284031215615436575f5ffd5b8151801515811461444d575f5ffd5b9687526020870195909552604086019390935260e09190911b6001600160e01b03191660608501526064840152608483015260a482015260c40190565b6001600160a01b03831681526040602082018190525f9061429b9083018461459356fe60a060405260405161088838038061088883398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f5160206108685f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f5160206108685f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516104436104255f395f601001526104435ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610069575f356001600160e01b03191663278f794360e11b146100615761005f61006d565b565b61005f61007d565b61005f5b61005f6100786100ab565b6100cf565b5f8061008c36600481846102ba565b81019061009991906102f5565b915091506100a782826100ed565b5050565b5f6100ca5f5160206103ee5f395f51905f52546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156100e9573d5ff35b3d5ffd5b6100f682610147565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561013f5761013a82826101aa565b505050565b6100a761021c565b806001600160a01b03163b5f0361017c5780604051634c9c8ce360e01b815260040161017391906103c3565b60405180910390fd5b5f5160206103ee5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516101c691906103d7565b5f60405180830381855af49150503d805f81146101fe576040519150601f19603f3d011682016040523d82523d5f602084013e610203565b606091505b509150915061021385838361023b565b95945050505050565b341561005f5760405163b398979f60e01b815260040160405180910390fd5b6060826102505761024b82610291565b61028a565b815115801561026757506001600160a01b0384163b155b156102875783604051639996b31560e01b815260040161017391906103c3565b50805b9392505050565b8051156102a15780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f858511156102c8575f5ffd5b838611156102d4575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610306575f5ffd5b82356001600160a01b038116811461031c575f5ffd5b915060208301356001600160401b03811115610336575f5ffd5b8301601f81018513610346575f5ffd5b80356001600160401b0381111561035f5761035f6102e1565b604051601f8201601f19908116603f011681016001600160401b038111828210171561038d5761038d6102e1565b6040528181528282016020018710156103a4575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b6001600160a01b0391909116815260200190565b5f82518060208501845e5f92019182525091905056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220a0a1c3c04daffeaf76a25e5512528e424455aa0b0d045a8ab4ba83a4edfd9cea64736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610373cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0d1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d366156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fbab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641ac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59062ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f47f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024983dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff2859519b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545ea0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a264697066735822122024ed65dbba4373924460ded420d353864dc072b8c2c1d0a973bc7ce94f07340964736f6c634300081c0033",
}

// AgglayermanagernotupgradedABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayermanagernotupgradedMetaData.ABI instead.
var AgglayermanagernotupgradedABI = AgglayermanagernotupgradedMetaData.ABI

// AgglayermanagernotupgradedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayermanagernotupgradedMetaData.Bin instead.
var AgglayermanagernotupgradedBin = AgglayermanagernotupgradedMetaData.Bin

// DeployAgglayermanagernotupgraded deploys a new Ethereum contract, binding an instance of Agglayermanagernotupgraded to it.
func DeployAgglayermanagernotupgraded(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Agglayermanagernotupgraded, error) {
	parsed, err := AgglayermanagernotupgradedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayermanagernotupgradedBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayermanagernotupgraded{AgglayermanagernotupgradedCaller: AgglayermanagernotupgradedCaller{contract: contract}, AgglayermanagernotupgradedTransactor: AgglayermanagernotupgradedTransactor{contract: contract}, AgglayermanagernotupgradedFilterer: AgglayermanagernotupgradedFilterer{contract: contract}}, nil
}

// Agglayermanagernotupgraded is an auto generated Go binding around an Ethereum contract.
type Agglayermanagernotupgraded struct {
	AgglayermanagernotupgradedCaller     // Read-only binding to the contract
	AgglayermanagernotupgradedTransactor // Write-only binding to the contract
	AgglayermanagernotupgradedFilterer   // Log filterer for contract events
}

// AgglayermanagernotupgradedCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayermanagernotupgradedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagernotupgradedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayermanagernotupgradedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagernotupgradedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayermanagernotupgradedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagernotupgradedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayermanagernotupgradedSession struct {
	Contract     *Agglayermanagernotupgraded // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AgglayermanagernotupgradedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayermanagernotupgradedCallerSession struct {
	Contract *AgglayermanagernotupgradedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// AgglayermanagernotupgradedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayermanagernotupgradedTransactorSession struct {
	Contract     *AgglayermanagernotupgradedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// AgglayermanagernotupgradedRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayermanagernotupgradedRaw struct {
	Contract *Agglayermanagernotupgraded // Generic contract binding to access the raw methods on
}

// AgglayermanagernotupgradedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayermanagernotupgradedCallerRaw struct {
	Contract *AgglayermanagernotupgradedCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayermanagernotupgradedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayermanagernotupgradedTransactorRaw struct {
	Contract *AgglayermanagernotupgradedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayermanagernotupgraded creates a new instance of Agglayermanagernotupgraded, bound to a specific deployed contract.
func NewAgglayermanagernotupgraded(address common.Address, backend bind.ContractBackend) (*Agglayermanagernotupgraded, error) {
	contract, err := bindAgglayermanagernotupgraded(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayermanagernotupgraded{AgglayermanagernotupgradedCaller: AgglayermanagernotupgradedCaller{contract: contract}, AgglayermanagernotupgradedTransactor: AgglayermanagernotupgradedTransactor{contract: contract}, AgglayermanagernotupgradedFilterer: AgglayermanagernotupgradedFilterer{contract: contract}}, nil
}

// NewAgglayermanagernotupgradedCaller creates a new read-only instance of Agglayermanagernotupgraded, bound to a specific deployed contract.
func NewAgglayermanagernotupgradedCaller(address common.Address, caller bind.ContractCaller) (*AgglayermanagernotupgradedCaller, error) {
	contract, err := bindAgglayermanagernotupgraded(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedCaller{contract: contract}, nil
}

// NewAgglayermanagernotupgradedTransactor creates a new write-only instance of Agglayermanagernotupgraded, bound to a specific deployed contract.
func NewAgglayermanagernotupgradedTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayermanagernotupgradedTransactor, error) {
	contract, err := bindAgglayermanagernotupgraded(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedTransactor{contract: contract}, nil
}

// NewAgglayermanagernotupgradedFilterer creates a new log filterer instance of Agglayermanagernotupgraded, bound to a specific deployed contract.
func NewAgglayermanagernotupgradedFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayermanagernotupgradedFilterer, error) {
	contract, err := bindAgglayermanagernotupgraded(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedFilterer{contract: contract}, nil
}

// bindAgglayermanagernotupgraded binds a generic wrapper to an already deployed contract.
func bindAgglayermanagernotupgraded(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayermanagernotupgradedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanagernotupgraded.Contract.AgglayermanagernotupgradedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AgglayermanagernotupgradedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AgglayermanagernotupgradedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanagernotupgraded.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.DEFAULTADMINROLE(&_Agglayermanagernotupgraded.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.DEFAULTADMINROLE(&_Agglayermanagernotupgraded.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) ROLLUPMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "ROLLUP_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanagernotupgraded.Contract.ROLLUPMANAGERVERSION(&_Agglayermanagernotupgraded.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanagernotupgraded.Contract.ROLLUPMANAGERVERSION(&_Agglayermanagernotupgraded.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.AggLayerGateway(&_Agglayermanagernotupgraded.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.AggLayerGateway(&_Agglayermanagernotupgraded.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.BridgeAddress(&_Agglayermanagernotupgraded.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.BridgeAddress(&_Agglayermanagernotupgraded.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.CalculateRewardPerBatch(&_Agglayermanagernotupgraded.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.CalculateRewardPerBatch(&_Agglayermanagernotupgraded.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.ChainIDToRollupID(&_Agglayermanagernotupgraded.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.ChainIDToRollupID(&_Agglayermanagernotupgraded.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.GetBatchFee(&_Agglayermanagernotupgraded.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.GetBatchFee(&_Agglayermanagernotupgraded.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.GetForcedBatchFee(&_Agglayermanagernotupgraded.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanagernotupgraded.Contract.GetForcedBatchFee(&_Agglayermanagernotupgraded.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetInputPessimisticBytes(&_Agglayermanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetInputPessimisticBytes(&_Agglayermanagernotupgraded.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetInputSnarkBytes(&_Agglayermanagernotupgraded.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetInputSnarkBytes(&_Agglayermanagernotupgraded.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.GetLastVerifiedBatch(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.GetLastVerifiedBatch(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRoleAdmin(&_Agglayermanagernotupgraded.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRoleAdmin(&_Agglayermanagernotupgraded.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupExitRoot(&_Agglayermanagernotupgraded.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupExitRoot(&_Agglayermanagernotupgraded.CallOpts)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupSequencedBatches(&_Agglayermanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanagernotupgraded.Contract.GetRollupSequencedBatches(&_Agglayermanagernotupgraded.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.GlobalExitRootManager(&_Agglayermanagernotupgraded.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.GlobalExitRootManager(&_Agglayermanagernotupgraded.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanagernotupgraded.Contract.HasRole(&_Agglayermanagernotupgraded.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanagernotupgraded.Contract.HasRole(&_Agglayermanagernotupgraded.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) IsEmergencyState() (bool, error) {
	return _Agglayermanagernotupgraded.Contract.IsEmergencyState(&_Agglayermanagernotupgraded.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) IsEmergencyState() (bool, error) {
	return _Agglayermanagernotupgraded.Contract.IsEmergencyState(&_Agglayermanagernotupgraded.CallOpts)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) IsRollupMigrating(opts *bind.CallOpts, rollupID uint32) (bool, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "isRollupMigrating", rollupID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanagernotupgraded.Contract.IsRollupMigrating(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanagernotupgraded.Contract.IsRollupMigrating(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.LastAggregationTimestamp(&_Agglayermanagernotupgraded.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.LastAggregationTimestamp(&_Agglayermanagernotupgraded.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanagernotupgraded.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanagernotupgraded.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) Pol() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.Pol(&_Agglayermanagernotupgraded.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) Pol() (common.Address, error) {
	return _Agglayermanagernotupgraded.Contract.Pol(&_Agglayermanagernotupgraded.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupAddressToID(&_Agglayermanagernotupgraded.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupAddressToID(&_Agglayermanagernotupgraded.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupCount() (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupCount(&_Agglayermanagernotupgraded.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupCount() (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupCount(&_Agglayermanagernotupgraded.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturn)).(*AgglayerManagerRollupDataReturn)

	return out0, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupData(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupData(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupIDToRollupDataDeserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataDeserialized", rollupID)

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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupIDToRollupDataV2(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataV2", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturnV2), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturnV2)).(*AgglayerManagerRollupDataReturnV2)

	return out0, err

}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataV2(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataV2(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupIDToRollupDataV2Deserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupIDToRollupDataV2Deserialized", rollupID)

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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanagernotupgraded.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanagernotupgraded.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupTypeCount(&_Agglayermanagernotupgraded.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanagernotupgraded.Contract.RollupTypeCount(&_Agglayermanagernotupgraded.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanagernotupgraded.Contract.RollupTypeMap(&_Agglayermanagernotupgraded.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanagernotupgraded.Contract.RollupTypeMap(&_Agglayermanagernotupgraded.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.TotalSequencedBatches(&_Agglayermanagernotupgraded.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.TotalSequencedBatches(&_Agglayermanagernotupgraded.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.TotalVerifiedBatches(&_Agglayermanagernotupgraded.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanagernotupgraded.Contract.TotalVerifiedBatches(&_Agglayermanagernotupgraded.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanagernotupgraded.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) Version() (string, error) {
	return _Agglayermanagernotupgraded.Contract.Version(&_Agglayermanagernotupgraded.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedCallerSession) Version() (string, error) {
	return _Agglayermanagernotupgraded.Contract.Version(&_Agglayermanagernotupgraded.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.ActivateEmergencyState(&_Agglayermanagernotupgraded.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.ActivateEmergencyState(&_Agglayermanagernotupgraded.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AddExistingRollup(&_Agglayermanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AddExistingRollup(&_Agglayermanagernotupgraded.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AddNewRollupType(&_Agglayermanagernotupgraded.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AddNewRollupType(&_Agglayermanagernotupgraded.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AttachAggchainToAL(&_Agglayermanagernotupgraded.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.AttachAggchainToAL(&_Agglayermanagernotupgraded.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.DeactivateEmergencyState(&_Agglayermanagernotupgraded.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.DeactivateEmergencyState(&_Agglayermanagernotupgraded.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.GrantRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.GrantRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) InitMigration(opts *bind.TransactOpts, rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "initMigration", rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.InitMigration(&_Agglayermanagernotupgraded.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.InitMigration(&_Agglayermanagernotupgraded.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.Initialize(&_Agglayermanagernotupgraded.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.Initialize(&_Agglayermanagernotupgraded.TransactOpts)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) Initialize0(opts *bind.TransactOpts, trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "initialize0", trustedAggregator, admin, timelock, emergencyCouncil)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) Initialize0(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.Initialize0(&_Agglayermanagernotupgraded.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) Initialize0(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.Initialize0(&_Agglayermanagernotupgraded.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.ObsoleteRollupType(&_Agglayermanagernotupgraded.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.ObsoleteRollupType(&_Agglayermanagernotupgraded.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.OnSequenceBatches(&_Agglayermanagernotupgraded.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.OnSequenceBatches(&_Agglayermanagernotupgraded.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RenounceRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RenounceRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RevokeRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RevokeRole(&_Agglayermanagernotupgraded.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RollbackBatches(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.RollbackBatches(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.SetBatchFee(&_Agglayermanagernotupgraded.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.SetBatchFee(&_Agglayermanagernotupgraded.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.UpdateRollup(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.UpdateRollup(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.UpdateRollupByRollupAdmin(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.UpdateRollupByRollupAdmin(&_Agglayermanagernotupgraded.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanagernotupgraded.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanagernotupgraded.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagernotupgraded.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanagernotupgraded.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// AgglayermanagernotupgradedAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedAddExistingRollupIterator struct {
	Event *AgglayermanagernotupgradedAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedAddExistingRollup)
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
		it.Event = new(AgglayermanagernotupgradedAddExistingRollup)
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
func (it *AgglayermanagernotupgradedAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedAddExistingRollup represents a AddExistingRollup event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedAddExistingRollup struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedAddExistingRollupIterator{contract: _Agglayermanagernotupgraded.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedAddExistingRollup)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseAddExistingRollup(log types.Log) (*AgglayermanagernotupgradedAddExistingRollup, error) {
	event := new(AgglayermanagernotupgradedAddExistingRollup)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedAddNewRollupTypeIterator struct {
	Event *AgglayermanagernotupgradedAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedAddNewRollupType)
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
		it.Event = new(AgglayermanagernotupgradedAddNewRollupType)
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
func (it *AgglayermanagernotupgradedAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedAddNewRollupType represents a AddNewRollupType event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedAddNewRollupType struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagernotupgradedAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedAddNewRollupTypeIterator{contract: _Agglayermanagernotupgraded.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedAddNewRollupType)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseAddNewRollupType(log types.Log) (*AgglayermanagernotupgradedAddNewRollupType, error) {
	event := new(AgglayermanagernotupgradedAddNewRollupType)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedCompletedMigrationIterator is returned from FilterCompletedMigration and is used to iterate over the raw logs and unpacked data for CompletedMigration events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCompletedMigrationIterator struct {
	Event *AgglayermanagernotupgradedCompletedMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedCompletedMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedCompletedMigration)
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
		it.Event = new(AgglayermanagernotupgradedCompletedMigration)
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
func (it *AgglayermanagernotupgradedCompletedMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedCompletedMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedCompletedMigration represents a CompletedMigration event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCompletedMigration struct {
	RollupID uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompletedMigration is a free log retrieval operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterCompletedMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedCompletedMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedCompletedMigrationIterator{contract: _Agglayermanagernotupgraded.contract, event: "CompletedMigration", logs: logs, sub: sub}, nil
}

// WatchCompletedMigration is a free log subscription operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchCompletedMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedCompletedMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedCompletedMigration)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseCompletedMigration(log types.Log) (*AgglayermanagernotupgradedCompletedMigration, error) {
	event := new(AgglayermanagernotupgradedCompletedMigration)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedCreateNewAggchainIterator is returned from FilterCreateNewAggchain and is used to iterate over the raw logs and unpacked data for CreateNewAggchain events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCreateNewAggchainIterator struct {
	Event *AgglayermanagernotupgradedCreateNewAggchain // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedCreateNewAggchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedCreateNewAggchain)
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
		it.Event = new(AgglayermanagernotupgradedCreateNewAggchain)
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
func (it *AgglayermanagernotupgradedCreateNewAggchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedCreateNewAggchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedCreateNewAggchain represents a CreateNewAggchain event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCreateNewAggchain struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterCreateNewAggchain(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedCreateNewAggchainIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedCreateNewAggchainIterator{contract: _Agglayermanagernotupgraded.contract, event: "CreateNewAggchain", logs: logs, sub: sub}, nil
}

// WatchCreateNewAggchain is a free log subscription operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchCreateNewAggchain(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedCreateNewAggchain, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedCreateNewAggchain)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseCreateNewAggchain(log types.Log) (*AgglayermanagernotupgradedCreateNewAggchain, error) {
	event := new(AgglayermanagernotupgradedCreateNewAggchain)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCreateNewRollupIterator struct {
	Event *AgglayermanagernotupgradedCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedCreateNewRollup)
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
		it.Event = new(AgglayermanagernotupgradedCreateNewRollup)
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
func (it *AgglayermanagernotupgradedCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedCreateNewRollup represents a CreateNewRollup event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedCreateNewRollup struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedCreateNewRollupIterator{contract: _Agglayermanagernotupgraded.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedCreateNewRollup)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseCreateNewRollup(log types.Log) (*AgglayermanagernotupgradedCreateNewRollup, error) {
	event := new(AgglayermanagernotupgradedCreateNewRollup)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedEmergencyStateActivatedIterator struct {
	Event *AgglayermanagernotupgradedEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedEmergencyStateActivated)
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
		it.Event = new(AgglayermanagernotupgradedEmergencyStateActivated)
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
func (it *AgglayermanagernotupgradedEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedEmergencyStateActivated represents a EmergencyStateActivated event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*AgglayermanagernotupgradedEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedEmergencyStateActivatedIterator{contract: _Agglayermanagernotupgraded.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedEmergencyStateActivated)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseEmergencyStateActivated(log types.Log) (*AgglayermanagernotupgradedEmergencyStateActivated, error) {
	event := new(AgglayermanagernotupgradedEmergencyStateActivated)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedEmergencyStateDeactivatedIterator struct {
	Event *AgglayermanagernotupgradedEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedEmergencyStateDeactivated)
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
		it.Event = new(AgglayermanagernotupgradedEmergencyStateDeactivated)
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
func (it *AgglayermanagernotupgradedEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*AgglayermanagernotupgradedEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedEmergencyStateDeactivatedIterator{contract: _Agglayermanagernotupgraded.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedEmergencyStateDeactivated)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseEmergencyStateDeactivated(log types.Log) (*AgglayermanagernotupgradedEmergencyStateDeactivated, error) {
	event := new(AgglayermanagernotupgradedEmergencyStateDeactivated)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedInitMigrationIterator is returned from FilterInitMigration and is used to iterate over the raw logs and unpacked data for InitMigration events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedInitMigrationIterator struct {
	Event *AgglayermanagernotupgradedInitMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedInitMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedInitMigration)
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
		it.Event = new(AgglayermanagernotupgradedInitMigration)
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
func (it *AgglayermanagernotupgradedInitMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedInitMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedInitMigration represents a InitMigration event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedInitMigration struct {
	RollupID        uint32
	NewRollupTypeID uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitMigration is a free log retrieval operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterInitMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedInitMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedInitMigrationIterator{contract: _Agglayermanagernotupgraded.contract, event: "InitMigration", logs: logs, sub: sub}, nil
}

// WatchInitMigration is a free log subscription operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchInitMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedInitMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedInitMigration)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "InitMigration", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseInitMigration(log types.Log) (*AgglayermanagernotupgradedInitMigration, error) {
	event := new(AgglayermanagernotupgradedInitMigration)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "InitMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedInitializedIterator struct {
	Event *AgglayermanagernotupgradedInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedInitialized)
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
		it.Event = new(AgglayermanagernotupgradedInitialized)
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
func (it *AgglayermanagernotupgradedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedInitialized represents a Initialized event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayermanagernotupgradedInitializedIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedInitializedIterator{contract: _Agglayermanagernotupgraded.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedInitialized)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseInitialized(log types.Log) (*AgglayermanagernotupgradedInitialized, error) {
	event := new(AgglayermanagernotupgradedInitialized)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedObsoleteRollupTypeIterator struct {
	Event *AgglayermanagernotupgradedObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedObsoleteRollupType)
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
		it.Event = new(AgglayermanagernotupgradedObsoleteRollupType)
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
func (it *AgglayermanagernotupgradedObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedObsoleteRollupType represents a ObsoleteRollupType event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagernotupgradedObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedObsoleteRollupTypeIterator{contract: _Agglayermanagernotupgraded.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedObsoleteRollupType)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseObsoleteRollupType(log types.Log) (*AgglayermanagernotupgradedObsoleteRollupType, error) {
	event := new(AgglayermanagernotupgradedObsoleteRollupType)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedOnSequenceBatchesIterator struct {
	Event *AgglayermanagernotupgradedOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedOnSequenceBatches)
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
		it.Event = new(AgglayermanagernotupgradedOnSequenceBatches)
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
func (it *AgglayermanagernotupgradedOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedOnSequenceBatches represents a OnSequenceBatches event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedOnSequenceBatchesIterator{contract: _Agglayermanagernotupgraded.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedOnSequenceBatches)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseOnSequenceBatches(log types.Log) (*AgglayermanagernotupgradedOnSequenceBatches, error) {
	event := new(AgglayermanagernotupgradedOnSequenceBatches)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleAdminChangedIterator struct {
	Event *AgglayermanagernotupgradedRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedRoleAdminChanged)
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
		it.Event = new(AgglayermanagernotupgradedRoleAdminChanged)
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
func (it *AgglayermanagernotupgradedRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedRoleAdminChanged represents a RoleAdminChanged event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgglayermanagernotupgradedRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedRoleAdminChangedIterator{contract: _Agglayermanagernotupgraded.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedRoleAdminChanged)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseRoleAdminChanged(log types.Log) (*AgglayermanagernotupgradedRoleAdminChanged, error) {
	event := new(AgglayermanagernotupgradedRoleAdminChanged)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleGrantedIterator struct {
	Event *AgglayermanagernotupgradedRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedRoleGranted)
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
		it.Event = new(AgglayermanagernotupgradedRoleGranted)
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
func (it *AgglayermanagernotupgradedRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedRoleGranted represents a RoleGranted event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagernotupgradedRoleGrantedIterator, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedRoleGrantedIterator{contract: _Agglayermanagernotupgraded.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedRoleGranted)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseRoleGranted(log types.Log) (*AgglayermanagernotupgradedRoleGranted, error) {
	event := new(AgglayermanagernotupgradedRoleGranted)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleRevokedIterator struct {
	Event *AgglayermanagernotupgradedRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedRoleRevoked)
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
		it.Event = new(AgglayermanagernotupgradedRoleRevoked)
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
func (it *AgglayermanagernotupgradedRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedRoleRevoked represents a RoleRevoked event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagernotupgradedRoleRevokedIterator, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedRoleRevokedIterator{contract: _Agglayermanagernotupgraded.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedRoleRevoked)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseRoleRevoked(log types.Log) (*AgglayermanagernotupgradedRoleRevoked, error) {
	event := new(AgglayermanagernotupgradedRoleRevoked)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRollbackBatchesIterator struct {
	Event *AgglayermanagernotupgradedRollbackBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedRollbackBatches)
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
		it.Event = new(AgglayermanagernotupgradedRollbackBatches)
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
func (it *AgglayermanagernotupgradedRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedRollbackBatches represents a RollbackBatches event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*AgglayermanagernotupgradedRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedRollbackBatchesIterator{contract: _Agglayermanagernotupgraded.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedRollbackBatches)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseRollbackBatches(log types.Log) (*AgglayermanagernotupgradedRollbackBatches, error) {
	event := new(AgglayermanagernotupgradedRollbackBatches)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedSetBatchFeeIterator struct {
	Event *AgglayermanagernotupgradedSetBatchFee // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedSetBatchFee)
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
		it.Event = new(AgglayermanagernotupgradedSetBatchFee)
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
func (it *AgglayermanagernotupgradedSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedSetBatchFee represents a SetBatchFee event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*AgglayermanagernotupgradedSetBatchFeeIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedSetBatchFeeIterator{contract: _Agglayermanagernotupgraded.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedSetBatchFee)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseSetBatchFee(log types.Log) (*AgglayermanagernotupgradedSetBatchFee, error) {
	event := new(AgglayermanagernotupgradedSetBatchFee)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedSetTrustedAggregatorIterator struct {
	Event *AgglayermanagernotupgradedSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedSetTrustedAggregator)
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
		it.Event = new(AgglayermanagernotupgradedSetTrustedAggregator)
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
func (it *AgglayermanagernotupgradedSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedSetTrustedAggregator represents a SetTrustedAggregator event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*AgglayermanagernotupgradedSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedSetTrustedAggregatorIterator{contract: _Agglayermanagernotupgraded.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedSetTrustedAggregator)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseSetTrustedAggregator(log types.Log) (*AgglayermanagernotupgradedSetTrustedAggregator, error) {
	event := new(AgglayermanagernotupgradedSetTrustedAggregator)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedUpdateRollupIterator struct {
	Event *AgglayermanagernotupgradedUpdateRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedUpdateRollup)
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
		it.Event = new(AgglayermanagernotupgradedUpdateRollup)
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
func (it *AgglayermanagernotupgradedUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedUpdateRollup represents a UpdateRollup event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagernotupgradedUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedUpdateRollupIterator{contract: _Agglayermanagernotupgraded.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedUpdateRollup)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseUpdateRollup(log types.Log) (*AgglayermanagernotupgradedUpdateRollup, error) {
	event := new(AgglayermanagernotupgradedUpdateRollup)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedUpdateRollupManagerVersionIterator is returned from FilterUpdateRollupManagerVersion and is used to iterate over the raw logs and unpacked data for UpdateRollupManagerVersion events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedUpdateRollupManagerVersionIterator struct {
	Event *AgglayermanagernotupgradedUpdateRollupManagerVersion // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedUpdateRollupManagerVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedUpdateRollupManagerVersion)
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
		it.Event = new(AgglayermanagernotupgradedUpdateRollupManagerVersion)
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
func (it *AgglayermanagernotupgradedUpdateRollupManagerVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedUpdateRollupManagerVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedUpdateRollupManagerVersion represents a UpdateRollupManagerVersion event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedUpdateRollupManagerVersion struct {
	RollupManagerVersion string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupManagerVersion is a free log retrieval operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterUpdateRollupManagerVersion(opts *bind.FilterOpts) (*AgglayermanagernotupgradedUpdateRollupManagerVersionIterator, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedUpdateRollupManagerVersionIterator{contract: _Agglayermanagernotupgraded.contract, event: "UpdateRollupManagerVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupManagerVersion is a free log subscription operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchUpdateRollupManagerVersion(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedUpdateRollupManagerVersion) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedUpdateRollupManagerVersion)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseUpdateRollupManagerVersion(log types.Log) (*AgglayermanagernotupgradedUpdateRollupManagerVersion, error) {
	event := new(AgglayermanagernotupgradedUpdateRollupManagerVersion)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator struct {
	Event *AgglayermanagernotupgradedVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedVerifyBatchesTrustedAggregator)
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
		it.Event = new(AgglayermanagernotupgradedVerifyBatchesTrustedAggregator)
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
func (it *AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedVerifyBatchesTrustedAggregator struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedVerifyBatchesTrustedAggregatorIterator{contract: _Agglayermanagernotupgraded.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedVerifyBatchesTrustedAggregator)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*AgglayermanagernotupgradedVerifyBatchesTrustedAggregator, error) {
	event := new(AgglayermanagernotupgradedVerifyBatchesTrustedAggregator)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator is returned from FilterVerifyPessimisticStateTransition and is used to iterate over the raw logs and unpacked data for VerifyPessimisticStateTransition events raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator struct {
	Event *AgglayermanagernotupgradedVerifyPessimisticStateTransition // Event containing the contract specifics and raw log

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
func (it *AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagernotupgradedVerifyPessimisticStateTransition)
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
		it.Event = new(AgglayermanagernotupgradedVerifyPessimisticStateTransition)
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
func (it *AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagernotupgradedVerifyPessimisticStateTransition represents a VerifyPessimisticStateTransition event raised by the Agglayermanagernotupgraded contract.
type AgglayermanagernotupgradedVerifyPessimisticStateTransition struct {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) FilterVerifyPessimisticStateTransition(opts *bind.FilterOpts, rollupID []uint32, trustedAggregator []common.Address) (*AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.FilterLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagernotupgradedVerifyPessimisticStateTransitionIterator{contract: _Agglayermanagernotupgraded.contract, event: "VerifyPessimisticStateTransition", logs: logs, sub: sub}, nil
}

// WatchVerifyPessimisticStateTransition is a free log subscription operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) WatchVerifyPessimisticStateTransition(opts *bind.WatchOpts, sink chan<- *AgglayermanagernotupgradedVerifyPessimisticStateTransition, rollupID []uint32, trustedAggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanagernotupgraded.contract.WatchLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagernotupgradedVerifyPessimisticStateTransition)
				if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
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
func (_Agglayermanagernotupgraded *AgglayermanagernotupgradedFilterer) ParseVerifyPessimisticStateTransition(log types.Log) (*AgglayermanagernotupgradedVerifyPessimisticStateTransition, error) {
	event := new(AgglayermanagernotupgradedVerifyPessimisticStateTransition)
	if err := _Agglayermanagernotupgraded.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

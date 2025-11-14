// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayermanager

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

// AgglayermanagerMetaData contains all meta data concerning the Agglayermanager contract.
var AgglayermanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewRollupTypeMustBePessimisticOrALGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"CompletedMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"CreateNewAggchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"InitMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"}],\"name\":\"VerifyPessimisticStateTransition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"isRollupMigrating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structAgglayerManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataDeserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structAgglayerManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2Deserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b50604051615f0f380380615f0f83398101604081905261002f91610192565b6001600160a01b038416158061004c57506001600160a01b038316155b8061005e57506001600160a01b038216155b8061007057506001600160a01b038116155b1561008e576040516342eda64b60e11b815260040160405180910390fd5b6001600160a01b0380851660805283811660c05282811660a052811660e0526100b56100be565b505050506101ee565b5f54610100900460ff16156101295760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610179575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461018f575f5ffd5b50565b5f5f5f5f608085870312156101a5575f5ffd5b84516101b08161017b565b60208601519094506101c18161017b565b60408601519093506101d28161017b565b60608601519092506101e38161017b565b939692955090935050565b60805160a05160c05160e051615cb66102595f395f818161075f0152610f7301525f81816108a501528181611ee8015261371f01525f818161072b015281816129a9015261381d01525f81816107eb01528181610af801528181610e0601526110800152615cb65ff3fe608060405234801561000f575f5ffd5b50600436106102b8575f3560e01c806391d1485411610177578063d5073f6f116100d5578063dfdb8c5e1161008f578063dfdb8c5e1461088d578063e46761c4146108a0578063e4f3d8f9146108c7578063e764a37314610977578063e80e50301461098a578063f4e926751461099d578063f9c4c2ae146109ad575f5ffd5b8063d5073f6f1461080d578063d547741f14610820578063d890581214610833578063dbc1697614610858578063dd0464b914610860578063dde0ff7714610873575f5ffd5b8063a3c573eb11610131578063a3c573eb14610726578063ab0475cf1461075a578063abcb519814610781578063c1acbc3414610794578063c4c928c2146107ae578063ceee281d146107c1578063d02103ca146107e6575f5ffd5b806391d14854146106d657806397d289a3146106e957806399f5634e146106fc5780639a908e7314610704578063a217fddf14610717578063a2967d991461071e575f5ffd5b8063477fa2701161022457806370603909116101de57806370603909146105a25780637222020f1461065057806374d9c244146106635780637975fcfe146106835780637fb6e76a146106965780638129fc1c146106bb5780638fd88cc2146106c3575f5ffd5b8063477fa2701461049957806354fd4d50146104a157806355a71ee0146104cc578063604691691461050c57806365c0504d146105145780636c7668771461058f575f5ffd5b8063248a9ca311610275578063248a9ca31461035e578063252801691461038e5780632f2ff15d1461043e57806330c27dde1461045157806336568abe146104645780633a7094bd14610477575f5ffd5b8063066ec012146102bc57806311f6b287146102ec5780631489ed10146102ff57806315064c96146103145780631796a1ae146103315780632072f6c514610356575b5f5ffd5b6084546102cf906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b6102cf6102fa366004614310565b6109cd565b61031261030d36600461435a565b6109fc565b005b606f546103219060ff1681565b60405190151581526020016102e3565b607e546103419063ffffffff1681565b60405163ffffffff90911681526020016102e3565b610312610bd8565b61038061036c3660046143e8565b5f9081526034602052604090206001015490565b6040519081526020016102e3565b61040b61039c3660046143ff565b60408051606080820183525f808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b039081169183019190915292820151909216908201526060016102e3565b61031261044c366004614430565b610ca3565b6087546102cf906001600160401b031681565b610312610472366004614430565b610ccc565b610321610485366004614310565b60886020525f908152604090205460ff1681565b608654610380565b604080518082019091526006815265076312e302e360d41b60208201525b6040516102e3919061448c565b6103806104da3660046143ff565b63ffffffff82165f9081526081602090815260408083206001600160401b038516845260020190915290205492915050565b610380610d03565b61057c610522366004614310565b607f6020525f908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b6040516102e397969594939291906144d2565b61031261059d366004614568565b610d18565b6106386105b0366004614310565b63ffffffff165f9081526081602052604090208054600182015460058301546006840154600785015460088601546009909601546001600160a01b0380871698600160a01b978890046001600160401b03908116999288169890970487169680861695600160401b908190048216958281169591810490921693600160801b90920460ff1692565b6040516102e39c9b9a99989796959493929190614608565b61031261065e366004614310565b61127b565b610676610671366004614310565b61136c565b6040516102e39190614699565b6104bf6106913660046147a7565b6114bb565b6103416106a4366004614802565b60836020525f908152604090205463ffffffff1681565b6103126114eb565b6103126106d136600461481b565b61161a565b6103216106e4366004614430565b611988565b6103126106f73660046148fa565b6119b2565b610380611ee4565b6102cf610712366004614953565b611fc0565b6103805f81565b6103806121ad565b61074d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516102e3919061497b565b61074d7f000000000000000000000000000000000000000000000000000000000000000081565b61031261078f36600461499d565b612473565b6084546102cf90600160801b90046001600160401b031681565b6103126107bc366004614a43565b612787565b6103416107cf366004614a6e565b60826020525f908152604090205463ffffffff1681565b61074d7f000000000000000000000000000000000000000000000000000000000000000081565b61031261081b3660046143e8565b61288f565b61031261082e366004614430565b61292d565b6104bf60405180604001604052806006815260200165076312e302e360d41b81525081565b610312612951565b6104bf61086e366004614a89565b612a17565b6084546102cf90600160401b90046001600160401b031681565b61031261089b366004614af4565b612a3e565b61074d7f000000000000000000000000000000000000000000000000000000000000000081565b61095f6108d5366004614310565b63ffffffff165f90815260816020526040902080546001820154600583015460068401546007909401546001600160a01b0380851696600160a01b958690046001600160401b03908116979286169690950485169480831693600160401b808504831694600160801b808204851695600160c01b9092048516948085169493840416920460ff1690565b6040516102e39c9b9a99989796959493929190614b1e565b610312610985366004614bc4565b612dc2565b610312610998366004614bdf565b612f2e565b6080546103419063ffffffff1681565b6109c06109bb366004614310565b61336d565b6040516102e39190614c60565b63ffffffff81165f90815260816020526040812060060154600160401b90046001600160401b03165b92915050565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4610a26816134c3565b6001600160401b03881615610a4e5760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610a8357610a8361449e565b14610aa1576040516390db0d0760e01b815260040160405180910390fd5b610ab0818989898989896134cd565b60068101805467ffffffffffffffff60401b1916600160401b6001600160401b038a16908102919091179091555f9081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610b2d6121ad565b6040518263ffffffff1660e01b8152600401610b4b91815260200190565b5f604051808303815f87803b158015610b62575f5ffd5b505af1158015610b74573d5f5f3e3d5ffd5b5050604080516001600160401b038b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b610c027f141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e33611988565b610c9957608454600160801b90046001600160401b03161580610c4e57506084544290610c439062093a8090600160801b90046001600160401b0316614da4565b6001600160401b0316115b80610c7b57506087544290610c709062093a80906001600160401b0316614da4565b6001600160401b0316115b15610c995760405163692baaad60e11b815260040160405180910390fd5b610ca161381b565b565b5f82815260346020526040902060010154610cbd816134c3565b610cc78383613891565b505050565b6001600160a01b0381163314610cf557604051630b4ad1cd60e31b815260040160405180910390fd5b610cff82826138f9565b5050565b5f6086546064610d139190614dc3565b905090565b7f084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4610d42816134c3565b610d4a61395f565b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610d7f57610d7f61449e565b03610d9d57604051635b6602b760e01b815260040160405180910390fd5b60016007820154600160801b900460ff166002811115610dbf57610dbf61449e565b148015610dcb57508215155b15610de957604051630ded782d60e01b815260040160405180910390fd5b60405163ef4eeb3560e01b815263ffffffff8a1660048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ef4eeb3590602401602060405180830381865afa158015610e53573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e779190614dda565b905080610e975760405163a60721e160e01b815260040160405180910390fd5b63ffffffff8b165f9081526088602052604090205460ff1615610f235781600501548914610ed8576040516306c6486360e11b815260040160405180910390fd5b5f6005830181905563ffffffff8c1680825260886020526040808320805460ff191690555190917f6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e91a25b5f610f338c84848d8d8b8b6139a6565b905060026007840154600160801b900460ff166002811115610f5757610f5761449e565b03610fdd5760405163a48fd37760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a48fd37790610fac9084908c908c90600401614e19565b5f6040518083038186803b158015610fc2575f5ffd5b505afa158015610fd4573d5f5f3e3d5ffd5b50505050611045565b6001830154600984015460405163020a49e360e51b81526001600160a01b03909216916341493c60916110189185908d908d90600401614e48565b5f6040518083038186803b15801561102e575f5ffd5b505afa158015611040573d5f5f3e3d5ffd5b505050505b6084805467ffffffffffffffff60801b1916600160801b426001600160401b031602179055600583018054908b9055600884018054908b90557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d6110b56121ad565b6040518263ffffffff1660e01b81526004016110d391815260200190565b5f604051808303815f87803b1580156110ea575f5ffd5b505af11580156110fc573d5f5f3e3d5ffd5b505050505f8c9050336001600160a01b03168f63ffffffff167fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d35f5f5f1b85604051611166939291906001600160401b039390931683526020830191909152604082015260600190565b60405180910390a3336001600160a01b03168f63ffffffff167fdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711848f87868b6040516111d4959493929190948552602085019390935260408401919091526060830152608082015260a00190565b60405180910390a360026007870154600160801b900460ff1660028111156111fe576111fe61449e565b03611262578554604051639ee4afa360e01b81526001600160a01b0390911690639ee4afa390611234908c908c90600401614e73565b5f604051808303815f87803b15801561124b575f5ffd5b505af115801561125d573d5f5f3e3d5ffd5b505050505b505050505050611270613b29565b505050505050505050565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd6112a5816134c3565b63ffffffff821615806112c35750607e5463ffffffff908116908316115b156112e157604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001810154600160e81b900460ff161561132257604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b6113cf60408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082019081526020015f81526020015f81525090565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b0380821686526001600160401b03600160a01b928390048116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b90819004821660c0850152600783015480831660e086015290810490911661010084015261012083019060ff600160801b90910416600281111561148b5761148b61449e565b9081600281111561149e5761149e61449e565b905250600881015461014083015260090154610160820152919050565b63ffffffff86165f9081526081602052604090206060906114e0908787878787613b40565b979650505050505050565b5f54600590610100900460ff1615801561150b57505f5460ff8083169116105b6115735760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805461ffff191660ff8316176101001790556040805180820182526006815265076312e302e360d41b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2916115cf9161448c565b60405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b61162261395f565b6116395f516020615c415f395f51905f5233611988565b1580156116b75750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015611687573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116ab9190614e86565b6001600160a01b031614155b156116d557604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b0382165f9081526082602052604081205463ffffffff1690819003611714576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff1660028111156117495761174961449e565b14611767576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b039081169084168111158061179f575060068201546001600160401b03600160401b9091048116908516105b156117bd5760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b031614611855576001600160401b038082165f908152600385016020526040902060010154600160401b9004811690861681101561182257604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091165f908152600384016020526040812090815560010180546001600160801b03191690556117bf565b60068301805467ffffffffffffffff19166001600160401b03871617905561187d8583614ea1565b608480545f906118979084906001600160401b0316614ea1565b82546101009290920a6001600160401b0381810219909316918316021790915586165f8181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece906044015f604051808303815f87803b158015611910575f5ffd5b505af1158015611922573d5f5f3e3d5ffd5b505050506001600160401b0385165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a350505050610cff613b29565b5f9182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd086119dc816134c3565b63ffffffff841615806119fa5750607e5463ffffffff908116908516115b15611a1857604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615611a5957604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b0385161115611a8757604051634c753f5760e01b815260040160405180910390fd5b6001600160401b0384165f9081526083602052604090205463ffffffff1615611ac3576040516337c8fe0960e11b815260040160405180910390fd5b608080545f91908290611adb9063ffffffff16614ec0565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f80825260208201928390529394506001600160a01b03909216913091611b26906142f0565b611b3293929190614ee4565b604051809103905ff080158015611b4b573d5f5f3e3d5ffd5b506001600160401b0387165f818152608360209081526040808320805463ffffffff1990811663ffffffff8a81169182179093556001600160a01b038816808752608286528487208054909316821790925585526081909352922080546001600160e01b031916909117600160a01b90930292909217825560078201805467ffffffffffffffff60401b198116928c16600160401b02928317825560018801549495509293600160e01b900460ff1692909168ffffffffffffffffff60401b1990911660ff60801b1990911617600160801b836002811115611c2f57611c2f61449e565b021790555060026001850154600160e01b900460ff166002811115611c5657611c5661449e565b03611d85578263ffffffff167f144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b89848a88600101601c9054906101000a900460ff166002811115611ca957611ca961449e565b8b604051611cbb959493929190614f18565b60405180910390a28263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a64189848a6001600160a01b03604051611d069493929190614f5f565b60405180910390a25f86806020019051810190611d239190614e86565b60405163b3a326f760e01b81529091506001600160a01b0384169063b3a326f790611d5290849060040161497b565b5f604051808303815f87803b158015611d69575f5ffd5b505af1158015611d7b573d5f5f3e3d5ffd5b5050505050611eda565b60018481018054918301805467ffffffffffffffff60a01b198116600160a01b948590046001600160401b0316909402938417825591546001600160a01b03166001600160e01b03199092166001600160a01b0319909316929092171790556002808501545f808052918301602090815260408320919091556003860154600984015587518291829182918291611e2391908d018101908d01614fe4565b945094509450945094508763ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418e898f87604051611e699493929190614f5f565b60405180910390a2604051633892b81160e11b81526001600160a01b03881690637125702290611ea790889088908d90899089908990600401615080565b5f604051808303815f87803b158015611ebe575f5ffd5b505af1158015611ed0573d5f5f3e3d5ffd5b5050505050505050505b5050505050505050565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611f32919061497b565b602060405180830381865afa158015611f4d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f719190614dda565b6084549091505f90611f95906001600160401b03600160401b820481169116614ea1565b6001600160401b03169050805f03611faf575f9250505090565b611fb981836150f2565b9250505090565b606f545f9060ff1615611fe657604051630bc011ff60e21b815260040160405180910390fd5b335f9081526082602052604081205463ffffffff169081900361201c576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03165f0361204557604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff16600281111561207a5761207a61449e565b14612098576040516390db0d0760e01b815260040160405180910390fd5b608480548691905f906120b59084906001600160401b0316614da4565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690505f6120e88783614da4565b6006840180546001600160401b0383811667ffffffffffffffff199092168217909255604080516060810182528a815242841660208083019182528886168385019081525f86815260038c018352859020935184559151600193909301805492518716600160401b026001600160801b03199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b6080545f9063ffffffff168082036121c657505f919050565b5f816001600160401b038111156121df576121df614837565b604051908082528060200260200182016040528015612208578160200160208202803683370190505b5090505f5b828110156122655760815f612223836001615105565b63ffffffff1663ffffffff1681526020019081526020015f206005015482828151811061225257612252615118565b602090810291909101015260010161220d565b505f60205b83600114612422575f61227e60028661512c565b6122896002876150f2565b6122939190615105565b90505f816001600160401b038111156122ae576122ae614837565b6040519080825280602002602001820160405280156122d7578160200160208202803683370190505b5090505f5b828110156123f2576122ef60018461513f565b81148015612307575061230360028861512c565b6001145b15612365576123428661231b836002614dc3565b8151811061232b5761232b615118565b6020026020010151865f9182526020526040902090565b82828151811061235457612354615118565b6020026020010181815250506123ea565b6123cb86612374836002614dc3565b8151811061238457612384615118565b60200260200101518783600261239a9190614dc3565b6123a5906001615105565b815181106123b5576123b5615118565b60200260200101515f9182526020526040902090565b8282815181106123dd576123dd615118565b6020026020010181815250505b6001016122dc565b5080945081955061240c84855f9182526020526040902090565b93508261241881615152565b935050505061226a565b5f835f8151811061243557612435615118565b602002602001015190505f5f90505b82811015612469575f9182526020849052604080832085845292209350600101612444565b5095945050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59061249d816134c3565b306001600160a01b038916036124c65760405163325c055b60e21b815260040160405180910390fd5b607e80545f919082906124de9063ffffffff16614ec0565b91906101000a81548163ffffffff021916908363ffffffff16021790559050600160028111156125105761251061449e565b8660028111156125225761252261449e565b0361254b578415612546576040516363d722e760e01b815260040160405180910390fd5b612605565b600286600281111561255f5761255f61449e565b036125b5576001600160a01b03881615158061258357506001600160401b03871615155b8061258d57508415155b8061259757508215155b15612546576040516363d722e760e01b815260040160405180910390fd5b5f8660028111156125c8576125c861449e565b036125ec578215612546576040516363d722e760e01b815260040160405180910390fd5b6040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b0316815260200187600281111561264f5761264f61449e565b81525f602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b031990911617825591850151600182018054948701516001600160401b0316600160a01b026001600160e01b031990951691909316179290921780825592840151919260ff60e01b1916600160e01b8360028111156126f2576126f261449e565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a90612774908c908c908c908c908c908c908c90615167565b60405180910390a2505050505050505050565b5f516020615c415f395f51905f5261279e816134c3565b6001600160a01b0384165f9081526082602090815260408083205463ffffffff168084526081909252909120600263ffffffff86165f908152607f6020526040902060010154600160e01b900460ff1660028111156127ff576127ff61449e565b1415801561285e575063ffffffff85165f908152607f6020526040902060010154600160e01b900460ff16600281111561283b5761283b61449e565b6007820154600160801b900460ff16600281111561285b5761285b61449e565b14155b1561287c57604051635aa0d5f160e11b815260040160405180910390fd5b612887868686613c43565b505050505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb6128b9816134c3565b683635c9adc5dea000008211806128d35750633b9aca0082105b156128f157604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b5f82815260346020526040902060010154612947816134c3565b610cc783836138f9565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f461297b816134c3565b6087805467ffffffffffffffff1916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc16976916004808301925f92919082900301818387803b1580156129f6575f5ffd5b505af1158015612a08573d5f5f3e3d5ffd5b50505050612a14613eee565b50565b63ffffffff86165f9081526081602052604090206060906114e090889088888888886139a6565b6001600160a01b0382165f9081526082602090815260408083205463ffffffff1683526081909152902060026007820154600160801b900460ff166002811115612a8a57612a8a61449e565b03612c3357336001600160a01b0316836001600160a01b0316637388c4366040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ad5573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612af99190614e86565b6001600160a01b031614612b205760405163660a7ce560e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020908152604091829020548251636e7fbce960e01b815292516001600160a01b0390911692636e7fbce99260048083019391928290030181865afa158015612b78573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b9c91906151c6565b6001600160f01b031916836001600160a01b0316636e7fbce96040518163ffffffff1660e01b8152600401602060405180830381865afa158015612be2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c0691906151c6565b6001600160f01b03191614612c2e57604051635aa0d5f160e11b815260040160405180910390fd5b612cfb565b336001600160a01b0316836001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015612c79573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c9d9190614e86565b6001600160a01b031614612cc45760405163696072e960e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612cfb5760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b031610612d3657604051633e37e23360e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020526040902060010154600160e01b900460ff166002811115612d6957612d6961449e565b6007820154600160801b900460ff166002811115612d8957612d8961449e565b14612da757604051635aa0d5f160e11b815260040160405180910390fd5b604080515f815260208101909152610cc79084908490613c43565b5f516020615c415f395f51905f52612dd9816134c3565b63ffffffff84165f908152608160205260408120906007820154600160801b900460ff166002811115612e0e57612e0e61449e565b14612e2c576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b9092041614612e635760405163664316a560e11b815260040160405180910390fd5b5f63ffffffff85165f908152607f6020526040902060010154600160e01b900460ff166002811115612e9757612e9761449e565b03612eb5576040516301ea4e5b60e01b815260040160405180910390fd5b63ffffffff85165f908152608860205260409020805460ff191660011790558054612eea906001600160a01b03168585613c43565b60405163ffffffff85811682528616907f3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de99060200160405180910390a25050505050565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e612f58816134c3565b6001600160401b0386165f9081526083602052604090205463ffffffff1615612f94576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b0387161115612fc257604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b0389165f9081526082602052604090205463ffffffff1615612ffe57604051630d409b9360e41b815260040160405180910390fd5b6001600160a01b03891630148061301d57506001600160a01b0389163b155b1561303b5760405163325c055b60e21b815260040160405180910390fd5b608080545f919082906130539063ffffffff16614ec0565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f896001600160401b03166001600160401b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8363ffffffff1663ffffffff1681526020019081526020015f2090508a815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550888160010160146101000a8154816001600160401b0302191690836001600160401b0316021790555089816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555087815f0160146101000a8154816001600160401b0302191690836001600160401b03160217905550858160070160106101000a81548160ff021916908360028111156131da576131da61449e565b021790555060018660028111156131f3576131f361449e565b0361323c576009810185905560088101849055600581018790556001600160a01b038a163b5f036132375760405163043103a360e21b815260040160405180910390fd5b613316565b60028660028111156132505761325061449e565b036132af576001600160a01b038a1615158061327457506001600160401b03891615155b8061327e57508415155b1561329c57604051636fc5557f60e01b815260040160405180910390fd5b6008810184905560058101879055613316565b841515806132bc57508315155b156132da57604051636fc5557f60e01b815260040160405180910390fd5b5f80805260028201602052604081208890556001600160a01b038b163b90036133165760405163043103a360e21b815260040160405180910390fd5b8163ffffffff167f4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e8a8d8b8a5f8b8b60405161335897969594939291906151ed565b60405180910390a25050505050505050505050565b6133d160408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b03908116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b808204831660c0860152600160801b808304841660e0870152600160c01b909204831661010086015260078401548084166101208701529081049092166101408501526101608401910460ff1660028111156134a6576134a661449e565b908160028111156134b9576134b961449e565b8152505050919050565b612a148133613f45565b5f5f6134eb89600601546001600160401b03600160401b9091041690565b60078a01549091506001600160401b0390811690891610156135205760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b0388165f90815260028a01602052604090205491508161355a576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b0316111561358d57604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b0316116135bf5760405163b9b18f5760e01b815260040160405180910390fd5b5f6135ce8a8a8a8a878b613b40565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516136029190615248565b602060405180830381855afa15801561361d573d5f5f3e3d5ffd5b5050506040513d601f19601f820116820180604052508101906136409190614dda565b61364a919061512c565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a9161368c9189919060040161525e565b602060405180830381865afa1580156136a7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136cb919061529a565b6136e8576040516309bde33960e01b815260040160405180910390fd5b5f6136f3848b614ea1565b905061374687826001600160401b031661370b611ee4565b6137159190614dc3565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190613f6c565b80608460088282829054906101000a90046001600160401b031661376a9190614da4565b82546101009290920a6001600160401b038181021990931691831602179091556084805467ffffffffffffffff60801b1916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064015f604051808303815f87803b1580156137f7575f5ffd5b505af1158015613809573d5f5f3e3d5ffd5b50505050505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613873575f5ffd5b505af1158015613885573d5f5f3e3d5ffd5b50505050610ca1613fbe565b61389b8282611988565b610cff575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6139038282611988565b15610cff575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f516020615c615f395f51905f525c1561398c57604051633ee5aeb560e01b815260040160405180910390fd5b610ca160015f516020615c615f395f51905f525b90614019565b606060026007880154600160801b900460ff1660028111156139ca576139ca61449e565b03613a7c578654604051631a957d9b60e21b81525f916001600160a01b031690636a55f66c90613a009087908790600401614e73565b602060405180830381865afa158015613a1b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a3f9190614dda565b600589015460088a0154604051929350613a65928a908d9086908c908c906020016152b9565b6040516020818303038152906040529150506114e0565b865460408051632b47b7cd60e21b815290515f926001600160a01b03169163ad1edf349160048083019260209291908290030181865afa158015613ac2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ae69190614dda565b600589015460088a0154604051929350613b0c928a908d9086908c908c906020016152b9565b604051602081830303815290604052915050979650505050505050565b610ca15f5f516020615c615f395f51905f526139a0565b6001600160401b038086165f8181526003890160205260408082205493881682529020546060929115801590613b74575081155b15613b925760405163340c614f60e11b815260040160405180910390fd5b80613bb0576040516366385b5160e01b815260040160405180910390fd5b613bb984614020565b613bd6576040516305dae44f60e21b815260040160405180910390fd5b3385838a8c5f0160149054906101000a90046001600160401b03168d60010160149054906101000a90046001600160401b031689878d8f604051602001613c269a999897969594939291906152f6565b604051602081830303815290604052925050509695505050505050565b63ffffffff82161580613c615750607e5463ffffffff908116908316115b15613c7f57604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526082602052604081205463ffffffff1690819003613cbe576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181165f908152608160205260409020600781015490918516600160401b9091046001600160401b031603613d0b57604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615613d4c57604051633b8d3d9960e01b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825582546001600160401b03600160a01b9182900416026001600160e01b03199091169093179290921790915560038201546009840155600783018054600160401b63ffffffff89160267ffffffffffffffff60401b19821681178355925460ff600160e01b909104169260ff60801b191668ffffffffffffffffff60401b1990911617600160801b836002811115613e0c57613e0c61449e565b02179055505f613e1b846109cd565b60078401805467ffffffffffffffff19166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef28692613e6d9216908990600401615392565b5f604051808303815f87803b158015613e84575f5ffd5b505af1158015613e96573d5f5f3e3d5ffd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff16613f1157604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b613f4f8282611988565b610cff57604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610cc79084906140a5565b606f5460ff1615613fe257604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b80825d5050565b5f67ffffffff000000016001600160401b038316108015614055575067ffffffff00000001604083901c6001600160401b0316105b8015614075575067ffffffff00000001608083901c6001600160401b0316105b801561408c575067ffffffff0000000160c083901c105b1561409957506001919050565b505f919050565b919050565b5f6140f9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166141769092919063ffffffff16565b805190915015610cc75780806020019051810190614117919061529a565b610cc75760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161156a565b606061418484845f8561418c565b949350505050565b6060824710156141ed5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161156a565b5f5f866001600160a01b031685876040516142089190615248565b5f6040518083038185875af1925050503d805f8114614242576040519150601f19603f3d011682016040523d82523d5f602084013e614247565b606091505b50915091506114e087838387606083156142c15782515f036142ba576001600160a01b0385163b6142ba5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161156a565b5081614184565b61418483838151156142d65781518083602001fd5b8060405162461bcd60e51b815260040161156a919061448c565b61088b806153b683390190565b803563ffffffff811681146140a0575f5ffd5b5f60208284031215614320575f5ffd5b614329826142fd565b9392505050565b80356001600160401b03811681146140a0575f5ffd5b6001600160a01b0381168114612a14575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b031215614372575f5ffd5b61437b896142fd565b975061438960208a01614330565b965061439760408a01614330565b95506143a560608a01614330565b94506080890135935060a0890135925060c08901356143c381614346565b91506103e089018a10156143d5575f5ffd5b60e0890190509295985092959890939650565b5f602082840312156143f8575f5ffd5b5035919050565b5f5f60408385031215614410575f5ffd5b614419836142fd565b915061442760208401614330565b90509250929050565b5f5f60408385031215614441575f5ffd5b82359150602083013561445381614346565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f614329602083018461445e565b634e487b7160e01b5f52602160045260245ffd5b600381106144ce57634e487b7160e01b5f52602160045260245ffd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e0810161450660608301876144b2565b931515608082015260a081019290925260c090910152949350505050565b5f5f83601f840112614534575f5ffd5b5081356001600160401b0381111561454a575f5ffd5b602083019150836020828501011115614561575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b03121561457f575f5ffd5b614588896142fd565b975061459660208a016142fd565b9650604089013595506060890135945060808901356001600160401b038111156145be575f5ffd5b6145ca8b828c01614524565b90955093505060a08901356001600160401b038111156145e8575f5ffd5b6145f48b828c01614524565b999c989b5096995094979396929594505050565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b03861661010083015261467d6101208301866144b2565b61014082019390935261016001529a9950505050505050505050565b81516001600160a01b03168152610180810160208301516146c560208401826001600160401b03169052565b5060408301516146e060408401826001600160a01b03169052565b5060608301516146fb60608401826001600160401b03169052565b506080830151608083015260a083015161472060a08401826001600160401b03169052565b5060c083015161473b60c08401826001600160401b03169052565b5060e083015161475660e08401826001600160401b03169052565b506101008301516147736101008401826001600160401b03169052565b506101208301516147886101208401826144b2565b5061014083015161014083015261016083015161016083015292915050565b5f5f5f5f5f5f60c087890312156147bc575f5ffd5b6147c5876142fd565b95506147d360208801614330565b94506147e160408801614330565b959894975094956060810135955060808101359460a0909101359350915050565b5f60208284031215614812575f5ffd5b61432982614330565b5f5f6040838503121561482c575f5ffd5b823561441981614346565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561487357614873614837565b604052919050565b5f6001600160401b0382111561489357614893614837565b50601f01601f191660200190565b5f6148b36148ae8461487b565b61484b565b90508281528383830111156148c6575f5ffd5b828260208301375f602084830101529392505050565b5f82601f8301126148eb575f5ffd5b614329838335602085016148a1565b5f5f5f6060848603121561490c575f5ffd5b614915846142fd565b925061492360208501614330565b915060408401356001600160401b0381111561493d575f5ffd5b614949868287016148dc565b9150509250925092565b5f5f60408385031215614964575f5ffd5b61496d83614330565b946020939093013593505050565b6001600160a01b0391909116815260200190565b8035600381106140a0575f5ffd5b5f5f5f5f5f5f5f60e0888a0312156149b3575f5ffd5b87356149be81614346565b965060208801356149ce81614346565b95506149dc60408901614330565b94506149ea6060890161498f565b93506080880135925060a08801356001600160401b03811115614a0b575f5ffd5b8801601f81018a13614a1b575f5ffd5b614a2a8a8235602084016148a1565b979a969950949793969295929450505060c09091013590565b5f5f5f60608486031215614a55575f5ffd5b8335614a6081614346565b9250614923602085016142fd565b5f60208284031215614a7e575f5ffd5b813561432981614346565b5f5f5f5f5f5f60a08789031215614a9e575f5ffd5b614aa7876142fd565b955060208701359450604087013593506060870135925060808701356001600160401b03811115614ad6575f5ffd5b614ae289828a01614524565b979a9699509497509295939492505050565b5f5f60408385031215614b05575f5ffd5b8235614b1081614346565b9150614427602084016142fd565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b0386166101008301526001600160401b0385166101208301526001600160401b038416610140830152614bb36101608301846144b2565b9d9c50505050505050505050505050565b5f5f5f60608486031215614bd6575f5ffd5b614a60846142fd565b5f5f5f5f5f5f5f5f610100898b031215614bf7575f5ffd5b8835614c0281614346565b97506020890135614c1281614346565b9650614c2060408a01614330565b9550614c2e60608a01614330565b945060808901359350614c4360a08a0161498f565b979a969950949793969295929450505060c08201359160e0013590565b81516001600160a01b0316815261018081016020830151614c8c60208401826001600160401b03169052565b506040830151614ca760408401826001600160a01b03169052565b506060830151614cc260608401826001600160401b03169052565b506080830151608083015260a0830151614ce760a08401826001600160401b03169052565b5060c0830151614d0260c08401826001600160401b03169052565b5060e0830151614d1d60e08401826001600160401b03169052565b50610100830151614d3a6101008401826001600160401b03169052565b50610120830151614d576101208401826001600160401b03169052565b50610140830151614d746101408401826001600160401b03169052565b50610160830151614d896101608401826144b2565b5092915050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b0381811683821601908111156109f6576109f6614d90565b80820281158282048414176109f6576109f6614d90565b5f60208284031215614dea575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f614e2b604083018661445e565b8281036020840152614e3e818587614df1565b9695505050505050565b848152606060208201525f614e60606083018661445e565b82810360408401526114e0818587614df1565b602081525f614184602083018486614df1565b5f60208284031215614e96575f5ffd5b815161432981614346565b6001600160401b0382811682821603908111156109f6576109f6614d90565b5f63ffffffff821663ffffffff8103614edb57614edb614d90565b60010192915050565b6001600160a01b038481168252831660208201526060604082018190525f90614f0f9083018461445e565b95945050505050565b63ffffffff861681526001600160a01b03851660208201526001600160401b038416604082015260ff8316606082015260a0608082018190525f906114e09083018461445e565b63ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b5f82601f830112614fa6575f5ffd5b8151614fb46148ae8261487b565b818152846020838601011115614fc8575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215614ff8575f5ffd5b855161500381614346565b602087015190955061501481614346565b604087015190945061502581614346565b60608701519093506001600160401b03811115615040575f5ffd5b61504c88828901614f97565b92505060808601516001600160401b03811115615067575f5ffd5b61507388828901614f97565b9150509295509295909350565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c0608082018190525f906150bf9083018561445e565b82810360a08401526150d1818561445e565b9998505050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f82615100576151006150de565b500490565b808201808211156109f6576109f6614d90565b634e487b7160e01b5f52603260045260245ffd5b5f8261513a5761513a6150de565b500690565b818103818111156109f6576109f6614d90565b5f8161516057615160614d90565b505f190190565b6001600160a01b038881168252871660208201526001600160401b038616604082015261519760608201866144b2565b83608082015260e060a08201525f6151b260e083018561445e565b90508260c083015298975050505050505050565b5f602082840312156151d6575f5ffd5b81516001600160f01b031981168114614329575f5ffd5b6001600160401b0388811682526001600160a01b03881660208301528616604082015260e0810161522160608301876144b2565b6001600160401b03851660808301528360a08301528260c083015298975050505050505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b6001811015615290578151835260209283019290910190600101615271565b5050509392505050565b5f602082840312156152aa575f5ffd5b81518015158114614329575f5ffd5b9687526020870195909552604086019390935260e09190911b6001600160e01b03191660608501526064840152608483015260a482015260c40190565b6bffffffffffffffffffffffff198b60601b1681528960148201528860348201526001600160401b0360c01b8860c01b1660548201526001600160401b0360c01b8760c01b16605c8201526001600160401b0360c01b8660c01b16606482015284606c82015283608c8201528260ac82015261538160cc82018360c01b6001600160c01b0319169052565b60d4019a9950505050505050505050565b6001600160a01b03831681526040602082018190525f906141849083018461445e56fe60a060405260405161088b38038061088b83398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f51602061086b5f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f51602061086b5f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516104466104255f395f601001526104465ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610069575f356001600160e01b03191663278f794360e11b146100615761005f61006d565b565b61005f61007d565b61005f5b61005f6100786100ab565b6100e2565b5f8061008c36600481846102ee565b8101906100999190610329565b915091506100a78282610100565b5050565b5f6100dd7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156100fc573d5ff35b3d5ffd5b6101098261015a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101525761014d82826101d5565b505050565b6100a7610247565b806001600160a01b03163b5f0361019457604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516101f191906103fa565b5f60405180830381855af49150503d805f8114610229576040519150601f19603f3d011682016040523d82523d5f602084013e61022e565b606091505b509150915061023e858383610266565b95945050505050565b341561005f5760405163b398979f60e01b815260040160405180910390fd5b60608261027b57610276826102c5565b6102be565b815115801561029257506001600160a01b0384163b155b156102bb57604051639996b31560e01b81526001600160a01b038516600482015260240161018b565b50805b9392505050565b8051156102d55780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f858511156102fc575f5ffd5b83861115610308575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561033a575f5ffd5b82356001600160a01b0381168114610350575f5ffd5b9150602083013567ffffffffffffffff81111561036b575f5ffd5b8301601f8101851361037b575f5ffd5b803567ffffffffffffffff81111561039557610395610315565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156103c4576103c4610315565b6040528181528282016020018710156103db575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220c0931976df4a26c2184205b53cad6447a86a5b19b88f07a03b24865328b06bac64736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610366156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212204420fb0708b8a61d7ab9da438547ada61398f7a9e4ee6dcc588cb57a3ba4422864736f6c634300081c0033",
}

// AgglayermanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayermanagerMetaData.ABI instead.
var AgglayermanagerABI = AgglayermanagerMetaData.ABI

// AgglayermanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayermanagerMetaData.Bin instead.
var AgglayermanagerBin = AgglayermanagerMetaData.Bin

// DeployAgglayermanager deploys a new Ethereum contract, binding an instance of Agglayermanager to it.
func DeployAgglayermanager(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Agglayermanager, error) {
	parsed, err := AgglayermanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayermanagerBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayermanager{AgglayermanagerCaller: AgglayermanagerCaller{contract: contract}, AgglayermanagerTransactor: AgglayermanagerTransactor{contract: contract}, AgglayermanagerFilterer: AgglayermanagerFilterer{contract: contract}}, nil
}

// Agglayermanager is an auto generated Go binding around an Ethereum contract.
type Agglayermanager struct {
	AgglayermanagerCaller     // Read-only binding to the contract
	AgglayermanagerTransactor // Write-only binding to the contract
	AgglayermanagerFilterer   // Log filterer for contract events
}

// AgglayermanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayermanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayermanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayermanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayermanagerSession struct {
	Contract     *Agglayermanager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgglayermanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayermanagerCallerSession struct {
	Contract *AgglayermanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AgglayermanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayermanagerTransactorSession struct {
	Contract     *AgglayermanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AgglayermanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayermanagerRaw struct {
	Contract *Agglayermanager // Generic contract binding to access the raw methods on
}

// AgglayermanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayermanagerCallerRaw struct {
	Contract *AgglayermanagerCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayermanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayermanagerTransactorRaw struct {
	Contract *AgglayermanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayermanager creates a new instance of Agglayermanager, bound to a specific deployed contract.
func NewAgglayermanager(address common.Address, backend bind.ContractBackend) (*Agglayermanager, error) {
	contract, err := bindAgglayermanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayermanager{AgglayermanagerCaller: AgglayermanagerCaller{contract: contract}, AgglayermanagerTransactor: AgglayermanagerTransactor{contract: contract}, AgglayermanagerFilterer: AgglayermanagerFilterer{contract: contract}}, nil
}

// NewAgglayermanagerCaller creates a new read-only instance of Agglayermanager, bound to a specific deployed contract.
func NewAgglayermanagerCaller(address common.Address, caller bind.ContractCaller) (*AgglayermanagerCaller, error) {
	contract, err := bindAgglayermanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerCaller{contract: contract}, nil
}

// NewAgglayermanagerTransactor creates a new write-only instance of Agglayermanager, bound to a specific deployed contract.
func NewAgglayermanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayermanagerTransactor, error) {
	contract, err := bindAgglayermanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerTransactor{contract: contract}, nil
}

// NewAgglayermanagerFilterer creates a new log filterer instance of Agglayermanager, bound to a specific deployed contract.
func NewAgglayermanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayermanagerFilterer, error) {
	contract, err := bindAgglayermanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerFilterer{contract: contract}, nil
}

// bindAgglayermanager binds a generic wrapper to an already deployed contract.
func bindAgglayermanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayermanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanager *AgglayermanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanager.Contract.AgglayermanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanager *AgglayermanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AgglayermanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanager *AgglayermanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AgglayermanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanager *AgglayermanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanager *AgglayermanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanager *AgglayermanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanager.Contract.DEFAULTADMINROLE(&_Agglayermanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanager.Contract.DEFAULTADMINROLE(&_Agglayermanager.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanager *AgglayermanagerCaller) ROLLUPMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "ROLLUP_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanager *AgglayermanagerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanager.Contract.ROLLUPMANAGERVERSION(&_Agglayermanager.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanager *AgglayermanagerCallerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanager.Contract.ROLLUPMANAGERVERSION(&_Agglayermanager.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanager *AgglayermanagerCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanager *AgglayermanagerSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanager.Contract.AggLayerGateway(&_Agglayermanager.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanager *AgglayermanagerCallerSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanager.Contract.AggLayerGateway(&_Agglayermanager.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanager *AgglayermanagerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanager *AgglayermanagerSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanager.Contract.BridgeAddress(&_Agglayermanager.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanager *AgglayermanagerCallerSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanager.Contract.BridgeAddress(&_Agglayermanager.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanager *AgglayermanagerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanager.Contract.CalculateRewardPerBatch(&_Agglayermanager.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanager.Contract.CalculateRewardPerBatch(&_Agglayermanager.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanager.Contract.ChainIDToRollupID(&_Agglayermanager.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanager.Contract.ChainIDToRollupID(&_Agglayermanager.CallOpts, chainID)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanager.Contract.GetBatchFee(&_Agglayermanager.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCallerSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanager.Contract.GetBatchFee(&_Agglayermanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanager.Contract.GetForcedBatchFee(&_Agglayermanager.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanager *AgglayermanagerCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanager.Contract.GetForcedBatchFee(&_Agglayermanager.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanager *AgglayermanagerCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanager *AgglayermanagerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanager.Contract.GetInputPessimisticBytes(&_Agglayermanager.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanager *AgglayermanagerCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanager.Contract.GetInputPessimisticBytes(&_Agglayermanager.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanager *AgglayermanagerCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanager *AgglayermanagerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanager.Contract.GetInputSnarkBytes(&_Agglayermanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanager *AgglayermanagerCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanager.Contract.GetInputSnarkBytes(&_Agglayermanager.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanager *AgglayermanagerCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanager.Contract.GetLastVerifiedBatch(&_Agglayermanager.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanager *AgglayermanagerCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanager.Contract.GetLastVerifiedBatch(&_Agglayermanager.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanager.Contract.GetRoleAdmin(&_Agglayermanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanager.Contract.GetRoleAdmin(&_Agglayermanager.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanager.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanager.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanager.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanager.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanager.Contract.GetRollupExitRoot(&_Agglayermanager.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanager *AgglayermanagerCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanager.Contract.GetRollupExitRoot(&_Agglayermanager.CallOpts)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanager *AgglayermanagerCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanager *AgglayermanagerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanager.Contract.GetRollupSequencedBatches(&_Agglayermanager.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanager *AgglayermanagerCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanager.Contract.GetRollupSequencedBatches(&_Agglayermanager.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanager *AgglayermanagerCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanager *AgglayermanagerSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanager.Contract.GlobalExitRootManager(&_Agglayermanager.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanager *AgglayermanagerCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanager.Contract.GlobalExitRootManager(&_Agglayermanager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanager *AgglayermanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanager *AgglayermanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanager.Contract.HasRole(&_Agglayermanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanager *AgglayermanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanager.Contract.HasRole(&_Agglayermanager.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanager *AgglayermanagerCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanager *AgglayermanagerSession) IsEmergencyState() (bool, error) {
	return _Agglayermanager.Contract.IsEmergencyState(&_Agglayermanager.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanager *AgglayermanagerCallerSession) IsEmergencyState() (bool, error) {
	return _Agglayermanager.Contract.IsEmergencyState(&_Agglayermanager.CallOpts)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanager *AgglayermanagerCaller) IsRollupMigrating(opts *bind.CallOpts, rollupID uint32) (bool, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "isRollupMigrating", rollupID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanager *AgglayermanagerSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanager.Contract.IsRollupMigrating(&_Agglayermanager.CallOpts, rollupID)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanager *AgglayermanagerCallerSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanager.Contract.IsRollupMigrating(&_Agglayermanager.CallOpts, rollupID)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanager.Contract.LastAggregationTimestamp(&_Agglayermanager.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanager.Contract.LastAggregationTimestamp(&_Agglayermanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanager.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanager.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanager.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanager *AgglayermanagerCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanager *AgglayermanagerSession) Pol() (common.Address, error) {
	return _Agglayermanager.Contract.Pol(&_Agglayermanager.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanager *AgglayermanagerCallerSession) Pol() (common.Address, error) {
	return _Agglayermanager.Contract.Pol(&_Agglayermanager.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanager.Contract.RollupAddressToID(&_Agglayermanager.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanager.Contract.RollupAddressToID(&_Agglayermanager.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerSession) RollupCount() (uint32, error) {
	return _Agglayermanager.Contract.RollupCount(&_Agglayermanager.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupCount() (uint32, error) {
	return _Agglayermanager.Contract.RollupCount(&_Agglayermanager.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanager *AgglayermanagerCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturn)).(*AgglayerManagerRollupDataReturn)

	return out0, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanager *AgglayermanagerSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanager.Contract.RollupIDToRollupData(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanager.Contract.RollupIDToRollupData(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanager *AgglayermanagerCaller) RollupIDToRollupDataDeserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanager.contract.Call(opts, &out, "rollupIDToRollupDataDeserialized", rollupID)

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
func (_Agglayermanager *AgglayermanagerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanager.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanager.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanager *AgglayermanagerCaller) RollupIDToRollupDataV2(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupIDToRollupDataV2", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturnV2), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturnV2)).(*AgglayerManagerRollupDataReturnV2)

	return out0, err

}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanager *AgglayermanagerSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanager.Contract.RollupIDToRollupDataV2(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanager.Contract.RollupIDToRollupDataV2(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanager *AgglayermanagerCaller) RollupIDToRollupDataV2Deserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanager.contract.Call(opts, &out, "rollupIDToRollupDataV2Deserialized", rollupID)

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
func (_Agglayermanager *AgglayermanagerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanager.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanager.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanager.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanager.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanager.Contract.RollupTypeCount(&_Agglayermanager.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanager.Contract.RollupTypeCount(&_Agglayermanager.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanager *AgglayermanagerCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Agglayermanager *AgglayermanagerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanager.Contract.RollupTypeMap(&_Agglayermanager.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanager *AgglayermanagerCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanager.Contract.RollupTypeMap(&_Agglayermanager.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanager.Contract.TotalSequencedBatches(&_Agglayermanager.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanager.Contract.TotalSequencedBatches(&_Agglayermanager.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanager.Contract.TotalVerifiedBatches(&_Agglayermanager.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanager *AgglayermanagerCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanager.Contract.TotalVerifiedBatches(&_Agglayermanager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanager *AgglayermanagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanager *AgglayermanagerSession) Version() (string, error) {
	return _Agglayermanager.Contract.Version(&_Agglayermanager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanager *AgglayermanagerCallerSession) Version() (string, error) {
	return _Agglayermanager.Contract.Version(&_Agglayermanager.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanager.Contract.ActivateEmergencyState(&_Agglayermanager.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanager.Contract.ActivateEmergencyState(&_Agglayermanager.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanager *AgglayermanagerTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanager *AgglayermanagerSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AddExistingRollup(&_Agglayermanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AddExistingRollup(&_Agglayermanager.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanager *AgglayermanagerTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanager *AgglayermanagerSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AddNewRollupType(&_Agglayermanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AddNewRollupType(&_Agglayermanager.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanager *AgglayermanagerTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanager *AgglayermanagerSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AttachAggchainToAL(&_Agglayermanager.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.AttachAggchainToAL(&_Agglayermanager.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanager.Contract.DeactivateEmergencyState(&_Agglayermanager.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanager.Contract.DeactivateEmergencyState(&_Agglayermanager.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.GrantRole(&_Agglayermanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.GrantRole(&_Agglayermanager.TransactOpts, role, account)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerTransactor) InitMigration(opts *bind.TransactOpts, rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "initMigration", rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.InitMigration(&_Agglayermanager.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.InitMigration(&_Agglayermanager.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanager *AgglayermanagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanager *AgglayermanagerSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanager.Contract.Initialize(&_Agglayermanager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanager.Contract.Initialize(&_Agglayermanager.TransactOpts)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.Contract.ObsoleteRollupType(&_Agglayermanager.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.Contract.ObsoleteRollupType(&_Agglayermanager.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanager *AgglayermanagerTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanager *AgglayermanagerSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.OnSequenceBatches(&_Agglayermanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanager *AgglayermanagerTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.OnSequenceBatches(&_Agglayermanager.TransactOpts, newSequencedBatches, newAccInputHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RenounceRole(&_Agglayermanager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RenounceRole(&_Agglayermanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RevokeRole(&_Agglayermanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RevokeRole(&_Agglayermanager.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanager *AgglayermanagerTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanager *AgglayermanagerSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RollbackBatches(&_Agglayermanager.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanager.Contract.RollbackBatches(&_Agglayermanager.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanager *AgglayermanagerTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanager *AgglayermanagerSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanager.Contract.SetBatchFee(&_Agglayermanager.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanager.Contract.SetBatchFee(&_Agglayermanager.TransactOpts, newBatchFee)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.UpdateRollup(&_Agglayermanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.UpdateRollup(&_Agglayermanager.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.Contract.UpdateRollupByRollupAdmin(&_Agglayermanager.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanager.Contract.UpdateRollupByRollupAdmin(&_Agglayermanager.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanager *AgglayermanagerTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanager *AgglayermanagerSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanager.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanager *AgglayermanagerTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanager.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanager *AgglayermanagerSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanager *AgglayermanagerTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanager.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanager.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// AgglayermanagerAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Agglayermanager contract.
type AgglayermanagerAddExistingRollupIterator struct {
	Event *AgglayermanagerAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerAddExistingRollup)
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
		it.Event = new(AgglayermanagerAddExistingRollup)
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
func (it *AgglayermanagerAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerAddExistingRollup represents a AddExistingRollup event raised by the Agglayermanager contract.
type AgglayermanagerAddExistingRollup struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerAddExistingRollupIterator{contract: _Agglayermanager.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
func (_Agglayermanager *AgglayermanagerFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagerAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerAddExistingRollup)
				if err := _Agglayermanager.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseAddExistingRollup(log types.Log) (*AgglayermanagerAddExistingRollup, error) {
	event := new(AgglayermanagerAddExistingRollup)
	if err := _Agglayermanager.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Agglayermanager contract.
type AgglayermanagerAddNewRollupTypeIterator struct {
	Event *AgglayermanagerAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerAddNewRollupType)
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
		it.Event = new(AgglayermanagerAddNewRollupType)
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
func (it *AgglayermanagerAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerAddNewRollupType represents a AddNewRollupType event raised by the Agglayermanager contract.
type AgglayermanagerAddNewRollupType struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagerAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerAddNewRollupTypeIterator{contract: _Agglayermanager.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Agglayermanager *AgglayermanagerFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagerAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerAddNewRollupType)
				if err := _Agglayermanager.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseAddNewRollupType(log types.Log) (*AgglayermanagerAddNewRollupType, error) {
	event := new(AgglayermanagerAddNewRollupType)
	if err := _Agglayermanager.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerCompletedMigrationIterator is returned from FilterCompletedMigration and is used to iterate over the raw logs and unpacked data for CompletedMigration events raised by the Agglayermanager contract.
type AgglayermanagerCompletedMigrationIterator struct {
	Event *AgglayermanagerCompletedMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerCompletedMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerCompletedMigration)
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
		it.Event = new(AgglayermanagerCompletedMigration)
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
func (it *AgglayermanagerCompletedMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerCompletedMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerCompletedMigration represents a CompletedMigration event raised by the Agglayermanager contract.
type AgglayermanagerCompletedMigration struct {
	RollupID uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompletedMigration is a free log retrieval operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanager *AgglayermanagerFilterer) FilterCompletedMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerCompletedMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerCompletedMigrationIterator{contract: _Agglayermanager.contract, event: "CompletedMigration", logs: logs, sub: sub}, nil
}

// WatchCompletedMigration is a free log subscription operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanager *AgglayermanagerFilterer) WatchCompletedMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagerCompletedMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerCompletedMigration)
				if err := _Agglayermanager.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseCompletedMigration(log types.Log) (*AgglayermanagerCompletedMigration, error) {
	event := new(AgglayermanagerCompletedMigration)
	if err := _Agglayermanager.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerCreateNewAggchainIterator is returned from FilterCreateNewAggchain and is used to iterate over the raw logs and unpacked data for CreateNewAggchain events raised by the Agglayermanager contract.
type AgglayermanagerCreateNewAggchainIterator struct {
	Event *AgglayermanagerCreateNewAggchain // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerCreateNewAggchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerCreateNewAggchain)
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
		it.Event = new(AgglayermanagerCreateNewAggchain)
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
func (it *AgglayermanagerCreateNewAggchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerCreateNewAggchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerCreateNewAggchain represents a CreateNewAggchain event raised by the Agglayermanager contract.
type AgglayermanagerCreateNewAggchain struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterCreateNewAggchain(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerCreateNewAggchainIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerCreateNewAggchainIterator{contract: _Agglayermanager.contract, event: "CreateNewAggchain", logs: logs, sub: sub}, nil
}

// WatchCreateNewAggchain is a free log subscription operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Agglayermanager *AgglayermanagerFilterer) WatchCreateNewAggchain(opts *bind.WatchOpts, sink chan<- *AgglayermanagerCreateNewAggchain, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerCreateNewAggchain)
				if err := _Agglayermanager.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseCreateNewAggchain(log types.Log) (*AgglayermanagerCreateNewAggchain, error) {
	event := new(AgglayermanagerCreateNewAggchain)
	if err := _Agglayermanager.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Agglayermanager contract.
type AgglayermanagerCreateNewRollupIterator struct {
	Event *AgglayermanagerCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerCreateNewRollup)
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
		it.Event = new(AgglayermanagerCreateNewRollup)
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
func (it *AgglayermanagerCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerCreateNewRollup represents a CreateNewRollup event raised by the Agglayermanager contract.
type AgglayermanagerCreateNewRollup struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerCreateNewRollupIterator{contract: _Agglayermanager.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Agglayermanager *AgglayermanagerFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagerCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerCreateNewRollup)
				if err := _Agglayermanager.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseCreateNewRollup(log types.Log) (*AgglayermanagerCreateNewRollup, error) {
	event := new(AgglayermanagerCreateNewRollup)
	if err := _Agglayermanager.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Agglayermanager contract.
type AgglayermanagerEmergencyStateActivatedIterator struct {
	Event *AgglayermanagerEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerEmergencyStateActivated)
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
		it.Event = new(AgglayermanagerEmergencyStateActivated)
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
func (it *AgglayermanagerEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerEmergencyStateActivated represents a EmergencyStateActivated event raised by the Agglayermanager contract.
type AgglayermanagerEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanager *AgglayermanagerFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*AgglayermanagerEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerEmergencyStateActivatedIterator{contract: _Agglayermanager.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanager *AgglayermanagerFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagerEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerEmergencyStateActivated)
				if err := _Agglayermanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseEmergencyStateActivated(log types.Log) (*AgglayermanagerEmergencyStateActivated, error) {
	event := new(AgglayermanagerEmergencyStateActivated)
	if err := _Agglayermanager.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Agglayermanager contract.
type AgglayermanagerEmergencyStateDeactivatedIterator struct {
	Event *AgglayermanagerEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerEmergencyStateDeactivated)
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
		it.Event = new(AgglayermanagerEmergencyStateDeactivated)
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
func (it *AgglayermanagerEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Agglayermanager contract.
type AgglayermanagerEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanager *AgglayermanagerFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*AgglayermanagerEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerEmergencyStateDeactivatedIterator{contract: _Agglayermanager.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanager *AgglayermanagerFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagerEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerEmergencyStateDeactivated)
				if err := _Agglayermanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseEmergencyStateDeactivated(log types.Log) (*AgglayermanagerEmergencyStateDeactivated, error) {
	event := new(AgglayermanagerEmergencyStateDeactivated)
	if err := _Agglayermanager.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerInitMigrationIterator is returned from FilterInitMigration and is used to iterate over the raw logs and unpacked data for InitMigration events raised by the Agglayermanager contract.
type AgglayermanagerInitMigrationIterator struct {
	Event *AgglayermanagerInitMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerInitMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerInitMigration)
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
		it.Event = new(AgglayermanagerInitMigration)
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
func (it *AgglayermanagerInitMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerInitMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerInitMigration represents a InitMigration event raised by the Agglayermanager contract.
type AgglayermanagerInitMigration struct {
	RollupID        uint32
	NewRollupTypeID uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitMigration is a free log retrieval operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanager *AgglayermanagerFilterer) FilterInitMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerInitMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerInitMigrationIterator{contract: _Agglayermanager.contract, event: "InitMigration", logs: logs, sub: sub}, nil
}

// WatchInitMigration is a free log subscription operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanager *AgglayermanagerFilterer) WatchInitMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagerInitMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerInitMigration)
				if err := _Agglayermanager.contract.UnpackLog(event, "InitMigration", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseInitMigration(log types.Log) (*AgglayermanagerInitMigration, error) {
	event := new(AgglayermanagerInitMigration)
	if err := _Agglayermanager.contract.UnpackLog(event, "InitMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayermanager contract.
type AgglayermanagerInitializedIterator struct {
	Event *AgglayermanagerInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerInitialized)
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
		it.Event = new(AgglayermanagerInitialized)
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
func (it *AgglayermanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerInitialized represents a Initialized event raised by the Agglayermanager contract.
type AgglayermanagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanager *AgglayermanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayermanagerInitializedIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerInitializedIterator{contract: _Agglayermanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanager *AgglayermanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayermanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerInitialized)
				if err := _Agglayermanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseInitialized(log types.Log) (*AgglayermanagerInitialized, error) {
	event := new(AgglayermanagerInitialized)
	if err := _Agglayermanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Agglayermanager contract.
type AgglayermanagerObsoleteRollupTypeIterator struct {
	Event *AgglayermanagerObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerObsoleteRollupType)
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
		it.Event = new(AgglayermanagerObsoleteRollupType)
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
func (it *AgglayermanagerObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerObsoleteRollupType represents a ObsoleteRollupType event raised by the Agglayermanager contract.
type AgglayermanagerObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanager *AgglayermanagerFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagerObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerObsoleteRollupTypeIterator{contract: _Agglayermanager.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanager *AgglayermanagerFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagerObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerObsoleteRollupType)
				if err := _Agglayermanager.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseObsoleteRollupType(log types.Log) (*AgglayermanagerObsoleteRollupType, error) {
	event := new(AgglayermanagerObsoleteRollupType)
	if err := _Agglayermanager.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Agglayermanager contract.
type AgglayermanagerOnSequenceBatchesIterator struct {
	Event *AgglayermanagerOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerOnSequenceBatches)
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
		it.Event = new(AgglayermanagerOnSequenceBatches)
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
func (it *AgglayermanagerOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerOnSequenceBatches represents a OnSequenceBatches event raised by the Agglayermanager contract.
type AgglayermanagerOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanager *AgglayermanagerFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerOnSequenceBatchesIterator{contract: _Agglayermanager.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanager *AgglayermanagerFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagerOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerOnSequenceBatches)
				if err := _Agglayermanager.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseOnSequenceBatches(log types.Log) (*AgglayermanagerOnSequenceBatches, error) {
	event := new(AgglayermanagerOnSequenceBatches)
	if err := _Agglayermanager.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Agglayermanager contract.
type AgglayermanagerRoleAdminChangedIterator struct {
	Event *AgglayermanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerRoleAdminChanged)
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
		it.Event = new(AgglayermanagerRoleAdminChanged)
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
func (it *AgglayermanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Agglayermanager contract.
type AgglayermanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanager *AgglayermanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgglayermanagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerRoleAdminChangedIterator{contract: _Agglayermanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanager *AgglayermanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgglayermanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerRoleAdminChanged)
				if err := _Agglayermanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseRoleAdminChanged(log types.Log) (*AgglayermanagerRoleAdminChanged, error) {
	event := new(AgglayermanagerRoleAdminChanged)
	if err := _Agglayermanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Agglayermanager contract.
type AgglayermanagerRoleGrantedIterator struct {
	Event *AgglayermanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerRoleGranted)
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
		it.Event = new(AgglayermanagerRoleGranted)
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
func (it *AgglayermanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerRoleGranted represents a RoleGranted event raised by the Agglayermanager contract.
type AgglayermanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanager *AgglayermanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagerRoleGrantedIterator, error) {

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

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerRoleGrantedIterator{contract: _Agglayermanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanager *AgglayermanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgglayermanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerRoleGranted)
				if err := _Agglayermanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseRoleGranted(log types.Log) (*AgglayermanagerRoleGranted, error) {
	event := new(AgglayermanagerRoleGranted)
	if err := _Agglayermanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Agglayermanager contract.
type AgglayermanagerRoleRevokedIterator struct {
	Event *AgglayermanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerRoleRevoked)
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
		it.Event = new(AgglayermanagerRoleRevoked)
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
func (it *AgglayermanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerRoleRevoked represents a RoleRevoked event raised by the Agglayermanager contract.
type AgglayermanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanager *AgglayermanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagerRoleRevokedIterator, error) {

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

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerRoleRevokedIterator{contract: _Agglayermanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanager *AgglayermanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgglayermanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerRoleRevoked)
				if err := _Agglayermanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseRoleRevoked(log types.Log) (*AgglayermanagerRoleRevoked, error) {
	event := new(AgglayermanagerRoleRevoked)
	if err := _Agglayermanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Agglayermanager contract.
type AgglayermanagerRollbackBatchesIterator struct {
	Event *AgglayermanagerRollbackBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerRollbackBatches)
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
		it.Event = new(AgglayermanagerRollbackBatches)
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
func (it *AgglayermanagerRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerRollbackBatches represents a RollbackBatches event raised by the Agglayermanager contract.
type AgglayermanagerRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanager *AgglayermanagerFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*AgglayermanagerRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerRollbackBatchesIterator{contract: _Agglayermanager.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanager *AgglayermanagerFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagerRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerRollbackBatches)
				if err := _Agglayermanager.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseRollbackBatches(log types.Log) (*AgglayermanagerRollbackBatches, error) {
	event := new(AgglayermanagerRollbackBatches)
	if err := _Agglayermanager.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Agglayermanager contract.
type AgglayermanagerSetBatchFeeIterator struct {
	Event *AgglayermanagerSetBatchFee // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerSetBatchFee)
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
		it.Event = new(AgglayermanagerSetBatchFee)
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
func (it *AgglayermanagerSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerSetBatchFee represents a SetBatchFee event raised by the Agglayermanager contract.
type AgglayermanagerSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanager *AgglayermanagerFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*AgglayermanagerSetBatchFeeIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerSetBatchFeeIterator{contract: _Agglayermanager.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanager *AgglayermanagerFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *AgglayermanagerSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerSetBatchFee)
				if err := _Agglayermanager.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseSetBatchFee(log types.Log) (*AgglayermanagerSetBatchFee, error) {
	event := new(AgglayermanagerSetBatchFee)
	if err := _Agglayermanager.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Agglayermanager contract.
type AgglayermanagerSetTrustedAggregatorIterator struct {
	Event *AgglayermanagerSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerSetTrustedAggregator)
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
		it.Event = new(AgglayermanagerSetTrustedAggregator)
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
func (it *AgglayermanagerSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerSetTrustedAggregator represents a SetTrustedAggregator event raised by the Agglayermanager contract.
type AgglayermanagerSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanager *AgglayermanagerFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*AgglayermanagerSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerSetTrustedAggregatorIterator{contract: _Agglayermanager.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanager *AgglayermanagerFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagerSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerSetTrustedAggregator)
				if err := _Agglayermanager.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseSetTrustedAggregator(log types.Log) (*AgglayermanagerSetTrustedAggregator, error) {
	event := new(AgglayermanagerSetTrustedAggregator)
	if err := _Agglayermanager.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Agglayermanager contract.
type AgglayermanagerUpdateRollupIterator struct {
	Event *AgglayermanagerUpdateRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerUpdateRollup)
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
		it.Event = new(AgglayermanagerUpdateRollup)
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
func (it *AgglayermanagerUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerUpdateRollup represents a UpdateRollup event raised by the Agglayermanager contract.
type AgglayermanagerUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanager *AgglayermanagerFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagerUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerUpdateRollupIterator{contract: _Agglayermanager.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanager *AgglayermanagerFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagerUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerUpdateRollup)
				if err := _Agglayermanager.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseUpdateRollup(log types.Log) (*AgglayermanagerUpdateRollup, error) {
	event := new(AgglayermanagerUpdateRollup)
	if err := _Agglayermanager.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerUpdateRollupManagerVersionIterator is returned from FilterUpdateRollupManagerVersion and is used to iterate over the raw logs and unpacked data for UpdateRollupManagerVersion events raised by the Agglayermanager contract.
type AgglayermanagerUpdateRollupManagerVersionIterator struct {
	Event *AgglayermanagerUpdateRollupManagerVersion // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerUpdateRollupManagerVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerUpdateRollupManagerVersion)
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
		it.Event = new(AgglayermanagerUpdateRollupManagerVersion)
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
func (it *AgglayermanagerUpdateRollupManagerVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerUpdateRollupManagerVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerUpdateRollupManagerVersion represents a UpdateRollupManagerVersion event raised by the Agglayermanager contract.
type AgglayermanagerUpdateRollupManagerVersion struct {
	RollupManagerVersion string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupManagerVersion is a free log retrieval operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanager *AgglayermanagerFilterer) FilterUpdateRollupManagerVersion(opts *bind.FilterOpts) (*AgglayermanagerUpdateRollupManagerVersionIterator, error) {

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerUpdateRollupManagerVersionIterator{contract: _Agglayermanager.contract, event: "UpdateRollupManagerVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupManagerVersion is a free log subscription operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanager *AgglayermanagerFilterer) WatchUpdateRollupManagerVersion(opts *bind.WatchOpts, sink chan<- *AgglayermanagerUpdateRollupManagerVersion) (event.Subscription, error) {

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerUpdateRollupManagerVersion)
				if err := _Agglayermanager.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseUpdateRollupManagerVersion(log types.Log) (*AgglayermanagerUpdateRollupManagerVersion, error) {
	event := new(AgglayermanagerUpdateRollupManagerVersion)
	if err := _Agglayermanager.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Agglayermanager contract.
type AgglayermanagerVerifyBatchesTrustedAggregatorIterator struct {
	Event *AgglayermanagerVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerVerifyBatchesTrustedAggregator)
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
		it.Event = new(AgglayermanagerVerifyBatchesTrustedAggregator)
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
func (it *AgglayermanagerVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Agglayermanager contract.
type AgglayermanagerVerifyBatchesTrustedAggregator struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*AgglayermanagerVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerVerifyBatchesTrustedAggregatorIterator{contract: _Agglayermanager.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Agglayermanager *AgglayermanagerFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagerVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerVerifyBatchesTrustedAggregator)
				if err := _Agglayermanager.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*AgglayermanagerVerifyBatchesTrustedAggregator, error) {
	event := new(AgglayermanagerVerifyBatchesTrustedAggregator)
	if err := _Agglayermanager.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagerVerifyPessimisticStateTransitionIterator is returned from FilterVerifyPessimisticStateTransition and is used to iterate over the raw logs and unpacked data for VerifyPessimisticStateTransition events raised by the Agglayermanager contract.
type AgglayermanagerVerifyPessimisticStateTransitionIterator struct {
	Event *AgglayermanagerVerifyPessimisticStateTransition // Event containing the contract specifics and raw log

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
func (it *AgglayermanagerVerifyPessimisticStateTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagerVerifyPessimisticStateTransition)
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
		it.Event = new(AgglayermanagerVerifyPessimisticStateTransition)
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
func (it *AgglayermanagerVerifyPessimisticStateTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagerVerifyPessimisticStateTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagerVerifyPessimisticStateTransition represents a VerifyPessimisticStateTransition event raised by the Agglayermanager contract.
type AgglayermanagerVerifyPessimisticStateTransition struct {
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
func (_Agglayermanager *AgglayermanagerFilterer) FilterVerifyPessimisticStateTransition(opts *bind.FilterOpts, rollupID []uint32, trustedAggregator []common.Address) (*AgglayermanagerVerifyPessimisticStateTransitionIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanager.contract.FilterLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagerVerifyPessimisticStateTransitionIterator{contract: _Agglayermanager.contract, event: "VerifyPessimisticStateTransition", logs: logs, sub: sub}, nil
}

// WatchVerifyPessimisticStateTransition is a free log subscription operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Agglayermanager *AgglayermanagerFilterer) WatchVerifyPessimisticStateTransition(opts *bind.WatchOpts, sink chan<- *AgglayermanagerVerifyPessimisticStateTransition, rollupID []uint32, trustedAggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanager.contract.WatchLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagerVerifyPessimisticStateTransition)
				if err := _Agglayermanager.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
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
func (_Agglayermanager *AgglayermanagerFilterer) ParseVerifyPessimisticStateTransition(log types.Log) (*AgglayermanagerVerifyPessimisticStateTransition, error) {
	event := new(AgglayermanagerVerifyPessimisticStateTransition)
	if err := _Agglayermanager.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

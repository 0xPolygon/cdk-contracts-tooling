// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agglayermanagermock

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

// AgglayermanagermockMetaData contains all meta data concerning the Agglayermanagermock contract.
var AgglayermanagermockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlOnlyCanRenounceRolesForSelf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressDoNotHaveRequiredRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainDataMustBeZeroForPessimisticVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllSequencedMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllzkEVMSequencedBatchesMustBeVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchFeeOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotUpdateWithUnconsolidatedPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIDOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyVerifySequencesData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceBelowLastVerifiedSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumSequenceDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitBatchMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceMustMatchCurrentForkID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequenceNumDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidConstructorInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputsForRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewLocalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPessimisticProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierZkGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeSequenceTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRollupType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVerifierType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustSequenceSomeBlob\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewRollupTypeMustBePessimisticOrALGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStateTransitionChains\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNumExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotEndOfSequence\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollbackBatchIsNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAddressAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupIDNotAscendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupMustExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupTypeObsolete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMustBeRollup\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StateTransitionChainsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateNotCompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToOldRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateToSameRollupTypeID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"zkGasPriceOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"AddExistingRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"AddNewRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"CompletedMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"CreateNewAggchain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasTokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"InitMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"ObsoleteRollupType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"}],\"name\":\"OnSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"SetBatchFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedAggregator\",\"type\":\"address\"}],\"name\":\"SetTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"}],\"name\":\"UpdateRollup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupManagerVersion\",\"type\":\"string\"}],\"name\":\"UpdateRollupManagerVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatchesTrustedAggregator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"prevLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"}],\"name\":\"VerifyPessimisticStateTransition\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROLLUP_MANAGER_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"initRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"initPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"addExistingRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"name\":\"addNewRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAgglayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"attachAggchainToAL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateRewardPerBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"chainIDToRollupID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStateRoot\",\"type\":\"uint256\"}],\"name\":\"exposed_checkStateRootInsidePrime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getForcedBatchFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"l1InfoTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getInputPessimisticBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"oldStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"}],\"name\":\"getInputSnarkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"getLastVerifiedBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupBatchNumToStateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"batchNum\",\"type\":\"uint64\"}],\"name\":\"getRollupSequencedBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"accInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequencedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"previousLastBatchSequenced\",\"type\":\"uint64\"}],\"internalType\":\"structLegacyZKEVMStateVariables.SequencedBatchData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"initMigration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trustedAggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emergencyCouncil\",\"type\":\"address\"}],\"name\":\"initializeMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"isRollupMigrating\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAggregationTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDeactivatedEmergencyStateTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"obsoleteRollupType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSequencedBatches\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"onSequenceBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"localExitRoots\",\"type\":\"bytes32[]\"}],\"name\":\"prepareMockCalculateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolygonRollupBase\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"}],\"name\":\"rollupAddressToID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"internalType\":\"structAgglayerManager.RollupDataReturn\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataDeserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingState\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"legacyLastPendingStateConsolidated\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"internalType\":\"structAgglayerManager.RollupDataReturnV2\",\"name\":\"rollupData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"}],\"name\":\"rollupIDToRollupDataV2Deserialized\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"lastBatchSequenced\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatchBeforeUpgrade\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"rollupTypeID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTypeCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupTypeID\",\"type\":\"uint32\"}],\"name\":\"rollupTypeMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"forkID\",\"type\":\"uint64\"},{\"internalType\":\"enumIAgglayerManager.VerifierType\",\"name\":\"rollupVerifierType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"obsolete\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"genesis\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"programVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBatchFee\",\"type\":\"uint256\"}],\"name\":\"setBatchFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"lastLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastPessimisticRoot\",\"type\":\"bytes32\"}],\"name\":\"setRollupData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSequencedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVerifiedBatches\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"upgradeData\",\"type\":\"bytes\"}],\"name\":\"updateRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITransparentUpgradeableProxy\",\"name\":\"rollupContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newRollupTypeID\",\"type\":\"uint32\"}],\"name\":\"updateRollupByRollupAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"pendingStateNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initNumBatch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"finalNewBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes32[24]\",\"name\":\"proof\",\"type\":\"bytes32[24]\"}],\"name\":\"verifyBatchesTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rollupID\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"l1InfoTreeLeafCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"newLocalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newPessimisticRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"verifyPessimisticTrustedAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161655c38038061655c83398101604081905261002f9161019a565b838383836001600160a01b038416158061005057506001600160a01b038316155b8061006257506001600160a01b038216155b8061007457506001600160a01b038116155b15610092576040516342eda64b60e11b815260040160405180910390fd5b6001600160a01b0380851660805283811660c05282811660a052811660e0526100b96100c6565b50505050505050506101f6565b5f54610100900460ff16156101315760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610181575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610197575f5ffd5b50565b5f5f5f5f608085870312156101ad575f5ffd5b84516101b881610183565b60208601519094506101c981610183565b60408601519093506101da81610183565b60608601519092506101eb81610183565b939692955090935050565b60805160a05160c05160e0516162fb6102615f395f81816107f5015261131f01525f818161093b015281816122b90152613af001525f81816107c101528181612d7a0152613bee01525f818161088101528181610b7b015281816111b2015261142c01526162fb5ff3fe608060405234801561000f575f5ffd5b50600436106102e4575f3560e01c80638f698ec51161018d578063d02103ca116100e0578063dfdb8c5e1161008f578063dfdb8c5e14610923578063e46761c414610936578063e4f3d8f91461095d578063e764a37314610a0d578063e80e503014610a20578063f4e9267514610a33578063f9c4c2ae14610a43575f5ffd5b8063d02103ca1461087c578063d5073f6f146108a3578063d547741f146108b6578063d8905812146108c9578063dbc16976146108ee578063dd0464b9146108f6578063dde0ff7714610909575f5ffd5b8063a2967d991161013c578063a2967d99146107b4578063a3c573eb146107bc578063ab0475cf146107f0578063abcb519814610817578063c1acbc341461082a578063c4c928c214610844578063ceee281d14610857575f5ffd5b80638f698ec5146107465780638fd88cc21461075957806391d148541461076c57806397d289a31461077f57806399f5634e146107925780639a908e731461079a578063a217fddf146107ad575f5ffd5b806354fd4d50116102455780636c766877116101f45780636c7668771461061257806370603909146106255780637222020f146106d357806374d9c244146106e65780637975fcfe146107065780637fb6e76a146107195780638129fc1c1461073e575f5ffd5b806354fd4d50146104cd57806355a71ee0146104f8578063604691691461053857806362d87e661461054057806365c0504d14610553578063680658a1146105ce5780636c0a51b7146105ff575f5ffd5b8063248a9ca3116102a1578063248a9ca31461038a57806325280169146103ba5780632f2ff15d1461046a57806330c27dde1461047d57806336568abe146104905780633a7094bd146104a3578063477fa270146104c5575f5ffd5b8063066ec012146102e857806311f6b287146103185780631489ed101461032b57806315064c96146103405780631796a1ae1461035d5780632072f6c514610382575b5f5ffd5b6084546102fb906001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b6102fb61032636600461479f565b610a63565b61033e6103393660046147e9565b610a92565b005b606f5461034d9060ff1681565b604051901515815260200161030f565b607e5461036d9063ffffffff1681565b60405163ffffffff909116815260200161030f565b61033e610c5b565b6103ac610398366004614877565b5f9081526034602052604090206001015490565b60405190815260200161030f565b6104376103c836600461488e565b60408051606080820183525f808352602080840182905292840181905263ffffffff959095168552608182528285206001600160401b03948516865260030182529382902082519485018352805485526001015480841691850191909152600160401b90049091169082015290565b60408051825181526020808401516001600160401b0390811691830191909152928201519092169082015260600161030f565b61033e6104783660046148bf565b610d13565b6087546102fb906001600160401b031681565b61033e61049e3660046148bf565b610d3c565b61034d6104b136600461479f565b60886020525f908152604090205460ff1681565b6086546103ac565b604080518082019091526006815265076312e302e360d41b60208201525b60405161030f919061491b565b6103ac61050636600461488e565b63ffffffff82165f9081526081602090815260408083206001600160401b038516845260020190915290205492915050565b6103ac610d73565b61034d61054e366004614877565b610d88565b6105bb61056136600461479f565b607f6020525f908152604090208054600182015460028301546003909301546001600160a01b0392831693928216926001600160401b03600160a01b8404169260ff600160e01b8204811693600160e81b90920416919087565b60405161030f9796959493929190614961565b61033e6105dc3660046149b3565b63ffffffff9092165f908152608160205260409020600581019190915560080155565b61033e61060d3660046149e3565b610d92565b61033e610620366004614a80565b6110d7565b6106bb61063336600461479f565b63ffffffff165f9081526081602052604090208054600182015460058301546006840154600785015460088601546009909601546001600160a01b0380871698600160a01b978890046001600160401b03908116999288169890970487169680861695600160401b908190048216958281169591810490921693600160801b90920460ff1692565b60405161030f9c9b9a99989796959493929190614b20565b61033e6106e136600461479f565b611627565b6106f96106f436600461479f565b611718565b60405161030f9190614bb1565b6104eb610714366004614cbf565b611867565b61036d610727366004614d1a565b60836020525f908152604090205463ffffffff1681565b61033e611897565b61033e610754366004614d77565b61197a565b61033e610767366004614e16565b6119eb565b61034d61077a3660046148bf565b611d59565b61033e61078d366004614eb1565b611d83565b6103ac6122b5565b6102fb6107a8366004614f0a565b612391565b6103ac5f81565b6103ac61257e565b6107e37f000000000000000000000000000000000000000000000000000000000000000081565b60405161030f9190614f32565b6107e37f000000000000000000000000000000000000000000000000000000000000000081565b61033e610825366004614f54565b612844565b6084546102fb90600160801b90046001600160401b031681565b61033e610852366004614ffa565b612b58565b61036d610865366004615025565b60826020525f908152604090205463ffffffff1681565b6107e37f000000000000000000000000000000000000000000000000000000000000000081565b61033e6108b1366004614877565b612c60565b61033e6108c43660046148bf565b612cfe565b6104eb60405180604001604052806006815260200165076312e302e360d41b81525081565b61033e612d22565b6104eb610904366004615040565b612de8565b6084546102fb90600160401b90046001600160401b031681565b61033e6109313660046150ab565b612e0f565b6107e37f000000000000000000000000000000000000000000000000000000000000000081565b6109f561096b36600461479f565b63ffffffff165f90815260816020526040902080546001820154600583015460068401546007909401546001600160a01b0380851696600160a01b958690046001600160401b03908116979286169690950485169480831693600160401b808504831694600160801b808204851695600160c01b9092048516948085169493840416920460ff1690565b60405161030f9c9b9a999897969594939291906150d5565b61033e610a1b36600461517b565b613193565b61033e610a2e366004615196565b6132ff565b60805461036d9063ffffffff1681565b610a56610a5136600461479f565b61373e565b60405161030f9190615217565b63ffffffff81165f90815260816020526040812060060154600160401b90046001600160401b03165b92915050565b5f5160206162a65f395f51905f52610aa981613894565b6001600160401b03881615610ad15760405163306dfc5760e11b815260040160405180910390fd5b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff166002811115610b0657610b0661492d565b14610b24576040516390db0d0760e01b815260040160405180910390fd5b610b338189898989898961389e565b60068101805467ffffffffffffffff60401b1916600160401b6001600160401b038a16908102919091179091555f9081526002820160205260409020859055600581018690557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d610bb061257e565b6040518263ffffffff1660e01b8152600401610bce91815260200190565b5f604051808303815f87803b158015610be5575f5ffd5b505af1158015610bf7573d5f5f3e3d5ffd5b5050604080516001600160401b038b1681526020810189905290810189905233925063ffffffff8d1691507fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d39060600160405180910390a350505050505050505050565b610c725f5160206162865f395f51905f5233611d59565b610d0957608454600160801b90046001600160401b03161580610cbe57506084544290610cb39062093a8090600160801b90046001600160401b031661535b565b6001600160401b0316115b80610ceb57506087544290610ce09062093a80906001600160401b031661535b565b6001600160401b0316115b15610d095760405163692baaad60e11b815260040160405180910390fd5b610d11613bec565b565b5f82815260346020526040902060010154610d2d81613894565b610d378383613c62565b505050565b6001600160a01b0381163314610d6557604051630b4ad1cd60e31b815260040160405180910390fd5b610d6f8282613cca565b5050565b5f6086546064610d83919061537a565b905090565b5f610a8c82613d30565b5f54600490610100900460ff16158015610db257505f5460ff8083169116105b610dd75760405162461bcd60e51b8152600401610dce90615391565b60405180910390fd5b5f805461ffff191660ff83161761010017905567016345785d8a0000608655610dfe613db5565b610e155f5160206162a65f395f51905f5286613e1f565b610e1f5f84613e1f565b610e497fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59084613e1f565b610e737f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e84613e1f565b610e8a5f5160206162465f395f51905f5284613e1f565b610eb47fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd85613e1f565b610ede7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd0885613e1f565b610f087f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f485613e1f565b610f327fa5c5790f581d443ed43873ab47cfb8c5d66a6db268e58b5971bb33fc66e07db185613e1f565b610f695f5160206162a65f395f51905f527f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f0613e29565b610f937f73cb0569fdbea2544dae03fdb2fe10eda92a72a2e8cd2bd496e85b762505a3f085613e1f565b610fbd7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb85613e1f565b610ff45f5160206162865f395f51905f527f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff285951613e29565b61100b5f5160206162865f395f51905f5283613e1f565b6110357f9b6f082d8d3644ae2f24a3c32e356d6f2d9b2844d9b26164fbc82663ff28595183613e1f565b61103f5f33613e1f565b6040805180820182526006815265076312e302e360d41b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2916110889161491b565b60405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050565b5f5160206162a65f395f51905f526110ee81613894565b6110f6613e73565b63ffffffff89165f908152608160205260408120906007820154600160801b900460ff16600281111561112b5761112b61492d565b0361114957604051635b6602b760e01b815260040160405180910390fd5b60016007820154600160801b900460ff16600281111561116b5761116b61492d565b14801561117757508215155b1561119557604051630ded782d60e01b815260040160405180910390fd5b60405163ef4eeb3560e01b815263ffffffff8a1660048201525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063ef4eeb3590602401602060405180830381865afa1580156111ff573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061122391906153df565b9050806112435760405163a60721e160e01b815260040160405180910390fd5b63ffffffff8b165f9081526088602052604090205460ff16156112cf5781600501548914611284576040516306c6486360e11b815260040160405180910390fd5b5f6005830181905563ffffffff8c1680825260886020526040808320805460ff191690555190917f6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e91a25b5f6112df8c84848d8d8b8b613eba565b905060026007840154600160801b900460ff1660028111156113035761130361492d565b036113895760405163a48fd37760e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a48fd377906113589084908c908c9060040161541e565b5f6040518083038186803b15801561136e575f5ffd5b505afa158015611380573d5f5f3e3d5ffd5b505050506113f1565b6001830154600984015460405163020a49e360e51b81526001600160a01b03909216916341493c60916113c49185908d908d9060040161544d565b5f6040518083038186803b1580156113da575f5ffd5b505afa1580156113ec573d5f5f3e3d5ffd5b505050505b6084805467ffffffffffffffff60801b1916600160801b426001600160401b031602179055600583018054908b9055600884018054908b90557f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166333d6247d61146161257e565b6040518263ffffffff1660e01b815260040161147f91815260200190565b5f604051808303815f87803b158015611496575f5ffd5b505af11580156114a8573d5f5f3e3d5ffd5b505050505f8c9050336001600160a01b03168f63ffffffff167fd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d35f5f5f1b85604051611512939291906001600160401b039390931683526020830191909152604082015260600190565b60405180910390a3336001600160a01b03168f63ffffffff167fdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711848f87868b604051611580959493929190948552602085019390935260408401919091526060830152608082015260a00190565b60405180910390a360026007870154600160801b900460ff1660028111156115aa576115aa61492d565b0361160e578554604051639ee4afa360e01b81526001600160a01b0390911690639ee4afa3906115e0908c908c90600401615478565b5f604051808303815f87803b1580156115f7575f5ffd5b505af1158015611609573d5f5f3e3d5ffd5b505050505b50505050505061161c61403d565b505050505050505050565b7fab66e11c4f712cd06ab11bf9339b48bef39e12d4a22eeef71d2860a0c90482bd61165181613894565b63ffffffff8216158061166f5750607e5463ffffffff908116908316115b1561168d57604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff82165f908152607f602052604090206001810154600160e81b900460ff16156116ce57604051633b8d3d9960e01b815260040160405180910390fd5b60018101805460ff60e81b1916600160e81b17905560405163ffffffff8416907f4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44905f90a2505050565b61177b60408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082019081526020015f81526020015f81525090565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b0380821686526001600160401b03600160a01b928390048116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b90819004821660c0850152600783015480831660e086015290810490911661010084015261012083019060ff600160801b9091041660028111156118375761183761492d565b9081600281111561184a5761184a61492d565b905250600881015461014083015260090154610160820152919050565b63ffffffff86165f90815260816020526040902060609061188c908787878787614054565b979650505050505050565b5f54600590610100900460ff161580156118b757505f5460ff8083169116105b6118d35760405162461bcd60e51b8152600401610dce90615391565b5f805461ffff191660ff8316176101001790556040805180820182526006815265076312e302e360d41b602082015290517f50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c29161192f9161491b565b60405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b80516080805463ffffffff191663ffffffff9092169190911790555f5b8151811015610d6f578181815181106119b2576119b261548b565b602002602001015160815f8360016119ca919061549f565b63ffffffff16815260208101919091526040015f2060050155600101611997565b6119f3613e73565b611a0a5f5160206162465f395f51905f5233611d59565b158015611a885750336001600160a01b0316826001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a58573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a7c91906154b2565b6001600160a01b031614155b15611aa657604051630d03687f60e11b815260040160405180910390fd5b6001600160a01b0382165f9081526082602052604081205463ffffffff1690819003611ae5576040516374a086a360e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff166002811115611b1a57611b1a61492d565b14611b38576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b0390811690841681111580611b70575060068201546001600160401b03600160401b9091048116908516105b15611b8e5760405163cb23ebdf60e01b815260040160405180910390fd5b805b846001600160401b0316816001600160401b031614611c26576001600160401b038082165f908152600385016020526040902060010154600160401b90048116908616811015611bf357604051639753965f60e01b815260040160405180910390fd5b6001600160401b039091165f908152600384016020526040812090815560010180546001600160801b0319169055611b90565b60068301805467ffffffffffffffff19166001600160401b038716179055611c4e85836154cd565b608480545f90611c689084906001600160401b03166154cd565b82546101009290920a6001600160401b0381810219909316918316021790915586165f8181526003860160205260409081902054905163334d6f6760e11b8152600481019290925260248201526001600160a01b038816915063669adece906044015f604051808303815f87803b158015611ce1575f5ffd5b505af1158015611cf3573d5f5f3e3d5ffd5b505050506001600160401b0385165f81815260038501602090815260409182902054915191825263ffffffff8716917f80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402910160405180910390a350505050610d6f61403d565b5f9182526034602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7fa0fab074aba36a6fa69f1a83ee86e5abfb8433966eb57efb13dc2fc2f24ddd08611dad81613894565b63ffffffff84161580611dcb5750607e5463ffffffff908116908516115b15611de957604051637512e5cb60e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff1615611e2a57604051633b8d3d9960e01b815260040160405180910390fd5b63ffffffff6001600160401b0385161115611e5857604051634c753f5760e01b815260040160405180910390fd5b6001600160401b0384165f9081526083602052604090205463ffffffff1615611e94576040516337c8fe0960e11b815260040160405180910390fd5b608080545f91908290611eac9063ffffffff166154ec565b825463ffffffff8281166101009490940a93840293021916919091179091558254604080515f80825260208201928390529394506001600160a01b03909216913091611ef79061477f565b611f0393929190615510565b604051809103905ff080158015611f1c573d5f5f3e3d5ffd5b506001600160401b0387165f818152608360209081526040808320805463ffffffff1990811663ffffffff8a81169182179093556001600160a01b038816808752608286528487208054909316821790925585526081909352922080546001600160e01b031916909117600160a01b90930292909217825560078201805467ffffffffffffffff60401b198116928c16600160401b02928317825560018801549495509293600160e01b900460ff1692909168ffffffffffffffffff60401b1990911660ff60801b1990911617600160801b8360028111156120005761200061492d565b021790555060026001850154600160e01b900460ff1660028111156120275761202761492d565b03612156578263ffffffff167f144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b89848a88600101601c9054906101000a900460ff16600281111561207a5761207a61492d565b8b60405161208c959493929190615544565b60405180910390a28263ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a64189848a6001600160a01b036040516120d7949392919061558b565b60405180910390a25f868060200190518101906120f491906154b2565b60405163b3a326f760e01b81529091506001600160a01b0384169063b3a326f790612123908490600401614f32565b5f604051808303815f87803b15801561213a575f5ffd5b505af115801561214c573d5f5f3e3d5ffd5b50505050506122ab565b60018481018054918301805467ffffffffffffffff60a01b198116600160a01b948590046001600160401b0316909402938417825591546001600160a01b03166001600160e01b03199092166001600160a01b0319909316929092171790556002808501545f8080529183016020908152604083209190915560038601546009840155875182918291829182916121f491908d018101908d01615610565b945094509450945094508763ffffffff167f194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a6418e898f8760405161223a949392919061558b565b60405180910390a2604051633892b81160e11b81526001600160a01b0388169063712570229061227890889088908d908990899089906004016156ac565b5f604051808303815f87803b15801561228f575f5ffd5b505af11580156122a1573d5f5f3e3d5ffd5b5050505050505050505b5050505050505050565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016123039190614f32565b602060405180830381865afa15801561231e573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061234291906153df565b6084549091505f90612366906001600160401b03600160401b8204811691166154cd565b6001600160401b03169050805f03612380575f9250505090565b61238a818361571e565b9250505090565b606f545f9060ff16156123b757604051630bc011ff60e21b815260040160405180910390fd5b335f9081526082602052604081205463ffffffff16908190036123ed576040516371653c1560e01b815260040160405180910390fd5b836001600160401b03165f0361241657604051632590ccf960e01b815260040160405180910390fd5b63ffffffff81165f908152608160205260408120906007820154600160801b900460ff16600281111561244b5761244b61492d565b14612469576040516390db0d0760e01b815260040160405180910390fd5b608480548691905f906124869084906001600160401b031661535b565b82546101009290920a6001600160401b0381810219909316918316021790915560068301541690505f6124b9878361535b565b6006840180546001600160401b0383811667ffffffffffffffff199092168217909255604080516060810182528a815242841660208083019182528886168385019081525f86815260038c018352859020935184559151600193909301805492518716600160401b026001600160801b03199093169390961692909217179093555190815291925063ffffffff8616917f1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25910160405180910390a29695505050505050565b6080545f9063ffffffff1680820361259757505f919050565b5f816001600160401b038111156125b0576125b0614d33565b6040519080825280602002602001820160405280156125d9578160200160208202803683370190505b5090505f5b828110156126365760815f6125f483600161549f565b63ffffffff1663ffffffff1681526020019081526020015f20600501548282815181106126235761262361548b565b60209081029190910101526001016125de565b505f60205b836001146127f3575f61264f600286615731565b61265a60028761571e565b612664919061549f565b90505f816001600160401b0381111561267f5761267f614d33565b6040519080825280602002602001820160405280156126a8578160200160208202803683370190505b5090505f5b828110156127c3576126c0600184615744565b811480156126d857506126d4600288615731565b6001145b1561273657612713866126ec83600261537a565b815181106126fc576126fc61548b565b6020026020010151865f9182526020526040902090565b8282815181106127255761272561548b565b6020026020010181815250506127bb565b61279c8661274583600261537a565b815181106127555761275561548b565b60200260200101518783600261276b919061537a565b61277690600161549f565b815181106127865761278661548b565b60200260200101515f9182526020526040902090565b8282815181106127ae576127ae61548b565b6020026020010181815250505b6001016126ad565b508094508195506127dd84855f9182526020526040902090565b9350826127e981615757565b935050505061263b565b5f835f815181106128065761280661548b565b602002602001015190505f5f90505b8281101561283a575f9182526020849052604080832085845292209350600101612815565b5095945050505050565b7fac75d24dbb35ea80e25fab167da4dea46c1915260426570db84f184891f5f59061286e81613894565b306001600160a01b038916036128975760405163325c055b60e21b815260040160405180910390fd5b607e80545f919082906128af9063ffffffff166154ec565b91906101000a81548163ffffffff021916908363ffffffff16021790559050600160028111156128e1576128e161492d565b8660028111156128f3576128f361492d565b0361291c578415612917576040516363d722e760e01b815260040160405180910390fd5b6129d6565b60028660028111156129305761293061492d565b03612986576001600160a01b03881615158061295457506001600160401b03871615155b8061295e57508415155b8061296857508215155b15612917576040516363d722e760e01b815260040160405180910390fd5b5f8660028111156129995761299961492d565b036129bd578215612917576040516363d722e760e01b815260040160405180910390fd5b6040516363d722e760e01b815260040160405180910390fd5b6040518060e001604052808a6001600160a01b03168152602001896001600160a01b03168152602001886001600160401b03168152602001876002811115612a2057612a2061492d565b81525f602080830182905260408084018a9052606093840188905263ffffffff86168352607f825291829020845181546001600160a01b039182166001600160a01b031990911617825591850151600182018054948701516001600160401b0316600160a01b026001600160e01b031990951691909316179290921780825592840151919260ff60e01b1916600160e01b836002811115612ac357612ac361492d565b02179055506080820151600182018054911515600160e81b0260ff60e81b1990921691909117905560a0820151600282015560c09091015160039091015560405163ffffffff8216907f9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a90612b45908c908c908c908c908c908c908c9061576c565b60405180910390a2505050505050505050565b5f5160206162465f395f51905f52612b6f81613894565b6001600160a01b0384165f9081526082602090815260408083205463ffffffff168084526081909252909120600263ffffffff86165f908152607f6020526040902060010154600160e01b900460ff166002811115612bd057612bd061492d565b14158015612c2f575063ffffffff85165f908152607f6020526040902060010154600160e01b900460ff166002811115612c0c57612c0c61492d565b6007820154600160801b900460ff166002811115612c2c57612c2c61492d565b14155b15612c4d57604051635aa0d5f160e11b815260040160405180910390fd5b612c58868686614157565b505050505050565b7f8cf807f6970720f8e2c208c7c5037595982c7bd9ed93c380d09df743d0dcc3fb612c8a81613894565b683635c9adc5dea00000821180612ca45750633b9aca0082105b15612cc257604051638586952560e01b815260040160405180910390fd5b60868290556040518281527ffb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b29060200160405180910390a15050565b5f82815260346020526040902060010154612d1881613894565b610d378383613cca565b7f62ba6ba2ffed8cfe316b583325ea41ac6e7ba9e5864d2bc6fabba7ac26d2f0f4612d4c81613894565b6087805467ffffffffffffffff1916426001600160401b031617905560408051636de0b4bb60e11b815290517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163dbc16976916004808301925f92919082900301818387803b158015612dc7575f5ffd5b505af1158015612dd9573d5f5f3e3d5ffd5b50505050612de5614402565b50565b63ffffffff86165f90815260816020526040902060609061188c9088908888888888613eba565b6001600160a01b0382165f9081526082602090815260408083205463ffffffff1683526081909152902060026007820154600160801b900460ff166002811115612e5b57612e5b61492d565b0361300457336001600160a01b0316836001600160a01b0316637388c4366040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ea6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612eca91906154b2565b6001600160a01b031614612ef15760405163660a7ce560e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020908152604091829020548251636e7fbce960e01b815292516001600160a01b0390911692636e7fbce99260048083019391928290030181865afa158015612f49573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f6d91906157cb565b6001600160f01b031916836001600160a01b0316636e7fbce96040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fb3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fd791906157cb565b6001600160f01b03191614612fff57604051635aa0d5f160e11b815260040160405180910390fd5b6130cc565b336001600160a01b0316836001600160a01b031663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561304a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061306e91906154b2565b6001600160a01b0316146130955760405163696072e960e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b90920416146130cc5760405163664316a560e11b815260040160405180910390fd5b600781015463ffffffff8316600160401b9091046001600160401b03161061310757604051633e37e23360e01b815260040160405180910390fd5b63ffffffff82165f908152607f6020526040902060010154600160e01b900460ff16600281111561313a5761313a61492d565b6007820154600160801b900460ff16600281111561315a5761315a61492d565b1461317857604051635aa0d5f160e11b815260040160405180910390fd5b604080515f815260208101909152610d379084908490614157565b5f5160206162465f395f51905f526131aa81613894565b63ffffffff84165f908152608160205260408120906007820154600160801b900460ff1660028111156131df576131df61492d565b146131fd576040516390db0d0760e01b815260040160405180910390fd5b60068101546001600160401b03808216600160401b90920416146132345760405163664316a560e11b815260040160405180910390fd5b5f63ffffffff85165f908152607f6020526040902060010154600160e01b900460ff1660028111156132685761326861492d565b03613286576040516301ea4e5b60e01b815260040160405180910390fd5b63ffffffff85165f908152608860205260409020805460ff1916600117905580546132bb906001600160a01b03168585614157565b60405163ffffffff85811682528616907f3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de99060200160405180910390a25050505050565b7f3dfe277d2a2c04b75fb2eb3743fa00005ae3678a20c299e65fdf4df76517f68e61332981613894565b6001600160401b0386165f9081526083602052604090205463ffffffff1615613365576040516337c8fe0960e11b815260040160405180910390fd5b63ffffffff6001600160401b038716111561339357604051634c753f5760e01b815260040160405180910390fd5b6001600160a01b0389165f9081526082602052604090205463ffffffff16156133cf57604051630d409b9360e41b815260040160405180910390fd5b6001600160a01b0389163014806133ee57506001600160a01b0389163b155b1561340c5760405163325c055b60e21b815260040160405180910390fd5b608080545f919082906134249063ffffffff166154ec565b91906101000a81548163ffffffff021916908363ffffffff160217905590508060835f896001600160401b03166001600160401b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055508060825f8c6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548163ffffffff021916908363ffffffff1602179055505f60815f8363ffffffff1663ffffffff1681526020019081526020015f2090508a815f015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550888160010160146101000a8154816001600160401b0302191690836001600160401b0316021790555089816001015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555087815f0160146101000a8154816001600160401b0302191690836001600160401b03160217905550858160070160106101000a81548160ff021916908360028111156135ab576135ab61492d565b021790555060018660028111156135c4576135c461492d565b0361360d576009810185905560088101849055600581018790556001600160a01b038a163b5f036136085760405163043103a360e21b815260040160405180910390fd5b6136e7565b60028660028111156136215761362161492d565b03613680576001600160a01b038a1615158061364557506001600160401b03891615155b8061364f57508415155b1561366d57604051636fc5557f60e01b815260040160405180910390fd5b60088101849055600581018790556136e7565b8415158061368d57508315155b156136ab57604051636fc5557f60e01b815260040160405180910390fd5b5f80805260028201602052604081208890556001600160a01b038b163b90036136e75760405163043103a360e21b815260040160405180910390fd5b8163ffffffff167f4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e8a8d8b8a5f8b8b60405161372997969594939291906157f2565b60405180910390a25050505050505050505050565b6137a260408051610180810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081018290529061016082015290565b63ffffffff82165f9081526081602090815260409182902080546001600160a01b038082168652600160a01b918290046001600160401b03908116948701949094526001830154908116948601949094529092048116606084015260058201546080840152600682015480821660a0850152600160401b808204831660c0860152600160801b808304841660e0870152600160c01b909204831661010086015260078401548084166101208701529081049092166101408501526101608401910460ff1660028111156138775761387761492d565b9081600281111561388a5761388a61492d565b8152505050919050565b612de58133614459565b5f5f6138bc89600601546001600160401b03600160401b9091041690565b60078a01549091506001600160401b0390811690891610156138f15760405163ead1340b60e01b815260040160405180910390fd5b6001600160401b0388165f90815260028a01602052604090205491508161392b576040516324cbdcc360e11b815260040160405180910390fd5b806001600160401b0316886001600160401b0316111561395e57604051630f2b74f160e11b815260040160405180910390fd5b806001600160401b0316876001600160401b0316116139905760405163b9b18f5760e01b815260040160405180910390fd5b5f61399f8a8a8a8a878b614054565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016002836040516139d3919061584d565b602060405180830381855afa1580156139ee573d5f5f3e3d5ffd5b5050506040513d601f19601f82011682018060405250810190613a1191906153df565b613a1b9190615731565b60018c0154604080516020810182528381529051634890ed4560e11b81529293506001600160a01b0390911691639121da8a91613a5d91899190600401615863565b602060405180830381865afa158015613a78573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a9c919061589f565b613ab9576040516309bde33960e01b815260040160405180910390fd5b5f613ac4848b6154cd565b9050613b1787826001600160401b0316613adc6122b5565b613ae6919061537a565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190614480565b80608460088282829054906101000a90046001600160401b0316613b3b919061535b565b82546101009290920a6001600160401b038181021990931691831602179091556084805467ffffffffffffffff60801b1916600160801b428416021790558d546040516332c2d15360e01b8152918d166004830152602482018b90523360448301526001600160a01b031691506332c2d153906064015f604051808303815f87803b158015613bc8575f5ffd5b505af1158015613bda573d5f5f3e3d5ffd5b50505050505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632072f6c56040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613c44575f5ffd5b505af1158015613c56573d5f5f3e3d5ffd5b50505050610d116144d2565b613c6c8282611d59565b610d6f575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b613cd48282611d59565b15610d6f575f8281526034602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b5f67ffffffff000000016001600160401b038316108015613d65575067ffffffff00000001604083901c6001600160401b0316105b8015613d85575067ffffffff00000001608083901c6001600160401b0316105b8015613d9c575067ffffffff0000000160c083901c105b15613da957506001919050565b505f919050565b919050565b5f54610100900460ff16610d115760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610dce565b610d6f8282613c62565b5f82815260346020526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b5f5160206162665f395f51905f525c15613ea057604051633ee5aeb560e01b815260040160405180910390fd5b610d1160015f5160206162665f395f51905f525b9061452d565b606060026007880154600160801b900460ff166002811115613ede57613ede61492d565b03613f90578654604051631a957d9b60e21b81525f916001600160a01b031690636a55f66c90613f149087908790600401615478565b602060405180830381865afa158015613f2f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f5391906153df565b600589015460088a0154604051929350613f79928a908d9086908c908c906020016158be565b60405160208183030381529060405291505061188c565b865460408051632b47b7cd60e21b815290515f926001600160a01b03169163ad1edf349160048083019260209291908290030181865afa158015613fd6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ffa91906153df565b600589015460088a0154604051929350614020928a908d9086908c908c906020016158be565b604051602081830303815290604052915050979650505050505050565b610d115f5f5160206162665f395f51905f52613eb4565b6001600160401b038086165f8181526003890160205260408082205493881682529020546060929115801590614088575081155b156140a65760405163340c614f60e11b815260040160405180910390fd5b806140c4576040516366385b5160e01b815260040160405180910390fd5b6140cd84613d30565b6140ea576040516305dae44f60e21b815260040160405180910390fd5b3385838a8c5f0160149054906101000a90046001600160401b03168d60010160149054906101000a90046001600160401b031689878d8f60405160200161413a9a999897969594939291906158fb565b604051602081830303815290604052925050509695505050505050565b63ffffffff821615806141755750607e5463ffffffff908116908316115b1561419357604051637512e5cb60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526082602052604081205463ffffffff16908190036141d2576040516374a086a360e01b815260040160405180910390fd5b63ffffffff8181165f908152608160205260409020600781015490918516600160401b9091046001600160401b03160361421f57604051634f61d51960e01b815260040160405180910390fd5b63ffffffff84165f908152607f602052604090206001810154600160e81b900460ff161561426057604051633b8d3d9960e01b815260040160405180910390fd5b6001808201805491840180546001600160a01b031981166001600160a01b03909416938417825582546001600160401b03600160a01b9182900416026001600160e01b03199091169093179290921790915560038201546009840155600783018054600160401b63ffffffff89160267ffffffffffffffff60401b19821681178355925460ff600160e01b909104169260ff60801b191668ffffffffffffffffff60401b1990911617600160801b8360028111156143205761432061492d565b02179055505f61432f84610a63565b60078401805467ffffffffffffffff19166001600160401b038316179055825460405163278f794360e11b81529192506001600160a01b0389811692634f1ef286926143819216908990600401615997565b5f604051808303815f87803b158015614398575f5ffd5b505af11580156143aa573d5f5f3e3d5ffd5b50506040805163ffffffff8a811682526001600160401b0386166020830152881693507ff585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d92500160405180910390a250505050505050565b606f5460ff1661442557604051635386698160e01b815260040160405180910390fd5b606f805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b6144638282611d59565b610d6f57604051637615be1f60e11b815260040160405180910390fd5b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610d37908490614534565b606f5460ff16156144f657604051630bc011ff60e21b815260040160405180910390fd5b606f805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b80825d5050565b5f614588826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166146059092919063ffffffff16565b805190915015610d3757808060200190518101906145a6919061589f565b610d375760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610dce565b606061461384845f8561461b565b949350505050565b60608247101561467c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610dce565b5f5f866001600160a01b03168587604051614697919061584d565b5f6040518083038185875af1925050503d805f81146146d1576040519150601f19603f3d011682016040523d82523d5f602084013e6146d6565b606091505b509150915061188c87838387606083156147505782515f03614749576001600160a01b0385163b6147495760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610dce565b5081614613565b61461383838151156147655781518083602001fd5b8060405162461bcd60e51b8152600401610dce919061491b565b61088b806159bb83390190565b803563ffffffff81168114613db0575f5ffd5b5f602082840312156147af575f5ffd5b6147b88261478c565b9392505050565b80356001600160401b0381168114613db0575f5ffd5b6001600160a01b0381168114612de5575f5ffd5b5f5f5f5f5f5f5f5f6103e0898b031215614801575f5ffd5b61480a8961478c565b975061481860208a016147bf565b965061482660408a016147bf565b955061483460608a016147bf565b94506080890135935060a0890135925060c0890135614852816147d5565b91506103e089018a1015614864575f5ffd5b60e0890190509295985092959890939650565b5f60208284031215614887575f5ffd5b5035919050565b5f5f6040838503121561489f575f5ffd5b6148a88361478c565b91506148b6602084016147bf565b90509250929050565b5f5f604083850312156148d0575f5ffd5b8235915060208301356148e2816147d5565b809150509250929050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6147b860208301846148ed565b634e487b7160e01b5f52602160045260245ffd5b6003811061495d57634e487b7160e01b5f52602160045260245ffd5b9052565b6001600160a01b038881168252871660208201526001600160401b038616604082015260e081016149956060830187614941565b931515608082015260a081019290925260c090910152949350505050565b5f5f5f606084860312156149c5575f5ffd5b6149ce8461478c565b95602085013595506040909401359392505050565b5f5f5f5f608085870312156149f6575f5ffd5b8435614a01816147d5565b93506020850135614a11816147d5565b92506040850135614a21816147d5565b91506060850135614a31816147d5565b939692955090935050565b5f5f83601f840112614a4c575f5ffd5b5081356001600160401b03811115614a62575f5ffd5b602083019150836020828501011115614a79575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f60c0898b031215614a97575f5ffd5b614aa08961478c565b9750614aae60208a0161478c565b9650604089013595506060890135945060808901356001600160401b03811115614ad6575f5ffd5b614ae28b828c01614a3c565b90955093505060a08901356001600160401b03811115614b00575f5ffd5b614b0c8b828c01614a3c565b999c989b5096995094979396929594505050565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b038616610100830152614b95610120830186614941565b61014082019390935261016001529a9950505050505050505050565b81516001600160a01b0316815261018081016020830151614bdd60208401826001600160401b03169052565b506040830151614bf860408401826001600160a01b03169052565b506060830151614c1360608401826001600160401b03169052565b506080830151608083015260a0830151614c3860a08401826001600160401b03169052565b5060c0830151614c5360c08401826001600160401b03169052565b5060e0830151614c6e60e08401826001600160401b03169052565b50610100830151614c8b6101008401826001600160401b03169052565b50610120830151614ca0610120840182614941565b5061014083015161014083015261016083015161016083015292915050565b5f5f5f5f5f5f60c08789031215614cd4575f5ffd5b614cdd8761478c565b9550614ceb602088016147bf565b9450614cf9604088016147bf565b959894975094956060810135955060808101359460a0909101359350915050565b5f60208284031215614d2a575f5ffd5b6147b8826147bf565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715614d6f57614d6f614d33565b604052919050565b5f60208284031215614d87575f5ffd5b81356001600160401b03811115614d9c575f5ffd5b8201601f81018413614dac575f5ffd5b80356001600160401b03811115614dc557614dc5614d33565b8060051b614dd560208201614d47565b91825260208184018101929081019087841115614df0575f5ffd5b6020850194505b8385101561188c57843580835260209586019590935090910190614df7565b5f5f60408385031215614e27575f5ffd5b82356148a8816147d5565b5f6001600160401b03821115614e4a57614e4a614d33565b50601f01601f191660200190565b5f614e6a614e6584614e32565b614d47565b9050828152838383011115614e7d575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112614ea2575f5ffd5b6147b883833560208501614e58565b5f5f5f60608486031215614ec3575f5ffd5b614ecc8461478c565b9250614eda602085016147bf565b915060408401356001600160401b03811115614ef4575f5ffd5b614f0086828701614e93565b9150509250925092565b5f5f60408385031215614f1b575f5ffd5b614f24836147bf565b946020939093013593505050565b6001600160a01b0391909116815260200190565b803560038110613db0575f5ffd5b5f5f5f5f5f5f5f60e0888a031215614f6a575f5ffd5b8735614f75816147d5565b96506020880135614f85816147d5565b9550614f93604089016147bf565b9450614fa160608901614f46565b93506080880135925060a08801356001600160401b03811115614fc2575f5ffd5b8801601f81018a13614fd2575f5ffd5b614fe18a823560208401614e58565b979a969950949793969295929450505060c09091013590565b5f5f5f6060848603121561500c575f5ffd5b8335615017816147d5565b9250614eda6020850161478c565b5f60208284031215615035575f5ffd5b81356147b8816147d5565b5f5f5f5f5f5f60a08789031215615055575f5ffd5b61505e8761478c565b955060208701359450604087013593506060870135925060808701356001600160401b0381111561508d575f5ffd5b61509989828a01614a3c565b979a9699509497509295939492505050565b5f5f604083850312156150bc575f5ffd5b82356150c7816147d5565b91506148b66020840161478c565b6001600160a01b038d811682526001600160401b038d81166020840152908c1660408301528a81166060830152608082018a905288811660a0830152871660c082015261018081016001600160401b03871660e08301526001600160401b0386166101008301526001600160401b0385166101208301526001600160401b03841661014083015261516a610160830184614941565b9d9c50505050505050505050505050565b5f5f5f6060848603121561518d575f5ffd5b6150178461478c565b5f5f5f5f5f5f5f5f610100898b0312156151ae575f5ffd5b88356151b9816147d5565b975060208901356151c9816147d5565b96506151d760408a016147bf565b95506151e560608a016147bf565b9450608089013593506151fa60a08a01614f46565b979a969950949793969295929450505060c08201359160e0013590565b81516001600160a01b031681526101808101602083015161524360208401826001600160401b03169052565b50604083015161525e60408401826001600160a01b03169052565b50606083015161527960608401826001600160401b03169052565b506080830151608083015260a083015161529e60a08401826001600160401b03169052565b5060c08301516152b960c08401826001600160401b03169052565b5060e08301516152d460e08401826001600160401b03169052565b506101008301516152f16101008401826001600160401b03169052565b5061012083015161530e6101208401826001600160401b03169052565b5061014083015161532b6101408401826001600160401b03169052565b50610160830151615340610160840182614941565b5092915050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190811115610a8c57610a8c615347565b8082028115828204841417610a8c57610a8c615347565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b5f602082840312156153ef575f5ffd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f61543060408301866148ed565b82810360208401526154438185876153f6565b9695505050505050565b848152606060208201525f61546560608301866148ed565b828103604084015261188c8185876153f6565b602081525f6146136020830184866153f6565b634e487b7160e01b5f52603260045260245ffd5b80820180821115610a8c57610a8c615347565b5f602082840312156154c2575f5ffd5b81516147b8816147d5565b6001600160401b038281168282160390811115610a8c57610a8c615347565b5f63ffffffff821663ffffffff810361550757615507615347565b60010192915050565b6001600160a01b038481168252831660208201526060604082018190525f9061553b908301846148ed565b95945050505050565b63ffffffff861681526001600160a01b03851660208201526001600160401b038416604082015260ff8316606082015260a0608082018190525f9061188c908301846148ed565b63ffffffff9490941684526001600160a01b0392831660208501526001600160401b0391909116604084015216606082015260800190565b5f82601f8301126155d2575f5ffd5b81516155e0614e6582614e32565b8181528460208386010111156155f4575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215615624575f5ffd5b855161562f816147d5565b6020870151909550615640816147d5565b6040870151909450615651816147d5565b60608701519093506001600160401b0381111561566c575f5ffd5b615678888289016155c3565b92505060808601516001600160401b03811115615693575f5ffd5b61569f888289016155c3565b9150509295509295909350565b6001600160a01b038781168252868116602083015263ffffffff861660408301528416606082015260c0608082018190525f906156eb908301856148ed565b82810360a08401526156fd81856148ed565b9998505050505050505050565b634e487b7160e01b5f52601260045260245ffd5b5f8261572c5761572c61570a565b500490565b5f8261573f5761573f61570a565b500690565b81810381811115610a8c57610a8c615347565b5f8161576557615765615347565b505f190190565b6001600160a01b038881168252871660208201526001600160401b038616604082015261579c6060820186614941565b83608082015260e060a08201525f6157b760e08301856148ed565b90508260c083015298975050505050505050565b5f602082840312156157db575f5ffd5b81516001600160f01b0319811681146147b8575f5ffd5b6001600160401b0388811682526001600160a01b03881660208301528616604082015260e081016158266060830187614941565b6001600160401b03851660808301528360a08301528260c083015298975050505050505050565b5f82518060208501845e5f920191825250919050565b61032081016103008483376103008201835f5b6001811015615895578151835260209283019290910190600101615876565b5050509392505050565b5f602082840312156158af575f5ffd5b815180151581146147b8575f5ffd5b9687526020870195909552604086019390935260e09190911b6001600160e01b03191660608501526064840152608483015260a482015260c40190565b6bffffffffffffffffffffffff198b60601b1681528960148201528860348201526001600160401b0360c01b8860c01b1660548201526001600160401b0360c01b8760c01b16605c8201526001600160401b0360c01b8660c01b16606482015284606c82015283608c8201528260ac82015261598660cc82018360c01b6001600160c01b0319169052565b60d4019a9950505050505050505050565b6001600160a01b03831681526040602082018190525f90614613908301846148ed56fe60a060405260405161088b38038061088b83398101604081905261002291610327565b828161002e8282610056565b50506001600160a01b03821660805261004e61004960805190565b6100b4565b50505061040e565b61005f82610121565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156100a8576100a3828261019f565b505050565b6100b0610212565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6100f35f51602061086b5f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161011e81610233565b50565b806001600160a01b03163b5f0361015b57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b60605f5f846001600160a01b0316846040516101bb91906103f8565b5f60405180830381855af49150503d805f81146101f3576040519150601f19603f3d011682016040523d82523d5f602084013e6101f8565b606091505b509092509050610209858383610270565b95945050505050565b34156102315760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661025c57604051633173bdd160e11b81525f6004820152602401610152565b805f51602061086b5f395f51905f5261017e565b60608261028557610280826102cf565b6102c8565b815115801561029c57506001600160a01b0384163b155b156102c557604051639996b31560e01b81526001600160a01b0385166004820152602401610152565b50805b9392505050565b8051156102df5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80516001600160a01b038116811461030e575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5f5f60608486031215610339575f5ffd5b610342846102f8565b9250610350602085016102f8565b60408501519092506001600160401b0381111561036b575f5ffd5b8401601f8101861361037b575f5ffd5b80516001600160401b0381111561039457610394610313565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103c2576103c2610313565b6040528181528282016020018810156103d9575f5ffd5b8160208401602083015e5f602083830101528093505050509250925092565b5f82518060208501845e5f920191825250919050565b6080516104466104255f395f601001526104465ff3fe608060405261000c61000e565b005b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163303610069575f356001600160e01b03191663278f794360e11b146100615761005f61006d565b565b61005f61007d565b61005f5b61005f6100786100ab565b6100e2565b5f8061008c36600481846102ee565b8101906100999190610329565b915091506100a78282610100565b5050565b5f6100dd7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f5f375f5f365f845af43d5f5f3e8080156100fc573d5ff35b3d5ffd5b6101098261015a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156101525761014d82826101d5565b505050565b6100a7610247565b806001600160a01b03163b5f0361019457604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b0316846040516101f191906103fa565b5f60405180830381855af49150503d805f8114610229576040519150601f19603f3d011682016040523d82523d5f602084013e61022e565b606091505b509150915061023e858383610266565b95945050505050565b341561005f5760405163b398979f60e01b815260040160405180910390fd5b60608261027b57610276826102c5565b6102be565b815115801561029257506001600160a01b0384163b155b156102bb57604051639996b31560e01b81526001600160a01b038516600482015260240161018b565b50805b9392505050565b8051156102d55780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5f5f858511156102fc575f5ffd5b83861115610308575f5ffd5b5050820193919092039150565b634e487b7160e01b5f52604160045260245ffd5b5f5f6040838503121561033a575f5ffd5b82356001600160a01b0381168114610350575f5ffd5b9150602083013567ffffffffffffffff81111561036b575f5ffd5b8301601f8101851361037b575f5ffd5b803567ffffffffffffffff81111561039557610395610315565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156103c4576103c4610315565b6040528181528282016020018710156103db575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220c0931976df4a26c2184205b53cad6447a86a5b19b88f07a03b24865328b06bac64736f6c634300081c0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610366156603fe29d13f97c6f3e3dff4ef71919f9aa61c555be0182d954e94221aac9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00141f8f32ce6198eee741f695cec728bfd32d289f1acf73621fb303581000545e084e94f375e9d647f87f5b2ceffba1e062c70f6009fdbcf80291e803b5c9edd4a2646970667358221220e433f392d282b507215a45b2cedb7c185a3b202c16e27f48fb00e4a47e3645eb64736f6c634300081c0033",
}

// AgglayermanagermockABI is the input ABI used to generate the binding from.
// Deprecated: Use AgglayermanagermockMetaData.ABI instead.
var AgglayermanagermockABI = AgglayermanagermockMetaData.ABI

// AgglayermanagermockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgglayermanagermockMetaData.Bin instead.
var AgglayermanagermockBin = AgglayermanagermockMetaData.Bin

// DeployAgglayermanagermock deploys a new Ethereum contract, binding an instance of Agglayermanagermock to it.
func DeployAgglayermanagermock(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Agglayermanagermock, error) {
	parsed, err := AgglayermanagermockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgglayermanagermockBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agglayermanagermock{AgglayermanagermockCaller: AgglayermanagermockCaller{contract: contract}, AgglayermanagermockTransactor: AgglayermanagermockTransactor{contract: contract}, AgglayermanagermockFilterer: AgglayermanagermockFilterer{contract: contract}}, nil
}

// Agglayermanagermock is an auto generated Go binding around an Ethereum contract.
type Agglayermanagermock struct {
	AgglayermanagermockCaller     // Read-only binding to the contract
	AgglayermanagermockTransactor // Write-only binding to the contract
	AgglayermanagermockFilterer   // Log filterer for contract events
}

// AgglayermanagermockCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgglayermanagermockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagermockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgglayermanagermockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagermockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgglayermanagermockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgglayermanagermockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgglayermanagermockSession struct {
	Contract     *Agglayermanagermock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AgglayermanagermockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgglayermanagermockCallerSession struct {
	Contract *AgglayermanagermockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AgglayermanagermockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgglayermanagermockTransactorSession struct {
	Contract     *AgglayermanagermockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AgglayermanagermockRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgglayermanagermockRaw struct {
	Contract *Agglayermanagermock // Generic contract binding to access the raw methods on
}

// AgglayermanagermockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgglayermanagermockCallerRaw struct {
	Contract *AgglayermanagermockCaller // Generic read-only contract binding to access the raw methods on
}

// AgglayermanagermockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgglayermanagermockTransactorRaw struct {
	Contract *AgglayermanagermockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgglayermanagermock creates a new instance of Agglayermanagermock, bound to a specific deployed contract.
func NewAgglayermanagermock(address common.Address, backend bind.ContractBackend) (*Agglayermanagermock, error) {
	contract, err := bindAgglayermanagermock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agglayermanagermock{AgglayermanagermockCaller: AgglayermanagermockCaller{contract: contract}, AgglayermanagermockTransactor: AgglayermanagermockTransactor{contract: contract}, AgglayermanagermockFilterer: AgglayermanagermockFilterer{contract: contract}}, nil
}

// NewAgglayermanagermockCaller creates a new read-only instance of Agglayermanagermock, bound to a specific deployed contract.
func NewAgglayermanagermockCaller(address common.Address, caller bind.ContractCaller) (*AgglayermanagermockCaller, error) {
	contract, err := bindAgglayermanagermock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockCaller{contract: contract}, nil
}

// NewAgglayermanagermockTransactor creates a new write-only instance of Agglayermanagermock, bound to a specific deployed contract.
func NewAgglayermanagermockTransactor(address common.Address, transactor bind.ContractTransactor) (*AgglayermanagermockTransactor, error) {
	contract, err := bindAgglayermanagermock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockTransactor{contract: contract}, nil
}

// NewAgglayermanagermockFilterer creates a new log filterer instance of Agglayermanagermock, bound to a specific deployed contract.
func NewAgglayermanagermockFilterer(address common.Address, filterer bind.ContractFilterer) (*AgglayermanagermockFilterer, error) {
	contract, err := bindAgglayermanagermock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockFilterer{contract: contract}, nil
}

// bindAgglayermanagermock binds a generic wrapper to an already deployed contract.
func bindAgglayermanagermock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgglayermanagermockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanagermock *AgglayermanagermockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanagermock.Contract.AgglayermanagermockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanagermock *AgglayermanagermockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AgglayermanagermockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanagermock *AgglayermanagermockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AgglayermanagermockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agglayermanagermock *AgglayermanagermockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agglayermanagermock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agglayermanagermock *AgglayermanagermockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agglayermanagermock *AgglayermanagermockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanagermock.Contract.DEFAULTADMINROLE(&_Agglayermanagermock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Agglayermanagermock.Contract.DEFAULTADMINROLE(&_Agglayermanagermock.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagermock *AgglayermanagermockCaller) ROLLUPMANAGERVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "ROLLUP_MANAGER_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagermock *AgglayermanagermockSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanagermock.Contract.ROLLUPMANAGERVERSION(&_Agglayermanagermock.CallOpts)
}

// ROLLUPMANAGERVERSION is a free data retrieval call binding the contract method 0xd8905812.
//
// Solidity: function ROLLUP_MANAGER_VERSION() view returns(string)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) ROLLUPMANAGERVERSION() (string, error) {
	return _Agglayermanagermock.Contract.ROLLUPMANAGERVERSION(&_Agglayermanagermock.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanagermock.Contract.AggLayerGateway(&_Agglayermanagermock.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) AggLayerGateway() (common.Address, error) {
	return _Agglayermanagermock.Contract.AggLayerGateway(&_Agglayermanagermock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanagermock.Contract.BridgeAddress(&_Agglayermanagermock.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) BridgeAddress() (common.Address, error) {
	return _Agglayermanagermock.Contract.BridgeAddress(&_Agglayermanagermock.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCaller) CalculateRewardPerBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "calculateRewardPerBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanagermock.Contract.CalculateRewardPerBatch(&_Agglayermanagermock.CallOpts)
}

// CalculateRewardPerBatch is a free data retrieval call binding the contract method 0x99f5634e.
//
// Solidity: function calculateRewardPerBatch() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) CalculateRewardPerBatch() (*big.Int, error) {
	return _Agglayermanagermock.Contract.CalculateRewardPerBatch(&_Agglayermanagermock.CallOpts)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockCaller) ChainIDToRollupID(opts *bind.CallOpts, chainID uint64) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "chainIDToRollupID", chainID)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanagermock.Contract.ChainIDToRollupID(&_Agglayermanagermock.CallOpts, chainID)
}

// ChainIDToRollupID is a free data retrieval call binding the contract method 0x7fb6e76a.
//
// Solidity: function chainIDToRollupID(uint64 chainID) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) ChainIDToRollupID(chainID uint64) (uint32, error) {
	return _Agglayermanagermock.Contract.ChainIDToRollupID(&_Agglayermanagermock.CallOpts, chainID)
}

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCaller) ExposedCheckStateRootInsidePrime(opts *bind.CallOpts, newStateRoot *big.Int) (bool, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "exposed_checkStateRootInsidePrime", newStateRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Agglayermanagermock *AgglayermanagermockSession) ExposedCheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Agglayermanagermock.Contract.ExposedCheckStateRootInsidePrime(&_Agglayermanagermock.CallOpts, newStateRoot)
}

// ExposedCheckStateRootInsidePrime is a free data retrieval call binding the contract method 0x62d87e66.
//
// Solidity: function exposed_checkStateRootInsidePrime(uint256 newStateRoot) pure returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) ExposedCheckStateRootInsidePrime(newStateRoot *big.Int) (bool, error) {
	return _Agglayermanagermock.Contract.ExposedCheckStateRootInsidePrime(&_Agglayermanagermock.CallOpts, newStateRoot)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanagermock.Contract.GetBatchFee(&_Agglayermanagermock.CallOpts)
}

// GetBatchFee is a free data retrieval call binding the contract method 0x477fa270.
//
// Solidity: function getBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetBatchFee() (*big.Int, error) {
	return _Agglayermanagermock.Contract.GetBatchFee(&_Agglayermanagermock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetForcedBatchFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getForcedBatchFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanagermock.Contract.GetForcedBatchFee(&_Agglayermanagermock.CallOpts)
}

// GetForcedBatchFee is a free data retrieval call binding the contract method 0x60469169.
//
// Solidity: function getForcedBatchFee() view returns(uint256)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetForcedBatchFee() (*big.Int, error) {
	return _Agglayermanagermock.Contract.GetForcedBatchFee(&_Agglayermanagermock.CallOpts)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetInputPessimisticBytes(opts *bind.CallOpts, rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getInputPessimisticBytes", rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanagermock.Contract.GetInputPessimisticBytes(&_Agglayermanagermock.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputPessimisticBytes is a free data retrieval call binding the contract method 0xdd0464b9.
//
// Solidity: function getInputPessimisticBytes(uint32 rollupID, bytes32 l1InfoTreeRoot, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes aggchainData) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetInputPessimisticBytes(rollupID uint32, l1InfoTreeRoot [32]byte, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, aggchainData []byte) ([]byte, error) {
	return _Agglayermanagermock.Contract.GetInputPessimisticBytes(&_Agglayermanagermock.CallOpts, rollupID, l1InfoTreeRoot, newLocalExitRoot, newPessimisticRoot, aggchainData)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetInputSnarkBytes(opts *bind.CallOpts, rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getInputSnarkBytes", rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanagermock.Contract.GetInputSnarkBytes(&_Agglayermanagermock.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetInputSnarkBytes is a free data retrieval call binding the contract method 0x7975fcfe.
//
// Solidity: function getInputSnarkBytes(uint32 rollupID, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 oldStateRoot, bytes32 newStateRoot) view returns(bytes)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetInputSnarkBytes(rollupID uint32, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, oldStateRoot [32]byte, newStateRoot [32]byte) ([]byte, error) {
	return _Agglayermanagermock.Contract.GetInputSnarkBytes(&_Agglayermanagermock.CallOpts, rollupID, initNumBatch, finalNewBatch, newLocalExitRoot, oldStateRoot, newStateRoot)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetLastVerifiedBatch(opts *bind.CallOpts, rollupID uint32) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getLastVerifiedBatch", rollupID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanagermock.Contract.GetLastVerifiedBatch(&_Agglayermanagermock.CallOpts, rollupID)
}

// GetLastVerifiedBatch is a free data retrieval call binding the contract method 0x11f6b287.
//
// Solidity: function getLastVerifiedBatch(uint32 rollupID) view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetLastVerifiedBatch(rollupID uint32) (uint64, error) {
	return _Agglayermanagermock.Contract.GetLastVerifiedBatch(&_Agglayermanagermock.CallOpts, rollupID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRoleAdmin(&_Agglayermanagermock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRoleAdmin(&_Agglayermanagermock.CallOpts, role)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetRollupBatchNumToStateRoot(opts *bind.CallOpts, rollupID uint32, batchNum uint64) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getRollupBatchNumToStateRoot", rollupID, batchNum)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupBatchNumToStateRoot is a free data retrieval call binding the contract method 0x55a71ee0.
//
// Solidity: function getRollupBatchNumToStateRoot(uint32 rollupID, uint64 batchNum) view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetRollupBatchNumToStateRoot(rollupID uint32, batchNum uint64) ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRollupBatchNumToStateRoot(&_Agglayermanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCaller) GetRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRollupExitRoot(&_Agglayermanagermock.CallOpts)
}

// GetRollupExitRoot is a free data retrieval call binding the contract method 0xa2967d99.
//
// Solidity: function getRollupExitRoot() view returns(bytes32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetRollupExitRoot() ([32]byte, error) {
	return _Agglayermanagermock.Contract.GetRollupExitRoot(&_Agglayermanagermock.CallOpts)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagermock *AgglayermanagermockCaller) GetRollupSequencedBatches(opts *bind.CallOpts, rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "getRollupSequencedBatches", rollupID, batchNum)

	if err != nil {
		return *new(LegacyZKEVMStateVariablesSequencedBatchData), err
	}

	out0 := *abi.ConvertType(out[0], new(LegacyZKEVMStateVariablesSequencedBatchData)).(*LegacyZKEVMStateVariablesSequencedBatchData)

	return out0, err

}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagermock *AgglayermanagermockSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanagermock.Contract.GetRollupSequencedBatches(&_Agglayermanagermock.CallOpts, rollupID, batchNum)
}

// GetRollupSequencedBatches is a free data retrieval call binding the contract method 0x25280169.
//
// Solidity: function getRollupSequencedBatches(uint32 rollupID, uint64 batchNum) view returns((bytes32,uint64,uint64))
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GetRollupSequencedBatches(rollupID uint32, batchNum uint64) (LegacyZKEVMStateVariablesSequencedBatchData, error) {
	return _Agglayermanagermock.Contract.GetRollupSequencedBatches(&_Agglayermanagermock.CallOpts, rollupID, batchNum)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanagermock.Contract.GlobalExitRootManager(&_Agglayermanagermock.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Agglayermanagermock.Contract.GlobalExitRootManager(&_Agglayermanagermock.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanagermock.Contract.HasRole(&_Agglayermanagermock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Agglayermanagermock.Contract.HasRole(&_Agglayermanagermock.CallOpts, role, account)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockSession) IsEmergencyState() (bool, error) {
	return _Agglayermanagermock.Contract.IsEmergencyState(&_Agglayermanagermock.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) IsEmergencyState() (bool, error) {
	return _Agglayermanagermock.Contract.IsEmergencyState(&_Agglayermanagermock.CallOpts)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCaller) IsRollupMigrating(opts *bind.CallOpts, rollupID uint32) (bool, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "isRollupMigrating", rollupID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanagermock.Contract.IsRollupMigrating(&_Agglayermanagermock.CallOpts, rollupID)
}

// IsRollupMigrating is a free data retrieval call binding the contract method 0x3a7094bd.
//
// Solidity: function isRollupMigrating(uint32 rollupID) view returns(bool)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) IsRollupMigrating(rollupID uint32) (bool, error) {
	return _Agglayermanagermock.Contract.IsRollupMigrating(&_Agglayermanagermock.CallOpts, rollupID)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCaller) LastAggregationTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "lastAggregationTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanagermock.Contract.LastAggregationTimestamp(&_Agglayermanagermock.CallOpts)
}

// LastAggregationTimestamp is a free data retrieval call binding the contract method 0xc1acbc34.
//
// Solidity: function lastAggregationTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) LastAggregationTimestamp() (uint64, error) {
	return _Agglayermanagermock.Contract.LastAggregationTimestamp(&_Agglayermanagermock.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCaller) LastDeactivatedEmergencyStateTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "lastDeactivatedEmergencyStateTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanagermock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanagermock.CallOpts)
}

// LastDeactivatedEmergencyStateTimestamp is a free data retrieval call binding the contract method 0x30c27dde.
//
// Solidity: function lastDeactivatedEmergencyStateTimestamp() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) LastDeactivatedEmergencyStateTimestamp() (uint64, error) {
	return _Agglayermanagermock.Contract.LastDeactivatedEmergencyStateTimestamp(&_Agglayermanagermock.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockSession) Pol() (common.Address, error) {
	return _Agglayermanagermock.Contract.Pol(&_Agglayermanagermock.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) Pol() (common.Address, error) {
	return _Agglayermanagermock.Contract.Pol(&_Agglayermanagermock.CallOpts)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupAddressToID(opts *bind.CallOpts, rollupAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupAddressToID", rollupAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanagermock.Contract.RollupAddressToID(&_Agglayermanagermock.CallOpts, rollupAddress)
}

// RollupAddressToID is a free data retrieval call binding the contract method 0xceee281d.
//
// Solidity: function rollupAddressToID(address rollupAddress) view returns(uint32 rollupID)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupAddressToID(rollupAddress common.Address) (uint32, error) {
	return _Agglayermanagermock.Contract.RollupAddressToID(&_Agglayermanagermock.CallOpts, rollupAddress)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockSession) RollupCount() (uint32, error) {
	return _Agglayermanagermock.Contract.RollupCount(&_Agglayermanagermock.CallOpts)
}

// RollupCount is a free data retrieval call binding the contract method 0xf4e92675.
//
// Solidity: function rollupCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupCount() (uint32, error) {
	return _Agglayermanagermock.Contract.RollupCount(&_Agglayermanagermock.CallOpts)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupIDToRollupData(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupIDToRollupData", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturn)).(*AgglayerManagerRollupDataReturn)

	return out0, err

}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagermock *AgglayermanagermockSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanagermock.Contract.RollupIDToRollupData(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupData is a free data retrieval call binding the contract method 0xf9c4c2ae.
//
// Solidity: function rollupIDToRollupData(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint64,uint64,uint8) rollupData)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupIDToRollupData(rollupID uint32) (AgglayerManagerRollupDataReturn, error) {
	return _Agglayermanagermock.Contract.RollupIDToRollupData(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupIDToRollupDataDeserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupIDToRollupDataDeserialized", rollupID)

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
func (_Agglayermanagermock *AgglayermanagermockSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanagermock.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataDeserialized is a free data retrieval call binding the contract method 0xe4f3d8f9.
//
// Solidity: function rollupIDToRollupDataDeserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 legacyLastPendingState, uint64 legacyLastPendingStateConsolidated, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupIDToRollupDataDeserialized(rollupID uint32) (struct {
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
	return _Agglayermanagermock.Contract.RollupIDToRollupDataDeserialized(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupIDToRollupDataV2(opts *bind.CallOpts, rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupIDToRollupDataV2", rollupID)

	if err != nil {
		return *new(AgglayerManagerRollupDataReturnV2), err
	}

	out0 := *abi.ConvertType(out[0], new(AgglayerManagerRollupDataReturnV2)).(*AgglayerManagerRollupDataReturnV2)

	return out0, err

}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagermock *AgglayermanagermockSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanagermock.Contract.RollupIDToRollupDataV2(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2 is a free data retrieval call binding the contract method 0x74d9c244.
//
// Solidity: function rollupIDToRollupDataV2(uint32 rollupID) view returns((address,uint64,address,uint64,bytes32,uint64,uint64,uint64,uint64,uint8,bytes32,bytes32) rollupData)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupIDToRollupDataV2(rollupID uint32) (AgglayerManagerRollupDataReturnV2, error) {
	return _Agglayermanagermock.Contract.RollupIDToRollupDataV2(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupIDToRollupDataV2Deserialized(opts *bind.CallOpts, rollupID uint32) (struct {
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
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupIDToRollupDataV2Deserialized", rollupID)

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
func (_Agglayermanagermock *AgglayermanagermockSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanagermock.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupIDToRollupDataV2Deserialized is a free data retrieval call binding the contract method 0x70603909.
//
// Solidity: function rollupIDToRollupDataV2Deserialized(uint32 rollupID) view returns(address rollupContract, uint64 chainID, address verifier, uint64 forkID, bytes32 lastLocalExitRoot, uint64 lastBatchSequenced, uint64 lastVerifiedBatch, uint64 lastVerifiedBatchBeforeUpgrade, uint64 rollupTypeID, uint8 rollupVerifierType, bytes32 lastPessimisticRoot, bytes32 programVKey)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupIDToRollupDataV2Deserialized(rollupID uint32) (struct {
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
	return _Agglayermanagermock.Contract.RollupIDToRollupDataV2Deserialized(&_Agglayermanagermock.CallOpts, rollupID)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupTypeCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupTypeCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanagermock.Contract.RollupTypeCount(&_Agglayermanagermock.CallOpts)
}

// RollupTypeCount is a free data retrieval call binding the contract method 0x1796a1ae.
//
// Solidity: function rollupTypeCount() view returns(uint32)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupTypeCount() (uint32, error) {
	return _Agglayermanagermock.Contract.RollupTypeCount(&_Agglayermanagermock.CallOpts)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanagermock *AgglayermanagermockCaller) RollupTypeMap(opts *bind.CallOpts, rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "rollupTypeMap", rollupTypeID)

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
func (_Agglayermanagermock *AgglayermanagermockSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanagermock.Contract.RollupTypeMap(&_Agglayermanagermock.CallOpts, rollupTypeID)
}

// RollupTypeMap is a free data retrieval call binding the contract method 0x65c0504d.
//
// Solidity: function rollupTypeMap(uint32 rollupTypeID) view returns(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bool obsolete, bytes32 genesis, bytes32 programVKey)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) RollupTypeMap(rollupTypeID uint32) (struct {
	ConsensusImplementation common.Address
	Verifier                common.Address
	ForkID                  uint64
	RollupVerifierType      uint8
	Obsolete                bool
	Genesis                 [32]byte
	ProgramVKey             [32]byte
}, error) {
	return _Agglayermanagermock.Contract.RollupTypeMap(&_Agglayermanagermock.CallOpts, rollupTypeID)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCaller) TotalSequencedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "totalSequencedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanagermock.Contract.TotalSequencedBatches(&_Agglayermanagermock.CallOpts)
}

// TotalSequencedBatches is a free data retrieval call binding the contract method 0x066ec012.
//
// Solidity: function totalSequencedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) TotalSequencedBatches() (uint64, error) {
	return _Agglayermanagermock.Contract.TotalSequencedBatches(&_Agglayermanagermock.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCaller) TotalVerifiedBatches(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "totalVerifiedBatches")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanagermock.Contract.TotalVerifiedBatches(&_Agglayermanagermock.CallOpts)
}

// TotalVerifiedBatches is a free data retrieval call binding the contract method 0xdde0ff77.
//
// Solidity: function totalVerifiedBatches() view returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) TotalVerifiedBatches() (uint64, error) {
	return _Agglayermanagermock.Contract.TotalVerifiedBatches(&_Agglayermanagermock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagermock *AgglayermanagermockCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Agglayermanagermock.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagermock *AgglayermanagermockSession) Version() (string, error) {
	return _Agglayermanagermock.Contract.Version(&_Agglayermanagermock.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Agglayermanagermock *AgglayermanagermockCallerSession) Version() (string, error) {
	return _Agglayermanagermock.Contract.Version(&_Agglayermanagermock.CallOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.ActivateEmergencyState(&_Agglayermanagermock.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.ActivateEmergencyState(&_Agglayermanagermock.TransactOpts)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) AddExistingRollup(opts *bind.TransactOpts, rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "addExistingRollup", rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AddExistingRollup(&_Agglayermanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddExistingRollup is a paid mutator transaction binding the contract method 0xe80e5030.
//
// Solidity: function addExistingRollup(address rollupAddress, address verifier, uint64 forkID, uint64 chainID, bytes32 initRoot, uint8 rollupVerifierType, bytes32 programVKey, bytes32 initPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) AddExistingRollup(rollupAddress common.Address, verifier common.Address, forkID uint64, chainID uint64, initRoot [32]byte, rollupVerifierType uint8, programVKey [32]byte, initPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AddExistingRollup(&_Agglayermanagermock.TransactOpts, rollupAddress, verifier, forkID, chainID, initRoot, rollupVerifierType, programVKey, initPessimisticRoot)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) AddNewRollupType(opts *bind.TransactOpts, consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "addNewRollupType", consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AddNewRollupType(&_Agglayermanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AddNewRollupType is a paid mutator transaction binding the contract method 0xabcb5198.
//
// Solidity: function addNewRollupType(address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) AddNewRollupType(consensusImplementation common.Address, verifier common.Address, forkID uint64, rollupVerifierType uint8, genesis [32]byte, description string, programVKey [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AddNewRollupType(&_Agglayermanagermock.TransactOpts, consensusImplementation, verifier, forkID, rollupVerifierType, genesis, description, programVKey)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) AttachAggchainToAL(opts *bind.TransactOpts, rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "attachAggchainToAL", rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AttachAggchainToAL(&_Agglayermanagermock.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// AttachAggchainToAL is a paid mutator transaction binding the contract method 0x97d289a3.
//
// Solidity: function attachAggchainToAL(uint32 rollupTypeID, uint64 chainID, bytes initializeBytesAggchain) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) AttachAggchainToAL(rollupTypeID uint32, chainID uint64, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.AttachAggchainToAL(&_Agglayermanagermock.TransactOpts, rollupTypeID, chainID, initializeBytesAggchain)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.DeactivateEmergencyState(&_Agglayermanagermock.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.DeactivateEmergencyState(&_Agglayermanagermock.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.GrantRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.GrantRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) InitMigration(opts *bind.TransactOpts, rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "initMigration", rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.InitMigration(&_Agglayermanagermock.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// InitMigration is a paid mutator transaction binding the contract method 0xe764a373.
//
// Solidity: function initMigration(uint32 rollupID, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) InitMigration(rollupID uint32, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.InitMigration(&_Agglayermanagermock.TransactOpts, rollupID, newRollupTypeID, upgradeData)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagermock *AgglayermanagermockSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.Initialize(&_Agglayermanagermock.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) Initialize() (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.Initialize(&_Agglayermanagermock.TransactOpts)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) InitializeMock(opts *bind.TransactOpts, trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "initializeMock", trustedAggregator, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) InitializeMock(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.InitializeMock(&_Agglayermanagermock.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// InitializeMock is a paid mutator transaction binding the contract method 0x6c0a51b7.
//
// Solidity: function initializeMock(address trustedAggregator, address admin, address timelock, address emergencyCouncil) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) InitializeMock(trustedAggregator common.Address, admin common.Address, timelock common.Address, emergencyCouncil common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.InitializeMock(&_Agglayermanagermock.TransactOpts, trustedAggregator, admin, timelock, emergencyCouncil)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) ObsoleteRollupType(opts *bind.TransactOpts, rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "obsoleteRollupType", rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.ObsoleteRollupType(&_Agglayermanagermock.TransactOpts, rollupTypeID)
}

// ObsoleteRollupType is a paid mutator transaction binding the contract method 0x7222020f.
//
// Solidity: function obsoleteRollupType(uint32 rollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) ObsoleteRollupType(rollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.ObsoleteRollupType(&_Agglayermanagermock.TransactOpts, rollupTypeID)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockTransactor) OnSequenceBatches(opts *bind.TransactOpts, newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "onSequenceBatches", newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.OnSequenceBatches(&_Agglayermanagermock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// OnSequenceBatches is a paid mutator transaction binding the contract method 0x9a908e73.
//
// Solidity: function onSequenceBatches(uint64 newSequencedBatches, bytes32 newAccInputHash) returns(uint64)
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) OnSequenceBatches(newSequencedBatches uint64, newAccInputHash [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.OnSequenceBatches(&_Agglayermanagermock.TransactOpts, newSequencedBatches, newAccInputHash)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) PrepareMockCalculateRoot(opts *bind.TransactOpts, localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "prepareMockCalculateRoot", localExitRoots)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) PrepareMockCalculateRoot(localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.PrepareMockCalculateRoot(&_Agglayermanagermock.TransactOpts, localExitRoots)
}

// PrepareMockCalculateRoot is a paid mutator transaction binding the contract method 0x8f698ec5.
//
// Solidity: function prepareMockCalculateRoot(bytes32[] localExitRoots) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) PrepareMockCalculateRoot(localExitRoots [][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.PrepareMockCalculateRoot(&_Agglayermanagermock.TransactOpts, localExitRoots)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RenounceRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RenounceRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RevokeRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RevokeRole(&_Agglayermanagermock.TransactOpts, role, account)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) RollbackBatches(opts *bind.TransactOpts, rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "rollbackBatches", rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RollbackBatches(&_Agglayermanagermock.TransactOpts, rollupContract, targetBatch)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x8fd88cc2.
//
// Solidity: function rollbackBatches(address rollupContract, uint64 targetBatch) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) RollbackBatches(rollupContract common.Address, targetBatch uint64) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.RollbackBatches(&_Agglayermanagermock.TransactOpts, rollupContract, targetBatch)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) SetBatchFee(opts *bind.TransactOpts, newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "setBatchFee", newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.SetBatchFee(&_Agglayermanagermock.TransactOpts, newBatchFee)
}

// SetBatchFee is a paid mutator transaction binding the contract method 0xd5073f6f.
//
// Solidity: function setBatchFee(uint256 newBatchFee) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) SetBatchFee(newBatchFee *big.Int) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.SetBatchFee(&_Agglayermanagermock.TransactOpts, newBatchFee)
}

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) SetRollupData(opts *bind.TransactOpts, rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "setRollupData", rollupID, lastLocalExitRoot, lastPessimisticRoot)
}

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) SetRollupData(rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.SetRollupData(&_Agglayermanagermock.TransactOpts, rollupID, lastLocalExitRoot, lastPessimisticRoot)
}

// SetRollupData is a paid mutator transaction binding the contract method 0x680658a1.
//
// Solidity: function setRollupData(uint32 rollupID, bytes32 lastLocalExitRoot, bytes32 lastPessimisticRoot) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) SetRollupData(rollupID uint32, lastLocalExitRoot [32]byte, lastPessimisticRoot [32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.SetRollupData(&_Agglayermanagermock.TransactOpts, rollupID, lastLocalExitRoot, lastPessimisticRoot)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) UpdateRollup(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "updateRollup", rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.UpdateRollup(&_Agglayermanagermock.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollup is a paid mutator transaction binding the contract method 0xc4c928c2.
//
// Solidity: function updateRollup(address rollupContract, uint32 newRollupTypeID, bytes upgradeData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) UpdateRollup(rollupContract common.Address, newRollupTypeID uint32, upgradeData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.UpdateRollup(&_Agglayermanagermock.TransactOpts, rollupContract, newRollupTypeID, upgradeData)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) UpdateRollupByRollupAdmin(opts *bind.TransactOpts, rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "updateRollupByRollupAdmin", rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.UpdateRollupByRollupAdmin(&_Agglayermanagermock.TransactOpts, rollupContract, newRollupTypeID)
}

// UpdateRollupByRollupAdmin is a paid mutator transaction binding the contract method 0xdfdb8c5e.
//
// Solidity: function updateRollupByRollupAdmin(address rollupContract, uint32 newRollupTypeID) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) UpdateRollupByRollupAdmin(rollupContract common.Address, newRollupTypeID uint32) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.UpdateRollupByRollupAdmin(&_Agglayermanagermock.TransactOpts, rollupContract, newRollupTypeID)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) VerifyBatchesTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "verifyBatchesTrustedAggregator", rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyBatchesTrustedAggregator is a paid mutator transaction binding the contract method 0x1489ed10.
//
// Solidity: function verifyBatchesTrustedAggregator(uint32 rollupID, uint64 pendingStateNum, uint64 initNumBatch, uint64 finalNewBatch, bytes32 newLocalExitRoot, bytes32 newStateRoot, address beneficiary, bytes32[24] proof) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) VerifyBatchesTrustedAggregator(rollupID uint32, pendingStateNum uint64, initNumBatch uint64, finalNewBatch uint64, newLocalExitRoot [32]byte, newStateRoot [32]byte, beneficiary common.Address, proof [24][32]byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.VerifyBatchesTrustedAggregator(&_Agglayermanagermock.TransactOpts, rollupID, pendingStateNum, initNumBatch, finalNewBatch, newLocalExitRoot, newStateRoot, beneficiary, proof)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactor) VerifyPessimisticTrustedAggregator(opts *bind.TransactOpts, rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.contract.Transact(opts, "verifyPessimisticTrustedAggregator", rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagermock *AgglayermanagermockSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanagermock.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// VerifyPessimisticTrustedAggregator is a paid mutator transaction binding the contract method 0x6c766877.
//
// Solidity: function verifyPessimisticTrustedAggregator(uint32 rollupID, uint32 l1InfoTreeLeafCount, bytes32 newLocalExitRoot, bytes32 newPessimisticRoot, bytes proof, bytes aggchainData) returns()
func (_Agglayermanagermock *AgglayermanagermockTransactorSession) VerifyPessimisticTrustedAggregator(rollupID uint32, l1InfoTreeLeafCount uint32, newLocalExitRoot [32]byte, newPessimisticRoot [32]byte, proof []byte, aggchainData []byte) (*types.Transaction, error) {
	return _Agglayermanagermock.Contract.VerifyPessimisticTrustedAggregator(&_Agglayermanagermock.TransactOpts, rollupID, l1InfoTreeLeafCount, newLocalExitRoot, newPessimisticRoot, proof, aggchainData)
}

// AgglayermanagermockAddExistingRollupIterator is returned from FilterAddExistingRollup and is used to iterate over the raw logs and unpacked data for AddExistingRollup events raised by the Agglayermanagermock contract.
type AgglayermanagermockAddExistingRollupIterator struct {
	Event *AgglayermanagermockAddExistingRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockAddExistingRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockAddExistingRollup)
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
		it.Event = new(AgglayermanagermockAddExistingRollup)
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
func (it *AgglayermanagermockAddExistingRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockAddExistingRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockAddExistingRollup represents a AddExistingRollup event raised by the Agglayermanagermock contract.
type AgglayermanagermockAddExistingRollup struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterAddExistingRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockAddExistingRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockAddExistingRollupIterator{contract: _Agglayermanagermock.contract, event: "AddExistingRollup", logs: logs, sub: sub}, nil
}

// WatchAddExistingRollup is a free log subscription operation binding the contract event 0x4da47f6e9bbd9ef91887183a576aaebcf1b9bb7d2a567b33b075044c6d36082e.
//
// Solidity: event AddExistingRollup(uint32 indexed rollupID, uint64 forkID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, uint64 lastVerifiedBatchBeforeUpgrade, bytes32 programVKey, bytes32 initPessimisticRoot)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchAddExistingRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockAddExistingRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "AddExistingRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockAddExistingRollup)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseAddExistingRollup(log types.Log) (*AgglayermanagermockAddExistingRollup, error) {
	event := new(AgglayermanagermockAddExistingRollup)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "AddExistingRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockAddNewRollupTypeIterator is returned from FilterAddNewRollupType and is used to iterate over the raw logs and unpacked data for AddNewRollupType events raised by the Agglayermanagermock contract.
type AgglayermanagermockAddNewRollupTypeIterator struct {
	Event *AgglayermanagermockAddNewRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockAddNewRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockAddNewRollupType)
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
		it.Event = new(AgglayermanagermockAddNewRollupType)
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
func (it *AgglayermanagermockAddNewRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockAddNewRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockAddNewRollupType represents a AddNewRollupType event raised by the Agglayermanagermock contract.
type AgglayermanagermockAddNewRollupType struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterAddNewRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagermockAddNewRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockAddNewRollupTypeIterator{contract: _Agglayermanagermock.contract, event: "AddNewRollupType", logs: logs, sub: sub}, nil
}

// WatchAddNewRollupType is a free log subscription operation binding the contract event 0x9eaf2ecbddb14889c9e141a63175c55ac25e0cd7cdea312cdfbd0397976b383a.
//
// Solidity: event AddNewRollupType(uint32 indexed rollupTypeID, address consensusImplementation, address verifier, uint64 forkID, uint8 rollupVerifierType, bytes32 genesis, string description, bytes32 programVKey)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchAddNewRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockAddNewRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "AddNewRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockAddNewRollupType)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseAddNewRollupType(log types.Log) (*AgglayermanagermockAddNewRollupType, error) {
	event := new(AgglayermanagermockAddNewRollupType)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "AddNewRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockCompletedMigrationIterator is returned from FilterCompletedMigration and is used to iterate over the raw logs and unpacked data for CompletedMigration events raised by the Agglayermanagermock contract.
type AgglayermanagermockCompletedMigrationIterator struct {
	Event *AgglayermanagermockCompletedMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockCompletedMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockCompletedMigration)
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
		it.Event = new(AgglayermanagermockCompletedMigration)
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
func (it *AgglayermanagermockCompletedMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockCompletedMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockCompletedMigration represents a CompletedMigration event raised by the Agglayermanagermock contract.
type AgglayermanagermockCompletedMigration struct {
	RollupID uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompletedMigration is a free log retrieval operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterCompletedMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockCompletedMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockCompletedMigrationIterator{contract: _Agglayermanagermock.contract, event: "CompletedMigration", logs: logs, sub: sub}, nil
}

// WatchCompletedMigration is a free log subscription operation binding the contract event 0x6f5e400d25cb6bdafe9f941c2ed83a700da8e0da29dfe15ad4b7ed56e6dd151e.
//
// Solidity: event CompletedMigration(uint32 indexed rollupID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchCompletedMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockCompletedMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "CompletedMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockCompletedMigration)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseCompletedMigration(log types.Log) (*AgglayermanagermockCompletedMigration, error) {
	event := new(AgglayermanagermockCompletedMigration)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "CompletedMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockCreateNewAggchainIterator is returned from FilterCreateNewAggchain and is used to iterate over the raw logs and unpacked data for CreateNewAggchain events raised by the Agglayermanagermock contract.
type AgglayermanagermockCreateNewAggchainIterator struct {
	Event *AgglayermanagermockCreateNewAggchain // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockCreateNewAggchainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockCreateNewAggchain)
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
		it.Event = new(AgglayermanagermockCreateNewAggchain)
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
func (it *AgglayermanagermockCreateNewAggchainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockCreateNewAggchainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockCreateNewAggchain represents a CreateNewAggchain event raised by the Agglayermanagermock contract.
type AgglayermanagermockCreateNewAggchain struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterCreateNewAggchain(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockCreateNewAggchainIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockCreateNewAggchainIterator{contract: _Agglayermanagermock.contract, event: "CreateNewAggchain", logs: logs, sub: sub}, nil
}

// WatchCreateNewAggchain is a free log subscription operation binding the contract event 0x144e3f9b5c63682a3bb7e9ad31e99c043890d3d540cd79dcebc3b5bdfba94c9b.
//
// Solidity: event CreateNewAggchain(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, uint8 rollupVerifierType, bytes initializeBytesAggchain)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchCreateNewAggchain(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockCreateNewAggchain, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "CreateNewAggchain", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockCreateNewAggchain)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseCreateNewAggchain(log types.Log) (*AgglayermanagermockCreateNewAggchain, error) {
	event := new(AgglayermanagermockCreateNewAggchain)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "CreateNewAggchain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockCreateNewRollupIterator is returned from FilterCreateNewRollup and is used to iterate over the raw logs and unpacked data for CreateNewRollup events raised by the Agglayermanagermock contract.
type AgglayermanagermockCreateNewRollupIterator struct {
	Event *AgglayermanagermockCreateNewRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockCreateNewRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockCreateNewRollup)
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
		it.Event = new(AgglayermanagermockCreateNewRollup)
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
func (it *AgglayermanagermockCreateNewRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockCreateNewRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockCreateNewRollup represents a CreateNewRollup event raised by the Agglayermanagermock contract.
type AgglayermanagermockCreateNewRollup struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterCreateNewRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockCreateNewRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockCreateNewRollupIterator{contract: _Agglayermanagermock.contract, event: "CreateNewRollup", logs: logs, sub: sub}, nil
}

// WatchCreateNewRollup is a free log subscription operation binding the contract event 0x194c983456df6701c6a50830b90fe80e72b823411d0d524970c9590dc277a641.
//
// Solidity: event CreateNewRollup(uint32 indexed rollupID, uint32 rollupTypeID, address rollupAddress, uint64 chainID, address gasTokenAddress)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchCreateNewRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockCreateNewRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "CreateNewRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockCreateNewRollup)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseCreateNewRollup(log types.Log) (*AgglayermanagermockCreateNewRollup, error) {
	event := new(AgglayermanagermockCreateNewRollup)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "CreateNewRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Agglayermanagermock contract.
type AgglayermanagermockEmergencyStateActivatedIterator struct {
	Event *AgglayermanagermockEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockEmergencyStateActivated)
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
		it.Event = new(AgglayermanagermockEmergencyStateActivated)
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
func (it *AgglayermanagermockEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockEmergencyStateActivated represents a EmergencyStateActivated event raised by the Agglayermanagermock contract.
type AgglayermanagermockEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*AgglayermanagermockEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockEmergencyStateActivatedIterator{contract: _Agglayermanagermock.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockEmergencyStateActivated)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseEmergencyStateActivated(log types.Log) (*AgglayermanagermockEmergencyStateActivated, error) {
	event := new(AgglayermanagermockEmergencyStateActivated)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Agglayermanagermock contract.
type AgglayermanagermockEmergencyStateDeactivatedIterator struct {
	Event *AgglayermanagermockEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockEmergencyStateDeactivated)
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
		it.Event = new(AgglayermanagermockEmergencyStateDeactivated)
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
func (it *AgglayermanagermockEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Agglayermanagermock contract.
type AgglayermanagermockEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*AgglayermanagermockEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockEmergencyStateDeactivatedIterator{contract: _Agglayermanagermock.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockEmergencyStateDeactivated)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseEmergencyStateDeactivated(log types.Log) (*AgglayermanagermockEmergencyStateDeactivated, error) {
	event := new(AgglayermanagermockEmergencyStateDeactivated)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockInitMigrationIterator is returned from FilterInitMigration and is used to iterate over the raw logs and unpacked data for InitMigration events raised by the Agglayermanagermock contract.
type AgglayermanagermockInitMigrationIterator struct {
	Event *AgglayermanagermockInitMigration // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockInitMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockInitMigration)
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
		it.Event = new(AgglayermanagermockInitMigration)
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
func (it *AgglayermanagermockInitMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockInitMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockInitMigration represents a InitMigration event raised by the Agglayermanagermock contract.
type AgglayermanagermockInitMigration struct {
	RollupID        uint32
	NewRollupTypeID uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitMigration is a free log retrieval operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterInitMigration(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockInitMigrationIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockInitMigrationIterator{contract: _Agglayermanagermock.contract, event: "InitMigration", logs: logs, sub: sub}, nil
}

// WatchInitMigration is a free log subscription operation binding the contract event 0x3fb14e3ae056c8bb24b0f03d618f2aada703672121f643107dac4783669a9de9.
//
// Solidity: event InitMigration(uint32 indexed rollupID, uint32 newRollupTypeID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchInitMigration(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockInitMigration, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "InitMigration", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockInitMigration)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "InitMigration", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseInitMigration(log types.Log) (*AgglayermanagermockInitMigration, error) {
	event := new(AgglayermanagermockInitMigration)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "InitMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Agglayermanagermock contract.
type AgglayermanagermockInitializedIterator struct {
	Event *AgglayermanagermockInitialized // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockInitialized)
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
		it.Event = new(AgglayermanagermockInitialized)
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
func (it *AgglayermanagermockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockInitialized represents a Initialized event raised by the Agglayermanagermock contract.
type AgglayermanagermockInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgglayermanagermockInitializedIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockInitializedIterator{contract: _Agglayermanagermock.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockInitialized) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockInitialized)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseInitialized(log types.Log) (*AgglayermanagermockInitialized, error) {
	event := new(AgglayermanagermockInitialized)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockObsoleteRollupTypeIterator is returned from FilterObsoleteRollupType and is used to iterate over the raw logs and unpacked data for ObsoleteRollupType events raised by the Agglayermanagermock contract.
type AgglayermanagermockObsoleteRollupTypeIterator struct {
	Event *AgglayermanagermockObsoleteRollupType // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockObsoleteRollupTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockObsoleteRollupType)
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
		it.Event = new(AgglayermanagermockObsoleteRollupType)
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
func (it *AgglayermanagermockObsoleteRollupTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockObsoleteRollupTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockObsoleteRollupType represents a ObsoleteRollupType event raised by the Agglayermanagermock contract.
type AgglayermanagermockObsoleteRollupType struct {
	RollupTypeID uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterObsoleteRollupType is a free log retrieval operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterObsoleteRollupType(opts *bind.FilterOpts, rollupTypeID []uint32) (*AgglayermanagermockObsoleteRollupTypeIterator, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockObsoleteRollupTypeIterator{contract: _Agglayermanagermock.contract, event: "ObsoleteRollupType", logs: logs, sub: sub}, nil
}

// WatchObsoleteRollupType is a free log subscription operation binding the contract event 0x4710d2ee567ef1ed6eb2f651dde4589524bcf7cebc62147a99b281cc836e7e44.
//
// Solidity: event ObsoleteRollupType(uint32 indexed rollupTypeID)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchObsoleteRollupType(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockObsoleteRollupType, rollupTypeID []uint32) (event.Subscription, error) {

	var rollupTypeIDRule []interface{}
	for _, rollupTypeIDItem := range rollupTypeID {
		rollupTypeIDRule = append(rollupTypeIDRule, rollupTypeIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "ObsoleteRollupType", rollupTypeIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockObsoleteRollupType)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseObsoleteRollupType(log types.Log) (*AgglayermanagermockObsoleteRollupType, error) {
	event := new(AgglayermanagermockObsoleteRollupType)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "ObsoleteRollupType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockOnSequenceBatchesIterator is returned from FilterOnSequenceBatches and is used to iterate over the raw logs and unpacked data for OnSequenceBatches events raised by the Agglayermanagermock contract.
type AgglayermanagermockOnSequenceBatchesIterator struct {
	Event *AgglayermanagermockOnSequenceBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockOnSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockOnSequenceBatches)
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
		it.Event = new(AgglayermanagermockOnSequenceBatches)
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
func (it *AgglayermanagermockOnSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockOnSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockOnSequenceBatches represents a OnSequenceBatches event raised by the Agglayermanagermock contract.
type AgglayermanagermockOnSequenceBatches struct {
	RollupID           uint32
	LastBatchSequenced uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOnSequenceBatches is a free log retrieval operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterOnSequenceBatches(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockOnSequenceBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockOnSequenceBatchesIterator{contract: _Agglayermanagermock.contract, event: "OnSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchOnSequenceBatches is a free log subscription operation binding the contract event 0x1d9f30260051d51d70339da239ea7b080021adcaabfa71c9b0ea339a20cf9a25.
//
// Solidity: event OnSequenceBatches(uint32 indexed rollupID, uint64 lastBatchSequenced)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchOnSequenceBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockOnSequenceBatches, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "OnSequenceBatches", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockOnSequenceBatches)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseOnSequenceBatches(log types.Log) (*AgglayermanagermockOnSequenceBatches, error) {
	event := new(AgglayermanagermockOnSequenceBatches)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "OnSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleAdminChangedIterator struct {
	Event *AgglayermanagermockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockRoleAdminChanged)
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
		it.Event = new(AgglayermanagermockRoleAdminChanged)
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
func (it *AgglayermanagermockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockRoleAdminChanged represents a RoleAdminChanged event raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AgglayermanagermockRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockRoleAdminChangedIterator{contract: _Agglayermanagermock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockRoleAdminChanged)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseRoleAdminChanged(log types.Log) (*AgglayermanagermockRoleAdminChanged, error) {
	event := new(AgglayermanagermockRoleAdminChanged)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleGrantedIterator struct {
	Event *AgglayermanagermockRoleGranted // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockRoleGranted)
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
		it.Event = new(AgglayermanagermockRoleGranted)
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
func (it *AgglayermanagermockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockRoleGranted represents a RoleGranted event raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagermockRoleGrantedIterator, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockRoleGrantedIterator{contract: _Agglayermanagermock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockRoleGranted)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseRoleGranted(log types.Log) (*AgglayermanagermockRoleGranted, error) {
	event := new(AgglayermanagermockRoleGranted)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleRevokedIterator struct {
	Event *AgglayermanagermockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockRoleRevoked)
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
		it.Event = new(AgglayermanagermockRoleRevoked)
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
func (it *AgglayermanagermockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockRoleRevoked represents a RoleRevoked event raised by the Agglayermanagermock contract.
type AgglayermanagermockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AgglayermanagermockRoleRevokedIterator, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockRoleRevokedIterator{contract: _Agglayermanagermock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockRoleRevoked)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseRoleRevoked(log types.Log) (*AgglayermanagermockRoleRevoked, error) {
	event := new(AgglayermanagermockRoleRevoked)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Agglayermanagermock contract.
type AgglayermanagermockRollbackBatchesIterator struct {
	Event *AgglayermanagermockRollbackBatches // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockRollbackBatches)
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
		it.Event = new(AgglayermanagermockRollbackBatches)
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
func (it *AgglayermanagermockRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockRollbackBatches represents a RollbackBatches event raised by the Agglayermanagermock contract.
type AgglayermanagermockRollbackBatches struct {
	RollupID               uint32
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterRollbackBatches(opts *bind.FilterOpts, rollupID []uint32, targetBatch []uint64) (*AgglayermanagermockRollbackBatchesIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockRollbackBatchesIterator{contract: _Agglayermanagermock.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x80a6d395a55aed8126079cb8247f0a6848b1440ca2cdca3b4386f250c3529402.
//
// Solidity: event RollbackBatches(uint32 indexed rollupID, uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockRollbackBatches, rollupID []uint32, targetBatch []uint64) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}
	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "RollbackBatches", rollupIDRule, targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockRollbackBatches)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseRollbackBatches(log types.Log) (*AgglayermanagermockRollbackBatches, error) {
	event := new(AgglayermanagermockRollbackBatches)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockSetBatchFeeIterator is returned from FilterSetBatchFee and is used to iterate over the raw logs and unpacked data for SetBatchFee events raised by the Agglayermanagermock contract.
type AgglayermanagermockSetBatchFeeIterator struct {
	Event *AgglayermanagermockSetBatchFee // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockSetBatchFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockSetBatchFee)
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
		it.Event = new(AgglayermanagermockSetBatchFee)
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
func (it *AgglayermanagermockSetBatchFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockSetBatchFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockSetBatchFee represents a SetBatchFee event raised by the Agglayermanagermock contract.
type AgglayermanagermockSetBatchFee struct {
	NewBatchFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetBatchFee is a free log retrieval operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterSetBatchFee(opts *bind.FilterOpts) (*AgglayermanagermockSetBatchFeeIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockSetBatchFeeIterator{contract: _Agglayermanagermock.contract, event: "SetBatchFee", logs: logs, sub: sub}, nil
}

// WatchSetBatchFee is a free log subscription operation binding the contract event 0xfb383653f53ee079978d0c9aff7aeff04a10166ce244cca9c9f9d8d96bed45b2.
//
// Solidity: event SetBatchFee(uint256 newBatchFee)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchSetBatchFee(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockSetBatchFee) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "SetBatchFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockSetBatchFee)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseSetBatchFee(log types.Log) (*AgglayermanagermockSetBatchFee, error) {
	event := new(AgglayermanagermockSetBatchFee)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "SetBatchFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockSetTrustedAggregatorIterator is returned from FilterSetTrustedAggregator and is used to iterate over the raw logs and unpacked data for SetTrustedAggregator events raised by the Agglayermanagermock contract.
type AgglayermanagermockSetTrustedAggregatorIterator struct {
	Event *AgglayermanagermockSetTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockSetTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockSetTrustedAggregator)
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
		it.Event = new(AgglayermanagermockSetTrustedAggregator)
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
func (it *AgglayermanagermockSetTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockSetTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockSetTrustedAggregator represents a SetTrustedAggregator event raised by the Agglayermanagermock contract.
type AgglayermanagermockSetTrustedAggregator struct {
	NewTrustedAggregator common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedAggregator is a free log retrieval operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterSetTrustedAggregator(opts *bind.FilterOpts) (*AgglayermanagermockSetTrustedAggregatorIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockSetTrustedAggregatorIterator{contract: _Agglayermanagermock.contract, event: "SetTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchSetTrustedAggregator is a free log subscription operation binding the contract event 0x61f8fec29495a3078e9271456f05fb0707fd4e41f7661865f80fc437d06681ca.
//
// Solidity: event SetTrustedAggregator(address newTrustedAggregator)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchSetTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockSetTrustedAggregator) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "SetTrustedAggregator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockSetTrustedAggregator)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseSetTrustedAggregator(log types.Log) (*AgglayermanagermockSetTrustedAggregator, error) {
	event := new(AgglayermanagermockSetTrustedAggregator)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "SetTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockUpdateRollupIterator is returned from FilterUpdateRollup and is used to iterate over the raw logs and unpacked data for UpdateRollup events raised by the Agglayermanagermock contract.
type AgglayermanagermockUpdateRollupIterator struct {
	Event *AgglayermanagermockUpdateRollup // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockUpdateRollupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockUpdateRollup)
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
		it.Event = new(AgglayermanagermockUpdateRollup)
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
func (it *AgglayermanagermockUpdateRollupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockUpdateRollupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockUpdateRollup represents a UpdateRollup event raised by the Agglayermanagermock contract.
type AgglayermanagermockUpdateRollup struct {
	RollupID                       uint32
	NewRollupTypeID                uint32
	LastVerifiedBatchBeforeUpgrade uint64
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollup is a free log retrieval operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterUpdateRollup(opts *bind.FilterOpts, rollupID []uint32) (*AgglayermanagermockUpdateRollupIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockUpdateRollupIterator{contract: _Agglayermanagermock.contract, event: "UpdateRollup", logs: logs, sub: sub}, nil
}

// WatchUpdateRollup is a free log subscription operation binding the contract event 0xf585e04c05d396901170247783d3e5f0ee9c1df23072985b50af089f5e48b19d.
//
// Solidity: event UpdateRollup(uint32 indexed rollupID, uint32 newRollupTypeID, uint64 lastVerifiedBatchBeforeUpgrade)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchUpdateRollup(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockUpdateRollup, rollupID []uint32) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "UpdateRollup", rollupIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockUpdateRollup)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseUpdateRollup(log types.Log) (*AgglayermanagermockUpdateRollup, error) {
	event := new(AgglayermanagermockUpdateRollup)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "UpdateRollup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockUpdateRollupManagerVersionIterator is returned from FilterUpdateRollupManagerVersion and is used to iterate over the raw logs and unpacked data for UpdateRollupManagerVersion events raised by the Agglayermanagermock contract.
type AgglayermanagermockUpdateRollupManagerVersionIterator struct {
	Event *AgglayermanagermockUpdateRollupManagerVersion // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockUpdateRollupManagerVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockUpdateRollupManagerVersion)
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
		it.Event = new(AgglayermanagermockUpdateRollupManagerVersion)
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
func (it *AgglayermanagermockUpdateRollupManagerVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockUpdateRollupManagerVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockUpdateRollupManagerVersion represents a UpdateRollupManagerVersion event raised by the Agglayermanagermock contract.
type AgglayermanagermockUpdateRollupManagerVersion struct {
	RollupManagerVersion string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupManagerVersion is a free log retrieval operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterUpdateRollupManagerVersion(opts *bind.FilterOpts) (*AgglayermanagermockUpdateRollupManagerVersionIterator, error) {

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockUpdateRollupManagerVersionIterator{contract: _Agglayermanagermock.contract, event: "UpdateRollupManagerVersion", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupManagerVersion is a free log subscription operation binding the contract event 0x50cadc0c001f05dd4b81db1e92b98d77e718fd2f103fb7b77293e867d329a4c2.
//
// Solidity: event UpdateRollupManagerVersion(string rollupManagerVersion)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchUpdateRollupManagerVersion(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockUpdateRollupManagerVersion) (event.Subscription, error) {

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "UpdateRollupManagerVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockUpdateRollupManagerVersion)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseUpdateRollupManagerVersion(log types.Log) (*AgglayermanagermockUpdateRollupManagerVersion, error) {
	event := new(AgglayermanagermockUpdateRollupManagerVersion)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "UpdateRollupManagerVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockVerifyBatchesTrustedAggregatorIterator is returned from FilterVerifyBatchesTrustedAggregator and is used to iterate over the raw logs and unpacked data for VerifyBatchesTrustedAggregator events raised by the Agglayermanagermock contract.
type AgglayermanagermockVerifyBatchesTrustedAggregatorIterator struct {
	Event *AgglayermanagermockVerifyBatchesTrustedAggregator // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockVerifyBatchesTrustedAggregatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockVerifyBatchesTrustedAggregator)
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
		it.Event = new(AgglayermanagermockVerifyBatchesTrustedAggregator)
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
func (it *AgglayermanagermockVerifyBatchesTrustedAggregatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockVerifyBatchesTrustedAggregatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockVerifyBatchesTrustedAggregator represents a VerifyBatchesTrustedAggregator event raised by the Agglayermanagermock contract.
type AgglayermanagermockVerifyBatchesTrustedAggregator struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterVerifyBatchesTrustedAggregator(opts *bind.FilterOpts, rollupID []uint32, aggregator []common.Address) (*AgglayermanagermockVerifyBatchesTrustedAggregatorIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockVerifyBatchesTrustedAggregatorIterator{contract: _Agglayermanagermock.contract, event: "VerifyBatchesTrustedAggregator", logs: logs, sub: sub}, nil
}

// WatchVerifyBatchesTrustedAggregator is a free log subscription operation binding the contract event 0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3.
//
// Solidity: event VerifyBatchesTrustedAggregator(uint32 indexed rollupID, uint64 numBatch, bytes32 stateRoot, bytes32 exitRoot, address indexed aggregator)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchVerifyBatchesTrustedAggregator(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockVerifyBatchesTrustedAggregator, rollupID []uint32, aggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "VerifyBatchesTrustedAggregator", rollupIDRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockVerifyBatchesTrustedAggregator)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseVerifyBatchesTrustedAggregator(log types.Log) (*AgglayermanagermockVerifyBatchesTrustedAggregator, error) {
	event := new(AgglayermanagermockVerifyBatchesTrustedAggregator)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "VerifyBatchesTrustedAggregator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgglayermanagermockVerifyPessimisticStateTransitionIterator is returned from FilterVerifyPessimisticStateTransition and is used to iterate over the raw logs and unpacked data for VerifyPessimisticStateTransition events raised by the Agglayermanagermock contract.
type AgglayermanagermockVerifyPessimisticStateTransitionIterator struct {
	Event *AgglayermanagermockVerifyPessimisticStateTransition // Event containing the contract specifics and raw log

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
func (it *AgglayermanagermockVerifyPessimisticStateTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgglayermanagermockVerifyPessimisticStateTransition)
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
		it.Event = new(AgglayermanagermockVerifyPessimisticStateTransition)
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
func (it *AgglayermanagermockVerifyPessimisticStateTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgglayermanagermockVerifyPessimisticStateTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgglayermanagermockVerifyPessimisticStateTransition represents a VerifyPessimisticStateTransition event raised by the Agglayermanagermock contract.
type AgglayermanagermockVerifyPessimisticStateTransition struct {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) FilterVerifyPessimisticStateTransition(opts *bind.FilterOpts, rollupID []uint32, trustedAggregator []common.Address) (*AgglayermanagermockVerifyPessimisticStateTransitionIterator, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.FilterLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return &AgglayermanagermockVerifyPessimisticStateTransitionIterator{contract: _Agglayermanagermock.contract, event: "VerifyPessimisticStateTransition", logs: logs, sub: sub}, nil
}

// WatchVerifyPessimisticStateTransition is a free log subscription operation binding the contract event 0xdf47e7dbf79874ec576f516c40bc1483f7c8ddf4b45bfd4baff4650f1229a711.
//
// Solidity: event VerifyPessimisticStateTransition(uint32 indexed rollupID, bytes32 prevPessimisticRoot, bytes32 newPessimisticRoot, bytes32 prevLocalExitRoot, bytes32 newLocalExitRoot, bytes32 l1InfoRoot, address indexed trustedAggregator)
func (_Agglayermanagermock *AgglayermanagermockFilterer) WatchVerifyPessimisticStateTransition(opts *bind.WatchOpts, sink chan<- *AgglayermanagermockVerifyPessimisticStateTransition, rollupID []uint32, trustedAggregator []common.Address) (event.Subscription, error) {

	var rollupIDRule []interface{}
	for _, rollupIDItem := range rollupID {
		rollupIDRule = append(rollupIDRule, rollupIDItem)
	}

	var trustedAggregatorRule []interface{}
	for _, trustedAggregatorItem := range trustedAggregator {
		trustedAggregatorRule = append(trustedAggregatorRule, trustedAggregatorItem)
	}

	logs, sub, err := _Agglayermanagermock.contract.WatchLogs(opts, "VerifyPessimisticStateTransition", rollupIDRule, trustedAggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgglayermanagermockVerifyPessimisticStateTransition)
				if err := _Agglayermanagermock.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
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
func (_Agglayermanagermock *AgglayermanagermockFilterer) ParseVerifyPessimisticStateTransition(log types.Log) (*AgglayermanagermockVerifyPessimisticStateTransition, error) {
	event := new(AgglayermanagermockVerifyPessimisticStateTransition)
	if err := _Agglayermanagermock.contract.UnpackLog(event, "VerifyPessimisticStateTransition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

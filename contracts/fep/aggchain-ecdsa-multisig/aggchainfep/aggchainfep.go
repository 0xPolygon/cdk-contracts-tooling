// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggchainfep

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

// AggchainBaseRemoveSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type AggchainBaseRemoveSignerInfo struct {
	Addr  common.Address
	Index *big.Int
}

// AggchainBaseSignerInfo is an auto generated low-level Go binding around an user-defined struct.
type AggchainBaseSignerInfo struct {
	Addr common.Address
	Url  string
}

// AggchainFEPInitParams is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPInitParams struct {
	L2BlockTime           *big.Int
	RollupConfigHash      [32]byte
	StartingOutputRoot    [32]byte
	StartingBlockNumber   *big.Int
	StartingTimestamp     *big.Int
	SubmissionInterval    *big.Int
	OptimisticModeManager common.Address
	AggregationVkey       [32]byte
	RangeVkeyCommitment   [32]byte
}

// AggchainFEPOpSuccinctConfig is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPOpSuccinctConfig struct {
	AggregationVkey     [32]byte
	RangeVkeyCommitment [32]byte
	RollupConfigHash    [32]byte
}

// AggchainFEPOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// AggchainfepMetaData contains all meta data concerning the Aggchainfep contract.
var AggchainfepMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggregationVkeyMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotProposeFutureL2Output\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfigDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAggchainSignersArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockNumberLessThanNextBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockTimeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2OutputRootCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RangeVkeyCommitmentMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupConfigHashMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartL2TimestampMustBeLessThanCurrentTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmissionIntervalMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHighAfterRemoval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"AcceptOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldAggregationVkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAggregationVkey\",\"type\":\"bytes32\"}],\"name\":\"AggregationVkeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configName\",\"type\":\"bytes32\"}],\"name\":\"OpSuccinctConfigDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configName\",\"type\":\"bytes32\"}],\"name\":\"OpSuccinctConfigSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"configName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"OpSuccinctConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRangeVkeyCommitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"RangeVkeyCommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRollupConfigHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfigHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSubmissionInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubmissionInterval\",\"type\":\"uint256\"}],\"name\":\"SubmissionIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"TransferOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_FEP_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_CONFIG_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacypendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacyvKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_configName\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rollupConfigHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"addOpSuccinctConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggchainSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainSignersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregationVkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_configName\",\"type\":\"bytes32\"}],\"name\":\"deleteOpSuccinctConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainParamsAndVKeySelector\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structAggchainFEP.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startingOutputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"optimisticModeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"internalType\":\"structAggchainFEP.InitParams\",\"name\":\"_initParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultVkeys\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultSigners\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_initOwnedAggchainVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"_initAggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startingOutputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"optimisticModeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"internalType\":\"structAggchainFEP.InitParams\",\"name\":\"_initParams\",\"type\":\"tuple\"}],\"name\":\"initializeFromECDSAMultisig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"startingOutputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"optimisticModeManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"internalType\":\"structAggchainFEP.InitParams\",\"name\":\"_initParams\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultVkeys\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultSigners\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_initOwnedAggchainVKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"_initAggchainVKeySelector\",\"type\":\"bytes4\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"initializeFromPessimisticConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"}],\"internalType\":\"structAggchainFEP.OpSuccinctConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"isValidOpSuccinctConfig\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"opSuccinctConfigs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggregationVkey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rangeVkeyCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupConfigHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOptimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeVkeyCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_configName\",\"type\":\"bytes32\"}],\"name\":\"selectOpSuccinctConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"selectedOpSuccinctConfigName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToURLs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"transferOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structAggchainBase.RemoveSignerInfo[]\",\"name\":\"_signersToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateSignersAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeFromPreviousFEP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultVkeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051615bf4380380615bf483398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e0516101005161594a6102aa5f395f8181610cbf0152818161100c015281816118cd015281816126fa0152818161299b01526139cb01525f818161098d01528181612bfa01528181612f4e01526135b201525f610c9801525f610dbb01525f610e06015261594a5ff3fe608060405234801561000f575f5ffd5b50600436106105e2575f3560e01c80636e05d2cd11610307578063be647d031161019d578063e46761c4116100f3578063f51f563a1161009e578063fc5014d611610079578063fc5014d614610ee6578063fd7d249314610f0b578063fdbbc19b14610f13575f5ffd5b8063f51f563a14610e87578063f72f606d14610e9a578063f851a44014610ec1575f5ffd5b8063ec5b2e3a116100ce578063ec5b2e3a14610e4d578063efe6c9f414610e60578063effb847914610e68575f5ffd5b8063e46761c414610e01578063e7a7ed0214610e28578063e90a340914610e3c575f5ffd5b8063cea5a4c011610153578063d1de856c1161012e578063d1de856c14610ddd578063dcec334814610df0578063e1a41bcf14610df8575f5ffd5b8063cea5a4c014610d8e578063cfa8ed4714610d96578063d02103ca14610db6575f5ffd5b8063c754c7ed11610183578063c754c7ed14610d4b578063c89e42df14610d73578063ca69e7dc14610d86575f5ffd5b8063be647d0314610d3a578063c32e4e3e14610d42575f5ffd5b806393991af31161025d578063a3c573eb11610208578063adb8696c116101e3578063adb8696c14610cf4578063b3a326f714610d14578063bdfbed7e14610d27575f5ffd5b8063a3c573eb14610c93578063ab0475cf14610cba578063ada8f91914610ce1575f5ffd5b80639f78f066116102385780639f78f06614610c28578063a2398fc614610c31578063a25ae55714610c44575f5ffd5b806393991af314610c0457806396a4f54614610c0d5780639ee4afa314610c15575f5ffd5b80637388c436116102bd57806381eb0baf1161029857806381eb0baf14610beb5780638878627214610bf35780638c3d730114610bfc575f5ffd5b80637388c43614610b9857806374f0b0c114610bb85780637df73e2714610bd8575f5ffd5b80636ff512cc116102ed5780636ff512cc14610b6957806370872aa514610b7c5780637125702214610b85575f5ffd5b80636e05d2cd14610b3a5780636e7fbce914610b43575f5ffd5b80633e1e01211161047c57806354fd4d50116103d257806369f16eec1161037d5780636abcf563116103585780636abcf56314610b0a5780636b8616ce14610b125780636d9a1c8b14610b31575f5ffd5b806369f16eec14610aa65780636a55f66c14610aae5780636a56620b14610ac1575f5ffd5b806360caf7a0116103ad57806360caf7a014610a5e57806363de5bc614610a6b578063680196ed14610a7e575f5ffd5b806354fd4d50146109f2578063558716c114610a2b5780635ecaca2b14610a3e575f5ffd5b806349185e0611610432578063527570f11161040d578063527570f1146109c2578063529933df146109e2578063542028d5146109ea575f5ffd5b806349185e061461097557806349b7b8021461098857806352076aca146109af575f5ffd5b8063456052671161046257806345605267146109215780634599c7881461095a57806347c37e9c14610962575f5ffd5b80633e1e01211461090357806342cde4e814610918575f5ffd5b80631d0b435e1161053c578063336c9e81116104e757806337d4d030116104c257806337d4d0301461086a5780633c351e10146108a65780633cbc795b146108c6575f5ffd5b8063336c9e811461083157806335acd6c21461084457806336cd6b5b14610857575f5ffd5b80632b31841e116105175780632b31841e146107f55780632c111c06146107fe578063314eb17b1461081e575f5ffd5b80631d0b435e146106e3578063267822471461076e57806326f9b76d1461078e575f5ffd5b80631489e7071161059c578063188d918011610577578063188d91801461065057806319451a8f146106865780631cf810ee14610699575f5ffd5b80631489e7071461063757806315981b291461063f57806317b7a9f014610647575f5ffd5b80630822dc61116105cc5780630822dc6114610610578063107bf28c1461061a578063126349001461062f575f5ffd5b80622134cc146105e657806301fcf6a0146105fd575b5f5ffd5b6078545b6040519081526020015b60405180910390f35b6105ea61060b366004614b63565b610f26565b610618611090565b005b610622611174565b6040516105f49190614b85565b610618611200565b61061861130f565b61061861141e565b6105ea60455481565b603e54610676907501000000000000000000000000000000000000000000900460ff1681565b60405190151581526020016105f4565b610618610694366004614bd8565b6114f5565b607c546106be90610100900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016105f4565b61073d6106f1366004614c31565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016105f4565b6001546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b6107c461079c366004614b63565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016105f4565b6105ea607a5481565b6008546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b61061861082c366004614bd8565b611656565b61061861083f366004614c62565b61178c565b6106be610852366004614c62565b611857565b610622610865366004614ca5565b61188c565b6106226040518060400160405280600681526020017f76332e302e30000000000000000000000000000000000000000000000000000081525081565b6009546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b6009546108ee9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016105f4565b61090b6118a4565b6040516105f49190614cc0565b6105ea60445481565b6007546109419068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016105f4565b6105ea6119e8565b610618610970366004614d18565b611a57565b610676610983366004614e10565b611d26565b6106be7f000000000000000000000000000000000000000000000000000000000000000081565b6106186109bd366004614c62565b611d4c565b6040546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b6077546105ea565b610622611e46565b60408051808201909152600681527f76332e302e3000000000000000000000000000000000000000000000000000006020820152610622565b610618610a393660046150b7565b611e53565b603d546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b607c546106769060ff1681565b610618610a793660046151dd565b6121a9565b610a91610a8c3660046151f8565b6123e3565b604080519283526020830191909152016105f4565b6105ea6126bf565b6105ea610abc3660046151f8565b6126d0565b610aef610acf366004614c62565b607e6020525f908152604090208054600182015460029092015490919083565b604080519384526020840192909252908201526060016105f4565b6074546105ea565b6105ea610b20366004615245565b60066020525f908152604090205481565b6105ea607b5481565b6105ea60055481565b6107c47e0100000000000000000000000000000000000000000000000000000000000081565b610618610b77366004614ca5565b612835565b6105ea60755481565b610618610b9336600461526c565b6128fe565b603f546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b603e546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b610676610be6366004614ca5565b612930565b610618612a3d565b6105ea60765481565b610618612b25565b6105ea60785481565b610618612bf8565b610618610c233660046151f8565b612f4c565b6105ea607f5481565b610618610c3f366004615320565b613107565b610c57610c52366004614c62565b613455565b60408051825181526020808401516fffffffffffffffffffffffffffffffff9081169183019190915292820151909216908201526060016105f4565b6106be7f000000000000000000000000000000000000000000000000000000000000000081565b6106be7f000000000000000000000000000000000000000000000000000000000000000081565b610618610cef366004614ca5565b6134e7565b607d546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b610618610d22366004614ca5565b6135b0565b610618610d35366004614ca5565b6136e6565b610618613800565b6105ea60795481565b60075461094190700100000000000000000000000000000000900467ffffffffffffffff1681565b610618610d813660046153be565b613911565b6105ea6139a3565b6108ee600181565b6002546106be9073ffffffffffffffffffffffffffffffffffffffff1681565b6106be7f000000000000000000000000000000000000000000000000000000000000000081565b6105ea610deb366004614c62565b613a5d565b6105ea613a86565b6105ea60775481565b6106be7f000000000000000000000000000000000000000000000000000000000000000081565b6007546109419067ffffffffffffffff1681565b6107c4610e4a366004614b63565b90565b610618610e5b366004614c62565b613a9c565b610618613b34565b6105ea610e76366004614b63565b60416020525f908152604090205481565b610618610e953660046153f0565b613c2b565b6105ea7fae8304f40f7123e0c87b97f8a600e94ff3a3a25be588fc66b8a3717c8959ce7781565b5f546106be9062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603e546106769074010000000000000000000000000000000000000000900460ff1681565b610618613c8c565b610618610f21366004614ca5565b613d84565b603e545f9074010000000000000000000000000000000000000000900460ff1615158103610fbc57507fffffffff0000000000000000000000000000000000000000000000000000000081165f9081526041602052604090205480610fb7576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015611066573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061108a91906154e3565b92915050565b607c54610100900473ffffffffffffffffffffffffffffffffffffffff1633146110e6576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c5460ff16611122576040517f873dabd200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac905f90a1565b60048054611181906154fa565b80601f01602080910402602001604051908101604052809291908181526020018280546111ad906154fa565b80156111f85780601f106111cf576101008083540402835291602001916111f8565b820191905f5260205f20905b8154815290600101906020018083116111db57829003601f168201915b505050505081565b607d5473ffffffffffffffffffffffffffffffffffffffff163314611251576040517f93f1e09400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c8054607d805473ffffffffffffffffffffffffffffffffffffffff8082166101009081027fffffffffffffffffffffff0000000000000000000000000000000000000000ff861617958690557fffffffffffffffffffffffff000000000000000000000000000000000000000090921690925560408051938290048316808552919094049091166020830152917f9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f331991015b60405180910390a150565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611360576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff16156113b5576040517f9573504e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829905f90a1565b60405473ffffffffffffffffffffffffffffffffffffffff16331461146f576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101611304565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611546576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8061157d576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f90815260416020526040902054156115e5576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260416020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146116a7576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205461170e576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152604160209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146117dd576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03611816576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60775460408051918252602082018390527fc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2910160405180910390a1607755565b60428181548110611866575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60436020525f908152604090208054611181906154fa565b603e546060907501000000000000000000000000000000000000000000900460ff161561197d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633e1e01216040518163ffffffff1660e01b81526004015f60405180830381865afa158015611933573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611978919081019061554b565b905090565b60428054806020026020016040519081016040528092919081815260200182805480156119de57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116119b3575b5050505050905090565b6074545f9015611a505760748054611a0290600190615612565b81548110611a1257611a12615625565b5f91825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b5060755490565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611aa8576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83611b3a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4c324f75747075744f7261636c653a20636f6e666967206e616d652063616e6e60448201527f6f7420626520656d70747900000000000000000000000000000000000000000060648201526084015b60405180910390fd5b611b7b607e5f8681526020019081526020015f206040518060600160405290815f820154815260200160018201548152602001600282015481525050611d26565b15611c08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4c324f75747075744f7261636c653a20636f6e66696720616c7265616479206560448201527f78697374730000000000000000000000000000000000000000000000000000006064820152608401611b31565b6040805160608101825283815260208101839052908101849052611c2b81611d26565b611cb7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f4c324f75747075744f7261636c653a20696e76616c6964204f5020537563636960448201527f6e637420636f6e66696775726174696f6e20706172616d6574657273000000006064820152608401611b31565b5f858152607e60209081526040918290208351815583820151600182015583830151600290910155815185815290810184905290810185905285907fea0123c726a665cb0ab5691444f929a7056c7a7709c60c0587829e8046b8d5149060600160405180910390a25050505050565b80515f9015801590611d3b5750602082015115155b801561108a57505060400151151590565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611d9d576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611dde607e5f8381526020019081526020015f206040518060600160405290815f820154815260200160018201548152602001600282015481525050611d26565b611e14576040517f7863d56700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607f81905560405181907f2a2ab116b5abc962503c3c7f941af94e3dc855231d07abb9bc4dc2105591a031905f90a250565b60038054611181906154fa565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611ea4576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600390610100900460ff16158015611ef457505f5460ff8083169116105b611f80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611b31565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c1615611fee576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b891561205f577fffffffff00000000000000000000000000000000000000000000000000000000871615158061202357508715155b1561205a576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120dd565b7fffff000000000000000000000000000000000000000000000000000000000000601088901b167e01000000000000000000000000000000000000000000000000000000000000146120dd576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6120ee86868686868f8f8f8f613ead565b6120f78d613fd5565b604080515f8082526020820190925261213d91612135565b604080518082019091525f808252602082015281526020019060019003908161210f5790505b508d8d614304565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050505050505050565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146121fa576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600390610100900460ff1615801561224a57505f5460ff8083169116105b6122d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611b31565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff808416919091176101001782556002915c1614612346576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745415612380576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61238982613fd5565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200161164a565b5f5f8251606014612420576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f858060200190518101906124379190615652565b919450925090507fffff000000000000000000000000000000000000000000000000000000000000601084901b167e01000000000000000000000000000000000000000000000000000000000000146124bc576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6124c4613a86565b8110156124fd576040517f541d595b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4261250782613a5d565b1061253e576040517f0dffe81800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81612575576040517f80bcf51500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607f545f908152607e602090815260409182902082516060810184528154815260018201549281019290925260020154918101919091526125b581611d26565b6125eb576040517f7863d56700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60746125f66126bf565b8154811061260657612606615625565b5f91825260209182902060029182020154604085810151607c5493548786015188518451978801959095529286018a905260608087018a9052608087019290925260ff909416151560f81b60a08601529290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660a184015260b583019190915260d582015260f5016040516020818303038152906040528051906020012090506126b285610f26565b9890975095505050505050565b6074545f9061197890600190615612565b603e545f9081907501000000000000000000000000000000000000000000900460ff161561278c577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166312b941ca6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612761573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061278591906154e3565b90506127c7565b50604554806127c7576040517fdd41f1ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f6127d2856123e3565b6040517c010000000000000000000000000000000000000000000000000000000060208201526024810183905260448101829052606481018690529193509150608401604051602081830303815290604052805190602001209350505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461288b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001611304565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e545f907501000000000000000000000000000000000000000000900460ff1615612a04576040517f7df73e2700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f00000000000000000000000000000000000000000000000000000000000000001690637df73e2790602401602060405180830381865afa1580156129e0573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061108a9190615686565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526043602052604081208054612a33906154fa565b9050119050919050565b607c54610100900473ffffffffffffffffffffffffffffffffffffffff163314612a93576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c5460ff1615612ad0576040517f98b3177900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a905f90a1565b60015473ffffffffffffffffffffffffffffffffffffffff163314612b76576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314612c67576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54600390610100900460ff16158015612c8757505f5460ff8083169116105b612d13576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611b31565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905560455415612d7d576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080516060810182526079548152607a546020808301918252607b549383019384527fae8304f40f7123e0c87b97f8a600e94ff3a3a25be588fc66b8a3717c8959ce775f819052607e90915291517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840bad55517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840bae5590517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840baf55607f5560025460038054612ee59273ffffffffffffffffffffffffffffffffffffffff169190612e64906154fa565b80601f0160208091040260200160405190810160405280929190818152602001828054612e90906154fa565b8015612edb5780601f10612eb257610100808354040283529160200191612edb565b820191905f5260205f20905b815481529060010190602001808311612ebe57829003601f168201915b50505050506144e0565b6001604455612ef261462c565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611304565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314612fbb576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8051606014612ff6576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f8280602001905181019061300c9190615652565b92509250508061301b60745490565b837fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e24260405161304d91815260200190565b60405180910390a4604080516060810182529283526fffffffffffffffffffffffffffffffff42811660208501908152928116918401918252607480546001810182555f91909152935160029094027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810194909455915190518216700100000000000000000000000000000000029116177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef8149091015550565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613158576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600390610100900460ff161580156131a857505f5460ff8083169116105b613234576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401611b31565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c166001146132a4576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8615613315577fffffffff0000000000000000000000000000000000000000000000000000000084161515806132d957508415155b15613310576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613393565b7fffff000000000000000000000000000000000000000000000000000000000000601085901b167e0100000000000000000000000000000000000000000000000000000000000014613393576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61339c88613fd5565b6133a8878787876146b0565b604080515f808252602082019092526133ee916133e6565b604080518082019091525f80825260208201528152602001906001900390816133c05790505b508484614304565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050505050565b604080516060810182525f80825260208201819052918101919091526074828154811061348457613484615625565b5f91825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff16331461353d576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001611304565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461361f576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811661366c576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101611304565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613737576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116613784576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101611304565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613851576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff16156138a7576040517f278d998800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790556040517f67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b905f90a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314613967576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600361397382826156ec565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516113049190614b85565b603e545f907501000000000000000000000000000000000000000000900460ff1615613a56577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ca69e7dc6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a32573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061197891906154e3565b5060425490565b5f60785460755483613a6f9190615612565b613a799190615803565b60765461108a919061581a565b5f607754613a926119e8565b611978919061581a565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613aed576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818152607e6020526040808220828155600181018390556002018290555182917f4432b02a2fcbed48d94e8d72723e155c6690e4b7f39afa41a2a8ff8c0aa425da91a250565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613b85576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff16613bd9576040517fa2c45f8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07905f90a1565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613c7c576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613c87838383614304565b505050565b603f5473ffffffffffffffffffffffffffffffffffffffff163314613cdd576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff16613d32576040517f5aa930a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1690556040517f4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be905f90a1565b607c54610100900473ffffffffffffffffffffffffffffffffffffffff163314613dda576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116613e27576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255607c5460408051610100909204909316815260208101919091527ff67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af799101611304565b5f54610100900460ff16613f43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611b31565b73ffffffffffffffffffffffffffffffffffffffff89161580613f7a575073ffffffffffffffffffffffffffffffffffffffff8816155b15613fb1576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fbe8989898989614806565b613fca848484846146b0565b505050505050505050565b60c081015173ffffffffffffffffffffffffffffffffffffffff16614026576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060a001515f03614063576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f0361409d576040517fff5f860000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b42816080015111156140db576040517f2403afcb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020810151614116576040517f4bf41e1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a081015160775580516078556074545f036141f85760408051606080820183528383015182526080840180516fffffffffffffffffffffffffffffffff9081166020850190815292860180518216958501958652607480546001810182555f91909152945160029095027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810195909555925194518116700100000000000000000000000000000000029416939093177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef8149092019190915551607555516076555b60c0810151607c80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff90931683021790556040805160608101825260e08401518152918301516020808401918252938401519183019182527fae8304f40f7123e0c87b97f8a600e94ff3a3a25be588fc66b8a3717c8959ce775f819052607e90945291517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840bad5590517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840bae55517fc6915fec25626e33a89ab02e326f2680fcbfd756b7692722484cab2a30840baf55607f55565b6001835111156143ac575f5b6001845161431e9190615612565b8110156143aa578361433182600161581a565b8151811061434157614341615625565b60200260200101516020015184828151811061435f5761435f615625565b602002602001015160200151116143a2576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600101614310565b505b5f5b8351811015614402576143fa8482815181106143cc576143cc615625565b60200260200101515f01518583815181106143e9576143e9615625565b6020026020010151602001516148e7565b6001016143ae565b505f5b82518110156144595761445183828151811061442357614423615625565b60200260200101515f015184838151811061444057614440615625565b6020026020010151602001516144e0565b600101614405565b5060425460ff1015614497576040517f5a7f382c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6042548111156144d3576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6044819055613c8761462c565b73ffffffffffffffffffffffffffffffffffffffff821661452d576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03614567576040517f8715f5fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61457082612930565b156145a7576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60428054600181019091557f38dfe4635b27babeca8be38d3b448cb5161a639b899a14825ba9c8d7892eb8c30180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f908152604360205260409020613c8782826156ec565b604454604260405160200161464292919061582d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209091012060458190556044547f66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa74392612bee9260429291615880565b5f54610100900460ff16614746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611b31565b603e80547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000951515959095027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1694909417750100000000000000000000000000000000000000000093151593909302929092179092557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff000000000000000000000000000000000000000016918616919091179055600361488d83826156ec565b50600461489a82826156ec565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b604254808210614923576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff166042838154811061494d5761494d615625565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16146149a5576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f9081526043602052604081206149d291614ad9565b60426149df600183615612565b815481106149ef576149ef615625565b5f918252602090912001546042805473ffffffffffffffffffffffffffffffffffffffff9092169184908110614a2757614a27615625565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506042805480614a7d57614a7d6158e7565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505050565b508054614ae5906154fa565b5f825580601f10614af4575050565b601f0160209004905f5260205f2090810190614b109190614b13565b50565b5b80821115614b27575f8155600101614b14565b5090565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114614b10575f5ffd5b8035610fb781614b2b565b5f60208284031215614b73575f5ffd5b8135614b7e81614b2b565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215614be9575f5ffd5b8235614bf481614b2b565b946020939093013593505050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610fb7575f5ffd5b5f5f60408385031215614c42575f5ffd5b614c4b83614c02565b9150614c5960208401614c02565b90509250929050565b5f60208284031215614c72575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114614b10575f5ffd5b8035610fb781614c79565b5f60208284031215614cb5575f5ffd5b8135614b7e81614c79565b602080825282518282018190525f918401906040840190835b81811015614d0d57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101614cd9565b509095945050505050565b5f5f5f5f60808587031215614d2b575f5ffd5b5050823594602084013594506040840135936060013592509050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610120810167ffffffffffffffff81118282101715614d9857614d98614d47565b60405290565b6040805190810167ffffffffffffffff81118282101715614d9857614d98614d47565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614e0857614e08614d47565b604052919050565b5f6060828403128015614e21575f5ffd5b506040516060810167ffffffffffffffff81118282101715614e4557614e45614d47565b60409081528335825260208085013590830152928301359281019290925250919050565b5f6101208284031215614e7a575f5ffd5b614e82614d74565b823581526020808401359082015260408084013590820152606080840135908201526080808401359082015260a080840135908201529050614ec660c08301614c9a565b60c082015260e082810135908201526101009182013591810191909152919050565b5f67ffffffffffffffff821115614f0157614f01614d47565b5060051b60200190565b5f5f67ffffffffffffffff841115614f2557614f25614d47565b50601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001614f5881614dc1565b915050828152838383011115614f6c575f5ffd5b828260208301375f602084830101529392505050565b5f82601f830112614f91575f5ffd5b614b7e83833560208501614f0b565b5f82601f830112614faf575f5ffd5b8135614fc2614fbd82614ee8565b614dc1565b8082825260208201915060208360051b860101925085831115614fe3575f5ffd5b602085015b8381101561509557803567ffffffffffffffff811115615006575f5ffd5b860160408189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215615039575f5ffd5b615041614d9e565b602082013561504f81614c79565b8152604082013567ffffffffffffffff81111561506a575f5ffd5b6150798a602083860101614f82565b6020830152508085525050602083019250602081019050614fe8565b5095945050505050565b8015158114614b10575f5ffd5b8035610fb78161509f565b5f5f5f5f5f5f5f5f5f5f5f5f6102808d8f0312156150d3575f5ffd5b6150dd8e8e614e69565b9b5067ffffffffffffffff6101208e013511156150f8575f5ffd5b6151098e6101208f01358f01614fa0565b9a506101408d013599506151206101608e016150ac565b985061512f6101808e016150ac565b97506101a08d013596506151466101c08e01614b58565b95506151556101e08e01614c9a565b94506151646102008e01614c9a565b93506151736102208e01614c9a565b925067ffffffffffffffff6102408e0135111561518e575f5ffd5b61519f8e6102408f01358f01614f82565b915067ffffffffffffffff6102608e013511156151ba575f5ffd5b6151cb8e6102608f01358f01614f82565b90509295989b509295989b509295989b565b5f61012082840312156151ee575f5ffd5b614b7e8383614e69565b5f60208284031215615208575f5ffd5b813567ffffffffffffffff81111561521e575f5ffd5b8201601f8101841361522e575f5ffd5b61523d84823560208401614f0b565b949350505050565b5f60208284031215615255575f5ffd5b813567ffffffffffffffff81168114614b7e575f5ffd5b5f5f5f5f5f5f60c08789031215615281575f5ffd5b863561528c81614c79565b9550602087013561529c81614c79565b9450604087013563ffffffff811681146152b4575f5ffd5b935060608701356152c481614c79565b9250608087013567ffffffffffffffff8111156152df575f5ffd5b6152eb89828a01614f82565b92505060a087013567ffffffffffffffff811115615307575f5ffd5b61531389828a01614f82565b9150509295509295509295565b5f5f5f5f5f5f5f6101e0888a031215615337575f5ffd5b6153418989614e69565b96506101208801356153528161509f565b95506101408801356153638161509f565b9450610160880135935061018088013561537c81614b2b565b92506101a088013567ffffffffffffffff811115615398575f5ffd5b6153a48a828b01614fa0565b979a96995094979396929592945050506101c09091013590565b5f602082840312156153ce575f5ffd5b813567ffffffffffffffff8111156153e4575f5ffd5b61523d84828501614f82565b5f5f5f60608486031215615402575f5ffd5b833567ffffffffffffffff811115615418575f5ffd5b8401601f81018613615428575f5ffd5b8035615436614fbd82614ee8565b8082825260208201915060208360061b850101925088831115615457575f5ffd5b6020840193505b828410156154a7576040848a031215615475575f5ffd5b61547d614d9e565b843561548881614c79565b815260208581013581830152908352604090940193919091019061545e565b9550505050602084013567ffffffffffffffff8111156154c5575f5ffd5b6154d186828701614fa0565b93969395505050506040919091013590565b5f602082840312156154f3575f5ffd5b5051919050565b600181811c9082168061550e57607f821691505b602082108103615545577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f6020828403121561555b575f5ffd5b815167ffffffffffffffff811115615571575f5ffd5b8201601f81018413615581575f5ffd5b805161558f614fbd82614ee8565b8082825260208201915060208360051b8501019250868311156155b0575f5ffd5b6020840193505b828410156155db5783516155ca81614c79565b8252602093840193909101906155b7565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561108a5761108a6155e5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f60608486031215615664575f5ffd5b835161566f81614b2b565b602085015160409095015190969495509392505050565b5f60208284031215615696575f5ffd5b8151614b7e8161509f565b601f821115613c8757805f5260205f20601f840160051c810160208510156156c65750805b601f840160051c820191505b818110156156e5575f81556001016156d2565b5050505050565b815167ffffffffffffffff81111561570657615706614d47565b61571a8161571484546154fa565b846156a1565b6020601f82116001811461576b575f83156157355750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556156e5565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156157b85787850151825560209485019460019092019101615798565b50848210156157f457868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b808202811582820484141761108a5761108a6155e5565b8082018082111561108a5761108a6155e5565b8281525f602082018354845f5260205f205f5b8281101561587457815473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019182019101615840565b50919695505050505050565b606080825284549082018190525f8581526020812090916080840190835b818110156158d257835473ffffffffffffffffffffffffffffffffffffffff1683526001938401936020909301920161589e565b50506020840195909552505060400152919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212203fbfca958ab65a6bc00f9a03672ef5e92c80dd2b655f742be451f0fb92860da664736f6c634300081c0033",
}

// AggchainfepABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainfepMetaData.ABI instead.
var AggchainfepABI = AggchainfepMetaData.ABI

// AggchainfepBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggchainfepMetaData.Bin instead.
var AggchainfepBin = AggchainfepMetaData.Bin

// DeployAggchainfep deploys a new Ethereum contract, binding an instance of Aggchainfep to it.
func DeployAggchainfep(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Aggchainfep, error) {
	parsed, err := AggchainfepMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggchainfepBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggchainfep{AggchainfepCaller: AggchainfepCaller{contract: contract}, AggchainfepTransactor: AggchainfepTransactor{contract: contract}, AggchainfepFilterer: AggchainfepFilterer{contract: contract}}, nil
}

// Aggchainfep is an auto generated Go binding around an Ethereum contract.
type Aggchainfep struct {
	AggchainfepCaller     // Read-only binding to the contract
	AggchainfepTransactor // Write-only binding to the contract
	AggchainfepFilterer   // Log filterer for contract events
}

// AggchainfepCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainfepCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfepTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainfepTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfepFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainfepFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfepSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainfepSession struct {
	Contract     *Aggchainfep      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggchainfepCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainfepCallerSession struct {
	Contract *AggchainfepCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AggchainfepTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainfepTransactorSession struct {
	Contract     *AggchainfepTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AggchainfepRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainfepRaw struct {
	Contract *Aggchainfep // Generic contract binding to access the raw methods on
}

// AggchainfepCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainfepCallerRaw struct {
	Contract *AggchainfepCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainfepTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainfepTransactorRaw struct {
	Contract *AggchainfepTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainfep creates a new instance of Aggchainfep, bound to a specific deployed contract.
func NewAggchainfep(address common.Address, backend bind.ContractBackend) (*Aggchainfep, error) {
	contract, err := bindAggchainfep(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggchainfep{AggchainfepCaller: AggchainfepCaller{contract: contract}, AggchainfepTransactor: AggchainfepTransactor{contract: contract}, AggchainfepFilterer: AggchainfepFilterer{contract: contract}}, nil
}

// NewAggchainfepCaller creates a new read-only instance of Aggchainfep, bound to a specific deployed contract.
func NewAggchainfepCaller(address common.Address, caller bind.ContractCaller) (*AggchainfepCaller, error) {
	contract, err := bindAggchainfep(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainfepCaller{contract: contract}, nil
}

// NewAggchainfepTransactor creates a new write-only instance of Aggchainfep, bound to a specific deployed contract.
func NewAggchainfepTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainfepTransactor, error) {
	contract, err := bindAggchainfep(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransactor{contract: contract}, nil
}

// NewAggchainfepFilterer creates a new log filterer instance of Aggchainfep, bound to a specific deployed contract.
func NewAggchainfepFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainfepFilterer, error) {
	contract, err := bindAggchainfep(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainfepFilterer{contract: contract}, nil
}

// bindAggchainfep binds a generic wrapper to an already deployed contract.
func bindAggchainfep(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainfepMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainfep *AggchainfepRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainfep.Contract.AggchainfepCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainfep *AggchainfepRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AggchainfepTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainfep *AggchainfepRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AggchainfepTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainfep *AggchainfepCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainfep.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainfep *AggchainfepTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainfep *AggchainfepTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainfep.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfep *AggchainfepCaller) AGGCHAINFEPVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "AGGCHAIN_FEP_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfep *AggchainfepSession) AGGCHAINFEPVERSION() (string, error) {
	return _Aggchainfep.Contract.AGGCHAINFEPVERSION(&_Aggchainfep.CallOpts)
}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfep *AggchainfepCallerSession) AGGCHAINFEPVERSION() (string, error) {
	return _Aggchainfep.Contract.AGGCHAINFEPVERSION(&_Aggchainfep.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfep *AggchainfepCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfep *AggchainfepSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPE(&_Aggchainfep.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfep *AggchainfepCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainfep.Contract.AGGCHAINTYPE(&_Aggchainfep.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainfep.Contract.CONSENSUSTYPE(&_Aggchainfep.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfep *AggchainfepCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainfep.Contract.CONSENSUSTYPE(&_Aggchainfep.CallOpts)
}

// GENESISCONFIGNAME is a free data retrieval call binding the contract method 0xf72f606d.
//
// Solidity: function GENESIS_CONFIG_NAME() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) GENESISCONFIGNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "GENESIS_CONFIG_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENESISCONFIGNAME is a free data retrieval call binding the contract method 0xf72f606d.
//
// Solidity: function GENESIS_CONFIG_NAME() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) GENESISCONFIGNAME() ([32]byte, error) {
	return _Aggchainfep.Contract.GENESISCONFIGNAME(&_Aggchainfep.CallOpts)
}

// GENESISCONFIGNAME is a free data retrieval call binding the contract method 0xf72f606d.
//
// Solidity: function GENESIS_CONFIG_NAME() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) GENESISCONFIGNAME() ([32]byte, error) {
	return _Aggchainfep.Contract.GENESISCONFIGNAME(&_Aggchainfep.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) L2BLOCKTIME() (*big.Int, error) {
	return _Aggchainfep.Contract.L2BLOCKTIME(&_Aggchainfep.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _Aggchainfep.Contract.L2BLOCKTIME(&_Aggchainfep.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _Aggchainfep.Contract.SUBMISSIONINTERVAL(&_Aggchainfep.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _Aggchainfep.Contract.SUBMISSIONINTERVAL(&_Aggchainfep.CallOpts)
}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) LegacypendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "_legacypendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) LegacypendingVKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.LegacypendingVKeyManager(&_Aggchainfep.CallOpts)
}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) LegacypendingVKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.LegacypendingVKeyManager(&_Aggchainfep.CallOpts)
}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) LegacyvKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "_legacyvKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) LegacyvKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.LegacyvKeyManager(&_Aggchainfep.CallOpts)
}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) LegacyvKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.LegacyvKeyManager(&_Aggchainfep.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfep *AggchainfepCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfep *AggchainfepSession) Admin() (common.Address, error) {
	return _Aggchainfep.Contract.Admin(&_Aggchainfep.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) Admin() (common.Address, error) {
	return _Aggchainfep.Contract.Admin(&_Aggchainfep.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfep *AggchainfepCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfep *AggchainfepSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainfep.Contract.AggLayerGateway(&_Aggchainfep.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainfep.Contract.AggLayerGateway(&_Aggchainfep.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) AggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "aggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) AggchainManager() (common.Address, error) {
	return _Aggchainfep.Contract.AggchainManager(&_Aggchainfep.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) AggchainManager() (common.Address, error) {
	return _Aggchainfep.Contract.AggchainManager(&_Aggchainfep.CallOpts)
}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainfep *AggchainfepCaller) AggchainSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "aggchainSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainfep *AggchainfepSession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Aggchainfep.Contract.AggchainSigners(&_Aggchainfep.CallOpts, arg0)
}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Aggchainfep.Contract.AggchainSigners(&_Aggchainfep.CallOpts, arg0)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) AggchainSignersHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "aggchainSignersHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) AggchainSignersHash() ([32]byte, error) {
	return _Aggchainfep.Contract.AggchainSignersHash(&_Aggchainfep.CallOpts)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) AggchainSignersHash() ([32]byte, error) {
	return _Aggchainfep.Contract.AggchainSignersHash(&_Aggchainfep.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) AggregationVkey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "aggregationVkey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) AggregationVkey() ([32]byte, error) {
	return _Aggchainfep.Contract.AggregationVkey(&_Aggchainfep.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) AggregationVkey() ([32]byte, error) {
	return _Aggchainfep.Contract.AggregationVkey(&_Aggchainfep.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfep *AggchainfepCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfep *AggchainfepSession) BridgeAddress() (common.Address, error) {
	return _Aggchainfep.Contract.BridgeAddress(&_Aggchainfep.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) BridgeAddress() (common.Address, error) {
	return _Aggchainfep.Contract.BridgeAddress(&_Aggchainfep.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfep *AggchainfepSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _Aggchainfep.Contract.ComputeL2Timestamp(&_Aggchainfep.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _Aggchainfep.Contract.ComputeL2Timestamp(&_Aggchainfep.CallOpts, _l2BlockNumber)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfep *AggchainfepCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfep *AggchainfepSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainfep.Contract.ForceBatchAddress(&_Aggchainfep.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainfep.Contract.ForceBatchAddress(&_Aggchainfep.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfep *AggchainfepCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfep *AggchainfepSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainfep.Contract.ForceBatchTimeout(&_Aggchainfep.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfep *AggchainfepCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainfep.Contract.ForceBatchTimeout(&_Aggchainfep.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainfep.Contract.ForcedBatches(&_Aggchainfep.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainfep.Contract.ForcedBatches(&_Aggchainfep.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfep *AggchainfepCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfep *AggchainfepSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainfep.Contract.GasTokenAddress(&_Aggchainfep.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainfep.Contract.GasTokenAddress(&_Aggchainfep.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfep *AggchainfepCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfep *AggchainfepSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainfep.Contract.GasTokenNetwork(&_Aggchainfep.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfep *AggchainfepCallerSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainfep.Contract.GasTokenNetwork(&_Aggchainfep.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainHash(&_Aggchainfep.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainHash(&_Aggchainfep.CallOpts, aggchainData)
}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) view returns(bytes32, bytes32)
func (_Aggchainfep *AggchainfepCaller) GetAggchainParamsAndVKeySelector(opts *bind.CallOpts, aggchainData []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainParamsAndVKeySelector", aggchainData)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) view returns(bytes32, bytes32)
func (_Aggchainfep *AggchainfepSession) GetAggchainParamsAndVKeySelector(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainParamsAndVKeySelector(&_Aggchainfep.CallOpts, aggchainData)
}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) view returns(bytes32, bytes32)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainParamsAndVKeySelector(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainParamsAndVKeySelector(&_Aggchainfep.CallOpts, aggchainData)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainfep *AggchainfepCaller) GetAggchainSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainfep *AggchainfepSession) GetAggchainSigners() ([]common.Address, error) {
	return _Aggchainfep.Contract.GetAggchainSigners(&_Aggchainfep.CallOpts)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainSigners() ([]common.Address, error) {
	return _Aggchainfep.Contract.GetAggchainSigners(&_Aggchainfep.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) GetAggchainSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Aggchainfep.Contract.GetAggchainSignersCount(&_Aggchainfep.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Aggchainfep.Contract.GetAggchainSignersCount(&_Aggchainfep.CallOpts)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepCaller) GetAggchainTypeFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainTypeFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfep.Contract.GetAggchainTypeFromSelector(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfep.Contract.GetAggchainTypeFromSelector(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainVKey", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKey(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKey(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfep *AggchainfepCaller) GetAggchainVKeySelector(opts *bind.CallOpts, aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainVKeySelector", aggchainVKeyVersion, aggchainType)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfep *AggchainfepSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKeySelector(&_Aggchainfep.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKeySelector(&_Aggchainfep.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepCaller) GetAggchainVKeyVersionFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getAggchainVKeyVersionFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfep *AggchainfepCallerSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfep.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfep *AggchainfepCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(AggchainFEPOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(AggchainFEPOutputProposal)).(*AggchainFEPOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfep *AggchainfepSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	return _Aggchainfep.Contract.GetL2Output(&_Aggchainfep.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfep *AggchainfepCallerSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPOutputProposal, error) {
	return _Aggchainfep.Contract.GetL2Output(&_Aggchainfep.CallOpts, _l2OutputIndex)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainfep.Contract.GlobalExitRootManager(&_Aggchainfep.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainfep.Contract.GlobalExitRootManager(&_Aggchainfep.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfep *AggchainfepCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfep *AggchainfepSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainfep.Contract.Initialize0(&_Aggchainfep.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfep *AggchainfepCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainfep.Contract.Initialize0(&_Aggchainfep.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainfep *AggchainfepCaller) IsSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "isSigner", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainfep *AggchainfepSession) IsSigner(_signer common.Address) (bool, error) {
	return _Aggchainfep.Contract.IsSigner(&_Aggchainfep.CallOpts, _signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) IsSigner(_signer common.Address) (bool, error) {
	return _Aggchainfep.Contract.IsSigner(&_Aggchainfep.CallOpts, _signer)
}

// IsValidOpSuccinctConfig is a free data retrieval call binding the contract method 0x49185e06.
//
// Solidity: function isValidOpSuccinctConfig((bytes32,bytes32,bytes32) _config) pure returns(bool)
func (_Aggchainfep *AggchainfepCaller) IsValidOpSuccinctConfig(opts *bind.CallOpts, _config AggchainFEPOpSuccinctConfig) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "isValidOpSuccinctConfig", _config)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOpSuccinctConfig is a free data retrieval call binding the contract method 0x49185e06.
//
// Solidity: function isValidOpSuccinctConfig((bytes32,bytes32,bytes32) _config) pure returns(bool)
func (_Aggchainfep *AggchainfepSession) IsValidOpSuccinctConfig(_config AggchainFEPOpSuccinctConfig) (bool, error) {
	return _Aggchainfep.Contract.IsValidOpSuccinctConfig(&_Aggchainfep.CallOpts, _config)
}

// IsValidOpSuccinctConfig is a free data retrieval call binding the contract method 0x49185e06.
//
// Solidity: function isValidOpSuccinctConfig((bytes32,bytes32,bytes32) _config) pure returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) IsValidOpSuccinctConfig(_config AggchainFEPOpSuccinctConfig) (bool, error) {
	return _Aggchainfep.Contract.IsValidOpSuccinctConfig(&_Aggchainfep.CallOpts, _config)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) L2BlockTime() (*big.Int, error) {
	return _Aggchainfep.Contract.L2BlockTime(&_Aggchainfep.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) L2BlockTime() (*big.Int, error) {
	return _Aggchainfep.Contract.L2BlockTime(&_Aggchainfep.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainfep.Contract.LastAccInputHash(&_Aggchainfep.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainfep.Contract.LastAccInputHash(&_Aggchainfep.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfep *AggchainfepCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfep *AggchainfepSession) LastForceBatch() (uint64, error) {
	return _Aggchainfep.Contract.LastForceBatch(&_Aggchainfep.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfep *AggchainfepCallerSession) LastForceBatch() (uint64, error) {
	return _Aggchainfep.Contract.LastForceBatch(&_Aggchainfep.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfep *AggchainfepCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfep *AggchainfepSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainfep.Contract.LastForceBatchSequenced(&_Aggchainfep.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfep *AggchainfepCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainfep.Contract.LastForceBatchSequenced(&_Aggchainfep.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) LatestBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.LatestBlockNumber(&_Aggchainfep.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.LatestBlockNumber(&_Aggchainfep.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) LatestOutputIndex() (*big.Int, error) {
	return _Aggchainfep.Contract.LatestOutputIndex(&_Aggchainfep.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _Aggchainfep.Contract.LatestOutputIndex(&_Aggchainfep.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfep *AggchainfepCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfep *AggchainfepSession) NetworkName() (string, error) {
	return _Aggchainfep.Contract.NetworkName(&_Aggchainfep.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfep *AggchainfepCallerSession) NetworkName() (string, error) {
	return _Aggchainfep.Contract.NetworkName(&_Aggchainfep.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) NextBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.NextBlockNumber(&_Aggchainfep.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) NextBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.NextBlockNumber(&_Aggchainfep.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) NextOutputIndex() (*big.Int, error) {
	return _Aggchainfep.Contract.NextOutputIndex(&_Aggchainfep.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) NextOutputIndex() (*big.Int, error) {
	return _Aggchainfep.Contract.NextOutputIndex(&_Aggchainfep.CallOpts)
}

// OpSuccinctConfigs is a free data retrieval call binding the contract method 0x6a56620b.
//
// Solidity: function opSuccinctConfigs(bytes32 ) view returns(bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepCaller) OpSuccinctConfigs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AggregationVkey     [32]byte
	RangeVkeyCommitment [32]byte
	RollupConfigHash    [32]byte
}, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "opSuccinctConfigs", arg0)

	outstruct := new(struct {
		AggregationVkey     [32]byte
		RangeVkeyCommitment [32]byte
		RollupConfigHash    [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AggregationVkey = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RangeVkeyCommitment = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.RollupConfigHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// OpSuccinctConfigs is a free data retrieval call binding the contract method 0x6a56620b.
//
// Solidity: function opSuccinctConfigs(bytes32 ) view returns(bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepSession) OpSuccinctConfigs(arg0 [32]byte) (struct {
	AggregationVkey     [32]byte
	RangeVkeyCommitment [32]byte
	RollupConfigHash    [32]byte
}, error) {
	return _Aggchainfep.Contract.OpSuccinctConfigs(&_Aggchainfep.CallOpts, arg0)
}

// OpSuccinctConfigs is a free data retrieval call binding the contract method 0x6a56620b.
//
// Solidity: function opSuccinctConfigs(bytes32 ) view returns(bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepCallerSession) OpSuccinctConfigs(arg0 [32]byte) (struct {
	AggregationVkey     [32]byte
	RangeVkeyCommitment [32]byte
	RollupConfigHash    [32]byte
}, error) {
	return _Aggchainfep.Contract.OpSuccinctConfigs(&_Aggchainfep.CallOpts, arg0)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfep *AggchainfepCaller) OptimisticMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "optimisticMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfep *AggchainfepSession) OptimisticMode() (bool, error) {
	return _Aggchainfep.Contract.OptimisticMode(&_Aggchainfep.CallOpts)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) OptimisticMode() (bool, error) {
	return _Aggchainfep.Contract.OptimisticMode(&_Aggchainfep.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) OptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "optimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) OptimisticModeManager() (common.Address, error) {
	return _Aggchainfep.Contract.OptimisticModeManager(&_Aggchainfep.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) OptimisticModeManager() (common.Address, error) {
	return _Aggchainfep.Contract.OptimisticModeManager(&_Aggchainfep.CallOpts)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfep *AggchainfepCaller) OwnedAggchainVKeys(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "ownedAggchainVKeys", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfep *AggchainfepSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.OwnedAggchainVKeys(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfep *AggchainfepCallerSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfep.Contract.OwnedAggchainVKeys(&_Aggchainfep.CallOpts, aggchainVKeySelector)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfep *AggchainfepCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfep *AggchainfepSession) PendingAdmin() (common.Address, error) {
	return _Aggchainfep.Contract.PendingAdmin(&_Aggchainfep.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) PendingAdmin() (common.Address, error) {
	return _Aggchainfep.Contract.PendingAdmin(&_Aggchainfep.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) PendingAggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "pendingAggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingAggchainManager(&_Aggchainfep.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingAggchainManager(&_Aggchainfep.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) PendingOptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "pendingOptimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) PendingOptimisticModeManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingOptimisticModeManager(&_Aggchainfep.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) PendingOptimisticModeManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingOptimisticModeManager(&_Aggchainfep.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfep *AggchainfepCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfep *AggchainfepSession) Pol() (common.Address, error) {
	return _Aggchainfep.Contract.Pol(&_Aggchainfep.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) Pol() (common.Address, error) {
	return _Aggchainfep.Contract.Pol(&_Aggchainfep.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) RangeVkeyCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "rangeVkeyCommitment")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) RangeVkeyCommitment() ([32]byte, error) {
	return _Aggchainfep.Contract.RangeVkeyCommitment(&_Aggchainfep.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) RangeVkeyCommitment() ([32]byte, error) {
	return _Aggchainfep.Contract.RangeVkeyCommitment(&_Aggchainfep.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) RollupConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "rollupConfigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) RollupConfigHash() ([32]byte, error) {
	return _Aggchainfep.Contract.RollupConfigHash(&_Aggchainfep.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) RollupConfigHash() ([32]byte, error) {
	return _Aggchainfep.Contract.RollupConfigHash(&_Aggchainfep.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) RollupManager() (common.Address, error) {
	return _Aggchainfep.Contract.RollupManager(&_Aggchainfep.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) RollupManager() (common.Address, error) {
	return _Aggchainfep.Contract.RollupManager(&_Aggchainfep.CallOpts)
}

// SelectedOpSuccinctConfigName is a free data retrieval call binding the contract method 0x9f78f066.
//
// Solidity: function selectedOpSuccinctConfigName() view returns(bytes32)
func (_Aggchainfep *AggchainfepCaller) SelectedOpSuccinctConfigName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "selectedOpSuccinctConfigName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SelectedOpSuccinctConfigName is a free data retrieval call binding the contract method 0x9f78f066.
//
// Solidity: function selectedOpSuccinctConfigName() view returns(bytes32)
func (_Aggchainfep *AggchainfepSession) SelectedOpSuccinctConfigName() ([32]byte, error) {
	return _Aggchainfep.Contract.SelectedOpSuccinctConfigName(&_Aggchainfep.CallOpts)
}

// SelectedOpSuccinctConfigName is a free data retrieval call binding the contract method 0x9f78f066.
//
// Solidity: function selectedOpSuccinctConfigName() view returns(bytes32)
func (_Aggchainfep *AggchainfepCallerSession) SelectedOpSuccinctConfigName() ([32]byte, error) {
	return _Aggchainfep.Contract.SelectedOpSuccinctConfigName(&_Aggchainfep.CallOpts)
}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainfep *AggchainfepCaller) SignerToURLs(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "signerToURLs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainfep *AggchainfepSession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Aggchainfep.Contract.SignerToURLs(&_Aggchainfep.CallOpts, arg0)
}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainfep *AggchainfepCallerSession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Aggchainfep.Contract.SignerToURLs(&_Aggchainfep.CallOpts, arg0)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) StartingBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.StartingBlockNumber(&_Aggchainfep.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _Aggchainfep.Contract.StartingBlockNumber(&_Aggchainfep.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) StartingTimestamp() (*big.Int, error) {
	return _Aggchainfep.Contract.StartingTimestamp(&_Aggchainfep.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) StartingTimestamp() (*big.Int, error) {
	return _Aggchainfep.Contract.StartingTimestamp(&_Aggchainfep.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) SubmissionInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "submissionInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) SubmissionInterval() (*big.Int, error) {
	return _Aggchainfep.Contract.SubmissionInterval(&_Aggchainfep.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) SubmissionInterval() (*big.Int, error) {
	return _Aggchainfep.Contract.SubmissionInterval(&_Aggchainfep.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainfep *AggchainfepCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainfep *AggchainfepSession) Threshold() (*big.Int, error) {
	return _Aggchainfep.Contract.Threshold(&_Aggchainfep.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainfep *AggchainfepCallerSession) Threshold() (*big.Int, error) {
	return _Aggchainfep.Contract.Threshold(&_Aggchainfep.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfep *AggchainfepCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfep *AggchainfepSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainfep.Contract.TrustedSequencer(&_Aggchainfep.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainfep.Contract.TrustedSequencer(&_Aggchainfep.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfep *AggchainfepCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfep *AggchainfepSession) TrustedSequencerURL() (string, error) {
	return _Aggchainfep.Contract.TrustedSequencerURL(&_Aggchainfep.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfep *AggchainfepCallerSession) TrustedSequencerURL() (string, error) {
	return _Aggchainfep.Contract.TrustedSequencerURL(&_Aggchainfep.CallOpts)
}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainfep *AggchainfepCaller) UseDefaultSigners(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "useDefaultSigners")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainfep *AggchainfepSession) UseDefaultSigners() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultSigners(&_Aggchainfep.CallOpts)
}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) UseDefaultSigners() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultSigners(&_Aggchainfep.CallOpts)
}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainfep *AggchainfepCaller) UseDefaultVkeys(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "useDefaultVkeys")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainfep *AggchainfepSession) UseDefaultVkeys() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultVkeys(&_Aggchainfep.CallOpts)
}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) UseDefaultVkeys() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultVkeys(&_Aggchainfep.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfep *AggchainfepCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfep *AggchainfepSession) Version() (string, error) {
	return _Aggchainfep.Contract.Version(&_Aggchainfep.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfep *AggchainfepCallerSession) Version() (string, error) {
	return _Aggchainfep.Contract.Version(&_Aggchainfep.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfep *AggchainfepTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfep *AggchainfepSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptAdminRole(&_Aggchainfep.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfep *AggchainfepTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptAdminRole(&_Aggchainfep.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactor) AcceptAggchainManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "acceptAggchainManagerRole")
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfep *AggchainfepSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptAggchainManagerRole(&_Aggchainfep.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactorSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptAggchainManagerRole(&_Aggchainfep.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactor) AcceptOptimisticModeManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "acceptOptimisticModeManagerRole")
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfep *AggchainfepSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptOptimisticModeManagerRole(&_Aggchainfep.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactorSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptOptimisticModeManagerRole(&_Aggchainfep.TransactOpts)
}

// AddOpSuccinctConfig is a paid mutator transaction binding the contract method 0x47c37e9c.
//
// Solidity: function addOpSuccinctConfig(bytes32 _configName, bytes32 _rollupConfigHash, bytes32 _aggregationVkey, bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfep *AggchainfepTransactor) AddOpSuccinctConfig(opts *bind.TransactOpts, _configName [32]byte, _rollupConfigHash [32]byte, _aggregationVkey [32]byte, _rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "addOpSuccinctConfig", _configName, _rollupConfigHash, _aggregationVkey, _rangeVkeyCommitment)
}

// AddOpSuccinctConfig is a paid mutator transaction binding the contract method 0x47c37e9c.
//
// Solidity: function addOpSuccinctConfig(bytes32 _configName, bytes32 _rollupConfigHash, bytes32 _aggregationVkey, bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfep *AggchainfepSession) AddOpSuccinctConfig(_configName [32]byte, _rollupConfigHash [32]byte, _aggregationVkey [32]byte, _rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName, _rollupConfigHash, _aggregationVkey, _rangeVkeyCommitment)
}

// AddOpSuccinctConfig is a paid mutator transaction binding the contract method 0x47c37e9c.
//
// Solidity: function addOpSuccinctConfig(bytes32 _configName, bytes32 _rollupConfigHash, bytes32 _aggregationVkey, bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfep *AggchainfepTransactorSession) AddOpSuccinctConfig(_configName [32]byte, _rollupConfigHash [32]byte, _aggregationVkey [32]byte, _rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName, _rollupConfigHash, _aggregationVkey, _rangeVkeyCommitment)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "addOwnedAggchainVKey", aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactorSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.AddOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// DeleteOpSuccinctConfig is a paid mutator transaction binding the contract method 0xec5b2e3a.
//
// Solidity: function deleteOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepTransactor) DeleteOpSuccinctConfig(opts *bind.TransactOpts, _configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "deleteOpSuccinctConfig", _configName)
}

// DeleteOpSuccinctConfig is a paid mutator transaction binding the contract method 0xec5b2e3a.
//
// Solidity: function deleteOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepSession) DeleteOpSuccinctConfig(_configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.DeleteOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName)
}

// DeleteOpSuccinctConfig is a paid mutator transaction binding the contract method 0xec5b2e3a.
//
// Solidity: function deleteOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepTransactorSession) DeleteOpSuccinctConfig(_configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.DeleteOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepTransactor) DisableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "disableOptimisticMode")
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableOptimisticMode(&_Aggchainfep.TransactOpts)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepTransactorSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableOptimisticMode(&_Aggchainfep.TransactOpts)
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) DisableUseDefaultSignersFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "disableUseDefaultSignersFlag")
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepSession) DisableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultSignersFlag(&_Aggchainfep.TransactOpts)
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) DisableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultSignersFlag(&_Aggchainfep.TransactOpts)
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) DisableUseDefaultVkeysFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "disableUseDefaultVkeysFlag")
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepSession) DisableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultVkeysFlag(&_Aggchainfep.TransactOpts)
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) DisableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultVkeysFlag(&_Aggchainfep.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepTransactor) EnableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "enableOptimisticMode")
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableOptimisticMode(&_Aggchainfep.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfep *AggchainfepTransactorSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableOptimisticMode(&_Aggchainfep.TransactOpts)
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) EnableUseDefaultSignersFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "enableUseDefaultSignersFlag")
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepSession) EnableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultSignersFlag(&_Aggchainfep.TransactOpts)
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) EnableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultSignersFlag(&_Aggchainfep.TransactOpts)
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) EnableUseDefaultVkeysFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "enableUseDefaultVkeysFlag")
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepSession) EnableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultVkeysFlag(&_Aggchainfep.TransactOpts)
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) EnableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultVkeysFlag(&_Aggchainfep.TransactOpts)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepTransactor) InitAggchainManager(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initAggchainManager", newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitAggchainManager(&_Aggchainfep.TransactOpts, newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepTransactorSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitAggchainManager(&_Aggchainfep.TransactOpts, newAggchainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x558716c1.
//
// Solidity: function initialize((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, (address,string)[] _signersToAdd, uint256 _newThreshold, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName) returns()
func (_Aggchainfep *AggchainfepTransactor) Initialize(opts *bind.TransactOpts, _initParams AggchainFEPInitParams, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initialize", _initParams, _signersToAdd, _newThreshold, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x558716c1.
//
// Solidity: function initialize((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, (address,string)[] _signersToAdd, uint256 _newThreshold, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName) returns()
func (_Aggchainfep *AggchainfepSession) Initialize(_initParams AggchainFEPInitParams, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, _initParams, _signersToAdd, _newThreshold, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x558716c1.
//
// Solidity: function initialize((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, (address,string)[] _signersToAdd, uint256 _newThreshold, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName) returns()
func (_Aggchainfep *AggchainfepTransactorSession) Initialize(_initParams AggchainFEPInitParams, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, _initParams, _signersToAdd, _newThreshold, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName)
}

// InitializeFromECDSAMultisig is a paid mutator transaction binding the contract method 0x63de5bc6.
//
// Solidity: function initializeFromECDSAMultisig((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams) returns()
func (_Aggchainfep *AggchainfepTransactor) InitializeFromECDSAMultisig(opts *bind.TransactOpts, _initParams AggchainFEPInitParams) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initializeFromECDSAMultisig", _initParams)
}

// InitializeFromECDSAMultisig is a paid mutator transaction binding the contract method 0x63de5bc6.
//
// Solidity: function initializeFromECDSAMultisig((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams) returns()
func (_Aggchainfep *AggchainfepSession) InitializeFromECDSAMultisig(_initParams AggchainFEPInitParams) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitializeFromECDSAMultisig(&_Aggchainfep.TransactOpts, _initParams)
}

// InitializeFromECDSAMultisig is a paid mutator transaction binding the contract method 0x63de5bc6.
//
// Solidity: function initializeFromECDSAMultisig((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams) returns()
func (_Aggchainfep *AggchainfepTransactorSession) InitializeFromECDSAMultisig(_initParams AggchainFEPInitParams) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitializeFromECDSAMultisig(&_Aggchainfep.TransactOpts, _initParams)
}

// InitializeFromPessimisticConsensus is a paid mutator transaction binding the contract method 0xa2398fc6.
//
// Solidity: function initializeFromPessimisticConsensus((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepTransactor) InitializeFromPessimisticConsensus(opts *bind.TransactOpts, _initParams AggchainFEPInitParams, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initializeFromPessimisticConsensus", _initParams, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _signersToAdd, _newThreshold)
}

// InitializeFromPessimisticConsensus is a paid mutator transaction binding the contract method 0xa2398fc6.
//
// Solidity: function initializeFromPessimisticConsensus((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepSession) InitializeFromPessimisticConsensus(_initParams AggchainFEPInitParams, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitializeFromPessimisticConsensus(&_Aggchainfep.TransactOpts, _initParams, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _signersToAdd, _newThreshold)
}

// InitializeFromPessimisticConsensus is a paid mutator transaction binding the contract method 0xa2398fc6.
//
// Solidity: function initializeFromPessimisticConsensus((uint256,bytes32,bytes32,uint256,uint256,uint256,address,bytes32,bytes32) _initParams, bool _useDefaultVkeys, bool _useDefaultSigners, bytes32 _initOwnedAggchainVKey, bytes4 _initAggchainVKeySelector, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepTransactorSession) InitializeFromPessimisticConsensus(_initParams AggchainFEPInitParams, _useDefaultVkeys bool, _useDefaultSigners bool, _initOwnedAggchainVKey [32]byte, _initAggchainVKeySelector [4]byte, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.InitializeFromPessimisticConsensus(&_Aggchainfep.TransactOpts, _initParams, _useDefaultVkeys, _useDefaultSigners, _initOwnedAggchainVKey, _initAggchainVKeySelector, _signersToAdd, _newThreshold)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfep *AggchainfepTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfep *AggchainfepSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.OnVerifyPessimistic(&_Aggchainfep.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfep *AggchainfepTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.OnVerifyPessimistic(&_Aggchainfep.TransactOpts, aggchainData)
}

// SelectOpSuccinctConfig is a paid mutator transaction binding the contract method 0x52076aca.
//
// Solidity: function selectOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepTransactor) SelectOpSuccinctConfig(opts *bind.TransactOpts, _configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "selectOpSuccinctConfig", _configName)
}

// SelectOpSuccinctConfig is a paid mutator transaction binding the contract method 0x52076aca.
//
// Solidity: function selectOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepSession) SelectOpSuccinctConfig(_configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SelectOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName)
}

// SelectOpSuccinctConfig is a paid mutator transaction binding the contract method 0x52076aca.
//
// Solidity: function selectOpSuccinctConfig(bytes32 _configName) returns()
func (_Aggchainfep *AggchainfepTransactorSession) SelectOpSuccinctConfig(_configName [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SelectOpSuccinctConfig(&_Aggchainfep.TransactOpts, _configName)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfep *AggchainfepTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfep *AggchainfepSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SetTrustedSequencer(&_Aggchainfep.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfep *AggchainfepTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SetTrustedSequencer(&_Aggchainfep.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfep *AggchainfepTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfep *AggchainfepSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SetTrustedSequencerURL(&_Aggchainfep.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfep *AggchainfepTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfep.Contract.SetTrustedSequencerURL(&_Aggchainfep.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfep *AggchainfepTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfep *AggchainfepSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferAdminRole(&_Aggchainfep.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfep *AggchainfepTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferAdminRole(&_Aggchainfep.TransactOpts, newPendingAdmin)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepTransactor) TransferAggchainManagerRole(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "transferAggchainManagerRole", newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferAggchainManagerRole(&_Aggchainfep.TransactOpts, newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfep *AggchainfepTransactorSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferAggchainManagerRole(&_Aggchainfep.TransactOpts, newAggchainManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfep *AggchainfepTransactor) TransferOptimisticModeManagerRole(opts *bind.TransactOpts, newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "transferOptimisticModeManagerRole", newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfep *AggchainfepSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferOptimisticModeManagerRole(&_Aggchainfep.TransactOpts, newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfep *AggchainfepTransactorSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferOptimisticModeManagerRole(&_Aggchainfep.TransactOpts, newOptimisticModeManager)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateOwnedAggchainVKey(&_Aggchainfep.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateSignersAndThreshold(opts *bind.TransactOpts, _signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateSignersAndThreshold", _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateSignersAndThreshold(&_Aggchainfep.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateSignersAndThreshold(&_Aggchainfep.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateSubmissionInterval(opts *bind.TransactOpts, _submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateSubmissionInterval", _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfep *AggchainfepSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateSubmissionInterval(&_Aggchainfep.TransactOpts, _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateSubmissionInterval(&_Aggchainfep.TransactOpts, _submissionInterval)
}

// UpgradeFromPreviousFEP is a paid mutator transaction binding the contract method 0x96a4f546.
//
// Solidity: function upgradeFromPreviousFEP() returns()
func (_Aggchainfep *AggchainfepTransactor) UpgradeFromPreviousFEP(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "upgradeFromPreviousFEP")
}

// UpgradeFromPreviousFEP is a paid mutator transaction binding the contract method 0x96a4f546.
//
// Solidity: function upgradeFromPreviousFEP() returns()
func (_Aggchainfep *AggchainfepSession) UpgradeFromPreviousFEP() (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpgradeFromPreviousFEP(&_Aggchainfep.TransactOpts)
}

// UpgradeFromPreviousFEP is a paid mutator transaction binding the contract method 0x96a4f546.
//
// Solidity: function upgradeFromPreviousFEP() returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpgradeFromPreviousFEP() (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpgradeFromPreviousFEP(&_Aggchainfep.TransactOpts)
}

// AggchainfepAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Aggchainfep contract.
type AggchainfepAcceptAdminRoleIterator struct {
	Event *AggchainfepAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAcceptAdminRole)
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
		it.Event = new(AggchainfepAcceptAdminRole)
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
func (it *AggchainfepAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAcceptAdminRole represents a AcceptAdminRole event raised by the Aggchainfep contract.
type AggchainfepAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainfep *AggchainfepFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainfepAcceptAdminRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAcceptAdminRoleIterator{contract: _Aggchainfep.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainfep *AggchainfepFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainfepAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAcceptAdminRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainfep *AggchainfepFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainfepAcceptAdminRole, error) {
	event := new(AggchainfepAcceptAdminRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Aggchainfep contract.
type AggchainfepAcceptAggchainManagerRoleIterator struct {
	Event *AggchainfepAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAcceptAggchainManagerRole)
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
		it.Event = new(AggchainfepAcceptAggchainManagerRole)
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
func (it *AggchainfepAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Aggchainfep contract.
type AggchainfepAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*AggchainfepAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAcceptAggchainManagerRoleIterator{contract: _Aggchainfep.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAcceptAggchainManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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

// ParseAcceptAggchainManagerRole is a log parse operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*AggchainfepAcceptAggchainManagerRole, error) {
	event := new(AggchainfepAcceptAggchainManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepAcceptOptimisticModeManagerRoleIterator is returned from FilterAcceptOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for AcceptOptimisticModeManagerRole events raised by the Aggchainfep contract.
type AggchainfepAcceptOptimisticModeManagerRoleIterator struct {
	Event *AggchainfepAcceptOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepAcceptOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAcceptOptimisticModeManagerRole)
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
		it.Event = new(AggchainfepAcceptOptimisticModeManagerRole)
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
func (it *AggchainfepAcceptOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAcceptOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAcceptOptimisticModeManagerRole represents a AcceptOptimisticModeManagerRole event raised by the Aggchainfep contract.
type AggchainfepAcceptOptimisticModeManagerRole struct {
	OldOptimisticModeManager common.Address
	NewOptimisticModeManager common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) FilterAcceptOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainfepAcceptOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAcceptOptimisticModeManagerRoleIterator{contract: _Aggchainfep.contract, event: "AcceptOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptOptimisticModeManagerRole is a free log subscription operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) WatchAcceptOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepAcceptOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAcceptOptimisticModeManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
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

// ParseAcceptOptimisticModeManagerRole is a log parse operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) ParseAcceptOptimisticModeManagerRole(log types.Log) (*AggchainfepAcceptOptimisticModeManagerRole, error) {
	event := new(AggchainfepAcceptOptimisticModeManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Aggchainfep contract.
type AggchainfepAcceptVKeyManagerRoleIterator struct {
	Event *AggchainfepAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAcceptVKeyManagerRole)
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
		it.Event = new(AggchainfepAcceptVKeyManagerRole)
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
func (it *AggchainfepAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Aggchainfep contract.
type AggchainfepAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfepAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAcceptVKeyManagerRoleIterator{contract: _Aggchainfep.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAcceptVKeyManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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

// ParseAcceptVKeyManagerRole is a log parse operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*AggchainfepAcceptVKeyManagerRole, error) {
	event := new(AggchainfepAcceptVKeyManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Aggchainfep contract.
type AggchainfepAddAggchainVKeyIterator struct {
	Event *AggchainfepAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainfepAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAddAggchainVKey)
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
		it.Event = new(AggchainfepAddAggchainVKey)
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
func (it *AggchainfepAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAddAggchainVKey represents a AddAggchainVKey event raised by the Aggchainfep contract.
type AggchainfepAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainfepAddAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainfepAddAggchainVKeyIterator{contract: _Aggchainfep.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainfepAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAddAggchainVKey)
				if err := _Aggchainfep.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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

// ParseAddAggchainVKey is a log parse operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainfepAddAggchainVKey, error) {
	event := new(AggchainfepAddAggchainVKey)
	if err := _Aggchainfep.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepAggregationVkeyUpdatedIterator is returned from FilterAggregationVkeyUpdated and is used to iterate over the raw logs and unpacked data for AggregationVkeyUpdated events raised by the Aggchainfep contract.
type AggchainfepAggregationVkeyUpdatedIterator struct {
	Event *AggchainfepAggregationVkeyUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepAggregationVkeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepAggregationVkeyUpdated)
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
		it.Event = new(AggchainfepAggregationVkeyUpdated)
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
func (it *AggchainfepAggregationVkeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepAggregationVkeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepAggregationVkeyUpdated represents a AggregationVkeyUpdated event raised by the Aggchainfep contract.
type AggchainfepAggregationVkeyUpdated struct {
	OldAggregationVkey [32]byte
	NewAggregationVkey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAggregationVkeyUpdated is a free log retrieval operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_Aggchainfep *AggchainfepFilterer) FilterAggregationVkeyUpdated(opts *bind.FilterOpts, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (*AggchainfepAggregationVkeyUpdatedIterator, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepAggregationVkeyUpdatedIterator{contract: _Aggchainfep.contract, event: "AggregationVkeyUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregationVkeyUpdated is a free log subscription operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_Aggchainfep *AggchainfepFilterer) WatchAggregationVkeyUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepAggregationVkeyUpdated, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (event.Subscription, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepAggregationVkeyUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
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

// ParseAggregationVkeyUpdated is a log parse operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_Aggchainfep *AggchainfepFilterer) ParseAggregationVkeyUpdated(log types.Log) (*AggchainfepAggregationVkeyUpdated, error) {
	event := new(AggchainfepAggregationVkeyUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepDisableOptimisticModeIterator is returned from FilterDisableOptimisticMode and is used to iterate over the raw logs and unpacked data for DisableOptimisticMode events raised by the Aggchainfep contract.
type AggchainfepDisableOptimisticModeIterator struct {
	Event *AggchainfepDisableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainfepDisableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepDisableOptimisticMode)
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
		it.Event = new(AggchainfepDisableOptimisticMode)
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
func (it *AggchainfepDisableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepDisableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepDisableOptimisticMode represents a DisableOptimisticMode event raised by the Aggchainfep contract.
type AggchainfepDisableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableOptimisticMode is a free log retrieval operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) FilterDisableOptimisticMode(opts *bind.FilterOpts) (*AggchainfepDisableOptimisticModeIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainfepDisableOptimisticModeIterator{contract: _Aggchainfep.contract, event: "DisableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchDisableOptimisticMode is a free log subscription operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) WatchDisableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainfepDisableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepDisableOptimisticMode)
				if err := _Aggchainfep.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
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

// ParseDisableOptimisticMode is a log parse operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) ParseDisableOptimisticMode(log types.Log) (*AggchainfepDisableOptimisticMode, error) {
	event := new(AggchainfepDisableOptimisticMode)
	if err := _Aggchainfep.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepDisableUseDefaultSignersFlagIterator is returned from FilterDisableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultSignersFlag events raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultSignersFlagIterator struct {
	Event *AggchainfepDisableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepDisableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepDisableUseDefaultSignersFlag)
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
		it.Event = new(AggchainfepDisableUseDefaultSignersFlag)
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
func (it *AggchainfepDisableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepDisableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepDisableUseDefaultSignersFlag represents a DisableUseDefaultSignersFlag event raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterDisableUseDefaultSignersFlag(opts *bind.FilterOpts) (*AggchainfepDisableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepDisableUseDefaultSignersFlagIterator{contract: _Aggchainfep.contract, event: "DisableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchDisableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepDisableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepDisableUseDefaultSignersFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
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

// ParseDisableUseDefaultSignersFlag is a log parse operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) ParseDisableUseDefaultSignersFlag(log types.Log) (*AggchainfepDisableUseDefaultSignersFlag, error) {
	event := new(AggchainfepDisableUseDefaultSignersFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepDisableUseDefaultVkeysFlagIterator is returned from FilterDisableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultVkeysFlag events raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultVkeysFlagIterator struct {
	Event *AggchainfepDisableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepDisableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepDisableUseDefaultVkeysFlag)
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
		it.Event = new(AggchainfepDisableUseDefaultVkeysFlag)
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
func (it *AggchainfepDisableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepDisableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepDisableUseDefaultVkeysFlag represents a DisableUseDefaultVkeysFlag event raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterDisableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*AggchainfepDisableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepDisableUseDefaultVkeysFlagIterator{contract: _Aggchainfep.contract, event: "DisableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchDisableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepDisableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepDisableUseDefaultVkeysFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
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

// ParseDisableUseDefaultVkeysFlag is a log parse operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) ParseDisableUseDefaultVkeysFlag(log types.Log) (*AggchainfepDisableUseDefaultVkeysFlag, error) {
	event := new(AggchainfepDisableUseDefaultVkeysFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepEnableOptimisticModeIterator is returned from FilterEnableOptimisticMode and is used to iterate over the raw logs and unpacked data for EnableOptimisticMode events raised by the Aggchainfep contract.
type AggchainfepEnableOptimisticModeIterator struct {
	Event *AggchainfepEnableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainfepEnableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepEnableOptimisticMode)
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
		it.Event = new(AggchainfepEnableOptimisticMode)
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
func (it *AggchainfepEnableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepEnableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepEnableOptimisticMode represents a EnableOptimisticMode event raised by the Aggchainfep contract.
type AggchainfepEnableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableOptimisticMode is a free log retrieval operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) FilterEnableOptimisticMode(opts *bind.FilterOpts) (*AggchainfepEnableOptimisticModeIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainfepEnableOptimisticModeIterator{contract: _Aggchainfep.contract, event: "EnableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchEnableOptimisticMode is a free log subscription operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) WatchEnableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainfepEnableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepEnableOptimisticMode)
				if err := _Aggchainfep.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
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

// ParseEnableOptimisticMode is a log parse operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_Aggchainfep *AggchainfepFilterer) ParseEnableOptimisticMode(log types.Log) (*AggchainfepEnableOptimisticMode, error) {
	event := new(AggchainfepEnableOptimisticMode)
	if err := _Aggchainfep.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepEnableUseDefaultSignersFlagIterator is returned from FilterEnableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultSignersFlag events raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultSignersFlagIterator struct {
	Event *AggchainfepEnableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepEnableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepEnableUseDefaultSignersFlag)
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
		it.Event = new(AggchainfepEnableUseDefaultSignersFlag)
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
func (it *AggchainfepEnableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepEnableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepEnableUseDefaultSignersFlag represents a EnableUseDefaultSignersFlag event raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterEnableUseDefaultSignersFlag(opts *bind.FilterOpts) (*AggchainfepEnableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepEnableUseDefaultSignersFlagIterator{contract: _Aggchainfep.contract, event: "EnableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchEnableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepEnableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepEnableUseDefaultSignersFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
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

// ParseEnableUseDefaultSignersFlag is a log parse operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Aggchainfep *AggchainfepFilterer) ParseEnableUseDefaultSignersFlag(log types.Log) (*AggchainfepEnableUseDefaultSignersFlag, error) {
	event := new(AggchainfepEnableUseDefaultSignersFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepEnableUseDefaultVkeysFlagIterator is returned from FilterEnableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultVkeysFlag events raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultVkeysFlagIterator struct {
	Event *AggchainfepEnableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepEnableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepEnableUseDefaultVkeysFlag)
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
		it.Event = new(AggchainfepEnableUseDefaultVkeysFlag)
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
func (it *AggchainfepEnableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepEnableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepEnableUseDefaultVkeysFlag represents a EnableUseDefaultVkeysFlag event raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterEnableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*AggchainfepEnableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepEnableUseDefaultVkeysFlagIterator{contract: _Aggchainfep.contract, event: "EnableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchEnableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepEnableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepEnableUseDefaultVkeysFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
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

// ParseEnableUseDefaultVkeysFlag is a log parse operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Aggchainfep *AggchainfepFilterer) ParseEnableUseDefaultVkeysFlag(log types.Log) (*AggchainfepEnableUseDefaultVkeysFlag, error) {
	event := new(AggchainfepEnableUseDefaultVkeysFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggchainfep contract.
type AggchainfepInitializedIterator struct {
	Event *AggchainfepInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainfepInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepInitialized)
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
		it.Event = new(AggchainfepInitialized)
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
func (it *AggchainfepInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepInitialized represents a Initialized event raised by the Aggchainfep contract.
type AggchainfepInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainfep *AggchainfepFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainfepInitializedIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainfepInitializedIterator{contract: _Aggchainfep.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainfep *AggchainfepFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainfepInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepInitialized)
				if err := _Aggchainfep.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggchainfep *AggchainfepFilterer) ParseInitialized(log types.Log) (*AggchainfepInitialized, error) {
	event := new(AggchainfepInitialized)
	if err := _Aggchainfep.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepOpSuccinctConfigDeletedIterator is returned from FilterOpSuccinctConfigDeleted and is used to iterate over the raw logs and unpacked data for OpSuccinctConfigDeleted events raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigDeletedIterator struct {
	Event *AggchainfepOpSuccinctConfigDeleted // Event containing the contract specifics and raw log

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
func (it *AggchainfepOpSuccinctConfigDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepOpSuccinctConfigDeleted)
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
		it.Event = new(AggchainfepOpSuccinctConfigDeleted)
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
func (it *AggchainfepOpSuccinctConfigDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepOpSuccinctConfigDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepOpSuccinctConfigDeleted represents a OpSuccinctConfigDeleted event raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigDeleted struct {
	ConfigName [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOpSuccinctConfigDeleted is a free log retrieval operation binding the contract event 0x4432b02a2fcbed48d94e8d72723e155c6690e4b7f39afa41a2a8ff8c0aa425da.
//
// Solidity: event OpSuccinctConfigDeleted(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) FilterOpSuccinctConfigDeleted(opts *bind.FilterOpts, configName [][32]byte) (*AggchainfepOpSuccinctConfigDeletedIterator, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "OpSuccinctConfigDeleted", configNameRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepOpSuccinctConfigDeletedIterator{contract: _Aggchainfep.contract, event: "OpSuccinctConfigDeleted", logs: logs, sub: sub}, nil
}

// WatchOpSuccinctConfigDeleted is a free log subscription operation binding the contract event 0x4432b02a2fcbed48d94e8d72723e155c6690e4b7f39afa41a2a8ff8c0aa425da.
//
// Solidity: event OpSuccinctConfigDeleted(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) WatchOpSuccinctConfigDeleted(opts *bind.WatchOpts, sink chan<- *AggchainfepOpSuccinctConfigDeleted, configName [][32]byte) (event.Subscription, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "OpSuccinctConfigDeleted", configNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepOpSuccinctConfigDeleted)
				if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigDeleted", log); err != nil {
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

// ParseOpSuccinctConfigDeleted is a log parse operation binding the contract event 0x4432b02a2fcbed48d94e8d72723e155c6690e4b7f39afa41a2a8ff8c0aa425da.
//
// Solidity: event OpSuccinctConfigDeleted(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) ParseOpSuccinctConfigDeleted(log types.Log) (*AggchainfepOpSuccinctConfigDeleted, error) {
	event := new(AggchainfepOpSuccinctConfigDeleted)
	if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepOpSuccinctConfigSelectedIterator is returned from FilterOpSuccinctConfigSelected and is used to iterate over the raw logs and unpacked data for OpSuccinctConfigSelected events raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigSelectedIterator struct {
	Event *AggchainfepOpSuccinctConfigSelected // Event containing the contract specifics and raw log

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
func (it *AggchainfepOpSuccinctConfigSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepOpSuccinctConfigSelected)
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
		it.Event = new(AggchainfepOpSuccinctConfigSelected)
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
func (it *AggchainfepOpSuccinctConfigSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepOpSuccinctConfigSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepOpSuccinctConfigSelected represents a OpSuccinctConfigSelected event raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigSelected struct {
	ConfigName [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOpSuccinctConfigSelected is a free log retrieval operation binding the contract event 0x2a2ab116b5abc962503c3c7f941af94e3dc855231d07abb9bc4dc2105591a031.
//
// Solidity: event OpSuccinctConfigSelected(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) FilterOpSuccinctConfigSelected(opts *bind.FilterOpts, configName [][32]byte) (*AggchainfepOpSuccinctConfigSelectedIterator, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "OpSuccinctConfigSelected", configNameRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepOpSuccinctConfigSelectedIterator{contract: _Aggchainfep.contract, event: "OpSuccinctConfigSelected", logs: logs, sub: sub}, nil
}

// WatchOpSuccinctConfigSelected is a free log subscription operation binding the contract event 0x2a2ab116b5abc962503c3c7f941af94e3dc855231d07abb9bc4dc2105591a031.
//
// Solidity: event OpSuccinctConfigSelected(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) WatchOpSuccinctConfigSelected(opts *bind.WatchOpts, sink chan<- *AggchainfepOpSuccinctConfigSelected, configName [][32]byte) (event.Subscription, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "OpSuccinctConfigSelected", configNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepOpSuccinctConfigSelected)
				if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigSelected", log); err != nil {
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

// ParseOpSuccinctConfigSelected is a log parse operation binding the contract event 0x2a2ab116b5abc962503c3c7f941af94e3dc855231d07abb9bc4dc2105591a031.
//
// Solidity: event OpSuccinctConfigSelected(bytes32 indexed configName)
func (_Aggchainfep *AggchainfepFilterer) ParseOpSuccinctConfigSelected(log types.Log) (*AggchainfepOpSuccinctConfigSelected, error) {
	event := new(AggchainfepOpSuccinctConfigSelected)
	if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepOpSuccinctConfigUpdatedIterator is returned from FilterOpSuccinctConfigUpdated and is used to iterate over the raw logs and unpacked data for OpSuccinctConfigUpdated events raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigUpdatedIterator struct {
	Event *AggchainfepOpSuccinctConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepOpSuccinctConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepOpSuccinctConfigUpdated)
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
		it.Event = new(AggchainfepOpSuccinctConfigUpdated)
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
func (it *AggchainfepOpSuccinctConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepOpSuccinctConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepOpSuccinctConfigUpdated represents a OpSuccinctConfigUpdated event raised by the Aggchainfep contract.
type AggchainfepOpSuccinctConfigUpdated struct {
	ConfigName          [32]byte
	AggregationVkey     [32]byte
	RangeVkeyCommitment [32]byte
	RollupConfigHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOpSuccinctConfigUpdated is a free log retrieval operation binding the contract event 0xea0123c726a665cb0ab5691444f929a7056c7a7709c60c0587829e8046b8d514.
//
// Solidity: event OpSuccinctConfigUpdated(bytes32 indexed configName, bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) FilterOpSuccinctConfigUpdated(opts *bind.FilterOpts, configName [][32]byte) (*AggchainfepOpSuccinctConfigUpdatedIterator, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "OpSuccinctConfigUpdated", configNameRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepOpSuccinctConfigUpdatedIterator{contract: _Aggchainfep.contract, event: "OpSuccinctConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchOpSuccinctConfigUpdated is a free log subscription operation binding the contract event 0xea0123c726a665cb0ab5691444f929a7056c7a7709c60c0587829e8046b8d514.
//
// Solidity: event OpSuccinctConfigUpdated(bytes32 indexed configName, bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) WatchOpSuccinctConfigUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepOpSuccinctConfigUpdated, configName [][32]byte) (event.Subscription, error) {

	var configNameRule []interface{}
	for _, configNameItem := range configName {
		configNameRule = append(configNameRule, configNameItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "OpSuccinctConfigUpdated", configNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepOpSuccinctConfigUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigUpdated", log); err != nil {
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

// ParseOpSuccinctConfigUpdated is a log parse operation binding the contract event 0xea0123c726a665cb0ab5691444f929a7056c7a7709c60c0587829e8046b8d514.
//
// Solidity: event OpSuccinctConfigUpdated(bytes32 indexed configName, bytes32 aggregationVkey, bytes32 rangeVkeyCommitment, bytes32 rollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) ParseOpSuccinctConfigUpdated(log types.Log) (*AggchainfepOpSuccinctConfigUpdated, error) {
	event := new(AggchainfepOpSuccinctConfigUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "OpSuccinctConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the Aggchainfep contract.
type AggchainfepOutputProposedIterator struct {
	Event *AggchainfepOutputProposed // Event containing the contract specifics and raw log

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
func (it *AggchainfepOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepOutputProposed)
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
		it.Event = new(AggchainfepOutputProposed)
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
func (it *AggchainfepOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepOutputProposed represents a OutputProposed event raised by the Aggchainfep contract.
type AggchainfepOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_Aggchainfep *AggchainfepFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*AggchainfepOutputProposedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepOutputProposedIterator{contract: _Aggchainfep.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_Aggchainfep *AggchainfepFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *AggchainfepOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepOutputProposed)
				if err := _Aggchainfep.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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

// ParseOutputProposed is a log parse operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_Aggchainfep *AggchainfepFilterer) ParseOutputProposed(log types.Log) (*AggchainfepOutputProposed, error) {
	event := new(AggchainfepOutputProposed)
	if err := _Aggchainfep.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepRangeVkeyCommitmentUpdatedIterator is returned from FilterRangeVkeyCommitmentUpdated and is used to iterate over the raw logs and unpacked data for RangeVkeyCommitmentUpdated events raised by the Aggchainfep contract.
type AggchainfepRangeVkeyCommitmentUpdatedIterator struct {
	Event *AggchainfepRangeVkeyCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepRangeVkeyCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepRangeVkeyCommitmentUpdated)
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
		it.Event = new(AggchainfepRangeVkeyCommitmentUpdated)
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
func (it *AggchainfepRangeVkeyCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepRangeVkeyCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepRangeVkeyCommitmentUpdated represents a RangeVkeyCommitmentUpdated event raised by the Aggchainfep contract.
type AggchainfepRangeVkeyCommitmentUpdated struct {
	OldRangeVkeyCommitment [32]byte
	NewRangeVkeyCommitment [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRangeVkeyCommitmentUpdated is a free log retrieval operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_Aggchainfep *AggchainfepFilterer) FilterRangeVkeyCommitmentUpdated(opts *bind.FilterOpts, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (*AggchainfepRangeVkeyCommitmentUpdatedIterator, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepRangeVkeyCommitmentUpdatedIterator{contract: _Aggchainfep.contract, event: "RangeVkeyCommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchRangeVkeyCommitmentUpdated is a free log subscription operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_Aggchainfep *AggchainfepFilterer) WatchRangeVkeyCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepRangeVkeyCommitmentUpdated, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (event.Subscription, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepRangeVkeyCommitmentUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
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

// ParseRangeVkeyCommitmentUpdated is a log parse operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_Aggchainfep *AggchainfepFilterer) ParseRangeVkeyCommitmentUpdated(log types.Log) (*AggchainfepRangeVkeyCommitmentUpdated, error) {
	event := new(AggchainfepRangeVkeyCommitmentUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepRollupConfigHashUpdatedIterator is returned from FilterRollupConfigHashUpdated and is used to iterate over the raw logs and unpacked data for RollupConfigHashUpdated events raised by the Aggchainfep contract.
type AggchainfepRollupConfigHashUpdatedIterator struct {
	Event *AggchainfepRollupConfigHashUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepRollupConfigHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepRollupConfigHashUpdated)
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
		it.Event = new(AggchainfepRollupConfigHashUpdated)
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
func (it *AggchainfepRollupConfigHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepRollupConfigHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepRollupConfigHashUpdated represents a RollupConfigHashUpdated event raised by the Aggchainfep contract.
type AggchainfepRollupConfigHashUpdated struct {
	OldRollupConfigHash [32]byte
	NewRollupConfigHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollupConfigHashUpdated is a free log retrieval operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) FilterRollupConfigHashUpdated(opts *bind.FilterOpts, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (*AggchainfepRollupConfigHashUpdatedIterator, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfepRollupConfigHashUpdatedIterator{contract: _Aggchainfep.contract, event: "RollupConfigHashUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupConfigHashUpdated is a free log subscription operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) WatchRollupConfigHashUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepRollupConfigHashUpdated, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (event.Subscription, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepRollupConfigHashUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
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

// ParseRollupConfigHashUpdated is a log parse operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_Aggchainfep *AggchainfepFilterer) ParseRollupConfigHashUpdated(log types.Log) (*AggchainfepRollupConfigHashUpdated, error) {
	event := new(AggchainfepRollupConfigHashUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Aggchainfep contract.
type AggchainfepSetTrustedSequencerIterator struct {
	Event *AggchainfepSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainfepSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepSetTrustedSequencer)
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
		it.Event = new(AggchainfepSetTrustedSequencer)
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
func (it *AggchainfepSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepSetTrustedSequencer represents a SetTrustedSequencer event raised by the Aggchainfep contract.
type AggchainfepSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainfep *AggchainfepFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainfepSetTrustedSequencerIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainfepSetTrustedSequencerIterator{contract: _Aggchainfep.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainfep *AggchainfepFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainfepSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepSetTrustedSequencer)
				if err := _Aggchainfep.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainfep *AggchainfepFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainfepSetTrustedSequencer, error) {
	event := new(AggchainfepSetTrustedSequencer)
	if err := _Aggchainfep.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Aggchainfep contract.
type AggchainfepSetTrustedSequencerURLIterator struct {
	Event *AggchainfepSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainfepSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepSetTrustedSequencerURL)
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
		it.Event = new(AggchainfepSetTrustedSequencerURL)
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
func (it *AggchainfepSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Aggchainfep contract.
type AggchainfepSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainfep *AggchainfepFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainfepSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainfepSetTrustedSequencerURLIterator{contract: _Aggchainfep.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainfep *AggchainfepFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainfepSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepSetTrustedSequencerURL)
				if err := _Aggchainfep.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainfep *AggchainfepFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainfepSetTrustedSequencerURL, error) {
	event := new(AggchainfepSetTrustedSequencerURL)
	if err := _Aggchainfep.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepSignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Aggchainfep contract.
type AggchainfepSignersAndThresholdUpdatedIterator struct {
	Event *AggchainfepSignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepSignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepSignersAndThresholdUpdated)
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
		it.Event = new(AggchainfepSignersAndThresholdUpdated)
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
func (it *AggchainfepSignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepSignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepSignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Aggchainfep contract.
type AggchainfepSignersAndThresholdUpdated struct {
	AggchainSigners        []common.Address
	NewThreshold           *big.Int
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainfep *AggchainfepFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*AggchainfepSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainfepSignersAndThresholdUpdatedIterator{contract: _Aggchainfep.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainfep *AggchainfepFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepSignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepSignersAndThresholdUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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

// ParseSignersAndThresholdUpdated is a log parse operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainfep *AggchainfepFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*AggchainfepSignersAndThresholdUpdated, error) {
	event := new(AggchainfepSignersAndThresholdUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepSubmissionIntervalUpdatedIterator is returned from FilterSubmissionIntervalUpdated and is used to iterate over the raw logs and unpacked data for SubmissionIntervalUpdated events raised by the Aggchainfep contract.
type AggchainfepSubmissionIntervalUpdatedIterator struct {
	Event *AggchainfepSubmissionIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfepSubmissionIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepSubmissionIntervalUpdated)
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
		it.Event = new(AggchainfepSubmissionIntervalUpdated)
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
func (it *AggchainfepSubmissionIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepSubmissionIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepSubmissionIntervalUpdated represents a SubmissionIntervalUpdated event raised by the Aggchainfep contract.
type AggchainfepSubmissionIntervalUpdated struct {
	OldSubmissionInterval *big.Int
	NewSubmissionInterval *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubmissionIntervalUpdated is a free log retrieval operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_Aggchainfep *AggchainfepFilterer) FilterSubmissionIntervalUpdated(opts *bind.FilterOpts) (*AggchainfepSubmissionIntervalUpdatedIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainfepSubmissionIntervalUpdatedIterator{contract: _Aggchainfep.contract, event: "SubmissionIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchSubmissionIntervalUpdated is a free log subscription operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_Aggchainfep *AggchainfepFilterer) WatchSubmissionIntervalUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfepSubmissionIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepSubmissionIntervalUpdated)
				if err := _Aggchainfep.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
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

// ParseSubmissionIntervalUpdated is a log parse operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_Aggchainfep *AggchainfepFilterer) ParseSubmissionIntervalUpdated(log types.Log) (*AggchainfepSubmissionIntervalUpdated, error) {
	event := new(AggchainfepSubmissionIntervalUpdated)
	if err := _Aggchainfep.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Aggchainfep contract.
type AggchainfepTransferAdminRoleIterator struct {
	Event *AggchainfepTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepTransferAdminRole)
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
		it.Event = new(AggchainfepTransferAdminRole)
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
func (it *AggchainfepTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepTransferAdminRole represents a TransferAdminRole event raised by the Aggchainfep contract.
type AggchainfepTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainfep *AggchainfepFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainfepTransferAdminRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransferAdminRoleIterator{contract: _Aggchainfep.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainfep *AggchainfepFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainfepTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepTransferAdminRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainfep *AggchainfepFilterer) ParseTransferAdminRole(log types.Log) (*AggchainfepTransferAdminRole, error) {
	event := new(AggchainfepTransferAdminRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Aggchainfep contract.
type AggchainfepTransferAggchainManagerRoleIterator struct {
	Event *AggchainfepTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepTransferAggchainManagerRole)
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
		it.Event = new(AggchainfepTransferAggchainManagerRole)
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
func (it *AggchainfepTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Aggchainfep contract.
type AggchainfepTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*AggchainfepTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransferAggchainManagerRoleIterator{contract: _Aggchainfep.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepTransferAggchainManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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

// ParseTransferAggchainManagerRole is a log parse operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainfep *AggchainfepFilterer) ParseTransferAggchainManagerRole(log types.Log) (*AggchainfepTransferAggchainManagerRole, error) {
	event := new(AggchainfepTransferAggchainManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepTransferOptimisticModeManagerRoleIterator is returned from FilterTransferOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for TransferOptimisticModeManagerRole events raised by the Aggchainfep contract.
type AggchainfepTransferOptimisticModeManagerRoleIterator struct {
	Event *AggchainfepTransferOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepTransferOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepTransferOptimisticModeManagerRole)
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
		it.Event = new(AggchainfepTransferOptimisticModeManagerRole)
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
func (it *AggchainfepTransferOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepTransferOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepTransferOptimisticModeManagerRole represents a TransferOptimisticModeManagerRole event raised by the Aggchainfep contract.
type AggchainfepTransferOptimisticModeManagerRole struct {
	CurrentOptimisticModeManager    common.Address
	NewPendingOptimisticModeManager common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTransferOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) FilterTransferOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainfepTransferOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransferOptimisticModeManagerRoleIterator{contract: _Aggchainfep.contract, event: "TransferOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferOptimisticModeManagerRole is a free log subscription operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) WatchTransferOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepTransferOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepTransferOptimisticModeManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
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

// ParseTransferOptimisticModeManagerRole is a log parse operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_Aggchainfep *AggchainfepFilterer) ParseTransferOptimisticModeManagerRole(log types.Log) (*AggchainfepTransferOptimisticModeManagerRole, error) {
	event := new(AggchainfepTransferOptimisticModeManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Aggchainfep contract.
type AggchainfepTransferVKeyManagerRoleIterator struct {
	Event *AggchainfepTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfepTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepTransferVKeyManagerRole)
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
		it.Event = new(AggchainfepTransferVKeyManagerRole)
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
func (it *AggchainfepTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Aggchainfep contract.
type AggchainfepTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfepTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfepTransferVKeyManagerRoleIterator{contract: _Aggchainfep.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfepTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepTransferVKeyManagerRole)
				if err := _Aggchainfep.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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

// ParseTransferVKeyManagerRole is a log parse operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainfep *AggchainfepFilterer) ParseTransferVKeyManagerRole(log types.Log) (*AggchainfepTransferVKeyManagerRole, error) {
	event := new(AggchainfepTransferVKeyManagerRole)
	if err := _Aggchainfep.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfepUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Aggchainfep contract.
type AggchainfepUpdateAggchainVKeyIterator struct {
	Event *AggchainfepUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainfepUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepUpdateAggchainVKey)
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
		it.Event = new(AggchainfepUpdateAggchainVKey)
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
func (it *AggchainfepUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Aggchainfep contract.
type AggchainfepUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainfepUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainfepUpdateAggchainVKeyIterator{contract: _Aggchainfep.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainfepUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepUpdateAggchainVKey)
				if err := _Aggchainfep.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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

// ParseUpdateAggchainVKey is a log parse operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainfep *AggchainfepFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainfepUpdateAggchainVKey, error) {
	event := new(AggchainfepUpdateAggchainVKey)
	if err := _Aggchainfep.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

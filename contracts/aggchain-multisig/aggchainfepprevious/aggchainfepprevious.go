// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggchainfepprevious

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

// AggchainFEPPreviousOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPPreviousOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// AggchainfeppreviousMetaData contains all meta data concerning the Aggchainfepprevious contract.
var AggchainfeppreviousMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGatewayPrevious\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdminCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggregationVkeyMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotProposeFutureL2Output\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockNumberLessThanNextBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockTimeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2OutputRootCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RangeVkeyCommitmentMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupConfigHashMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartL2TimestampMustBeLessThanCurrentTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmissionIntervalMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"AcceptOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldAggregationVkey\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAggregationVkey\",\"type\":\"bytes32\"}],\"name\":\"AggregationVkeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRangeVkeyCommitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"RangeVkeyCommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRollupConfigHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfigHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSubmissionInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubmissionInterval\",\"type\":\"uint256\"}],\"name\":\"SubmissionIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"TransferOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_FEP_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGatewayPrevious\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregationVkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structAggchainFEPPrevious.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOptimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeVkeyCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"transferOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_aggregationVkey\",\"type\":\"bytes32\"}],\"name\":\"updateAggregationVkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rangeVkeyCommitment\",\"type\":\"bytes32\"}],\"name\":\"updateRangeVkeyCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"updateRollupConfigHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051613e90380380613e9083398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e05161010051613c096102875f395f8181610a380152610d8e01525f818161080401528181612102015261241801525f610a1101525f610b6a01525f610bdd0152613c095ff3fe608060405234801561000f575f5ffd5b50600436106104dc575f3560e01c806370872aa511610291578063c4cb03ec11610162578063e1a41bcf116100dd578063e90a340911610093578063f851a44011610079578063f851a44014610c4b578063fdbbc19b14610c70578063ff90407914610c83575f5ffd5b8063e90a340914610c1b578063effb847914610c2c575f5ffd5b8063e46761c4116100c3578063e46761c414610bd8578063e631476c14610bff578063e7a7ed0214610c07575f5ffd5b8063e1a41bcf14610baf578063e279984e14610bb8575f5ffd5b8063cfa8ed4711610132578063d1de856c11610118578063d1de856c14610b8c578063dc8c424914610b9f578063dcec334814610ba7575f5ffd5b8063cfa8ed4714610b45578063d02103ca14610b65575f5ffd5b8063c4cb03ec14610aef578063c754c7ed14610b02578063c89e42df14610b2a578063cea5a4c014610b3d575f5ffd5b8063a25ae5571161020c578063b3a326f7116101c2578063bdfbed7e116101a8578063bdfbed7e14610ab3578063bfb193b614610ac6578063c32e4e3e14610ae6575f5ffd5b8063b3a326f714610a8d578063bc91ce3314610aa0575f5ffd5b8063ab0475cf116101f2578063ab0475cf14610a33578063ada8f91914610a5a578063adb8696c14610a6d575f5ffd5b8063a25ae557146109bd578063a3c573eb14610a0c575f5ffd5b806385018182116102615780638c3d7301116102475780638c3d73011461099957806393991af3146109a15780639ee4afa3146109aa575f5ffd5b8063850181821461097d5780638878627214610990575f5ffd5b806370872aa51461093957806371257022146109425780637388c4361461095557806381eb0baf14610975575f5ffd5b80633c351e10116103cb57806354fd4d50116103465780636b8616ce116102fc5780636e05d2cd116102e25780636e05d2cd146108f75780636e7fbce9146109005780636ff512cc14610926575f5ffd5b80636b8616ce146108cf5780636d9a1c8b146108ee575f5ffd5b806369f16eec1161032c57806369f16eec146108ac5780636a55f66c146108b45780636abcf563146108c7575f5ffd5b806354fd4d501461085657806360caf7a01461088f575f5ffd5b80634599c7881161039b578063527570f111610381578063527570f114610826578063529933df14610846578063542028d51461084e575f5ffd5b80634599c788146107f757806349b7b802146107ff575f5ffd5b80633c351e101461074e5780633cbc795b1461076e578063439fab91146107ab57806345605267146107be575f5ffd5b80631d0b435e1161045b5780632c111c061161042b578063336c9e8111610411578063336c9e81146106f7578063368c822c1461070a57806337d4d03014610712575f5ffd5b80632c111c06146106c4578063314eb17b146106e4575f5ffd5b80631d0b435e146105a9578063267822471461063457806326f9b76d146106545780632b31841e146106bb575f5ffd5b806312634900116104b057806319451a8f1161049657806319451a8f146105395780631bdd450c1461054c5780631cf810ee1461055f575f5ffd5b8063126349001461052957806315981b2914610531575f5ffd5b80622134cc146104e057806301fcf6a0146104f75780630822dc611461050a578063107bf28c14610514575b5f5ffd5b6078545b6040519081526020015b60405180910390f35b6104e461050536600461330a565b610ca8565b610512610e12565b005b61051c610ef6565b6040516104ee919061332c565b610512610f82565b610512611091565b61051261054736600461337f565b611168565b61051261055a3660046133a9565b6112c9565b607c5461058490610100900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016104ee565b6106036105b73660046133ef565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016104ee565b6001546105849073ffffffffffffffffffffffffffffffffffffffff1681565b61068a61066236600461330a565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016104ee565b6104e4607a5481565b6008546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6105126106f236600461337f565b611384565b6105126107053660046133a9565b6114ba565b610512611585565b61051c6040518060400160405280600f81526020017f76322e302e302d70726576696f7573000000000000000000000000000000000081525081565b6009546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6009546107969074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016104ee565b6105126107b9366004613546565b611661565b6007546107de9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016104ee565b6104e4611a33565b6105847f000000000000000000000000000000000000000000000000000000000000000081565b6040546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6077546104e4565b61051c611aa4565b60408051808201909152600f81527f76322e302e302d70726576696f75730000000000000000000000000000000000602082015261051c565b607c5461089c9060ff1681565b60405190151581526020016104ee565b6104e4611ab1565b6104e46108c2366004613546565b611ac2565b6074546104e4565b6104e46108dd366004613593565b60066020525f908152604090205481565b6104e4607b5481565b6104e460055481565b61068a7e0100000000000000000000000000000000000000000000000000000000000081565b6105126109343660046135db565b611d79565b6104e460755481565b610512610950366004613614565b611e42565b603f546105849073ffffffffffffffffffffffffffffffffffffffff1681565b610512611e74565b61051261098b3660046135db565b611f5c565b6104e460765481565b61051261202e565b6104e460785481565b6105126109b8366004613546565b612100565b6109d06109cb3660046133a9565b6122bb565b60408051825181526020808401516fffffffffffffffffffffffffffffffff9081169183019190915292820151909216908201526060016104ee565b6105847f000000000000000000000000000000000000000000000000000000000000000081565b6105847f000000000000000000000000000000000000000000000000000000000000000081565b610512610a683660046135db565b61234d565b607d546105849073ffffffffffffffffffffffffffffffffffffffff1681565b610512610a9b3660046135db565b612416565b610512610aae3660046133a9565b61254c565b610512610ac13660046135db565b612607565b603e546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6104e460795481565b610512610afd3660046133a9565b612721565b6007546107de90700100000000000000000000000000000000900467ffffffffffffffff1681565b610512610b383660046136c8565b6127dc565b610796600181565b6002546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6105847f000000000000000000000000000000000000000000000000000000000000000081565b6104e4610b9a3660046133a9565b61286e565b610512612897565b6104e461298e565b6104e460775481565b603d546105849073ffffffffffffffffffffffffffffffffffffffff1681565b6105847f000000000000000000000000000000000000000000000000000000000000000081565b6105126129a4565b6007546107de9067ffffffffffffffff1681565b61068a610c2936600461330a565b90565b6104e4610c3a36600461330a565b60416020525f908152604090205481565b5f546105849062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b610512610c7e3660046135db565b612ab3565b603e5461089c9074010000000000000000000000000000000000000000900460ff1681565b603e545f9074010000000000000000000000000000000000000000900460ff1615158103610d3e57507fffffffff0000000000000000000000000000000000000000000000000000000081165f9081526041602052604090205480610d39576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610de8573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e0c91906136fa565b92915050565b607c54610100900473ffffffffffffffffffffffffffffffffffffffff163314610e68576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c5460ff16610ea4576040517f873dabd200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac905f90a1565b60048054610f0390613711565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2f90613711565b8015610f7a5780601f10610f5157610100808354040283529160200191610f7a565b820191905f5260205f20905b815481529060010190602001808311610f5d57829003601f168201915b505050505081565b607d5473ffffffffffffffffffffffffffffffffffffffff163314610fd3576040517f93f1e09400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c8054607d805473ffffffffffffffffffffffffffffffffffffffff8082166101009081027fffffffffffffffffffffff0000000000000000000000000000000000000000ff861617958690557fffffffffffffffffffffffff000000000000000000000000000000000000000090921690925560408051938290048316808552919094049091166020830152917f9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f331991015b60405180910390a150565b60405473ffffffffffffffffffffffffffffffffffffffff1633146110e2576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101611086565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146111b9576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806111f0576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205415611258576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260416020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461131a576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611351576040517f4bf41e1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607b546040518291907f5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b905f90a3607b55565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146113d5576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205461143c576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152604160209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461150b576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f03611544576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60775460408051918252602082018390527fc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2910160405180910390a1607755565b603e5473ffffffffffffffffffffffffffffffffffffffff1633146115d6576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d8054603e805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe29101611086565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146116b2576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561170257505f5460ff8083169116105b611793576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c1690036118d9575f5f5f5f5f5f5f5f5f5f8b8060200190518101906117ee9190613853565b9950995099509950995099509950995099509950600160f01b7dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916611857887fffff00000000000000000000000000000000000000000000000000000000000060109190911b1690565b7fffff00000000000000000000000000000000000000000000000000000000000016146118b0576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118b98a612bdc565b6118ca85858585858e8e8e8e612e61565b505050505050505050506119d9565b60ff5f5c166001036119a7575f5f5f5f5f868060200190518101906118fe919061393d565b9398509196509450925090507fffff000000000000000000000000000000000000000000000000000000000000601083901b167e0100000000000000000000000000000000000000000000000000000000000014611988576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61199185612bdc565b61199d84848484613056565b50505050506119d9565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016112bd565b6074545f9015611a9b5760748054611a4d906001906139d2565b81548110611a5d57611a5d6139e5565b5f91825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6075545b905090565b60038054610f0390613711565b6074545f90611a9f906001906139d2565b5f8151606014611afe576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f84806020019051810190611b159190613a12565b919450925090507fffff000000000000000000000000000000000000000000000000000000000000601084901b167e0100000000000000000000000000000000000000000000000000000000000014611b9a576040517f4570795000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ba261298e565b811015611bdb576040517f541d595b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b42611be58261286e565b10611c1c576040517f0dffe81800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81611c53576040517f80bcf51500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6074611c5e611ab1565b81548110611c6e57611c6e6139e5565b5f91825260209182902060029182020154607b54607c549254607a5460795460408051978801959095529386018990526060808701899052608087019390935260ff909416151560f81b60a08601527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000911b1660a184015260b583019190915260d582015260f5016040516020818303038152906040528051906020012090506001611d1985610ca8565b60405160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016602083015260248201526044810182905260640160405160208183030381529060405280519060200120945050505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611dcf576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001611086565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c54610100900473ffffffffffffffffffffffffffffffffffffffff163314611eca576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c5460ff1615611f07576040517f98b3177900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607c80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a905f90a1565b603d5473ffffffffffffffffffffffffffffffffffffffff163314611fad576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255603d546040805191909316815260208101919091527fc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b29101611086565b60015473ffffffffffffffffffffffffffffffffffffffff16331461207f576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16331461216f576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516060146121aa576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f828060200190518101906121c09190613a12565b9250925050806121cf60745490565b837fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e24260405161220191815260200190565b60405180910390a4604080516060810182529283526fffffffffffffffffffffffffffffffff42811660208501908152928116918401918252607480546001810182555f91909152935160029094027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810194909455915190518216700100000000000000000000000000000000029116177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef8149091015550565b604080516060810182525f8082526020820181905291810191909152607482815481106122ea576122ea6139e5565b5f91825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146123a3576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001611086565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314612485576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166124d2576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101611086565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461259d576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806125d4576040517f87b761a200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607a546040518291907fbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156905f90a3607a55565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612658576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166126a5576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101611086565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612772576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806127a9576040517f580a5fa800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6079546040518291907f390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0905f90a3607955565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612832576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600361283e8282613a92565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051611086919061332c565b5f6078546075548361288091906139d2565b61288a9190613ba9565b607654610e0c9190613bc0565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146128e8576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1661293c576040517f62de044500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e905f90a1565b5f60775461299a611a33565b611a9f9190613bc0565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146129f5576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1615612a4a576040517f93be805100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517fb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84905f90a1565b607c54610100900473ffffffffffffffffffffffffffffffffffffffff163314612b09576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116612b56576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255607c5460408051610100909204909316815260208101919091527ff67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af799101611086565b60c081015173ffffffffffffffffffffffffffffffffffffffff16612c2d576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060a001515f03612c6a576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03612ca4576040517fff5f860000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4281608001511115612ce2576040517f2403afcb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020810151612d1d576040517f4bf41e1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a081015160775580516078556074545f03612dff5760408051606080820183528383015182526080840180516fffffffffffffffffffffffffffffffff9081166020850190815292860180518216958501958652607480546001810182555f91909152945160029095027f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef813810195909555925194518116700100000000000000000000000000000000029416939093177f19a0b39aa25ac793b5f6e9a0534364cc0b3fd1ea9b651e79c7f50a59d48ef8149092019190915551607555516076555b6020810151607b5560c0810151607c80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff909316830217905560e08201516079550151607a55565b5f54610100900460ff16612ef7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161178a565b73ffffffffffffffffffffffffffffffffffffffff89161580612f2e575073ffffffffffffffffffffffffffffffffffffffff8816155b80612f4d575073ffffffffffffffffffffffffffffffffffffffff8116155b15612f84576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f9189898989896131ac565b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000095151595909502949094179093557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555050505050565b5f54610100900460ff166130ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161178a565b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000095151595909502949094179093557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055565b73ffffffffffffffffffffffffffffffffffffffff85166131f9576040517fe6cd565400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001691861691909117905560036132808382613a92565b50600461328d8282613a92565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081168114613307575f5ffd5b50565b5f6020828403121561331a575f5ffd5b8135613325816132da565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215613390575f5ffd5b823561339b816132da565b946020939093013593505050565b5f602082840312156133b9575f5ffd5b5035919050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610d39575f5ffd5b5f5f60408385031215613400575f5ffd5b613409836133c0565b9150613417602084016133c0565b90509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610120810167ffffffffffffffff8111828210171561347157613471613420565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156134be576134be613420565b604052919050565b5f67ffffffffffffffff8211156134df576134df613420565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f61351d613518846134c6565b613477565b9050828152838383011115613530575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215613556575f5ffd5b813567ffffffffffffffff81111561356c575f5ffd5b8201601f8101841361357c575f5ffd5b61358b8482356020840161350b565b949350505050565b5f602082840312156135a3575f5ffd5b813567ffffffffffffffff81168114613325575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114613307575f5ffd5b5f602082840312156135eb575f5ffd5b8135613325816135ba565b5f82601f830112613605575f5ffd5b6133258383356020850161350b565b5f5f5f5f5f5f60c08789031215613629575f5ffd5b8635613634816135ba565b95506020870135613644816135ba565b9450604087013563ffffffff8116811461365c575f5ffd5b9350606087013561366c816135ba565b9250608087013567ffffffffffffffff811115613687575f5ffd5b61369389828a016135f6565b92505060a087013567ffffffffffffffff8111156136af575f5ffd5b6136bb89828a016135f6565b9150509295509295509295565b5f602082840312156136d8575f5ffd5b813567ffffffffffffffff8111156136ee575f5ffd5b61358b848285016135f6565b5f6020828403121561370a575f5ffd5b5051919050565b600181811c9082168061372557607f821691505b60208210810361375c577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8051610d39816135ba565b5f610120828403121561377e575f5ffd5b61378661344d565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015290506137ca60c08301613762565b60c082015260e082810151908201526101009182015191810191909152919050565b80518015158114610d39575f5ffd5b8051610d39816132da565b5f82601f830112613815575f5ffd5b8151613823613518826134c6565b818152846020838601011115613837575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f5f6102408b8d03121561386d575f5ffd5b6138778c8c61376d565b99506138866101208c016137ec565b6101408c0151909950975061389e6101608c016137fb565b96506138ad6101808c01613762565b95506138bc6101a08c01613762565b94506138cb6101c08c01613762565b93506138da6101e08c01613762565b92506102008b015167ffffffffffffffff8111156138f6575f5ffd5b6139028d828e01613806565b9250506102208b015167ffffffffffffffff81111561391f575f5ffd5b61392b8d828e01613806565b9150509295989b9194979a5092959850565b5f5f5f5f5f6101a08688031215613952575f5ffd5b61395c878761376d565b945061396b61012087016137ec565b6101408701516101608801519195509350613985816132da565b610180870151909250613997816135ba565b809150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610e0c57610e0c6139a5565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f60608486031215613a24575f5ffd5b8351613a2f816132da565b602085015160409095015190969495509392505050565b601f821115613a8d57805f5260205f20601f840160051c81016020851015613a6b5750805b601f840160051c820191505b81811015613a8a575f8155600101613a77565b50505b505050565b815167ffffffffffffffff811115613aac57613aac613420565b613ac081613aba8454613711565b84613a46565b6020601f821160018114613b11575f8315613adb5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455613a8a565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015613b5e5787850151825560209485019460019092019101613b3e565b5084821015613b9a57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b8082028115828204841417610e0c57610e0c6139a5565b80820180821115610e0c57610e0c6139a556fea2646970667358221220f4faefe131fe869e281d480d49b7af6c5bd38db19e72cbff78d279ed7225e14264736f6c634300081c0033",
}

// AggchainfeppreviousABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainfeppreviousMetaData.ABI instead.
var AggchainfeppreviousABI = AggchainfeppreviousMetaData.ABI

// AggchainfeppreviousBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggchainfeppreviousMetaData.Bin instead.
var AggchainfeppreviousBin = AggchainfeppreviousMetaData.Bin

// DeployAggchainfepprevious deploys a new Ethereum contract, binding an instance of Aggchainfepprevious to it.
func DeployAggchainfepprevious(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Aggchainfepprevious, error) {
	parsed, err := AggchainfeppreviousMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggchainfeppreviousBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggchainfepprevious{AggchainfeppreviousCaller: AggchainfeppreviousCaller{contract: contract}, AggchainfeppreviousTransactor: AggchainfeppreviousTransactor{contract: contract}, AggchainfeppreviousFilterer: AggchainfeppreviousFilterer{contract: contract}}, nil
}

// Aggchainfepprevious is an auto generated Go binding around an Ethereum contract.
type Aggchainfepprevious struct {
	AggchainfeppreviousCaller     // Read-only binding to the contract
	AggchainfeppreviousTransactor // Write-only binding to the contract
	AggchainfeppreviousFilterer   // Log filterer for contract events
}

// AggchainfeppreviousCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainfeppreviousCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfeppreviousTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainfeppreviousTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfeppreviousFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainfeppreviousFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainfeppreviousSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainfeppreviousSession struct {
	Contract     *Aggchainfepprevious // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AggchainfeppreviousCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainfeppreviousCallerSession struct {
	Contract *AggchainfeppreviousCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AggchainfeppreviousTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainfeppreviousTransactorSession struct {
	Contract     *AggchainfeppreviousTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AggchainfeppreviousRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainfeppreviousRaw struct {
	Contract *Aggchainfepprevious // Generic contract binding to access the raw methods on
}

// AggchainfeppreviousCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainfeppreviousCallerRaw struct {
	Contract *AggchainfeppreviousCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainfeppreviousTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainfeppreviousTransactorRaw struct {
	Contract *AggchainfeppreviousTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainfepprevious creates a new instance of Aggchainfepprevious, bound to a specific deployed contract.
func NewAggchainfepprevious(address common.Address, backend bind.ContractBackend) (*Aggchainfepprevious, error) {
	contract, err := bindAggchainfepprevious(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggchainfepprevious{AggchainfeppreviousCaller: AggchainfeppreviousCaller{contract: contract}, AggchainfeppreviousTransactor: AggchainfeppreviousTransactor{contract: contract}, AggchainfeppreviousFilterer: AggchainfeppreviousFilterer{contract: contract}}, nil
}

// NewAggchainfeppreviousCaller creates a new read-only instance of Aggchainfepprevious, bound to a specific deployed contract.
func NewAggchainfeppreviousCaller(address common.Address, caller bind.ContractCaller) (*AggchainfeppreviousCaller, error) {
	contract, err := bindAggchainfepprevious(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousCaller{contract: contract}, nil
}

// NewAggchainfeppreviousTransactor creates a new write-only instance of Aggchainfepprevious, bound to a specific deployed contract.
func NewAggchainfeppreviousTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainfeppreviousTransactor, error) {
	contract, err := bindAggchainfepprevious(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousTransactor{contract: contract}, nil
}

// NewAggchainfeppreviousFilterer creates a new log filterer instance of Aggchainfepprevious, bound to a specific deployed contract.
func NewAggchainfeppreviousFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainfeppreviousFilterer, error) {
	contract, err := bindAggchainfepprevious(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousFilterer{contract: contract}, nil
}

// bindAggchainfepprevious binds a generic wrapper to an already deployed contract.
func bindAggchainfepprevious(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainfeppreviousMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainfepprevious *AggchainfeppreviousRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainfepprevious.Contract.AggchainfeppreviousCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainfepprevious *AggchainfeppreviousRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AggchainfeppreviousTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainfepprevious *AggchainfeppreviousRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AggchainfeppreviousTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainfepprevious *AggchainfeppreviousCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainfepprevious.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainfepprevious *AggchainfeppreviousTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainfepprevious *AggchainfeppreviousTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) AGGCHAINFEPVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "AGGCHAIN_FEP_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousSession) AGGCHAINFEPVERSION() (string, error) {
	return _Aggchainfepprevious.Contract.AGGCHAINFEPVERSION(&_Aggchainfepprevious.CallOpts)
}

// AGGCHAINFEPVERSION is a free data retrieval call binding the contract method 0x37d4d030.
//
// Solidity: function AGGCHAIN_FEP_VERSION() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) AGGCHAINFEPVERSION() (string, error) {
	return _Aggchainfepprevious.Contract.AGGCHAINFEPVERSION(&_Aggchainfepprevious.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainfepprevious.Contract.AGGCHAINTYPE(&_Aggchainfepprevious.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainfepprevious.Contract.AGGCHAINTYPE(&_Aggchainfepprevious.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainfepprevious.Contract.CONSENSUSTYPE(&_Aggchainfepprevious.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainfepprevious.Contract.CONSENSUSTYPE(&_Aggchainfepprevious.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) L2BLOCKTIME() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.L2BLOCKTIME(&_Aggchainfepprevious.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.L2BLOCKTIME(&_Aggchainfepprevious.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.SUBMISSIONINTERVAL(&_Aggchainfepprevious.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.SUBMISSIONINTERVAL(&_Aggchainfepprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) Admin() (common.Address, error) {
	return _Aggchainfepprevious.Contract.Admin(&_Aggchainfepprevious.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) Admin() (common.Address, error) {
	return _Aggchainfepprevious.Contract.Admin(&_Aggchainfepprevious.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainfepprevious.Contract.AggLayerGateway(&_Aggchainfepprevious.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainfepprevious.Contract.AggLayerGateway(&_Aggchainfepprevious.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) AggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "aggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) AggchainManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.AggchainManager(&_Aggchainfepprevious.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) AggchainManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.AggchainManager(&_Aggchainfepprevious.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) AggregationVkey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "aggregationVkey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) AggregationVkey() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.AggregationVkey(&_Aggchainfepprevious.CallOpts)
}

// AggregationVkey is a free data retrieval call binding the contract method 0xc32e4e3e.
//
// Solidity: function aggregationVkey() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) AggregationVkey() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.AggregationVkey(&_Aggchainfepprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) BridgeAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.BridgeAddress(&_Aggchainfepprevious.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) BridgeAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.BridgeAddress(&_Aggchainfepprevious.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _Aggchainfepprevious.Contract.ComputeL2Timestamp(&_Aggchainfepprevious.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _Aggchainfepprevious.Contract.ComputeL2Timestamp(&_Aggchainfepprevious.CallOpts, _l2BlockNumber)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.ForceBatchAddress(&_Aggchainfepprevious.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.ForceBatchAddress(&_Aggchainfepprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainfepprevious.Contract.ForceBatchTimeout(&_Aggchainfepprevious.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainfepprevious.Contract.ForceBatchTimeout(&_Aggchainfepprevious.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.ForcedBatches(&_Aggchainfepprevious.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.ForcedBatches(&_Aggchainfepprevious.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.GasTokenAddress(&_Aggchainfepprevious.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainfepprevious.Contract.GasTokenAddress(&_Aggchainfepprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainfepprevious.Contract.GasTokenNetwork(&_Aggchainfepprevious.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainfepprevious.Contract.GasTokenNetwork(&_Aggchainfepprevious.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainHash(&_Aggchainfepprevious.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainHash(&_Aggchainfepprevious.CallOpts, aggchainData)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetAggchainTypeFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getAggchainTypeFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainTypeFromSelector(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainTypeFromSelector(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getAggchainVKey", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKey(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKey(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetAggchainVKeySelector(opts *bind.CallOpts, aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getAggchainVKeySelector", aggchainVKeyVersion, aggchainType)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKeySelector(&_Aggchainfepprevious.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKeySelector(&_Aggchainfepprevious.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetAggchainVKeyVersionFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getAggchainVKeyVersionFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainfepprevious.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (AggchainFEPPreviousOutputProposal, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(AggchainFEPPreviousOutputProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(AggchainFEPPreviousOutputProposal)).(*AggchainFEPPreviousOutputProposal)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfepprevious *AggchainfeppreviousSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPPreviousOutputProposal, error) {
	return _Aggchainfepprevious.Contract.GetL2Output(&_Aggchainfepprevious.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((bytes32,uint128,uint128))
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GetL2Output(_l2OutputIndex *big.Int) (AggchainFEPPreviousOutputProposal, error) {
	return _Aggchainfepprevious.Contract.GetL2Output(&_Aggchainfepprevious.CallOpts, _l2OutputIndex)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.GlobalExitRootManager(&_Aggchainfepprevious.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.GlobalExitRootManager(&_Aggchainfepprevious.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfepprevious *AggchainfeppreviousCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainfepprevious.Contract.Initialize0(&_Aggchainfepprevious.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainfepprevious.Contract.Initialize0(&_Aggchainfepprevious.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) L2BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "l2BlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) L2BlockTime() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.L2BlockTime(&_Aggchainfepprevious.CallOpts)
}

// L2BlockTime is a free data retrieval call binding the contract method 0x93991af3.
//
// Solidity: function l2BlockTime() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) L2BlockTime() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.L2BlockTime(&_Aggchainfepprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.LastAccInputHash(&_Aggchainfepprevious.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.LastAccInputHash(&_Aggchainfepprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousSession) LastForceBatch() (uint64, error) {
	return _Aggchainfepprevious.Contract.LastForceBatch(&_Aggchainfepprevious.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) LastForceBatch() (uint64, error) {
	return _Aggchainfepprevious.Contract.LastForceBatch(&_Aggchainfepprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainfepprevious.Contract.LastForceBatchSequenced(&_Aggchainfepprevious.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainfepprevious.Contract.LastForceBatchSequenced(&_Aggchainfepprevious.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) LatestBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.LatestBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.LatestBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) LatestOutputIndex() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.LatestOutputIndex(&_Aggchainfepprevious.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.LatestOutputIndex(&_Aggchainfepprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousSession) NetworkName() (string, error) {
	return _Aggchainfepprevious.Contract.NetworkName(&_Aggchainfepprevious.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) NetworkName() (string, error) {
	return _Aggchainfepprevious.Contract.NetworkName(&_Aggchainfepprevious.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) NextBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.NextBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) NextBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.NextBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) NextOutputIndex() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.NextOutputIndex(&_Aggchainfepprevious.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) NextOutputIndex() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.NextOutputIndex(&_Aggchainfepprevious.CallOpts)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) OptimisticMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "optimisticMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousSession) OptimisticMode() (bool, error) {
	return _Aggchainfepprevious.Contract.OptimisticMode(&_Aggchainfepprevious.CallOpts)
}

// OptimisticMode is a free data retrieval call binding the contract method 0x60caf7a0.
//
// Solidity: function optimisticMode() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) OptimisticMode() (bool, error) {
	return _Aggchainfepprevious.Contract.OptimisticMode(&_Aggchainfepprevious.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) OptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "optimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) OptimisticModeManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.OptimisticModeManager(&_Aggchainfepprevious.CallOpts)
}

// OptimisticModeManager is a free data retrieval call binding the contract method 0x1cf810ee.
//
// Solidity: function optimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) OptimisticModeManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.OptimisticModeManager(&_Aggchainfepprevious.CallOpts)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) OwnedAggchainVKeys(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "ownedAggchainVKeys", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.OwnedAggchainVKeys(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainfepprevious.Contract.OwnedAggchainVKeys(&_Aggchainfepprevious.CallOpts, aggchainVKeySelector)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) PendingAdmin() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingAdmin(&_Aggchainfepprevious.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) PendingAdmin() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingAdmin(&_Aggchainfepprevious.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) PendingAggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "pendingAggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingAggchainManager(&_Aggchainfepprevious.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingAggchainManager(&_Aggchainfepprevious.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) PendingOptimisticModeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "pendingOptimisticModeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) PendingOptimisticModeManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingOptimisticModeManager(&_Aggchainfepprevious.CallOpts)
}

// PendingOptimisticModeManager is a free data retrieval call binding the contract method 0xadb8696c.
//
// Solidity: function pendingOptimisticModeManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) PendingOptimisticModeManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingOptimisticModeManager(&_Aggchainfepprevious.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) PendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "pendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingVKeyManager(&_Aggchainfepprevious.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.PendingVKeyManager(&_Aggchainfepprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) Pol() (common.Address, error) {
	return _Aggchainfepprevious.Contract.Pol(&_Aggchainfepprevious.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) Pol() (common.Address, error) {
	return _Aggchainfepprevious.Contract.Pol(&_Aggchainfepprevious.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) RangeVkeyCommitment(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "rangeVkeyCommitment")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) RangeVkeyCommitment() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.RangeVkeyCommitment(&_Aggchainfepprevious.CallOpts)
}

// RangeVkeyCommitment is a free data retrieval call binding the contract method 0x2b31841e.
//
// Solidity: function rangeVkeyCommitment() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) RangeVkeyCommitment() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.RangeVkeyCommitment(&_Aggchainfepprevious.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) RollupConfigHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "rollupConfigHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousSession) RollupConfigHash() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.RollupConfigHash(&_Aggchainfepprevious.CallOpts)
}

// RollupConfigHash is a free data retrieval call binding the contract method 0x6d9a1c8b.
//
// Solidity: function rollupConfigHash() view returns(bytes32)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) RollupConfigHash() ([32]byte, error) {
	return _Aggchainfepprevious.Contract.RollupConfigHash(&_Aggchainfepprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) RollupManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.RollupManager(&_Aggchainfepprevious.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) RollupManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.RollupManager(&_Aggchainfepprevious.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) StartingBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.StartingBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.StartingBlockNumber(&_Aggchainfepprevious.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) StartingTimestamp() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.StartingTimestamp(&_Aggchainfepprevious.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) StartingTimestamp() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.StartingTimestamp(&_Aggchainfepprevious.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) SubmissionInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "submissionInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousSession) SubmissionInterval() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.SubmissionInterval(&_Aggchainfepprevious.CallOpts)
}

// SubmissionInterval is a free data retrieval call binding the contract method 0xe1a41bcf.
//
// Solidity: function submissionInterval() view returns(uint256)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) SubmissionInterval() (*big.Int, error) {
	return _Aggchainfepprevious.Contract.SubmissionInterval(&_Aggchainfepprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainfepprevious.Contract.TrustedSequencer(&_Aggchainfepprevious.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainfepprevious.Contract.TrustedSequencer(&_Aggchainfepprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousSession) TrustedSequencerURL() (string, error) {
	return _Aggchainfepprevious.Contract.TrustedSequencerURL(&_Aggchainfepprevious.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) TrustedSequencerURL() (string, error) {
	return _Aggchainfepprevious.Contract.TrustedSequencerURL(&_Aggchainfepprevious.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousSession) UseDefaultGateway() (bool, error) {
	return _Aggchainfepprevious.Contract.UseDefaultGateway(&_Aggchainfepprevious.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) UseDefaultGateway() (bool, error) {
	return _Aggchainfepprevious.Contract.UseDefaultGateway(&_Aggchainfepprevious.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) VKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "vKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousSession) VKeyManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.VKeyManager(&_Aggchainfepprevious.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) VKeyManager() (common.Address, error) {
	return _Aggchainfepprevious.Contract.VKeyManager(&_Aggchainfepprevious.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainfepprevious.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousSession) Version() (string, error) {
	return _Aggchainfepprevious.Contract.Version(&_Aggchainfepprevious.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainfepprevious *AggchainfeppreviousCallerSession) Version() (string, error) {
	return _Aggchainfepprevious.Contract.Version(&_Aggchainfepprevious.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptAdminRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptAdminRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) AcceptAggchainManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "acceptAggchainManagerRole")
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptAggchainManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptAggchainManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) AcceptOptimisticModeManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "acceptOptimisticModeManagerRole")
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptOptimisticModeManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0x12634900.
//
// Solidity: function acceptOptimisticModeManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) AcceptOptimisticModeManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptOptimisticModeManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) AcceptVKeyManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "acceptVKeyManagerRole")
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptVKeyManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AcceptVKeyManagerRole(&_Aggchainfepprevious.TransactOpts)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "addOwnedAggchainVKey", aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AddOwnedAggchainVKey(&_Aggchainfepprevious.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.AddOwnedAggchainVKey(&_Aggchainfepprevious.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) DisableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "disableOptimisticMode")
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.DisableOptimisticMode(&_Aggchainfepprevious.TransactOpts)
}

// DisableOptimisticMode is a paid mutator transaction binding the contract method 0x0822dc61.
//
// Solidity: function disableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) DisableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.DisableOptimisticMode(&_Aggchainfepprevious.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.DisableUseDefaultGatewayFlag(&_Aggchainfepprevious.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.DisableUseDefaultGatewayFlag(&_Aggchainfepprevious.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) EnableOptimisticMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "enableOptimisticMode")
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.EnableOptimisticMode(&_Aggchainfepprevious.TransactOpts)
}

// EnableOptimisticMode is a paid mutator transaction binding the contract method 0x81eb0baf.
//
// Solidity: function enableOptimisticMode() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) EnableOptimisticMode() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.EnableOptimisticMode(&_Aggchainfepprevious.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.EnableUseDefaultGatewayFlag(&_Aggchainfepprevious.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.EnableUseDefaultGatewayFlag(&_Aggchainfepprevious.TransactOpts)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) InitAggchainManager(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "initAggchainManager", newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.InitAggchainManager(&_Aggchainfepprevious.TransactOpts, newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.InitAggchainManager(&_Aggchainfepprevious.TransactOpts, newAggchainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) Initialize(opts *bind.TransactOpts, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "initialize", initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.Initialize(&_Aggchainfepprevious.TransactOpts, initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.Initialize(&_Aggchainfepprevious.TransactOpts, initializeBytesAggchain)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.OnVerifyPessimistic(&_Aggchainfepprevious.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.OnVerifyPessimistic(&_Aggchainfepprevious.TransactOpts, aggchainData)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.SetTrustedSequencer(&_Aggchainfepprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.SetTrustedSequencer(&_Aggchainfepprevious.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.SetTrustedSequencerURL(&_Aggchainfepprevious.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.SetTrustedSequencerURL(&_Aggchainfepprevious.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferAdminRole(&_Aggchainfepprevious.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferAdminRole(&_Aggchainfepprevious.TransactOpts, newPendingAdmin)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) TransferAggchainManagerRole(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "transferAggchainManagerRole", newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferAggchainManagerRole(&_Aggchainfepprevious.TransactOpts, newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferAggchainManagerRole(&_Aggchainfepprevious.TransactOpts, newAggchainManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) TransferOptimisticModeManagerRole(opts *bind.TransactOpts, newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "transferOptimisticModeManagerRole", newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferOptimisticModeManagerRole(&_Aggchainfepprevious.TransactOpts, newOptimisticModeManager)
}

// TransferOptimisticModeManagerRole is a paid mutator transaction binding the contract method 0xfdbbc19b.
//
// Solidity: function transferOptimisticModeManagerRole(address newOptimisticModeManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) TransferOptimisticModeManagerRole(newOptimisticModeManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferOptimisticModeManagerRole(&_Aggchainfepprevious.TransactOpts, newOptimisticModeManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) TransferVKeyManagerRole(opts *bind.TransactOpts, newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "transferVKeyManagerRole", newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferVKeyManagerRole(&_Aggchainfepprevious.TransactOpts, newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.TransferVKeyManagerRole(&_Aggchainfepprevious.TransactOpts, newVKeyManager)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) UpdateAggregationVkey(opts *bind.TransactOpts, _aggregationVkey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "updateAggregationVkey", _aggregationVkey)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) UpdateAggregationVkey(_aggregationVkey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateAggregationVkey(&_Aggchainfepprevious.TransactOpts, _aggregationVkey)
}

// UpdateAggregationVkey is a paid mutator transaction binding the contract method 0xc4cb03ec.
//
// Solidity: function updateAggregationVkey(bytes32 _aggregationVkey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) UpdateAggregationVkey(_aggregationVkey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateAggregationVkey(&_Aggchainfepprevious.TransactOpts, _aggregationVkey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateOwnedAggchainVKey(&_Aggchainfepprevious.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateOwnedAggchainVKey(&_Aggchainfepprevious.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) UpdateRangeVkeyCommitment(opts *bind.TransactOpts, _rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "updateRangeVkeyCommitment", _rangeVkeyCommitment)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) UpdateRangeVkeyCommitment(_rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateRangeVkeyCommitment(&_Aggchainfepprevious.TransactOpts, _rangeVkeyCommitment)
}

// UpdateRangeVkeyCommitment is a paid mutator transaction binding the contract method 0xbc91ce33.
//
// Solidity: function updateRangeVkeyCommitment(bytes32 _rangeVkeyCommitment) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) UpdateRangeVkeyCommitment(_rangeVkeyCommitment [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateRangeVkeyCommitment(&_Aggchainfepprevious.TransactOpts, _rangeVkeyCommitment)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) UpdateRollupConfigHash(opts *bind.TransactOpts, _rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "updateRollupConfigHash", _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateRollupConfigHash(&_Aggchainfepprevious.TransactOpts, _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateRollupConfigHash(&_Aggchainfepprevious.TransactOpts, _rollupConfigHash)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactor) UpdateSubmissionInterval(opts *bind.TransactOpts, _submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfepprevious.contract.Transact(opts, "updateSubmissionInterval", _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfepprevious *AggchainfeppreviousSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateSubmissionInterval(&_Aggchainfepprevious.TransactOpts, _submissionInterval)
}

// UpdateSubmissionInterval is a paid mutator transaction binding the contract method 0x336c9e81.
//
// Solidity: function updateSubmissionInterval(uint256 _submissionInterval) returns()
func (_Aggchainfepprevious *AggchainfeppreviousTransactorSession) UpdateSubmissionInterval(_submissionInterval *big.Int) (*types.Transaction, error) {
	return _Aggchainfepprevious.Contract.UpdateSubmissionInterval(&_Aggchainfepprevious.TransactOpts, _submissionInterval)
}

// AggchainfeppreviousAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptAdminRoleIterator struct {
	Event *AggchainfeppreviousAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAcceptAdminRole)
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
		it.Event = new(AggchainfeppreviousAcceptAdminRole)
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
func (it *AggchainfeppreviousAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAcceptAdminRole represents a AcceptAdminRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainfeppreviousAcceptAdminRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAcceptAdminRoleIterator{contract: _Aggchainfepprevious.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAcceptAdminRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainfeppreviousAcceptAdminRole, error) {
	event := new(AggchainfeppreviousAcceptAdminRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptAggchainManagerRoleIterator struct {
	Event *AggchainfeppreviousAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAcceptAggchainManagerRole)
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
		it.Event = new(AggchainfeppreviousAcceptAggchainManagerRole)
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
func (it *AggchainfeppreviousAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAcceptAggchainManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAcceptAggchainManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*AggchainfeppreviousAcceptAggchainManagerRole, error) {
	event := new(AggchainfeppreviousAcceptAggchainManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator is returned from FilterAcceptOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for AcceptOptimisticModeManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator struct {
	Event *AggchainfeppreviousAcceptOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAcceptOptimisticModeManagerRole)
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
		it.Event = new(AggchainfeppreviousAcceptOptimisticModeManagerRole)
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
func (it *AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAcceptOptimisticModeManagerRole represents a AcceptOptimisticModeManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptOptimisticModeManagerRole struct {
	OldOptimisticModeManager common.Address
	NewOptimisticModeManager common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAcceptOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAcceptOptimisticModeManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "AcceptOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptOptimisticModeManagerRole is a free log subscription operation binding the contract event 0x9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f3319.
//
// Solidity: event AcceptOptimisticModeManagerRole(address oldOptimisticModeManager, address newOptimisticModeManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAcceptOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAcceptOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AcceptOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAcceptOptimisticModeManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAcceptOptimisticModeManagerRole(log types.Log) (*AggchainfeppreviousAcceptOptimisticModeManagerRole, error) {
	event := new(AggchainfeppreviousAcceptOptimisticModeManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptVKeyManagerRoleIterator struct {
	Event *AggchainfeppreviousAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAcceptVKeyManagerRole)
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
		it.Event = new(AggchainfeppreviousAcceptVKeyManagerRole)
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
func (it *AggchainfeppreviousAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAcceptVKeyManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAcceptVKeyManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*AggchainfeppreviousAcceptVKeyManagerRole, error) {
	event := new(AggchainfeppreviousAcceptVKeyManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAddAggchainVKeyIterator struct {
	Event *AggchainfeppreviousAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAddAggchainVKey)
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
		it.Event = new(AggchainfeppreviousAddAggchainVKey)
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
func (it *AggchainfeppreviousAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAddAggchainVKey represents a AddAggchainVKey event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainfeppreviousAddAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAddAggchainVKeyIterator{contract: _Aggchainfepprevious.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAddAggchainVKey)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainfeppreviousAddAggchainVKey, error) {
	event := new(AggchainfeppreviousAddAggchainVKey)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousAggregationVkeyUpdatedIterator is returned from FilterAggregationVkeyUpdated and is used to iterate over the raw logs and unpacked data for AggregationVkeyUpdated events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAggregationVkeyUpdatedIterator struct {
	Event *AggchainfeppreviousAggregationVkeyUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousAggregationVkeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousAggregationVkeyUpdated)
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
		it.Event = new(AggchainfeppreviousAggregationVkeyUpdated)
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
func (it *AggchainfeppreviousAggregationVkeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousAggregationVkeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousAggregationVkeyUpdated represents a AggregationVkeyUpdated event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousAggregationVkeyUpdated struct {
	OldAggregationVkey [32]byte
	NewAggregationVkey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAggregationVkeyUpdated is a free log retrieval operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterAggregationVkeyUpdated(opts *bind.FilterOpts, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (*AggchainfeppreviousAggregationVkeyUpdatedIterator, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousAggregationVkeyUpdatedIterator{contract: _Aggchainfepprevious.contract, event: "AggregationVkeyUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregationVkeyUpdated is a free log subscription operation binding the contract event 0x390b73b2b067afcef04d30b573e4590c6e565519e370927dd777ca0ce8a55db0.
//
// Solidity: event AggregationVkeyUpdated(bytes32 indexed oldAggregationVkey, bytes32 indexed newAggregationVkey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchAggregationVkeyUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousAggregationVkeyUpdated, oldAggregationVkey [][32]byte, newAggregationVkey [][32]byte) (event.Subscription, error) {

	var oldAggregationVkeyRule []interface{}
	for _, oldAggregationVkeyItem := range oldAggregationVkey {
		oldAggregationVkeyRule = append(oldAggregationVkeyRule, oldAggregationVkeyItem)
	}
	var newAggregationVkeyRule []interface{}
	for _, newAggregationVkeyItem := range newAggregationVkey {
		newAggregationVkeyRule = append(newAggregationVkeyRule, newAggregationVkeyItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "AggregationVkeyUpdated", oldAggregationVkeyRule, newAggregationVkeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousAggregationVkeyUpdated)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseAggregationVkeyUpdated(log types.Log) (*AggchainfeppreviousAggregationVkeyUpdated, error) {
	event := new(AggchainfeppreviousAggregationVkeyUpdated)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "AggregationVkeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousDisableOptimisticModeIterator is returned from FilterDisableOptimisticMode and is used to iterate over the raw logs and unpacked data for DisableOptimisticMode events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousDisableOptimisticModeIterator struct {
	Event *AggchainfeppreviousDisableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousDisableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousDisableOptimisticMode)
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
		it.Event = new(AggchainfeppreviousDisableOptimisticMode)
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
func (it *AggchainfeppreviousDisableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousDisableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousDisableOptimisticMode represents a DisableOptimisticMode event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousDisableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableOptimisticMode is a free log retrieval operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterDisableOptimisticMode(opts *bind.FilterOpts) (*AggchainfeppreviousDisableOptimisticModeIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousDisableOptimisticModeIterator{contract: _Aggchainfepprevious.contract, event: "DisableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchDisableOptimisticMode is a free log subscription operation binding the contract event 0x334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac.
//
// Solidity: event DisableOptimisticMode()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchDisableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousDisableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "DisableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousDisableOptimisticMode)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseDisableOptimisticMode(log types.Log) (*AggchainfeppreviousDisableOptimisticMode, error) {
	event := new(AggchainfeppreviousDisableOptimisticMode)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "DisableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousDisableUseDefaultGatewayFlagIterator struct {
	Event *AggchainfeppreviousDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousDisableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainfeppreviousDisableUseDefaultGatewayFlag)
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
func (it *AggchainfeppreviousDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainfeppreviousDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousDisableUseDefaultGatewayFlagIterator{contract: _Aggchainfepprevious.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousDisableUseDefaultGatewayFlag)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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

// ParseDisableUseDefaultGatewayFlag is a log parse operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*AggchainfeppreviousDisableUseDefaultGatewayFlag, error) {
	event := new(AggchainfeppreviousDisableUseDefaultGatewayFlag)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousEnableOptimisticModeIterator is returned from FilterEnableOptimisticMode and is used to iterate over the raw logs and unpacked data for EnableOptimisticMode events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousEnableOptimisticModeIterator struct {
	Event *AggchainfeppreviousEnableOptimisticMode // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousEnableOptimisticModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousEnableOptimisticMode)
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
		it.Event = new(AggchainfeppreviousEnableOptimisticMode)
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
func (it *AggchainfeppreviousEnableOptimisticModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousEnableOptimisticModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousEnableOptimisticMode represents a EnableOptimisticMode event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousEnableOptimisticMode struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableOptimisticMode is a free log retrieval operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterEnableOptimisticMode(opts *bind.FilterOpts) (*AggchainfeppreviousEnableOptimisticModeIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousEnableOptimisticModeIterator{contract: _Aggchainfepprevious.contract, event: "EnableOptimisticMode", logs: logs, sub: sub}, nil
}

// WatchEnableOptimisticMode is a free log subscription operation binding the contract event 0x26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a.
//
// Solidity: event EnableOptimisticMode()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchEnableOptimisticMode(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousEnableOptimisticMode) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "EnableOptimisticMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousEnableOptimisticMode)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseEnableOptimisticMode(log types.Log) (*AggchainfeppreviousEnableOptimisticMode, error) {
	event := new(AggchainfeppreviousEnableOptimisticMode)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "EnableOptimisticMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousEnableUseDefaultGatewayFlagIterator struct {
	Event *AggchainfeppreviousEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousEnableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainfeppreviousEnableUseDefaultGatewayFlag)
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
func (it *AggchainfeppreviousEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainfeppreviousEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousEnableUseDefaultGatewayFlagIterator{contract: _Aggchainfepprevious.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousEnableUseDefaultGatewayFlag)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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

// ParseEnableUseDefaultGatewayFlag is a log parse operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*AggchainfeppreviousEnableUseDefaultGatewayFlag, error) {
	event := new(AggchainfeppreviousEnableUseDefaultGatewayFlag)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousInitializedIterator struct {
	Event *AggchainfeppreviousInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousInitialized)
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
		it.Event = new(AggchainfeppreviousInitialized)
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
func (it *AggchainfeppreviousInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousInitialized represents a Initialized event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainfeppreviousInitializedIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousInitializedIterator{contract: _Aggchainfepprevious.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousInitialized)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseInitialized(log types.Log) (*AggchainfeppreviousInitialized, error) {
	event := new(AggchainfeppreviousInitialized)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousOutputProposedIterator is returned from FilterOutputProposed and is used to iterate over the raw logs and unpacked data for OutputProposed events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousOutputProposedIterator struct {
	Event *AggchainfeppreviousOutputProposed // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousOutputProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousOutputProposed)
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
		it.Event = new(AggchainfeppreviousOutputProposed)
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
func (it *AggchainfeppreviousOutputProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousOutputProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousOutputProposed represents a OutputProposed event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousOutputProposed struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputProposed is a free log retrieval operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterOutputProposed(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*AggchainfeppreviousOutputProposedIterator, error) {

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

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousOutputProposedIterator{contract: _Aggchainfepprevious.contract, event: "OutputProposed", logs: logs, sub: sub}, nil
}

// WatchOutputProposed is a free log subscription operation binding the contract event 0xa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e2.
//
// Solidity: event OutputProposed(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchOutputProposed(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousOutputProposed, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "OutputProposed", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousOutputProposed)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "OutputProposed", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseOutputProposed(log types.Log) (*AggchainfeppreviousOutputProposed, error) {
	event := new(AggchainfeppreviousOutputProposed)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "OutputProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator is returned from FilterRangeVkeyCommitmentUpdated and is used to iterate over the raw logs and unpacked data for RangeVkeyCommitmentUpdated events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator struct {
	Event *AggchainfeppreviousRangeVkeyCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousRangeVkeyCommitmentUpdated)
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
		it.Event = new(AggchainfeppreviousRangeVkeyCommitmentUpdated)
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
func (it *AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousRangeVkeyCommitmentUpdated represents a RangeVkeyCommitmentUpdated event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousRangeVkeyCommitmentUpdated struct {
	OldRangeVkeyCommitment [32]byte
	NewRangeVkeyCommitment [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRangeVkeyCommitmentUpdated is a free log retrieval operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterRangeVkeyCommitmentUpdated(opts *bind.FilterOpts, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (*AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousRangeVkeyCommitmentUpdatedIterator{contract: _Aggchainfepprevious.contract, event: "RangeVkeyCommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchRangeVkeyCommitmentUpdated is a free log subscription operation binding the contract event 0xbf8cab6317796bfa97fea82b6d27c9542a08fa0821813cf2a57e7cff7fdc8156.
//
// Solidity: event RangeVkeyCommitmentUpdated(bytes32 indexed oldRangeVkeyCommitment, bytes32 indexed newRangeVkeyCommitment)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchRangeVkeyCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousRangeVkeyCommitmentUpdated, oldRangeVkeyCommitment [][32]byte, newRangeVkeyCommitment [][32]byte) (event.Subscription, error) {

	var oldRangeVkeyCommitmentRule []interface{}
	for _, oldRangeVkeyCommitmentItem := range oldRangeVkeyCommitment {
		oldRangeVkeyCommitmentRule = append(oldRangeVkeyCommitmentRule, oldRangeVkeyCommitmentItem)
	}
	var newRangeVkeyCommitmentRule []interface{}
	for _, newRangeVkeyCommitmentItem := range newRangeVkeyCommitment {
		newRangeVkeyCommitmentRule = append(newRangeVkeyCommitmentRule, newRangeVkeyCommitmentItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "RangeVkeyCommitmentUpdated", oldRangeVkeyCommitmentRule, newRangeVkeyCommitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousRangeVkeyCommitmentUpdated)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseRangeVkeyCommitmentUpdated(log types.Log) (*AggchainfeppreviousRangeVkeyCommitmentUpdated, error) {
	event := new(AggchainfeppreviousRangeVkeyCommitmentUpdated)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "RangeVkeyCommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousRollupConfigHashUpdatedIterator is returned from FilterRollupConfigHashUpdated and is used to iterate over the raw logs and unpacked data for RollupConfigHashUpdated events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousRollupConfigHashUpdatedIterator struct {
	Event *AggchainfeppreviousRollupConfigHashUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousRollupConfigHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousRollupConfigHashUpdated)
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
		it.Event = new(AggchainfeppreviousRollupConfigHashUpdated)
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
func (it *AggchainfeppreviousRollupConfigHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousRollupConfigHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousRollupConfigHashUpdated represents a RollupConfigHashUpdated event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousRollupConfigHashUpdated struct {
	OldRollupConfigHash [32]byte
	NewRollupConfigHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollupConfigHashUpdated is a free log retrieval operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterRollupConfigHashUpdated(opts *bind.FilterOpts, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (*AggchainfeppreviousRollupConfigHashUpdatedIterator, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousRollupConfigHashUpdatedIterator{contract: _Aggchainfepprevious.contract, event: "RollupConfigHashUpdated", logs: logs, sub: sub}, nil
}

// WatchRollupConfigHashUpdated is a free log subscription operation binding the contract event 0x5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b.
//
// Solidity: event RollupConfigHashUpdated(bytes32 indexed oldRollupConfigHash, bytes32 indexed newRollupConfigHash)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchRollupConfigHashUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousRollupConfigHashUpdated, oldRollupConfigHash [][32]byte, newRollupConfigHash [][32]byte) (event.Subscription, error) {

	var oldRollupConfigHashRule []interface{}
	for _, oldRollupConfigHashItem := range oldRollupConfigHash {
		oldRollupConfigHashRule = append(oldRollupConfigHashRule, oldRollupConfigHashItem)
	}
	var newRollupConfigHashRule []interface{}
	for _, newRollupConfigHashItem := range newRollupConfigHash {
		newRollupConfigHashRule = append(newRollupConfigHashRule, newRollupConfigHashItem)
	}

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "RollupConfigHashUpdated", oldRollupConfigHashRule, newRollupConfigHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousRollupConfigHashUpdated)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseRollupConfigHashUpdated(log types.Log) (*AggchainfeppreviousRollupConfigHashUpdated, error) {
	event := new(AggchainfeppreviousRollupConfigHashUpdated)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "RollupConfigHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSetTrustedSequencerIterator struct {
	Event *AggchainfeppreviousSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousSetTrustedSequencer)
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
		it.Event = new(AggchainfeppreviousSetTrustedSequencer)
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
func (it *AggchainfeppreviousSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousSetTrustedSequencer represents a SetTrustedSequencer event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainfeppreviousSetTrustedSequencerIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousSetTrustedSequencerIterator{contract: _Aggchainfepprevious.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousSetTrustedSequencer)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainfeppreviousSetTrustedSequencer, error) {
	event := new(AggchainfeppreviousSetTrustedSequencer)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSetTrustedSequencerURLIterator struct {
	Event *AggchainfeppreviousSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousSetTrustedSequencerURL)
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
		it.Event = new(AggchainfeppreviousSetTrustedSequencerURL)
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
func (it *AggchainfeppreviousSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainfeppreviousSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousSetTrustedSequencerURLIterator{contract: _Aggchainfepprevious.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousSetTrustedSequencerURL)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainfeppreviousSetTrustedSequencerURL, error) {
	event := new(AggchainfeppreviousSetTrustedSequencerURL)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousSubmissionIntervalUpdatedIterator is returned from FilterSubmissionIntervalUpdated and is used to iterate over the raw logs and unpacked data for SubmissionIntervalUpdated events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSubmissionIntervalUpdatedIterator struct {
	Event *AggchainfeppreviousSubmissionIntervalUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousSubmissionIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousSubmissionIntervalUpdated)
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
		it.Event = new(AggchainfeppreviousSubmissionIntervalUpdated)
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
func (it *AggchainfeppreviousSubmissionIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousSubmissionIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousSubmissionIntervalUpdated represents a SubmissionIntervalUpdated event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousSubmissionIntervalUpdated struct {
	OldSubmissionInterval *big.Int
	NewSubmissionInterval *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubmissionIntervalUpdated is a free log retrieval operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterSubmissionIntervalUpdated(opts *bind.FilterOpts) (*AggchainfeppreviousSubmissionIntervalUpdatedIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousSubmissionIntervalUpdatedIterator{contract: _Aggchainfepprevious.contract, event: "SubmissionIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchSubmissionIntervalUpdated is a free log subscription operation binding the contract event 0xc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2.
//
// Solidity: event SubmissionIntervalUpdated(uint256 oldSubmissionInterval, uint256 newSubmissionInterval)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchSubmissionIntervalUpdated(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousSubmissionIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "SubmissionIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousSubmissionIntervalUpdated)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseSubmissionIntervalUpdated(log types.Log) (*AggchainfeppreviousSubmissionIntervalUpdated, error) {
	event := new(AggchainfeppreviousSubmissionIntervalUpdated)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "SubmissionIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferAdminRoleIterator struct {
	Event *AggchainfeppreviousTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousTransferAdminRole)
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
		it.Event = new(AggchainfeppreviousTransferAdminRole)
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
func (it *AggchainfeppreviousTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousTransferAdminRole represents a TransferAdminRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainfeppreviousTransferAdminRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousTransferAdminRoleIterator{contract: _Aggchainfepprevious.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousTransferAdminRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseTransferAdminRole(log types.Log) (*AggchainfeppreviousTransferAdminRole, error) {
	event := new(AggchainfeppreviousTransferAdminRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferAggchainManagerRoleIterator struct {
	Event *AggchainfeppreviousTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousTransferAggchainManagerRole)
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
		it.Event = new(AggchainfeppreviousTransferAggchainManagerRole)
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
func (it *AggchainfeppreviousTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousTransferAggchainManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousTransferAggchainManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseTransferAggchainManagerRole(log types.Log) (*AggchainfeppreviousTransferAggchainManagerRole, error) {
	event := new(AggchainfeppreviousTransferAggchainManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousTransferOptimisticModeManagerRoleIterator is returned from FilterTransferOptimisticModeManagerRole and is used to iterate over the raw logs and unpacked data for TransferOptimisticModeManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferOptimisticModeManagerRoleIterator struct {
	Event *AggchainfeppreviousTransferOptimisticModeManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousTransferOptimisticModeManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousTransferOptimisticModeManagerRole)
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
		it.Event = new(AggchainfeppreviousTransferOptimisticModeManagerRole)
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
func (it *AggchainfeppreviousTransferOptimisticModeManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousTransferOptimisticModeManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousTransferOptimisticModeManagerRole represents a TransferOptimisticModeManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferOptimisticModeManagerRole struct {
	CurrentOptimisticModeManager    common.Address
	NewPendingOptimisticModeManager common.Address
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTransferOptimisticModeManagerRole is a free log retrieval operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterTransferOptimisticModeManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousTransferOptimisticModeManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousTransferOptimisticModeManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "TransferOptimisticModeManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferOptimisticModeManagerRole is a free log subscription operation binding the contract event 0xf67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af79.
//
// Solidity: event TransferOptimisticModeManagerRole(address currentOptimisticModeManager, address newPendingOptimisticModeManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchTransferOptimisticModeManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousTransferOptimisticModeManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "TransferOptimisticModeManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousTransferOptimisticModeManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseTransferOptimisticModeManagerRole(log types.Log) (*AggchainfeppreviousTransferOptimisticModeManagerRole, error) {
	event := new(AggchainfeppreviousTransferOptimisticModeManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferOptimisticModeManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferVKeyManagerRoleIterator struct {
	Event *AggchainfeppreviousTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousTransferVKeyManagerRole)
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
		it.Event = new(AggchainfeppreviousTransferVKeyManagerRole)
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
func (it *AggchainfeppreviousTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainfeppreviousTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousTransferVKeyManagerRoleIterator{contract: _Aggchainfepprevious.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousTransferVKeyManagerRole)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseTransferVKeyManagerRole(log types.Log) (*AggchainfeppreviousTransferVKeyManagerRole, error) {
	event := new(AggchainfeppreviousTransferVKeyManagerRole)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainfeppreviousUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Aggchainfepprevious contract.
type AggchainfeppreviousUpdateAggchainVKeyIterator struct {
	Event *AggchainfeppreviousUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainfeppreviousUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfeppreviousUpdateAggchainVKey)
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
		it.Event = new(AggchainfeppreviousUpdateAggchainVKey)
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
func (it *AggchainfeppreviousUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfeppreviousUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfeppreviousUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Aggchainfepprevious contract.
type AggchainfeppreviousUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainfeppreviousUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainfepprevious.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainfeppreviousUpdateAggchainVKeyIterator{contract: _Aggchainfepprevious.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainfeppreviousUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainfepprevious.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfeppreviousUpdateAggchainVKey)
				if err := _Aggchainfepprevious.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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
func (_Aggchainfepprevious *AggchainfeppreviousFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainfeppreviousUpdateAggchainVKey, error) {
	event := new(AggchainfeppreviousUpdateAggchainVKey)
	if err := _Aggchainfepprevious.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

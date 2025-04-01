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

// AggchainFEPOutputProposal is an auto generated low-level Go binding around an user-defined struct.
type AggchainFEPOutputProposal struct {
	OutputRoot    [32]byte
	Timestamp     *big.Int
	L2BlockNumber *big.Int
}

// AggchainfepMetaData contains all meta data concerning the Aggchainfep contract.
var AggchainfepMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotProposeFutureL2Output\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockNumberLessThanNextBlockNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2BlockTimeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2OutputRootCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingOptimisticModeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimisticModeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupConfigHashMustBeDifferentThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartL2TimestampMustBeGreaterThanCurrentTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmissionIntervalMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultGatewayAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"AcceptOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableOptimisticMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultGatewayFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"oldRollupConfigHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfigHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSubmissionInterval\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSubmissionInterval\",\"type\":\"uint256\"}],\"name\":\"SubmissionIntervalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentOptimisticModeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"TransferOptimisticModeManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableOptimisticMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultGatewayFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structAggchainFEP.OutputProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initializeBytesAggchain\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOptimisticModeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupConfigHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOptimisticModeManager\",\"type\":\"address\"}],\"name\":\"transferOptimisticModeManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"transferVKeyManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_rollupConfigHash\",\"type\":\"bytes32\"}],\"name\":\"updateRollupConfigHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"}],\"name\":\"updateSubmissionInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultGateway\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161371838038061371883398101604081905261002f9161017b565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100a7565b5050506001600160a01b038216905061008b5760405163ae7e6e0f60e01b815260040160405180910390fd5b6001600160a01b031661010052506101ec975050505050505050565b5f54610100900460ff16156101125760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161015610162575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610178575f5ffd5b50565b5f5f5f5f5f60a0868803121561018f575f5ffd5b855161019a81610164565b60208701519095506101ab81610164565b60408701519094506101bc81610164565b60608701519093506101cd81610164565b60808701519092506101de81610164565b809150509295509295909350565b60805160a05160c05160e051610100516134da61023e5f395f818161089a0152610b9d01525f81816106690152818161147c0152611dbc01525f61087301525f61098a01525f6109fd01526134da5ff3fe608060405234801561000f575f5ffd5b50600436106103f2575f3560e01c806370872aa51161020f578063c754c7ed1161012e578063e1a41bcf116100c3578063e7a7ed0211610093578063f851a44011610079578063f851a44014610a5a578063fdbbc19b14610a7f578063ff90407914610a92575f5ffd5b8063e7a7ed0214610a27578063effb847914610a3b575f5ffd5b8063e1a41bcf146109cf578063e279984e146109d8578063e46761c4146109f8578063e631476c14610a1f575f5ffd5b8063d02103ca116100fe578063d02103ca14610985578063d1de856c146109ac578063dc8c4249146109bf578063dcec3348146109c7575f5ffd5b8063c754c7ed14610922578063c89e42df1461094a578063cea5a4c01461095d578063cfa8ed4714610965575f5ffd5b80639ee4afa3116101a4578063ada8f91911610174578063ada8f919146108bc578063adb8696c146108cf578063bdfbed7e146108ef578063bfb193b614610902575f5ffd5b80639ee4afa31461080c578063a25ae5571461081f578063a3c573eb1461086e578063ab0475cf14610895575f5ffd5b806385018182116101df57806385018182146107df57806388786272146107f25780638c3d7301146107fb57806393991af314610803575f5ffd5b806370872aa514610796578063712570221461079f5780637388c436146107b257806381eb0baf146107d7575f5ffd5b80633cbc795b1161031557806360caf7a0116102aa5780636b8616ce1161027a5780636e05d2cd116102605780636e05d2cd146107235780636e7fbce91461072c5780636ff512cc14610783575f5ffd5b80636b8616ce146106fb5780636d9a1c8b1461071a575f5ffd5b806360caf7a0146106bb57806369f16eec146106d85780636a55f66c146106e05780636abcf563146106f3575f5ffd5b806349b7b802116102e557806349b7b80214610664578063527570f11461068b578063529933df146106ab578063542028d5146106b3575f5ffd5b80633cbc795b146105d3578063439fab911461061057806345605267146106235780634599c7881461065c575f5ffd5b80631cf810ee1161038b578063314eb17b1161035b578063314eb17b14610585578063336c9e8114610598578063368c822c146105ab5780633c351e10146105b3575f5ffd5b80631cf810ee146104755780631d0b435e146104ba57806326782247146105455780632c111c0614610565575f5ffd5b806312634900116103c6578063126349001461043f57806315981b291461044757806319451a8f1461044f5780631bdd450c14610462575f5ffd5b80622134cc146103f657806301fcf6a01461040d5780630822dc6114610420578063107bf28c1461042a575b5f5ffd5b6075545b6040519081526020015b60405180910390f35b6103fa61041b366004612bdb565b610ab7565b610428610c21565b005b610432610d00565b6040516104049190612bfb565b610428610d8c565b610428610e6f565b61042861045d366004612c4e565b610f77565b610428610470366004612c76565b6110d8565b6079546104959073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610404565b6105146104c8366004612cbd565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610404565b6001546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6008546104959073ffffffffffffffffffffffffffffffffffffffff1681565b610428610593366004612c4e565b611198565b6104286105a6366004612c76565b6112ce565b61042861139e565b6009546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6009546105fb9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610404565b61042861061e366004612e1a565b61147a565b6007546106439068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610404565b6103fa61177c565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6078546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6074546103fa565b6104326117ed565b6077546106c89060ff1681565b6040519015158152602001610404565b6103fa6117fa565b6103fa6106ee366004612e1a565b61180b565b6071546103fa565b6103fa610709366004612e67565b60066020525f908152604090205481565b6103fa60765481565b6103fa60055481565b6107527e0100000000000000000000000000000000000000000000000000000000000081565b6040517fffff0000000000000000000000000000000000000000000000000000000000009091168152602001610404565b610428610791366004612eaf565b611a38565b6103fa60725481565b6104286107ad366004612ee8565b611b01565b60775461049590610100900473ffffffffffffffffffffffffffffffffffffffff1681565b610428611b33565b6104286107ed366004612eaf565b611c16565b6103fa60735481565b610428611ce8565b6103fa60755481565b61042861081a366004612e1a565b611dba565b61083261082d366004612c76565b611f3a565b60408051825181526020808401516fffffffffffffffffffffffffffffffff908116918301919091529282015190921690820152606001610404565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6104286108ca366004612eaf565b611fcc565b607a546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6104286108fd366004612eaf565b612095565b603d546104959073ffffffffffffffffffffffffffffffffffffffff1681565b60075461064390700100000000000000000000000000000000900467ffffffffffffffff1681565b610428610958366004612f9c565b612171565b6105fb600181565b6002546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b6103fa6109ba366004612c76565b612203565b61042861222c565b6103fa612323565b6103fa60745481565b603c546104959073ffffffffffffffffffffffffffffffffffffffff1681565b6104957f000000000000000000000000000000000000000000000000000000000000000081565b610428612339565b6007546106439067ffffffffffffffff1681565b6103fa610a49366004612bdb565b603e6020525f908152604090205481565b5f546104959062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b610428610a8d366004612eaf565b612448565b603d546106c89074010000000000000000000000000000000000000000900460ff1681565b603d545f9074010000000000000000000000000000000000000000900460ff1615158103610b4d57507fffffffff0000000000000000000000000000000000000000000000000000000081165f908152603e602052604090205480610b48576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610bf7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c1b9190612fce565b92915050565b60795473ffffffffffffffffffffffffffffffffffffffff163314610c72576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60775460ff16610cae576040517f873dabd200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556040517f334fa04f09bf04163481cd42794a867682f0b5ccb521db4fc4dbcca8a1e755ac905f90a1565b60048054610d0d90612fe5565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3990612fe5565b8015610d845780601f10610d5b57610100808354040283529160200191610d84565b820191905f5260205f20905b815481529060010190602001808311610d6757829003601f168201915b505050505081565b607a5473ffffffffffffffffffffffffffffffffffffffff163314610ddd576040517f93f1e09400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60798054607a805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527f9a58f1fef974b760afdc36e96f8d4af9162ba9fec7cd8ce7ca397aa3399f331991015b60405180910390a150565b60785473ffffffffffffffffffffffffffffffffffffffff163314610ec0576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607780546078805473ffffffffffffffffffffffffffffffffffffffff8082166101009081027fffffffffffffffffffffff0000000000000000000000000000000000000000ff861617958690557fffffffffffffffffffffffff000000000000000000000000000000000000000090921690925560408051938290048316808552919094049091166020830152917f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610e64565b603c5473ffffffffffffffffffffffffffffffffffffffff163314610fc8576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610fff576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e602052604090205415611067576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e6020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a7991015b60405180910390a15050565b607754610100900473ffffffffffffffffffffffffffffffffffffffff16331461112e576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611165576040517f4bf41e1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6076546040518291907f5d9ebe9f09b0810b3546b30781ba9a51092b37dd6abada4b830ce54a41ac6a4b905f90a3607655565b603c5473ffffffffffffffffffffffffffffffffffffffff1633146111e9576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f908152603e6020526040902054611250576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152603e60209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b607754610100900473ffffffffffffffffffffffffffffffffffffffff163314611324576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f0361135d576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60745460408051918252602082018390527fc1bf9abfb57ea01ed9ecb4f45e9cefa7ba44b2e6778c3ce7281409999f1af1b2910160405180910390a1607455565b603d5473ffffffffffffffffffffffffffffffffffffffff1633146113ef576040517f05882cf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c8054603d805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe29101610e64565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146114e9576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561153957505f5460ff8083169116105b6115ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255815c169003611682575f5f5f5f5f5f5f5f5f5f8b8060200190518101906116259190613124565b99509950995099509950995099509950995099506116428a61251a565b61167385858585858e8e8e8e7e01000000000000000000000000000000000000000000000000000000000000612773565b50505050505050505050611722565b60ff5f5c166001036116f0575f5f5f5f5f868060200190518101906116a7919061320e565b945094509450945094506116ba8561251a565b6116e6848484847e01000000000000000000000000000000000000000000000000000000000000612928565b5050505050611722565b6040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016110cc565b6071545f90156117e45760718054611796906001906132a3565b815481106117a6576117a66132b6565b5f91825260209091206002909102016001015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6072545b905090565b60038054610d0d90612fe5565b6071545f906117e8906001906132a3565b5f5f5f5f8480602001905181019061182391906132e3565b925092509250611831612323565b81101561186a576040517f541d595b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4261187482612203565b106118ab576040517f0dffe81800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816118e2576040517f80bcf51500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7c01000000000000000000000000000000000000000000000000000000007fffff0000000000000000000000000000000000000000000000000000000000008416175f607161192f6117fa565b8154811061193f5761193f6132b6565b5f91825260209182902060029182020154607654607754925460408051958601939093529184018890526060808501889052608085019190915260ff909216151560f81b60a08401527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000911b1660a182015260b50160405160208183030381529060405280519060200120905060016119d783610ab7565b60405160e09290921b7fffffffff000000000000000000000000000000000000000000000000000000001660208301526024820152604481018290526064016040516020818303038152906040528051906020012095505050505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611a8e576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610e64565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60795473ffffffffffffffffffffffffffffffffffffffff163314611b84576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60775460ff1615611bc1576040517f98b3177900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556040517f26cf5e39429c85f7657b1e1f24aa2eb5a5882942a3f4a0dcd42844579bf7850a905f90a1565b603c5473ffffffffffffffffffffffffffffffffffffffff163314611c67576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255603c546040805191909316815260208101919091527fc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b29101610e64565b60015473ffffffffffffffffffffffffffffffffffffffff163314611d39576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611e29576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f82806020019051810190611e3f91906132e3565b925092505080611e4e60715490565b837fa7aaf2512769da4e444e3de247be2564225c2e7a8f74cfe528e46e17d24868e242604051611e8091815260200190565b60405180910390a4604080516060810182529283526fffffffffffffffffffffffffffffffff42811660208501908152928116918401918252607180546001810182555f91909152935160029094027fa1fcd19bfe8c32a61095b6bfbb2664842857e148fcbb5188386c8cd40348d5b6810194909455915190518216700100000000000000000000000000000000029116177fa1fcd19bfe8c32a61095b6bfbb2664842857e148fcbb5188386c8cd40348d5b79091015550565b604080516060810182525f808252602082018190529181019190915260718281548110611f6957611f696132b6565b5f91825260209182902060408051606081018252600290930290910180548352600101546fffffffffffffffffffffffffffffffff8082169484019490945270010000000000000000000000000000000090049092169181019190915292915050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612022576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610e64565b607754610100900473ffffffffffffffffffffffffffffffffffffffff1633146120eb576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560775460408051610100909204909316815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101610e64565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146121c7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036121d38282613363565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610e649190612bfb565b5f6075546072548361221591906132a3565b61221f919061347a565b607354610c1b9190613491565b603c5473ffffffffffffffffffffffffffffffffffffffff16331461227d576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff166122d1576040517f62de044500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e905f90a1565b5f60745461232f61177c565b6117e89190613491565b603c5473ffffffffffffffffffffffffffffffffffffffff16331461238a576040517fe4d753bd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d5474010000000000000000000000000000000000000000900460ff16156123df576040517f93be805100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517fb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84905f90a1565b60795473ffffffffffffffffffffffffffffffffffffffff163314612499576040517f4382608900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b607a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092556079546040805191909316815260208101919091527ff67c2e74a956fb061c1a9c17172d5a9197efc33c180fac0319ce5cd90702af799101610e64565b8060a001515f03612557576040517fd685d8e700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03612591576040517fff5f860000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b42816080015111156125cf576040517f906872e500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081015161260a576040517f4bf41e1700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60a081015160745580516075556071545f036126ec5760408051606080820183528383015182526080840180516fffffffffffffffffffffffffffffffff9081166020850190815292860180518216958501958652607180546001810182555f91909152945160029095027fa1fcd19bfe8c32a61095b6bfbb2664842857e148fcbb5188386c8cd40348d5b6810195909555925194518116700100000000000000000000000000000000029416939093177fa1fcd19bfe8c32a61095b6bfbb2664842857e148fcbb5188386c8cd40348d5b79092019190915551607255516073555b602081015160765560c0810151607780547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff9384160217905560e090910151607980547fffffffffffffffffffffffff00000000000000000000000000000000000000001691909216179055565b5f54610100900460ff16612809576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016115c1565b6128168a8a8a8a8a612acb565b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000008715150217905583603e5f601084901c7dffff00000000000000000000000000000000000000000000000000000000167fffff0000000000000000000000000000000000000000000000000000000000008716177fffffffff0000000000000000000000000000000000000000000000000000000016815260208101919091526040015f205550603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790555050505050505050565b5f54610100900460ff166129be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016115c1565b603d80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000008715150217905583603e5f601084901c7dffff00000000000000000000000000000000000000000000000000000000167fffff0000000000000000000000000000000000000000000000000000000000008716177fffffffff0000000000000000000000000000000000000000000000000000000016815260208101919091526040015f205550603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055505050565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003612b528382613363565b506004612b5f8282613363565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610b48575f5ffd5b5f60208284031215612beb575f5ffd5b612bf482612bac565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215612c5f575f5ffd5b612c6883612bac565b946020939093013593505050565b5f60208284031215612c86575f5ffd5b5035919050565b7fffff00000000000000000000000000000000000000000000000000000000000081168114612cba575f5ffd5b50565b5f5f60408385031215612cce575f5ffd5b8235612cd981612c8d565b91506020830135612ce981612c8d565b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610100810167ffffffffffffffff81118282101715612d4557612d45612cf4565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612d9257612d92612cf4565b604052919050565b5f67ffffffffffffffff821115612db357612db3612cf4565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f612df1612dec84612d9a565b612d4b565b9050828152838383011115612e04575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215612e2a575f5ffd5b813567ffffffffffffffff811115612e40575f5ffd5b8201601f81018413612e50575f5ffd5b612e5f84823560208401612ddf565b949350505050565b5f60208284031215612e77575f5ffd5b813567ffffffffffffffff81168114612bf4575f5ffd5b73ffffffffffffffffffffffffffffffffffffffff81168114612cba575f5ffd5b5f60208284031215612ebf575f5ffd5b8135612bf481612e8e565b5f82601f830112612ed9575f5ffd5b612bf483833560208501612ddf565b5f5f5f5f5f5f60c08789031215612efd575f5ffd5b8635612f0881612e8e565b95506020870135612f1881612e8e565b9450604087013563ffffffff81168114612f30575f5ffd5b93506060870135612f4081612e8e565b9250608087013567ffffffffffffffff811115612f5b575f5ffd5b612f6789828a01612eca565b92505060a087013567ffffffffffffffff811115612f83575f5ffd5b612f8f89828a01612eca565b9150509295509295509295565b5f60208284031215612fac575f5ffd5b813567ffffffffffffffff811115612fc2575f5ffd5b612e5f84828501612eca565b5f60208284031215612fde575f5ffd5b5051919050565b600181811c90821680612ff957607f821691505b602082108103613030577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8051610b4881612e8e565b5f6101008284031215613052575f5ffd5b61305a612d21565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c08301519091506130a181612e8e565b60c08201526130b260e08301613036565b60e082015292915050565b80518015158114610b48575f5ffd5b8051610b4881612c8d565b5f82601f8301126130e6575f5ffd5b81516130f4612dec82612d9a565b818152846020838601011115613108575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f5f5f5f5f5f6102208b8d03121561313e575f5ffd5b6131488c8c613041565b99506131576101008c016130bd565b6101208c0151909950975061316f6101408c016130cc565b965061317e6101608c01613036565b955061318d6101808c01613036565b945061319c6101a08c01613036565b93506131ab6101c08c01613036565b92506101e08b015167ffffffffffffffff8111156131c7575f5ffd5b6131d38d828e016130d7565b9250506102008b015167ffffffffffffffff8111156131f0575f5ffd5b6131fc8d828e016130d7565b9150509295989b9194979a5092959850565b5f5f5f5f5f6101808688031215613223575f5ffd5b61322d8787613041565b945061323c61010087016130bd565b610120870151610140880151919550935061325681612c8d565b61016087015190925061326881612e8e565b809150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610c1b57610c1b613276565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f5f5f606084860312156132f5575f5ffd5b835161330081612c8d565b602085015160409095015190969495509392505050565b601f82111561335e57805f5260205f20601f840160051c8101602085101561333c5750805b601f840160051c820191505b8181101561335b575f8155600101613348565b50505b505050565b815167ffffffffffffffff81111561337d5761337d612cf4565b6133918161338b8454612fe5565b84613317565b6020601f8211600181146133e2575f83156133ac5750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b17845561335b565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b8281101561342f578785015182556020948501946001909201910161340f565b508482101561346b57868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b8082028115828204841417610c1b57610c1b613276565b80820180821115610c1b57610c1b61327656fea26469706673582212203b77562a5f1f17f1e013a68a006b221b4800dc60239d869db94812edf49233e264736f6c634300081c0033",
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

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) PendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "pendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingVKeyManager(&_Aggchainfep.CallOpts)
}

// PendingVKeyManager is a free data retrieval call binding the contract method 0xbfb193b6.
//
// Solidity: function pendingVKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) PendingVKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.PendingVKeyManager(&_Aggchainfep.CallOpts)
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

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfep *AggchainfepCaller) UseDefaultGateway(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "useDefaultGateway")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfep *AggchainfepSession) UseDefaultGateway() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultGateway(&_Aggchainfep.CallOpts)
}

// UseDefaultGateway is a free data retrieval call binding the contract method 0xff904079.
//
// Solidity: function useDefaultGateway() view returns(bool)
func (_Aggchainfep *AggchainfepCallerSession) UseDefaultGateway() (bool, error) {
	return _Aggchainfep.Contract.UseDefaultGateway(&_Aggchainfep.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCaller) VKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainfep.contract.Call(opts, &out, "vKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepSession) VKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.VKeyManager(&_Aggchainfep.CallOpts)
}

// VKeyManager is a free data retrieval call binding the contract method 0xe279984e.
//
// Solidity: function vKeyManager() view returns(address)
func (_Aggchainfep *AggchainfepCallerSession) VKeyManager() (common.Address, error) {
	return _Aggchainfep.Contract.VKeyManager(&_Aggchainfep.CallOpts)
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

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactor) AcceptVKeyManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "acceptVKeyManagerRole")
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfep *AggchainfepSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptVKeyManagerRole(&_Aggchainfep.TransactOpts)
}

// AcceptVKeyManagerRole is a paid mutator transaction binding the contract method 0x368c822c.
//
// Solidity: function acceptVKeyManagerRole() returns()
func (_Aggchainfep *AggchainfepTransactorSession) AcceptVKeyManagerRole() (*types.Transaction, error) {
	return _Aggchainfep.Contract.AcceptVKeyManagerRole(&_Aggchainfep.TransactOpts)
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

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) DisableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "disableUseDefaultGatewayFlag")
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultGatewayFlag(&_Aggchainfep.TransactOpts)
}

// DisableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xdc8c4249.
//
// Solidity: function disableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) DisableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.DisableUseDefaultGatewayFlag(&_Aggchainfep.TransactOpts)
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

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepTransactor) EnableUseDefaultGatewayFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "enableUseDefaultGatewayFlag")
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultGatewayFlag(&_Aggchainfep.TransactOpts)
}

// EnableUseDefaultGatewayFlag is a paid mutator transaction binding the contract method 0xe631476c.
//
// Solidity: function enableUseDefaultGatewayFlag() returns()
func (_Aggchainfep *AggchainfepTransactorSession) EnableUseDefaultGatewayFlag() (*types.Transaction, error) {
	return _Aggchainfep.Contract.EnableUseDefaultGatewayFlag(&_Aggchainfep.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfep *AggchainfepTransactor) Initialize(opts *bind.TransactOpts, initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "initialize", initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfep *AggchainfepSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, initializeBytesAggchain)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes initializeBytesAggchain) returns()
func (_Aggchainfep *AggchainfepTransactorSession) Initialize(initializeBytesAggchain []byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.Initialize(&_Aggchainfep.TransactOpts, initializeBytesAggchain)
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

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfep *AggchainfepTransactor) TransferVKeyManagerRole(opts *bind.TransactOpts, newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "transferVKeyManagerRole", newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfep *AggchainfepSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferVKeyManagerRole(&_Aggchainfep.TransactOpts, newVKeyManager)
}

// TransferVKeyManagerRole is a paid mutator transaction binding the contract method 0x85018182.
//
// Solidity: function transferVKeyManagerRole(address newVKeyManager) returns()
func (_Aggchainfep *AggchainfepTransactorSession) TransferVKeyManagerRole(newVKeyManager common.Address) (*types.Transaction, error) {
	return _Aggchainfep.Contract.TransferVKeyManagerRole(&_Aggchainfep.TransactOpts, newVKeyManager)
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

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfep *AggchainfepTransactor) UpdateRollupConfigHash(opts *bind.TransactOpts, _rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.contract.Transact(opts, "updateRollupConfigHash", _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfep *AggchainfepSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateRollupConfigHash(&_Aggchainfep.TransactOpts, _rollupConfigHash)
}

// UpdateRollupConfigHash is a paid mutator transaction binding the contract method 0x1bdd450c.
//
// Solidity: function updateRollupConfigHash(bytes32 _rollupConfigHash) returns()
func (_Aggchainfep *AggchainfepTransactorSession) UpdateRollupConfigHash(_rollupConfigHash [32]byte) (*types.Transaction, error) {
	return _Aggchainfep.Contract.UpdateRollupConfigHash(&_Aggchainfep.TransactOpts, _rollupConfigHash)
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

// AggchainfepDisableUseDefaultGatewayFlagIterator is returned from FilterDisableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultGatewayFlag events raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultGatewayFlagIterator struct {
	Event *AggchainfepDisableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepDisableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepDisableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainfepDisableUseDefaultGatewayFlag)
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
func (it *AggchainfepDisableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepDisableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepDisableUseDefaultGatewayFlag represents a DisableUseDefaultGatewayFlag event raised by the Aggchainfep contract.
type AggchainfepDisableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterDisableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainfepDisableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepDisableUseDefaultGatewayFlagIterator{contract: _Aggchainfep.contract, event: "DisableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0x67dd1717a1952df380cb73eecb312e949df6d6a086bd7f88669005341972528e.
//
// Solidity: event DisableUseDefaultGatewayFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchDisableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepDisableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "DisableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepDisableUseDefaultGatewayFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainfep *AggchainfepFilterer) ParseDisableUseDefaultGatewayFlag(log types.Log) (*AggchainfepDisableUseDefaultGatewayFlag, error) {
	event := new(AggchainfepDisableUseDefaultGatewayFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "DisableUseDefaultGatewayFlag", log); err != nil {
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

// AggchainfepEnableUseDefaultGatewayFlagIterator is returned from FilterEnableUseDefaultGatewayFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultGatewayFlag events raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultGatewayFlagIterator struct {
	Event *AggchainfepEnableUseDefaultGatewayFlag // Event containing the contract specifics and raw log

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
func (it *AggchainfepEnableUseDefaultGatewayFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainfepEnableUseDefaultGatewayFlag)
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
		it.Event = new(AggchainfepEnableUseDefaultGatewayFlag)
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
func (it *AggchainfepEnableUseDefaultGatewayFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainfepEnableUseDefaultGatewayFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainfepEnableUseDefaultGatewayFlag represents a EnableUseDefaultGatewayFlag event raised by the Aggchainfep contract.
type AggchainfepEnableUseDefaultGatewayFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultGatewayFlag is a free log retrieval operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainfep *AggchainfepFilterer) FilterEnableUseDefaultGatewayFlag(opts *bind.FilterOpts) (*AggchainfepEnableUseDefaultGatewayFlagIterator, error) {

	logs, sub, err := _Aggchainfep.contract.FilterLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainfepEnableUseDefaultGatewayFlagIterator{contract: _Aggchainfep.contract, event: "EnableUseDefaultGatewayFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultGatewayFlag is a free log subscription operation binding the contract event 0xb6563aed80fde357e737eb0d19f246a58cb6bfd469933d05701ecbad0f2dca84.
//
// Solidity: event EnableUseDefaultGatewayFlag()
func (_Aggchainfep *AggchainfepFilterer) WatchEnableUseDefaultGatewayFlag(opts *bind.WatchOpts, sink chan<- *AggchainfepEnableUseDefaultGatewayFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainfep.contract.WatchLogs(opts, "EnableUseDefaultGatewayFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainfepEnableUseDefaultGatewayFlag)
				if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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
func (_Aggchainfep *AggchainfepFilterer) ParseEnableUseDefaultGatewayFlag(log types.Log) (*AggchainfepEnableUseDefaultGatewayFlag, error) {
	event := new(AggchainfepEnableUseDefaultGatewayFlag)
	if err := _Aggchainfep.contract.UnpackLog(event, "EnableUseDefaultGatewayFlag", log); err != nil {
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

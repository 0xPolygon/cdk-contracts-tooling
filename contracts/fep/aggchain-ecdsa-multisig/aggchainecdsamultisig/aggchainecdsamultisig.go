// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggchainecdsamultisig

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

// AggchainecdsamultisigMetaData contains all meta data concerning the Aggchainecdsamultisig contract.
var AggchainecdsamultisigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"},{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"_aggLayerGateway\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AggchainManagerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersHashNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainSignersTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAggchainSignersArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndicesNotInDescendingOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggLayerGatewayAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAggchainType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoTreeLeafCountInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAggchainManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVKeyManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnedAggchainVKeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerURLCannotBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdTooHighAfterRemoval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultSignersAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseDefaultVkeysAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroValueAggchainVKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"AcceptAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVKeyManager\",\"type\":\"address\"}],\"name\":\"AcceptVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"AddAggchainVKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DisableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultSignersFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EnableUseDefaultVkeysFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OnVerifyPessimisticECDSAMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"aggchainSigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainSignersHash\",\"type\":\"bytes32\"}],\"name\":\"SignersAndThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentAggchainManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAggchainManager\",\"type\":\"address\"}],\"name\":\"TransferAggchainManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentVKeyManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVKeyManager\",\"type\":\"address\"}],\"name\":\"TransferVKeyManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousAggchainVKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"UpdateAggchainVKey\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGCHAIN_ECDSA_MULTISIG_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGGCHAIN_TYPE\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONSENSUS_TYPE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacypendingVKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_legacyvKeyManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"newAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"addOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggLayerGateway\",\"outputs\":[{\"internalType\":\"contractIAggLayerGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"aggchainSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggchainSignersHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultSignersFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableUseDefaultVkeysFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"getAggchainParamsAndVKeySelector\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggchainSignersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainTypeFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"aggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"aggchainVKeyVersion\",\"type\":\"bytes2\"},{\"internalType\":\"bytes2\",\"name\":\"aggchainType\",\"type\":\"bytes2\"}],\"name\":\"getAggchainVKeySelector\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"getAggchainVKeyVersionFromSelector\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"initAggchainManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_useDefaultSigners\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateFromPessimisticConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"aggchainData\",\"type\":\"bytes\"}],\"name\":\"onVerifyPessimistic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"}],\"name\":\"ownedAggchainVKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"ownedAggchainVKey\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAggchainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToURLs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAggchainManager\",\"type\":\"address\"}],\"name\":\"transferAggchainManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"aggchainVKeySelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"updatedAggchainVKey\",\"type\":\"bytes32\"}],\"name\":\"updateOwnedAggchainVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structAggchainBase.RemoveSignerInfo[]\",\"name\":\"_signersToRemove\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structAggchainBase.SignerInfo[]\",\"name\":\"_signersToAdd\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"updateSignersAndThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useDefaultVkeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b50604051613d09380380613d0983398101604081905261002f916101c4565b6001600160a01b0380861660a05280851660805280841660c052821660e05284848484848484848461005f6100f0565b505050506001600160a01b038116158061008057506001600160a01b038516155b8061009257506001600160a01b038416155b806100a457506001600160a01b038316155b806100b657506001600160a01b038216155b156100d45760405163f6b2911f60e01b815260040160405180910390fd5b6001600160a01b03166101005250610235975050505050505050565b5f54610100900460ff161561015b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811610156101ab575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c1575f5ffd5b50565b5f5f5f5f5f60a086880312156101d8575f5ffd5b85516101e3816101ad565b60208701519095506101f4816101ad565b6040870151909450610205816101ad565b6060870151909350610216816101ad565b6080870151909250610227816101ad565b809150509295509295909350565b60805160a05160c05160e05161010051613a5f6102aa5f395f81816108a301528181610b560152818161115f01528181611556015281816117f7015261222801525f818161067f0152818161189b01528181611c730152611e0f01525f61087c01525f61097601525f61099d0152613a5f5ff3fe608060405234801561000f575f5ffd5b5060043610610393575f3560e01c80636ff512cc116101df578063c754c7ed11610109578063e7a7ed02116100a9578063f51f563a11610079578063f51f563a14610a0b578063f851a44014610a1e578063fc5014d614610a43578063fd7d249314610a68575f5ffd5b8063e7a7ed02146109bf578063e90a3409146109d3578063efe6c9f4146109e4578063effb8479146109ec575f5ffd5b8063cea5a4c0116100e4578063cea5a4c014610949578063cfa8ed4714610951578063d02103ca14610971578063e46761c414610998575f5ffd5b8063c754c7ed14610906578063c89e42df1461092e578063ca69e7dc14610941575f5ffd5b80638c3d73011161017f578063ada8f9191161014f578063ada8f919146108c5578063b3a326f7146108d8578063bdfbed7e146108eb578063be647d03146108fe575f5ffd5b80638c3d73011461085c5780639ee4afa314610864578063a3c573eb14610877578063ab0475cf1461089e575f5ffd5b806374f0b0c1116101ba57806374f0b0c1146107e5578063759664d4146108055780637df73e271461084157806387b2b55914610854575f5ffd5b80636ff512cc1461079f57806371257022146107b25780637388c436146107c5575f5ffd5b80633cbc795b116102c057806354fd4d50116102605780636a55f66c116102305780636a55f66c1461075d5780636b8616ce146107705780636e05d2cd1461078f5780636e7fbce914610798575f5ffd5b806354fd4d50146106c95780635ecaca2b14610702578063680196ed14610722578063697427f61461074a575f5ffd5b8063456052671161029b578063456052671461064157806349b7b8021461067a578063527570f1146106a1578063542028d5146106c1575f5ffd5b80633cbc795b146105e65780633e1e01211461062357806342cde4e814610638575f5ffd5b80631d0b435e11610336578063314eb17b11610306578063314eb17b1461058d57806335acd6c2146105a057806336cd6b5b146105b35780633c351e10146105c6575f5ffd5b80631d0b435e1461043657806326782247146104c157806326f9b76d146105065780632c111c061461056d575f5ffd5b806315981b291161037157806315981b29146103dc57806317b7a9f0146103e4578063188d9180146103ed57806319451a8f14610423575f5ffd5b806301fcf6a014610397578063107bf28c146103bd5780631489e707146103d2575b5f5ffd5b6103aa6103a5366004612e88565b610a70565b6040519081526020015b60405180910390f35b6103c5610bda565b6040516103b49190612ea8565b6103da610c66565b005b6103da610d75565b6103aa60455481565b603e54610413907501000000000000000000000000000000000000000000900460ff1681565b60405190151581526020016103b4565b6103da610431366004612efb565b610e53565b610490610444366004612f52565b60101c7dffff00000000000000000000000000000000000000000000000000000000167fffff000000000000000000000000000000000000000000000000000000000000919091161790565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103b4565b6001546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016103b4565b61053c610514366004612e88565b60101b7fffff0000000000000000000000000000000000000000000000000000000000001690565b6040517fffff00000000000000000000000000000000000000000000000000000000000090911681526020016103b4565b6008546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b6103da61059b366004612efb565b610fb3565b6104e16105ae366004612f83565b6110e9565b6103c56105c1366004612fc6565b61111e565b6009546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b60095461060e9074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff90911681526020016103b4565b61062b611136565b6040516103b49190612fe1565b6103aa60445481565b6007546106619068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016103b4565b6104e17f000000000000000000000000000000000000000000000000000000000000000081565b6040546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b6103c561127a565b60408051808201909152600681527f76312e302e30000000000000000000000000000000000000000000000000000060208201526103c5565b603d546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b610735610730366004613155565b611287565b604080519283526020830191909152016103b4565b6103da6107583660046132fa565b6112cd565b6103aa61076b366004613155565b61152c565b6103aa61077e3660046133d5565b60066020525f908152604090205481565b6103aa60055481565b61053c5f81565b6103da6107ad366004612fc6565b611691565b6103da6107c03660046133fc565b61175a565b603f546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b603e546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b6103c56040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b61041361084f366004612fc6565b61178c565b6103da611899565b6103da611b9e565b6103da6108723660046134b0565b611c71565b6104e17f000000000000000000000000000000000000000000000000000000000000000081565b6104e17f000000000000000000000000000000000000000000000000000000000000000081565b6103da6108d3366004612fc6565b611d44565b6103da6108e6366004612fc6565b611e0d565b6103da6108f9366004612fc6565b611f43565b6103da61205d565b60075461066190700100000000000000000000000000000000900467ffffffffffffffff1681565b6103da61093c36600461351e565b61216e565b6103aa612200565b61060e600181565b6002546104e19073ffffffffffffffffffffffffffffffffffffffff1681565b6104e17f000000000000000000000000000000000000000000000000000000000000000081565b6104e17f000000000000000000000000000000000000000000000000000000000000000081565b6007546106619067ffffffffffffffff1681565b61053c6109e1366004612e88565b90565b6103da6122ba565b6103aa6109fa366004612e88565b60416020525f908152604090205481565b6103da610a19366004613550565b6123b1565b5f546104e19062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b603e546104139074010000000000000000000000000000000000000000900460ff1681565b6103da612412565b603e545f9074010000000000000000000000000000000000000000900460ff1615158103610b0657507fffffffff0000000000000000000000000000000000000000000000000000000081165f9081526041602052604090205480610b01576040517f925e5a3a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040517f6cabfdab0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690636cabfdab90602401602060405180830381865afa158015610bb0573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bd49190613643565b92915050565b60048054610be79061365a565b80601f0160208091040260200160405190810160405280929190818152602001828054610c139061365a565b8015610c5e5780601f10610c3557610100808354040283529160200191610c5e565b820191905f5260205f20905b815481529060010190602001808311610c4157829003601f168201915b505050505081565b603f5473ffffffffffffffffffffffffffffffffffffffff163314610cb7576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1615610d0c576040517f9573504e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829905f90a1565b60405473ffffffffffffffffffffffffffffffffffffffff163314610dc6576040517f3ac87ac900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80546040805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff000000000000000000000000000000000000000080861682179096559490911682558151921680835260208301939093527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f91015b60405180910390a150565b603f5473ffffffffffffffffffffffffffffffffffffffff163314610ea4576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80610edb576040517fe1dbcf2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205415610f43576040517fe3cc761000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f81815260416020908152604091829020849055815192835282018390527f6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79910160405180910390a15050565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611004576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f9081526041602052604090205461106b576040517ff360deaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082165f818152604160209081526040918290208054908590558251938452908301819052908201839052907f0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f9060600160405180910390a1505050565b604281815481106110f8575f80fd5b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60436020525f908152604090208054610be79061365a565b603e546060907501000000000000000000000000000000000000000000900460ff161561120f577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633e1e01216040518163ffffffff1660e01b81526004015f60405180830381865afa1580156111c5573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261120a91908101906136ab565b905090565b604280548060200260200160405190810160405280929190818152602001828054801561127057602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611245575b5050505050905090565b60038054610be79061365a565b5f5f82515f146112c3576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505f928392509050565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461131e576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561136e57505f5460ff8083169116105b6113ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c161561146d576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61147e89898989895f8a818061250a565b604080515f808252602082019092526114c4916114bc565b604080518082019091525f80825260208201528152602001906001900390816114965790505b508484612632565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050565b603e545f9081907501000000000000000000000000000000000000000000900460ff16156115e8577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166312b941ca6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115e19190613643565b9050611623565b5060455480611623576040517fdd41f1ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f61162e85611287565b6040517c010000000000000000000000000000000000000000000000000000000060208201526024810183905260448101829052606481018690529193509150608401604051602081830303815290604052805190602001209350505050919050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146116e7576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610e48565b6040517ff57ac68300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e545f907501000000000000000000000000000000000000000000900460ff1615611860576040517f7df73e2700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f00000000000000000000000000000000000000000000000000000000000000001690637df73e2790602401602060405180830381865afa15801561183c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bd49190613745565b73ffffffffffffffffffffffffffffffffffffffff82165f908152604360205260408120805461188f9061365a565b9050119050919050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611908576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805460ff16907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00815c168217905d505f54600290610100900460ff1615801561195857505f5460ff8083169116105b6119e4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016113f6565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff80841691909117610100178255905c16600114611a54576040517fadc06ae700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff62010000909304831617905560025460038054611b37939290921691611ab69061365a565b80601f0160208091040260200160405190810160405280929190818152602001828054611ae29061365a565b8015611b2d5780601f10611b0457610100808354040283529160200191611b2d565b820191905f5260205f20905b815481529060010190602001808311611b1057829003601f168201915b505050505061280e565b6001604455611b4461295a565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610e48565b60015473ffffffffffffffffffffffffffffffffffffffff163314611bef576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e906020015b60405180910390a1565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611ce0576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8015611d18576040517f3063965400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8905f90a15050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611d9a576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610e48565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611e7c576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116611ec9576040517fd6bdac3f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604080515f815260208101929092527f67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f9101610e48565b603f5473ffffffffffffffffffffffffffffffffffffffff163314611f94576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116611fe1576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182178355603f5483519116815260208101919091527fa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab69101610e48565b603f5473ffffffffffffffffffffffffffffffffffffffff1633146120ae576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff1615612104576040517f278d998800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790556040517f67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b905f90a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146121c4576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036121d082826137ab565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610e489190612ea8565b603e545f907501000000000000000000000000000000000000000000900460ff16156122b3577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ca69e7dc6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561228f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061120a9190613643565b5060425490565b603f5473ffffffffffffffffffffffffffffffffffffffff16331461230b576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e5474010000000000000000000000000000000000000000900460ff1661235f576040517fa2c45f8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07905f90a1565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612402576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61240d838383612632565b505050565b603f5473ffffffffffffffffffffffffffffffffffffffff163314612463576040517f660a7ce500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e547501000000000000000000000000000000000000000000900460ff166124b8576040517f5aa930a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603e80547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1690556040517f4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be905f90a1565b5f54610100900460ff166125a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016113f6565b73ffffffffffffffffffffffffffffffffffffffff891615806125d7575073ffffffffffffffffffffffffffffffffffffffff8816155b1561260e576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61261b89898989896129de565b61262784848484612abf565b505050505050505050565b6001835111156126da575f5b6001845161264c91906138ef565b8110156126d8578361265f826001613902565b8151811061266f5761266f613915565b60200260200101516020015184828151811061268d5761268d613915565b602002602001015160200151116126d0576040517fb9a11d3100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60010161263e565b505b5f5b8351811015612730576127288482815181106126fa576126fa613915565b60200260200101515f015185838151811061271757612717613915565b602002602001015160200151612c15565b6001016126dc565b505f5b82518110156127875761277f83828151811061275157612751613915565b60200260200101515f015184838151811061276e5761276e613915565b60200260200101516020015161280e565b600101612733565b5060425460ff10156127c5576040517f5a7f382c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604254811115612801576040517faabd5a0900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604481905561240d61295a565b73ffffffffffffffffffffffffffffffffffffffff821661285b576040517f7b3a0df600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80515f03612895576040517f8715f5fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61289e8261178c565b156128d5576040517f38615ecc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60428054600181019091557f38dfe4635b27babeca8be38d3b448cb5161a639b899a14825ba9c8d7892eb8c30180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84169081179091555f90815260436020526040902061240d82826137ab565b6044546042604051602001612970929190613942565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052805160209091012060458190556044547f66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa74392611c679260429291613995565b5f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff8881169190910291909117909155600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000169186169190911790556003612a6583826137ab565b506004612a7282826137ab565b5050600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9390931692909217909155505050565b5f54610100900460ff16612b55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016113f6565b603e80547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000951515959095027fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1694909417750100000000000000000000000000000000000000000093151593909302929092179092557fffffffff00000000000000000000000000000000000000000000000000000000165f90815260416020526040902055565b604254808210612c51576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1660428381548110612c7b57612c7b613915565b5f9182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614612cd3576040517fd244b30700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83165f908152604360205260408120612d0091612e07565b6042612d0d6001836138ef565b81548110612d1d57612d1d613915565b5f918252602090912001546042805473ffffffffffffffffffffffffffffffffffffffff9092169184908110612d5557612d55613915565b905f5260205f20015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506042805480612dab57612dab6139fc565b5f8281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505050565b508054612e139061365a565b5f825580601f10612e22575050565b601f0160209004905f5260205f2090810190612e3e9190612e41565b50565b5b80821115612e55575f8155600101612e42565b5090565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610b01575f5ffd5b5f60208284031215612e98575f5ffd5b612ea182612e59565b9392505050565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60408385031215612f0c575f5ffd5b612f1583612e59565b946020939093013593505050565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610b01575f5ffd5b5f5f60408385031215612f63575f5ffd5b612f6c83612f23565b9150612f7a60208401612f23565b90509250929050565b5f60208284031215612f93575f5ffd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff81168114612e3e575f5ffd5b8035610b0181612f9a565b5f60208284031215612fd6575f5ffd5b8135612ea181612f9a565b602080825282518282018190525f918401906040840190835b8181101561302e57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612ffa565b509095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff8111828210171561308957613089613039565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156130d6576130d6613039565b604052919050565b5f5f67ffffffffffffffff8411156130f8576130f8613039565b50601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200161312b8161308f565b91505082815283838301111561313f575f5ffd5b828260208301375f602084830101529392505050565b5f60208284031215613165575f5ffd5b813567ffffffffffffffff81111561317b575f5ffd5b8201601f8101841361318b575f5ffd5b61319a848235602084016130de565b949350505050565b5f82601f8301126131b1575f5ffd5b612ea1838335602085016130de565b8015158114612e3e575f5ffd5b8035610b01816131c0565b5f67ffffffffffffffff8211156131f1576131f1613039565b5060051b60200190565b5f82601f83011261320a575f5ffd5b813561321d613218826131d8565b61308f565b8082825260208201915060208360051b86010192508583111561323e575f5ffd5b602085015b838110156132f057803567ffffffffffffffff811115613261575f5ffd5b860160408189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613294575f5ffd5b61329c613066565b60208201356132aa81612f9a565b8152604082013567ffffffffffffffff8111156132c5575f5ffd5b6132d48a6020838601016131a2565b6020830152508085525050602083019250602081019050613243565b5095945050505050565b5f5f5f5f5f5f5f5f610100898b031215613312575f5ffd5b61331b89612fbb565b975061332960208a01612fbb565b965061333760408a01612fbb565b9550606089013567ffffffffffffffff811115613352575f5ffd5b61335e8b828c016131a2565b955050608089013567ffffffffffffffff81111561337a575f5ffd5b6133868b828c016131a2565b94505061339560a08a016131cd565b925060c089013567ffffffffffffffff8111156133b0575f5ffd5b6133bc8b828c016131fb565b989b979a50959894979396929550929360e00135925050565b5f602082840312156133e5575f5ffd5b813567ffffffffffffffff81168114612ea1575f5ffd5b5f5f5f5f5f5f60c08789031215613411575f5ffd5b863561341c81612f9a565b9550602087013561342c81612f9a565b9450604087013563ffffffff81168114613444575f5ffd5b9350606087013561345481612f9a565b9250608087013567ffffffffffffffff81111561346f575f5ffd5b61347b89828a016131a2565b92505060a087013567ffffffffffffffff811115613497575f5ffd5b6134a389828a016131a2565b9150509295509295509295565b5f5f602083850312156134c1575f5ffd5b823567ffffffffffffffff8111156134d7575f5ffd5b8301601f810185136134e7575f5ffd5b803567ffffffffffffffff8111156134fd575f5ffd5b85602082840101111561350e575f5ffd5b6020919091019590945092505050565b5f6020828403121561352e575f5ffd5b813567ffffffffffffffff811115613544575f5ffd5b61319a848285016131a2565b5f5f5f60608486031215613562575f5ffd5b833567ffffffffffffffff811115613578575f5ffd5b8401601f81018613613588575f5ffd5b8035613596613218826131d8565b8082825260208201915060208360061b8501019250888311156135b7575f5ffd5b6020840193505b82841015613607576040848a0312156135d5575f5ffd5b6135dd613066565b84356135e881612f9a565b81526020858101358183015290835260409094019391909101906135be565b9550505050602084013567ffffffffffffffff811115613625575f5ffd5b613631868287016131fb565b93969395505050506040919091013590565b5f60208284031215613653575f5ffd5b5051919050565b600181811c9082168061366e57607f821691505b6020821081036136a5577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b5f602082840312156136bb575f5ffd5b815167ffffffffffffffff8111156136d1575f5ffd5b8201601f810184136136e1575f5ffd5b80516136ef613218826131d8565b8082825260208201915060208360051b850101925086831115613710575f5ffd5b6020840193505b8284101561373b57835161372a81612f9a565b825260209384019390910190613717565b9695505050505050565b5f60208284031215613755575f5ffd5b8151612ea1816131c0565b601f82111561240d57805f5260205f20601f840160051c810160208510156137855750805b601f840160051c820191505b818110156137a4575f8155600101613791565b5050505050565b815167ffffffffffffffff8111156137c5576137c5613039565b6137d9816137d3845461365a565b84613760565b6020601f82116001811461382a575f83156137f45750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b1784556137a4565b5f848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b828110156138775787850151825560209485019460019092019101613857565b50848210156138b357868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610bd457610bd46138c2565b80820180821115610bd457610bd46138c2565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8281525f602082018354845f5260205f205f5b8281101561398957815473ffffffffffffffffffffffffffffffffffffffff16845260209093019260019182019101613955565b50919695505050505050565b606080825284549082018190525f8581526020812090916080840190835b818110156139e757835473ffffffffffffffffffffffffffffffffffffffff168352600193840193602090930192016139b3565b50506020840195909552505060400152919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea2646970667358221220e260c777a8d20c71588294b5cc6edf2510f3e038ca154bf77fad6fcb17a09dcc64736f6c634300081c0033",
}

// AggchainecdsamultisigABI is the input ABI used to generate the binding from.
// Deprecated: Use AggchainecdsamultisigMetaData.ABI instead.
var AggchainecdsamultisigABI = AggchainecdsamultisigMetaData.ABI

// AggchainecdsamultisigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AggchainecdsamultisigMetaData.Bin instead.
var AggchainecdsamultisigBin = AggchainecdsamultisigMetaData.Bin

// DeployAggchainecdsamultisig deploys a new Ethereum contract, binding an instance of Aggchainecdsamultisig to it.
func DeployAggchainecdsamultisig(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address, _aggLayerGateway common.Address) (common.Address, *types.Transaction, *Aggchainecdsamultisig, error) {
	parsed, err := AggchainecdsamultisigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AggchainecdsamultisigBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager, _aggLayerGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Aggchainecdsamultisig{AggchainecdsamultisigCaller: AggchainecdsamultisigCaller{contract: contract}, AggchainecdsamultisigTransactor: AggchainecdsamultisigTransactor{contract: contract}, AggchainecdsamultisigFilterer: AggchainecdsamultisigFilterer{contract: contract}}, nil
}

// Aggchainecdsamultisig is an auto generated Go binding around an Ethereum contract.
type Aggchainecdsamultisig struct {
	AggchainecdsamultisigCaller     // Read-only binding to the contract
	AggchainecdsamultisigTransactor // Write-only binding to the contract
	AggchainecdsamultisigFilterer   // Log filterer for contract events
}

// AggchainecdsamultisigCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggchainecdsamultisigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsamultisigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggchainecdsamultisigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsamultisigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggchainecdsamultisigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggchainecdsamultisigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggchainecdsamultisigSession struct {
	Contract     *Aggchainecdsamultisig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AggchainecdsamultisigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggchainecdsamultisigCallerSession struct {
	Contract *AggchainecdsamultisigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AggchainecdsamultisigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggchainecdsamultisigTransactorSession struct {
	Contract     *AggchainecdsamultisigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AggchainecdsamultisigRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggchainecdsamultisigRaw struct {
	Contract *Aggchainecdsamultisig // Generic contract binding to access the raw methods on
}

// AggchainecdsamultisigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggchainecdsamultisigCallerRaw struct {
	Contract *AggchainecdsamultisigCaller // Generic read-only contract binding to access the raw methods on
}

// AggchainecdsamultisigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggchainecdsamultisigTransactorRaw struct {
	Contract *AggchainecdsamultisigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggchainecdsamultisig creates a new instance of Aggchainecdsamultisig, bound to a specific deployed contract.
func NewAggchainecdsamultisig(address common.Address, backend bind.ContractBackend) (*Aggchainecdsamultisig, error) {
	contract, err := bindAggchainecdsamultisig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggchainecdsamultisig{AggchainecdsamultisigCaller: AggchainecdsamultisigCaller{contract: contract}, AggchainecdsamultisigTransactor: AggchainecdsamultisigTransactor{contract: contract}, AggchainecdsamultisigFilterer: AggchainecdsamultisigFilterer{contract: contract}}, nil
}

// NewAggchainecdsamultisigCaller creates a new read-only instance of Aggchainecdsamultisig, bound to a specific deployed contract.
func NewAggchainecdsamultisigCaller(address common.Address, caller bind.ContractCaller) (*AggchainecdsamultisigCaller, error) {
	contract, err := bindAggchainecdsamultisig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigCaller{contract: contract}, nil
}

// NewAggchainecdsamultisigTransactor creates a new write-only instance of Aggchainecdsamultisig, bound to a specific deployed contract.
func NewAggchainecdsamultisigTransactor(address common.Address, transactor bind.ContractTransactor) (*AggchainecdsamultisigTransactor, error) {
	contract, err := bindAggchainecdsamultisig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigTransactor{contract: contract}, nil
}

// NewAggchainecdsamultisigFilterer creates a new log filterer instance of Aggchainecdsamultisig, bound to a specific deployed contract.
func NewAggchainecdsamultisigFilterer(address common.Address, filterer bind.ContractFilterer) (*AggchainecdsamultisigFilterer, error) {
	contract, err := bindAggchainecdsamultisig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigFilterer{contract: contract}, nil
}

// bindAggchainecdsamultisig binds a generic wrapper to an already deployed contract.
func bindAggchainecdsamultisig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggchainecdsamultisigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainecdsamultisig *AggchainecdsamultisigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainecdsamultisig.Contract.AggchainecdsamultisigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainecdsamultisig *AggchainecdsamultisigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AggchainecdsamultisigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainecdsamultisig *AggchainecdsamultisigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AggchainecdsamultisigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggchainecdsamultisig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.contract.Transact(opts, method, params...)
}

// AGGCHAINECDSAMULTISIGVERSION is a free data retrieval call binding the contract method 0x759664d4.
//
// Solidity: function AGGCHAIN_ECDSA_MULTISIG_VERSION() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AGGCHAINECDSAMULTISIGVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "AGGCHAIN_ECDSA_MULTISIG_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AGGCHAINECDSAMULTISIGVERSION is a free data retrieval call binding the contract method 0x759664d4.
//
// Solidity: function AGGCHAIN_ECDSA_MULTISIG_VERSION() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AGGCHAINECDSAMULTISIGVERSION() (string, error) {
	return _Aggchainecdsamultisig.Contract.AGGCHAINECDSAMULTISIGVERSION(&_Aggchainecdsamultisig.CallOpts)
}

// AGGCHAINECDSAMULTISIGVERSION is a free data retrieval call binding the contract method 0x759664d4.
//
// Solidity: function AGGCHAIN_ECDSA_MULTISIG_VERSION() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AGGCHAINECDSAMULTISIGVERSION() (string, error) {
	return _Aggchainecdsamultisig.Contract.AGGCHAINECDSAMULTISIGVERSION(&_Aggchainecdsamultisig.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AGGCHAINTYPE(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "AGGCHAIN_TYPE")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.AGGCHAINTYPE(&_Aggchainecdsamultisig.CallOpts)
}

// AGGCHAINTYPE is a free data retrieval call binding the contract method 0x6e7fbce9.
//
// Solidity: function AGGCHAIN_TYPE() view returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AGGCHAINTYPE() ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.AGGCHAINTYPE(&_Aggchainecdsamultisig.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) CONSENSUSTYPE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "CONSENSUS_TYPE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainecdsamultisig.Contract.CONSENSUSTYPE(&_Aggchainecdsamultisig.CallOpts)
}

// CONSENSUSTYPE is a free data retrieval call binding the contract method 0xcea5a4c0.
//
// Solidity: function CONSENSUS_TYPE() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) CONSENSUSTYPE() (uint32, error) {
	return _Aggchainecdsamultisig.Contract.CONSENSUSTYPE(&_Aggchainecdsamultisig.CallOpts)
}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) LegacypendingVKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "_legacypendingVKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) LegacypendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.LegacypendingVKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// LegacypendingVKeyManager is a free data retrieval call binding the contract method 0x74f0b0c1.
//
// Solidity: function _legacypendingVKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) LegacypendingVKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.LegacypendingVKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) LegacyvKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "_legacyvKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) LegacyvKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.LegacyvKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// LegacyvKeyManager is a free data retrieval call binding the contract method 0x5ecaca2b.
//
// Solidity: function _legacyvKeyManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) LegacyvKeyManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.LegacyvKeyManager(&_Aggchainecdsamultisig.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Admin() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Admin(&_Aggchainecdsamultisig.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Admin() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Admin(&_Aggchainecdsamultisig.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggLayerGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggLayerGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggLayerGateway(&_Aggchainecdsamultisig.CallOpts)
}

// AggLayerGateway is a free data retrieval call binding the contract method 0xab0475cf.
//
// Solidity: function aggLayerGateway() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggLayerGateway() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggLayerGateway(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainManager(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainManager is a free data retrieval call binding the contract method 0x7388c436.
//
// Solidity: function aggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainManager(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainSigners(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// AggchainSigners is a free data retrieval call binding the contract method 0x35acd6c2.
//
// Solidity: function aggchainSigners(uint256 ) view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainSigners(arg0 *big.Int) (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.AggchainSigners(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) AggchainSignersHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "aggchainSignersHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AggchainSignersHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.AggchainSignersHash(&_Aggchainecdsamultisig.CallOpts)
}

// AggchainSignersHash is a free data retrieval call binding the contract method 0x17b7a9f0.
//
// Solidity: function aggchainSignersHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) AggchainSignersHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.AggchainSignersHash(&_Aggchainecdsamultisig.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) BridgeAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.BridgeAddress(&_Aggchainecdsamultisig.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) BridgeAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.BridgeAddress(&_Aggchainecdsamultisig.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.ForceBatchAddress(&_Aggchainecdsamultisig.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.ForceBatchAddress(&_Aggchainecdsamultisig.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.ForceBatchTimeout(&_Aggchainecdsamultisig.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.ForceBatchTimeout(&_Aggchainecdsamultisig.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.ForcedBatches(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.ForcedBatches(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GasTokenAddress(&_Aggchainecdsamultisig.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GasTokenAddress() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GasTokenAddress(&_Aggchainecdsamultisig.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainecdsamultisig.Contract.GasTokenNetwork(&_Aggchainecdsamultisig.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GasTokenNetwork() (uint32, error) {
	return _Aggchainecdsamultisig.Contract.GasTokenNetwork(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainHash(opts *bind.CallOpts, aggchainData []byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainHash", aggchainData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainHash(&_Aggchainecdsamultisig.CallOpts, aggchainData)
}

// GetAggchainHash is a free data retrieval call binding the contract method 0x6a55f66c.
//
// Solidity: function getAggchainHash(bytes aggchainData) view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainHash(aggchainData []byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainHash(&_Aggchainecdsamultisig.CallOpts, aggchainData)
}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainParamsAndVKeySelector(opts *bind.CallOpts, aggchainData []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainParamsAndVKeySelector", aggchainData)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainParamsAndVKeySelector(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainParamsAndVKeySelector(&_Aggchainecdsamultisig.CallOpts, aggchainData)
}

// GetAggchainParamsAndVKeySelector is a free data retrieval call binding the contract method 0x680196ed.
//
// Solidity: function getAggchainParamsAndVKeySelector(bytes aggchainData) pure returns(bytes32, bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainParamsAndVKeySelector(aggchainData []byte) ([32]byte, [32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainParamsAndVKeySelector(&_Aggchainecdsamultisig.CallOpts, aggchainData)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainSigners() ([]common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSigners(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainSigners is a free data retrieval call binding the contract method 0x3e1e0121.
//
// Solidity: function getAggchainSigners() view returns(address[])
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainSigners() ([]common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSigners(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainSignersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainSignersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSignersCount(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainSignersCount is a free data retrieval call binding the contract method 0xca69e7dc.
//
// Solidity: function getAggchainSignersCount() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainSignersCount() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainSignersCount(&_Aggchainecdsamultisig.CallOpts)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainTypeFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainTypeFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainTypeFromSelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GetAggchainTypeFromSelector is a free data retrieval call binding the contract method 0x26f9b76d.
//
// Solidity: function getAggchainTypeFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainTypeFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainTypeFromSelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainVKey(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainVKey", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKey(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKey is a free data retrieval call binding the contract method 0x01fcf6a0.
//
// Solidity: function getAggchainVKey(bytes4 aggchainVKeySelector) view returns(bytes32 aggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainVKey(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKey(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainVKeySelector(opts *bind.CallOpts, aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainVKeySelector", aggchainVKeyVersion, aggchainType)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKeySelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeySelector is a free data retrieval call binding the contract method 0x1d0b435e.
//
// Solidity: function getAggchainVKeySelector(bytes2 aggchainVKeyVersion, bytes2 aggchainType) pure returns(bytes4)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainVKeySelector(aggchainVKeyVersion [2]byte, aggchainType [2]byte) ([4]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKeySelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeyVersion, aggchainType)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GetAggchainVKeyVersionFromSelector(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([2]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "getAggchainVKeyVersionFromSelector", aggchainVKeySelector)

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GetAggchainVKeyVersionFromSelector is a free data retrieval call binding the contract method 0xe90a3409.
//
// Solidity: function getAggchainVKeyVersionFromSelector(bytes4 aggchainVKeySelector) pure returns(bytes2)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GetAggchainVKeyVersionFromSelector(aggchainVKeySelector [4]byte) ([2]byte, error) {
	return _Aggchainecdsamultisig.Contract.GetAggchainVKeyVersionFromSelector(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GlobalExitRootManager(&_Aggchainecdsamultisig.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.GlobalExitRootManager(&_Aggchainecdsamultisig.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Initialize0(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainecdsamultisig.Contract.Initialize0(&_Aggchainecdsamultisig.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a free data retrieval call binding the contract method 0x71257022.
//
// Solidity: function initialize(address , address , uint32 , address , string , string ) pure returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Initialize0(arg0 common.Address, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 string, arg5 string) error {
	return _Aggchainecdsamultisig.Contract.Initialize0(&_Aggchainecdsamultisig.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) IsSigner(opts *bind.CallOpts, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "isSigner", _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) IsSigner(_signer common.Address) (bool, error) {
	return _Aggchainecdsamultisig.Contract.IsSigner(&_Aggchainecdsamultisig.CallOpts, _signer)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address _signer) view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) IsSigner(_signer common.Address) (bool, error) {
	return _Aggchainecdsamultisig.Contract.IsSigner(&_Aggchainecdsamultisig.CallOpts, _signer)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.LastAccInputHash(&_Aggchainecdsamultisig.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.LastAccInputHash(&_Aggchainecdsamultisig.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) LastForceBatch() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.LastForceBatch(&_Aggchainecdsamultisig.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) LastForceBatch() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.LastForceBatch(&_Aggchainecdsamultisig.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.LastForceBatchSequenced(&_Aggchainecdsamultisig.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Aggchainecdsamultisig.Contract.LastForceBatchSequenced(&_Aggchainecdsamultisig.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) NetworkName() (string, error) {
	return _Aggchainecdsamultisig.Contract.NetworkName(&_Aggchainecdsamultisig.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) NetworkName() (string, error) {
	return _Aggchainecdsamultisig.Contract.NetworkName(&_Aggchainecdsamultisig.CallOpts)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) OwnedAggchainVKeys(opts *bind.CallOpts, aggchainVKeySelector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "ownedAggchainVKeys", aggchainVKeySelector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.OwnedAggchainVKeys(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// OwnedAggchainVKeys is a free data retrieval call binding the contract method 0xeffb8479.
//
// Solidity: function ownedAggchainVKeys(bytes4 aggchainVKeySelector) view returns(bytes32 ownedAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) OwnedAggchainVKeys(aggchainVKeySelector [4]byte) ([32]byte, error) {
	return _Aggchainecdsamultisig.Contract.OwnedAggchainVKeys(&_Aggchainecdsamultisig.CallOpts, aggchainVKeySelector)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) PendingAdmin() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingAdmin(&_Aggchainecdsamultisig.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) PendingAdmin() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingAdmin(&_Aggchainecdsamultisig.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) PendingAggchainManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "pendingAggchainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingAggchainManager(&_Aggchainecdsamultisig.CallOpts)
}

// PendingAggchainManager is a free data retrieval call binding the contract method 0x527570f1.
//
// Solidity: function pendingAggchainManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) PendingAggchainManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.PendingAggchainManager(&_Aggchainecdsamultisig.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Pol() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Pol(&_Aggchainecdsamultisig.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Pol() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.Pol(&_Aggchainecdsamultisig.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) RollupManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.RollupManager(&_Aggchainecdsamultisig.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) RollupManager() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.RollupManager(&_Aggchainecdsamultisig.CallOpts)
}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) SignerToURLs(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "signerToURLs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Aggchainecdsamultisig.Contract.SignerToURLs(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// SignerToURLs is a free data retrieval call binding the contract method 0x36cd6b5b.
//
// Solidity: function signerToURLs(address ) view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) SignerToURLs(arg0 common.Address) (string, error) {
	return _Aggchainecdsamultisig.Contract.SignerToURLs(&_Aggchainecdsamultisig.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Threshold() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.Threshold(&_Aggchainecdsamultisig.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Threshold() (*big.Int, error) {
	return _Aggchainecdsamultisig.Contract.Threshold(&_Aggchainecdsamultisig.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.TrustedSequencer(&_Aggchainecdsamultisig.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) TrustedSequencer() (common.Address, error) {
	return _Aggchainecdsamultisig.Contract.TrustedSequencer(&_Aggchainecdsamultisig.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) TrustedSequencerURL() (string, error) {
	return _Aggchainecdsamultisig.Contract.TrustedSequencerURL(&_Aggchainecdsamultisig.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) TrustedSequencerURL() (string, error) {
	return _Aggchainecdsamultisig.Contract.TrustedSequencerURL(&_Aggchainecdsamultisig.CallOpts)
}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) UseDefaultSigners(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "useDefaultSigners")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UseDefaultSigners() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultSigners(&_Aggchainecdsamultisig.CallOpts)
}

// UseDefaultSigners is a free data retrieval call binding the contract method 0x188d9180.
//
// Solidity: function useDefaultSigners() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) UseDefaultSigners() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultSigners(&_Aggchainecdsamultisig.CallOpts)
}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) UseDefaultVkeys(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "useDefaultVkeys")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UseDefaultVkeys() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultVkeys(&_Aggchainecdsamultisig.CallOpts)
}

// UseDefaultVkeys is a free data retrieval call binding the contract method 0xfc5014d6.
//
// Solidity: function useDefaultVkeys() view returns(bool)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) UseDefaultVkeys() (bool, error) {
	return _Aggchainecdsamultisig.Contract.UseDefaultVkeys(&_Aggchainecdsamultisig.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggchainecdsamultisig.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Version() (string, error) {
	return _Aggchainecdsamultisig.Contract.Version(&_Aggchainecdsamultisig.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Aggchainecdsamultisig *AggchainecdsamultisigCallerSession) Version() (string, error) {
	return _Aggchainecdsamultisig.Contract.Version(&_Aggchainecdsamultisig.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptAdminRole(&_Aggchainecdsamultisig.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptAdminRole(&_Aggchainecdsamultisig.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AcceptAggchainManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "acceptAggchainManagerRole")
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptAggchainManagerRole(&_Aggchainecdsamultisig.TransactOpts)
}

// AcceptAggchainManagerRole is a paid mutator transaction binding the contract method 0x15981b29.
//
// Solidity: function acceptAggchainManagerRole() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AcceptAggchainManagerRole() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AcceptAggchainManagerRole(&_Aggchainecdsamultisig.TransactOpts)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) AddOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "addOwnedAggchainVKey", aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddOwnedAggchainVKey(&_Aggchainecdsamultisig.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// AddOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x19451a8f.
//
// Solidity: function addOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 newAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) AddOwnedAggchainVKey(aggchainVKeySelector [4]byte, newAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.AddOwnedAggchainVKey(&_Aggchainecdsamultisig.TransactOpts, aggchainVKeySelector, newAggchainVKey)
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) DisableUseDefaultSignersFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "disableUseDefaultSignersFlag")
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) DisableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultSignersFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// DisableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xfd7d2493.
//
// Solidity: function disableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) DisableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultSignersFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) DisableUseDefaultVkeysFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "disableUseDefaultVkeysFlag")
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) DisableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// DisableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0xefe6c9f4.
//
// Solidity: function disableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) DisableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.DisableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) EnableUseDefaultSignersFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "enableUseDefaultSignersFlag")
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) EnableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultSignersFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultSignersFlag is a paid mutator transaction binding the contract method 0xbe647d03.
//
// Solidity: function enableUseDefaultSignersFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) EnableUseDefaultSignersFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultSignersFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) EnableUseDefaultVkeysFlag(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "enableUseDefaultVkeysFlag")
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) EnableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// EnableUseDefaultVkeysFlag is a paid mutator transaction binding the contract method 0x1489e707.
//
// Solidity: function enableUseDefaultVkeysFlag() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) EnableUseDefaultVkeysFlag() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.EnableUseDefaultVkeysFlag(&_Aggchainecdsamultisig.TransactOpts)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) InitAggchainManager(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "initAggchainManager", newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.InitAggchainManager(&_Aggchainecdsamultisig.TransactOpts, newAggchainManager)
}

// InitAggchainManager is a paid mutator transaction binding the contract method 0xb3a326f7.
//
// Solidity: function initAggchainManager(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) InitAggchainManager(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.InitAggchainManager(&_Aggchainecdsamultisig.TransactOpts, newAggchainManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x697427f6.
//
// Solidity: function initialize(address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName, bool _useDefaultSigners, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "initialize", _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x697427f6.
//
// Solidity: function initialize(address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName, bool _useDefaultSigners, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) Initialize(_admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x697427f6.
//
// Solidity: function initialize(address _admin, address _trustedSequencer, address _gasTokenAddress, string _trustedSequencerURL, string _networkName, bool _useDefaultSigners, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) Initialize(_admin common.Address, _trustedSequencer common.Address, _gasTokenAddress common.Address, _trustedSequencerURL string, _networkName string, _useDefaultSigners bool, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.Initialize(&_Aggchainecdsamultisig.TransactOpts, _admin, _trustedSequencer, _gasTokenAddress, _trustedSequencerURL, _networkName, _useDefaultSigners, _signersToAdd, _newThreshold)
}

// MigrateFromPessimisticConsensus is a paid mutator transaction binding the contract method 0x87b2b559.
//
// Solidity: function migrateFromPessimisticConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) MigrateFromPessimisticConsensus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "migrateFromPessimisticConsensus")
}

// MigrateFromPessimisticConsensus is a paid mutator transaction binding the contract method 0x87b2b559.
//
// Solidity: function migrateFromPessimisticConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) MigrateFromPessimisticConsensus() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.MigrateFromPessimisticConsensus(&_Aggchainecdsamultisig.TransactOpts)
}

// MigrateFromPessimisticConsensus is a paid mutator transaction binding the contract method 0x87b2b559.
//
// Solidity: function migrateFromPessimisticConsensus() returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) MigrateFromPessimisticConsensus() (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.MigrateFromPessimisticConsensus(&_Aggchainecdsamultisig.TransactOpts)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) OnVerifyPessimistic(opts *bind.TransactOpts, aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "onVerifyPessimistic", aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.OnVerifyPessimistic(&_Aggchainecdsamultisig.TransactOpts, aggchainData)
}

// OnVerifyPessimistic is a paid mutator transaction binding the contract method 0x9ee4afa3.
//
// Solidity: function onVerifyPessimistic(bytes aggchainData) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) OnVerifyPessimistic(aggchainData []byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.OnVerifyPessimistic(&_Aggchainecdsamultisig.TransactOpts, aggchainData)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetTrustedSequencer(&_Aggchainecdsamultisig.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetTrustedSequencer(&_Aggchainecdsamultisig.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetTrustedSequencerURL(&_Aggchainecdsamultisig.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.SetTrustedSequencerURL(&_Aggchainecdsamultisig.TransactOpts, newTrustedSequencerURL)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferAdminRole(&_Aggchainecdsamultisig.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferAdminRole(&_Aggchainecdsamultisig.TransactOpts, newPendingAdmin)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) TransferAggchainManagerRole(opts *bind.TransactOpts, newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "transferAggchainManagerRole", newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferAggchainManagerRole(&_Aggchainecdsamultisig.TransactOpts, newAggchainManager)
}

// TransferAggchainManagerRole is a paid mutator transaction binding the contract method 0xbdfbed7e.
//
// Solidity: function transferAggchainManagerRole(address newAggchainManager) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) TransferAggchainManagerRole(newAggchainManager common.Address) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.TransferAggchainManagerRole(&_Aggchainecdsamultisig.TransactOpts, newAggchainManager)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) UpdateOwnedAggchainVKey(opts *bind.TransactOpts, aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "updateOwnedAggchainVKey", aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsamultisig.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateOwnedAggchainVKey is a paid mutator transaction binding the contract method 0x314eb17b.
//
// Solidity: function updateOwnedAggchainVKey(bytes4 aggchainVKeySelector, bytes32 updatedAggchainVKey) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) UpdateOwnedAggchainVKey(aggchainVKeySelector [4]byte, updatedAggchainVKey [32]byte) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateOwnedAggchainVKey(&_Aggchainecdsamultisig.TransactOpts, aggchainVKeySelector, updatedAggchainVKey)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactor) UpdateSignersAndThreshold(opts *bind.TransactOpts, _signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.contract.Transact(opts, "updateSignersAndThreshold", _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateSignersAndThreshold(&_Aggchainecdsamultisig.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// UpdateSignersAndThreshold is a paid mutator transaction binding the contract method 0xf51f563a.
//
// Solidity: function updateSignersAndThreshold((address,uint256)[] _signersToRemove, (address,string)[] _signersToAdd, uint256 _newThreshold) returns()
func (_Aggchainecdsamultisig *AggchainecdsamultisigTransactorSession) UpdateSignersAndThreshold(_signersToRemove []AggchainBaseRemoveSignerInfo, _signersToAdd []AggchainBaseSignerInfo, _newThreshold *big.Int) (*types.Transaction, error) {
	return _Aggchainecdsamultisig.Contract.UpdateSignersAndThreshold(&_Aggchainecdsamultisig.TransactOpts, _signersToRemove, _signersToAdd, _newThreshold)
}

// AggchainecdsamultisigAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptAdminRoleIterator struct {
	Event *AggchainecdsamultisigAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAcceptAdminRole)
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
		it.Event = new(AggchainecdsamultisigAcceptAdminRole)
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
func (it *AggchainecdsamultisigAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAcceptAdminRole represents a AcceptAdminRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*AggchainecdsamultisigAcceptAdminRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAcceptAdminRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAcceptAdminRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAcceptAdminRole(log types.Log) (*AggchainecdsamultisigAcceptAdminRole, error) {
	event := new(AggchainecdsamultisigAcceptAdminRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigAcceptAggchainManagerRoleIterator is returned from FilterAcceptAggchainManagerRole and is used to iterate over the raw logs and unpacked data for AcceptAggchainManagerRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptAggchainManagerRoleIterator struct {
	Event *AggchainecdsamultisigAcceptAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAcceptAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAcceptAggchainManagerRole)
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
		it.Event = new(AggchainecdsamultisigAcceptAggchainManagerRole)
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
func (it *AggchainecdsamultisigAcceptAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAcceptAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAcceptAggchainManagerRole represents a AcceptAggchainManagerRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptAggchainManagerRole struct {
	OldAggchainManager common.Address
	NewAggchainManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAcceptAggchainManagerRole is a free log retrieval operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAcceptAggchainManagerRole(opts *bind.FilterOpts) (*AggchainecdsamultisigAcceptAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAcceptAggchainManagerRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "AcceptAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAggchainManagerRole is a free log subscription operation binding the contract event 0x67c02ffba2f5329171ad235a360497af6ac3cfe82f1412866fbbf2dd3556ed3f.
//
// Solidity: event AcceptAggchainManagerRole(address oldAggchainManager, address newAggchainManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAcceptAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAcceptAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AcceptAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAcceptAggchainManagerRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAcceptAggchainManagerRole(log types.Log) (*AggchainecdsamultisigAcceptAggchainManagerRole, error) {
	event := new(AggchainecdsamultisigAcceptAggchainManagerRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigAcceptVKeyManagerRoleIterator is returned from FilterAcceptVKeyManagerRole and is used to iterate over the raw logs and unpacked data for AcceptVKeyManagerRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptVKeyManagerRoleIterator struct {
	Event *AggchainecdsamultisigAcceptVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAcceptVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAcceptVKeyManagerRole)
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
		it.Event = new(AggchainecdsamultisigAcceptVKeyManagerRole)
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
func (it *AggchainecdsamultisigAcceptVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAcceptVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAcceptVKeyManagerRole represents a AcceptVKeyManagerRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAcceptVKeyManagerRole struct {
	OldVKeyManager common.Address
	NewVKeyManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAcceptVKeyManagerRole is a free log retrieval operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAcceptVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsamultisigAcceptVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAcceptVKeyManagerRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "AcceptVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptVKeyManagerRole is a free log subscription operation binding the contract event 0xbb3b066505f14a628f4ba4187a046abd4dd17e96591d7a9ed31c91c79322ffe2.
//
// Solidity: event AcceptVKeyManagerRole(address oldVKeyManager, address newVKeyManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAcceptVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAcceptVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AcceptVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAcceptVKeyManagerRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAcceptVKeyManagerRole(log types.Log) (*AggchainecdsamultisigAcceptVKeyManagerRole, error) {
	event := new(AggchainecdsamultisigAcceptVKeyManagerRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AcceptVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigAddAggchainVKeyIterator is returned from FilterAddAggchainVKey and is used to iterate over the raw logs and unpacked data for AddAggchainVKey events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAddAggchainVKeyIterator struct {
	Event *AggchainecdsamultisigAddAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigAddAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigAddAggchainVKey)
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
		it.Event = new(AggchainecdsamultisigAddAggchainVKey)
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
func (it *AggchainecdsamultisigAddAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigAddAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigAddAggchainVKey represents a AddAggchainVKey event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigAddAggchainVKey struct {
	Selector        [4]byte
	NewAggchainVKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddAggchainVKey is a free log retrieval operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterAddAggchainVKey(opts *bind.FilterOpts) (*AggchainecdsamultisigAddAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigAddAggchainVKeyIterator{contract: _Aggchainecdsamultisig.contract, event: "AddAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchAddAggchainVKey is a free log subscription operation binding the contract event 0x6cd6ce07b60b06519523b9a97add34c2dcaa32dad22d44eb738554d81dfe2a79.
//
// Solidity: event AddAggchainVKey(bytes4 selector, bytes32 newAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchAddAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigAddAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "AddAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigAddAggchainVKey)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseAddAggchainVKey(log types.Log) (*AggchainecdsamultisigAddAggchainVKey, error) {
	event := new(AggchainecdsamultisigAddAggchainVKey)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "AddAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigDisableUseDefaultSignersFlagIterator is returned from FilterDisableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultSignersFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultSignersFlagIterator struct {
	Event *AggchainecdsamultisigDisableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigDisableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigDisableUseDefaultSignersFlag)
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
		it.Event = new(AggchainecdsamultisigDisableUseDefaultSignersFlag)
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
func (it *AggchainecdsamultisigDisableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigDisableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigDisableUseDefaultSignersFlag represents a DisableUseDefaultSignersFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterDisableUseDefaultSignersFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigDisableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigDisableUseDefaultSignersFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "DisableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x4c75580a56c734245a7418eb07d8a311e1bff79f982fed747da3589630e414be.
//
// Solidity: event DisableUseDefaultSignersFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchDisableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigDisableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "DisableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigDisableUseDefaultSignersFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseDisableUseDefaultSignersFlag(log types.Log) (*AggchainecdsamultisigDisableUseDefaultSignersFlag, error) {
	event := new(AggchainecdsamultisigDisableUseDefaultSignersFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator is returned from FilterDisableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for DisableUseDefaultVkeysFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator struct {
	Event *AggchainecdsamultisigDisableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigDisableUseDefaultVkeysFlag)
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
		it.Event = new(AggchainecdsamultisigDisableUseDefaultVkeysFlag)
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
func (it *AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigDisableUseDefaultVkeysFlag represents a DisableUseDefaultVkeysFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigDisableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterDisableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigDisableUseDefaultVkeysFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "DisableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchDisableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0x922aeecd3505b46415820aae489ed9dac9e250e74d497b14c33e8360b581ac07.
//
// Solidity: event DisableUseDefaultVkeysFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchDisableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigDisableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "DisableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigDisableUseDefaultVkeysFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseDisableUseDefaultVkeysFlag(log types.Log) (*AggchainecdsamultisigDisableUseDefaultVkeysFlag, error) {
	event := new(AggchainecdsamultisigDisableUseDefaultVkeysFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "DisableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigEnableUseDefaultSignersFlagIterator is returned from FilterEnableUseDefaultSignersFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultSignersFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultSignersFlagIterator struct {
	Event *AggchainecdsamultisigEnableUseDefaultSignersFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigEnableUseDefaultSignersFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigEnableUseDefaultSignersFlag)
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
		it.Event = new(AggchainecdsamultisigEnableUseDefaultSignersFlag)
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
func (it *AggchainecdsamultisigEnableUseDefaultSignersFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigEnableUseDefaultSignersFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigEnableUseDefaultSignersFlag represents a EnableUseDefaultSignersFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultSignersFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultSignersFlag is a free log retrieval operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterEnableUseDefaultSignersFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigEnableUseDefaultSignersFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigEnableUseDefaultSignersFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "EnableUseDefaultSignersFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultSignersFlag is a free log subscription operation binding the contract event 0x67ec953bdc8546ede08f8ee91e5205a1d1814e126cb8f5d00a918ddb1eaa292b.
//
// Solidity: event EnableUseDefaultSignersFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchEnableUseDefaultSignersFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigEnableUseDefaultSignersFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "EnableUseDefaultSignersFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigEnableUseDefaultSignersFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseEnableUseDefaultSignersFlag(log types.Log) (*AggchainecdsamultisigEnableUseDefaultSignersFlag, error) {
	event := new(AggchainecdsamultisigEnableUseDefaultSignersFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultSignersFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator is returned from FilterEnableUseDefaultVkeysFlag and is used to iterate over the raw logs and unpacked data for EnableUseDefaultVkeysFlag events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator struct {
	Event *AggchainecdsamultisigEnableUseDefaultVkeysFlag // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigEnableUseDefaultVkeysFlag)
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
		it.Event = new(AggchainecdsamultisigEnableUseDefaultVkeysFlag)
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
func (it *AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigEnableUseDefaultVkeysFlag represents a EnableUseDefaultVkeysFlag event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigEnableUseDefaultVkeysFlag struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEnableUseDefaultVkeysFlag is a free log retrieval operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterEnableUseDefaultVkeysFlag(opts *bind.FilterOpts) (*AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigEnableUseDefaultVkeysFlagIterator{contract: _Aggchainecdsamultisig.contract, event: "EnableUseDefaultVkeysFlag", logs: logs, sub: sub}, nil
}

// WatchEnableUseDefaultVkeysFlag is a free log subscription operation binding the contract event 0xaacf3fb6dd8daa3bebb71f5548f782413c3f3531625c6f9057c0f3d751b83829.
//
// Solidity: event EnableUseDefaultVkeysFlag()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchEnableUseDefaultVkeysFlag(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigEnableUseDefaultVkeysFlag) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "EnableUseDefaultVkeysFlag")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigEnableUseDefaultVkeysFlag)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseEnableUseDefaultVkeysFlag(log types.Log) (*AggchainecdsamultisigEnableUseDefaultVkeysFlag, error) {
	event := new(AggchainecdsamultisigEnableUseDefaultVkeysFlag)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "EnableUseDefaultVkeysFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigInitializedIterator struct {
	Event *AggchainecdsamultisigInitialized // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigInitialized)
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
		it.Event = new(AggchainecdsamultisigInitialized)
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
func (it *AggchainecdsamultisigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigInitialized represents a Initialized event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggchainecdsamultisigInitializedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigInitializedIterator{contract: _Aggchainecdsamultisig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigInitialized) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigInitialized)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseInitialized(log types.Log) (*AggchainecdsamultisigInitialized, error) {
	event := new(AggchainecdsamultisigInitialized)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator is returned from FilterOnVerifyPessimisticECDSAMultisig and is used to iterate over the raw logs and unpacked data for OnVerifyPessimisticECDSAMultisig events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator struct {
	Event *AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig)
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
		it.Event = new(AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig)
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
func (it *AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig represents a OnVerifyPessimisticECDSAMultisig event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOnVerifyPessimisticECDSAMultisig is a free log retrieval operation binding the contract event 0xc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8.
//
// Solidity: event OnVerifyPessimisticECDSAMultisig()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterOnVerifyPessimisticECDSAMultisig(opts *bind.FilterOpts) (*AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "OnVerifyPessimisticECDSAMultisig")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigOnVerifyPessimisticECDSAMultisigIterator{contract: _Aggchainecdsamultisig.contract, event: "OnVerifyPessimisticECDSAMultisig", logs: logs, sub: sub}, nil
}

// WatchOnVerifyPessimisticECDSAMultisig is a free log subscription operation binding the contract event 0xc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8.
//
// Solidity: event OnVerifyPessimisticECDSAMultisig()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchOnVerifyPessimisticECDSAMultisig(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "OnVerifyPessimisticECDSAMultisig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "OnVerifyPessimisticECDSAMultisig", log); err != nil {
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

// ParseOnVerifyPessimisticECDSAMultisig is a log parse operation binding the contract event 0xc118263690a4306c74bd1bc80b55962addc2d9e61619ac0b2c2883badbbd01d8.
//
// Solidity: event OnVerifyPessimisticECDSAMultisig()
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseOnVerifyPessimisticECDSAMultisig(log types.Log) (*AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig, error) {
	event := new(AggchainecdsamultisigOnVerifyPessimisticECDSAMultisig)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "OnVerifyPessimisticECDSAMultisig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetTrustedSequencerIterator struct {
	Event *AggchainecdsamultisigSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSetTrustedSequencer)
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
		it.Event = new(AggchainecdsamultisigSetTrustedSequencer)
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
func (it *AggchainecdsamultisigSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSetTrustedSequencer represents a SetTrustedSequencer event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*AggchainecdsamultisigSetTrustedSequencerIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSetTrustedSequencerIterator{contract: _Aggchainecdsamultisig.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSetTrustedSequencer)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSetTrustedSequencer(log types.Log) (*AggchainecdsamultisigSetTrustedSequencer, error) {
	event := new(AggchainecdsamultisigSetTrustedSequencer)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetTrustedSequencerURLIterator struct {
	Event *AggchainecdsamultisigSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSetTrustedSequencerURL)
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
		it.Event = new(AggchainecdsamultisigSetTrustedSequencerURL)
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
func (it *AggchainecdsamultisigSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*AggchainecdsamultisigSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSetTrustedSequencerURLIterator{contract: _Aggchainecdsamultisig.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSetTrustedSequencerURL)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSetTrustedSequencerURL(log types.Log) (*AggchainecdsamultisigSetTrustedSequencerURL, error) {
	event := new(AggchainecdsamultisigSetTrustedSequencerURL)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigSignersAndThresholdUpdatedIterator is returned from FilterSignersAndThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignersAndThresholdUpdated events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignersAndThresholdUpdatedIterator struct {
	Event *AggchainecdsamultisigSignersAndThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigSignersAndThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigSignersAndThresholdUpdated)
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
		it.Event = new(AggchainecdsamultisigSignersAndThresholdUpdated)
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
func (it *AggchainecdsamultisigSignersAndThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigSignersAndThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigSignersAndThresholdUpdated represents a SignersAndThresholdUpdated event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigSignersAndThresholdUpdated struct {
	AggchainSigners        []common.Address
	NewThreshold           *big.Int
	NewAggchainSignersHash [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSignersAndThresholdUpdated is a free log retrieval operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterSignersAndThresholdUpdated(opts *bind.FilterOpts) (*AggchainecdsamultisigSignersAndThresholdUpdatedIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigSignersAndThresholdUpdatedIterator{contract: _Aggchainecdsamultisig.contract, event: "SignersAndThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersAndThresholdUpdated is a free log subscription operation binding the contract event 0x66d7b0647fdd512b69cbf4f8e1ce8068bfe0b236168e2704ba13b07425eaa743.
//
// Solidity: event SignersAndThresholdUpdated(address[] aggchainSigners, uint256 newThreshold, bytes32 newAggchainSignersHash)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchSignersAndThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigSignersAndThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "SignersAndThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigSignersAndThresholdUpdated)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseSignersAndThresholdUpdated(log types.Log) (*AggchainecdsamultisigSignersAndThresholdUpdated, error) {
	event := new(AggchainecdsamultisigSignersAndThresholdUpdated)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "SignersAndThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferAdminRoleIterator struct {
	Event *AggchainecdsamultisigTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigTransferAdminRole)
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
		it.Event = new(AggchainecdsamultisigTransferAdminRole)
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
func (it *AggchainecdsamultisigTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigTransferAdminRole represents a TransferAdminRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*AggchainecdsamultisigTransferAdminRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigTransferAdminRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigTransferAdminRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseTransferAdminRole(log types.Log) (*AggchainecdsamultisigTransferAdminRole, error) {
	event := new(AggchainecdsamultisigTransferAdminRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigTransferAggchainManagerRoleIterator is returned from FilterTransferAggchainManagerRole and is used to iterate over the raw logs and unpacked data for TransferAggchainManagerRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferAggchainManagerRoleIterator struct {
	Event *AggchainecdsamultisigTransferAggchainManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigTransferAggchainManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigTransferAggchainManagerRole)
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
		it.Event = new(AggchainecdsamultisigTransferAggchainManagerRole)
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
func (it *AggchainecdsamultisigTransferAggchainManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigTransferAggchainManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigTransferAggchainManagerRole represents a TransferAggchainManagerRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferAggchainManagerRole struct {
	CurrentAggchainManager    common.Address
	NewPendingAggchainManager common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterTransferAggchainManagerRole is a free log retrieval operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterTransferAggchainManagerRole(opts *bind.FilterOpts) (*AggchainecdsamultisigTransferAggchainManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigTransferAggchainManagerRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "TransferAggchainManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferAggchainManagerRole is a free log subscription operation binding the contract event 0xa3d8e5d045432398be30f83ce7c35a7bfc220c1b66cc5bf3f4dd4d539d93fab6.
//
// Solidity: event TransferAggchainManagerRole(address currentAggchainManager, address newPendingAggchainManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchTransferAggchainManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigTransferAggchainManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "TransferAggchainManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigTransferAggchainManagerRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseTransferAggchainManagerRole(log types.Log) (*AggchainecdsamultisigTransferAggchainManagerRole, error) {
	event := new(AggchainecdsamultisigTransferAggchainManagerRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferAggchainManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigTransferVKeyManagerRoleIterator is returned from FilterTransferVKeyManagerRole and is used to iterate over the raw logs and unpacked data for TransferVKeyManagerRole events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferVKeyManagerRoleIterator struct {
	Event *AggchainecdsamultisigTransferVKeyManagerRole // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigTransferVKeyManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigTransferVKeyManagerRole)
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
		it.Event = new(AggchainecdsamultisigTransferVKeyManagerRole)
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
func (it *AggchainecdsamultisigTransferVKeyManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigTransferVKeyManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigTransferVKeyManagerRole represents a TransferVKeyManagerRole event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigTransferVKeyManagerRole struct {
	CurrentVKeyManager    common.Address
	NewPendingVKeyManager common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTransferVKeyManagerRole is a free log retrieval operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterTransferVKeyManagerRole(opts *bind.FilterOpts) (*AggchainecdsamultisigTransferVKeyManagerRoleIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigTransferVKeyManagerRoleIterator{contract: _Aggchainecdsamultisig.contract, event: "TransferVKeyManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferVKeyManagerRole is a free log subscription operation binding the contract event 0xc54ae01017d0b80bd8af833f66387d6eb547dc16c8206faf13d0b72764aab8b2.
//
// Solidity: event TransferVKeyManagerRole(address currentVKeyManager, address newPendingVKeyManager)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchTransferVKeyManagerRole(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigTransferVKeyManagerRole) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "TransferVKeyManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigTransferVKeyManagerRole)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseTransferVKeyManagerRole(log types.Log) (*AggchainecdsamultisigTransferVKeyManagerRole, error) {
	event := new(AggchainecdsamultisigTransferVKeyManagerRole)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "TransferVKeyManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggchainecdsamultisigUpdateAggchainVKeyIterator is returned from FilterUpdateAggchainVKey and is used to iterate over the raw logs and unpacked data for UpdateAggchainVKey events raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigUpdateAggchainVKeyIterator struct {
	Event *AggchainecdsamultisigUpdateAggchainVKey // Event containing the contract specifics and raw log

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
func (it *AggchainecdsamultisigUpdateAggchainVKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggchainecdsamultisigUpdateAggchainVKey)
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
		it.Event = new(AggchainecdsamultisigUpdateAggchainVKey)
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
func (it *AggchainecdsamultisigUpdateAggchainVKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggchainecdsamultisigUpdateAggchainVKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggchainecdsamultisigUpdateAggchainVKey represents a UpdateAggchainVKey event raised by the Aggchainecdsamultisig contract.
type AggchainecdsamultisigUpdateAggchainVKey struct {
	Selector             [4]byte
	PreviousAggchainVKey [32]byte
	NewAggchainVKey      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateAggchainVKey is a free log retrieval operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) FilterUpdateAggchainVKey(opts *bind.FilterOpts) (*AggchainecdsamultisigUpdateAggchainVKeyIterator, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.FilterLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return &AggchainecdsamultisigUpdateAggchainVKeyIterator{contract: _Aggchainecdsamultisig.contract, event: "UpdateAggchainVKey", logs: logs, sub: sub}, nil
}

// WatchUpdateAggchainVKey is a free log subscription operation binding the contract event 0x0aa5f73c189fb0b0a7cc98ae5fa89dfc16595480396208483518178435ed5b4f.
//
// Solidity: event UpdateAggchainVKey(bytes4 selector, bytes32 previousAggchainVKey, bytes32 newAggchainVKey)
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) WatchUpdateAggchainVKey(opts *bind.WatchOpts, sink chan<- *AggchainecdsamultisigUpdateAggchainVKey) (event.Subscription, error) {

	logs, sub, err := _Aggchainecdsamultisig.contract.WatchLogs(opts, "UpdateAggchainVKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggchainecdsamultisigUpdateAggchainVKey)
				if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
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
func (_Aggchainecdsamultisig *AggchainecdsamultisigFilterer) ParseUpdateAggchainVKey(log types.Log) (*AggchainecdsamultisigUpdateAggchainVKey, error) {
	event := new(AggchainecdsamultisigUpdateAggchainVKey)
	if err := _Aggchainecdsamultisig.contract.UnpackLog(event, "UpdateAggchainVKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
